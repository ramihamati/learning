package main

import (
	"errors"
	"log"
	"sync"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

type NetAddr struct {
	string
}

func (n NetAddr) Equals(addr NetAddr) bool {
	return n.string == addr.string
}

type RPC struct {
	Payload []byte
}

type Connection struct {
	peer      *Peer
	transport Transport
}

func NewConnection(peer *Peer, transport Transport) *Connection {
	return &Connection{peer: peer, transport: transport}
}

type Peer struct {
	lock           sync.RWMutex
	connectedPeers []*Connection
	Address        NetAddr
}

func NewPeer(address NetAddr) *Peer {
	return &Peer{
		lock:           sync.RWMutex{},
		connectedPeers: make([]*Connection, 0),
		Address:        address,
	}
}

func (p *Peer) AddPeer(peer *Peer, transport Transport) {
	p.lock.Lock()
	p.connectedPeers = append(peer.connectedPeers, NewConnection(
		peer,
		transport))
	p.lock.Unlock()
}

func (p *Peer) StartListening(peer *Peer) error {
	for i := range p.connectedPeers {
		if p.connectedPeers[i].peer.Address.Equals(peer.Address) {
			p.connectedPeers[i].transport.Subscribe(peer, func(rpc RPC) {
				log.Println(string(rpc.Payload))
			})
			return nil
		}
	}
	return errors.New("peer not connected")
}

func (p *Peer) Send(peer *Peer, payload RPC) error {

	for i := range p.connectedPeers {
		if p.connectedPeers[i].peer.Address.Equals(peer.Address) {
			p.connectedPeers[i].transport.Send(peer, payload)
			return nil
		}
	}
	return errors.New("peer not found")
}

func (p *Peer) Broadcast(payload RPC) {
	for i := range p.connectedPeers {
		p.connectedPeers[i].transport.Send(p.connectedPeers[i].peer, payload)
	}
}

type Transport interface {
	Send(*Peer, RPC)
	Subscribe(*Peer, func(rpc RPC))
}

type LocalTransport struct {
	Sender   chan<- RPC
	Receiver <-chan RPC
}

func NewLocalTransport() *LocalTransport {
	channel := make(chan RPC)
	return &LocalTransport{
		Sender:   channel,
		Receiver: channel,
	}
}

func (t *LocalTransport) Send(peer *Peer, payload RPC) {
	t.Sender <- payload
}

func (t *LocalTransport) Subscribe(peer *Peer, handler func(rpc RPC)) {
	// Should create a subscription for disposal probably and return
	// it in current's function response
	go (func() {
		for msg := range t.Receiver {
			handler(msg)
		}
	})()
}

func main() {
	//fmt.Print("test")
	log.Print("test")
	peer1 := NewPeer(NetAddr{"first"})
	peer2 := NewPeer(NetAddr{"second"})
	transport := NewLocalTransport()

	peer1.AddPeer(peer2, transport)
	peer2.AddPeer(peer1, transport)

	if err1 := peer1.StartListening(peer2); err1 != nil {
		panic(err1)
	}
	if err1 := peer2.StartListening(peer1); err1 != nil {
		panic(err1)
	}

	peer1.Broadcast(RPC{[]byte("hello")})

	//time.Sleep(30 * time.Second)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

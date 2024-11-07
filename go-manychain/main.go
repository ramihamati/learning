package main

import (
	"awesomeProject/network"
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

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
	Address        network.NetAddr
}

func NewPeer(address network.NetAddr) *Peer {
	return &Peer{
		lock:           sync.RWMutex{},
		connectedPeers: make([]*Connection, 0),
		Address:        address,
	}
}

func (p *Peer) RegisterPeer(peer *Peer, transport Transport) {
	p.lock.Lock()
	p.connectedPeers = append(peer.connectedPeers, NewConnection(
		peer,
		transport))
	p.lock.Unlock()
}

func (p *Peer) StartListening(peer *Peer) (ITransportListener, error) {
	for i := range p.connectedPeers {
		if p.connectedPeers[i].peer.Address.Equals(peer.Address) {
			return p.connectedPeers[i].transport.Subscribe(peer, func(rpc network.RPC) {
				log.Println(string(rpc.Payload))
			}), nil
		}
	}
	return nil, errors.New("peer not connected")
}

func (p *Peer) Send(peer *Peer, payload network.RPC) error {

	for i := range p.connectedPeers {
		if p.connectedPeers[i].peer.Address.Equals(peer.Address) {
			p.connectedPeers[i].transport.Send(peer, payload)
			return nil
		}
	}
	return errors.New("peer not found")
}

func (p *Peer) Broadcast(payload network.RPC) {
	for i := range p.connectedPeers {
		p.connectedPeers[i].transport.Send(p.connectedPeers[i].peer, payload)
	}
}

type ITransportListener interface {
	Close()
}

type LocalTransportListener[T network.RPC] struct {
	ctx       context.Context
	cancel    context.CancelFunc
	handler   func(rpc network.RPC)
	receiver  <-chan network.RPC
	lock      sync.RWMutex
	isStarted bool
}

func NewTransportListener[T network.RPC](handler func(rpc network.RPC), receiver <-chan network.RPC) LocalTransportListener[T] {
	ctx, cancel := context.WithCancel(context.Background())
	return LocalTransportListener[T]{
		ctx:       ctx,
		cancel:    cancel,
		handler:   handler,
		receiver:  receiver,
		lock:      sync.RWMutex{},
		isStarted: false,
	}
}

func (t *LocalTransportListener[T]) Start() {
	t.lock.Lock()

	if t.isStarted {
		log.Print("LocalTransportListener already started")
		t.lock.Unlock()
		return
	}
	t.isStarted = true
	log.Print("LocalTransportListener starting")
	go (func() {
		for {
			select {
			case <-t.ctx.Done():
				log.Println("LocalTransportListener closing")
				return
			case msg := <-t.receiver:
				log.Println("LocalTransportListener processing")
				// should we wrap this in a go also to alow for cancellations to occur?
				go t.handler(msg)
			}
		}
		//for msg := range t.receiver {
		//	t.handler(msg)
		//}
	})()
	t.lock.Unlock()
}

func (t *LocalTransportListener[T]) Close() {
	t.cancel()
}

type Transport interface {
	Send(*Peer, network.RPC)
	Subscribe(*Peer, func(rpc network.RPC)) ITransportListener
}

type LocalTransport struct {
	Sender   chan<- network.RPC
	Receiver <-chan network.RPC
}

func NewLocalTransport() *LocalTransport {
	channel := make(chan network.RPC)
	return &LocalTransport{
		Sender:   channel,
		Receiver: channel,
	}
}

func (t *LocalTransport) Send(peer *Peer, payload network.RPC) {
	t.Sender <- payload
}

func (t *LocalTransport) Subscribe(peer *Peer, handler func(rpc network.RPC)) ITransportListener {
	// Should create a subscription for disposal probably and return
	// it in current's function response
	listener := NewTransportListener(handler, t.Receiver)
	listener.Start()
	return &listener
}

func main() {
	//fmt.Print("test")
	log.Print("test")
	peer1 := NewPeer(network.NetAddr{Value: "first"})
	peer2 := NewPeer(network.NetAddr{Value: "second"})
	transport := NewLocalTransport()

	peer1.RegisterPeer(peer2, transport)
	peer2.RegisterPeer(peer1, transport)

	listener1, err1 := peer1.StartListening(peer2)
	if err1 != nil {
		panic(err1)
	}

	listener2, err1 := peer2.StartListening(peer1)
	if err1 != nil {
		panic(err1)
	}

	peer1.Broadcast(network.RPC{Payload: []byte("hello1")})
	peer1.Broadcast(network.RPC{Payload: []byte("hello2")})
	listener2.Close()
	listener1.Close()
	time.Sleep(1 * time.Second)

	peer1.Broadcast(network.RPC{[]byte("hello3")})
	peer1.Broadcast(network.RPC{[]byte("hello4")})

	time.Sleep(30 * time.Second)
	//time.Sleep(30 * time.Second)
	listener1.Close()
}

// TODO: add a ping between peers
//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

package network

import (
	"errors"
	"log"
	"sync"
)

type Peer struct {
	lock           sync.RWMutex
	Address        NetAddr
	Transport      ITransport
	connectedPeers []*RemotePeer
}

func NewPeer(address NetAddr, transport ITransport) *Peer {
	peer := &Peer{
		lock:           sync.RWMutex{},
		connectedPeers: make([]*RemotePeer, 0),
		Address:        address,
		Transport:      transport,
	}

	transport.Receive(func(rpc RPC) {
		log.Printf("Peer {%s} - Received {%s}\n", address.Value, rpc)
	})

	return peer
}

func (p *Peer) RegisterPeer(address NetAddr, transport ITransport) *RemotePeer {
	p.lock.Lock()
	remotePeer := &RemotePeer{
		Address:   address,
		Transport: transport,
	}
	p.connectedPeers = append(p.connectedPeers, remotePeer)
	p.lock.Unlock()
	return remotePeer
}

func (p *Peer) Send(peer *RemotePeer, payload RPC) error {

	for i := range p.connectedPeers {
		if p.connectedPeers[i].Address.Equals(peer.Address) {
			p.connectedPeers[i].Transport.Send(payload)
			return nil
		}
	}
	return errors.New("peer not found")
}

func (p *Peer) Broadcast(payload RPC) {
	for i := range p.connectedPeers {
		p.connectedPeers[i].Transport.Send(payload)
	}
}

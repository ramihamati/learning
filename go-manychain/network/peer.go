package network

import (
	"errors"
	"sync"
)

type Peer struct {
	lock           sync.RWMutex
	Address        NetAddr
	Transport      ITransport
	connectedPeers []*RemotePeer
}

func NewPeer(address NetAddr, transport ITransport) *Peer {
	return &Peer{
		lock:           sync.RWMutex{},
		connectedPeers: make([]*RemotePeer, 0),
		Address:        address,
		Transport:      transport,
	}
}

func (p *Peer) RegisterPeer(address NetAddr, transport ITransport) {
	p.lock.Lock()
	p.connectedPeers = append(p.connectedPeers, &RemotePeer{
		Address:   address,
		Transport: transport,
	})
	p.lock.Unlock()
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

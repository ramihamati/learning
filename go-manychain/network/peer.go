package network

import (
	"errors"
	"sync"
)

type Peer struct {
	lock           sync.RWMutex
	Address        NetworkAddress
	connectedPeers []*RemotePeer
}

func NewPeer(address NetworkAddress) *Peer {
	peer := &Peer{
		lock:           sync.RWMutex{},
		connectedPeers: make([]*RemotePeer, 0),
		Address:        address,
	}

	//transport.Receive(func(rpc RPC) {
	//	log.Printf("Peer {%s} - Received {%s}\n", address.Value, rpc)
	//})

	return peer
}

func (p *Peer) RegisterPeer(node INetworkNode, transport INetworkTransport) *RemotePeer {
	p.lock.Lock()
	remotePeer := &RemotePeer{
		Node:      node,
		Transport: transport,
	}
	p.connectedPeers = append(p.connectedPeers, remotePeer)
	p.lock.Unlock()
	return remotePeer
}

func (p *Peer) Send(peer *RemotePeer, payload RPC) error {

	// sending only if it's a known and registered peer
	for i := range p.connectedPeers {
		if p.connectedPeers[i].Node.Addr().Equals(peer.Node.Addr()) {
			p.connectedPeers[i].Transport.Send(p.connectedPeers[i].Node, payload)
			return nil
		}
	}
	return errors.New("peer not found")
}

func (p *Peer) Broadcast(payload RPC) {
	for i := range p.connectedPeers {
		p.connectedPeers[i].Transport.Send(p.connectedPeers[i].Node, payload)
	}
}

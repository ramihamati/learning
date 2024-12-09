package network

import (
	"errors"
	"sync"
)

type NodeServer struct {
	lock           sync.RWMutex
	Address        NetworkAddress
	connectedPeers []*NodeRemote
}

func NewNodeServer(address NetworkAddress) *NodeServer {
	peer := &NodeServer{
		lock:           sync.RWMutex{},
		connectedPeers: make([]*NodeRemote, 0),
		Address:        address,
	}

	return peer
}

func (p *NodeServer) AddPeer(node INetworkNode, transport INetworkTransport) *NodeRemote {
	p.lock.Lock()
	remotePeer := NewNodeRemote(transport, node)
	p.connectedPeers = append(p.connectedPeers, remotePeer)
	p.lock.Unlock()
	return remotePeer
}

func (p *NodeServer) Send(peer *NodeRemote, payload RPC) error {

	// sending only if it's a known and registered peer
	for i := range p.connectedPeers {
		if p.connectedPeers[i].Node.Addr().Equals(peer.Node.Addr()) {
			p.connectedPeers[i].Transport.Send(p.connectedPeers[i].Node, payload)
			return nil
		}
	}
	return errors.New("peer not found")
}

func (p *NodeServer) Broadcast(payload RPC) {
	for i := range p.connectedPeers {
		p.connectedPeers[i].Transport.Send(p.connectedPeers[i].Node, payload)
	}
}

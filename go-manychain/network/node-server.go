package network

import (
	"sync"
)

type NodeServer struct {
	lock           sync.RWMutex
	Address        INetworkEndpoint
	connectedPeers []IConnection
}

func NewNodeServer(address INetworkEndpoint) *NodeServer {
	peer := &NodeServer{
		lock:           sync.RWMutex{},
		connectedPeers: make([]IConnection, 0),
		Address:        address,
	}

	return peer
}

func (p *NodeServer) AddPeer(node IConnection) {
	p.lock.Lock()
	p.connectedPeers = append(p.connectedPeers, node)
	p.lock.Unlock()
}

//func (p *NodeServer) Send(peer *NodeRemote, payload RPC) error {
//
//	// sending only if it's a known and registered peer
//	for i := range p.connectedPeers {
//		if p.connectedPeers[i].Node.Addr().Equals(peer.Node.Addr()) {
//			p.connectedPeers[i].Transport.Send(p.connectedPeers[i].Node, payload)
//			return nil
//		}
//	}
//	return errors.New("peer not found")
//}

func (p *NodeServer) Broadcast(payload RPC) {
	for i := range p.connectedPeers {
		(p.connectedPeers[i]).Send(payload)
	}
}

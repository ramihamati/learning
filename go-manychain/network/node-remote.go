package network

type NodeRemote struct {
	Node      INetworkNode
	Transport INetworkTransport
}

func NewNodeRemote(transport INetworkTransport, node INetworkNode) *NodeRemote {
	return &NodeRemote{
		Transport: transport,
		Node:      node,
	}
}

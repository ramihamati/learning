package network

type RemotePeer struct {
	Node      INetworkNode
	Transport INetworkTransport
}

func NewRemotePeer(transport INetworkTransport, node INetworkNode) *RemotePeer {
	return &RemotePeer{
		Transport: transport,
		Node:      node,
	}
}

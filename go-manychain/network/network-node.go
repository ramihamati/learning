package network

type INetworkNode interface {
	Addr() NetworkAddress
}

// Local Node

type LocalNode struct {
	_addr    NetworkAddress
	_channel chan RPC
	Sender   chan<- RPC
	Receiver <-chan RPC
	IsClosed bool
}

func (t LocalNode) Addr() NetworkAddress {
	return t._addr
}

func NewLocalNode(name string) *LocalNode {
	channel := make(chan RPC, 1024)
	return &LocalNode{
		Sender:   channel,
		Receiver: channel,
		_channel: channel,
		_addr: NetworkAddress{
			Value: name,
		},
	}
}

func (t LocalNode) Close() {
	t.IsClosed = true
	close(t._channel)
}

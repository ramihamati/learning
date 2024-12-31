package network

type IStream interface {
	Send(RPC)
	Receive(func(rpc RPC))
}

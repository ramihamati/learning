package network

type IConnection interface {
	Send(RPC)
	Receive(func(rpc RPC))
	Close()
}

package network

type ITransport interface {
	Send(RPC)
	Receive(func(rpc RPC))
}

type LocalTransport struct {
	Endpoint INetworkEndpoint
}

func NewLocalTransport() *LocalTransport {
	return &LocalTransport{
		Endpoint: NewLocalEndpoint(),
	}
}

func (t *LocalTransport) Send(payload RPC) {
	t.Endpoint.Write(payload)
}

func (t *LocalTransport) Receive(function func(RPC)) {
	t.Endpoint.Listen(function)
}

func (t *LocalTransport) Close() {
	t.Endpoint.Close()
}

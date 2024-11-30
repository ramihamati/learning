package network

type ITransport interface {
	Send(RPC)
}

type LocalTransport struct {
	Writer ITransportWriter
}

func NewLocalTransport() *LocalTransport {
	return &LocalTransport{
		Writer: NewLocalTransportWriter(),
	}
}

func (t *LocalTransport) Send(payload RPC) {
	t.Writer.Write(payload)
}

func (t *LocalTransport) Close() {
	t.Writer.Close()
}

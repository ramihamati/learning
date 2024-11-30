package network

type ITransportWriter interface {
	Write(payload RPC)
	Close()
}

type LocalTransportWriter struct {
	Sender chan<- RPC
}

func NewLocalTransportWriter() *LocalTransportWriter {
	return &LocalTransportWriter{
		Sender: make(chan RPC),
	}
}

func (t LocalTransportWriter) Write(payload RPC) {
	t.Sender <- payload
}

func (t LocalTransportWriter) Close() {
	close(t.Sender)
}

package network

import "time"

type IConnection interface {
	Send(RPC)
	Receive(func(rpc RPC))
}

type IStream interface {
	Send(RPC)
	Receive(func(rpc RPC))
}

type LocalStream struct {
	channel  chan RPC
	isClosed bool
}

func (t *LocalStream) Send(payload RPC) {
	t.channel <- payload
}

func (t *LocalStream) Receive(action func(RPC)) {
	// TODO: should be deferred closed
	go func() {
		for {
			if t.isClosed {
				return
			}
			select {
			case vr := <-t.channel:
				action(vr)
			default: // If none are ready currently, we end up here
				time.Sleep(time.Millisecond * 1)
				if t.isClosed {
					return
				}
			}
		}
	}()
}

func (t *LocalStream) Close() {
	if !t.isClosed {
		t.isClosed = true
		close(t.channel)
	}
}

func NewLocalStream() *LocalStream {
	return &LocalStream{
		channel:  make(chan RPC),
		isClosed: false,
	}
}

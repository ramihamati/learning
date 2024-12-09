package network

import (
	"time"
)

type INetworkEndpoint interface {
	Write(payload RPC)
	Listen(func(rpc RPC))
	Close()
}

type LocalNetworkEndpoint struct {
	_channel chan RPC
	Sender   chan<- RPC
	Receiver <-chan RPC
	IsClosed bool
}

func NewLocalEndpoint() *LocalNetworkEndpoint {
	channel := make(chan RPC, 1024)
	return &LocalNetworkEndpoint{
		Sender:   channel,
		Receiver: channel,
		_channel: channel,
	}
}

func (t LocalNetworkEndpoint) Close() {
	t.IsClosed = true
	close(t._channel)
}

func (t LocalNetworkEndpoint) Write(payload RPC) {
	t.Sender <- payload
}

func (t LocalNetworkEndpoint) Listen(action func(rpc RPC)) {
	go func() {
		for {
			if t.IsClosed {
				return
			}
			select {
			case vr := <-t.Receiver:
				action(vr)
			default: // If none are ready currently, we end up here
				time.Sleep(time.Millisecond * 1)
				if t.IsClosed {
					return
				}
			}
		}
	}()
}

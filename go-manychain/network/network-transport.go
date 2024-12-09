package network

import (
	"fmt"
	"time"
)

type INetworkTransport interface {
	Send(INetworkNode, RPC)
	Receive(INetworkNode, func(rpc RPC))
}

// Local Transport

type LocalNetworkTransport struct {
}

func NewLocalTransport() *LocalNetworkTransport {
	return &LocalNetworkTransport{}
}

func (t *LocalNetworkTransport) Send(node INetworkNode, payload RPC) {
	localNode, ok := node.(*LocalNode)
	if !ok {
		panic(fmt.Errorf("this transport works only on LocalNode"))
	}

	localNode.Sender <- payload
}

func (t *LocalNetworkTransport) Receive(node INetworkNode, action func(RPC)) {
	localNode, ok := node.(*LocalNode)
	if !ok {
		panic(fmt.Errorf("this transport works only on LocalNode"))
	}

	// TODO: should be deferred closed
	go func() {
		for {
			if localNode.IsClosed {
				return
			}
			select {
			case vr := <-localNode.Receiver:
				action(vr)
			default: // If none are ready currently, we end up here
				time.Sleep(time.Millisecond * 1)
				if localNode.IsClosed {
					return
				}
			}
		}
	}()
}

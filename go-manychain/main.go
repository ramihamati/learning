package main

import (
	"awesomeProject/network"
	"log"
	"time"
)

func main() {
	//fmt.Print("test")
	log.Print("test")

	peer1 := network.NewPeer(network.NetworkAddress{Value: "peer1"})

	transport := network.NewLocalTransport()
	node1 := network.NewLocalNode("local1")
	node2 := network.NewLocalNode("local2")
	node3 := network.NewLocalNode("local3")

	rp1 := peer1.RegisterPeer(node1, transport)
	peer1.RegisterPeer(node2, transport)
	peer1.RegisterPeer(node3, transport)

	transport.Receive(node1, func(rpc network.RPC) {
		log.Printf("Peer {%s} - Received {%s}\n", node1.Addr().Value, rpc)
	})
	transport.Receive(node2, func(rpc network.RPC) {
		log.Printf("Peer {%s} - Received {%s}\n", node2.Addr().Value, rpc)
	})
	transport.Receive(node3, func(rpc network.RPC) {
		log.Printf("Peer {%s} - Received {%s}\n", node3.Addr().Value, rpc)
	})

	peer1.Broadcast(network.RPC{Payload: []byte("hello1")})
	peer1.Send(rp1, network.RPC{Payload: []byte("hello - direct - 1")})
	time.Sleep(30 * time.Second)
	//time.Sleep(30 * time.Second)
}

// TODO: add a ping between peers
//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

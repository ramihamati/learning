package main

import (
	"log"
	"manychain/core"
	"manychain/network"
	"time"
)

func main() {
	//fmt.Print("test")
	log.Print("test")

	node := network.NewLocalStream()

	node.Receive(func(rpc network.RPC) {
		println("{}", rpc.Payload)
	})

	node.Send(network.RPC{
		Payload: []byte("test"),
	})
	node.Send(network.RPC{
		Payload: []byte("test"),
	})
	time.Sleep(30 * time.Second)
}

func TestAddress() {
	address := core.NewAddress("chain.address.rami")
	log.Printf("%x", address)
	log.Printf("%d", len(address.Value()))
}

func TestBlocks() {
	header := core.NewHeader(0, core.HashFromBytes(make([]byte, 0)), core.Timestamp(0), 0, 1)
	block := core.NewBlock(header)
	hash := block.Hash()
	log.Printf("{%x}", hash)
}

func TestNetwork() {
	peer1 := network.NewNodeServer(network.NetworkAddress{Value: "peer1"})
	transport := network.NewLocalTransport()
	node1 := network.NewLocalNode("local1")
	node2 := network.NewLocalNode("local2")
	node3 := network.NewLocalNode("local3")

	rp1 := peer1.AddPeer(node1, transport)
	peer1.AddPeer(node2, transport)
	peer1.AddPeer(node3, transport)

	transport.Receive(node1, func(rpc network.RPC) {
		log.Printf("NodeServer {%s} - Received {%s}\n", node1.Addr().Value, rpc)
	})
	transport.Receive(node2, func(rpc network.RPC) {
		log.Printf("NodeServer {%s} - Received {%s}\n", node2.Addr().Value, rpc)
	})
	transport.Receive(node3, func(rpc network.RPC) {
		log.Printf("NodeServer {%s} - Received {%s}\n", node3.Addr().Value, rpc)
	})

	peer1.Broadcast(network.RPC{Payload: []byte("hello1")})
	peer1.Send(rp1, network.RPC{Payload: []byte("hello - direct - 1")})
	time.Sleep(30 * time.Second)
}

// TODO: add a ping between peers
//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

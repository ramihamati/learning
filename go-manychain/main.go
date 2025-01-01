package main

import (
	"log"
	"manychain/core"
	"manychain/crypto"
	"manychain/network"
	"time"
)

func main() {
	//fmt.Print("test")
	log.Print("test")
	TestBlocks()
	time.Sleep(30 * time.Second)
}

func TestSignature() {
	bytes := ([]byte)("test")
	ds := crypto.NewDigitalSignatureKeys()
	signature, err := ds.Sign(bytes)
	if err != nil {
		panic(err)
	}
	println(signature)
	verified := signature.Verify(ds.PublicKey(), bytes)
	println(verified)
}

func TestStreams() {
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
}

func TestAddress() {
	address := core.NewAddress("chain.address.rami")
	log.Printf("%x", address)
	log.Printf("%d", len(address.Value()))
}

func TestBlocks() {
	header := core.NewHeader(0, core.HashFromBytes(make([]byte, 0)), core.Timestamp(0), 0, 1)
	block := core.NewBlock(header)
	validator := crypto.NewDigitalSignatureKeys()
	block.Sign(validator)

	hash := block.Hash()
	log.Printf("{%x}", hash)
}

func TestNetwork() {
	serverEp := network.NewLocalEndpoint("localserver")
	server := network.NewNodeServer(serverEp)

	node1 := network.NewLocalConnection(network.NewLocalEndpoint("node1"))
	node2 := network.NewLocalConnection(network.NewLocalEndpoint("node2"))
	node3 := network.NewLocalConnection(network.NewLocalEndpoint("node3"))

	server.AddPeer(node1)
	server.AddPeer(node2)
	server.AddPeer(node3)

	node1.Receive(func(rpc network.RPC) {
		log.Printf("NodeServer {%s} - Received {%s}\n", node1.Name(), rpc)
	})
	node2.Receive(func(rpc network.RPC) {
		log.Printf("NodeServer {%s} - Received {%s}\n", node2.Name(), rpc)
	})
	node3.Receive(func(rpc network.RPC) {
		log.Printf("NodeServer {%s} - Received {%s}\n", node3.Name(), rpc)
	})

	server.Broadcast(network.RPC{Payload: []byte("hello1")})
	server.Broadcast(network.RPC{Payload: []byte("hello2")})
	//server.Send(rp1, network.RPC{Payload: []byte("hello - direct - 1")})
	time.Sleep(30 * time.Second)
}

// TODO: add a ping between peers
//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

package main

import (
	"awesomeProject/network"
	"log"
	"time"
)

func main() {
	//fmt.Print("test")
	log.Print("test")

	transport1 := network.NewLocalTransport()
	peer1 := network.NewPeer(network.NetAddr{Value: "first"}, transport1)

	transport2 := network.NewLocalTransport()
	peer2 := network.NewPeer(network.NetAddr{Value: "second"}, transport2)

	peer1.RegisterPeer(peer2.Address, transport2)
	peer2.RegisterPeer(peer1.Address, transport1)

	peer1.Broadcast(network.RPC{Payload: []byte("hello1")})
	peer2.Broadcast(network.RPC{Payload: []byte("hello2")})
	peer1.Broadcast(network.RPC{Payload: []byte("hello3")})
	peer2.Broadcast(network.RPC{Payload: []byte("hello4")})

	time.Sleep(30 * time.Second)
	//time.Sleep(30 * time.Second)
}

// TODO: add a ping between peers
//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

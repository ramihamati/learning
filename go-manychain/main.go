package main

import (
	"awesomeProject/network"
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	//fmt.Print("test")
	log.Print("test")

	transport1 := network.NewLocalTransport()
	peer1 := network.NewPeer(network.NetAddr{Value: "first"}, transport1)

	transport2 := network.NewLocalTransport()
	peer2 := network.NewPeer(network.NetAddr{Value: "second"}, transport2)

	// configure 2 as remote
	remotePeer2 := network.NewRemotePeer(transport1, peer2.Address)

	peer1.RegisterPeer(&remotePeer2, transport1)
	peer2.RegisterPeer(peer1, transport1)

	listener1, err1 := peer1.StartListening(peer2)
	if err1 != nil {
		panic(err1)
	}

	listener2, err1 := peer2.StartListening(peer1)
	if err1 != nil {
		panic(err1)
	}

	peer1.Broadcast(network.RPC{Payload: []byte("hello1")})
	peer1.Broadcast(network.RPC{Payload: []byte("hello2")})
	listener2.Close()
	listener1.Close()
	time.Sleep(1 * time.Second)

	peer1.Broadcast(network.RPC{[]byte("hello3")})
	peer1.Broadcast(network.RPC{[]byte("hello4")})

	time.Sleep(30 * time.Second)
	//time.Sleep(30 * time.Second)
	listener1.Close()
}

// TODO: add a ping between peers
//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.

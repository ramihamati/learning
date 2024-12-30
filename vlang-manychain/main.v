module main

import time
import network

fn main() {
	conn1 := network.LocalConnection.new(network.LocalEndpoint.new('local')) as network.LocalConnection
	conn1.consume(fn (message network.RPC) {
		eprintln('${message.payload.bytestr()}')
	})
	bytes := 'tet'.bytes()
	conn1.send_message(network.RPC{
		payload: bytes
	})
	conn1.send_message(network.RPC{
		payload: bytes
	})
	conn1.send_message(network.RPC{
		payload: bytes
	})
	println('Hello, World!')
	time.sleep(1 * time.second)
}

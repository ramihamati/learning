module network

pub interface IConnection {
	send_message(payload RPC)
	consume(fn (RPC))
}

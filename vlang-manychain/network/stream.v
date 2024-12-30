module network

interface IStream {
	send_message(RPC)
	consume(fn (RPC))
}

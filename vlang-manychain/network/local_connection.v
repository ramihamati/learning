module network

pub struct LocalConnection implements IConnection {
pub:
	endpoint LocalEndpoint
	stream   LocalStream
}

pub fn LocalConnection.new(endpoint LocalEndpoint) LocalConnection {
	return LocalConnection{
		endpoint: endpoint
		stream:   LocalStream.new()
	}
}

pub fn (conn &LocalConnection) send_message(payload RPC) {
	conn.stream.send_message(payload)
}

pub fn (conn &LocalConnection) consume(consumer fn (RPC)) {
	conn.stream.consume(consumer)
}

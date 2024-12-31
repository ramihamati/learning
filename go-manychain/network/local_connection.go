package network

// LocalConnection implements IConnection
type LocalConnection struct {
	endpoint LocalEndpoint
	stream   LocalStream
}

func NewLocalConnection(endpoint LocalEndpoint) *LocalConnection {
	return &LocalConnection{
		endpoint: endpoint,
		stream:   *NewLocalStream(),
	}
}

// Send is an implementation for IConnection
func (c *LocalConnection) Send(payload RPC) {
	c.stream.Send(payload)
}

// Receive is an implementation for IConnection
func (c *LocalConnection) Receive(action func(RPC)) {
	c.stream.Receive(action)
}

// Close is an implementation for IConnection
func (c *LocalConnection) Close() {
	c.stream.Close()
}

func (c *LocalConnection) Name() string {
	return c.endpoint.Name()
}

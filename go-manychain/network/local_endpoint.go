package network

// LocalEndpoint implements INetworkEndpoint
type LocalEndpoint struct {
	name string
}

func (obj LocalEndpoint) Name() string {
	return obj.name
}

func NewLocalEndpoint(name string) LocalEndpoint {
	return LocalEndpoint{
		name: name,
	}
}

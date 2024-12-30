module network

import net

pub struct LocalEndpoint implements INetworkEndpoint {
pub:
	name string
}

pub fn LocalEndpoint.new(name string) LocalEndpoint {
	return LocalEndpoint{
		name: name
	}
}

pub struct SocketNetworkEndpoint implements INetworkEndpoint {
pub:
	address net.Addr
	port    u16
}

pub fn SocketNetworkEndpoint.new(address net.Addr,
	port u16) SocketNetworkEndpoint {
	return SocketNetworkEndpoint{
		address: address
		port:    port
	}
}

module network

import net

@[noinit]
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

module network

import net

@[noinit]
pub struct LocalEndpoint implements INetworkEndpoint {
pub:
	name string
}

pub fn LocalEndpoint.new(name string) LocalEndpoint {
	return LocalEndpoint{
		name: name
	}
}

use std::hash::{Hash};
use std::net::IpAddr;
use crate::network::network_endpoint::INetworkEndpoint;

#[derive(Debug, Eq, Hash, PartialEq, Clone)]
pub struct SocketNetworkEndpoint {
    address: IpAddr,
    port: u16,
}

impl SocketNetworkEndpoint {
    pub fn new(address: IpAddr, port: u16) -> SocketNetworkEndpoint {
        SocketNetworkEndpoint { address, port }
    }
    fn formatted_address(self) -> String {
        format!("{}:{}", self.address, self.port)
    }
}

impl INetworkEndpoint for SocketNetworkEndpoint {
}

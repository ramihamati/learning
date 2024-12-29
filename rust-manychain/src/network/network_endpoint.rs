use std::hash::{Hash, Hasher};
use std::net::IpAddr;

pub trait INetworkEndpoint : Eq + Hash {
    fn formatted_address(self) -> String;
}

pub struct SocketNetworkEndpoint {
    address: IpAddr,
    port: u16,
}

impl SocketNetworkEndpoint {
    pub fn new(address: IpAddr, port: u16) -> SocketNetworkEndpoint {
        SocketNetworkEndpoint { address, port }
    }
}

impl Eq for SocketNetworkEndpoint {}

impl PartialEq<Self> for SocketNetworkEndpoint {
    fn eq(&self, other: &Self) -> bool {
        self.address == other.address && self.port == other.port
    }
}

impl Hash for SocketNetworkEndpoint {
    fn hash<H: Hasher>(&self, state: &mut H) {
        self.address.hash(state);
        self.port.hash(state);
    }
}

impl INetworkEndpoint for SocketNetworkEndpoint {
    fn formatted_address(self) -> String {
        format!("{}:{}", self.address, self.port)
    }
}
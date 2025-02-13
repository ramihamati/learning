use std::hash::{Hash};
use crate::network::endpoints::i_network_endpoint::INetworkEndpoint;

#[derive(Debug, Eq, Hash, PartialEq, Clone)]
pub struct LocalNetworkEndpoint {
    name: String,
}

impl LocalNetworkEndpoint {
    pub fn new(name: String) -> LocalNetworkEndpoint {
        LocalNetworkEndpoint { name }
    }
    fn formatted_address(self) -> String {
        return self.name;
    }
}
impl INetworkEndpoint for LocalNetworkEndpoint {}

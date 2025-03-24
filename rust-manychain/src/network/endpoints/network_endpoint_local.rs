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
   pub fn formatted_address<'a>(self) -> &'a String {
        return &self.name;
    }
}
impl INetworkEndpoint for LocalNetworkEndpoint {}

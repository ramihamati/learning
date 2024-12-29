use std::hash::{Hash, Hasher};
use crate::network::network_endpoint::INetworkEndpoint;

pub struct LocalNetworkEndpoint {
    name: String,
}

impl LocalNetworkEndpoint {
    pub fn new(name: String) -> LocalNetworkEndpoint {
        LocalNetworkEndpoint { name }
    }
}

impl Eq for LocalNetworkEndpoint {}

impl PartialEq<Self> for LocalNetworkEndpoint {
    fn eq(&self, other: &Self) -> bool {
        self.name.eq(&other.name)
    }
}

impl Hash for LocalNetworkEndpoint {
    fn hash<H: Hasher>(&self, state: &mut H) {
        // this was auto completed
        self.name.hash(state)
    }
}

impl INetworkEndpoint for LocalNetworkEndpoint {
    fn formatted_address(self) -> String {
        return self.name;
    }

}
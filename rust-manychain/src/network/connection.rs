use crate::network::network_endpoint::INetworkEndpoint;

pub trait IConnection {
    type T : INetworkEndpoint;
    async fn send_message(&self, payload: Vec<u8>);
    async fn consume(&mut self);
}
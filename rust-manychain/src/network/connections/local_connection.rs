use crate::network::connections::i_connection::IConnection;
use crate::network::endpoints::network_endpoint_local::LocalNetworkEndpoint;
use crate::network::stream::Stream;
use crate::network::streams::stream_local::LocalStream;

pub struct LocalConnection{
    endpoint: LocalNetworkEndpoint,
    stream: LocalStream
}

impl LocalConnection {
    pub fn new<'a>(endpoint: LocalNetworkEndpoint) ->  LocalConnection {
        LocalConnection{
            endpoint: endpoint.clone(),
            stream: LocalStream::new(endpoint.clone())
        }
    }
}

impl IConnection for LocalConnection {
    type T = LocalNetworkEndpoint;
    async fn send_message(&self, payload: Vec<u8>) {
        self.stream.send_message(payload).await;
    }
    async fn consume(&mut self) {
        self.stream.consume().await;
    }
}

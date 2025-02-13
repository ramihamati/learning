use std::future::IntoFuture;
use tokio::sync::{mpsc, RwLock};
use crate::network::endpoints::network_endpoint_local::LocalNetworkEndpoint;
use crate::network::streams::rpc::RPC;
use crate::network::streams::stream::Stream;

pub struct LocalStream{
    consume: mpsc::Receiver<RPC>,
    produce: mpsc::Sender<RPC>,
    lock : RwLock<()>,
    addr: LocalNetworkEndpoint,
}

impl LocalStream {
    pub fn new(addr: LocalNetworkEndpoint) -> LocalStream  {
        let (tx,  rx) = mpsc::channel::<RPC>(1024);
        LocalStream{
            consume : rx,
            produce : tx,
            lock : RwLock::new(()),
            addr
        }
    }
}
impl Stream for LocalStream {

    async fn send_message(&self, payload: Vec<u8>){
        self.produce.send(RPC::new(payload)).await.unwrap();
    }

    async fn consume(&mut self) {
        println!("local stream consumer thread starting");
        while let Some(message) = self.consume.recv().into_future().await {
            println!("Received")
        }
        println!("local stream consumer thread stoping");

    }
}


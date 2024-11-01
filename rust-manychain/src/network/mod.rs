mod rpc;

use std::collections::HashMap;
use std::sync::Arc;
use tokio::sync::{mpsc, RwLock};
use tokio::sync::mpsc::Receiver;
use std::thread;
#[derive(Debug, Eq, Hash, PartialEq)]
pub struct  NetAddr {
    addr: String,
}

impl NetAddr {
    fn clone(&self) -> Self {
        return self.clone();
    }
}

pub struct RPC{
    from : NetAddr,
    payload : Vec<u8>
}
pub trait Transport<T : Transport<T> + Send + Sync>{
    async fn connect(&mut self, transport: Box<T>);
    async fn start(&mut self);
    async fn broadcast(self, payload: Vec<u8>);
    async fn send_message(self, addr: NetAddr, payload: Vec<u8>);
    fn addr(self) -> Box<NetAddr>;
}
pub struct MyError{

}
pub struct LocalTransport{
    addr : NetAddr,
    consume: mpsc::Receiver<RPC>,
    produce: mpsc::Sender<RPC>,
    lock : RwLock<()>,
    peers: HashMap<NetAddr, Box<LocalTransport>>,
}

impl LocalTransport {
    pub fn new(addr: NetAddr) -> LocalTransport  {
        let (tx,  rx) = mpsc::channel::<RPC>(1024);

        return LocalTransport{
            addr,
            peers: HashMap::new(),
            consume : rx,
            produce : tx,
            lock : RwLock::new(()),
        }
    }
}

impl Transport<LocalTransport> for LocalTransport{
    async fn connect(&mut self, transport: Box<LocalTransport>)  {
        self.lock.write().await;
        self.peers.insert(transport.addr.clone(),  transport);
    }

    async fn start(&mut self){
        loop {
            match self.consume.recv().await {
                Some(_) => println!("Received"),
                None => {
                    println!("Channel closed, exiting loop.");
                    break;
                }
            }
        }
    }

  async  fn broadcast(self, payload: Vec<u8>) {
        for peer in self.peers {
            peer.1.produce.send(RPC {
                from: self.addr.clone(),
                payload: payload.clone()
            }).await.expect("TODO: panic message");
        }
    }

   async fn send_message(self, addr: NetAddr, payload: Vec<u8>){

        let peer = self.peers.get(&addr).unwrap();
        peer.produce.send(RPC{
            from: self.addr.clone(),
            payload: payload.clone()
        }).await.expect("failed to send message");
    }

    fn addr(self) -> Box<NetAddr> {
        todo!()
    }
}

pub fn Network(){
    let (tx, mut rx) = mpsc::channel::<NetAddr>(10);
}
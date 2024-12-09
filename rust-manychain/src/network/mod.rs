mod rpc;
use std::collections::{HashMap, HashSet};
use std::hash::{Hash, Hasher};
use std::ops::Deref;
use std::sync::Arc;
use tokio::sync::{mpsc, RwLock};

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
    payload : Vec<u8>
}
pub struct ConnectedPeer {
    addr: NetAddr,
    transport: Box<dyn Transport>,
}

impl ConnectedPeer {
    pub fn new(addr: NetAddr, transport: Box<dyn Transport>) -> ConnectedPeer {
        ConnectedPeer{
            transport,
            addr,
        }
    }
}

pub struct Peer {
    addr: NetAddr,
    lock : RwLock<()>,
    peers: HashSet<Box<ConnectedPeer>>,
}

impl Peer {
    pub fn new(addr : NetAddr) -> Peer {
        Peer{
            lock: RwLock::new(()),
            peers: HashSet::new(),
            addr,
        }
    }
    pub fn add_peer(&mut self, addr: NetAddr, transport: Box<dyn Transport>) {
        let connected = ConnectedPeer::new(addr, transport);
        self.peers.insert(Box::new(connected));
    }
}

pub trait Transport {
    async fn send_message(self, addr: NetAddr, payload: Vec<u8>);
    async fn consume(self);
}

pub struct LocalTransport{
    consume: mpsc::Receiver<RPC>,
    produce: mpsc::Sender<RPC>,
    lock : RwLock<()>,
}

impl LocalTransport {
    pub fn new(addr: NetAddr) -> LocalTransport  {
        let (tx,  rx) = mpsc::channel::<RPC>(1024);
        return LocalTransport{
            consume : rx,
            produce : tx,
            lock : RwLock::new(()),
        }
    }
}

impl Transport for LocalTransport{
   async fn send_message(self, addr: NetAddr, payload: Vec<u8>){
        self.produce.send(RPC{
            payload: payload.clone(),
        }).await.unwrap();
    }

    async fn consume(mut self) {
        loop{
            match  self.consume.recv().await {
                Some(_) => println!("Received"),
                None => {
                    println!("Channel closed, exiting loop.");
                    break;
                }
            }
        }
    }
}

pub fn Network(){
    let (tx, mut rx) = mpsc::channel::<NetAddr>(10);
}
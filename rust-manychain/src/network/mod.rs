mod rpc;
use std::collections::{HashMap, HashSet};
use std::hash::{Hash, Hasher};
use std::ops::Deref;
use std::rc::Rc;
use std::sync::Arc;
use async_trait::async_trait;
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
#[derive(Eq, Hash, PartialEq)]
pub struct ConnectedPeer<T : Transport> {
    addr: NetAddr,
    transport: Arc<T>,
}

impl<T: Transport> ConnectedPeer<T> {
    pub fn new(addr: NetAddr, transport: Arc<T>) -> ConnectedPeer <T>{
        ConnectedPeer{
            transport,
            addr,
        }
    }
}

pub struct Peer<T : Transport> {
    addr: NetAddr,
    lock : RwLock<()>,
    peers: HashSet<ConnectedPeer<T>>,
}

impl<T: Transport> Peer<T> {
    pub fn new(addr : NetAddr) -> Peer<T> {
        Peer{
            lock: RwLock::new(()),
            peers: HashSet::new(),
            addr,
        }
    }
    pub fn add_peer(&mut self, addr: NetAddr, transport: Arc<T>) {
        let connected = ConnectedPeer::new(addr, transport);
        self.peers.insert(connected);
    }
}
pub trait Transport : Eq +  Hash {
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

impl Eq for LocalTransport {}

impl PartialEq<Self> for LocalTransport {
    fn eq(&self, other: &Self) -> bool {
        todo!()
    }
}

impl Hash for LocalTransport {
    fn hash<H: Hasher>(&self, state: &mut H) {
        todo!()
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
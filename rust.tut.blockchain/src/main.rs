use std::collections::HashMap;
use std::net::SocketAddr;
use std::sync::mpsc::{Receiver,  SyncSender};
use std::error::Error;
use std::sync::{Arc, RwLock};
use std::sync::mpsc;

struct RPC{
    From : String,
    Payload : Vec<u8>
}

trait T_Connection{

    fn Listen(&self) -> &Receiver<RPC>;
    fn SendMessage(&self, bytes: &[u8]);
}

trait T_Node{
}

struct TcpNode {
    Addr : SocketAddr
}

struct LocalConnection{
    Sender : SyncSender<RPC>,
    Receiver : Receiver<RPC>
}

impl LocalConnection{
    fn New() -> LocalConnection{
        let (sender, receiver) = mpsc::sync_channel::<RPC>(1024);
        LocalConnection{
            Receiver : receiver,
            Sender : sender
        }
    }
}

impl T_Connection for LocalConnection {
    fn Listen(&self) -> &Receiver<RPC> {
         return &self.Receiver
    }
    fn SendMessage(&self, bytes: &[u8]) {

    }
}
trait Hub {
    fn AddPeer(&self, nodeLeft : dyn T_Node);
}

struct LocalHub{
    Peers :   HashMap<SocketAddr, Box<dyn Transport>>,
    ConsumeCh : Receiver<RPC>,
    SenderCh : SyncSender<RPC>,
    Lock : Arc<RwLock<()>>
}

impl LocalTransport {
    fn New(addr: SocketAddr) -> LocalTransport{
        let (sender, receiver) = mpsc::sync_channel(1024);

        LocalTransport{
            Addr : addr,
            Peers: HashMap::new(),
            ConsumeCh: receiver,
            SenderCh: sender,
            Lock: Arc::new(Default::default()),
        }
    }
}

impl Transport for LocalTransport{
    fn Consume(&self) -> &Receiver<RPC> {
        return &self.ConsumeCh;
    }

    fn SendMessage(&self, addr: SocketAddr, bytes: &[u8]) {
        todo!()
    }
    fn Addr(&self) -> SocketAddr {
        return self.Addr;
    }
    fn Connect(&mut self, other: &dyn Transport) -> Result<(), Box<dyn Error>> {
        self.Peers.insert(other.Addr(), Box::new(other));
        todo!()
    }
}

fn main() {
    println!("Hello, world!");
}

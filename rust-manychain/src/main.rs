mod block;
mod network;
mod transaction;

use {
    sha2::{Digest, Sha256},
    std::fmt,
};
use bincode::{config, Decode, Encode};
use std::sync::Arc;
use std::time::Duration;
use tokio::sync::Mutex;
use tokio::time::sleep;
use crate::network::connections::local_connection::LocalConnection;
use crate::network::endpoints::network_endpoint_local::LocalNetworkEndpoint;
use crate::network::i_connection::IConnection;

const HASH_BYTES: usize = 32;
pub struct Hash(pub [u8; HASH_BYTES]);

#[derive(Default)]
pub struct Hasher {
    hasher: Sha256,
}

impl Hasher {
    pub fn hash(&mut self, val: &[u8]) {
        self.hasher.update(val);
    }
    pub fn hashv(&mut self, vals: &[&[u8]]) {
        for val in vals {
            self.hash(val);
        }
    }
    pub fn result(self) -> Hash {
        Hash(self.hasher.finalize().into())
    }
}

impl fmt::Display for Hash {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", bs58::encode(self.0).into_string())
    }
}
#[tokio::main]
async fn main() {
    // let mut hasher = Hasher::default();
    // let vals = &["Gaggablaghblagh!".as_ref(), "flurbos".as_ref()];
    // hasher.hashv(vals);
    // let bhc = BlockHashContent {
    //     previous_hash: "test".to_string(),
    //     index: 1,
    //     transactions: vec!["test".to_string()],
    //     timestamp: 1,
    // };
    // let config = config::standard();
    // let a = bincode::encode_to_vec(bhc, config).unwrap();
    // let mut hasher2 = Hasher::default();
    // hasher2.hash(&a[..]);
    // println!("Hello, world! {}",  hasher.result());
    // println!("Hello, world! {}",  hasher2.result());

    let node1 = LocalNetworkEndpoint::new("node1".to_string());

    let node2 = LocalNetworkEndpoint::new("node2".to_string());
    let conn2 = LocalConnection::new(node1.clone());

    let conn1 = Arc::new(Mutex::new(LocalConnection::new(node1.clone())));

    tokio::join!(
        send_message(Arc::clone(&conn1), String::from("hey 1")),
        send_message(Arc::clone(&conn1), String::from("hey 2")),
        recv_message(Arc::clone(&conn1)));

    tokio::time::sleep(tokio::time::Duration::from_secs(1)).await;
    sleep(Duration::from_secs(10)).await;
}

async fn send_message(arc : Arc<Mutex<LocalConnection>>, message: String) {
    let mut conn1_lock = arc.lock().await;
    conn1_lock.send_message(Vec::from(message)).await;
}

async fn recv_message(arc : Arc<Mutex<LocalConnection>>) {
    let mut conn1_lock = arc.lock().await;
    conn1_lock.consume().await;
}
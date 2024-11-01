mod block;
mod network;

use {
    sha2::{Digest, Sha256},
    std::fmt,
};
use crate::block::BlockHashContent;
use bincode::{config, Decode, Encode};
use std::str;
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
fn main() {
    let mut hasher = Hasher::default();
    let vals = &["Gaggablaghblagh!".as_ref(), "flurbos".as_ref()];
    hasher.hashv(vals);
    let bhc = BlockHashContent {
        previous_hash: "test".to_string(),
        index: 1,
        transactions: vec!["test".to_string()],
        timestamp: 1,
    };
    let config = config::standard();
    let a = bincode::encode_to_vec(bhc, config).unwrap();
    let mut hasher2 = Hasher::default();
    hasher2.hash(&a[..]);
    println!("Hello, world! {}",  hasher.result());
    println!("Hello, world! {}",  hasher2.result());
}

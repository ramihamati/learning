use bincode::Encode;
use chrono::{Utc, DateTime};
use sha2::{Sha256, Digest};
use data_encoding::HEXLOWER;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone)]
pub struct Block {
    pub index : u64,
    pub timestamp : DateTime<Utc>,
    pub transactions: Vec<String>,
    pub previous_hash : String,
    pub hash : String,
}

#[derive(Serialize, Deserialize, Clone, Debug, Encode)]
pub struct BlockHashContent {
    pub index : u64,
    pub timestamp : i64,
    pub transactions: Vec<String>,
    pub previous_hash : String,
}


impl Block {
    pub fn new(
        index: u64,
        transactions: Vec<String>,
        timestamp: DateTime<Utc>,
        previous_hash : String) -> Block {

        let mut block = Block{
            index,
            timestamp,
            transactions,
            previous_hash,
            hash : String::from(""),
        };
        block.hash = block.calculate_hash();
        block
    }

    pub fn calculate_hash(&self) -> String {

        let bhc = BlockHashContent {
            previous_hash: self.previous_hash.clone(),
            index: self.index,
            transactions: self.transactions.clone(),
            timestamp: self.timestamp.timestamp(),
        };

        let serialized = serde_json::to_string(&bhc).expect("Failed to serialize object");

        let mut hasher = Sha256::new();
        hasher.update(serialized);

        HEXLOWER.encode(&hasher.finalize())
    }
}

#[cfg(test)]
mod tests{
    use super::*;

    #[test]
    fn hash_should_be_consistent(){

        let timestamp = DateTime::<Utc>::from_timestamp(10, 10).unwrap();

        let block = Block::new(
            1,
            vec![String::from("transaction1")],
            timestamp,
            String::from("previous_hash"));

        let hash = block.calculate_hash();

        assert_eq!(block.hash, hash);
        assert_eq!(block.calculate_hash(), "60edf88efa560ba99da5fe015164a002ac6f746eb3426e72805d03633254c0c5");
        assert_eq!(block.hash, "60edf88efa560ba99da5fe015164a002ac6f746eb3426e72805d03633254c0c5");
    }
}
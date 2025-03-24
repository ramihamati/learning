#[derive(Debug, Clone, PartialEq, Eq, Hash)]
pub struct RPC{
   pub payload : Vec<u8>,
}

impl RPC{
    pub fn new(payload : Vec<u8>) -> Self{
        RPC{payload}
    }
}
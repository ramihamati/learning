pub trait  Stream {
    async fn send_message(&self, payload: Vec<u8>);
    async fn consume(&mut self);
}

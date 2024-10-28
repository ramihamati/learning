impl :: bincode :: Encode for BlockHashContent
{
    fn encode < __E : :: bincode :: enc :: Encoder >
    (& self, encoder : & mut __E) ->core :: result :: Result < (), :: bincode
    :: error :: EncodeError >
    {
        :: bincode :: Encode :: encode(&self.index, encoder) ?; :: bincode ::
        Encode :: encode(&self.timestamp, encoder) ?; :: bincode :: Encode ::
        encode(&self.transactions, encoder) ?; :: bincode :: Encode ::
        encode(&self.previous_hash, encoder) ?; Ok(())
    }
}
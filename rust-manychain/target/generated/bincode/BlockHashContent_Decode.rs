impl :: bincode :: Decode for BlockHashContent
{
    fn decode < __D : :: bincode :: de :: Decoder > (decoder : & mut __D)
    ->core :: result :: Result < Self, :: bincode :: error :: DecodeError >
    {
        Ok(Self
        {
            index : :: bincode :: Decode :: decode(decoder) ?, timestamp : ::
            bincode :: Decode :: decode(decoder) ?, transactions : :: bincode
            :: Decode :: decode(decoder) ?, previous_hash : :: bincode ::
            Decode :: decode(decoder) ?,
        })
    }
} impl < '__de > :: bincode :: BorrowDecode < '__de > for BlockHashContent
{
    fn borrow_decode < __D : :: bincode :: de :: BorrowDecoder < '__de > >
    (decoder : & mut __D) ->core :: result :: Result < Self, :: bincode ::
    error :: DecodeError >
    {
        Ok(Self
        {
            index : :: bincode :: BorrowDecode :: borrow_decode(decoder) ?,
            timestamp : :: bincode :: BorrowDecode :: borrow_decode(decoder)
            ?, transactions : :: bincode :: BorrowDecode ::
            borrow_decode(decoder) ?, previous_hash : :: bincode ::
            BorrowDecode :: borrow_decode(decoder) ?,
        })
    }
}
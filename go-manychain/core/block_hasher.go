package core

import (
	"bytes"
)

type Hasher[T any] interface {
}

type BlockHasher struct {
}

func NewBlockHasher() BlockHasher {
	return BlockHasher{}
}

func (bh BlockHasher) Hash(bc *Block) Hash {
	// GOB commented implementation
	// GOB depends on exported fields and i don't like this approach
	// because it drives visibility to respect the third party desire instead of
	// the logic requirements

	//buf := &bytes.Buffer{}
	//enc := gob.NewEncoder(buf)
	//err := enc.Encode(bc)
	//if err != nil {
	//	panic(err)
	//}
	//return HashFromBytes(buf.Bytes())

	buffer := &bytes.Buffer{}

	if err := bc.Signature.EncodeBinary(buffer); err != nil {
		panic(err)
	}

	if err := bc.Validator.EncodeBinary(buffer); err != nil {
		panic(err)
	}

	if err := bc.Header.EncodeBinary(buffer); err != nil {
		panic(err)
	}

	for _, tx := range bc.Transactions {
		if err := tx.EncodeBinary(buffer); err != nil {
			panic(err)
		}
	}

	return HashFromBytes(buffer.Bytes())
}

//
//func (h *BasicTransaction) EncodeBinary(w io.Writer) error {
//	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
//		return err
//	}
//	if err := binary.Write(w, binary.LittleEndian, &h.Data); err != nil {
//		return err
//	}
//	return nil
//}

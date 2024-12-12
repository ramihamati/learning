package core

import (
	"bytes"
	"encoding/binary"
	"io"
)

func (h *BasicTransaction) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Data); err != nil {
		return err
	}
	return nil
}

type Block struct {
	Header       Header
	Transactions []Transaction
	cachedHash   *Hash
}

func (b Block) Hash() Hash {

	if b.cachedHash != nil {
		return *b.cachedHash
	}

	buffer := &bytes.Buffer{}
	if err := b.Header.EncodeBinary(buffer); err != nil {
		panic(err)
	}

	for _, tx := range b.Transactions {
		if err := tx.EncodeBinary(buffer); err != nil {
			panic(err)
		}
	}

	hash := HashFromBytes(buffer.Bytes())
	b.cachedHash = &hash
	return *b.cachedHash
}

func NewBlock(header Header) *Block {
	return &Block{
		Header:       header,
		Transactions: make([]Transaction, 0),
	}
}

func (b Block) Clone() Block {
	return Block{
		Header:       b.Header.Clone(),
		Transactions: b.Transactions,
	}
}

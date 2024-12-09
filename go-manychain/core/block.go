package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"io"
)

type Hash [32]byte
type Timestamp uint64

func HashFromBytes(data []byte) Hash {
	// Compute the SHA-256 hash
	return sha256.Sum256(data)
}

type Header struct {
	Version   uint32
	PrevHash  Hash
	Timestamp Timestamp
	Height    uint32
	Nonce     uint64
}

func NewHeader(version uint32, prevHash Hash, timestamp Timestamp, height uint32, noonce uint64) Header {
	return Header{
		Version:   version,
		PrevHash:  prevHash,
		Timestamp: timestamp,
		Height:    height,
		Nonce:     noonce,
	}
}

func (h *Header) Clone() Header {
	return Header{
		Version:   h.Version,
		PrevHash:  h.PrevHash,
		Timestamp: h.Timestamp,
		Height:    h.Height,
		Nonce:     h.Nonce,
	}
}

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.PrevHash); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, &h.Nonce)
}

type Transaction interface {
	EncodeBinary(w io.Writer) error
}

type BasicTransaction struct {
	Timestamp Timestamp
	Data      []byte
}

func NewBasicTransaction(
	timestamp Timestamp,
	data []byte) *BasicTransaction {
	return &BasicTransaction{
		Timestamp: timestamp,
		Data:      data,
	}
}

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

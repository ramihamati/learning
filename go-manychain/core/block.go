package core

import (
	"bytes"
	"manychain/crypto"
)

type Block struct {
	Header       Header
	Transactions []ITransaction
	// Storing the validator public key and signature
	Validator  *crypto.PublicKey
	Signature  *crypto.DigitalSignature
	cachedHash *Hash
}

func (b *Block) Sign(validator crypto.DigitalSignatureKeys) {
	pub := validator.PublicKey()
	b.Validator = &pub
	buffer := &bytes.Buffer{}
	for _, tx := range b.Transactions {
		tx.EncodeBinary(buffer)
	}

	signature, err := validator.Sign(buffer.Bytes())
	if err != nil {
		panic("Failed to sign block")
		panic(err)
	}
	b.Signature = signature
}

func (b Block) Hash() Hash {

	if b.cachedHash != nil {
		return *b.cachedHash
	}
	hash := BlockHasher{}.Hash(&b)
	b.cachedHash = &hash
	return *b.cachedHash
}

func NewBlock(header Header) *Block {
	return &Block{
		Header:       header,
		Transactions: make([]ITransaction, 0),
	}
}

func (b Block) Clone() Block {
	return Block{
		Header:       b.Header.Clone(),
		Transactions: b.Transactions,
	}
}

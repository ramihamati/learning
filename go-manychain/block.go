package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"time"
)

type Block struct {
	Index        uint64
	Timestamp    time.Time
	Transactions []string
	PreviousHash string
	Hash         string
}

type BlockHashContent struct {
	Index        uint64
	Timestamp    time.Time
	Transactions []string
	PreviousHash string
}

func NewBlock(index uint64, timestamp time.Time, transactions []string, previousHash string) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    timestamp,
		Transactions: transactions,
		PreviousHash: previousHash,
	}

	block.Hash = block.CalculateHash()
	return block
}

func (b *Block) CalculateHash() string {
	hashGenerator := sha256.New()
	blockHashContent := BlockHashContent{
		Index:        b.Index,
		Timestamp:    b.Timestamp,
		Transactions: b.Transactions,
		PreviousHash: b.PreviousHash,
	}

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(blockHashContent)
	if err != nil {
		panic(err)
	}
	hashGenerator.Write(buffer.Bytes())
	return hex.EncodeToString(hashGenerator.Sum(nil))
}

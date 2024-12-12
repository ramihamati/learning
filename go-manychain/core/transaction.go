package core

import "io"

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

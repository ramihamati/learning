package core

import (
	"encoding/binary"
	"io"
)

// TODO: transaction can have only one Id, a Signature and Data associated to it can
// be stored sepparaetlly
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

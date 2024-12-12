package core

import (
	"encoding/binary"
	"io"
)

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

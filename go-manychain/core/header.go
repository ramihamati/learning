package core

import (
	"encoding/binary"
	"io"
)

type Header struct {
	version   uint32
	prevHash  Hash
	timestamp Timestamp
	height    uint32
	nonce     uint64
}

func (h *Header) Version() uint32 {
	return h.version
}
func (h *Header) PrevHash() Hash {
	return h.prevHash
}
func (h *Header) Timestamp() Timestamp {
	return h.timestamp
}
func (h *Header) Nonce() uint64 {
	return h.nonce
}
func (h *Header) Height() uint32 {
	return h.height
}

func NewHeader(version uint32, prevHash Hash, timestamp Timestamp, height uint32, noOnce uint64) Header {
	return Header{
		version:   version,
		prevHash:  prevHash,
		timestamp: timestamp,
		height:    height,
		nonce:     noOnce,
	}
}

func (h *Header) Clone() Header {
	return Header{
		version:   h.version,
		prevHash:  h.prevHash,
		timestamp: h.timestamp,
		height:    h.height,
		nonce:     h.nonce,
	}
}

func (h *Header) EncodeBinary(w io.Writer) error {
	if err := binary.Write(w, binary.LittleEndian, &h.version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.prevHash); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.LittleEndian, &h.height); err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, &h.nonce)
}

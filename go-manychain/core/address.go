package core

import (
	"crypto/sha256"
	"manychain/crypto"
)

// TODO: address can be more complex then this I think. It can have smart data, it can be a signed object
type Address [20]uint8

func (a Address) ToSlice() []byte {
	b := make([]byte, 20)
	for i := 0; i < 20; i++ {
		b[i] = a[i]
	}
	return b
}

func NewAddressFromPublicKey(publicKey crypto.PublicKey) Address {
	bytes := publicKey.Bytes()
	h := sha256.Sum256(bytes)
	v := make([]byte, 20)
	for i := 0; i < 20; i++ {
		v[i] = h[i]
	}

	return Address(v)
}

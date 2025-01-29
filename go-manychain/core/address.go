package core

import (
	"crypto/sha256"
	"manychain/crypto"
)

// TODO: address can be more complex then this I think. It can have smart data, it can be a signed object
// TODO: can everything be an object with an address?

type Address struct {
	value [20]byte
}

func (a Address) Value() [20]byte {
	return a.value
}

func New() Address {
	ds := crypto.NewDigitalSignatureKeys()
	address := FromKey(ds.PublicKey())
	return address
}

func FromKey(publicKey crypto.PublicKey) Address {
	bytes := publicKey.Bytes()
	h := sha256.Sum256(bytes)
	var v [20]byte
	for i := 0; i < 20; i++ {
		v[i] = h[i]
	}

	return Address{
		value: v,
	}
}

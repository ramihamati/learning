package core

import (
	"crypto/sha256"
	"manychain/crypto"
)

// TODO: address can be more complex then this I think. It can have smart data, it can be a signed object
// TODO: can everything be an object with an address?

type Address struct {
	value [20]byte
	name  string
}

func (a Address) Value() [20]byte {
	return a.value
}

func (a Address) Name() string {
	return a.name
}

//func (a Address) ToSlice() Address {
//	b := make([]byte, 20)
//	for i := 0; i < 20; i++ {
//		b[i] = a.value[i]
//	}
//	return b
//}

func NewAddress(name string) Address {
	ds := crypto.NewDigitalSignatureKeys()
	address := newAddressFromPublicKey(ds.PublicKey(), name)
	return address
}

func newAddressFromPublicKey(publicKey crypto.PublicKey, name string) Address {
	bytes := publicKey.Bytes()
	h := sha256.Sum256(bytes)
	var v [20]byte
	for i := 0; i < 20; i++ {
		v[i] = h[i]
	}

	return Address{
		value: v,
		name:  name,
	}
}

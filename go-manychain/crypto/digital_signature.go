package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type DigitalSignature struct {
	publicKey  PublicKey
	privateKey PrivateKey
}

func NewDigitalSignature() DigitalSignature {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	privateKey := PrivateKey{
		key: key,
	}
	publicKey := PublicKey{
		key: &key.PublicKey,
	}

	return DigitalSignature{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (ds DigitalSignature) PublicKey() PublicKey {
	return ds.publicKey
}
func (ds DigitalSignature) PrivateKey() PrivateKey {
	return ds.privateKey
}

package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type DigitalSignatureKeys struct {
	publicKey  PublicKey
	privateKey PrivateKey
}

func NewDigitalSignatureKeys() DigitalSignatureKeys {
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

	return DigitalSignatureKeys{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (ds DigitalSignatureKeys) PublicKey() PublicKey {
	return ds.publicKey
}
func (ds DigitalSignatureKeys) PrivateKey() PrivateKey {
	return ds.privateKey
}

func (ds DigitalSignatureKeys) Sign(msg []byte) (*DigitalSignature, error) {
	r, s, error := ecdsa.Sign(rand.Reader, ds.privateKey.key, msg)

	if error != nil {
		return nil, error
	}

	return &DigitalSignature{
		r: r,
		s: s,
	}, nil
}

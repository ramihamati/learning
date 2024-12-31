package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type DigitalSignaturePair struct {
	publicKey  PublicKey
	privateKey PrivateKey
}

func NewDigitalSignature() DigitalSignaturePair {
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

	return DigitalSignaturePair{
		publicKey:  publicKey,
		privateKey: privateKey,
	}
}

func (ds DigitalSignaturePair) PublicKey() PublicKey {
	return ds.publicKey
}
func (ds DigitalSignaturePair) PrivateKey() PrivateKey {
	return ds.privateKey
}

func (ds DigitalSignaturePair) Sign(msg []byte) (*DigitalSignature, error) {
	sig, error := ecdsa.SignASN1(rand.Reader, ds.privateKey.key, msg)

	if error != nil {
		return nil, error
	}

	return &DigitalSignature{
		signature: sig,
	}, nil
}

type DigitalSignature struct {
	signature []byte
}

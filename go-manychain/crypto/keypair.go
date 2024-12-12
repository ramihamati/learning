package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (pub *PublicKey) Bytes() []byte {
	return elliptic.MarshalCompressed(elliptic.P256(), pub.key.X, pub.key.Y)
}

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func NewPrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		key: key,
	}
}

func NewPublicKey(key PrivateKey) PublicKey {
	return PublicKey{
		key: &key.key.PublicKey,
	}
}

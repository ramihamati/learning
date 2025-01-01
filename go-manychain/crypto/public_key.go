package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
)

type PublicKey struct {
	key *ecdsa.PublicKey
}

func (pub *PublicKey) Bytes() []byte {
	return elliptic.MarshalCompressed(elliptic.P256(), pub.key.X, pub.key.Y)
}

func (pub *PublicKey) Key() *ecdsa.PublicKey {
	return pub.key
}

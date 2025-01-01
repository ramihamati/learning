package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"io"
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

func (pub *PublicKey) EncodeBinary(w io.Writer) error {
	if pub.key.X == nil {
		panic("Invalid public key. X is nil. This can happen if the hash was requested on a non-signed block")
	}
	if pub.key.Y == nil {
		panic("Invalid public key. Y is nil. This can happen if the hash was requested on a non-signed block")
	}
	if _, err := w.Write(pub.key.X.Bytes()); err != nil {
		return err
	}
	if _, err := w.Write(pub.key.Y.Bytes()); err != nil {
		return err
	}
	return nil
}

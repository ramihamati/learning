package crypto

import (
	"crypto/ecdsa"
	"math/big"
)

type DigitalSignature struct {
	r *big.Int
	s *big.Int
}

func NewDigitalSignature(r *big.Int, s *big.Int) DigitalSignature {
	return DigitalSignature{
		r: r,
		s: s,
	}
}

func (ds DigitalSignature) Verify(publicKey PublicKey, data []byte) bool {
	return ecdsa.Verify(publicKey.Key(), data, ds.r, ds.s)
}

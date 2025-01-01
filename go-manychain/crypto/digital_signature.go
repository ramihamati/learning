package crypto

import (
	"crypto/ecdsa"
	"io"
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
func (ds DigitalSignature) EncodeBinary(w io.Writer) error {
	if ds.r == nil {
		panic("Invalid signature. R is nil. This can happen if the hash was requested on a non-signed block")
	}
	if ds.s == nil {
		panic("Invalid signature. S is nil. This can happen if the hash was requested on a non-signed block")
	}
	if _, err := w.Write(ds.r.Bytes()); err != nil {
		return err
	}
	if _, err := w.Write(ds.s.Bytes()); err != nil {
		return err
	}
	return nil
}

package crypto

import "crypto/ecdsa"

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

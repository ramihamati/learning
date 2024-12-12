package core

import "crypto/sha256"

type Hash [32]byte

func HashFromBytes(data []byte) Hash {
	// Compute the SHA-256 hash
	return sha256.Sum256(data)
}

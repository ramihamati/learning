package main

import (
	"testing"
	"time"
)

func TestHashShouldBeConsistent(t *testing.T) {
	timestamp := time.Unix(10, 10) // Equivalent to Rust's `from_timestamp`

	block := NewBlock(1, timestamp, []string{"transaction1"}, "previous_hash")

	expectedHash := "f5ca3dd92a39fa972eb914fbefa9f3e9c3f73d1508e2d92dbcfcdf4ae1385d34"

	if block.Hash != expectedHash {
		t.Errorf("Expected block hash to be %s, but got %s", expectedHash, block.Hash)
	}

	hash := block.CalculateHash()
	if hash != expectedHash {
		t.Errorf("Expected hash to be %s, but got %s", expectedHash, hash)
	}
}

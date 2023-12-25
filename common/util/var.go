package util

import (
	"crypto/rand"
	"log"
	"math/big"
)

// Parses a base 16 string and returns *big.Int.
func FromHex(hex string) (n *big.Int) {
	n, ok := new(big.Int).SetString(hex, 16)
	if !ok {
		log.Fatalf("fail to parse hex number")
	}
	return n
}

// Generates a random int within a specific order and returns the underlying
// sequence of bytes.
func RandomBytes(order *big.Int) []byte {
	random, err := rand.Int(rand.Reader, order)
	if err != nil {
		log.Fatalf("fail create random bytes: %v", err)
	}
	return random.Bytes()
}

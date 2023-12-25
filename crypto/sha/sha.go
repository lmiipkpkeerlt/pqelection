package sha

import (
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"log"
	"math/big"
)

// Generates a hash-based commitment to a key and returns the commitment and
// randomness used.
func Commit(key []byte, order *big.Int) (com []byte, r []byte) {
	random, err := rand.Int(rand.Reader, order)
	if err != nil {
		log.Fatalf("fail to get random int, field <= 0: %v", err)
	}
	r = random.Bytes()
	sha := sha512.New()
	_, err = sha.Write(append(key, r...))
	if err != nil {
		log.Fatalf("fail to compute sha-512: %v", err)
	}
	return sha.Sum(nil), r
}

// Verifies a commitment and returns a boolean indicating success/failure.
func VerifyCommit(com []byte, key []byte, r []byte) bool {
	sha := sha512.New()
	_, err := sha.Write(append(key, r...))
	if err != nil {
		log.Fatalf("fail to compute sha-512: %v", err)
	}
	return bytes.Equal(sha.Sum(nil), com)
}

// Hashes bytes and returns the digest.
func Hash(bytes []byte) []byte {
	sha := sha512.New()
	_, err := sha.Write(bytes)
	if err != nil {
		log.Fatalf("fail to compute sha-512: %v", err)
	}
	return sha.Sum(nil)
}

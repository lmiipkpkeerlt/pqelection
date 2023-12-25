package edsa

import (
	"crypto/ed25519"
	"log"
)

// Returns an ed25519 digital signature keypair.
func GenerateKeypair() (ed25519.PublicKey, ed25519.PrivateKey) {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Fatalf("fail to generate ed25519 keypair: %v", err)
	}
	return pub, priv
}

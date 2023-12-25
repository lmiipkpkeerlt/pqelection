// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package elgamal

import (
	"bsc/src/common/util"
	"bsc/src/config"
	"bytes"
	"crypto/rand"
	"math/big"
	"testing"
)

const (
	primeHex     = config.PRIME
	generatorHex = config.GENERATOR
)

func TestEncryptDecrypt(t *testing.T) {
	priv := &PrivateKey{
		PublicKey: PublicKey{
			G: util.FromHex(generatorHex),
			P: util.FromHex(primeHex),
		},
		X: util.FromHex("42"),
	}
	priv.Y = new(big.Int).Exp(priv.G, priv.X, priv.P)

	message := []byte("hello world")
	c1, c2, err := Encrypt(rand.Reader, &priv.PublicKey, message)
	if err != nil {
		t.Errorf("error encrypting: %s", err)
	}
	message2, err := Decrypt(priv, c1, c2)
	if err != nil {
		t.Errorf("error decrypting: %s", err)
	}
	if !bytes.Equal(message2, message) {
		t.Errorf("decryption failed, got: %x, want: %x", message2, message)
	}
}

func TestDecryptBadKey(t *testing.T) {
	priv := &PrivateKey{
		PublicKey: PublicKey{
			G: util.FromHex(generatorHex),
			P: util.FromHex("2"),
		},
		X: util.FromHex("42"),
	}
	priv.Y = new(big.Int).Exp(priv.G, priv.X, priv.P)
	c1, c2 := util.FromHex("8"), util.FromHex("8")
	if _, err := Decrypt(priv, c1, c2); err == nil {
		t.Errorf("unexpected success decrypting")
	}
}

// Package diffiehellman implements Diffie-Hellman key exchange
//
// This one is very easy in Go because the big.Int type supports
// all the necessary operations for calculating keys, and the
// crypto/rand package supports seeding the keys.
package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey generates a random private key such that 1 < key < p
func PrivateKey(p *big.Int) (key *big.Int) {
	upperLimit := (&big.Int{}).Sub(p, big.NewInt(2))
	key, _ = rand.Int(rand.Reader, upperLimit)
	key.Add(key, big.NewInt(2))
	return
}

// PublicKey generates a public key such that key = g**a % p
func PublicKey(a, p *big.Int, g int64) (key *big.Int) {
	key = (&big.Int{}).Exp(big.NewInt(g), a, p)
	return
}

// SecretKey generates a secret key such that key = B**a % p
func SecretKey(a, B, p *big.Int) (key *big.Int) {
	key = (&big.Int{}).Exp(B, a, p)
	return
}

// NewPair generates a private and public key pair
func NewPair(p *big.Int, g int64) (a, A *big.Int) {
	a = PrivateKey(p)
	A = PublicKey(a, p, g)
	return
}

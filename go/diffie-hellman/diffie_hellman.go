package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	max := big.NewInt(0).Sub(p, big.NewInt(2))
	pk, _ := rand.Int(rand.Reader, max)
	return big.NewInt(0).Add(pk, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	pk := PrivateKey(p)
	return pk, PublicKey(pk, p, g)
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return big.NewInt(0).Exp(public2, private1, p)
}

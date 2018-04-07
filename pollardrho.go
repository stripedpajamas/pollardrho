package pollardrho

import (
	"errors"
	"math/big"
)

// DefaultPolynomial is the commonly used polynomial g(x) = (x^2 + 1) mod n
func DefaultPolynomial(x, n *big.Int) *big.Int {
	one := big.NewInt(1)
	two := big.NewInt(2)
	x2 := new(big.Int).Exp(x, two, n) // x^2 mod n
	x2.Add(x2, one)                   // (x^2 mod n) + 1
	x2.Mod(x2, n)                     // (x^2 + 1) mod n
	return x2
}

// Rho is an implementation of Pollard's rho factorization algorithm
// using the default parameters x = y = 2
func Rho(n *big.Int, g func(*big.Int, *big.Int) *big.Int) (*big.Int, error) {
	one := big.NewInt(1)
	x, y, d := big.NewInt(2), big.NewInt(2), big.NewInt(1)

	for d.Cmp(one) == 0 {
		x = g(x, n)
		y = g(g(y, n), n)
		sub := new(big.Int).Sub(x, y)
		d.GCD(nil, nil, sub.Abs(sub), n) // gcd(|x - y|, n)
	}

	if d.Cmp(n) == 0 {
		return nil, errors.New("algorithm failed with default parameters (x = y = 2)")
	}

	return d, nil
}

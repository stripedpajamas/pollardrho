package pollardrho

import (
	"math/big"
	"testing"
)

func TestDefaultPolynomial(t *testing.T) {
	n := big.NewInt(8051)
	x := big.NewInt(677)

	result := DefaultPolynomial(x, n)
	want := big.NewInt(7474)

	if result.Cmp(want) != 0 {
		t.Fail()
	}
}

func TestRho(t *testing.T) {
	n := big.NewInt(8051) // 97 * 83
	factor, err := Rho(n, DefaultPolynomial)
	if err != nil {
		t.Error("Factorization failed:", err.Error())
	}
	if factor.Cmp(big.NewInt(97)) != 0 && factor.Cmp(big.NewInt(83)) != 0 {
		t.Fail()
	}

	n.SetInt64(10403) // 101 * 103
	factor, err = Rho(n, DefaultPolynomial)
	if err != nil {
		t.Error("Factorization failed:", err.Error())
	}
	if factor.Cmp(big.NewInt(101)) != 0 && factor.Cmp(big.NewInt(103)) != 0 {
		t.Fail()
	}
}

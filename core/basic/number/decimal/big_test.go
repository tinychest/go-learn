package decimal

import (
	"math/big"
	"testing"
)

// big 包可以提高计算的精度，但

func TestBig(t *testing.T) {
	f1 := big.NewFloat(0.1)
	f2 := big.NewFloat(0.2)
	f3 := big.NewFloat(0.3)

	a1, o1 := f1.Add(f1, f2).Float64()
	t.Log(a1, o1)
	a2, o2 := f3.Float64()
	t.Log(a2, o2)

	r := f1.Cmp(f3)
	t.Log(r) // → f1 + f2 > f3
}

package basic

import "testing"

func TestIota(t *testing.T) {
	const (
		A = 1 << (5 - iota) // 32
		B                   // 16
	)
}

package todo

import (
	"encoding/binary"
	"math"
	"testing"
)

func TestFloat642Byte(t *testing.T) {
	f := 2.33
	bs := Float642Byte(f)
	t.Log(string(bs))
	f = Byte2Float64(bs)
	t.Log(f)
}

func Float642Byte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func Byte2Float64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}

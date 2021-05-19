package bitreader

import (
	"errors"
	"math"
)

type Reader struct {
	bits []uint8
}

// NewReader returns a new Reader
func NewReader(input []byte) *Reader {
	bits := bytesToBits(input)
	return &Reader{
		bits: bits,
	}
}

// SliceToInt reads length bits from offset and returns them
func (r *Reader) SliceToInt(offset uint64, length uint64) (uint64, error) {
	if length < 1 || length > 64 {
		return 0, errors.New("invalid length")
	}
	if uint64(len(r.bits)) < offset+length {
		return 0, errors.New("invalid sum of offset and length value")
	}
	result := bitsToInt(r.bits[offset : offset+length])
	return result, nil
}

func bytesToBits(data []byte) []uint8 {
	result := make([]uint8, 0, 8*len(data))
	for _, v := range data {
		for i := 0; i < 8; i++ {
			move := uint(7 - i)
			value := (v >> move) & 1
			result = append(result, uint8(value))
		}
	}
	return result
}

func bitsToInt(bits []uint8) uint64 {
	var result uint64
	for i, v := range bits {
		result = result + uint64(float64(v)*math.Pow(2, float64(len(bits)-i-1)))
	}
	return result
}

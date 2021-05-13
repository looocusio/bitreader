package bitreader

import (
	"errors"
	"math"
)

type Reader struct {
	bits []int
}

// NewReader returns a new Reader
func NewReader(input []byte) *Reader {
	bits := bytesToBits(input)
	return &Reader{
		bits: bits,
	}
}

// SliceToInt reads length bits from offset and returns them
func (r *Reader) SliceToInt(offset int, length int) (int, error) {
	if len(r.bits) < offset+length {
		return 0, errors.New("invalid offset and length value")
	}
	result := bitsToInt(r.bits[offset : offset+length])
	return result, nil
}

func bytesToBits(data []byte) []int {
	result := make([]int, 0)
	for _, v := range data {
		for i := 0; i < 8; i++ {
			move := uint(7 - i)
			value := (v >> move) & 1
			result = append(result, int(value))
		}
	}
	return result
}

func bitsToInt(bits []int) int {
	r := reverseIntSlice(bits)
	var result int
	for i, v := range r {
		result = result + v*int(math.Pow(2, float64(i)))
	}
	return result
}

func reverseIntSlice(input []int) []int {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
	return input
}

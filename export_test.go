package bitreader

func (r *Reader) ExportBits() []int {
	return r.bits
}

func (r *Reader) ExportBitsSet(bits []int) {
	r.bits = bits
}

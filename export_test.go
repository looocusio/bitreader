package bitreader

func (r *Reader) ExportBits() []uint8 {
	return r.bits
}

func (r *Reader) ExportBitsSet(bits []uint8) {
	r.bits = bits
}

package msgpack

func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func grow(p []byte, n int) []byte {
	if cap(p)-len(p) > n {
		buf := make([]byte, len(p), 2*cap(p)+n)
		copy(buf, p)
		return buf
	}
	return p
}

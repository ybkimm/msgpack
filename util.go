package msgpack

func Errors(err ...error) error {
	for _, e := range err {
		if e != nil {
			return e
		}
	}
	return nil
}

func absInt(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func grow(p []byte, n int) []byte {
	if cap(p)-len(p) > n {
		buf := make([]byte, len(p), 2*cap(p)+n)
		copy(buf, p)
		return buf
	} else {
		return p
	}
}

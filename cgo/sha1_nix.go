package cgo

func (d *digest) Write(p []byte) (nn int, err error) {
	if len(p) == 0 {
		return 0, nil
	}

	return len(p), nil
}

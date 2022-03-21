package crypto

// Pkcs5pad is a padding function that uses the PKCS5 method.
func Pkcs5pad(data []byte, blocksize int) []byte {
	pad := blocksize - len(data)%blocksize
	b := make([]byte, pad, pad)
	for i := 0; i < pad; i++ {
		b[i] = uint8(pad)
	}
	return append(data, b...)
}

// Pkcs5unpad is a stripping function that reverts the PKCS5 method.
func Pkcs5unpad(data []byte) []byte {
	pad := int(data[len(data)-1])
	// FIXME: check that the padding bytes are all what we expect
	return data[:len(data)-pad]
}

package crypto

import (
	"crypto/rand"
	"github.com/minio/sio"
	"io"
)

/*
NewSecretBox returns a SecretBox object with a provided key.

Here is the validation code written for Python:
	import os,binascii
	key=binascii.b2a_base64(os.urandom(32),newline=False)
	print(key.decode('us-ascii'))
*/
func NewSecretBox(key []byte) SecretBox {
	streamConfig := sio.Config{
		Key:  key,
		Rand: rand.Reader,
	}
	var (
		encryptReader = func(src io.Reader) (io.Reader, error) {
			return sio.EncryptReader(src, streamConfig)
		}
		decryptReader = func(src io.Reader) (io.Reader, error) {
			return sio.DecryptReader(src, streamConfig)
		}
		encryptWriter = func(dst io.Writer) (io.WriteCloser, error) {
			return sio.EncryptWriter(dst, streamConfig)
		}
		decryptWriter = func(dst io.Writer) (io.WriteCloser, error) {
			return sio.DecryptWriter(dst, streamConfig)
		}
	)
	return &secretBoxImpl{
		encryptReaderGen: encryptReader,
		decryptReaderGen: decryptReader,
		encryptWriterGen: encryptWriter,
		decryptWriterGen: decryptWriter,
	}
}

// SecretBox is an encryption and decryption provider.
type SecretBox interface {
	// NewEncryptReader returns EncryptReader from given io.Reader.
	NewEncryptReader(io.Reader) (io.Reader, error)
	// NewDecryptReader returns DecryptReader from given io.Reader.
	NewDecryptReader(src io.Reader) (io.Reader, error)
	// NewEncryptWriter returns EncryptWriter from given io.Writer.
	NewEncryptWriter(dst io.Writer) (io.WriteCloser, error)
	// NewDecryptWriter returns DecryptWriter from given io.Writer.
	NewDecryptWriter(dst io.Writer) (io.WriteCloser, error)
	// EncryptedSize returns encrypted data size from given raw data size.
	EncryptedSize(size uint64) (uint64, error)
	// DecryptedSize returns decrypted data size from given raw data size.
	DecryptedSize(size uint64) (uint64, error)
}

type secretBoxImpl struct {
	encryptReaderGen func(src io.Reader) (io.Reader, error)
	decryptReaderGen func(src io.Reader) (io.Reader, error)
	encryptWriterGen func(dst io.Writer) (io.WriteCloser, error)
	decryptWriterGen func(dst io.Writer) (io.WriteCloser, error)
}

// 인터페이스가 실제 dto랑 호환되는가
var _ SecretBox = (*secretBoxImpl)(nil)

func (x *secretBoxImpl) NewEncryptReader(r io.Reader) (io.Reader, error) {
	return x.encryptReaderGen(r)
}
func (x *secretBoxImpl) NewDecryptReader(r io.Reader) (io.Reader, error) {
	return x.decryptReaderGen(r)
}
func (x *secretBoxImpl) NewEncryptWriter(w io.Writer) (io.WriteCloser, error) {
	return x.encryptWriterGen(w)
}
func (x *secretBoxImpl) NewDecryptWriter(w io.Writer) (io.WriteCloser, error) {
	return x.decryptWriterGen(w)
}
func (x *secretBoxImpl) EncryptedSize(size uint64) (uint64, error) { return sio.EncryptedSize(size) }
func (x *secretBoxImpl) DecryptedSize(size uint64) (uint64, error) { return sio.DecryptedSize(size) }

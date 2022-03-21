package crypto

import (
	"bytes"
	"crypto/cipher"
	"errors"
	"io"
)

// ErrShortCiphertext is an error that has text too short.
var ErrShortCiphertext = errors.New("input too short to be valid ciphertext")

/*
NewCryptoWriter returns block cipher writer.
from https://github.com/acasajus/dkeyczar/blob/master/streams.go
*/
func NewCryptoWriter(bm cipher.BlockMode, sink io.WriteCloser) io.WriteCloser {
	return &cryptoWriter{
		bm:     bm,
		sink:   sink,
		buffer: bytes.NewBuffer(nil),
	}
}

type cryptoWriter struct {
	bm     cipher.BlockMode
	buffer *bytes.Buffer
	sink   io.WriteCloser
	count  int
}

func (c *cryptoWriter) Write(data []byte) (int, error) {
	if _, err := c.buffer.Write(data); err != nil {
		return 0, err
	}
	bL := c.buffer.Len() - c.buffer.Len()%c.bm.BlockSize()
	tmp := c.buffer.Next(bL)
	c.bm.CryptBlocks(tmp, tmp)
	wL := 0
	for wL < len(tmp) {
		n, err := c.sink.Write(tmp[wL:])
		if err != nil {
			return 0, err
		}
		wL += n
	}
	c.count += wL
	return len(data), nil
}

func (c *cryptoWriter) Close() error {
	tmp := Pkcs5pad(c.buffer.Next(c.buffer.Len()), c.bm.BlockSize())
	c.bm.CryptBlocks(tmp, tmp)
	wL := 0
	for wL < len(tmp) {
		n, err := c.sink.Write(tmp[wL:])
		if err != nil {
			return err
		}
		wL += n
	}
	c.count += wL
	return c.sink.Close()
}

func NewCryptoReader(bm cipher.BlockMode, source io.ReadCloser) io.ReadCloser {
	return &cryptoReader{
		bm:     bm,
		source: source,
		outBuf: bytes.NewBuffer(nil),
		inBuf:  bytes.NewBuffer(nil),
		eof:    false,
	}
}

type cryptoReader struct {
	bm     cipher.BlockMode
	outBuf *bytes.Buffer
	inBuf  *bytes.Buffer
	source io.ReadCloser
	eof    bool
}

func (cr *cryptoReader) Read(data []byte) (int, error) {
	missing := len(data) - cr.outBuf.Len()
	for !cr.eof && missing > 0 {
		toRead := missing + cr.bm.BlockSize() + 1 //Always go beyond the required data to be able to unpad when eof'ed
		if off := toRead % cr.bm.BlockSize(); off > 0 {
			toRead += cr.bm.BlockSize() - off //Make sure we read in multiples of blocksize
		}
		cr.inBuf.Grow(toRead)
		n, err := io.CopyN(cr.inBuf, cr.source, int64(toRead))
		if err == io.EOF {
			cr.eof = true
		} else if err != nil {
			return 0, err
		}
		readBytes := int(n)
		if readBytes%cr.bm.BlockSize() > 0 && cr.eof {
			return 0, ErrShortCiphertext
		}
		bytesToDec := readBytes - readBytes%cr.bm.BlockSize()
		tmpdata := cr.inBuf.Next(bytesToDec)
		cr.bm.CryptBlocks(tmpdata, tmpdata)
		if _, err := cr.outBuf.Write(tmpdata); err != nil {
			return 0, err
		}
		if cr.eof {
			pad := cr.outBuf.Bytes()[cr.outBuf.Len()-1]
			cr.outBuf.Truncate(cr.outBuf.Len() - int(pad))
		}
		missing = len(data) - cr.outBuf.Len()
	}
	return cr.outBuf.Read(data)
}

func (cr *cryptoReader) Close() error {
	cr.eof = true
	return cr.source.Close()
}

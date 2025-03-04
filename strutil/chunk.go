package strutil

import (
	"io"
)

const pemLineLength = 64

// LineBreaker is an io.Writer that advances a newline if one line exceeds 64 bytes.
// this source originally comes from the url(https://cs.opensource.google/go/go/+/refs/tags/go1.24.0:src/encoding/pem/pem.go;l=89)
type LineBreaker struct {
	// Out os
	Out  io.Writer
	line [pemLineLength]byte
	used int
}

var nl = []byte{'\n'}

// Write writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written from p (0 <= n <= len(p))
// and any error encountered that caused the write to stop early.
// Write must return a non-nil error if it returns n < len(p).
// Write must not modify the slice data, even temporarily.
func (l *LineBreaker) Write(b []byte) (n int, err error) {
	if l.used+len(b) < pemLineLength {
		copy(l.line[l.used:], b)
		l.used += len(b)
		return len(b), nil
	}

	_, err = l.Out.Write(l.line[0:l.used:l.used])
	if err != nil {
		return
	}
	brk := pemLineLength - l.used

	var nn int
	for len(b) >= brk {
		nn, err = l.Out.Write(b[0:brk:brk])
		n += nn
		if err != nil {
			return
		}

		_, err = l.Out.Write(nl)
		if err != nil {
			return
		}
		b = b[brk:]
		brk = pemLineLength
	}

	l.used = len(b)
	copy(l.line[:], b)
	n += len(b)
	return
}

// Close flushes any pending output from the writer.
// It is an error to call Write after calling Close.
func (l *LineBreaker) Close() (err error) {
	if l.used > 0 {
		_, err = l.Out.Write(l.line[0:l.used])
		if err != nil {
			return
		}
		_, err = l.Out.Write(nl)
	}
	if err != nil {
		return
	} else if closer, ok := l.Out.(io.Closer); ok {
		err = closer.Close()
	}
	return
}

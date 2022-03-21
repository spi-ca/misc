package strutil

import (
	"io"
)

const pemLineLength = 64

// LineBreaker is an io.Writer that advances a newline if one line exceeds 64 bytes.
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

	n, err = l.Out.Write(l.line[0:l.used])
	if err != nil {
		return
	}
	excess := pemLineLength - l.used
	l.used = 0

	n, err = l.Out.Write(b[0:excess])
	if err != nil {
		return
	}

	n, err = l.Out.Write(nl)
	if err != nil {
		return
	}

	return l.Out.Write(b[excess:])
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

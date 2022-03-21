package dbg

import "bytes"

import (
	"fmt"
	"golang.org/x/text/width"
	"unicode/utf8"
)

// HexDump returns hexadecimal data dump data.
func HexDump(by []byte) string {
	n := len(by)
	rowcount := 0
	stop := (n / 16) * 16
	k := 0
	buf := &bytes.Buffer{}

	for i := 0; i <= stop; i += 16 {
		k++
		if i+16 < n {
			rowcount = 16
		} else {
			rowcount = min(k*16, n) % 16
		}

		fmt.Fprintf(buf, "%08x ", i)
		for j := 0; j < rowcount; j++ {
			if j%8 == 0 {
				fmt.Fprintf(buf, " %02x ", by[i+j])
			} else {
				fmt.Fprintf(buf, "%02x ", by[i+j])
			}

		}

		for j := rowcount; j < 16; j++ {
			if j%8 == 0 {
				fmt.Fprintf(buf, "    ")
			} else {
				fmt.Fprintf(buf, "   ")
			}
		}
		buf.WriteRune('|')
		viewString(by[i:(i+rowcount)], buf)
		buf.WriteRune('|')
		buf.WriteRune('\n')
		buf.WriteRune('\r')
	}
	return buf.String()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GuessUnicodeWidth returns the size of bytes for a single rune.
func GuessUnicodeWidth(char rune) (realSize int) {
	prop := width.LookupRune(char)
	switch prop.Kind() {
	case width.EastAsianFullwidth:
		fallthrough
	case width.EastAsianWide:
		realSize = 2
	case width.EastAsianHalfwidth:
		fallthrough
	case width.EastAsianNarrow:
		realSize = 2
	case width.EastAsianAmbiguous:
		fallthrough
	case width.Neutral:
		fallthrough
	default:
		realSize = 1
	}
	return
}

func viewString(b []byte, buf *bytes.Buffer) {
	for {
		if r, size := utf8.DecodeRune(b); size == 0 {
			return
		} else if r == utf8.RuneError {
			for i := 0; i < size; i++ {
				buf.WriteRune('_')
			}
			b = b[size:]
		} else if r < 32 {
			for i := 0; i < size; i++ {
				buf.WriteRune('.')
			}
			b = b[size:]
		} else {
			buf.WriteRune(r)
			pad := max(0, size-GuessUnicodeWidth(r))
			for i := 0; i < pad; i++ {
				buf.WriteRune('.')
			}
			b = b[size:]
		}
	}
}

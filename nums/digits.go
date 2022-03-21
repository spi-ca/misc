package nums

import (
	"github.com/spi-ca/misc/strutil"
	"strings"
)

// SplitSignFromInt64 decomposes a signed integer with a sign and positive integer.
func SplitSignFromInt64(src int64) (sign bool, dst uint64) {
	signedInteger := uint64(src)
	sign = 1 == (signedInteger >> 63)

	//Two's complement!!
	if sign {
		dst = 1 + (signedInteger ^ (1<<64 - 1))
	} else {
		dst = signedInteger & (1<<63 - 1)
	}
	return
}

// MergeSignFromUInt64 composes a signed integer with a sign and positive integer.
func MergeSignFromUInt64(sign bool, src uint64) (dst int64) {
	//Two's complement!!
	if sign {
		dst = -int64(src & (1<<63 - 1))
	} else {
		dst = int64(src & (1<<63 - 1))
	}
	return
}

// FractionalStringify is a formatting function for a fractional part number.
func FractionalStringify(fractional, fractionalLength int, showPoint bool) (formatted string) {
	if fractionalLength == 0 {
		return
	}
	var builder strings.Builder

	partialFormatted := strutil.FormatIntToStringReversed(fractional)
	if leftDigits := fractionalLength - len(partialFormatted); leftDigits < 0 {
		builder.WriteString(partialFormatted[-leftDigits:])
	} else {
		builder.WriteString(partialFormatted)
		for ; leftDigits > 0; leftDigits-- {
			builder.WriteByte('0')
		}
	}

	if showPoint {
		builder.WriteByte('.')
	}

	return strutil.Reverse(builder.String())
}

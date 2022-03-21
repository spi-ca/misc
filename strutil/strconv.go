package strutil

import (
	"bytes"
	"math"
	"strings"
)

const (
	fingerNum = 10
	zeroChr   = '0'
)

// Reverse is a function that returns a string in reverse order.
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// FormatIntToStringReversed is a function that converts number to decimal string reversed order(LSB).
func FormatIntToStringReversed(number int) (out string) {
	if number == 0 {
		return "0"
	}
	var (
		buf   strings.Builder
		digit int
		sign  = number < 0
	)
	if sign {
		number = -number
	}
	for radix := 0; number > 0; radix++ {
		digit, number = number%fingerNum, number/fingerNum
		buf.WriteByte(zeroChr + byte(digit))
	}
	if sign {
		buf.WriteByte('-')
	}
	return buf.String()
}

// FormatIntToString is a function that converts number to decimal string(MSB).
func FormatIntToString(number int) (out string) {
	if number == 0 {
		return "0"
	}
	var (
		buf   strings.Builder
		digit int
		sign  = number < 0
	)

	if sign {
		number = -number
		buf.WriteByte('-')
	}

	var (
		maxDigit = int(math.Log10(float64(number))) + 1
		numBuf   = make([]byte, maxDigit)
	)

	for radix := 0; number > 0; radix++ {
		digit, number = number%fingerNum, number/fingerNum
		numBuf[maxDigit-radix-1] = zeroChr + byte(digit)
	}
	buf.Write(numBuf)

	return buf.String()
}

// FormatIntToBytesReversed is a function that converts number to decimal byte array(ASCII code based) reversed order(LSB).
func FormatIntToBytesReversed(number int) (out []byte) {
	if number == 0 {
		return []byte("0")
	}
	var (
		buf   bytes.Buffer
		digit int
		sign  = number < 0
	)

	if sign {
		number = -number
	}
	for radix := 0; number > 0; radix++ {
		digit, number = number%fingerNum, number/fingerNum
		buf.WriteByte(zeroChr + byte(digit))
	}

	if sign {
		buf.WriteByte('-')
	}
	return buf.Bytes()
}

// FormatIntToBytes is a function that converts number to decimal byte array(ASCII code based,MSB).
func FormatIntToBytes(number int) (out []byte) {
	if number == 0 {
		return []byte("0")
	}

	var (
		digit    int
		sign     = number < 0
		maxDigit int
	)

	if sign {
		number = -number
		maxDigit = int(math.Log10(float64(number))) + 2
	} else {
		maxDigit = int(math.Log10(float64(number))) + 1
	}

	out = make([]byte, maxDigit)
	if sign {
		out[0] = '-'
	}

	for radix := 0; number > 0; radix++ {
		digit, number = number%fingerNum, number/fingerNum
		out[maxDigit-radix-1] = zeroChr + byte(digit)
	}
	return
}

// ParseStringToInt is a function that converts a decimal string to integer.
func ParseStringToInt(in string) (number int) {
	if len(in) == 0 {
		return
	}

	var (
		out  []uint8
		sign = in[0] == '-'
	)
	if sign {
		out = make([]uint8, 0, len(in)-1)
		in = in[1:]
	} else {
		out = make([]uint8, 0, len(in))
	}
	for _, digit := range in {
		if digit < '0' || digit > '9' {
			continue
		}
		out = append(out, byte(digit-'0'))
	}
	scanned := len(out)
	pow := 1

	for i := scanned; i > 0; i-- {
		number += pow * int(out[i-1])
		pow *= 10
	}

	if sign {
		number = -number
	}
	return
}

// ParseBytesToInt is a function that converts a decimal byte array(ASCII based) to integer.
func ParseBytesToInt(in []byte) (number int) {
	if len(in) == 0 {
		return -1
	}

	var (
		out  []uint8
		sign = in[0] == '-'
	)

	if sign {
		out = make([]uint8, 0, len(in)-1)
		in = in[1:]
	} else {
		out = make([]uint8, 0, len(in))
	}

	for _, digit := range in {
		if digit < '0' || digit > '9' {
			continue
		}
		out = append(out, digit-'0')
	}
	scanned := len(out)
	if scanned == 0 {
		return -1
	}
	pow := 1

	for i := scanned; i > 0; i-- {
		number += pow * int(out[i-1])
		pow *= 10
	}

	if sign {
		number = -number
	}
	return
}

// FormatUnsignedIntToStringReversed is a function that converts number to decimal string reversed order(LSB).
func FormatUnsignedIntToStringReversed(sign bool, number uint) (out string) {
	if number == 0 {
		return "0"
	}
	var (
		buf   strings.Builder
		digit uint
	)

	for radix := 0; number > 0; radix++ {
		digit, number = number%fingerNum, number/fingerNum
		buf.WriteByte(zeroChr + byte(digit))
	}
	if sign {
		buf.WriteByte('-')
	}
	return buf.String()
}

// FormatUnsignedIntToString is a function that converts number to decimal string(MSB).
func FormatUnsignedIntToString(sign bool, number uint) (out string) {
	if number == 0 {
		return "0"
	}
	var (
		buf      strings.Builder
		digit    uint
		maxDigit = int(math.Log10(float64(number))) + 1
		numBuf   = make([]byte, maxDigit)
	)
	if sign {
		buf.WriteByte('-')
	}
	for radix := 0; number > 0; radix++ {
		digit, number = number%fingerNum, number/fingerNum
		numBuf[maxDigit-radix-1] = zeroChr + byte(digit)
	}

	buf.Write(numBuf)

	return buf.String()
}

// FormatUnsignedIntToBytesReversed is a function that converts number to decimal byte array(ASCII code based) reversed order(LSB).
func FormatUnsignedIntToBytesReversed(sign bool, number uint) (out []byte) {
	if number == 0 {
		return []byte("0")
	}
	var (
		buf   bytes.Buffer
		digit uint
	)
	for radix := 0; number > 0; radix++ {
		digit, number = number%fingerNum, number/fingerNum
		buf.WriteByte(zeroChr + byte(digit))
	}
	if sign {
		buf.WriteByte('-')
	}
	return buf.Bytes()
}

// FormatUnsignedIntToBytes is a function that converts number to decimal byte array(ASCII code based, MSB).
func FormatUnsignedIntToBytes(sign bool, number uint) (out []byte) {
	if number == 0 {
		return []byte("0")
	}
	var (
		digit    uint
		maxDigit = int(math.Log10(float64(number))) + 1
	)

	if sign {
		maxDigit++
		out = make([]byte, maxDigit)
		out[0] = '-'
	} else {
		out = make([]byte, maxDigit)
	}

	for radix := 0; number > 0; radix++ {
		digit, number = number%fingerNum, number/fingerNum
		out[maxDigit-radix-1] = zeroChr + byte(digit)
	}
	return
}

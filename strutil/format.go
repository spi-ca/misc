package strutil

import (
	"math"
	"strings"
)

// FillBytes fill the destination byte array with the given pattern.
func FillBytes(dst []byte, pattern []byte) {
	for i := 0; i < len(dst); i++ {
		dst[i] = pattern[i%len(pattern)]
	}
}

// FillByte fills the destination byte array with a single byte.
func FillByte(dst []byte, pattern byte) {
	for i := 0; i < len(dst); i++ {
		dst[i] = pattern
	}
}

// StrPad method pads the input string with the padString until the resulting string reaches the given length
func StrPad(input string, padLength int, padString string, rightPad bool) (output string) {
	var (
		inputLength     = len(input)
		padStringLength = len(padString)
	)
	if inputLength >= padLength {
		return input
	}

	var (
		repeat  = int(math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength)))
		builder strings.Builder
	)
	builder.Grow(inputLength + padStringLength*repeat)
	if rightPad {
		builder.WriteString(input)
		for i := 0; i < repeat; i++ {
			builder.WriteString(padString)
		}
		output = builder.String()[:padLength]
	} else {
		for i := 0; i < repeat; i++ {
			builder.WriteString(padString)
		}
		builder.WriteString(input)
		output = builder.String()[builder.Len()-padLength:]
	}

	return
}

// StrPadSingle method pads the input string with a single character until the resulting string reaches the given length
func StrPadSingle(input string, padLength int, pad byte, rightPad bool) (output string) {
	var (
		inputLength = len(input)
	)
	if inputLength >= padLength {
		return input
	}

	var (
		builder strings.Builder
	)
	builder.Grow(inputLength + padLength)
	if rightPad {
		builder.WriteString(input)
		for i := 0; i < padLength; i++ {
			builder.WriteByte(pad)
		}
		output = builder.String()[:padLength]
	} else {
		for i := 0; i < padLength; i++ {
			builder.WriteByte(pad)
		}
		builder.WriteString(input)
		output = builder.String()[builder.Len()-padLength:]
	}

	return
}

// BytesPad method pads the input byte array with the padData byte array until the resulting string reaches the given length
func BytesPad(input []byte, padLength int, padData []byte, rightPad bool) (output []byte) {
	var (
		inputLength   = len(input)
		padDataLength = len(padData)
	)
	if inputLength >= padLength {
		output = make([]byte, inputLength)
		copy(output, input)
		return
	}

	var (
		repeat            = int(math.Ceil(float64(1) + (float64(padLength-padDataLength))/float64(padDataLength)))
		maxFillLength     = repeat * padDataLength
		bufSize           = inputLength + repeat*padDataLength
		padArea, fillArea []byte
	)

	output = make([]byte, bufSize)

	if rightPad {
		fillArea, padArea = output[0:inputLength], output[inputLength:bufSize]
		output = output[:padLength]
	} else {
		fillArea, padArea = output[maxFillLength:bufSize], output[0:maxFillLength]
		output = output[inputLength+maxFillLength-padLength:]
	}

	// copy data
	copy(fillArea, input)
	// fill pad
	for i := 0; i < repeat; i++ {
		copy(padArea[:padDataLength], padData)
		padArea = padArea[padDataLength:]
	}

	return
}

// BytesPad method pads the input byte array with a single byte until the resulting string reaches the given length
func BytesPadSingle(input []byte, padLength int, pad byte, rightPad bool) (output []byte) {
	var (
		inputLength = len(input)
	)
	if inputLength >= padLength {
		output = make([]byte, inputLength)
		copy(output, input)
		return
	}

	var (
		maxFillLength     = padLength
		bufSize           = inputLength + padLength
		padArea, fillArea []byte
	)

	output = make([]byte, bufSize)

	if rightPad {
		fillArea, padArea = output[0:inputLength], output[inputLength:bufSize]
		output = output[:padLength]
	} else {
		fillArea, padArea = output[maxFillLength:bufSize], output[0:maxFillLength]
		output = output[inputLength+maxFillLength-padLength:]
	}

	// copy data
	copy(fillArea, input)
	// fill pad
	FillByte(padArea, pad)
	return
}

// BytesUnPadSingle method removes a given pad character from both ends.
func BytesUnPadSingle(input []byte, pad byte, rightPad bool, copyData bool) (output []byte) {
	var (
		limit  = len(input)
		offset = 0
	)
	if rightPad {
		for ; limit > offset; limit-- {
			if input[limit-1] != pad {
				break
			}
		}
	} else {
		for ; offset < limit; offset++ {
			if input[offset] != pad {
				break
			}
		}
	}
	if limit-offset < 1 {
		return
	}
	if copyData {
		output = make([]byte, limit-offset)
		copy(output, input[offset:limit])
	} else {
		output = input[offset:limit]
	}
	return
}

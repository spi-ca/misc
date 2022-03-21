package strutil

import (
	"math"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	// EmailRegex is the standard email address format.
	EmailRegex = regexp.MustCompile(`^([^<>()[\]\\.,;:\s@"]+(?:\.[^<>()[\]\\.,;:\s@"]+)*)@(?:(?:[a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,})$`)
	// DigitRegex represents a digit.
	DigitRegex = regexp.MustCompile(`\d`)
)

// MaskText is a masking function for hide sensitive data.
func MaskText(src string) (val string) {

	var dst strings.Builder
	defer func(ptr *string) { *ptr = dst.String() }(&val)

	// email
	if ret := EmailRegex.FindAllStringSubmatchIndex(src, -1); len(ret) > 0 {
		start := src[:ret[0][2]]
		username := src[ret[0][2]:ret[0][3]]
		end := src[ret[0][3]:]
		length := utf8.RuneCountInString(username)
		_, firstCharLength := utf8.DecodeRuneInString(username)
		_, lastCharLength := utf8.DecodeLastRuneInString(username)
		_, secondLastCharLength := utf8.DecodeLastRuneInString(username[:len(username)-lastCharLength])
		dst.WriteString(start)
		if length > 3 {
			dst.WriteString(username[:firstCharLength])
			for i := 0; i < length-3; i++ {
				dst.WriteByte('*')
			}
			dst.WriteString(username[len(username)-(secondLastCharLength+lastCharLength):])
		} else {
			dst.WriteByte('*')
			dst.WriteString(username[firstCharLength:])
		}
		dst.WriteString(end)
	} else if ret = DigitRegex.FindAllStringSubmatchIndex(src, -1); len(ret) > 2 {
		var (
			offset     = 0
			factor     = float64(len(ret)) / 3.0
			pos        = int(math.Floor(factor))
			fillLength = int(math.Ceil(factor))
		)
		for i := pos; i < pos+fillLength; i++ {
			dst.WriteString(src[offset:ret[i][0]])
			dst.WriteByte('*')
			offset = ret[i][1]
		}
		dst.WriteString(src[offset:])
	} else {
		length := utf8.RuneCountInString(src)
		_, firstCharLength := utf8.DecodeRuneInString(src)
		_, lastCharLength := utf8.DecodeLastRuneInString(src)
		if length > 2 {
			dst.WriteString(src[:firstCharLength])
			for i := 0; i < length-2; i++ {
				dst.WriteByte('*')
			}
			dst.WriteString(src[len(src)-lastCharLength:])
		} else if length > 0 {
			dst.WriteByte('*')
			dst.WriteString(src[firstCharLength:])
		}
	}

	return
}

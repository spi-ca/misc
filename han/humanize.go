package han

import (
	"amuz.es/src/go/misc/strutil"
	"strings"
)

const (
	koreanPluralizeSubunit = 10
	koreanPluralizeUnit    = 10000
	zeroChr                = '0'
)

var (
	koreanPluralizeUnitLabels    = []string{"", "만", "억", "조"}
	koreanPluralizeSubunitLabels = []string{"", "십", "백", "천"}
)

func transKoreanPluralizeSubunit(number int) (out string) {
	if number == 0 {
		return "0"
	}
	var (
		buf   strings.Builder
		digit int
	)
	for radix := 0; number > 0; radix++ {
		digit, number = number%koreanPluralizeSubunit, number/koreanPluralizeSubunit
		if digit > 0 {
			buf.WriteString(koreanPluralizeSubunitLabels[radix%len(koreanPluralizeSubunitLabels)])
		}
		if digit > 1 || (digit == 1 && radix == 0) {
			buf.WriteByte(zeroChr + byte(digit))
		}
	}
	return buf.String()
}

func transKoreanPluralizeUnit(number int, subunitmapper func(int) string, omitLastOne bool) (out string) {
	if number == 0 {
		return "0"
	}
	var (
		buf              strings.Builder
		digit, lastDigit int
	)
	for radix := 0; number > 0; radix++ {
		digit = number % koreanPluralizeUnit
		number = number / koreanPluralizeUnit
		lastDigit = digit % koreanPluralizeSubunit
		if digit > 0 {
			buf.WriteString(koreanPluralizeUnitLabels[radix%len(koreanPluralizeUnitLabels)])
		}
		// 1 뭉개기
		if omitLastOne && radix > 0 && lastDigit == 1 {
			digit--
		}
		//나머지 처리
		if digit > 1 || (digit == 1 && (radix == 0 || !omitLastOne)) {
			buf.WriteString(subunitmapper(digit))
		}
	}
	return buf.String()
}

// KoreanPluralizeSubunit is a function for translating into Korean after omitting the most significant digit when the most significant digit is 1.
func KoreanPluralizeSubunit(number int) (out string) {
	out = strutil.Reverse(transKoreanPluralizeSubunit(number))
	return
}

// KoreanPluralizeUnit is a function for translating into Korean after omitting the most significant digit when the most significant digit is 1.
func KoreanPluralizeUnit(number int) (out string) {
	out = strutil.Reverse(transKoreanPluralizeUnit(number, transKoreanPluralizeSubunit, true))
	return
}

// KoreanPluralizeUnitType2 function is a function that translates only each 10,000 units into Korean.
func KoreanPluralizeUnitType2(number int) (out string) {
	out = strutil.Reverse(transKoreanPluralizeUnit(number, strutil.FormatIntToStringReversed, false))
	return
}

//KoreanPluralizeUnitType3 is a function for translating into Korean.
func KoreanPluralizeUnitType3(number int) (out string) {
	if number == 0 {
		return "0"
	}
	var (
		below10thousand                        = number % koreanPluralizeUnit
		above10thousand                        = (number / koreanPluralizeUnit) * koreanPluralizeUnit
		outBelow10thousand, outAbove10thousand string
	)

	if below10thousand > 0 {
		outBelow10thousand = strutil.FormatIntToStringReversed(below10thousand)
	}
	if above10thousand > 0 {
		outAbove10thousand = transKoreanPluralizeUnit(above10thousand, transKoreanPluralizeSubunit, false)
	}
	out = strutil.Reverse(outBelow10thousand + outAbove10thousand)
	return
}

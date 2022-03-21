package strutil

import (
	"context"
	"crypto/rand"
	"math"
	"math/big"
	"regexp"
	"strconv"
	"time"
)

// PasswordCheck checks a password candidate with minimum length and regexp conditions.
func PasswordCheck(source []byte, minLength int, condition ...*regexp.Regexp) (invalid bool) {
	if len(source) < minLength {
		return true
	}
	for _, cond := range condition {
		if !cond.Match(source) {
			return true
		}
	}
	return
}

// RandomDigits returns authenticate code with given length.
func RandomDigits(digitLength, limit int) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	longest := func(source string) (longest int) {
		var (
			current int32
			count   = 0
		)
		for _, digitRune := range source {
			if digitRune == current {
				count++
			} else {
				count = 1
				current = digitRune
			}
			if count > longest {
				longest = count
			}
		}
		return
	}
	getNonce := func(digits int) string {
		seed, _ := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
		return strconv.FormatUint(seed.Uint64(), 10)
	}
	nonce := getNonce(digitLength)
	for longest(nonce) >= limit {
		select {
		case <-ctx.Done():
			panic(ctx.Err())
		default:
			nonce = getNonce(digitLength)
		}
	}
	return nonce
}

package han

import (
	"github.com/valyala/fasthttp"
	"go.uber.org/multierr"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
	"net/url"
	"strings"
	"unicode"
)

const (
	hexConstUpper = "0123456789ABCDEF"
)

// ConvertFormEUCKRFasthttp is translate method from euc-kr encoded fasthttp.Args to utf-8 encoded fasthttp.Args.
func ConvertFormEUCKRFasthttp(rawParsed *fasthttp.Args) (parsed *fasthttp.Args, err error) {
	var (
		//rawKey, rawValue string
		decoder = korean.EUCKR.NewDecoder()
	)

	parsed = &fasthttp.Args{}
	rawParsed.VisitAll(func(k, v []byte) {
		if rawKey, decodeErr := decoder.Bytes(k); decodeErr != nil {
			err = multierr.Append(err, decodeErr)
		} else if rawValue, decodeErr := decoder.Bytes(v); decodeErr != nil {
			err = multierr.Append(err, decodeErr)
		} else {
			parsed.SetBytesKV(rawKey, rawValue)
		}
	})

	return
}

// ConvertFormEUCKR is translate method from euc-kr encoded url.Values to utf-8 encoded url.Values.
func ConvertFormEUCKR(rawParsed url.Values) (parsed url.Values, err error) {
	var (
		rawKey, rawValue string
		decoder          = korean.EUCKR.NewDecoder()
	)

	parsed = make(url.Values, len(rawParsed))
	for k, v := range rawParsed {
		rawKey, err = decoder.String(k)
		if err != nil {
			return
		}
		valueList := make([]string, 0, len(v))
		for _, vi := range v {
			rawValue, err = decoder.String(vi)
			if err != nil {
				return
			}
			valueList = append(valueList, rawValue)
		}
		parsed[rawKey] = valueList
	}
	return
}

// ConvertToEUCKRUrlEncoded is translate method from utf-8 encoded string to euc-kr encoded string.
func ConvertToEUCKRUrlEncoded(str string) (converted string, err error) {
	var (
		buf     strings.Builder
		encoder = korean.EUCKR.NewEncoder()
		size    int
	)
	for _, r := range str {
		switch {
		case unicode.IsNumber(r) ||
			unicode.Is(unicode.Latin, r) ||
			r == '-' ||
			r == '_' ||
			r == '.':
			buf.WriteRune(r)
		case unicode.IsSpace(r):
			buf.WriteByte('+')
		case r == '$' ||
			r == '&' ||
			r == '+' ||
			r == ',' ||
			r == '/' ||
			r == ':' ||
			r == ';' ||
			r == '=' ||
			r == '?' ||
			r == '@':
			buf.WriteByte('%')
			buf.WriteByte(hexConstUpper[byte(r)>>4])
			buf.WriteByte(hexConstUpper[byte(r)&15])
		default:
			converted, size, err = transform.String(encoder, string(r))
			if err != nil {
				return
			}
			buf.WriteByte('%')
			buf.WriteByte(hexConstUpper[(converted[0]>>4)&0xf])
			buf.WriteByte(hexConstUpper[converted[0]&0xf])
			if size > 1 {
				buf.WriteByte('%')
				buf.WriteByte(hexConstUpper[(converted[1]>>4)&0xf])
				buf.WriteByte(hexConstUpper[converted[1]&0xf])
			}
		}
	}
	converted = buf.String()
	return
}

// ConvertFromEUCKRUrlEncoded is translate method from euc-kr encoded string to utf-8 encoded string.
func ConvertFromEUCKRUrlEncoded(str string) (converted string, err error) {
	dst, err := url.QueryUnescape(str)
	if err != nil {
		return
	}
	converted, _, err = transform.String(korean.EUCKR.NewDecoder(), dst)
	return
}

package strutil

import (
	"io"
	"strings"
	"testing"
)

type lineBreakerTest struct {
	in, out string
}

const sixtyFourCharString = "0123456789012345678901234567890123456789012345678901234567890123"

var lineBreakerTests = []lineBreakerTest{
	{"", ""},
	{"a", "a\n"},
	{"ab", "ab\n"},
	{sixtyFourCharString, sixtyFourCharString + "\n"},
	{sixtyFourCharString + "X", sixtyFourCharString + "\nX\n"},
	{sixtyFourCharString + sixtyFourCharString, sixtyFourCharString + "\n" + sixtyFourCharString + "\n"},
}

func TestLineBreaker(t *testing.T) {
	for i, test := range lineBreakerTests {
		buf := new(strings.Builder)
		var breaker LineBreaker
		breaker.Out = buf
		_, err := breaker.Write([]byte(test.in))
		if err != nil {
			t.Errorf("#%d: error from Write: %s", i, err)
			continue
		}
		err = breaker.Close()
		if err != nil {
			t.Errorf("#%d: error from Close: %s", i, err)
			continue
		}

		if got := buf.String(); got != test.out {
			t.Errorf("#%d: got:%s want:%s", i, got, test.out)
		}
	}

	for i, test := range lineBreakerTests {
		buf := new(strings.Builder)
		var breaker LineBreaker
		breaker.Out = buf

		for i := 0; i < len(test.in); i++ {
			_, err := breaker.Write([]byte(test.in[i : i+1]))
			if err != nil {
				t.Errorf("#%d: error from Write (byte by byte): %s", i, err)
				continue
			}
		}
		err := breaker.Close()
		if err != nil {
			t.Errorf("#%d: error from Close (byte by byte): %s", i, err)
			continue
		}

		if got := buf.String(); got != test.out {
			t.Errorf("#%d: (byte by byte) got:%s want:%s", i, got, test.out)
		}
	}
}

func FuzzLineBreaker(f *testing.F) {
	for _, test := range lineBreakerTests {
		f.Add(test.in, len(test.in))
		f.Add(test.in, 1)
	}

	f.Fuzz(func(t *testing.T, in string, chunkSize int) {
		if chunkSize <= 0 || chunkSize > len(in)+1 {
			return
		}
		var out strings.Builder
		var l LineBreaker
		l.Out = &out
		chunk := make([]byte, chunkSize)
		n, err := io.CopyBuffer(&l, strings.NewReader(in), chunk)
		if err != nil {
			t.Fatal(err)
		}
		if n != int64(len(in)) {
			t.Errorf("invalid written count: got %d, expected %d", n, len(in))
		}
		l.Close()
		if len(in) > 0 && out.Len() != len(in)+1+(len(in)-1)/pemLineLength {
			t.Fatalf("invalid final size: got %d, expected %d", out.Len(), len(in)+1+(len(in)-1)/pemLineLength)
		}
	})
}

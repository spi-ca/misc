package strutil

import (
	"reflect"
	"testing"
)

func TestFormatIntToBytesReversed(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
	}{
		{
			"simple",
			args{1234567890},
			[]byte("0987654321"),
		},
		{
			"-simple",
			args{-1234567890},
			[]byte("0987654321-"),
		},
		{
			"0",
			args{0},
			[]byte("0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := FormatIntToBytesReversed(tt.args.number); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("FormatIntToBytesReversed() = %v, want %v", string(gotOut), string(tt.wantOut))
			}
		})
	}
}
func TestFormatIntToBytes(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
	}{
		{
			"simple",
			args{1234567890},
			[]byte("1234567890"),
		},
		{
			"-simple",
			args{-1234567890},
			[]byte("-1234567890"),
		},
		{
			"0",
			args{0},
			[]byte("0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := FormatIntToBytes(tt.args.number); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("FormatIntToBytes() = %v, want %v", string(gotOut), string(tt.wantOut))
			}
		})
	}
}

func TestFormatIntToStringReversed(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			"simple",
			args{1234567890},
			"0987654321",
		},
		{
			"-simple",
			args{-1234567890},
			"0987654321-",
		},
		{
			"0",
			args{0},
			"0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := FormatIntToStringReversed(tt.args.number); gotOut != tt.wantOut {
				t.Errorf("FormatIntToStringReversed() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestFormatIntToString(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			"simple",
			args{1234567890},
			"1234567890",
		},
		{
			"simple",
			args{-1234567890},
			"-1234567890",
		},
		{
			"0",
			args{0},
			"0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := FormatIntToString(tt.args.number); gotOut != tt.wantOut {
				t.Errorf("FormatIntToString() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestParseBytesToInt(t *testing.T) {
	type args struct {
		in []byte
	}
	tests := []struct {
		name       string
		args       args
		wantNumber int
	}{
		{
			"simple",
			args{[]byte("1234567890")},
			1234567890,
		},
		{
			"simple",
			args{[]byte("-1234567890")},
			-1234567890,
		},
		{
			"0",
			args{[]byte("0")},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumber := ParseBytesToInt(tt.args.in); gotNumber != tt.wantNumber {
				t.Errorf("ParseBytesToInt() = %v, want %v", gotNumber, tt.wantNumber)
			}
		})
	}
}

func TestParseStringToInt(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name       string
		args       args
		wantNumber int
	}{
		{
			"simple",
			args{"1234567890"},
			1234567890,
		},
		{
			"-simple",
			args{"-1234567890"},
			-1234567890,
		},
		{
			"0",
			args{"0"},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNumber := ParseStringToInt(tt.args.in); gotNumber != tt.wantNumber {
				t.Errorf("ParseStringToInt() = %v, want %v", gotNumber, tt.wantNumber)
			}
		})
	}
}

func TestFormatUnsignedIntToStringReversed(t *testing.T) {
	type args struct {
		sign   bool
		number uint
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			"simple",
			args{false, 1234567890},
			"0987654321",
		},
		{
			"-simple",
			args{true, 1234567890},
			"0987654321-",
		},
		{
			"0",
			args{false, 0},
			"0",
		},
		{
			"-0",
			args{true, 0},
			"0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := FormatUnsignedIntToStringReversed(tt.args.sign, tt.args.number); gotOut != tt.wantOut {
				t.Errorf("FormatUnsignedIntToStringReversed() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestFormatUnsignedIntToString(t *testing.T) {
	type args struct {
		sign   bool
		number uint
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			"simple",
			args{false, 1234567890},
			"1234567890",
		},
		{
			"-simple",
			args{true, 1234567890},
			"-1234567890",
		},
		{
			"0",
			args{false, 0},
			"0",
		},
		{
			"-0",
			args{true, 0},
			"0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := FormatUnsignedIntToString(tt.args.sign, tt.args.number); gotOut != tt.wantOut {
				t.Errorf("FormatUnsignedIntToString() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestFormatUnsignedIntToBytesReversed(t *testing.T) {
	type args struct {
		sign   bool
		number uint
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
	}{
		{
			"simple",
			args{false, 1234567890},
			[]byte("0987654321"),
		},
		{
			"-simple",
			args{true, 1234567890},
			[]byte("0987654321-"),
		},
		{
			"0",
			args{false, 0},
			[]byte("0"),
		},
		{
			"0",
			args{true, 0},
			[]byte("0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := FormatUnsignedIntToBytesReversed(tt.args.sign, tt.args.number); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("FormatUnsignedIntToBytesReversed() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func TestFormatUnsignedIntToBytes(t *testing.T) {
	type args struct {
		sign   bool
		number uint
	}
	tests := []struct {
		name    string
		args    args
		wantOut []byte
	}{{
		"simple",
		args{false, 1234567890},
		[]byte("1234567890"),
	},
		{
			"-simple",
			args{true, 1234567890},
			[]byte("-1234567890"),
		},
		{
			"0",
			args{false, 0},
			[]byte("0"),
		},
		{
			"-0",
			args{true, 0},
			[]byte("0"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := FormatUnsignedIntToBytes(tt.args.sign, tt.args.number); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("FormatUnsignedIntToBytes() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

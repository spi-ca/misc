package strutil

import (
	"reflect"
	"testing"
)

func TestStrPad(t *testing.T) {
	const input = "Codes"
	type args struct {
		input     string
		padLength int
		padString string
		rightPad  bool
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			"rpad test",
			args{input, 10, " ", true},
			"Codes     ",
		},
		{
			"lpad test",
			args{input, 10, "-=", false},
			"=-=-=Codes",
		},
		{
			"rpad filled test",
			args{input, len(input), "_*", true},
			input,
		},
		{
			"lpad filled test",
			args{input, len(input), "*", false},
			input,
		},
		{
			"rpad overflow test",
			args{input, 3, "_*", true},
			input,
		},
		{
			"lpad overflow test",
			args{input, 3, "*", false},
			input,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := StrPad(tt.args.input, tt.args.padLength, tt.args.padString, tt.args.rightPad); gotOutput != tt.wantOutput {
				t.Errorf("StrPad() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestStrPadSingle(t *testing.T) {
	const input = "Codes"
	type args struct {
		input     string
		padLength int
		pad       byte
		rightPad  bool
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
	}{
		{
			"rpad test",
			args{input, 10, ' ', true},
			"Codes     ",
		},
		{
			"lpad test",
			args{input, 10, '-', false},
			"-----Codes",
		},
		{
			"rpad filled test",
			args{input, len(input), '*', true},
			input,
		},
		{
			"lpad filled test",
			args{input, len(input), '*', false},
			input,
		},
		{
			"rpad overflow test",
			args{input, 3, '*', true},
			input,
		},
		{
			"lpad overflow test",
			args{input, 3, '*', false},
			input,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := StrPadSingle(tt.args.input, tt.args.padLength, tt.args.pad, tt.args.rightPad); gotOutput != tt.wantOutput {
				t.Errorf("StrPad() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
func TestBytesPad(t *testing.T) {
	const input = "Codes"
	type args struct {
		input     []byte
		padLength int
		padData   []byte
		rightPad  bool
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []byte
	}{
		{
			"rpad test",
			args{[]byte(input), 10, []byte(" "), true},
			[]byte("Codes     "),
		},
		{
			"lpad test",
			args{[]byte(input), 10, []byte("-="), false},
			[]byte("=-=-=Codes"),
		},
		{
			"rpad filled test",
			args{[]byte(input), len(input), []byte("_*"), true},
			[]byte(input),
		},
		{
			"lpad filled test",
			args{[]byte(input), len(input), []byte("*"), false},
			[]byte(input),
		},
		{
			"rpad overflow test",
			args{[]byte(input), 3, []byte("_*"), true},
			[]byte(input),
		},
		{
			"lpad overflow test",
			args{[]byte(input), 3, []byte("*"), false},
			[]byte(input),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := BytesPad(tt.args.input, tt.args.padLength, tt.args.padData, tt.args.rightPad); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("BytesPad() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

func TestBytesPadSingle(t *testing.T) {
	const input = "Codes"
	type args struct {
		input     []byte
		padLength int
		padData   byte
		rightPad  bool
	}
	tests := []struct {
		name       string
		args       args
		wantOutput []byte
	}{
		{
			"rpad test",
			args{[]byte(input), 10, ' ', true},
			[]byte("Codes     "),
		},
		{
			"lpad test",
			args{[]byte(input), 10, '_', false},
			[]byte("_____Codes"),
		},
		{
			"rpad filled test",
			args{[]byte(input), len(input), '*', true},
			[]byte(input),
		},
		{
			"lpad filled test",
			args{[]byte(input), len(input), '*', false},
			[]byte(input),
		},
		{
			"rpad overflow test",
			args{[]byte(input), 3, '*', true},
			[]byte(input),
		},
		{
			"lpad overflow test",
			args{[]byte(input), 3, '*', false},
			[]byte(input),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOutput := BytesPadSingle(tt.args.input, tt.args.padLength, tt.args.padData, tt.args.rightPad); !reflect.DeepEqual(gotOutput, tt.wantOutput) {
				t.Errorf("BytesPad() = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}

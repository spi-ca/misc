package nums

import (
	"github.com/ericlagergren/decimal"
	"reflect"
	"testing"
)

func TestBigSeperateFractionalPotions(t *testing.T) {
	type args struct {
		src *decimal.Big
	}
	tests := []struct {
		name                 string
		args                 args
		wantSign             bool
		wantInteger          uint64
		wantFractional       uint64
		wantFractionalLength uint8
	}{
		{
			"11.111",
			args{
				decimal.New(11111, 3),
			},
			false,
			11,
			111,
			3,
		},
		{
			"0.111",
			args{
				decimal.New(111, 3),
			},
			false,
			0,
			111,
			3,
		},
		{
			"11",
			args{
				decimal.New(11000, 3),
			},
			false,
			11,
			0,
			0,
		},
		{
			"10.01",
			args{
				decimal.New(10010, 3),
			},
			false,
			10,
			1,
			2,
		},
		{
			"0.013",
			args{
				decimal.New(13, 3),
			},
			false,
			0,
			13,
			3,
		},
		{
			"-11.111",
			args{
				decimal.New(-11111, 3),
			},
			true,
			11,
			111,
			3,
		},
		{
			"-0.111",
			args{
				decimal.New(-111, 3),
			},
			true,
			0,
			111,
			3,
		},
		{
			"-11",
			args{
				decimal.New(-11000, 3),
			},
			true,
			11,
			0,
			0,
		},
		{
			"-10.01",
			args{
				decimal.New(-10010, 3),
			},
			true,
			10,
			1,
			2,
		},
		{
			"-0.013",
			args{
				decimal.New(-13, 3),
			},
			true,
			0,
			13,
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSign, gotInteger, gotFractional, gotFractionalLength := BigSeperateFractionalPotions(tt.args.src)
			t.Logf("BigSeperateFractionalPotions() sign =%v, integer = %v, fractional = %v, fractionalLength = %v", gotSign, gotInteger, gotFractional, gotFractionalLength)
			if gotSign != tt.wantSign {
				t.Errorf("BigSeperateFractionalPotions() gotSign = %v, want %v", gotSign, tt.wantSign)
			}
			if gotInteger != tt.wantInteger {
				t.Errorf("BigSeperateFractionalPotions() gotInteger = %v, want %v", gotInteger, tt.wantInteger)
			}
			if gotFractional != tt.wantFractional {
				t.Errorf("BigSeperateFractionalPotions() gotFractional = %v, want %v", gotFractional, tt.wantFractional)
			}
			if gotFractionalLength != tt.wantFractionalLength {
				t.Errorf("BigSeperateFractionalPotions() gotFractionalLength = %v, want %v", gotFractionalLength, tt.wantFractionalLength)
			}
		})
	}
}

func TestBigMergeFractionalPotions(t *testing.T) {
	type args struct {
		sign             bool
		integer          uint64
		fractional       uint64
		fractionalLength uint8
	}
	tests := []struct {
		name    string
		args    args
		wantDst *decimal.Big
	}{
		{
			"11.111",
			args{
				false,
				11,
				111,
				3,
			},
			decimal.New(11111, 3),
		},
		{
			"0.111",
			args{
				false,
				0,
				111,
				3,
			},
			decimal.New(111, 3),
		},
		{
			"11",
			args{
				false,
				11,
				0,
				0,
			},
			decimal.New(11, 0),
		},
		{
			"10.01",
			args{
				false,
				10,
				1,
				2,
			},
			decimal.New(1001, 2),
		},
		{
			"0.013",
			args{
				false,
				0,
				13,
				3,
			},
			decimal.New(13, 3),
		},
		{
			"-11.111",
			args{
				true,
				11,
				111,
				3,
			},
			decimal.New(-11111, 3),
		},
		{
			"-0.111",
			args{
				true,
				0,
				111,
				3,
			},
			decimal.New(-111, 3),
		},
		{
			"-11",
			args{
				true,
				11,
				0,
				0,
			},
			decimal.New(-11, 0),
		},
		{
			"-10.01",
			args{
				true,
				10,
				1,
				2,
			},
			decimal.New(-1001, 2),
		},
		{
			"-0.013",
			args{
				true,
				0,
				13,
				3,
			},
			decimal.New(-13, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDst := BigMergeFractionalPotions(tt.args.sign, tt.args.integer, tt.args.fractional, tt.args.fractionalLength); !reflect.DeepEqual(gotDst, tt.wantDst) {
				t.Errorf("BigMergeFractionalPotions() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

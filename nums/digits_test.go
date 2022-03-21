package nums

import "testing"

func TestSplitSignFromInt64(t *testing.T) {
	type args struct {
		src int64
	}
	tests := []struct {
		name     string
		args     args
		wantSign bool
		wantDst  uint64
	}{
		{
			"0",
			args{0},
			false, 0,
		},
		{
			"-0",
			args{-0},
			false, 0,
		},
		{
			"-223",
			args{-223},
			true, 223,
		},
		{
			"112",
			args{112},
			false, 112,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSign, gotDst := SplitSignFromInt64(tt.args.src)
			if gotSign != tt.wantSign {
				t.Errorf("SplitSignFromInt64() gotSign = %v, want %v", gotSign, tt.wantSign)
			}
			if gotDst != tt.wantDst {
				t.Errorf("SplitSignFromInt64() gotDst = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestMergeSignFromUInt64(t *testing.T) {
	type args struct {
		sign bool
		src  uint64
	}
	tests := []struct {
		name    string
		args    args
		wantDst int64
	}{
		{
			"0",
			args{false, 0},
			0,
		},
		{
			"-0",
			args{true, 0},
			0,
		},
		{
			"-223",
			args{true, 223},
			-223,
		},
		{
			"112",
			args{false, 112},
			112,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDst := MergeSignFromUInt64(tt.args.sign, tt.args.src); gotDst != tt.wantDst {
				t.Errorf("MergeSignFromUInt64() = %v, want %v", gotDst, tt.wantDst)
			}
		})
	}
}

func TestFractionalStringify(t *testing.T) {
	type args struct {
		fractional       int
		fractionalLength int
		showPoint        bool
	}
	tests := []struct {
		name          string
		args          args
		wantFormatted string
	}{
		{
			".123",
			args{
				123, 3, true,
			},
			".123",
		},
		{
			"length 0",
			args{
				123, 0, true,
			},
			"",
		},
		{
			".1",
			args{
				123, 1, true,
			},
			".1",
		},
		{
			".12",
			args{
				123, 2, true,
			},
			".12",
		},
		{
			".0123",
			args{
				123, 4, true,
			},
			".0123",
		},
		{
			".01200",
			args{
				1200, 5, true,
			},
			".01200",
		},
		{
			"123",
			args{
				123, 3, false,
			},
			"123",
		},
		{
			"length 0",
			args{
				123, 0, false,
			},
			"",
		},
		{
			"1",
			args{
				123, 1, false,
			},
			"1",
		},
		{
			"12",
			args{
				123, 2, false,
			},
			"12",
		},
		{
			"0123",
			args{
				123, 4, false,
			},
			"0123",
		},
		{
			"01200",
			args{
				1200, 5, false,
			},
			"01200",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFormatted := FractionalStringify(tt.args.fractional, tt.args.fractionalLength, tt.args.showPoint); gotFormatted != tt.wantFormatted {
				t.Errorf("FractionalStringify() = %v, want %v", gotFormatted, tt.wantFormatted)
			} else {
				t.Logf("FractionalStringify() = %v", gotFormatted)
			}
		})
	}
}

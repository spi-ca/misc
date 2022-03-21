package strutil

import (
	"testing"
)

func TestMaskText(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		wantVal string
	}{
		{
			"address",
			args{"test@example.com"},
			"t*st@example.com",
		},
		{
			"address_long",
			args{"testtesttesttest@example.com"},
			"t*************st@example.com",
		},
		{
			"address_short",
			args{"tst@example.com"},
			"*st@example.com",
		},
		{
			"address_short_2",
			args{"tt@example.com"},
			"*t@example.com",
		},
		{
			"address_very_short",
			args{"t@example.com"},
			"*@example.com",
		},
		{
			"non-formal",
			args{"texample.com"},
			"t**********m",
		},
		{
			"non-formal2",
			args{"te🧭ample.co🧭"},
			"t**********🧭",
		},
		{
			"non-formal short",
			args{"com"},
			"c*m",
		},
		{
			"non-formal short 2",
			args{"⏰om"},
			"⏰*m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotVal := MaskText(tt.args.input); gotVal != tt.wantVal {
				t.Errorf("MaskText() = %v, want %v", gotVal, tt.wantVal)
			} else {
				t.Log("==> ", gotVal)
			}
		})
	}
}

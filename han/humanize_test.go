package han

import (
	"testing"
)

func TestKoreanPluralizeSubunit(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			"0",
			args{0},
			"0",
		},
		{
			"1",
			args{1},
			"1",
		},
		{
			"9",
			args{9},
			"9",
		},
		{
			"10",
			args{10},
			"십",
		},
		{
			"12",
			args{12},
			"십2",
		},
		{
			"23",
			args{23},
			"2십3",
		},
		{
			"40",
			args{40},
			"4십",
		},
		{
			"100",
			args{100},
			"백",
		},
		{
			"106",
			args{106},
			"백6",
		},
		{
			"123",
			args{123},
			"백2십3",
		},
		{
			"170",
			args{170},
			"백7십",
		},
		{
			"234",
			args{234},
			"2백3십4",
		},
		{
			"500",
			args{500},
			"5백",
		},
		{
			"509",
			args{509},
			"5백9",
		},
		{
			"519",
			args{519},
			"5백십9",
		},
		{
			"1000",
			args{1000},
			"천",
		},
		{
			"1001",
			args{1001},
			"천1",
		},
		{
			"1005",
			args{1005},
			"천5",
		},
		{
			"1010",
			args{1010},
			"천십",
		},
		{
			"1015",
			args{1015},
			"천십5",
		},
		{
			"1060",
			args{1060},
			"천6십",
		},
		{
			"1064",
			args{1064},
			"천6십4",
		},
		{
			"1100",
			args{1100},
			"천백",
		},
		{
			"1104",
			args{1104},
			"천백4",
		},
		{
			"1140",
			args{1140},
			"천백4십",
		},
		{
			"1123",
			args{1123},
			"천백2십3",
		},
		{
			"1200",
			args{1200},
			"천2백",
		},
		{
			"1201",
			args{1201},
			"천2백1",
		},
		{
			"1204",
			args{1204},
			"천2백4",
		},
		{
			"1240",
			args{1240},
			"천2백4십",
		},
		{
			"1234",
			args{1234},
			"천2백3십4",
		},

		{
			"6000",
			args{6000},
			"6천",
		},
		{
			"6001",
			args{6001},
			"6천1",
		},
		{
			"6005",
			args{6005},
			"6천5",
		},
		{
			"6010",
			args{6010},
			"6천십",
		},
		{
			"6015",
			args{6015},
			"6천십5",
		},
		{
			"6060",
			args{6060},
			"6천6십",
		},
		{
			"6064",
			args{6064},
			"6천6십4",
		},
		{
			"6100",
			args{6100},
			"6천백",
		},
		{
			"6104",
			args{6104},
			"6천백4",
		},
		{
			"6140",
			args{6140},
			"6천백4십",
		},
		{
			"6123",
			args{6123},
			"6천백2십3",
		},
		{
			"6200",
			args{6200},
			"6천2백",
		},
		{
			"6201",
			args{6201},
			"6천2백1",
		},
		{
			"6204",
			args{6204},
			"6천2백4",
		},
		{
			"6240",
			args{6240},
			"6천2백4십",
		},
		{
			"6234",
			args{6234},
			"6천2백3십4",
		},
		{
			"987654321",
			args{987654321},
			"98천7백6십54천3백2십1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := KoreanPluralizeSubunit(tt.args.number); gotOut != tt.wantOut {
				t.Errorf("KoreanPluralizeSubunit() = %v, want %v", gotOut, tt.wantOut)
			} else {
				t.Logf("KoreanPluralizeSubunit() = %v", gotOut)
			}
		})
	}
}

func TestKoreanPluralizeUnit(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			"0",
			args{0},
			"0",
		},
		{
			"1",
			args{1},
			"1",
		},
		{
			"9",
			args{9},
			"9",
		},
		{
			"10",
			args{10},
			"십",
		},
		{
			"12",
			args{12},
			"십2",
		},
		{
			"23",
			args{23},
			"2십3",
		},
		{
			"40",
			args{40},
			"4십",
		},
		{
			"100",
			args{100},
			"백",
		},
		{
			"106",
			args{106},
			"백6",
		},
		{
			"123",
			args{123},
			"백2십3",
		},
		{
			"170",
			args{170},
			"백7십",
		},
		{
			"234",
			args{234},
			"2백3십4",
		},
		{
			"500",
			args{500},
			"5백",
		},
		{
			"509",
			args{509},
			"5백9",
		},
		{
			"519",
			args{519},
			"5백십9",
		},
		{
			"1000",
			args{1000},
			"천",
		},
		{
			"1001",
			args{1001},
			"천1",
		},
		{
			"1005",
			args{1005},
			"천5",
		},
		{
			"1010",
			args{1010},
			"천십",
		},
		{
			"1015",
			args{1015},
			"천십5",
		},
		{
			"1060",
			args{1060},
			"천6십",
		},
		{
			"1064",
			args{1064},
			"천6십4",
		},
		{
			"1100",
			args{1100},
			"천백",
		},
		{
			"1104",
			args{1104},
			"천백4",
		},
		{
			"1140",
			args{1140},
			"천백4십",
		},
		{
			"1123",
			args{1123},
			"천백2십3",
		},
		{
			"1200",
			args{1200},
			"천2백",
		},
		{
			"1201",
			args{1201},
			"천2백1",
		},
		{
			"1204",
			args{1204},
			"천2백4",
		},
		{
			"1240",
			args{1240},
			"천2백4십",
		},
		{
			"1234",
			args{1234},
			"천2백3십4",
		},

		{
			"6000",
			args{6000},
			"6천",
		},
		{
			"6001",
			args{6001},
			"6천1",
		},
		{
			"6005",
			args{6005},
			"6천5",
		},
		{
			"6010",
			args{6010},
			"6천십",
		},
		{
			"6015",
			args{6015},
			"6천십5",
		},
		{
			"6060",
			args{6060},
			"6천6십",
		},
		{
			"6064",
			args{6064},
			"6천6십4",
		},
		{
			"6100",
			args{6100},
			"6천백",
		},
		{
			"6104",
			args{6104},
			"6천백4",
		},
		{
			"6140",
			args{6140},
			"6천백4십",
		},
		{
			"6123",
			args{6123},
			"6천백2십3",
		},
		{
			"6200",
			args{6200},
			"6천2백",
		},
		{
			"6201",
			args{6201},
			"6천2백1",
		},
		{
			"6204",
			args{6204},
			"6천2백4",
		},
		{
			"6240",
			args{6240},
			"6천2백4십",
		},
		{
			"6234",
			args{6234},
			"6천2백3십4",
		},
		{
			"10000",
			args{10000},
			"만",
		},
		{
			"20000",
			args{20000},
			"2만",
		},
		{
			"11110000",
			args{11110000},
			"천백십만",
		},
		{
			"11111111",
			args{11111111},
			"천백십만천백십1",
		},
		{
			"10101010",
			args{10101010},
			"천십만천십",
		},
		{
			"10101012",
			args{10101012},
			"천십만천십2",
		},
		{
			"99990000",
			args{99990000},
			"9천9백9십9만",
		},
		{
			"99999999",
			args{99999999},
			"9천9백9십9만9천9백9십9",
		},
		{
			"100000000",
			args{100000000},
			"억",
		},
		{
			"100010000",
			args{100010000},
			"억만",
		},
		{
			"111110000",
			args{111110000},
			"억천백십만",
		},
		{
			"111111111",
			args{111111111},
			"억천백십만천백십1",
		},
		{
			"110101010",
			args{110101010},
			"억천십만천십",
		},
		{
			"110101012",
			args{110101012},
			"억천십만천십2",
		},
		{
			"199990000",
			args{199990000},
			"억9천9백9십9만",
		},
		{
			"199999999",
			args{199999999},
			"억9천9백9십9만9천9백9십9",
		},
		{
			"200000000",
			args{200000000},
			"2억",
		},
		{
			"200010000",
			args{200010000},
			"2억만",
		},
		{
			"211110000",
			args{211110000},
			"2억천백십만",
		},
		{
			"211111111",
			args{211111111},
			"2억천백십만천백십1",
		},
		{
			"210101010",
			args{210101010},
			"2억천십만천십",
		},
		{
			"210101012",
			args{210101012},
			"2억천십만천십2",
		},
		{
			"299990000",
			args{299990000},
			"2억9천9백9십9만",
		},
		{
			"299999999",
			args{299999999},
			"2억9천9백9십9만9천9백9십9",
		},
		{
			"606060606060",
			args{606060606060},
			"6천6십억6천6십만6천6십",
		},
		{
			"606060606066",
			args{606060606066},
			"6천6십억6천6십만6천6십6",
		},
		{
			"666600000000",
			args{666600000000},
			"6천6백6십6억",
		},
		{
			"1000000000000",
			args{1000000000000},
			"조",
		},
		{
			"1000000000001",
			args{1000000000001},
			"조1",
		},
		{
			"1000000000009",
			args{1000000000009},
			"조9",
		},
		{
			"1000000000010",
			args{1000000000010},
			"조십",
		},
		{
			"1000000000012",
			args{1000000000012},
			"조십2",
		},
		{
			"1000000000023",
			args{1000000000023},
			"조2십3",
		},
		{
			"1000000000040",
			args{1000000000040},
			"조4십",
		},
		{
			"1000000000100",
			args{1000000000100},
			"조백",
		},
		{
			"1000000000106",
			args{1000000000106},
			"조백6",
		},
		{
			"1000000000123",
			args{1000000000123},
			"조백2십3",
		},
		{
			"1000000000170",
			args{1000000000170},
			"조백7십",
		},
		{
			"1000000000234",
			args{1000000000234},
			"조2백3십4",
		},
		{
			"1000000000500",
			args{1000000000500},
			"조5백",
		},
		{
			"1000000000509",
			args{1000000000509},
			"조5백9",
		},
		{
			"1000000000519",
			args{1000000000519},
			"조5백십9",
		},
		{
			"1000000001000",
			args{1000000001000},
			"조천",
		},
		{
			"1000000001001",
			args{1000000001001},
			"조천1",
		},
		{
			"1000000001005",
			args{1000000001005},
			"조천5",
		},
		{
			"1000000001010",
			args{1000000001010},
			"조천십",
		},
		{
			"1000000001015",
			args{1000000001015},
			"조천십5",
		},
		{
			"1000000001060",
			args{1000000001060},
			"조천6십",
		},
		{
			"1000000001064",
			args{1000000001064},
			"조천6십4",
		},
		{
			"1000000001100",
			args{1000000001100},
			"조천백",
		},
		{
			"1000000001104",
			args{1000000001104},
			"조천백4",
		},
		{
			"1000000001140",
			args{1000000001140},
			"조천백4십",
		},
		{
			"1000000001123",
			args{1000000001123},
			"조천백2십3",
		},
		{
			"1000000001200",
			args{1000000001200},
			"조천2백",
		},
		{
			"1000000001201",
			args{1000000001201},
			"조천2백1",
		},
		{
			"1000000001204",
			args{1000000001204},
			"조천2백4",
		},
		{
			"1000000001240",
			args{1000000001240},
			"조천2백4십",
		},
		{
			"1000000001234",
			args{1000000001234},
			"조천2백3십4",
		},
		{
			"1000000006000",
			args{1000000006000},
			"조6천",
		},
		{
			"1000000006001",
			args{1000000006001},
			"조6천1",
		},
		{
			"1000000006005",
			args{1000000006005},
			"조6천5",
		},
		{
			"1000000006010",
			args{1000000006010},
			"조6천십",
		},
		{
			"1000000006015",
			args{1000000006015},
			"조6천십5",
		},
		{
			"1000000006060",
			args{1000000006060},
			"조6천6십",
		},
		{
			"1000000006064",
			args{1000000006064},
			"조6천6십4",
		},
		{
			"1000000006100",
			args{1000000006100},
			"조6천백",
		},
		{
			"1000000006104",
			args{1000000006104},
			"조6천백4",
		},
		{
			"1000000006140",
			args{1000000006140},
			"조6천백4십",
		},
		{
			"1000000006123",
			args{1000000006123},
			"조6천백2십3",
		},
		{
			"1000000006200",
			args{1000000006200},
			"조6천2백",
		},
		{
			"1000000006201",
			args{1000000006201},
			"조6천2백1",
		},
		{
			"1000000006204",
			args{1000000006204},
			"조6천2백4",
		},
		{
			"1000000006240",
			args{1000000006240},
			"조6천2백4십",
		},
		{
			"1000000006234",
			args{1000000006234},
			"조6천2백3십4",
		},
		{
			"1000000010000",
			args{1000000010000},
			"조만",
		},
		{
			"1000000020000",
			args{1000000020000},
			"조2만",
		},
		{
			"1000011110000",
			args{1000011110000},
			"조천백십만",
		},
		{
			"1000011111111",
			args{1000011111111},
			"조천백십만천백십1",
		},
		{
			"1000010101010",
			args{1000010101010},
			"조천십만천십",
		},
		{
			"1000010101012",
			args{1000010101012},
			"조천십만천십2",
		},
		{
			"1000099990000",
			args{1000099990000},
			"조9천9백9십9만",
		},
		{
			"1000099999999",
			args{1000099999999},
			"조9천9백9십9만9천9백9십9",
		},
		{
			"1000100000000",
			args{1000100000000},
			"조억",
		},
		{
			"1000100010000",
			args{1000100010000},
			"조억만",
		},
		{
			"1000111110000",
			args{1000111110000},
			"조억천백십만",
		},
		{
			"1000111111111",
			args{1000111111111},
			"조억천백십만천백십1",
		},
		{
			"1000110101010",
			args{1000110101010},
			"조억천십만천십",
		},
		{
			"1000110101012",
			args{1000110101012},
			"조억천십만천십2",
		},
		{
			"1000199990000",
			args{1000199990000},
			"조억9천9백9십9만",
		},
		{
			"1000199999999",
			args{1000199999999},
			"조억9천9백9십9만9천9백9십9",
		},
		{
			"1000200000000",
			args{1000200000000},
			"조2억",
		},
		{
			"1000200010000",
			args{1000200010000},
			"조2억만",
		},
		{
			"1000211110000",
			args{1000211110000},
			"조2억천백십만",
		},
		{
			"1000211111111",
			args{1000211111111},
			"조2억천백십만천백십1",
		},
		{
			"1000210101010",
			args{1000210101010},
			"조2억천십만천십",
		},
		{
			"1000210101012",
			args{1000210101012},
			"조2억천십만천십2",
		},
		{
			"1000299990000",
			args{1000299990000},
			"조2억9천9백9십9만",
		},
		{
			"1000299999999",
			args{1000299999999},
			"조2억9천9백9십9만9천9백9십9",
		},
		{
			"1111111111111",
			args{1111111111111},
			"조천백십억천백십만천백십1",
		},
		{
			"1606060606060",
			args{1606060606060},
			"조6천6십억6천6십만6천6십",
		},
		{
			"1606060606066",
			args{1606060606066},
			"조6천6십억6천6십만6천6십6",
		},
		{
			"1666600000000",
			args{1666600000000},
			"조6천6백6십6억",
		},
		{
			"9876543210123456",
			args{9876543210123456},
			"9천8백7십6조5천4백3십2억천십2만3천4백5십6",
		},
		{
			"9999999999999999",
			args{9999999999999999},
			"9천9백9십9조9천9백9십9억9천9백9십9만9천9백9십9",
		},
		{
			"999999999999999999",
			args{999999999999999999},
			"9십99천9백9십9조9천9백9십9억9천9백9십9만9천9백9십9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := KoreanPluralizeUnit(tt.args.number); gotOut != tt.wantOut {
				t.Errorf("KoreanPluralizeUnit() = %v, want %v", gotOut, tt.wantOut)
			} else {
				t.Logf("KoreanPluralizeUnit() = %v", gotOut)
			}
		})
	}
}

func TestKoreanPluralizeUnitType2(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			"0",
			args{0},
			"0",
		},
		{
			"1",
			args{1},
			"1",
		},
		{
			"9",
			args{9},
			"9",
		},
		{
			"10",
			args{10},
			"10",
		},
		{
			"12",
			args{12},
			"12",
		},
		{
			"23",
			args{23},
			"23",
		},
		{
			"40",
			args{40},
			"40",
		},
		{
			"100",
			args{100},
			"100",
		},
		{
			"106",
			args{106},
			"106",
		},
		{
			"123",
			args{123},
			"123",
		},
		{
			"170",
			args{170},
			"170",
		},
		{
			"234",
			args{234},
			"234",
		},
		{
			"500",
			args{500},
			"500",
		},
		{
			"509",
			args{509},
			"509",
		},
		{
			"519",
			args{519},
			"519",
		},
		{
			"1000",
			args{1000},
			"1000",
		},
		{
			"1001",
			args{1001},
			"1001",
		},
		{
			"1005",
			args{1005},
			"1005",
		},
		{
			"1010",
			args{1010},
			"1010",
		},
		{
			"1015",
			args{1015},
			"1015",
		},
		{
			"1060",
			args{1060},
			"1060",
		},
		{
			"1064",
			args{1064},
			"1064",
		},
		{
			"1100",
			args{1100},
			"1100",
		},
		{
			"1104",
			args{1104},
			"1104",
		},
		{
			"1140",
			args{1140},
			"1140",
		},
		{
			"1123",
			args{1123},
			"1123",
		},
		{
			"1200",
			args{1200},
			"1200",
		},
		{
			"1201",
			args{1201},
			"1201",
		},
		{
			"1204",
			args{1204},
			"1204",
		},
		{
			"1240",
			args{1240},
			"1240",
		},
		{
			"1234",
			args{1234},
			"1234",
		},

		{
			"6000",
			args{6000},
			"6000",
		},
		{
			"6001",
			args{6001},
			"6001",
		},
		{
			"6005",
			args{6005},
			"6005",
		},
		{
			"6010",
			args{6010},
			"6010",
		},
		{
			"6015",
			args{6015},
			"6015",
		},
		{
			"6060",
			args{6060},
			"6060",
		},
		{
			"6064",
			args{6064},
			"6064",
		},
		{
			"6100",
			args{6100},
			"6100",
		},
		{
			"6104",
			args{6104},
			"6104",
		},
		{
			"6140",
			args{6140},
			"6140",
		},
		{
			"6123",
			args{6123},
			"6123",
		},
		{
			"6200",
			args{6200},
			"6200",
		},
		{
			"6201",
			args{6201},
			"6201",
		},
		{
			"6204",
			args{6204},
			"6204",
		},
		{
			"6240",
			args{6240},
			"6240",
		},
		{
			"6234",
			args{6234},
			"6234",
		},
		{
			"10000",
			args{10000},
			"1만",
		},
		{
			"20000",
			args{20000},
			"2만",
		},
		{
			"11110000",
			args{11110000},
			"1111만",
		},
		{
			"11111111",
			args{11111111},
			"1111만1111",
		},
		{
			"10101010",
			args{10101010},
			"1010만1010",
		},
		{
			"10101012",
			args{10101012},
			"1010만1012",
		},
		{
			"99990000",
			args{99990000},
			"9999만",
		},
		{
			"99999999",
			args{99999999},
			"9999만9999",
		},
		{
			"100000000",
			args{100000000},
			"1억",
		},
		{
			"100010000",
			args{100010000},
			"1억1만",
		},
		{
			"111110000",
			args{111110000},
			"1억1111만",
		},
		{
			"111111111",
			args{111111111},
			"1억1111만1111",
		},
		{
			"110101010",
			args{110101010},
			"1억1010만1010",
		},
		{
			"110101012",
			args{110101012},
			"1억1010만1012",
		},
		{
			"199990000",
			args{199990000},
			"1억9999만",
		},
		{
			"199999999",
			args{199999999},
			"1억9999만9999",
		},
		{
			"200000000",
			args{200000000},
			"2억",
		},
		{
			"200010000",
			args{200010000},
			"2억1만",
		},
		{
			"211110000",
			args{211110000},
			"2억1111만",
		},
		{
			"211111111",
			args{211111111},
			"2억1111만1111",
		},
		{
			"210101010",
			args{210101010},
			"2억1010만1010",
		},
		{
			"210101012",
			args{210101012},
			"2억1010만1012",
		},
		{
			"299990000",
			args{299990000},
			"2억9999만",
		},
		{
			"299999999",
			args{299999999},
			"2억9999만9999",
		},
		{
			"606060606060",
			args{606060606060},
			"6060억6060만6060",
		},
		{
			"606060606066",
			args{606060606066},
			"6060억6060만6066",
		},
		{
			"666600000000",
			args{666600000000},
			"6666억",
		},
		{
			"1000000000000",
			args{1000000000000},
			"1조",
		},
		{
			"1000000000001",
			args{1000000000001},
			"1조1",
		},
		{
			"1000000000009",
			args{1000000000009},
			"1조9",
		},
		{
			"1000000000010",
			args{1000000000010},
			"1조10",
		},
		{
			"1000000000012",
			args{1000000000012},
			"1조12",
		},
		{
			"1000000000023",
			args{1000000000023},
			"1조23",
		},
		{
			"1000000000040",
			args{1000000000040},
			"1조40",
		},
		{
			"1000000000100",
			args{1000000000100},
			"1조100",
		},
		{
			"1000000000106",
			args{1000000000106},
			"1조106",
		},
		{
			"1000000000123",
			args{1000000000123},
			"1조123",
		},
		{
			"1000000000170",
			args{1000000000170},
			"1조170",
		},
		{
			"1000000000234",
			args{1000000000234},
			"1조234",
		},
		{
			"1000000000500",
			args{1000000000500},
			"1조500",
		},
		{
			"1000000000509",
			args{1000000000509},
			"1조509",
		},
		{
			"1000000000519",
			args{1000000000519},
			"1조519",
		},
		{
			"1000000001000",
			args{1000000001000},
			"1조1000",
		},
		{
			"1000000001001",
			args{1000000001001},
			"1조1001",
		},
		{
			"1000000001005",
			args{1000000001005},
			"1조1005",
		},
		{
			"1000000001010",
			args{1000000001010},
			"1조1010",
		},
		{
			"1000000001015",
			args{1000000001015},
			"1조1015",
		},
		{
			"1000000001060",
			args{1000000001060},
			"1조1060",
		},
		{
			"1000000001064",
			args{1000000001064},
			"1조1064",
		},
		{
			"1000000001100",
			args{1000000001100},
			"1조1100",
		},
		{
			"1000000001104",
			args{1000000001104},
			"1조1104",
		},
		{
			"1000000001140",
			args{1000000001140},
			"1조1140",
		},
		{
			"1000000001123",
			args{1000000001123},
			"1조1123",
		},
		{
			"1000000001200",
			args{1000000001200},
			"1조1200",
		},
		{
			"1000000001201",
			args{1000000001201},
			"1조1201",
		},
		{
			"1000000001204",
			args{1000000001204},
			"1조1204",
		},
		{
			"1000000001240",
			args{1000000001240},
			"1조1240",
		},
		{
			"1000000001234",
			args{1000000001234},
			"1조1234",
		},
		{
			"1000000006000",
			args{1000000006000},
			"1조6000",
		},
		{
			"1000000006001",
			args{1000000006001},
			"1조6001",
		},
		{
			"1000000006005",
			args{1000000006005},
			"1조6005",
		},
		{
			"1000000006010",
			args{1000000006010},
			"1조6010",
		},
		{
			"1000000006015",
			args{1000000006015},
			"1조6015",
		},
		{
			"1000000006060",
			args{1000000006060},
			"1조6060",
		},
		{
			"1000000006064",
			args{1000000006064},
			"1조6064",
		},
		{
			"1000000006100",
			args{1000000006100},
			"1조6100",
		},
		{
			"1000000006104",
			args{1000000006104},
			"1조6104",
		},
		{
			"1000000006140",
			args{1000000006140},
			"1조6140",
		},
		{
			"1000000006123",
			args{1000000006123},
			"1조6123",
		},
		{
			"1000000006200",
			args{1000000006200},
			"1조6200",
		},
		{
			"1000000006201",
			args{1000000006201},
			"1조6201",
		},
		{
			"1000000006204",
			args{1000000006204},
			"1조6204",
		},
		{
			"1000000006240",
			args{1000000006240},
			"1조6240",
		},
		{
			"1000000006234",
			args{1000000006234},
			"1조6234",
		},
		{
			"1000000010000",
			args{1000000010000},
			"1조1만",
		},
		{
			"1000000020000",
			args{1000000020000},
			"1조2만",
		},
		{
			"1000011110000",
			args{1000011110000},
			"1조1111만",
		},
		{
			"1000011111111",
			args{1000011111111},
			"1조1111만1111",
		},
		{
			"1000010101010",
			args{1000010101010},
			"1조1010만1010",
		},
		{
			"1000010101012",
			args{1000010101012},
			"1조1010만1012",
		},
		{
			"1000099990000",
			args{1000099990000},
			"1조9999만",
		},
		{
			"1000099999999",
			args{1000099999999},
			"1조9999만9999",
		},
		{
			"1000100000000",
			args{1000100000000},
			"1조1억",
		},
		{
			"1000100010000",
			args{1000100010000},
			"1조1억1만",
		},
		{
			"1000111110000",
			args{1000111110000},
			"1조1억1111만",
		},
		{
			"1000111111111",
			args{1000111111111},
			"1조1억1111만1111",
		},
		{
			"1000110101010",
			args{1000110101010},
			"1조1억1010만1010",
		},
		{
			"1000110101012",
			args{1000110101012},
			"1조1억1010만1012",
		},
		{
			"1000199990000",
			args{1000199990000},
			"1조1억9999만",
		},
		{
			"1000199999999",
			args{1000199999999},
			"1조1억9999만9999",
		},
		{
			"1000200000000",
			args{1000200000000},
			"1조2억",
		},
		{
			"1000200010000",
			args{1000200010000},
			"1조2억1만",
		},
		{
			"1000211110000",
			args{1000211110000},
			"1조2억1111만",
		},
		{
			"1000211111111",
			args{1000211111111},
			"1조2억1111만1111",
		},
		{
			"1000210101010",
			args{1000210101010},
			"1조2억1010만1010",
		},
		{
			"1000210101012",
			args{1000210101012},
			"1조2억1010만1012",
		},
		{
			"1000299990000",
			args{1000299990000},
			"1조2억9999만",
		},
		{
			"1000299999999",
			args{1000299999999},
			"1조2억9999만9999",
		},
		{
			"1111111111111",
			args{1111111111111},
			"1조1111억1111만1111",
		},
		{
			"1606060606060",
			args{1606060606060},
			"1조6060억6060만6060",
		},
		{
			"1606060606066",
			args{1606060606066},
			"1조6060억6060만6066",
		},
		{
			"1666600000000",
			args{1666600000000},
			"1조6666억",
		},
		{
			"9876543210123456",
			args{9876543210123456},
			"9876조5432억1012만3456",
		},
		{
			"9999999999999999",
			args{9999999999999999},
			"9999조9999억9999만9999",
		},
		{
			"999999999999999999",
			args{999999999999999999},
			"999999조9999억9999만9999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := KoreanPluralizeUnitType2(tt.args.number); gotOut != tt.wantOut {
				t.Errorf("KoreanPluralizeUnitType2() = %v, want %v", gotOut, tt.wantOut)
			} else {
				t.Logf("KoreanPluralizeUnitType2() = %v", gotOut)
			}
		})
	}
}

func TestKoreanPluralizeUnitType3(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
	}{
		{
			"0",
			args{0},
			"0",
		},
		{
			"1",
			args{1},
			"1",
		},
		{
			"9",
			args{9},
			"9",
		},
		{
			"10",
			args{10},
			"10",
		},
		{
			"12",
			args{12},
			"12",
		},
		{
			"23",
			args{23},
			"23",
		},
		{
			"40",
			args{40},
			"40",
		},
		{
			"100",
			args{100},
			"100",
		},
		{
			"106",
			args{106},
			"106",
		},
		{
			"123",
			args{123},
			"123",
		},
		{
			"170",
			args{170},
			"170",
		},
		{
			"234",
			args{234},
			"234",
		},
		{
			"500",
			args{500},
			"500",
		},
		{
			"509",
			args{509},
			"509",
		},
		{
			"519",
			args{519},
			"519",
		},
		{
			"1000",
			args{1000},
			"1000",
		},
		{
			"1001",
			args{1001},
			"1001",
		},
		{
			"1005",
			args{1005},
			"1005",
		},
		{
			"1010",
			args{1010},
			"1010",
		},
		{
			"1015",
			args{1015},
			"1015",
		},
		{
			"1060",
			args{1060},
			"1060",
		},
		{
			"1064",
			args{1064},
			"1064",
		},
		{
			"1100",
			args{1100},
			"1100",
		},
		{
			"1104",
			args{1104},
			"1104",
		},
		{
			"1140",
			args{1140},
			"1140",
		},
		{
			"1123",
			args{1123},
			"1123",
		},
		{
			"1200",
			args{1200},
			"1200",
		},
		{
			"1201",
			args{1201},
			"1201",
		},
		{
			"1204",
			args{1204},
			"1204",
		},
		{
			"1240",
			args{1240},
			"1240",
		},
		{
			"1234",
			args{1234},
			"1234",
		},

		{
			"6000",
			args{6000},
			"6000",
		},
		{
			"6001",
			args{6001},
			"6001",
		},
		{
			"6005",
			args{6005},
			"6005",
		},
		{
			"6010",
			args{6010},
			"6010",
		},
		{
			"6015",
			args{6015},
			"6015",
		},
		{
			"6060",
			args{6060},
			"6060",
		},
		{
			"6064",
			args{6064},
			"6064",
		},
		{
			"6100",
			args{6100},
			"6100",
		},
		{
			"6104",
			args{6104},
			"6104",
		},
		{
			"6140",
			args{6140},
			"6140",
		},
		{
			"6123",
			args{6123},
			"6123",
		},
		{
			"6200",
			args{6200},
			"6200",
		},
		{
			"6201",
			args{6201},
			"6201",
		},
		{
			"6204",
			args{6204},
			"6204",
		},
		{
			"6240",
			args{6240},
			"6240",
		},
		{
			"6234",
			args{6234},
			"6234",
		},
		{
			"10000",
			args{10000},
			"1만",
		},
		{
			"20000",
			args{20000},
			"2만",
		},
		{
			"11110000",
			args{11110000},
			"천백십1만",
		},
		{
			"11111111",
			args{11111111},
			"천백십1만1111",
		},
		{
			"10101010",
			args{10101010},
			"천십만1010",
		},
		{
			"10101012",
			args{10101012},
			"천십만1012",
		},
		{
			"99990000",
			args{99990000},
			"9천9백9십9만",
		},
		{
			"99999999",
			args{99999999},
			"9천9백9십9만9999",
		},
		{
			"100000000",
			args{100000000},
			"1억",
		},
		{
			"100010000",
			args{100010000},
			"1억1만",
		},
		{
			"111110000",
			args{111110000},
			"1억천백십1만",
		},
		{
			"111111111",
			args{111111111},
			"1억천백십1만1111",
		},
		{
			"110101010",
			args{110101010},
			"1억천십만1010",
		},
		{
			"110101012",
			args{110101012},
			"1억천십만1012",
		},
		{
			"199990000",
			args{199990000},
			"1억9천9백9십9만",
		},
		{
			"199999999",
			args{199999999},
			"1억9천9백9십9만9999",
		},
		{
			"200000000",
			args{200000000},
			"2억",
		},
		{
			"200010000",
			args{200010000},
			"2억1만",
		},
		{
			"211110000",
			args{211110000},
			"2억천백십1만",
		},
		{
			"211111111",
			args{211111111},
			"2억천백십1만1111",
		},
		{
			"210101010",
			args{210101010},
			"2억천십만1010",
		},
		{
			"210101012",
			args{210101012},
			"2억천십만1012",
		},
		{
			"299990000",
			args{299990000},
			"2억9천9백9십9만",
		},
		{
			"299999999",
			args{299999999},
			"2억9천9백9십9만9999",
		},
		{
			"606060606060",
			args{606060606060},
			"6천6십억6천6십만6060",
		},
		{
			"606060606066",
			args{606060606066},
			"6천6십억6천6십만6066",
		},
		{
			"666600000000",
			args{666600000000},
			"6천6백6십6억",
		},
		{
			"1000000000000",
			args{1000000000000},
			"1조",
		},
		{
			"1000000000001",
			args{1000000000001},
			"1조1",
		},
		{
			"1000000000009",
			args{1000000000009},
			"1조9",
		},
		{
			"1000000000010",
			args{1000000000010},
			"1조10",
		},
		{
			"1000000000012",
			args{1000000000012},
			"1조12",
		},
		{
			"1000000000023",
			args{1000000000023},
			"1조23",
		},
		{
			"1000000000040",
			args{1000000000040},
			"1조40",
		},
		{
			"1000000000100",
			args{1000000000100},
			"1조100",
		},
		{
			"1000000000106",
			args{1000000000106},
			"1조106",
		},
		{
			"1000000000123",
			args{1000000000123},
			"1조123",
		},
		{
			"1000000000170",
			args{1000000000170},
			"1조170",
		},
		{
			"1000000000234",
			args{1000000000234},
			"1조234",
		},
		{
			"1000000000500",
			args{1000000000500},
			"1조500",
		},
		{
			"1000000000509",
			args{1000000000509},
			"1조509",
		},
		{
			"1000000000519",
			args{1000000000519},
			"1조519",
		},
		{
			"1000000001000",
			args{1000000001000},
			"1조1000",
		},
		{
			"1000000001001",
			args{1000000001001},
			"1조1001",
		},
		{
			"1000000001005",
			args{1000000001005},
			"1조1005",
		},
		{
			"1000000001010",
			args{1000000001010},
			"1조1010",
		},
		{
			"1000000001015",
			args{1000000001015},
			"1조1015",
		},
		{
			"1000000001060",
			args{1000000001060},
			"1조1060",
		},
		{
			"1000000001064",
			args{1000000001064},
			"1조1064",
		},
		{
			"1000000001100",
			args{1000000001100},
			"1조1100",
		},
		{
			"1000000001104",
			args{1000000001104},
			"1조1104",
		},
		{
			"1000000001140",
			args{1000000001140},
			"1조1140",
		},
		{
			"1000000001123",
			args{1000000001123},
			"1조1123",
		},
		{
			"1000000001200",
			args{1000000001200},
			"1조1200",
		},
		{
			"1000000001201",
			args{1000000001201},
			"1조1201",
		},
		{
			"1000000001204",
			args{1000000001204},
			"1조1204",
		},
		{
			"1000000001240",
			args{1000000001240},
			"1조1240",
		},
		{
			"1000000001234",
			args{1000000001234},
			"1조1234",
		},
		{
			"1000000006000",
			args{1000000006000},
			"1조6000",
		},
		{
			"1000000006001",
			args{1000000006001},
			"1조6001",
		},
		{
			"1000000006005",
			args{1000000006005},
			"1조6005",
		},
		{
			"1000000006010",
			args{1000000006010},
			"1조6010",
		},
		{
			"1000000006015",
			args{1000000006015},
			"1조6015",
		},
		{
			"1000000006060",
			args{1000000006060},
			"1조6060",
		},
		{
			"1000000006064",
			args{1000000006064},
			"1조6064",
		},
		{
			"1000000006100",
			args{1000000006100},
			"1조6100",
		},
		{
			"1000000006104",
			args{1000000006104},
			"1조6104",
		},
		{
			"1000000006140",
			args{1000000006140},
			"1조6140",
		},
		{
			"1000000006123",
			args{1000000006123},
			"1조6123",
		},
		{
			"1000000006200",
			args{1000000006200},
			"1조6200",
		},
		{
			"1000000006201",
			args{1000000006201},
			"1조6201",
		},
		{
			"1000000006204",
			args{1000000006204},
			"1조6204",
		},
		{
			"1000000006240",
			args{1000000006240},
			"1조6240",
		},
		{
			"1000000006234",
			args{1000000006234},
			"1조6234",
		},
		{
			"1000000010000",
			args{1000000010000},
			"1조1만",
		},
		{
			"1000000020000",
			args{1000000020000},
			"1조2만",
		},
		{
			"1000011110000",
			args{1000011110000},
			"1조천백십1만",
		},
		{
			"1000011111111",
			args{1000011111111},
			"1조천백십1만1111",
		},
		{
			"1000010101010",
			args{1000010101010},
			"1조천십만1010",
		},
		{
			"1000010101012",
			args{1000010101012},
			"1조천십만1012",
		},
		{
			"1000099990000",
			args{1000099990000},
			"1조9천9백9십9만",
		},
		{
			"1000099999999",
			args{1000099999999},
			"1조9천9백9십9만9999",
		},
		{
			"1000100000000",
			args{1000100000000},
			"1조1억",
		},
		{
			"1000100010000",
			args{1000100010000},
			"1조1억1만",
		},
		{
			"1000111110000",
			args{1000111110000},
			"1조1억천백십1만",
		},
		{
			"1000111111111",
			args{1000111111111},
			"1조1억천백십1만1111",
		},
		{
			"1000110101010",
			args{1000110101010},
			"1조1억천십만1010",
		},
		{
			"1000110101012",
			args{1000110101012},
			"1조1억천십만1012",
		},
		{
			"1000199990000",
			args{1000199990000},
			"1조1억9천9백9십9만",
		},
		{
			"1000199999999",
			args{1000199999999},
			"1조1억9천9백9십9만9999",
		},
		{
			"1000200000000",
			args{1000200000000},
			"1조2억",
		},
		{
			"1000200010000",
			args{1000200010000},
			"1조2억1만",
		},
		{
			"1000211110000",
			args{1000211110000},
			"1조2억천백십1만",
		},
		{
			"1000211111111",
			args{1000211111111},
			"1조2억천백십1만1111",
		},
		{
			"1000210101010",
			args{1000210101010},
			"1조2억천십만1010",
		},
		{
			"1000210101012",
			args{1000210101012},
			"1조2억천십만1012",
		},
		{
			"1000299990000",
			args{1000299990000},
			"1조2억9천9백9십9만",
		},
		{
			"1000299999999",
			args{1000299999999},
			"1조2억9천9백9십9만9999",
		},
		{
			"1111111111111",
			args{1111111111111},
			"1조천백십1억천백십1만1111",
		},
		{
			"1606060606060",
			args{1606060606060},
			"1조6천6십억6천6십만6060",
		},
		{
			"1606060606066",
			args{1606060606066},
			"1조6천6십억6천6십만6066",
		},
		{
			"1666600000000",
			args{1666600000000},
			"1조6천6백6십6억",
		},
		{
			"9876543210123456",
			args{9876543210123456},
			"9천8백7십6조5천4백3십2억천십2만3456",
		},
		{
			"9999999999999999",
			args{9999999999999999},
			"9천9백9십9조9천9백9십9억9천9백9십9만9999",
		},
		{
			"999999999999999999",
			args{999999999999999999},
			"9십99천9백9십9조9천9백9십9억9천9백9십9만9999",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := KoreanPluralizeUnitType3(tt.args.number); gotOut != tt.wantOut {
				t.Errorf("KoreanPluralizeUnitType3() = %v, want %v", gotOut, tt.wantOut)
			} else {
				t.Logf("KoreanPluralizeUnitType3() = %v", gotOut)
			}
		})
	}
}

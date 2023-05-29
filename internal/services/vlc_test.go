package services

import (
	"testing"
)

func Test_prepareText(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected string
	}{
		{
			name:     "test #1",
			str:      "My test string",
			expected: "!my test string",
		},
		{
			name:     "test #2",
			str:      "My TeSt strinG",
			expected: "!my !te!st strin!g",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r := prepareBeforeCompressText(tt.str); r != tt.expected {
				t.Errorf("bin() = %v, expected %v", r, tt.expected)
			}
		})
	}
}

func Test_bin(t *testing.T) {
	tests := []struct {
		name     string
		char     rune
		expected string
	}{
		{
			name:     "test #1",
			char:     'y',
			expected: "0000001",
		},
		{
			name:     "test #2",
			char:     'a',
			expected: "011",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r := bin(tt.char); r != tt.expected {
				t.Errorf("bin() = %v, expected %v", r, tt.expected)
			}
		})
	}
}

func Test_encodeBinary(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected string
	}{
		{
			name:     "test #1",
			str:      "tat",
			expected: "10010111001",
		},
		{
			name:     "test #2",
			str:      "tat ",
			expected: "1001011100111",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r := encodeBinary(tt.str); r != tt.expected {
				t.Errorf("encodeBinary() = %v, expected %v", r, tt.expected)
			}
		})
	}
}

func Test_Encode(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected string
	}{
		{
			name:     "test #1",
			str:      "Mty test strinG",
			expected: "20 39 03 CD 59 D6 50 98 10 10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r := Encode(tt.str); r != tt.expected {
				t.Errorf("Encode() = %v, expected %v", r, tt.expected)
			}
		})
	}
}

func Test_prepareBeforeUncompressText(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want string
	}{
		{
			name: "test #1",
			str:  "!my test",
			want: "My test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prepareBeforeUncompressText(tt.str); got != tt.want {
				t.Errorf("prepareBeforeUncompressText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	type args struct {
		encodedString string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test #1",
			args: args{
				"20 39 03 CD 59 D6 50 98 10 10",
			},
			want: "Mty test strinG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decode(tt.args.encodedString); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

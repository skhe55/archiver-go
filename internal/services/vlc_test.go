package services

import (
	"reflect"
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
			if r := prepareText(tt.str); r != tt.expected {
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

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}

	tests := []struct {
		name     string
		args     args
		expected BinaryChunks
	}{
		{
			name: "test #1",
			args: args{
				"1000100010001000",
				chunkSize,
			},
			expected: BinaryChunks{"10001000", "10001000"},
		},
		{
			name: "test #2",
			args: args{
				"100010001000100011",
				chunkSize,
			},
			expected: BinaryChunks{"10001000", "10001000", "11000000"},
		},
		{
			name: "test #3",
			args: args{
				"1001",
				chunkSize,
			},
			expected: BinaryChunks{"10010000"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r := splitByChunks(tt.args.bStr, tt.args.chunkSize); !reflect.DeepEqual(tt.expected, r) {
				t.Errorf("splitByChunks() = %v, expected %v", r, tt.expected)
			}
		})
	}
}

func Test_chunkToHex(t *testing.T) {
	tests := []struct {
		name     string
		bChunk   BinaryChunk
		expected HexChunk
	}{
		{
			name:     "test #1",
			bChunk:   BinaryChunk("11110000"),
			expected: HexChunk("F0"),
		}, {
			name:     "test #2",
			bChunk:   BinaryChunk("10000000"),
			expected: HexChunk("80"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r := tt.bChunk.ToHex(); r != tt.expected {
				t.Errorf("chunkToHex = %v, expected %v", r, tt.expected)
			}
		})
	}
}

func Test_chunksToHex(t *testing.T) {
	tests := []struct {
		name     string
		bChunks  BinaryChunks
		expected HexChunks
	}{
		{
			name:     "test #1",
			bChunks:  BinaryChunks{"10000000", "11110000"},
			expected: HexChunks{"80", "F0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r := tt.bChunks.ToHex(); !reflect.DeepEqual(tt.expected, r) {
				t.Errorf("chunksToHex() = %v, expected %v", r, tt.expected)
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

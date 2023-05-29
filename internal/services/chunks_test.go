package services

import (
	"reflect"
	"testing"
)

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

func Test_newHexChunks(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected HexChunks
	}{
		{
			name:     "test #1",
			str:      "20 39 03 CD 59",
			expected: HexChunks{"20", "39", "03", "CD", "59"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r := NewHexChunks(tt.str); !reflect.DeepEqual(tt.expected, r) {
				t.Errorf("newHexChunks() = %v, expected %v", r, tt.expected)
			}
		})
	}
}

func Test_chunkToBinary(t *testing.T) {
	tests := []struct {
		name     string
		hChunk   HexChunk
		expected BinaryChunk
	}{
		{
			name:     "test #1",
			hChunk:   HexChunk("80"),
			expected: BinaryChunk("10000000"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r := tt.hChunk.ToBinary(); r != tt.expected {
				t.Errorf("chunkToBinary() = %v, expected %v", r, tt.expected)
			}
		})
	}
}

func Test_chunksToBinary(t *testing.T) {
	tests := []struct {
		name     string
		hChunks  HexChunks
		expected BinaryChunks
	}{
		{
			name:     "test #1",
			hChunks:  HexChunks{"80", "F0"},
			expected: BinaryChunks{"10000000", "11110000"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r := tt.hChunks.ToBinary(); !reflect.DeepEqual(tt.expected, r) {
				t.Errorf("chunksToBinary() = %v, expected %v", r, tt.expected)
			}
		})
	}
}

func Test_Join(t *testing.T) {
	tests := []struct {
		name     string
		bChunks  BinaryChunks
		expected string
	}{
		{
			name:     "test #1",
			bChunks:  BinaryChunks{"10000000", "11110000"},
			expected: "1000000011110000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r := tt.bChunks.Join(); r != tt.expected {
				t.Errorf("Join() = %v, expected %v", r, tt.expected)
			}
		})
	}
}

package vlc

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

func TestNewBinChunks(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want BinaryChunks
	}{
		{
			name: "test #1",
			args: args{
				data: []byte{20, 30, 60, 18},
			},
			want: BinaryChunks{"00010100", "00011110", "00111100", "00010010"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinChunks(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

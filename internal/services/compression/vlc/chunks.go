package vlc

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type BinaryChunk string

const chunkSize = 8

// joins chunks into one line and returns string
func (bChunks BinaryChunks) Join() string {
	var buf strings.Builder

	for _, chunk := range bChunks {
		buf.WriteString(string(chunk))
	}

	return buf.String()
}

func (bChunks BinaryChunks) Bytes() []byte {
	res := make([]byte, 0, len(bChunks))

	for _, bc := range bChunks {
		res = append(res, bc.Byte())
	}

	return res
}

func (bChunk BinaryChunk) Byte() byte {
	num, err := strconv.ParseUint(string(bChunk), 2, chunkSize)
	if err != nil {
		panic("can't parse binary chunk: " + err.Error())
	}

	return byte(num)
}

func NewBinChunks(data []byte) BinaryChunks {

	res := make(BinaryChunks, 0, len(data))

	for _, part := range data {
		res = append(res, NewBinChunk(part))
	}

	return res
}

func NewBinChunk(code byte) BinaryChunk {
	return BinaryChunk(fmt.Sprintf("%08b", code))
}

// splitByChunks splits binary string by chunks with given size,
// i.g: '1001100110011001' -> '10011001 10011001'
func splitByChunks(bStr string, chunkSize int) BinaryChunks {
	strLen := utf8.RuneCountInString(bStr)

	chunksCount := strLen / chunkSize

	if strLen%chunkSize != 0 {
		chunksCount++
	}

	res := make(BinaryChunks, 0, chunksCount)

	var buf strings.Builder

	for i, char := range bStr {
		buf.WriteString(string(char))

		if (i+1)%chunkSize == 0 {
			res = append(res, BinaryChunk(buf.String()))
			buf.Reset()
		}
	}

	if buf.Len() != 0 {
		lastChunk := buf.String()

		lastChunk += strings.Repeat("0", chunkSize-len(lastChunk))

		res = append(res, BinaryChunk(lastChunk))
	}

	return res
}

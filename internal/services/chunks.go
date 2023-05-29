package services

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type HexChunks []HexChunk

type BinaryChunk string

type HexChunk string

const chunkSize = 8

const chunksSeparator = " "

func (bChunks BinaryChunks) ToHex() HexChunks {
	res := make(HexChunks, 0, len(bChunks))

	for _, chunk := range bChunks {
		hexChunk := chunk.ToHex()

		res = append(res, hexChunk)
	}

	return res
}

func (bChunk BinaryChunk) ToHex() HexChunk {
	num, err := strconv.ParseUint(string(bChunk), 2, chunkSize)
	if err != nil {
		panic("can't parse binary chunk" + err.Error())
	}

	res := strings.ToUpper(fmt.Sprintf("%x", num))

	if len(res) == 1 {
		res = "0" + res
	}

	return HexChunk(res)
}

func (hChunks HexChunks) ToBinary() BinaryChunks {
	res := make(BinaryChunks, 0, len(hChunks))

	for _, chunk := range hChunks {
		bChunk := chunk.ToBinary()
		res = append(res, bChunk)
	}

	return res
}

func (hChunk HexChunk) ToBinary() BinaryChunk {
	num, err := strconv.ParseUint(string(hChunk), 16, chunkSize)
	if err != nil {
		panic("can't parse hex chunk" + err.Error())
	}

	res := fmt.Sprintf("%08b", num)

	return BinaryChunk(res)
}

func (hChunks HexChunks) ToString() string {
	switch len(hChunks) {
	case 0:
		return ""
	case 1:
		return string(hChunks[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hChunks[0]))

	for _, hChunk := range hChunks[1:] {
		buf.WriteString(chunksSeparator)
		buf.WriteString(string(hChunk))
	}

	return buf.String()
}

// joins chunks into one line and returns string
func (bChunks BinaryChunks) Join() string {
	var buf strings.Builder

	for _, chunk := range bChunks {
		buf.WriteString(string(chunk))
	}

	return buf.String()
}

func NewHexChunks(str string) HexChunks {
	parts := strings.Split(str, chunksSeparator)

	res := make(HexChunks, 0, len(parts))

	for _, part := range parts {
		res = append(res, HexChunk(part))
	}

	return res
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

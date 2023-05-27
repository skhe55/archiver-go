package services

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type BinaryChunks []BinaryChunk

type HexChunks []HexChunk

type BinaryChunk string

type HexChunk string

var encodingTable = map[rune]string{
	' ': "11",
	't': "1001",
	'n': "10000",
	's': "0101",
	'r': "01000",
	'd': "00101",
	'!': "001000",
	'c': "000101",
	'm': "000011",
	'g': "0000100",
	'b': "0000010",
	'v': "00000001",
	'k': "0000000001",
	'q': "000000000001",
	'e': "101",
	'o': "10001",
	'a': "011",
	'i': "01001",
	'h': "0011",
	'l': "001001",
	'u': "00011",
	'f': "000100",
	'p': "0000101",
	'w': "0000011",
	'y': "0000001",
	'j': "000000001",
	'x': "00000000001",
	'z': "000000000000",
}

const chunkSize = 8

func Encode(str string) string {
	str = prepareText(str)

	chunks := splitByChunks(encodeBinary(str), chunkSize)

	return chunks.ToHex().ToString()
}

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

func (hChunks HexChunks) ToString() string {
	const sep = " "

	switch len(hChunks) {
	case 0:
		return ""
	case 1:
		return string(hChunks[0])
	}

	var buf strings.Builder

	buf.WriteString(string(hChunks[0]))

	for _, hChunk := range hChunks[1:] {
		buf.WriteString(sep)
		buf.WriteString(string(hChunk))
	}

	return buf.String()
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

func prepareText(str string) string {
	var buf strings.Builder

	for _, char := range str {
		if unicode.IsUpper(char) {
			buf.WriteRune('!')
			buf.WriteRune(unicode.ToLower(char))
		} else {
			buf.WriteRune(char)
		}
	}

	return buf.String()
}

// encodeBinary encodes str into binary codes string w/o spaces
func encodeBinary(str string) string {
	var buf strings.Builder

	for _, char := range str {
		buf.WriteString(bin(char))
	}

	return buf.String()
}

// get binary code in string format from encoding table
func bin(char rune) string {
	table := getEncodingTable()

	res, ok := table[char]
	if !ok {
		panic("unknow character: " + string(char))
	}

	return res
}

func getEncodingTable() map[rune]string {
	return encodingTable
}

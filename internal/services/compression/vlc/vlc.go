package vlc

import (
	"strings"
	"unicode"
)

var encodingTable = EncodingTable{
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

type EncoderDecoder struct{}

func New() EncoderDecoder {
	return EncoderDecoder{}
}

func (_ EncoderDecoder) Encode(str string) []byte {
	str = prepareBeforeCompressText(str)

	chunks := splitByChunks(encodeBinary(str), chunkSize)

	return chunks.Bytes()
}

func (_ EncoderDecoder) Decode(encodedData []byte) string {
	bString := NewBinChunks(encodedData).Join()

	dTree := getEncodingTable().DecodingTree()

	return prepareBeforeUncompressText(dTree.Decode(bString))
}

func prepareBeforeCompressText(str string) string {
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

func prepareBeforeUncompressText(str string) string {
	var buf strings.Builder

	for i := 0; i < len(str); i++ {
		if str[i] == '!' {
			buf.WriteRune(unicode.ToUpper(rune(str[i+1])))
			i++
		} else {
			buf.WriteByte(str[i])
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

func getEncodingTable() EncodingTable {
	return encodingTable
}

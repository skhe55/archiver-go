package services

import "strings"

type DecodingTree struct {
	Value string
	Left  *DecodingTree
	Right *DecodingTree
}

type EncodingTable map[rune]string

func (et EncodingTable) DecodingTree() DecodingTree {
	res := DecodingTree{}

	for ch, code := range et {
		res.add(code, ch)
	}

	return res
}

func (dt *DecodingTree) add(code string, value rune) {
	currNode := dt

	for _, ch := range code {
		switch ch {
		case '0':
			if currNode.Left == nil {
				currNode.Left = &DecodingTree{}
			}
			currNode = currNode.Left
		case '1':
			if currNode.Right == nil {
				currNode.Right = &DecodingTree{}
			}
			currNode = currNode.Right
		}
	}

	currNode.Value = string(value)
}

func (dt *DecodingTree) Decode(str string) string {
	var buf strings.Builder

	currNode := dt

	for _, ch := range str {
		if currNode.Value != "" {
			buf.WriteString(currNode.Value)
			currNode = dt
		}
		switch ch {
		case '0':
			currNode = currNode.Left
		case '1':
			currNode = currNode.Right
		}
	}

	if currNode.Value != "" {
		buf.WriteString(currNode.Value)
	}

	return buf.String()
}

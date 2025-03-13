package huffmango

import (
	"fmt"
	"strconv"
)

func CompressWithMapCode(mapCode map[rune]string, originalString string) ([]byte, int) {
	var strEncodedString string

	for _, rune := range originalString {
		strEncodedString += mapCode[rune]
	}

	encodedStr := binaryStringToByteSlice(strEncodedString)

	return encodedStr, len(originalString)
}

func PrintCompressWithMapCode(mapCode map[rune]string, originalString string) {
	var strEncodedString string

	for _, rune := range originalString {
		strEncodedString += mapCode[rune]
	}

	println(strEncodedString)

	encodedStr := binaryStringToByteSlice(strEncodedString)

	fmt.Printf("%v\n", encodedStr)
}

func printCode(root node, steps []int) {
	if root.isLeaf() {
		var code string
		for _, s := range steps {
			code += strconv.Itoa(s)
		}
		fmt.Printf("%v %v\n", root.char, code)
		return
	}

	if root.left != nil {
		steps = append(steps, 0)
		printCode(*root.left, steps)
	}

	if root.right != nil {
		steps = append(steps[:len(steps)-1], 1)
		printCode(*root.right, steps)
	}
}

func fillMapCode(root node, steps []int, mapCode map[rune]string) {
	if root.isLeaf() {
		var code string
		for _, s := range steps {
			code += strconv.Itoa(s)
		}
		mapCode[root.char] = code
		return
	}

	if root.left != nil {
		steps = append(steps, 0)
		fillMapCode(*root.left, steps, mapCode)
	}

	if root.right != nil {
		steps = append(steps[:len(steps)-1], 1)
		fillMapCode(*root.right, steps, mapCode)
	}
}

func CreateHuffmanMapCodeFromString(s string) map[rune]string {
	nodes := makeNodes(s)

	for len(nodes) > 1 {
		nodes = iteration(nodes)
	}

	mapCode := make(map[rune]string)

	fillMapCode(nodes[0], []int{}, mapCode)

	return mapCode
}

func PrintHuffmanCodeFromString(s string) {
	nodes := makeNodes(s)

	for len(nodes) > 1 {
		nodes = iteration(nodes)
	}

	printCode(nodes[0], []int{})
}

func decode(treeRoot *node, encodedData []byte, ol int) string {
	var decodedData string

	ref := treeRoot
	leafCount := 0

begin:
	for _, bytes := range encodedData {
		bitStr := fmt.Sprintf("%08b", bytes)

		for _, rune := range bitStr {
			if ref.isLeaf() {
				leafCount++
				decodedData += string(ref.char)

				if leafCount == ol {
					break begin
				}

				ref = treeRoot
			}
			switch rune {
			case '0':
				if ref.left != nil {
					ref = ref.left
				}
			case '1':
				if ref.right != nil {
					ref = ref.right
				}
			}
		}
	}

	return decodedData
}

func DecodeHuffmanFromMapCode(mapCode map[rune]string, bs []byte, ol int) string {
	root := makeTreeFromCodeMap(mapCode)

	decodeString := decode(root, bs, ol)

	return decodeString
}

func PrintDecodeHuffmanFromMapCode(mapCode map[rune]string, bs []byte, ol int) {
	root := makeTreeFromCodeMap(mapCode)

	decodeString := decode(root, bs, ol)

	fmt.Printf("%v\n", decodeString)
}

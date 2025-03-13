package huffmango

import "strconv"

func parseBinToHex(s string) uint64 {
	ui, err := strconv.ParseUint(s, 2, 64)
	if err != nil {
		panic(err)
	}

	return ui
}

func binaryStringToByteSlice(binaryString string) []byte {
	var byteSlice []byte
	for len(binaryString) >= 8 {
		byteSlice = append(byteSlice, byte(parseBinToHex(binaryString[:8])))
		binaryString = binaryString[8:]
		if len(binaryString) < 8 {
			left := 8 - len(binaryString)
			for left > 8 {
				binaryString += "0"
				left--
			}
			byteSlice = append(byteSlice, byte(parseBinToHex(binaryString)))
			break
		}
		if len(binaryString) == 8 {
			byteSlice = append(byteSlice, byte(parseBinToHex(binaryString)))
			break
		}
	}

	return byteSlice
}

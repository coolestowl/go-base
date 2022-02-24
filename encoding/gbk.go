package encoding

import (
	"golang.org/x/text/encoding/simplifiedchinese"
)

func IsGBK(data []byte) bool {
	length := len(data)
	var i int = 0
	for i < length {
		if data[i] <= 0x7f {
			i++
			continue
		} else {
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[i+1] >= 0x40 &&
				data[i+1] <= 0xfe &&
				data[i+1] != 0xf7 {
				i += 2
				continue
			} else {
				return false
			}
		}
	}
	return true
}

func GBKToUTF8(data []byte) ([]byte, error) {
	return simplifiedchinese.GBK.NewDecoder().Bytes(data)
}

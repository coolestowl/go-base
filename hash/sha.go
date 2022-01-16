package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func SHA256(src []byte) []byte {
	h := sha256.New()
	h.Write(src)
	return h.Sum(nil)
}

func HexSHA256(src []byte) string {
	return hex.EncodeToString(SHA256(src))
}

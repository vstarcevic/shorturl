package app

import (
	"bytes"
)

const base58_chars = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func EncodeBase58(num int64) string {
	var b bytes.Buffer
	for num > 0 {
		remainder := num % 58
		b.WriteByte(base58_chars[remainder])
		num /= 58
	}

	reverseString(b.Bytes())
	return b.String()
}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

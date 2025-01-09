package util

import (
	"bytes"
	"crypto/rand"
	"math/big"
	rand2 "math/rand"
	"strings"
)

const (
	LOWER_CASE_LETTERS = "abcdefghijklmnopqrstuvwxyz"
	UPPER_CASE_LETTERS = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NUMBERS            = "0123456789"
	SYMBOLS            = "!@#$%^&*()_+-=[]{};':\"\\|,.<>/?"
)

func RandString(length int, charsets ...string) string {
	buf := bytes.Buffer{}
	if len(charsets) == 0 {
		buf.WriteString(LOWER_CASE_LETTERS)
		buf.WriteString(UPPER_CASE_LETTERS)
		buf.WriteString(NUMBERS)
		buf.WriteString(SYMBOLS)
	} else {
		for _, charset := range charsets {
			buf.WriteString(charset)
		}
	}
	charBytes := buf.Bytes()
	maxIdx := big.NewInt(int64(len(charBytes)))
	res := strings.Builder{}
	for i := 0; i < length; i++ {
		idx, _ := rand.Int(rand.Reader, maxIdx)
		res.WriteByte(charBytes[idx.Int64()])
	}
	return res.String()
}

// RandStringImportAllType 生成随机字符串的函数，确保至少包含一个输入charset字符
func RandStringImportAllType(n int, charsets ...string) string {
	buf := bytes.Buffer{}
	if len(charsets) == 0 {
		charsets = []string{LOWER_CASE_LETTERS, UPPER_CASE_LETTERS, NUMBERS, SYMBOLS}
	}
	if n < len(charsets) {
		n = len(charsets)
	}
	for _, charset := range charsets {
		buf.WriteString(charset)
	}
	allChars := buf.Bytes()
	b := make([]byte, n)
	// 预先选择一个字符集中的字符，确保至少包含每种类型的字符
	for i, charset := range charsets {
		maxIdx := big.NewInt(int64(len(charset)))
		idx, _ := rand.Int(rand.Reader, maxIdx)
		b[i] = charset[idx.Int64()]
	}

	// 随机分配其余位置的字符
	maxIdx := big.NewInt(int64(len(allChars)))
	for i := len(charsets); i < n; i++ {
		idx, _ := rand.Int(rand.Reader, maxIdx)
		b[i] = allChars[idx.Int64()]
	}

	// 对结果进行打乱，以确保分布是随机的
	rand2.Shuffle(len(b), func(i, j int) {
		b[i], b[j] = b[j], b[i]
	})

	return string(b)
}

package go_stringable

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
	"unicode"
)

const STR_PAD_LEFT = 0

const STR_PAD_RIGHT = 1

func StrPadLeft(value string, pad rune, length int) string {
	res, _ := StrPad(value, pad, length, STR_PAD_LEFT)

	return res
}

func StrPadRight(value string, pad rune, length int) string {
	res, _ := StrPad(value, pad, length, STR_PAD_RIGHT)

	return res
}

func StrPad(value string, pad rune, length int, padType int) (string, error) {
	padCount := length - len(value)
	if padCount <= 0 {
		return value, nil
	}

	paddedStr := make([]rune, length)

	if padType == STR_PAD_LEFT {
		copy(paddedStr[padCount:], []rune(value))
		for i := 0; i < padCount; i++ {
			paddedStr[i] = pad
		}
	} else if padType == STR_PAD_RIGHT {
		copy(paddedStr[0:], []rune(value))
		for i := len(value); i < length; i++ {
			paddedStr[i] = pad
		}
	} else {
		return "", errors.New("The padType is invalid")
	}

	return string(paddedStr), nil
}

func Md5(str string) string {
	return Md5ByteArray([]byte(str))
}

func Md5ByteArray(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func Random(n int) (string, error) {
	return RandomLetters(n, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

func RandomLetters(n int, letters string) (string, error) {
	result := make([]byte, n)
	for i := range result {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		result[i] = letters[num.Int64()]
	}
	return string(result), nil
}

func RandomNumber(n int) (string, error) {
	return RandomLetters(n, "0123456789")
}

func MaskByStar(value string, offset int, length int) string {
	return Mask(value, offset, length, "*")
}

func MaskLastByStar(value string, offset int) string {
	return Mask(value, offset, 0, "*")
}

func Mask(value string, offset int, length int, replacement string) string {
	l := len(value)
	absOffset := abs(offset)
	if length < 0 || l <= absOffset {
		return value
	}

	replaceLength := length
	if length == 0 {
		replaceLength = l - absOffset
	}

	values := []rune(value)
	replacements := []rune(Repeat(replacement, replaceLength))
	var res []rune
	if offset >= 0 {
		res = append(values[0:offset], replacements...)
		res = append(res, values[offset+replaceLength:]...)
		return string(res)
	}

	if l+offset-replaceLength > 0 {
		res = append(values[0:l+offset-replaceLength], replacements...)
	} else {
		res = replacements
	}
	res = append(res, values[l+offset:]...)
	return string(res)
}

func RepeatRunes(value []rune, length int) []rune {
	var res []rune
	for i := 0; i < length; i++ {
		res = append(res, value...)
	}
	return res
}

func Repeat(value string, n int) string {
	if value == "" || n <= 0 {
		return ""
	}

	return string(RepeatRunes([]rune(value), n))
}

func Snake(value string) string {
	return SnakeWithDelimiter(value, '_')
}

func SnakeWithDelimiter(value string, delimiter byte) string {
	var builder strings.Builder
	for i, r := range value {
		if unicode.IsUpper(r) && i != 0 {
			builder.WriteByte(delimiter)
		}
		builder.WriteRune(unicode.ToLower(r))
	}
	return builder.String()
}

func Studly(value string) string {
	parts := strings.Split(value, "_")
	var builder strings.Builder
	for _, part := range parts {
		if part != "" {
			ps := strings.Split(part, "-")
			for _, p := range ps {
				if p != "" {
					builder.WriteString(strings.Title(strings.ToLower(p)))
				}
			}
		}
	}

	return builder.String()
}

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}

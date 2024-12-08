package go_stringable

import (
	"errors"
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

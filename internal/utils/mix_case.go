package utils

import "unicode"

func MixCase(text string) string {
	chars := []rune(text)

	for i, char := range chars {
		if i%2 == 0 {
			chars[i] = unicode.ToLower(char)
		} else {
			chars[i] = unicode.ToUpper(char)
		}
	}

	return string(chars)
}

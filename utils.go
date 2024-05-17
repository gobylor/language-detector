package langdetect

import "unicode"

func isStopChar(r rune) bool {
	if unicode.IsSymbol(r) || unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsDigit(r) {
		return true
	}
	return false
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

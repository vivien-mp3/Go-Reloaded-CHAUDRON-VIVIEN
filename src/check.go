package src

func IsLetter(r rune) bool {
	return ((r >= 65 && r <= 90) || (r >= 97 && r <= 122))
}

func IsNumber(r rune) bool {
	return (r >= 48 && r <= 57)
}

func IsSymbol(r rune) bool {
	return ((r >= 33 && r <= 47) || (r >= 58 && r <= 64)) && r != 39
}

func IsSpace(r rune) bool {
	return r == 32
}

func IsEndOfSentence(r rune) bool {
	return r == '.' || r == '!' || r == '?'
}

func isApos(r rune) bool {
	return r == 39
}

func isAn(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'h' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' || r == 'H'
}
package src

/*
Ce code a été écrit avec l'espoir d'un homme.
Bonne chance
*/
//import "fmt"

var aspos bool

func CutText(t string) []string {

	result := []string{}
	str := ""
	for i := 0; i < len(t); i++ {
		r := rune(t[i])
		prev := i - 1
		if prev < 0 {
			prev = 0
		}
		//chr := string(r)
		if IsLetter(r) {
			if t[i] > 0 && !IsLetter(rune(t[prev])) {
				result, str = strAppend(str)
			}

			if len(result) > 0 {
				if IsEndOfSentence(rune(result[len(result)-1][0])) && (r >= 97 && r <= 122) && !IsLetter(rune(t[prev]))  {
					r -= 32
				}
			} else if i == 0 && (r >= 97 && r <= 122) {
					r -= 32
			}
			str += string(r)

		} else if IsNumber(r) {
			if t[i] > 0 && !IsNumber(rune(t[prev])) {
				result, str = strAppend(str)
			}
			str += string(r)

		} else if IsSymbol(r) {
			if r == 40 {
				i = GetAction(t, result, i)
				continue
			} else if r == 39 {
				aspos = true
			}

			if t[i] > 0 && IsSpace(rune(t[prev])) {
				result, str = strAppend(str)
			}
			str += string(r)

		} else {
			result, str = strAppend(str)
		}
	}
	result, str = strAppend(str)

	return result
}

func GetAction(t string, tb []string, id int) int {
	act := ""
	nbr := 0
	for k, _ := range t {
		if rune(t[id+k]) == 41 {
			if nbr == 0 {
				nbr = 1
			}
			switch act {
			case "up":
				Upper(tb, nbr)
			case "low":
				Lower(tb, nbr)
			case "cap":
				Capitalize(tb, nbr)
			case "hex":
				Hex(tb, nbr)
			case "bin":
				Bin(tb, nbr)
			}
			return id + k
		}
		if IsLetter(rune(t[id+k])) {
			act += string(t[id+k])
		}
		if IsNumber(rune(t[id+k])) {
			nbr = nbr*10 + int(t[id+k]) - 48
		}
	}

	return id
}

var result2 []string

func strAppend(str string) ([]string, string) {
	if str != "" {
		if aspos && str[len(str)-1] != 39 {
			str += "'"
			aspos = false
		}
		result2 = append(result2, str)
	}
	str = ""
	return result2, str
}

func IsLetter(r rune) bool {
	return ((r >= 65 && r <= 90) || (r >= 97 && r <= 122))
}

func IsNumber(r rune) bool {
	return (r >= 48 && r <= 57)
}

func IsSymbol(r rune) bool {
	return ((r >= 33 && r <= 47) || (r >= 58 && r <= 64))
}

func IsSpace(r rune) bool {
	return r == 32
}

func IsEndOfSentence(r rune) bool {
	return r == '.' || r == '!' || r == '?'
}
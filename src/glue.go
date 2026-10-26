package src

import "fmt"

func GlueText(t []string) string {
	result := ""

	for i := 0; i < len(t)-1; i++ {
		
		fmt.Println(t[i])

		if t[i] == "a" || t[i] == "A" {
			if isAn(rune(t[i+1][0])) {
				t[i] += "n"
			}
		}

		result += t[i]

		if IsSymbol(rune(t[i+1][0])) && t[i+1][0] != 39 {
			continue
		} else if isApos(rune(t[i][0])) {
			if IsSymbol(rune(t[i+1][len(t[i+1])-1])) {
				last := t[i+1][len(t[i+1])-1]
				t[i+1] = t[i+1][:len(t[i+1])-1]
				t[i+1] += "'" + string(last)
			} else {
				t[i+1] += "'"
			}
			
		}

		if !isApos(rune(t[i][0])) {
			result += " "
		}
		
	}
	result += t[len(t)-1]
	return result
}
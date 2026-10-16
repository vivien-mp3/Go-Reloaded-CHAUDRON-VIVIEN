package src

//import "fmt"

func GlueText(t []string) string {
	result := ""

	for i := 0; i < len(t)-1; i++ {
		

		result += t[i]

		if IsSymbol(rune(t[i+1][0])) && t[i+1][0] != 39 {
			continue
		}
		result += " "
	}
	result += t[len(t)-1]
	return result
}
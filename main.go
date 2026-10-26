package main

import (
	"fmt"
	"go-reloaded/src"
	"os"
)

func main() {


	dat, _ := os.ReadFile(os.Args[1])
	fmt.Println(string(dat))

	cuttedText := src.CutText(string(dat))
	newData := []byte(src.GlueText(cuttedText))
	
	

	resultFile := "result.txt"
	if len(os.Args) >= 3 {
		resultFile = os.Args[2]
		if len(os.Args[2]) >= 3 && os.Args[2][len(os.Args[2])-4:] != ".txt" {
			resultFile += ".txt"
		}
	}

	if resultFile == os.Args[1] {
		fmt.Printf("%v: Les deux arguments partagent le même noms. Veuillez différencier les noms.", src.ErrSameName)
	}

	if _, err := os.Stat(resultFile); !os.IsNotExist(err) {
		os.Remove(resultFile)
	}

	var perm os.FileMode
	os.WriteFile(resultFile, newData, perm)

	dat2, _ := os.ReadFile(resultFile)
	fmt.Println(string(dat2))
}

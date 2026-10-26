package main

import (
	"fmt"
	"go-reloaded/src"
	"os"
)

func main() {


	dat, _ := os.ReadFile(os.Args[1])
	fmt.Println(string(dat))

	correctedText := src.CutText(string(dat))
	//fmt.Println(src.GlueText(correctedText))
	newData := []byte(src.GlueText(correctedText))
	
	if _, err := os.Stat("result.txt"); !os.IsNotExist(err) {
		os.Remove("result.txt")
	}

	resultFile := "result.txt"
	if len(os.Args) >= 3 {
		resultFile = os.Args[2]
		if len(os.Args[2]) >= 3 && os.Args[2][len(os.Args[2])-4:] != ".txt" {
			resultFile += ".txt"
		}
	}

	var perm os.FileMode
	os.WriteFile(resultFile, newData, perm)

	dat2, _ := os.ReadFile(resultFile)
	fmt.Println(string(dat2))
}

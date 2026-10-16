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
	fmt.Println(src.GlueText(correctedText))
	newData := []byte(src.GlueText(correctedText))
	
	if _, err := os.Stat("result.txt"); !os.IsNotExist(err) {
		os.Remove("result.txt")
	}
	var perm os.FileMode
	os.WriteFile("result.txt", newData, perm)

	dat2, _ := os.ReadFile("result.txt")
	fmt.Println(string(dat2))
}

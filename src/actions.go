package src

import (
	"fmt"
	"strconv"
	"strings"
)

func Upper(t []string, n int) {
	for i := (len(t)-(n)); i < len(t); i++ {
		t[i] = strings.ToUpper(t[i])
	}
}

func Lower(t []string, n int) {
	for i := (len(t)-(n)); i < len(t); i++ {
		t[i] = strings.ToLower(t[i])
	}
}

func Capitalize(t []string, n int) {
	for i := (len(t)-(n)); i < len(t); i++ {
		if (t[i][0] >= 97 && t[i][0] <= 122) {
			cap := []rune(t[i])
			cap[0] -= 32
			t[i] = string(cap)

		}
	}
}

func Bin(t []string, n int) {
	nbr := 0
	for i := (len(t)-(n)); i < len(t); i++ {
		str := t[i]
		for k := len(str)-1; k >= 0; k-- {
			if IsNumber(rune(str[k])) {
				nbr = nbr * 10 + (int(str[k])-48)
			}
		}
		bin := strconv.FormatInt(int64(nbr), 2)
		fmt.Println(bin)
		t[i] = bin
	}
}

func Hex(t []string, n int) {
	nbr := 0
	for i := (len(t)-(n)); i < len(t); i++ {
		str := t[i]
		for k := len(str)-1; k >= 0; k-- {
			if IsNumber(rune(str[k])) {
				nbr = nbr * 10 + (int(str[k])-48)
			}
		}
		hex := strconv.FormatInt(int64(nbr), 16)
		fmt.Println(hex)
		t[i] = hex
	}
}
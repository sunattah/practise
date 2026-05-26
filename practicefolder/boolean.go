package main

import (
	"fmt"
	"os"
)

func printStr(s string) {
	for _, r := range s {
		fmt.Println(r)
	}
	fmt.Println('\n')
}

func isEven(nbr int) bool {
	if (nbr) == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	if len(os.Args) == 1 {
		print("good")
	} else {
		printStr("the arguments should be two")
	}
}

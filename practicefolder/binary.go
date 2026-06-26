package main

import "fmt"

func BinaryCheck(nbr int32) int {
	if nbr%2 == 0 {
		return 1
	}
	return 0

}

func main() {
	fmt.Println(BinaryCheck(5))
	fmt.Println(BinaryCheck(0))
	fmt.Println(BinaryCheck(8))
	fmt.Println(BinaryCheck(-9))
	fmt.Println(BinaryCheck(-4))
}

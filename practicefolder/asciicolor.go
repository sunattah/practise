package main

import "fmt"

func findAll(line, sub string) []int {
	lenLine := len(line)
	lenSub := len(sub)
	lenslice := []int{}
	for i := 0; i <= lenLine- lenSub; i++ {
		if line[i:i+lenSub] == sub {
			lenslice = append(lenslice, i)

		}
	}
	return lenslice
}
func main() {
	fmt.Println(findAll("a king kitten have kit", "kit"))
}

package main

import "fmt"

func findAll(line, sub string) []int {
	lenLine := len(line)
	lenSub := len(sub)
	lenslice := []int{}
	for i := 0; i <= lenLine-lenSub; i++ {
		if line[i:i+lenSub] == sub {
			lenslice = append(lenslice, i)

		}
	}
	return lenslice
}
func markRanges(line, sub string) []bool {
	indecies := findAll(line, sub)
	colored := make([]bool, len(line))

	for _, ch := range indecies {
		for j := ch; j < ch+len(sub); j++ {
			colored[j] = true
		}
	}
	return colored

}
func compareLengths(str string) {
	fmt.Println(len(str))
	fmt.Println(len([]rune(str)))
}

func main() {
	fmt.Println(findAll("a king kitten have kit", "kit"))
	
	fmt.Println(markRanges("a king kitten have kit", "kit"))
	compareLengths("hello")
	compareLengths("café")
}

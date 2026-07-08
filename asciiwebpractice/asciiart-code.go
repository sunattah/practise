package main

import (
	"fmt"
	"os"
	"strings"
)

func banner(input string) (map[rune][]string, error) {
	file, err := os.ReadFile(input)
	if err != nil {
		return nil, err

	}
	lines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")

	bannermap := map[rune][]string{}
	currentchar := rune(32)

	for i := 1; i+8 < len(lines); i += 9 {
		bannermap[currentchar] = lines[i : i+8]
		currentchar++
	}
	return bannermap, err
}
func render(input string, banner map[rune][]string) []string {
	slicebanner := make([]string, 8)
	for i := 0; i < 8; i++ {
		for _, ch := range input {
			slicebanner[i] += banner[ch][i]
		}
	}
	return slicebanner
}
func spli(input string) []string {
	line := strings.ReplaceAll(input, "\\n", "\n")
	return strings.Split(line, "\n")
}
func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage go run . <TEXT> [STRING]")
	// }
	input := os.Args[1]
	bannerm := "standard.txt"
	if len(os.Args) == 3 {
		bannerm = os.Args[2] + ".txt"
	}
	g, err := banner(bannerm)
	if err != nil {
		fmt.Printf("no such file %s", err)
		os.Exit(0)
	}
	splitf := spli(input)
	for _, ch := range splitf {
		if ch == "" {
			fmt.Println()
			continue
		}
		renderline := render(ch, g)
		for _, char := range renderline {
			fmt.Println(char)
		}
	}

}

package main

import "strings"

func StringToArt(input string) string {
	if input == "" {
		return ""
	}

	digits := map[rune][]string{
		'0': {
			"_____",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2': {
			"_____",
			"    |",
			"____|",
			"|    ",
			"|____",
		},
		'3': {
			"_____",
			"    |",
			"____|",
			"    |",
			"____|",
		},
		'4': {
			"|   |",
			"|___|",
			"    |",
			"    |",
			"    |",
		},
		'5': {
			"_____",
			"|    ",
			"|____",
			"    |",
			"____|",
		},
		'6': {
			"_____",
			"|    ",
			"|____",
			"|   |",
			"|___|",
		},
		'7': {
			"_____",
			"    |",
			"    |",
			"    |",
			"    |",
		},
		'8': {
			"_____",
			"|   |",
			"|___|",
			"|   |",
			"|___|",
		},
		'9': {
			"_____",
			"|   |",
			"|___|",
			"    |",
			"____|",
		},
	}

	var b strings.Builder

	lines := strings.Split(input, "\n")

	for _, lineInput := range lines {

		for _, ch := range lineInput {
			if ch < '0' || ch > '9' {
				return ""
			}
		}

		for row := 0; row < 5; row++ {
			for _, ch := range lineInput {
				b.WriteString(digits[ch][row])
			}
			b.WriteString("\n")
		}
	}

	return b.String()
}

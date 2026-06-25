package main

import (
	"strconv"
	"strings"
)

func capitalizeletter(input string) string {
	fieldString := strings.Fields(input)
	for ch := 0; ch < len(fieldString); ch++ {
		if fieldString[ch] == "(cap)" {
			if ch >= 0 {
				fieldString[ch-1] = strings.Title(fieldString[ch-1])

			}
			fieldString = append(fieldString[:ch], fieldString[ch+1:]...)

		}
		if fieldString[ch] == "(cap," {
			num := strings.TrimSuffix(fieldString[ch+1], ")")
			number, err := strconv.Atoi(num)
			if err == nil {
				for j := ch - 1; j >= 0 && j >= ch-number; j-- {
					word := fieldString[j]
					fieldString[j] = strings.Title(word[:1] + word[1:])
				}
			}
			fieldString = append(fieldString[:ch], fieldString[ch+2:]...)
			ch--

		}
	}
	return strings.Join(fieldString, " ")
}

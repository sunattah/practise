package main

import "strings"

func Camelcase(input string) string {
	if input == "" {
		return ""
	}
	camel_case := strings.Fields(input)
	for _, ch := range camel_case {
		if ch != "" {
			return strings.ToUpper(ch)
		}
		if input != camel_case[0] {
			return input
		}
		if input == camel_case[0] {
			return input + "_"

		}
	}
	return input

}

package main

import (
	"strings"
)

func Art(input string) string {
	f := strings.Fields(input)
	for i := 0; i < len(f); i++ {
		word := f[i]
		if word == "a" && strings.ContainsAny(string(f[i]), "aeiouhAEIOUH") {
			f[i] = "an"
		} else if word == "an" && strings.ContainsAny(string(f[i]), "aeiouhAEIOUH") {
			f[i] = "a"

		} else if word == "A" && strings.ContainsAny(string(f[i]), "aeiouhAEIOUH") {
			f[i] = "An"

		} else if word == "An" && strings.ContainsAny(string(f[i]), "aeiouhAEIOUH") {
			f[i] = "A"
		}
	}
	return strings.Join(f, " ")
}

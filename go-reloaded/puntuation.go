package main

import (
	"regexp"
	"strings"
)

func punc(input string) string {
	s := regexp.MustCompile(`\s+([.,?':;])`).ReplaceAllString(input, "$1")
	regexp.MustCompile(`\s*(['"])([A-Z a-z])(['"])\s*`).ReplaceAllString(s, "$1 $2 $3")
	return strings.Join(strings.Fields(s), " ")
}

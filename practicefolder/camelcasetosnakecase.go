// if invalid camelCase
//     return original string

// result := ""

// for every letter
//     if uppercase
//         if not first letter
//             add "_"
//         add lowercase version
//     else
//         add letter

// return result

package main

func Camelcase(input string) string {
	if input == "" {
		return ""
	}
	for i := 0; i < len(input); i++ {
		c := input[i]
		if !(c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z') {
			return input
		}
	}
}

package main

func AlphaPosition(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	}
	if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 1)
	}
	return -1
}

// func Any(f func(string) bool, a []string) bool {

// }

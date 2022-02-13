package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(isAnagram(os.Args[1], os.Args[2]))
}

func isAnagram(s1, s2 string) bool {
	r1 := make(map[rune]int)
	for _, r := range s1 {
		r1[r]++
	}

	r2 := make(map[rune]int)
	for _, r := range s2 {
		r2[r]++
	}

	for r, freq := range r1 {
		if r2[r] != freq {
			return false
		}
	}

	for r, freq := range r2 {
		if r1[r] != freq {
			return false
		}
	}

	return true
}

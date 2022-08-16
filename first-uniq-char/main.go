package main

import (
	"fmt"
)

func main() {
	fmt.Println("aaa")
	test := "eloveleetcode"
	// k := 3
	fmt.Println(firstUniqChar(test))
}

func firstUniqChar(s string) int {
	counter := map[rune]int{}

	for _, letter := range s {
		if _, ok := counter[letter]; !ok {
			counter[letter] = 0
		}
		counter[letter]++
	}
	for i, letter := range s {
		if count := counter[letter]; count == 1 {
			return i
		}
	}
	return -1
}

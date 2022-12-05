package main

import (
	"fmt"
)

func main() {
	fmt.Println("aaa")
	test := []int{1, 2, 3, 4, 2}
	fmt.Println(containsDuplicate(test))
}

func containsDuplicate(nums []int) bool {
	numMap := make(map[int]int)

	for _, num := range nums {
		if _, ifExist := numMap[num]; ifExist {
			return true
		} else {
			numMap[num] = 1
		}
	}

	return false
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("aaa")
	test := []int{1, 2, 3, 1, 2, 3}
	k := 3
	fmt.Println(containsNearbyDuplicate(test, k))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		fmt.Println(numMap)
		if _, found := numMap[nums[i]]; found {
			return true
		} else {
			numMap[nums[i]] = i
			if len(numMap) > k {
				delete(numMap, nums[i-k])
			}
		}
	}
	return false
}

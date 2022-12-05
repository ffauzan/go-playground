package main

import (
	"fmt"
	"sort"
)

func main() {
	array := []int32{1, 5, 6, 4, 8, 4, 5, 2}
	fmt.Println(array)
	sort.Slice(array, func(i, j int) bool { return array[i] < array[j] })
	fmt.Println(array)
}

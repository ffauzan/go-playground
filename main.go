package main

import (
	"fmt"
)

func main() {
	fmt.Println("aa")
	arr := [][]int32{
		{112, 42, 83, 119},
		{56, 125, 56, 49},
		{15, 78, 101, 43},
		{62, 98, 114, 108},
	}

	flippingMatrix(arr)
}

func flippingMatrix(matrix [][]int32) int32 {
	// Write your code here
	quadranSize := len(matrix) / 2

	// row

	for i := 0; i < len(matrix); i++ {
		// First half sum
		firstHalfRowSum := 0

		sectHalfRowSum := 0

		for j := 0; j < len(matrix); j++ {
			if i < len(matrix)/2 {
				firstHalfRowSum += int(matrix[i][j])
			} else {
				sectHalfRowSum += int(matrix[i][j])
			}
		}

		// switch if needed
		if firstHalfRowSum < sectHalfRowSum {
			// switch

		}

	}

}

func reverseRow(row []int32) []int32 {

}

package main

import "fmt"

func main() {
	houseStart := 2
	houseEnd := 3

	appleTreeLoc := 1
	orangeTreeLoc := 5

	// appleCount := 3
	// orangeCount := 2

	appleLoc := []int32{-2}
	orangeLoc := []int32{-1}

	countApplesAndOranges(int32(houseStart), int32(houseEnd), int32(appleTreeLoc), int32(orangeTreeLoc), appleLoc, orangeLoc)
}

// https://www.hackerrank.com/challenges/apple-and-orange/problem?isFullScreen=true

/*
 * Complete the 'countApplesAndOranges' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER s
 *  2. INTEGER t
 *  3. INTEGER a
 *  4. INTEGER b
 *  5. INTEGER_ARRAY apples
 *  6. INTEGER_ARRAY oranges
 */
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	// Write your code here
	appleFall := 0
	orangeFall := 0

	// Get locations of apples
	for i := 0; i < len(apples); i++ {
		// If the location is within s and t (inclusive), add it to the count
		if (s <= apples[i]+a) && (apples[i]+a <= t) {
			appleFall++
		}
	}

	// Get locations of orange
	for i := 0; i < len(oranges); i++ {
		// If the location is within s and t (inclusive), add it to the count
		if (s <= oranges[i]+b) && (oranges[i]+b <= t) {
			orangeFall++
		}
	}

	fmt.Println(appleFall)
	fmt.Println(orangeFall)
}

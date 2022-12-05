package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("j")

	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}

	res := compareTriplets(a, b)

	fmt.Println(res)
}

// Source: Hackerrank
func compareTriplets(a []int32, b []int32) []int32 {
	aliceScore := 0
	bobScore := 0

	for i := 0; i < 3; i++ {
		// Kalau seri ga ada yang dapet nilai
		if a[i] == b[i] {
			continue
		}

		// Yang menang dapet nilai 1
		if a[i] > b[i] {
			aliceScore = aliceScore + 1
		} else {
			bobScore = bobScore + 1
		}
	}

	return []int32{int32(aliceScore), int32(bobScore)}
}

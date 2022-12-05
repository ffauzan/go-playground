package main

import (
	"fmt"
)

func main() {
	h := 5
	m := 4

	res := timeInWords(int32(h), int32(m))
	fmt.Println(res)
}

func timeInWords(h int32, m int32) string {
	// Write your code here
	hourPartInWord := ""
	minutePartInWord := ""
	relationWord := "" // past/to
	finalWords := ""

	// Create word representation of number
	numberToWordMap := make(map[int]string)

	numberToWordMap[1] = "one"
	numberToWordMap[2] = "two"
	numberToWordMap[3] = "three"
	numberToWordMap[4] = "four"
	numberToWordMap[5] = "five"
	numberToWordMap[6] = "six"
	numberToWordMap[7] = "seven"
	numberToWordMap[8] = "eight"
	numberToWordMap[9] = "nine"
	numberToWordMap[10] = "ten"

	numberToWordMap[11] = "eleven"
	numberToWordMap[12] = "twelve"
	numberToWordMap[13] = "thirteen"

	for i := 14; i <= 19; i++ {
		if (i == 18) || (i == 15) {
			continue
		}

		numberToWordMap[i] = numberToWordMap[i-10] + "teen"
	}

	numberToWordMap[15] = "quarter"
	numberToWordMap[18] = "eighteen"
	numberToWordMap[20] = "twenty"
	numberToWordMap[30] = "half"
	numberToWordMap[45] = "quarter"

	for i := 21; i <= 29; i++ {
		numberToWordMap[i] = "twenty " + numberToWordMap[i-20]
	}

	// Do the thing
	switch {
	case m == 0:
		minutePartInWord = "o' clock"
		relationWord = ""
		hourPartInWord = numberToWordMap[int(h)]
		finalWords = hourPartInWord + " " + minutePartInWord
	case (m < 30):
		minutePartInWord = numberToWordMap[int(m)]
		hourPartInWord = numberToWordMap[int(h)]
		if hourPartInWord == "one" {
			relationWord = "minute past"
		} else if m == 15 {
			relationWord = "past"
		} else {
			relationWord = "minutes past"
		}
		finalWords = minutePartInWord + " " + relationWord + " " + hourPartInWord
	case m == 30:
		minutePartInWord = "half"
		hourPartInWord = numberToWordMap[int(h)]
		relationWord = "past"
		finalWords = minutePartInWord + " " + relationWord + " " + hourPartInWord
	case m > 30:
		minutePartInWord = numberToWordMap[60-int(m)]
		hourPartInWord = numberToWordMap[int(h)+1]
		if hourPartInWord == "one" {
			relationWord = "minute to"
		} else if m == 45 {
			relationWord = "to"
		} else {
			relationWord = "minutes to"
		}
		finalWords = minutePartInWord + " " + relationWord + " " + hourPartInWord
	}

	return finalWords
}

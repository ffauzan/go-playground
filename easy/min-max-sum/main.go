package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
	// Write your code here
	var min, max, sum, minSum, maxSum int64

	for i := 0; i < len(arr); i++ {
		sum += int64(arr[i])

		if i == 0 {
			min = int64(arr[i])
			max = int64(arr[i])
			continue
		}

		if int64(arr[i]) < min {
			min = int64(arr[i])
		}

		if int64(arr[i]) > max {
			max = int64(arr[i])
		}
	}

	minSum = sum - max
	maxSum = sum - min

	fmt.Printf("%d %d", minSum, maxSum)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

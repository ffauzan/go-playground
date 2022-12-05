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
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
	// Write your code here
	isAm := strings.Contains(s, "AM")
	s = strings.Trim(s, "APM")

	arr := strings.Split(s, ":")

	if !isAm && arr[0] == "12" {
		return fmt.Sprintf("%s:%s:%s", arr[0], arr[1], arr[2])
	}

	if isAm && arr[0] == "12" {
		return fmt.Sprintf("%s:%s:%s", "00", arr[1], arr[2])
	}

	if isAm {
		return fmt.Sprintf("%s:%s:%s", arr[0], arr[1], arr[2])
	}

	newHour, _ := strconv.Atoi(arr[0])
	newHour += 12

	return fmt.Sprintf("%d:%s:%s", newHour, arr[1], arr[2])
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
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

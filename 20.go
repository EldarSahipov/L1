package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"

	fmt.Println(reversedStr4(str))
}

func reversedStr(str string) string {
	sliceStr := strings.Split(str, " ")
	newSlice := make([]string, len(sliceStr))

	for i := len(sliceStr) - 1; i >= 0; i-- {
		newSlice[len(sliceStr)-1-i] = sliceStr[i]
	}

	return strings.Join(newSlice, " ")
}

func reversedStr2(str string) string {
	sliceStr := strings.Split(str, " ")
	newBuilder := strings.Builder{}

	for i := len(sliceStr) - 1; i >= 0; i-- {
		newBuilder.WriteString(sliceStr[i])
		if i > 0 {
			newBuilder.WriteString(" ")
		}
	}

	return newBuilder.String()
}

func reversedStr3(str string) string {
	sliceStr := strings.Split(str, " ")
	newBuffer := bytes.Buffer{}

	for i := len(sliceStr) - 1; i >= 0; i-- {
		newBuffer.WriteString(sliceStr[i])
		if i > 0 {
			newBuffer.WriteString(" ")
		}
	}

	return newBuffer.String()
}

func reversedStr4(str string) string {
	sliceStr := strings.Split(str, " ")

	for i, j := 0, len(sliceStr)-1; i < j; i, j = i+1, j-1 {
		sliceStr[i], sliceStr[j] = sliceStr[j], sliceStr[i]
	}

	return strings.Join(sliceStr, " ")
}

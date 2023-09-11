package main

import "strings"

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
}

func createHugeString(i int) string {
	str := strings.Builder{}
	str.Grow(i)
	for j := 0; j < i; j++ {
		str.WriteString("A")
	}
	result := str.String()
	str.Reset() // Сбрасываем содержимое Builder, чтобы освободить память
	return result
}

func main() {
	someFunc()
}

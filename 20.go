package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str := "snow dog sun"

	fmt.Println(reversedStr(str))
	fmt.Println(reversedStr2(str))
	fmt.Println(reversedStr3(str))
	fmt.Println(reversedStr4(str))
}

// reversedStr переворачивает слова в строке, используя разделитель пробела.
func reversedStr(str string) string {
	sliceStr := strings.Split(str, " ") // Разбиваем строку на срез слов.
	newSlice := make([]string, len(sliceStr))

	// Переворачиваем срез слов.
	for i := len(sliceStr) - 1; i >= 0; i-- {
		newSlice[len(sliceStr)-1-i] = sliceStr[i]
	}

	return strings.Join(newSlice, " ") // Объединяем перевернутые слова в строку.
}

// reversedStr2 выполняет переворот слов, используя strings.Builder.
func reversedStr2(str string) string {
	sliceStr := strings.Split(str, " ") // Разбиваем строку на срез слов.
	newBuilder := strings.Builder{}

	// Переворачиваем срез слов и записываем их в Builder.
	for i := len(sliceStr) - 1; i >= 0; i-- {
		newBuilder.WriteString(sliceStr[i])
		if i > 0 {
			newBuilder.WriteString(" ") // Добавляем пробел между словами.
		}
	}

	return newBuilder.String() // Получаем строку из Builder.
}

// reversedStr3 выполняет переворот слов, используя bytes.Buffer.
func reversedStr3(str string) string {
	sliceStr := strings.Split(str, " ") // Разбиваем строку на срез слов.
	newBuffer := bytes.Buffer{}

	// Переворачиваем срез слов и записываем их в Buffer.
	for i := len(sliceStr) - 1; i >= 0; i-- {
		newBuffer.WriteString(sliceStr[i])
		if i > 0 {
			newBuffer.WriteString(" ") // Добавляем пробел между словами.
		}
	}

	return newBuffer.String() // Получаем строку из Buffer.
}

func reversedStr4(str string) string {
	sliceStr := strings.Split(str, " ") // Разбиваем строку на срез слов.

	// Переворачиваем срез слов с использованием двух указателей.
	for i, j := 0, len(sliceStr)-1; i < j; i, j = i+1, j-1 {
		sliceStr[i], sliceStr[j] = sliceStr[j], sliceStr[i]
	}

	return strings.Join(sliceStr, " ") // Объединяем перевернутые слова в строку.
}

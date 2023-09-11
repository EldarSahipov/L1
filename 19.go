package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "главрыба"

	// Выводим результаты двух функций переворачивания строки.
	fmt.Println(reverse(str))
	fmt.Println(reverse2(str))
}

// reverse переворачивает строку, используя руны (Unicode символы) и возвращает результат.
func reverse(str string) string {
	// Преобразуем строку в руны для правильной обработки символов Unicode.
	runes := []rune(str)

	// Переворачиваем строку, меняя местами символы с начала и с конца.
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Преобразуем руны обратно в строку.
	return string(runes)
}

// reverse2 также переворачивает строку, используя разделение строк и конкатенацию их в обратном порядке.
func reverse2(str string) string {
	// Разбиваем строку на слайс строк (по символам).
	sliceStr := strings.Split(str, "")
	// Создаем слайс для хранения символов в обратном порядке.
	revSliceStr := make([]string, len(sliceStr))

	// Переворачиваем строку, помещая символы из начального слайса в обратном порядке в новый слайс.
	for i := len(sliceStr) - 1; i >= 0; i-- {
		revSliceStr[len(sliceStr)-1-i] = sliceStr[i]
	}

	// Конкатенируем символы в новой строке и возвращаем результат.
	return strings.Join(revSliceStr, "")
}

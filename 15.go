package main

import "strings"

var justString string

func someFunc() {
	v := createHugeString(1 << 10) // Создаем огромную строку (1024 байта).
	justString = v[:100]           // Присваиваем только первые 100 байт этой строки переменной justString.
}

// Основная проблема заключается в том, что createHugeString создает строку, выделяя память для нее,
// и затем возвращает строку. Это может привести к большому потреблению памяти,
// если создавать большие строки
func createHugeString(i int) string {
	str := strings.Builder{} // Создаем новый strings.Builder, который позволит нам создавать большие строки.
	str.Grow(i)              // Выделяем память для строки на i байт.
	for j := 0; j < i; j++ {
		str.WriteString("A") // Используем WriteByte для добавления одного байта
	}
	result := str.String()
	str.Reset() // Сбрасываем содержимое Builder, чтобы освободить память
	return result
}

func main() {
	someFunc()
}

package main

import "fmt"

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}

	mySet := make(map[string]bool) // Создаем мапу для представления множества строк.

	// Итерируемся по всем строкам в последовательности.
	for _, value := range slice {
		mySet[value] = true // Добавляем строку как ключ в мапу с значением true.
	}

	// Выводим содержимое множества.
	fmt.Print("Множество: ")
	for key := range mySet {
		fmt.Print(key, ", ") // Выводим уникальные строки из множества.
	}
}

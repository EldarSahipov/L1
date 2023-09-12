package main

import "fmt"

func main() {
	// Исходный срез
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(deleteElementByIndex(slice, 5))
	fmt.Println(deleteElementByIndex2(slice, 5))
	fmt.Println(deleteElementByIndex3(slice, 5))
}

// deleteElementByIndex удаляет элемент из среза s по заданному индексу index
func deleteElementByIndex(s []int, index int) []int {
	var sliceNew []int

	// Копируем элементы до индекса index
	for i := 0; i < index; i++ {
		sliceNew = append(sliceNew, s[i])
	}

	// Копируем элементы после индекса index
	for i := index + 1; i < len(s); i++ {
		sliceNew = append(sliceNew, s[i])
	}

	return sliceNew
}

// deleteElementByIndex2 удаляет элемент из среза s по заданному индексу index,
// используя операцию среза
func deleteElementByIndex2(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// deleteElementByIndex3 удаляет элемент из среза s по заданному индексу index,
// создавая новый срез и копируя элементы
func deleteElementByIndex3(s []int, index int) (slice []int) {
	slice = make([]int, len(s)-1)

	// Копируем элементы до индекса index
	copy(slice, s[:index])

	// Копируем элементы после индекса index
	copy(slice[index:], s[index:])
	return
}

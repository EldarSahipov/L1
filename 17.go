package main

import "fmt"

func main() {
	arr := []int{1, 5, 7, 56, 57, 155, 544, 1058, 5455}
	fmt.Println(bin(arr, 5455))

}

// bin выполняет бинарный поиск значения n в упорядоченном массиве arr.
// Возвращает индекс элемента, если найден, или -1, если не найден.
func bin(arr []int, n int) int {
	left := 0             // Индекс начала подмассива.
	right := len(arr) - 1 // Индекс конца подмассива.

	for left <= right {
		middle := (right + left) / 2 // Находим середину подмассива.
		if n == arr[middle] {
			return middle // Значение найдено, возвращаем его индекс.
		} else if n > arr[middle] {
			left = middle + 1 // Значение больше, сужаем поиск к правой части подмассива.
		} else if n < arr[middle] {
			right = middle - 1 // Значение меньше, сужаем поиск к левой части подмассива.
		}
	}
	return -1 // Значение не найдено, возвращаем -1.
}

package main

import "fmt"

// doSort выполняет сортировку подмассива items от fst до lst включительно.
func doSort(items []int, fst, lst int) {
	if fst >= lst {
		return // Базовый случай: массив имеет 1 элемент или пуст.
	}

	i := fst
	j := lst
	x := items[(fst+lst)/2] // Выбираем опорный элемент.

	// Разделение элементов на две группы: меньше и больше опорного элемента.
	for i < j {
		for items[i] < x {
			i++
		}
		for items[j] > x {
			j--
		}
		if i <= j {
			items[i], items[j] = items[j], items[i] // Меняем элементы местами.
			i++
			j--
		}
	}

	// Рекурсивно сортируем подмассивы.
	doSort(items, fst, j)
	doSort(items, i, lst)
}

// quicksort выполняет сортировку массива arr и возвращает новый отсортированный массив.
func quicksort(arr []int) []int {
	items := make([]int, len(arr))
	copy(items, arr)
	doSort(items, 0, len(arr)-1) // Вызываем рекурсивную функцию сортировки.
	return items
}

func main() {
	// Исходный массив чисел.
	items := []int{1, 64, 64, 63, 135, 48, 4687, 413, 1, 654, 6541, 684, 11, 32, 1, 1, 12, 15, 45, 15, 4, 846, 51, 6, 5, 4, 654, 6513, 21, 65416}

	// Вызываем функцию быстрой сортировки.
	sortItems := quicksort(items)

	// Выводим отсортированный массив.
	fmt.Println(sortItems)

}

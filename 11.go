package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Создание двух неупорядоченных множеств set1 и set2.
	set1 := make(map[int]bool, 20)
	set2 := make(map[int]bool, 20)

	// Заполнение множеств случайными элементами.
	for i := 0; i < 20; i++ {
		key1 := rand.Intn(100) // Генерация случайного индекса для первого множества.
		key2 := rand.Intn(100) // Генерация случайного индекса для второго множества.
		set1[key1] = true      // Добавление элемента в первое множество.
		set2[key2] = true      // Добавление элемента во второе множество.
	}
	// Вывод на экран элементов обоих множеств.
	fmt.Println("Первое множество:", set1)
	fmt.Println("Второе множество:", set2)

	// Нахождение и вывод пересечения множеств с использованием двух разных методов.
	fmt.Println("Пересечение методом 1:", intersection1(set1, set2))
	fmt.Println("Пересечение методом 2:", intersection2(set1, set2))

}

// Метод intersection1 находит пересечение множеств, используя мапу для хранения результатов.
func intersection1(m1, m2 map[int]bool) (mapp map[int]bool) {
	mapp = make(map[int]bool)

	for key := range m1 {
		if m2[key] {
			mapp[key] = true
		}
	}
	return
}

// Метод intersection2 также находит пересечение множеств, но использует мапу для подсчета количества вхождений элементов.
// Затем он создает срез и добавляет элементы, которые встречаются дважды (присутствуют в обоих множествах).
func intersection2(m1, m2 map[int]bool) []int {
	mapp := make(map[int]int)
	slice := make([]int, 0)

	// Подсчет количества вхождений элементов из первого множества.
	for key := range m1 {
		mapp[key] += 1
	}

	// Подсчет количества вхождений элементов из второго множества.
	for key := range m2 {
		mapp[key] += 1
	}

	// Проверка, какие элементы встречаются дважды (пересекаются) и добавление их в срез.
	for key := range mapp {
		if mapp[key] == 2 {
			slice = append(slice, key)
		}
	}

	return slice
}

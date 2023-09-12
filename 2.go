package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем массив чисел
	arr := [5]int{2, 4, 6, 8, 10}

	// Запускаем функции для конкурентных вычислений
	runAllGoroutines(arr)
	runMultipleGoroutines(arr, 2)
}

// Функция для вычисления квадрата числа и вывода его в stdout
func Square(value int, group *sync.WaitGroup) {
	defer group.Done()
	fmt.Println(value * value)
}

// Функция для запуска горутин для всех элементов массива
func runAllGoroutines(arr [5]int) {
	var wg sync.WaitGroup

	// Итерируемся по массиву и запускаем горутины для каждого элемента
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go Square(arr[i], &wg)
	}
	wg.Wait() // Ждем завершения всех горутин
	fmt.Println("End!")
}

// Функция для запуска горутин с ограничением по количеству одновременных горутин (runners)
func runMultipleGoroutines(arr [5]int, runners int) {
	var wg sync.WaitGroup
	wg.Add(runners)
	var c = make(chan int)

	// Создаем и запускаем горутины для обработки элементов из канала
	for i := 0; i < runners; i++ {
		go func(cha <-chan int, group *sync.WaitGroup) {
			defer group.Done()
			for value := range cha {
				fmt.Println(value * value)
			}
		}(c, &wg)
	}

	// Помещаем элементы массива в канал
	for _, v := range arr {
		c <- v
	}
	close(c)  // Закрываем канал, чтобы завершить горутины
	wg.Wait() // Ждем завершения всех горутин
	fmt.Println("End!")
}

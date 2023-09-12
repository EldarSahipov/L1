package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	a(arr)    // Вызываем функцию a с горутинами
	b(arr, 5) // Вызываем функцию b с определенным количеством горутин
}

func a(arr [5]int) {
	var wg sync.WaitGroup
	var mutex sync.Mutex // Мьютекс для синхронизации доступа к sum
	sum := 0             // Переменная для хранения суммы квадратов чисел
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(value int, group *sync.WaitGroup) {
			defer wg.Done()
			mutex.Lock() // Захватываем мьютекс перед обновлением sum
			sum += value * value
			mutex.Unlock() // Освобождаем мьютекс после обновления sum
		}(arr[i], &wg)
	}
	wg.Wait()        // Ждем завершения всех горутин
	fmt.Println(sum) // Выводим сумму
}

func b(arr [5]int, runners int) {
	var wg sync.WaitGroup
	var mutex sync.Mutex // Мьютекс для синхронизации доступа к sum
	ch := make(chan int) // Канал для передачи значений в горутины
	sum := 0             // Переменная для хранения суммы квадратов чисел

	wg.Add(runners)
	for i := 0; i < runners; i++ {
		go func(i int) {
			defer wg.Done()
			for value := range ch {
				square := value * value
				mutex.Lock() // Захватываем мьютекс перед обновлением sum
				sum += square
				mutex.Unlock() // Освобождаем мьютекс после обновления sum
			}
		}(i)
	}
	for i := 0; i < len(arr); i++ {
		ch <- arr[i] // Помещаем элементы массива в канал
	}
	close(ch) // Закрываем канал после отправки всех значений в горутины

	wg.Wait()        // Ждем завершения всех горутин
	fmt.Println(sum) // Выводим сумму
}

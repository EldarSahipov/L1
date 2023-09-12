package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	input := make(chan int)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11} // Канал для входных данных (числа)
	out := make(chan int)                             // Канал для результатов операции умножения на 2

	// Горутина для умножения чисел на 2 и отправки в out канал.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range input {
			out <- v * 2
		}
		close(out) // Закрываем out канал после обработки всех чисел.
	}()

	// Горутина для вывода чисел из out канала в stdout.
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range out {
			fmt.Println(v)
		}
	}()

	// Запись чисел в input канал.
	for _, v := range slice {
		input <- v
	}
	close(input) // Закрываем input канал, чтобы завершить работу первой горутины.
	wg.Wait()    // Ожидание завершения всех горутин.
}

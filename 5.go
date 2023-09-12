package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	timeSecond = 3   // Время работы программы в секундах.
	maxValue   = 100 // Максимальное случайное значение, которое отправляется в канал.
)

func main() {
	wg := sync.WaitGroup{}  // WaitGroup для отслеживания завершения горутины.
	ch := make(chan int, 1) // Создаем буферизированный канал с емкостью 1.

	ticker := time.NewTicker(time.Second * time.Duration(timeSecond)) // Таймер с периодом времени.
	defer ticker.Stop()                                               // Остановка таймера после завершения работы.

	wg.Add(1)
	go goroutine(ch, &wg) // Запуск горутины для чтения данных из канала.

	for {
		select {
		case <-ticker.C:
			close(ch) // Закрытие канала при истечении времени.
			wg.Wait() // Ожидание завершения горутины.
			fmt.Println("Program completed")
			return
		default:
			ch <- rand.Intn(maxValue) // Отправка случайных значений в канал.
		}
	}
}

func goroutine(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v) // Чтение и вывод значений из канала.
	}
}

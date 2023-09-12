package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var (
	workers = 10 // Количество воркеров, которые будут читать данные из канала.
)

func main() {
	// Создаем канал для передачи данных между главным потоком и воркерами.
	ch := make(chan int)

	// Создаем канал для приема сигналов завершения (Ctrl+C)
	quit := make(chan os.Signal, 1)

	// Регистрируем обработчики сигналов SIGINT и SIGTERM.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Создаем WaitGroup для отслеживания завершения работы всех воркеров.
	wg := sync.WaitGroup{}

	// Запускаем воркеры.
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(ch, &wg, i)
	}

	// Основной цикл главного потока.
	for {
		select {
		case <-quit: // Получение сигнала завершения
			close(ch) // Закрываем канал, чтобы завершить работу воркеров.
			fmt.Println("Channel closed")
			wg.Wait() // Ожидаем завершения всех воркеров.
			fmt.Println("Goroutines have completed their work")
			return
		default:
			ch <- rand.Intn(100)               // Генерируем случайное число и отправляем в канал.
			time.Sleep(time.Millisecond * 300) // Задержка между отправками данных.
		}
	}
}

func worker(ch chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v) // Воркер читает данные из канала и выводит их.
	}
	fmt.Println("Goroutine", i, "completed work")
}

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Создание контекста и горутины для writeValue с использованием context.
	ctx, cancel := context.WithCancel(context.Background())
	go1 := Goroutine{wg: sync.WaitGroup{}}
	ch := make(chan int, 1)
	go1.wg.Add(2)
	go writeValue(ctx, ch, &go1.wg)
	go go1.chanelClose(ch)

	// Закрытие канала и ожидание завершения горутины.
	time.Sleep(time.Second * 1)
	close(ch)
	cancel()
	go1.wg.Wait()

	// Повторяем аналогичные шаги для остальных способов остановки горутин.
	// (chanelCloseWithRange, useAdditionalChannel, useSignal, useTime)

	// Создание контекста и горутины для chanelCloseWithRange с использованием context.
	ctx, cancel = context.WithCancel(context.Background())
	go2 := Goroutine{wg: sync.WaitGroup{}}
	ch = make(chan int, 2)
	go2.wg.Add(2)
	go writeValue(ctx, ch, &go2.wg)
	go go2.chanelCloseWithRange(ch)

	// Закрытие канала и ожидание завершения горутины.
	time.Sleep(time.Second * 1)
	close(ch)
	cancel()
	go2.wg.Wait()

	// Создание контекста и горутины для useAdditionalChannel с использованием context.
	ctx, cancel = context.WithCancel(context.Background())
	go3 := Goroutine{wg: sync.WaitGroup{}}
	ch = make(chan int, 1)
	done := make(chan bool, 1)
	go3.wg.Add(2)
	go writeValue(ctx, ch, &go3.wg)
	go go3.useAdditionalChannel(ch, done)

	// Отправка сигнала в done канал и ожидание завершения горутины.
	time.Sleep(time.Second * 1)
	done <- true
	close(done)
	cancel()
	close(ch)
	go3.wg.Wait()

	// Создание контекста и горутины для useSignal с использованием context.
	ctx, cancel = context.WithCancel(context.Background())
	go4 := Goroutine{wg: sync.WaitGroup{}}
	ch = make(chan int, 1)
	go4.wg.Add(2)
	go writeValue(ctx, ch, &go4.wg)
	go go4.useSignal(ctx, ch)

	// Отмена контекста и ожидание завершения горутины.
	time.Sleep(time.Second * 1)
	cancel()
	close(ch)
	go4.wg.Wait()

	// Создание контекста и горутины для useTime с использованием context.
	ctx, cancel = context.WithCancel(context.Background())
	go5 := Goroutine{wg: sync.WaitGroup{}}
	ch = make(chan int, 1)
	go5.wg.Add(2)
	go writeValue(ctx, ch, &go5.wg)
	go go5.useTime(ch)

	// Отмена контекста и ожидание завершения горутины.
	time.Sleep(1 * time.Second)
	cancel()
	close(ch)
	go5.wg.Wait()
}

// Goroutine - структура, представляющая горутину и используемую для ожидания ее завершения.
type Goroutine struct {
	wg sync.WaitGroup
}

// writeValue - горутина, которая пишет случайные значения в канал до получения сигнала от контекста ctx.Done().
func writeValue(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return // Горутина завершает работу при получении сигнала от контекста.
		default:
			ch <- rand.Intn(10)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// chanelClose - горутина, которая читает из канала до его закрытия.
func (g *Goroutine) chanelClose(ch chan int) {
	defer g.wg.Done()
	for {
		if v, ok := <-ch; ok {
			fmt.Println(v)
		} else {
			fmt.Println("closed chanel")
			return
		}
	}
}

// chanelCloseWithRange - горутина, которая читает из канала с использованием цикла range до закрытия канала.
func (g *Goroutine) chanelCloseWithRange(ch chan int) {
	defer g.wg.Done()

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("closed chanel (range)")
	return
}

// useAdditionalChannel - горутина, которая читает из канала до получения сигнала из done канала.
func (g *Goroutine) useAdditionalChannel(ch chan int, done chan bool) {
	defer g.wg.Done()
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

// useSignal - горутина, которая читает из канала до получения сигнала от контекста ctx.Done().
func (g *Goroutine) useSignal(ctx context.Context, ch chan int) {
	defer g.wg.Done()

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-ctx.Done():
			fmt.Println("context cancel")
			return
		}
	}
}

// useTime - горутина, которая читает из канала до получения сигнала от таймера ticker.
func (g *Goroutine) useTime(ch chan int) {
	defer g.wg.Done()

	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-ticker.C:
			fmt.Println("time!")
			return
		}
	}
}

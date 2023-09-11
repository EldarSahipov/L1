package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	timeSecond = 3
	maxValue   = 100
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 1)

	ticker := time.NewTicker(time.Second * time.Duration(timeSecond))
	defer ticker.Stop()

	wg.Add(1)
	go goroutine(ch, &wg)

	for {
		select {
		case <-ticker.C:
			close(ch)
			wg.Wait()
			fmt.Println("Program completed")
			return
		default:
			ch <- rand.Intn(maxValue)
		}
	}
}

func goroutine(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}

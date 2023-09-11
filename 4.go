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

func main() {
	const workers = 10
	ch := make(chan int)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	wg := sync.WaitGroup{}

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(ch, &wg, i)
	}

	for {
		select {
		case <-quit:
			close(ch)
			fmt.Println("Channel closed")
			wg.Wait()
			fmt.Println("Goroutines have completed their work")
			return
		default:
			ch <- rand.Intn(100)
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func worker(ch chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("Goroutine", i, "completed work")
}

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go1 := Goroutine{wg: sync.WaitGroup{}}
	ch := make(chan int, 1)
	go1.wg.Add(2)
	go writeValue(ctx, ch, &go1.wg)
	go go1.chanelClose(ch)
	time.Sleep(time.Second * 1)
	close(ch)
	cancel()
	go1.wg.Wait()

	ctx, cancel = context.WithCancel(context.Background())
	go2 := Goroutine{wg: sync.WaitGroup{}}
	ch = make(chan int, 2)
	go2.wg.Add(2)
	go writeValue(ctx, ch, &go2.wg)
	go go2.chanelCloseWithRange(ch)
	time.Sleep(time.Second * 1)
	close(ch)
	cancel()
	go2.wg.Wait()

	ctx, cancel = context.WithCancel(context.Background())
	go3 := Goroutine{wg: sync.WaitGroup{}}
	ch = make(chan int, 1)
	done := make(chan bool, 1)
	go3.wg.Add(2)
	go writeValue(ctx, ch, &go3.wg)
	go go3.useAdditionalChannel(ch, done)
	time.Sleep(time.Second * 1)
	done <- true
	close(done)
	cancel()
	close(ch)
	go3.wg.Wait()

	ctx, cancel = context.WithCancel(context.Background())
	go4 := Goroutine{wg: sync.WaitGroup{}}
	ch = make(chan int, 1)
	go4.wg.Add(2)
	go writeValue(ctx, ch, &go4.wg)
	go go4.useSignal(ctx, ch)
	time.Sleep(time.Second * 1)
	cancel()
	close(ch)
	go4.wg.Wait()

	ctx, cancel = context.WithCancel(context.Background())
	go5 := Goroutine{wg: sync.WaitGroup{}}
	ch = make(chan int, 1)
	go5.wg.Add(2)
	go writeValue(ctx, ch, &go5.wg)
	go go5.useTime(ch)
	time.Sleep(1 * time.Second)
	cancel()
	close(ch)
	go5.wg.Wait()
}

type Goroutine struct {
	wg sync.WaitGroup
}

func writeValue(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- rand.Intn(10)
			time.Sleep(200 * time.Millisecond)
		}
	}
}

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

func (g *Goroutine) chanelCloseWithRange(ch chan int) {
	defer g.wg.Done()

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("closed chanel (range)")
	return
}

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

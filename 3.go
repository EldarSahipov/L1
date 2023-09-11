package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	a(arr)
	b(arr, 2)
}

func a(arr [5]int) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	sum := 0
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(value int, group *sync.WaitGroup) {
			defer wg.Done()
			mutex.Lock()
			sum += value * value
			mutex.Unlock()
		}(arr[i], &wg)
	}
	wg.Wait()
	fmt.Println(sum)
}

func b(arr [5]int, runners int) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	ch := make(chan int)
	sum := 0

	wg.Add(runners)
	for i := 0; i < runners; i++ {
		go func(i int) {
			defer wg.Done()
			for value := range ch {
				square := value * value
				mutex.Lock()
				sum += square
				mutex.Unlock()
			}
		}(i)
	}
	for i := 0; i < len(arr); i++ {
		ch <- arr[i]
	}
	close(ch) // Закрыть канал после отправки всех значений.

	wg.Wait()
	fmt.Println(sum)
}

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	input := make(chan int)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	out := make(chan int)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range input {
			out <- v * 2
		}
		close(out)

	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range out {
			fmt.Println(v)
		}
	}()

	for _, v := range slice {
		input <- v
	}
	close(input)
	wg.Wait()
}

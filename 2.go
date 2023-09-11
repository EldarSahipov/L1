package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	runAllGoroutines(arr)
	runMultipleGoroutines(arr, 2)
}

func Square(value int, group *sync.WaitGroup) {
	defer group.Done()
	fmt.Println(int(math.Pow(float64(value), 2)))
}

func runAllGoroutines(arr [5]int) {
	var wg sync.WaitGroup
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go Square(arr[i], &wg)
	}
	wg.Wait()
	fmt.Println("End!")
}

func runMultipleGoroutines(arr [5]int, runners int) {
	var wg sync.WaitGroup
	wg.Add(runners)
	var c = make(chan int, 666)
	for i := 0; i < runners; i++ {
		go func(cha <-chan int, group *sync.WaitGroup) {
			defer group.Done()
			for value := range cha {
				fmt.Println(value * value)
			}
		}(c, &wg)
	}

	for _, v := range arr {
		c <- v
	}
	close(c)
	wg.Wait()

}

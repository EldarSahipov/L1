package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

var (
	workers = 9999
)

func main() {
	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}
	mapp := make(map[int]string)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go writeToMap(mapp, &rw, &wg)
	}

	wg.Wait()
	fmt.Println(mapp)
	fmt.Println("End!")

}

func writeToMap(mapp map[int]string, rw *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	key, str := genValues()
	rw.Lock()
	defer rw.Unlock()
	mapp[key] = str
}

func genValues() (key int, str string) {
	key = rand.Intn(100)
	str = strconv.Itoa(rand.Intn(100))
	return
}

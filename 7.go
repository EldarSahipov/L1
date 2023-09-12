package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

var (
	workers = 100
)

func main() {
	wg := sync.WaitGroup{}
	rw := sync.RWMutex{}         // Создаем RWMutex для защиты доступа к map.
	mapp := make(map[int]string) // Создаем пустую map для хранения данных.

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go writeToMap(mapp, &rw, &wg) // Запускаем горутины для записи данных в map.
	}

	wg.Wait()
	fmt.Println(mapp)
	fmt.Println("End!")

}

func writeToMap(mapp map[int]string, rw *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	key, str := genValues() // Генерируем случайные ключ и строку для записи в map.
	rw.Lock()               // Блокируем мьютекс для записи данных в map.
	defer rw.Unlock()       // Разблокируем мьютекс после завершения операции
	mapp[key] = str         // Записываем данные в map.
}

func genValues() (key int, str string) {
	key = rand.Intn(100)               // Генерируем случайный ключ (целое число до 100).
	str = strconv.Itoa(rand.Intn(100)) // Генерируем случайную строку (целое число до 100, преобразованное в строку).
	return
}

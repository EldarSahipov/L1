package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Counter представляет счетчик с использованием sync.RWMutex.
type Counter struct {
	rw    *sync.RWMutex
	count int
}

// NewCounter создает новый экземпляр Counter и инициализирует его.
func NewCounter() *Counter {
	return &Counter{
		rw:    &sync.RWMutex{},
		count: 0,
	}
}

// Increment инкрементирует счетчик на единицу в защищенной блокировкой области.
func (c *Counter) Increment() {
	c.rw.Lock()
	defer c.rw.Unlock()
	c.count++
}

// GetCount возвращает текущее значение счетчика в защищенной блокировкой области.
func (c *Counter) GetCount() int {
	c.rw.RLock()
	defer c.rw.RUnlock()
	return c.count
}

// CounterAtomic представляет счетчик с использованием sync/atomic.
type CounterAtomic struct {
	count atomic.Int64
}

// NewCounterAtomic создает новый экземпляр CounterAtomic и инициализирует его.
func NewCounterAtomic() *CounterAtomic {
	return &CounterAtomic{count: atomic.Int64{}}
}

// IncrementCountAtomic инкрементирует счетчик на единицу с использованием atomic операции.
func (c *CounterAtomic) IncrementCountAtomic() {
	c.count.Add(1)
}

// GetCountAtomic возвращает текущее значение счетчика с использованием atomic операции.
func (c *CounterAtomic) GetCountAtomic() int64 {
	return c.count.Load()
}

func main() {
	start := time.Now()
	wg := sync.WaitGroup{}
	counter := NewCounter()

	// Запуск 1000 горутин для инкрементирования счетчика с использованием sync.RWMutex.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}
	wg.Wait()

	// Вывод значения счетчика и времени выполнения.
	fmt.Println(counter.GetCount())
	fmt.Println(time.Now().Sub(start).Seconds())

	start = time.Now()
	wg = sync.WaitGroup{}
	counterAtomic := NewCounterAtomic()

	// Запуск 1000 горутин для инкрементирования счетчика с использованием sync/atomic.
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counterAtomic.IncrementCountAtomic()
		}()
	}
	wg.Wait()

	// Вывод значения счетчика и времени выполнения.
	fmt.Println(counterAtomic.GetCountAtomic())
	fmt.Println(time.Now().Sub(start).Seconds())
}

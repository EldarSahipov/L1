package main

import (
	"fmt"
	"reflect"
)

func main() {
	objects := []any{true, "a", make(chan int), 5, func() {}}

	// 1 способ через пакет fmt. Под капотом он тоже использует reflect
	for _, v := range objects {
		xType := fmt.Sprintf("%T", v)
		fmt.Println(xType)
	}

	// 2 способ через reflect
	for _, v := range objects {
		fmt.Println(reflect.TypeOf(v))
	}
}

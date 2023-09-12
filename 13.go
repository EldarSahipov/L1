package main

import "fmt"

func main() {
	a, b := 0, 666

	b, a = a, b

	fmt.Println(a, b)
}

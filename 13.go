package main

import "fmt"

func main() {
	a := 0
	b := 666

	b, a = a, b

	fmt.Println(a, b)
}

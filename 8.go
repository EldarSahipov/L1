package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num, i int64
	i = 5
	num = 6546

	fmt.Println(strconv.FormatInt(num, 2))
	fmt.Println(strconv.FormatInt(i, 2))

	num |= 1 << i
	fmt.Println(strconv.FormatInt(num, 2))

	num &^= 1 << i
	fmt.Println(strconv.FormatInt(num, 2))
}

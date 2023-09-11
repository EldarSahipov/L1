package main

import "fmt"

func main() {
	arr := []int{1, 5, 7, 56, 57, 155, 544, 1058, 5455}
	fmt.Println(bin(arr, 5455))

}

func bin(arr []int, n int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := (right + left) / 2
		if n == arr[middle] {
			return middle
		} else if n > arr[middle] {
			left = middle + 1
		} else if n < arr[middle] {
			right = middle - 1
		}
	}
	return -1
}

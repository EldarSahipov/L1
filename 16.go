package main

import "fmt"

func main() {
	slice := []int{51, 5, 8, 1, 514, 545, 4, 15, 1545, 6465, 487, 1, 21, 654, 6546, 161, 64, 984, 7, 351, 65, 465, 46, 84, 6451, 65, 416, 5168}
	quickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)

}

func quickSort(s []int, l, r int) {
	p := (r - l) / 2
	if s[l] > s[p] {
		s[l], s[p] = s[p], s[l]
		l++
	}
	if s[r] < s[p] {
		s[r], s[p] = s[p], s[r]
		r--
	}
	quickSort(s, l, r)
}

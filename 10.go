package main

import "fmt"

func main() {
	slice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	mapTemp := make(map[int][]float64)

	for _, v := range slice {
		step := int(v) / 10 * 10
		mapTemp[step] = append(mapTemp[step], v)
	}

	fmt.Println(mapTemp)
}

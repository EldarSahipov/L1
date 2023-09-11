package main

import "fmt"

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}

	mySet := make(map[string]bool)

	for _, value := range slice {
		mySet[value] = true
	}

	fmt.Print("Множество: ")
	for key := range mySet {
		fmt.Print(key, ", ")
	}
}

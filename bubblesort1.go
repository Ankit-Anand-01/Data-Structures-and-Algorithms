package main

import "fmt"

func bubblesort(input []int) {
	n := len(input)
	swap := true
	for swap {
		swap = false
		for i := 1; i < n; i++ {
			if input[i-1] > input[i] {
				fmt.Println("swapping")
				input[i-1], input[i] = input[i], input[i-1]
				swap = true
			}
		}
	}
	fmt.Println("sorted array:", input)
}
func main() {
	var a []int = []int{4, 8, 23, 0, 5, 9, 3, 15}
	b := []int{1, 53, 8, 6, 4, 0}

	bubblesort(a)
	bubblesort(b)
}

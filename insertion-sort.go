package main

import "fmt"

func InsertionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < i+1; j++ {
			//compare element present at index i with every element present
			//  left of it place it in right place so that array on the
			//left remains   sorted
			if numbers[j] > numbers[i] {

				intermediate := numbers[j]
				numbers[j] = numbers[i]
				numbers[i] = intermediate
			}
		}
		//fmt.Println(numbers)
	}
	return numbers
}

func main() {
	k := []int{5, 9, 3, 11, 2, 4}
	fmt.Println(InsertionSort(k))
}

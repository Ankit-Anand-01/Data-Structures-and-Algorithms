package main

import "fmt"

func Binarysearch(data []int, value int) bool {
	size := len(data)
	low := 0
	high := size - 1
	mid := 0
	for low <= high {
		mid = low + (high-low)/2
		if data[mid] == value {
			return true
		} else if data[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

func main() {
	a := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	fmt.Println(Binarysearch(a, 6))
}

//Mergesort implementation in Golang
package main

import (
	"fmt"
)
//defining func for mergesort
func MergeSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}
	mid := int(len(numbers) / 2)
	//calling mergesort func recursively 
	return merge(MergeSort(numbers[mid:]), MergeSort(numbers[:mid]))
}
//defining another func for merging
func merge(a, b []int) []int {
	size := len(a) + len(b)
	i := 0
	j := 0
	res := make([]int, size)
	for k := 0; k < size; k++ {
		//switch case
		switch true {
		case i == len(a):
			res[k] = b[j]
			j++
		case j == len(b):
			res[k] = a[i]
			i++
		case a[i] > b[j]:
			res[k] = b[j]
			j++
		case a[i] < b[j]:
			res[k] = a[i]
			i++
		case a[i] == b[j]:
			res[k] = a[i]
			i++
		}
	}
	return res

}
func main() {
	a := []int{8, 5, 2, 7, 5, 12}
	//calling function for mergesort
	fmt.Println(MergeSort(a))

}

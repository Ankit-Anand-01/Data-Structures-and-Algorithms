//Quick sort in Golang
package main

import (
	"fmt"
)

//defining func for dividing the arr into low and high
func partition(a []int, lo, hi int) int {
	p := a[hi]
	for j := lo; j < hi; j++ {
		if a[j] < p {
			a[j], a[lo] = a[lo], a[j]
			lo++
		}
	}

	a[lo], a[hi] = a[hi], a[lo]
	return lo
}
//defining func for quicksorting
func quickSort(a []int, lo, hi int) {
	if lo > hi {
		return
	}

	p := partition(a, lo, hi)
	quickSort(a, lo, p-1)
	quickSort(a, p+1, hi)
}

func main() {
	list := []int{7, 9, 3, 1, 0, 4, 9, 6}

	quickSort(list, 0, len(list)-1)
	fmt.Println(list)
}

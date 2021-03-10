//linear search in Golang
package main

import "fmt"
//defining func for linear search
func linearsearch(list []int, key int) bool {
	for _, item := range list {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	item := []int{8, 5, 6, 44, 10}
	//calling func
	fmt.Println(linearsearch(item, 15))
}

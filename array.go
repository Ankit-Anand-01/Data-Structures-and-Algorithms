//array in Golang
package main

import "fmt"

func main() {
	// Creating an array of string type
	// Using var keyword
	var arr [10]int
	//puting the elements in the array using index value
	arr[0] = 12
	arr[1] = 24
	arr[2] = 36
	arr[3] = 48
	arr[4] = 60
	arr[5] = 72
	arr[6] = 84
	arr[7] = 96
	arr[8] = 108
	arr[9] = 120
	fmt.Println(arr)
	//shorthand declaration of array
	array := [5]string{"Ankit", "Rahul", "Vinay", "Shubham", "Prem"}
	fmt.Println(array)
	//print a element using index value
	fmt.Println(array[0])
	// Accessing the elements of
	// the array Using for loop
	for i := 0; i < 5; i++ {
		fmt.Println(array[i])
	}
}

//multi-dimensional array in Golang
package main

import "fmt"

func main() {
	//creating and initializing 2D array
	arr2D := [2][2]int{{2, 4}, {4, 8}}
	fmt.Println(arr2D)
	//creating and initializing 3D array
	arr3D := [2][2][2]int{{{4, 5}, {6, 7}}, {{8, 9}, {10, 11}}}
	fmt.Println(arr3D)
	//print a element using index value in 3D array
	fmt.Println(arr3D[1][1][0])

}

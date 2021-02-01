//slice in Golang
package main

import "fmt"

func main() {
	slice := []int{5, 10, 15, 20, 25}
	fmt.Println(slice)
	//define slice using make() func
	S := make([]int, 5)
	S[0] = 3
	S[1] = 6
	S[2] = 9
	S[3] = 12
	S[4] = 15
	fmt.Println(S)
	//add to a slice
	S = append(S, 18)
	fmt.Println(S)
	//appent more than one
	S = append(S, 21, 24, 27, 30)
	fmt.Println(S)
	//delete from a slice
	S = append(S[:3], S[3+1:]...)
	fmt.Println(S)
	//replace a element
	S[2] = S[len(S)-2]
	fmt.Println(S)
}

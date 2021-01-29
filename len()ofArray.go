package main

//length of an array
import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	//length of the array by using len() method
	fmt.Println(len(arr))
	//if ellipsis ‘‘…’’ become visible at the place of length, then the length of the array is equal to the number of the elements.
	new_arr := [...]string{"virat", "Dravid", "Sachin", "Dhoni", "Karthik"}
	fmt.Println(len(new_arr))
}

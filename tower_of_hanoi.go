//Tower of Hanoi in Golang
package main

import (
	"fmt"
)

//func for disk management using recursion
func diskmanager(num int, from string, to string, temp string) {
	if num < 1 {
		return
	}
	//calling func recursively
	diskmanager(num-1, from, temp, to)
	fmt.Println("move disk", num, "kg from rod", from, "to rod", to)
	diskmanager(num-1, temp, to, from)
}
func TowerOfHanoi(num int) {
	diskmanager(num, "A", "C", "B")
}
func main() {
	TowerOfHanoi(3)
}

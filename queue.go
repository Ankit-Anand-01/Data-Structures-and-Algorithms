//Queue implementation in Go
package main

import "fmt"

//Queue represents a queue that holds a slice
type Queue struct {
	items []int
}

//Enqueue adds the item in the Queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)

}

//Dequeue removes a value at the front
func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	//adding values in the queue using enqueue func
	myQueue.Enqueue(11)
	myQueue.Enqueue(22)
	myQueue.Enqueue(33)
	myQueue.Enqueue(44)
	myQueue.Enqueue(55)
	myQueue.Enqueue(66)
	myQueue.Enqueue(77)
	myQueue.Enqueue(88)
	fmt.Println(myQueue)
	//removing value using dequeue
	myQueue.Dequeue()
	fmt.Println(myQueue)

}

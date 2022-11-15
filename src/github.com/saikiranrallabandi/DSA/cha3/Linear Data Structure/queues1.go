package main
import (
	"fmt"
)


// Queue represents a queue that holds a slice

type Queue struct {
	items []int
}

//Enqueue adds a value at the end
func (q *Queue) Enqueue(i int)  {
	q.items = append(q.items, i)
}

//Dequeu removes a value at the front
// and Returns the removed value

func (q *Queue) Dequeu() int  {
	toRemove := q.items[0]
	q.items = q.items[:1]
	return toRemove

}


func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	fmt.Println(myQueue)
	myQueue.Dequeu()
	fmt.Println(myQueue)

}
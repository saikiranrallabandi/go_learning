package main

import (
	"container/ring"
	"fmt"
)
func main() {
	var integers []int
	integers = []int{1,3,5,7}
	var circular_list *ring.Ring
	circular_list= ring.New(len(integers))
	var i int
	for i=0; i < circular_list.Len(); i++ {
		circular_list.Value = integers[i]
		circular_list = circular_list.Next()
	}
	circular_list.Do(func(element interface{}) {
			fmt.Println(element,",")
		})
		fmt.Println()
    // revers of the circular list
	for i := 0; i < circular_list.Len(); i++ {
		fmt.Println(circular_list.Value,",")
		circular_list = circular_list.Prev()
	}
	fmt.Println()
	// move two elements forward in the circular list
	circular_list = circular_list.Move(2)
	circular_list.Do(func (element interface{})  {
		fmt.Println(element,",")
	})
	fmt.Println()
}
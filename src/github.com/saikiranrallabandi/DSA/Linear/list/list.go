package main
import (
	"fmt"
	"container/list"
)
func main() {
	// Creating the list
	var intList list.List
	//Pushing element to the list
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(34)

	// list is iterated through the for loop, and elment value is accessed
	for element := intList.Front(); element != nil; element=element.Next() {
		fmt.Println(element.Value.(int))

		if element.Value !=1 {
			// Remove elements to the list
			intList.Remove(element)
		}
	}
}
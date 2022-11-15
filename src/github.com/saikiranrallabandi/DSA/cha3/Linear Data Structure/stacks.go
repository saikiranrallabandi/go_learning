package main
import (
	"fmt"
)

type Stack struct {
	items[]int
}

//push will add the value at the end
func (s *Stack) Push(i int)  {
     s.items = append(s.items,i)
}




//pop will remove a value at the end
// and Returns the removed value
func (s *Stack) Pop() int {
	toRemove := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return toRemove


}



func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)

}
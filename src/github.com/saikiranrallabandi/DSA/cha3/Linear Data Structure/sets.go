package main
import (
	"fmt"
)
//set Class
type Set struct {
	integerMap map[int]bool
}

//create the map of integer and bool
func (set *Set) New() {
   set.integerMap = make(map[int]bool)
}

// adds the element to the set
func (set *Set) AddElement(element int){
 if !set.ContainsElement(element) {
  set.integerMap[element] = true
 }
}


// DeleteElement method
func (set *Set) DeleteElement(element int)  {
	delete(set.integerMap, element)

}

//check if element is in the set
func (set *Set) ContainsElement(element int) bool  {
     var exists bool
	 _, exists = set.integerMap[element]
	 return exists
}

func main() {
	var set *Set
	set = &Set{}
	set.New()
	set.AddElement(1)
	set.AddElement(2)
	fmt.Println(set)
	fmt.Println(set.ContainsElement(1))
}
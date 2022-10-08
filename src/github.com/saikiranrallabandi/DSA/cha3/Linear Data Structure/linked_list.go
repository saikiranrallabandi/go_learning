package main
import (
	"fmt"
)

// //Node class
type Node struct {
	property int
	nextNode *Node
}

// //LinkedList class
type LinkedList struct {
	headNode *Node
}

//AddToHead method of LinkedList class
// Addto head method adds the node to start of linked list. The addtohead method of LinkedList class
//has a parameter integer
func (linkedList *LinkedList) AddToHead(property int)  {
	//initialize new Node
	var node = Node{} // a new node is instantiated
	node.property = property
	if node.nextNode !=nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node

}




func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.AddToHead(1)
	linkedList.AddToHead(3)
	fmt.Println(linkedList.headNode.property)
}
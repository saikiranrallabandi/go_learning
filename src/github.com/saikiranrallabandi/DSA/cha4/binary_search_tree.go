package main
import (
	"fmt"
)

// Node represents the components of a binary search tree
type Node struct {
	Key int
	Left *Node
	Right *Node
}
// Inseart will add a node to the tree
// the key to add should not be already
func (n *Node) Insert (k int)  {
	if n.Key < k {
		//move right
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	}else if n.Key > k {
			//move left
			if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
		}
	}
// Search will take in a key value
// and RETURN true if there is a node with that value
func (n *Node) Search (k int) bool{

	if n == nil {
		return false
	}

   if n.Key < k {
	// move right
	return n.Right.Search(k)
   } else if n.Key > k {
	//move left
	return n.Left.Search(k)
   }
   return true
}

func mains() {
	tree := &Node{Key: 100}
	tree.Insert(50)
	tree.Insert(150)
	tree.Insert(250)
	tree.Insert(550)
	tree.Insert(1150)
	tree.Insert(950)

	fmt.Println(tree.Search(350))




}
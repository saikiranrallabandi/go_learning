/*
Problem: How will you convert a binary tree into a doubly-linked list
Time complexity: O(N). Each node is processed exactly once. Space complexity: O(N). We have to keep a recursion stack of the size of the tree height

We use Depth First Search In Order Traversal.

inOrder(node.Left) node.Val inOrder(node.Right)
Here is the algorithm :

Initiate the first and the last nodes as nulls.
Call the standard inorder recursion helper(root) :
            If node is not null :
                Call the recursion for the left subtree helper(node.left).
                If the last node is not null, link the last and the current node nodes.
                Else initiate the first node.
                Mark the current node as the last one : last = node.
                Call the recursion for the right subtree helper(node.right).
                Link the first and the last nodes to close DLL ring and then return the first node
*/

package main
import (
	"fmt"
)

type Node struct {
	val int
	Left *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return root
	}

	var first, last *Node

	//function declaration
	var dfs func(node *Node)

	//function implementation
	// dfs - in order
	// inorder(node.Left)
	// node.val
	// inorder(node.Right)

	dfs = func(node *Node) {
		if node !=nil {
			dfs(node.Left)

			if last !=nil {
				//draw last node and current node connections
				last.Right = node
				node.Left = last
			} else {
				// if there is no last node, you are dealing with the very first node!
				first = node
			}
			// all done, now the last node is the current node
			last = node
			dfs(node.Right)
		}
	}
	// function call
	dfs(root)

	//making it circular
	last.Right = first
	first.left = last

	return first


	}


func main() {
}
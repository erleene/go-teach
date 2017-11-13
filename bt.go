package main

import (
	"errors"
	"fmt"
)

//a binary tree is made up of nodes
//each node contains a left and right pointer and a data element
//root pointer points to the topmost node of the tree
//left and right of tree recursively points to smaller "subtrees" on either side
//null pointer represents a binary tree with no elements = empty tree

type Node struct {
	Left  *Node
	Right *Node
	Value int
	Data  string
}

func (t *Node) Insert(val int, data string) *Node {
	//we want to insert a value to the tree
	//compare new value with current node's value
	//check if the value is the same with the left
	switch {
	case val < t.Value:
		//Left node
		if t.Left == nil { // if there's on left child, create a new one
			t.Left = &Node{Value: val, Data: data}
			return t.Left
		} else {
			t.Left.Insert(val, data) //repeat method.
		}
	case val > t.Value:
		//Right node
		if t.Right == nil { // if there's on left child, create a new one
			t.Right = &Node{Value: val, Data: data}
			return t.Right
		} else {
			t.Right.Insert(val, data) //repeat method.
		}
	default:
		return nil
	}
	return nil
}

//func (t *Tree) Delete(val) int {
//delete from the tree a given value if it exists
//assume the node to be deleted is the right child of it's parent node
// steps also apply if the node is the left child; you only need to swap "right"
//for "left", and "large" for "small"
//
//parents must not be nil
//if the node has no children , remote it from its parents
// //if val < t.Value {
// 	return t.Left.Delete(val)
// }
// if val > t.Value {
// 	return t.Right.Delete(val)
// }
//
// }

//create helper functions
func (t *Node) findMax(n *Node) (*Node, *Node) {
	if n.Right == nil {
		//we're at the bottom
		return n, t
	}
	return n.Right.findMax(n)
}

//replace node function
func (current *Node) replaceNode(parent *Node, r *Node) error {
	//
	if current == nil {
		return errors.New("No node")
	}
	//
	if current == parent.Left {
		parent.Left = r
	}
	parent.Right = r
	return nil
}

func (current *Node) deleteNode(parent *Node) error {
	//
	if current == nil {
		return errors.New("No current node")
	}
	switch {
	case parent.Value < current.Value:
		return current.Left.deleteNode(current)
	case parent.Value > current.Value:
		return current.Right.deleteNode(parent)
	}
	if current.Left == nil && current.Right == nil {
		current.replaceNode(parent, nil)
		return nil
	}
	//leaf node deletion
	if current.Left == nil {
		current.replaceNode(parent, current.Right)
	}
	if current.Right == nil {
		current.replaceNode(parent, current.Left)
	}
	//inner deletion - max size of subtrees
	replacement, replacementParent := current.Left.findMax(current)

	current.Value = replacement.Value
	current.Data = replacement.Data

	return replacement.deleteNode(replacementParent)
}

func main() {

	//create a new tree
	oneTree := &Node{Value: 8, Data: "tree"}
	oneTree.Insert(2, "dog")
	deleter := oneTree.Insert(3, "cat")
	deletee := oneTree.Insert(4, "fish")

	deleter.deleteNode(deletee)

	fmt.Printf("%v", *oneTree)

}

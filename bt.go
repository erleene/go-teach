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
	Left, Right *Node
	Value       int
	Data        string
}

func (t *Node) Insert(val int, data string) *Node {
	//we want to insert a value to the tree
	//compare new value with current node's value
	//check if the value is the same with the left
	switch {
	case val < t.Value:
		//Left node
		if t.Left == nil { // if there's no left child
			t.Left = &Node{Value: val, Data: data} //create a new one
			return t.Left
		} else { //if not empty
			t.Left.Insert(val, data) //repeat method, and return with newly inserted value
		}
	case val > t.Value:
		//Right node
		if t.Right == nil { // if there's no right child
			t.Right = &Node{Value: val, Data: data} //create a new one
			return t.Right                          //return new node with value
		} else { //if not empty
			t.Right.Insert(val, data) //repeat method, return newly inserted value
		}
	default: //if no tree
		return nil
	}
	return nil //function Insert returns itself
}

//delete from the tree a given value if it exists
//assume the node to be deleted is the right child of it's parent node
// steps also apply if the node is the left child; you only need to swap "right"
//for "left", and "large" for "small"
//parents must not be nil
//if the node has no children , remove it from its parents
//
// In the node’s left subtree, find the node with the largest value. Let’s call this node “Node B”.
// Replace the node’s value with B’s value.
// If B is a leaf node or a half-leaf node, delete it as described above for the leaf and half-leaf cases.
// If B is an inner node, recursively call the Delete method on this node.

//create helper functions
// findMax finds the maximum element in a (sub-)tree.
// Its value replaces the value of the to-be-deleted node.
// Return values: the node itself and its parent node.
func (t *Node) findMax(parent *Node) (*Node, *Node) { //func is applied to a Node type, returns parent and maximum node
	if parent.Right == nil { //if empty
		//we're at the bottom
		return parent, t //return the Max
	} //if not,
	return parent.Right.findMax(parent) //call the method on the parent node.
}

//replace node function
//replaceNode replaces the parent’s child pointer
//to n with a pointer to the replacement node. parent must not be nil.
func (child *Node) replaceNode(parent *Node, replacement *Node) error {
	//
	if child == nil { //if it doesnt exist
		return errors.New("No child node")
	}
	//
	if child == parent.Left { //if node is parent's left child
		parent.Left = replacement //replace the parent's child with replacement node
	} //child is parent's right node
	parent.Right = replacement //replace parent's right child with replacement node
	return nil
}

//
// Delete removes an element from the tree.
// It is an error to try deleting an element that does not exist.
// In order to remove an element properly, Delete needs to know the node’s parent node.
// parent must not be nil.

// Search the node to be deleted.

func (child *Node) deleteNode(parent *Node) error {
	//
	if child == nil {
		return errors.New("No child node")
	}
	switch {
	case parent.Value < child.Value: //if right node: replace the node with the parent's left child node.
		return child.Left.deleteNode(child)
	case parent.Value > child.Value: //if left node: replace the parent's left node, with the child's left.Right node.
		return child.Right.deleteNode(parent) //?
	}
	if child.Left == nil && child.Right == nil {
		child.replaceNode(parent, nil)
		return nil
	}
	//leaf node deletion
	if child.Left == nil {
		child.replaceNode(parent, child.Right)
	}
	if child.Right == nil {
		child.replaceNode(parent, child.Left)
	}
	//inner deletion - max size of subtrees
	replacement, replacementParent := child.Left.findMax(child)

	child.Value = replacement.Value
	child.Data = replacement.Data

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

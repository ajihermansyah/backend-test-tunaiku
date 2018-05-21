package main

import (
	_ "bufio"
	"fmt"
	_ "os"
)

type Node struct {
	data        int
	left, right *Node
}

func main() {

	node3 := Node{data: 3, left: nil, right: nil}
	node5 := Node{data: 5, left: nil, right: nil}
	node4 := Node{data: 4, left: nil, right: nil}
	node2 := Node{data: 2, left: &node4, right: &node5}
	node1 := Node{data: 1, left: &node2, right: &node3}
	fmt.Println("Input (Preorder)")
	node1.Preorder()
	fmt.Println("Output (Inorder)")
	node1.Inorder()
}

// Preorder (Root, Left, Right)
func (root *Node) Preorder() {
	fmt.Println(root.data)
	if root.left != nil {
		root.left.Preorder()
	}
	if root.right != nil {
		root.right.Preorder()
	}
}

// Inorder (Left, Root, Right)
func (root *Node) Inorder() {
	if root.left != nil {
		root.left.Inorder()
	}
	fmt.Println(root.data)
	if root.right != nil {
		root.right.Inorder()
	}
}

package bst

import (
	"fmt"

	"github.com/dineshtbits/data-structures-in-go/stack"
)

type BinarySearchTreeNode struct {
	data  interface{}
	left  *BinarySearchTreeNode
	right *BinarySearchTreeNode
}

func (n *BinarySearchTreeNode) Insert(data int) *BinarySearchTreeNode {
	if n == nil {
		n = &BinarySearchTreeNode{data: data}
	} else if n.data == nil {
		n.data = data
	} else if data > n.data.(int) {
		n.right = n.right.Insert(data)
	} else {
		n.left = n.left.Insert(data)
	}
	return n
}

func (n *BinarySearchTreeNode) ShowPreOrder() {
	fmt.Printf("Showing Preorder ")
	s := stack.Stack{}
	for n != nil || s.Size() > 0 {
		for n != nil {
			fmt.Printf("%v ", n.data)
			s.Push(n)
			n = n.left
		}
		n = s.Pop().(*BinarySearchTreeNode).right
	}
}

func (n *BinarySearchTreeNode) ShowInOrder() {
	fmt.Printf("Showing Inorder ")
	s := stack.Stack{}
	for n != nil || s.Size() > 0 {
		for n != nil {
			s.Push(n)
			n = n.left
		}
		n = s.Pop().(*BinarySearchTreeNode)
		fmt.Printf("%v ", n.data)
		n = n.right
	}
}

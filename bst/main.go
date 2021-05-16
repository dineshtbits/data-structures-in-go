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

func Build(arr []int) *BinarySearchTreeNode {
	if len(arr) == 0 {
		return nil
	} else if len(arr) == 1 {
		return &BinarySearchTreeNode{data: arr[0]}
	} else {
		mid := len(arr) / 2
		n := &BinarySearchTreeNode{data: arr[mid]}
		n.left = Build(arr[:mid])
		n.right = Build(arr[mid+1:])
		return n
	}
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

func (n *BinarySearchTreeNode) Delete(data int) *BinarySearchTreeNode {
	if n == nil {
		return nil
	} else if data > n.data.(int) {
		return n.right.Delete(data)
	} else if data < n.data.(int) {
		return n.right.Delete(data)
	} else {
		// equal

		if n.left != nil && n.right != nil {
			temp := n.left.Max()
			n.data = temp.data
			n.left = n.left.Delete(n.data.(int))
		}

		if n.left == nil {
			n = n.right
		}

		if n.right == nil {
			n = n.left
		}
		return n
	}
}

func (n *BinarySearchTreeNode) Max() *BinarySearchTreeNode {
	if n == nil || n.right == nil {
		return n
	} else {
		return n.right.Max()
	}
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

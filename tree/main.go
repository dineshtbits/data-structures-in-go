package tree

import (
	// "fmt"
	"github.com/dineshtbits/data-structures-in-go/queue"
)

type BinaryTreeNode struct {
	data interface{}
	left *BinaryTreeNode
	right *BinaryTreeNode
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func (b *BinaryTree) Insert(element int) {
	node := &BinaryTreeNode{data: element}
	if (b.root == nil) {
		b.root = node
	} else {
		q := queue.Queue{}
		q.Enqueue(b.root)
		for q.Size() > 0 {
			n := q.Dequeue().(*BinaryTreeNode)
			if (n.left == nil) {
				n.left = node
				break
			} else {
				q.Enqueue(n.left)
			}

			if (n.right == nil) {
				n.right = node
				break
			} else {
				q.Enqueue(n.right)
			}

		}
	}
}
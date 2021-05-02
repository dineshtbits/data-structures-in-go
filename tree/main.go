package tree

import (
	// "fmt"

	"math"

	"github.com/dineshtbits/data-structures-in-go/queue"
)

type BinaryTreeNode struct {
	data  interface{}
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

type BinaryTree struct {
	root *BinaryTreeNode
}

func (b *BinaryTree) Max() *BinaryTreeNode {
	if b.root == nil {
		panic("emty tree")
		// return nil
	} else {
		return b.root.max()
	}
}

func (n *BinaryTreeNode) max() *BinaryTreeNode {
	maxNode := &BinaryTreeNode{data: math.MinInt64}
	if n != nil {
		leftMax := n.left.max()
		rightMax := n.right.max()
		if rightMax.data.(int) > leftMax.data.(int) {
			maxNode = rightMax
		} else {
			maxNode = leftMax
		}
		if n.data.(int) > maxNode.data.(int) {
			maxNode = n
		}
	}
	return maxNode
}

func (b *BinaryTree) Insert(element int) {
	node := &BinaryTreeNode{data: element}
	if b.root == nil {
		b.root = node
	} else {
		q := queue.Queue{}
		q.Enqueue(b.root)
		for q.Size() > 0 {
			n := q.Dequeue().(*BinaryTreeNode)
			if n.left == nil {
				n.left = node
				break
			} else {
				q.Enqueue(n.left)
			}

			if n.right == nil {
				n.right = node
				break
			} else {
				q.Enqueue(n.right)
			}

		}
	}
}

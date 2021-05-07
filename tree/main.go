package tree

import (
	"fmt"
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

func (b *BinaryTree) Height() int {
	if b.root == nil {
		panic("tree is empty")
	} else {
		return b.root.height()
	}
}

func (b *BinaryTree) HeightUsingLevelOrder() int {
	if b.root == nil {
		panic("tree is empty")
	} else {
		return b.root.heightUsingLevelOrder()
	}
}

func (n *BinaryTreeNode) heightUsingLevelOrder() int {
	level := 0
	q := queue.Queue{}
	q.Enqueue(n)
	q.Enqueue(nil)
	for q.Size() > 0 {
		node := q.Dequeue()
		if node == nil {
			level++
			if q.Size() > 0 {
				q.Enqueue(nil)
			}
		} else {
			if node.(*BinaryTreeNode).left != nil {
				q.Enqueue(node.(*BinaryTreeNode).left)
			}
			if node.(*BinaryTreeNode).right != nil {
				q.Enqueue(node.(*BinaryTreeNode).right)
			}
		}
	}
	return level
}

func (n *BinaryTreeNode) height() int {
	if n == nil {
		return 0
	}
	leftHeight := n.left.height()
	rightHeight := n.right.height()
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}

func (b *BinaryTree) Find(data int) bool {
	if b.root == nil {
		panic("empty tree")
	} else {
		return b.root.find(data)
	}
}

func (b *BinaryTree) FindUsingLevelOrder(data int) bool {
	if b.root == nil {
		panic("empty tree")
	} else {
		return b.root.findUsingLevelOrder(data)
	}
}

func (n *BinaryTreeNode) findUsingLevelOrder(data int) bool {
	q := queue.Queue{}
	q.Enqueue(n)

	for q.Size() > 0 {
		node := q.Dequeue().(*BinaryTreeNode)
		if node.data == data {
			return true
		} else {
			if node.left != nil {
				q.Enqueue(node.left)
			}
			if node.right != nil {
				q.Enqueue(node.right)
			}
		}
	}
	return false
}

func (n *BinaryTreeNode) find(data int) bool {
	fmt.Printf("n: %v data: %v\n", n, data)
	if n == nil {
		return false
	} else if n.data == data {
		return true
	} else {
		found := n.left.find(data)
		if !found {
			return n.right.find(data)
		}
		return found
	}
}

func (b *BinaryTree) Max() *BinaryTreeNode {
	if b.root == nil {
		panic("emty tree")
		// return nil
	} else {
		return b.root.max()
	}
}

func (b *BinaryTree) MaxUsingLevelOrder() *BinaryTreeNode {
	if b.root == nil {
		panic("emty tree")
		// return nil
	} else {
		return b.root.maxUsingLevelOrder()
	}
}

func (n *BinaryTreeNode) maxUsingLevelOrder() *BinaryTreeNode {
	maxNode := &BinaryTreeNode{data: math.MinInt64}
	q := queue.Queue{}
	q.Enqueue(n)
	for q.Size() > 0 {
		node := q.Dequeue().(*BinaryTreeNode)
		if node.data.(int) > maxNode.data.(int) {
			maxNode = node
		}

		if node.left != nil {
			q.Enqueue(node.left)
		}

		if node.right != nil {
			q.Enqueue(node.right)
		}
	}
	return maxNode
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

package advancedtree

import (
	"fmt"
	"math"

	"github.com/dineshtbits/data-structures-in-go/advancedqueue"
	"github.com/dineshtbits/data-structures-in-go/stack"
)

type BinaryTreeNode struct {
	data  interface{}
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

func (n *BinaryTreeNode) Insert(data interface{}) {
	q := advancedqueue.Queue{}
	q.Enqueue(n)

	for q.Size() > 0 {
		node, _ := q.Dequeue()
		n := node.(*BinaryTreeNode)
		if n == nil {
			n = &BinaryTreeNode{data: data}
		} else if n.data == nil {
			n.data = data
		} else {

			if n.left == nil {
				n.left = &BinaryTreeNode{data: data}
				break
			} else {
				q.Enqueue(n.left)
			}

			if n.right == nil {
				n.right = &BinaryTreeNode{data: data}
				break
			} else {
				q.Enqueue(n.right)
			}
		}
	}
}

func (n *BinaryTreeNode) ShowLevelOrder() {
	fmt.Println()
	q := advancedqueue.Queue{}
	q.Enqueue(n)
	for q.Size() > 0 {
		node, _ := q.Dequeue()
		n := node.(*BinaryTreeNode)
		fmt.Printf("%v ", n.data)
		if n.left != nil {
			q.Enqueue(n.left)
		}
		if n.right != nil {
			q.Enqueue(n.right)
		}
	}
}

func (n *BinaryTreeNode) ShowInOrder() {
	fmt.Println()
	s := stack.Stack{}
	for s.Size() > 0 || n != nil {
		for n != nil {
			s.Push(n)
			n = n.left
		}
		n = s.Pop().(*BinaryTreeNode)
		fmt.Printf("%v ", n.data)
		n = n.right
	}
}

func (n *BinaryTreeNode) ShowPreOrder() {
	fmt.Println()
	s := stack.Stack{}
	for s.Size() > 0 || n != nil {
		for n != nil {
			fmt.Printf("%v ", n.data)
			s.Push(n)
			n = n.left
		}
		n = s.Pop().(*BinaryTreeNode)
		n = n.right
	}
}

func (n *BinaryTreeNode) ShowPostOrder() {
	fmt.Println()
	s := stack.Stack{}
	var previous *BinaryTreeNode
	for s.Size() > 0 || n != nil {
		for n != nil {
			s.Push(n)
			n = n.left
		}

		n = s.Top().(*BinaryTreeNode)
		if n.right == nil || n.right == previous {
			fmt.Printf("%v ", n.data)
			s.Pop()
			previous = n
			n = nil
		} else {
			n = n.right
		}
	}
}

func (n *BinaryTreeNode) Max() int {

	leftMax, rightMax, max := math.MinInt64, math.MinInt64, math.MinInt64

	if n.left == nil && n.right == nil {
		return n.data.(int)
	} else {
		if n.left != nil {
			leftMax = n.left.Max()
		}
		if n.right != nil {
			rightMax = n.right.Max()
		}

		if leftMax > rightMax {
			max = leftMax
		} else {
			max = rightMax
		}

		if n.data.(int) > max {
			max = n.data.(int)
		}

	}
	return max
}

func (n *BinaryTreeNode) MaxUsingLevelOrder() int {
	q := advancedqueue.Queue{}
	q.Enqueue(n)
	max := math.MinInt64
	for q.Size() > 0 {
		node, _ := q.Dequeue()
		n := node.(*BinaryTreeNode)
		if n.data.(int) > max {
			max = n.data.(int)
		}
		if n.left != nil {
			q.Enqueue(n.left)
		}
		if n.right != nil {
			q.Enqueue(n.right)
		}
	}
	return max
}

func (n *BinaryTreeNode) Size() int {
	if n == nil {
		return 0
	} else {
		return n.left.Size() + 1 + n.right.Size()
	}
}

func (n *BinaryTreeNode) Height() int {
	if n == nil {
		return 0
	} else {
		max := math.Max(float64(n.left.Height()), float64(n.right.Height()))
		return int(max) + 1
	}
}

func (n *BinaryTreeNode) HeightUsingLevelOrder() int {
	q := advancedqueue.Queue{}
	result := 0
	q.Enqueue(n)
	q.Enqueue(nil)
	for q.Size() > 0 {
		node, _ := q.Dequeue()
		if node == nil {
			result++
			if q.Size() > 0 {
				q.Enqueue(nil)
			}
		} else {
			n := node.(*BinaryTreeNode)
			if n.left != nil {
				q.Enqueue(n.left)
			}
			if n.right != nil {
				q.Enqueue(n.right)
			}
		}
	}
	return result
}

func (n *BinaryTreeNode) PrintPaths(path []int) {
	path = append(path, n.data.(int))
	if n.left == nil && n.right == nil {
		fmt.Println(path)
	} else {
		if n.left != nil {
			n.left.PrintPaths(path)
		}
		if n.right != nil {
			n.right.PrintPaths(path)
		}
	}
}

func (n *BinaryTreeNode) Mirror() {
	if n != nil {
		n.left.Mirror()
		n.right.Mirror()
		n.left, n.right = n.right, n.left
	}
}

func (n *BinaryTreeNode) VerticalSum(column int, m map[int]int) {
	if n == nil {
		return
	}
	n.left.VerticalSum(column-1, m)
	m[column] += n.data.(int)
	n.right.VerticalSum(column+1, m)
}

func (n *BinaryTreeNode) LCA(n1, n2 int) *BinaryTreeNode {

	if n == nil {
		return nil
	}
	if n.data.(int) == n1 || n.data.(int) == n2 {
		return n
	}

	left := n.left.LCA(n1, n2)
	right := n.right.LCA(n1, n2)
	if left != nil && right != nil {
		return n
	} else {
		if left != nil {
			return left
		} else {
			return right
		}
	}
}

func BuildFromPreOrder(s string) *BinaryTreeNode {
	if len(s) == 0 {
		return nil
	}

	node := &BinaryTreeNode{data: string(s[0])}
	if string(s[0]) == "L" {
		return node
	}

	node.left = BuildFromPreOrder(string(s[1:]))
	node.right = BuildFromPreOrder(string(s[2:]))
	return node
}

func (n *BinaryTreeNode) PrintAncestorsOf(target int) bool {
	if n == nil {
		return false
	}
	if n.data.(int) == target || n.left.PrintAncestorsOf(target) || n.right.PrintAncestorsOf(target) {
		fmt.Printf("%v ", n.data.(int))
		return true
	}
	return false
}

func (n *BinaryTreeNode) ZigzagTraversal() {
	fmt.Println()
	q := advancedqueue.Queue{}
	q.Enqueue(n)
	q.Enqueue(nil)
	order := 0
	s := stack.Stack{}
	for q.Size() > 0 {
		node, _ := q.Dequeue()
		if node == nil {
			// end of level reverse direction

			if order == 1 {
				// pop and print
				for s.Size() > 0 {
					fmt.Printf("%v ", s.Pop().(*BinaryTreeNode).data.(int))
				}
			}

			if q.Size() > 0 {
				q.Enqueue(nil)
			}
			order = 1 - order

		} else {
			n := node.(*BinaryTreeNode)
			if order == 0 { //leftToRight
				fmt.Printf("%v ", n.data.(int))
			} else {
				s.Push(n)
			}
			if n.left != nil {
				q.Enqueue(n.left)
			}
			if n.right != nil {
				q.Enqueue(n.right)
			}
		}
	}
}

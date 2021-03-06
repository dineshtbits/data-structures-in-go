package main

import (
	"fmt"

	"github.com/dineshtbits/data-structures-in-go/advancedqueue"
	"github.com/dineshtbits/data-structures-in-go/advancedtree"
	"github.com/dineshtbits/data-structures-in-go/bst"
	"github.com/dineshtbits/data-structures-in-go/graph"
	"github.com/dineshtbits/data-structures-in-go/heap"
	"github.com/dineshtbits/data-structures-in-go/linkedlist"
	"github.com/dineshtbits/data-structures-in-go/queue"
	"github.com/dineshtbits/data-structures-in-go/stack"
	"github.com/dineshtbits/data-structures-in-go/tree"
)

func main() {
	// linkedlistFunctions()
	// queueFunctions()
	// stackFunctions()
	// heapFunctions()
	// treeFunctions()
	// bstFunctions()
	// graphFunctions()
	// advancedQueueFunctions()
	advancedTreeFunctions()

}

func advancedTreeFunctions() {
	node := &advancedtree.BinaryTreeNode{}
	for i := 1; i <= 7; i++ {
		node.Insert(i)
	}
	node.ShowLevelOrder()
	node.ShowInOrder()
	node.ShowPreOrder()
	node.ShowPostOrder()
	fmt.Printf("\nMax is %v\n", node.Max())
	fmt.Printf("MaxUsingLevelOrder is %v\n", node.MaxUsingLevelOrder())
	fmt.Printf("Size is %v\n", node.Size())
	fmt.Printf("Height is %v\n", node.Height())
	fmt.Printf("HeightUsingLevelOrder is %v\n", node.HeightUsingLevelOrder())
	node.PrintPaths([]int{})
	// node.ShowLevelOrder()
	// node.Mirror()
	// node.ShowLevelOrder()
	m := make(map[int]int)
	node.VerticalSum(0, m)
	fmt.Printf("\nVertical sum is %v\n", m)
	fmt.Printf("\nLCA %v\n", node.LCA(4, 7))
	n := advancedtree.BuildFromPreOrder("ILILL")
	n.ShowPreOrder()
	fmt.Println()
	node.PrintAncestorsOf(4)

	node.ZigzagTraversal()
}

func graphFunctions() {
	g := graph.New()
	// fmt.Println(g)
	// g.AddNode()
	// g.AddNode()
	// g.AddNode()
	// g.AddNode()
	// g.ShowNodes()

	// fmt.Println(g.ShowEdges())
	// g.AddEdge(0, 1, 99)
	// g.AddEdge(0, 2, 99)
	// g.AddEdge(1, 2, 99)
	// g.AddEdge(2, 0, 99)
	// g.AddEdge(2, 3, 99)
	// g.AddEdge(3, 3, 99)

	// g.ShowNodes()
	// fmt.Println(g.ShowEdges())
	// g.DFS()
	// fmt.Println()
	// g.BFS()

	g = graph.New()
	g.AddNode()
	g.AddNode()
	g.AddNode()
	g.AddNode()
	g.AddNode()
	g.AddNode()
	g.AddEdge(5, 2, 99)
	g.AddEdge(5, 0, 99)
	g.AddEdge(4, 0, 99)
	g.AddEdge(4, 1, 99)
	g.AddEdge(2, 3, 99)
	g.AddEdge(3, 1, 99)
	g.TopologicalSort()

}

func treeFunctions() {
	b := &tree.BinaryTree{}
	b.Insert(5)
	b.Insert(20)
	b.Insert(12)
	b.Insert(14)
	b.Insert(15)
	fmt.Printf("b.Max() %v\n", b.Max())
	fmt.Printf("b.MaxUsingLevelOrder() %v\n", b.MaxUsingLevelOrder())
	fmt.Printf("b.Find(20) %v\n", b.Find(20))
	fmt.Printf("b.FindUsingLevelOrder(20) %v\n", b.FindUsingLevelOrder(20))
	fmt.Printf("b.Height() %v\n", b.Height())
	fmt.Printf("b.HeightUsingLevelOrder() %v\n", b.HeightUsingLevelOrder())
	inOrder := []int{4, 2, 5, 1, 6, 3}
	preOrder := []int{1, 2, 4, 5, 3, 6}
	t := tree.BuildBinaryTreeNode(preOrder, inOrder)
	fmt.Printf("Building a tree using inOrder: %v, preorder: %v, result: %v", inOrder, preOrder, t)
}

func bstFunctions() {
	b := &bst.BinarySearchTreeNode{}
	b.Insert(6)
	b.Insert(5)
	b.Insert(7)
	b.Insert(4)
	b.Insert(3)
	b.Insert(1)
	b.Insert(2)
	b.Insert(8)
	b.Insert(9)
	b.Insert(10)
	b.Insert(11)
	b.ShowPreOrder()
	fmt.Println()
	b.ShowInOrder()
	fmt.Println()
	fmt.Printf("Max: %v\n", b.Max())
	// b.Delete(6)
	// b.ShowPreOrder()
	// fmt.Println()
	// b.ShowInOrder()
	// fmt.Println()
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Printf("Building BST using sorted array: %v\n", arr)
	node := bst.Build(arr)
	node.ShowInOrder()
	fmt.Println()
}

func stackFunctions() {
	s := &stack.Stack{}
	s.Push(10)
	s.Push(9)
	s.Push(8)
	s.Push(7)
	s.Push(6)
	fmt.Printf("stack is %v\n", s)
	fmt.Printf("removed %v - %v\n", s.Pop(), s.Show())
	fmt.Printf("removed %v - %v\n", s.Pop(), s.Show())
	fmt.Printf("removed %v - %v\n", s.Pop(), s.Show())
	fmt.Printf("removed %v - %v\n", s.Pop(), s.Show())
}

func heapFunctions() {
	h := &heap.Heap{}
	fmt.Printf("initial value: %v \n", h)
	arr := []int{5, 7, 2, 9, 3, 6, 1, 4, 10}
	for _, v := range arr {
		h.Insert(v)
	}
	fmt.Printf("heapified value: %v \n", h)
	h.Extract()
	fmt.Printf("Extracted value: %v \n", h)
	h.Extract()
	fmt.Printf("Extracted value: %v \n", h)
	h.Extract()
	fmt.Printf("Extracted value: %v \n", h)
}

func linkedlistFunctions() {
	sl := linkedlist.SinglyList{}
	sl.Insert(36)
	sl.Insert(35)
	sl.Insert(34)
	sl.Show()
	sl.InsertAt(1, 37)
	sl.InsertAt(1, 38)
	sl.InsertAt(1, 39)
	sl.Show()
	sl.Append(40)
	sl.Append(41)
	sl.Show()
	sl.Delete()
	sl.Show()
	sl.Pop()
	sl.Show()
	sl.DeleteAt(1)
	sl.Show()
	n := 1
	fmt.Printf("\n%v Element from end is %v", n, sl.FindNthNodeFromEnd(n))
	fmt.Printf("\nLoop exists %v", sl.CheckIfLoopsExists())
	// sl.CreateLoop()
	// fmt.Printf("\n Loop exists %v", sl.CheckIfLoopsExists())
	sl.Reverse()
	sl.Append(34)
	sl.Show()
	fmt.Printf("\nMiddle is %v\n", sl.FindMiddle())

	// Merging two sorted lists test
	list1 := &linkedlist.SinglyList{}
	list1.Insert(3)
	list1.Insert(2)
	list1.Insert(1)
	list2 := &linkedlist.SinglyList{}
	list2.Insert(6)
	list2.Insert(5)
	list2.Insert(4)
	list1.Show()
	list2.Show()
	linkedlist.MergeSortedLists(list1, list2).Show()

	fmt.Println("\nStart: Testing Reverse in Blocks")
	list3 := &linkedlist.SinglyList{}
	list3.Insert(10)
	list3.Insert(9)
	list3.Insert(8)
	list3.Insert(7)
	list3.Insert(6)
	list3.Insert(5)
	list3.Insert(4)
	list3.Insert(3)
	list3.Insert(2)
	list3.Insert(1)
	list3.Show()
	list3.ReverseInBlocks(3)
	list3.Show()
	fmt.Println("\nEnd: Testing Reverse in Blocks")

	fmt.Println()
	fmt.Println("Doubly linked lists below")
	dl := linkedlist.DoublyList{}
	dl.Insert(4, 0)
	dl.Insert(3, 0)
	dl.Insert(2, 0)
	dl.Insert(1, 0)
	dl.Insert(5, 3)
	dl.Show()
	dl.Delete(3)
	dl.Show()
	dl.Delete(3)
	dl.Show()
}

func queueFunctions() {
	q := &queue.Queue{}
	q.Enqueue(7)
	q.Enqueue(2)
	q.Enqueue(6)
	q.Enqueue(1)
	q.Enqueue(0)
	q.Enqueue(87)
	fmt.Printf("Queue : %v\n", q)
	fmt.Printf("Dequeuing : %v --> %v\n", q.Dequeue(), q)
	fmt.Printf("Dequeuing : %v --> %v\n", q.Dequeue(), q)
	fmt.Printf("Dequeuing : %v --> %v\n", q.Dequeue(), q)
	fmt.Printf("Dequeuing : %v --> %v\n", q.Dequeue(), q)
	fmt.Printf("Dequeuing : %v --> %v\n", q.Dequeue(), q)
}

func advancedQueueFunctions() {
	q := &advancedqueue.Queue{}
	q.Enqueue(7)
	q.Enqueue(2)
	q.Enqueue(6)
	q.Enqueue(1)
	q.Enqueue(0)
	q.Enqueue(87)
	q.Show()

	element, err := q.Dequeue()
	if err != nil {
		fmt.Println(element)
	}
	q.Show()

	element, err = q.Dequeue()
	if err != nil {
		fmt.Println(element)
	}
	q.Show()
	q.ReverstFirstNElements(2)
	q.Show()
}

package queue

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(element int) {
	q.items = append(q.items, element)
}

func (q *Queue) Dequeue() int {
	dequed := q.items[0]
	q.items = q.items[1:]
	return dequed
}
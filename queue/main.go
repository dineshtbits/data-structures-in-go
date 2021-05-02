package queue

type Queue struct {
	items []interface{}
}

func (q *Queue) Enqueue(element interface{}) {
	q.items = append(q.items, element)
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Dequeue() interface{} {
	dequed := q.items[0]
	q.items = q.items[1:]
	return dequed
}
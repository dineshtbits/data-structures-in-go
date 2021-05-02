package heap

type Heap struct {
	items []int
}

func (h *Heap) Insert(element int) {
	h.items = append(h.items, element)
	h.heapifyUp()
}

func (h *Heap) Extract() {
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.heapifyDown()
}

func (h *Heap) heapifyUp() {
	for start := len(h.items) - 1; h.items[parent(start)] < h.items[start]; start = parent(start) {
		h.swap(parent(start), start)
	}
}

func (h *Heap) heapifyDown() {
	for start := 0; true; {
		largerChild := h.largerChildIndex(start)
		if largerChild <= -1 {
			return
		}
		if h.items[largerChild] > h.items[start] {
			h.swap(largerChild, start)
		}
		start = largerChild
	}
}

func (h *Heap) largerChildIndex(start int) int {
	left := 2*start + 1
	right := 2*start + 2

	if left >= len(h.items) && right >= len(h.items) {
		return -1
	} else if left < len(h.items) && right < len(h.items) {
		if h.items[left] > h.items[right] {
			return left
		} else {
			return right
		}
	} else if right >= len(h.items) {
		return left
	} else {
		return right
	}
}

func (h *Heap) swap(left, right int) {
	h.items[left], h.items[right] = h.items[right], h.items[left]
}

func parent(index int) int {
	return (index - 1) / 2
}
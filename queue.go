package queue

// Queue is the priority queue based on insertion sort
type Queue[T Item[T]] []T

// NewQueue creates an instance of the priority queue
func NewQueue[T Item[T]]() Queue[T] {
	return make(Queue[T], 0)
}

// Len returns the length of the queue
func (q *Queue[T]) Len() int {
	return len(*q)
}

// Index returns the element at the specified index, if any.
// NOTE: panics if out of bounds
func (q *Queue[T]) Index(index int) T {
	return (*q)[index]
}

// Push adds a new element to the priority queue
func (q *Queue[T]) Push(item T) {
	*q = append(*q, item)
	for i := len(*q) - 1; i > 0; i-- {
		if (*q)[i].Less((*q)[i-1]) {
			(*q)[i], (*q)[i-1] = (*q)[i-1], (*q)[i]
		} else {
			// queue is sorted, no need to continue iteration
			break
		}
	}
}

// Fix makes sure the priority queue is properly sorted
func (q *Queue[T]) Fix() {
	for i := 1; i < len(*q); i++ {
		for j := i - 1; j >= 0; j-- {
			if (*q)[j].Less((*q)[j+1]) {
				break
			}

			(*q)[j], (*q)[j+1] = (*q)[j+1], (*q)[j]
		}
	}
}

// PopFront removes the first element in the queue, if any
func (q *Queue[T]) PopFront() *T {
	if len(*q) == 0 {
		return nil
	}

	el := (*q)[0]
	*q = (*q)[1:]

	return &el
}

// PopBack removes the last element in the queue, if any
func (q *Queue[T]) PopBack() *T {
	if len(*q) == 0 {
		return nil
	}

	el := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]

	return &el
}

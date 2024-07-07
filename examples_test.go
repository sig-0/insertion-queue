package queue

import "fmt"

type number struct {
	value int
}

func (i number) Less(other number) bool {
	return i.value < other.value
}

func ExampleQueue_Push() {
	q := NewQueue[number]()

	q.Push(number{1}) // 1
	q.Push(number{2}) // 2
	q.Push(number{3}) // 3

	fmt.Println(q[0].value)
	// Output: 1
}

func ExampleQueue_PopFront() {
	q := NewQueue[number]()

	q.Push(number{1}) // 1
	q.Push(number{2}) // 2
	q.Push(number{3}) // 3

	popped := q.PopFront()
	fmt.Println(popped.value)
	// Output: 1
}

func ExampleQueue_PopBack() {
	q := NewQueue[number]()

	q.Push(number{1}) // 1
	q.Push(number{2}) // 2
	q.Push(number{3}) // 3

	popped := q.PopBack()
	fmt.Println(popped.value)
	// Output: 3
}

func ExampleQueue_Fix() {
	q := NewQueue[number]()

	q.Push(number{1}) // 1
	q.Push(number{2}) // 2
	q.Push(number{3}) // 3

	q[0] = number{4} // 1 -> 4
	q.Fix()

	fmt.Println(q[0].value)
	// Output: 2
}

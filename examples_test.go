package queue

import "fmt"

type number struct {
	value int
}

func (i number) Less(other Item) bool {
	return i.value < other.(number).value
}

func ExampleQueue_Push() {
	q := NewQueue()

	q.Push(number{1}) // 1
	q.Push(number{2}) // 2
	q.Push(number{3}) // 3

	fmt.Println(q[0].(number).value)
	// Output: 1
}

func ExampleQueue_PopFront() {
	q := NewQueue()

	q.Push(number{1}) // 1
	q.Push(number{2}) // 2
	q.Push(number{3}) // 3

	popped := q.PopFront()
	fmt.Println(popped.(number).value)
	// Output: 1
}

func ExampleQueue_PopBack() {
	q := NewQueue()

	q.Push(number{1}) // 1
	q.Push(number{2}) // 2
	q.Push(number{3}) // 3

	popped := q.PopBack()
	fmt.Println(popped.(number).value)
	// Output: 3
}

func ExampleQueue_Fix() {
	q := NewQueue()

	q.Push(number{1}) // 1
	q.Push(number{2}) // 2
	q.Push(number{3}) // 3

	q[0] = number{4} // 1 -> 4
	q.Fix()

	fmt.Println(q[0].(number).value)
	// Output: 2
}

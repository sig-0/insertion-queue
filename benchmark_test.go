package queue

import (
	"testing"
)

func benchmarkPush(b *testing.B, items []*mockItem) {
	b.Helper()

	for i := 0; i < b.N; i++ {
		q := NewQueue[*mockItem]()
		for _, item := range items {
			q.Push(item)
		}
	}
}

func BenchmarkHeap_Push10(b *testing.B) {
	items := generateRandomItems(10)

	b.ResetTimer()
	benchmarkPush(b, items)
}

func BenchmarkHeap_Push100(b *testing.B) {
	items := generateRandomItems(100)

	b.ResetTimer()
	benchmarkPush(b, items)
}

func benchmarkPop(
	b *testing.B,
	items []*mockItem,
	popCallback func(q *Queue[*mockItem]),
) {
	b.Helper()

	for i := 0; i < b.N; i++ {
		q := NewQueue[*mockItem]()

		b.StopTimer()

		for _, item := range items {
			q.Push(item)
		}

		b.StartTimer()

		for j := 0; j < len(items); j++ {
			popCallback(&q)
		}
	}
}

func BenchmarkHeap_PopFront10(b *testing.B) {
	items := generateRandomItems(10)

	b.ResetTimer()
	benchmarkPop(b, items, func(q *Queue[*mockItem]) {
		q.PopFront()
	})
}

func BenchmarkHeap_PopFront100(b *testing.B) {
	items := generateRandomItems(100)

	b.ResetTimer()
	benchmarkPop(b, items, func(q *Queue[*mockItem]) {
		q.PopFront()
	})
}

func BenchmarkHeap_PopBack10(b *testing.B) {
	items := generateRandomItems(10)

	b.ResetTimer()
	benchmarkPop(b, items, func(q *Queue[*mockItem]) {
		q.PopBack()
	})
}

func BenchmarkHeap_PopBack100(b *testing.B) {
	items := generateRandomItems(100)

	b.ResetTimer()
	benchmarkPop(b, items, func(q *Queue[*mockItem]) {
		q.PopBack()
	})
}

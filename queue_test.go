package queue

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// isSorted checks if the queue is sorted
func isSorted[T Item[T]](q Queue[T]) bool {
	for i := 0; i < len(q)-1; i++ {
		if !q[i].Less(q[i+1]) {
			return false
		}
	}

	return true
}

// generateRandomItems generates a random item set
func generateRandomItems(count int) []*mockItem {
	var (
		//nolint:gosec // Fine to use this for testing
		r     = rand.New(rand.NewSource(time.Now().UnixNano()))
		items = make([]*mockItem, count)
	)

	for i := 0; i < count; i++ {
		items[i] = &mockItem{
			value: r.Intn(100),
		}
	}

	return items
}

func TestQueue_Insert(t *testing.T) {
	t.Parallel()

	testTable := []struct {
		name      string
		ascending bool
	}{
		{
			"min-priority queue",
			true,
		},
		{
			"max-priority queue",
			false,
		},
	}

	for _, testCase := range testTable {
		testCase := testCase

		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			var (
				numItems = 100
				items    = generateRandomItems(numItems)
			)

			// Create a new queue
			q := NewQueue[*mockItem]()

			// Push items
			for _, item := range items {
				item := item

				if !testCase.ascending {
					// Prep the items if needed (min / max)
					item.lessFn = func(i *mockItem) bool {
						return item.value >= i.value
					}
				}

				q.Push(item)
			}

			assert.Len(t, q, numItems)
			assert.True(t, isSorted(q))
		})
	}
}

func TestQueue_PopFront(t *testing.T) {
	t.Parallel()

	var (
		numItems = 100
		items    = generateRandomItems(numItems)
	)

	// Create a new queue
	q := NewQueue[*mockItem]()

	// Push items
	for _, item := range items {
		q.Push(item)
	}

	assert.Len(t, q, numItems)

	// Locally sort original item queue
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Less(items[j])
	})

	// Start popping from the front
	for _, item := range items {
		popped := q.PopFront()

		assert.Equal(t, item, *popped)
	}

	assert.Len(t, q, 0)
	assert.Nil(t, q.PopFront())
}

func TestQueue_PopBack(t *testing.T) {
	t.Parallel()

	var (
		numItems = 100
		items    = generateRandomItems(numItems)
	)

	// Create a new queue
	q := NewQueue[*mockItem]()

	// Push items
	for _, item := range items {
		q.Push(item)
	}

	assert.Len(t, q, numItems)

	// Locally sort original item queue
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Less(items[j])
	})

	// Start popping from the back
	for i := len(items) - 1; i >= 0; i-- {
		popped := q.PopBack()

		assert.Equal(t, items[i], *popped)
	}

	assert.Len(t, q, 0)
	assert.Nil(t, q.PopBack())
}

func TestQueue_Fix(t *testing.T) {
	t.Parallel()

	var (
		numItems = 100
		items    = generateRandomItems(numItems)
	)

	// Create a new queue
	q := NewQueue[*mockItem]()

	// Push items
	for _, item := range items {
		q.Push(item)
	}

	assert.Len(t, q, numItems)

	newItem := &mockItem{
		value: -1, // lowest
	}

	// Manually modify the element
	q[10] = newItem

	// Resort the queue because of the manual change
	q.Fix()

	assert.Equal(t, q.Len(), numItems)
	assert.True(t, isSorted(q))
	assert.Equal(t, newItem, q.Index(0))
}

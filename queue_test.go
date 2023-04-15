package queue

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// isSorted checks if the queue is sorted
func isSorted(q Queue, ascending bool) bool {
	cmp := func(a, b *mockItem) bool {
		return a.value <= b.value
	}

	if !ascending {
		cmp = func(a, b *mockItem) bool {
			return a.value >= b.value
		}
	}

	for i := 0; i < len(q)-1; i++ {
		if !cmp(q[i].(*mockItem), q[i+1].(*mockItem)) {
			return false
		}

	}

	return true
}

// generateRandomItems generates a random item set
func generateRandomItems(count int) []*mockItem {
	var (
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
			q := NewQueue()

			// Insert items
			for _, item := range items {
				item := item

				if !testCase.ascending {
					// Prep the items if needed (min / max)
					item.lessFn = func(i Item) bool {
						other, _ := i.(*mockItem)

						return item.value >= other.value
					}
				}

				q.Insert(item)
			}

			assert.Len(t, q, numItems)
			assert.True(t, isSorted(q, testCase.ascending))
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
	q := NewQueue()

	// Insert items
	for _, item := range items {
		q.Insert(item)
	}

	assert.Len(t, q, numItems)

	// Locally sort original item queue
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Less(items[j])
	})

	// Start popping from the front
	for _, item := range items {
		popped := q.PopFront()

		assert.Equal(t, item, popped)
	}
}

func TestQueue_PopBack(t *testing.T) {
	t.Parallel()

	var (
		numItems = 100
		items    = generateRandomItems(numItems)
	)

	// Create a new queue
	q := NewQueue()

	// Insert items
	for _, item := range items {
		q.Insert(item)
	}

	assert.Len(t, q, numItems)

	// Locally sort original item queue
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Less(items[j])
	})

	// Start popping from the back
	for i := len(items) - 1; i > 0; i-- {
		popped := q.PopBack()

		assert.Equal(t, items[i], popped)
	}
}

func TestQueue_Fix(t *testing.T) {
	t.Parallel()

	var (
		numItems = 100
		items    = generateRandomItems(numItems)
	)

	// Create a new queue
	q := NewQueue()

	// Insert items
	for _, item := range items {
		q.Insert(item)
	}

	assert.Len(t, q, numItems)

	newItem := &mockItem{
		value: -1, // lowest
	}

	// Manually modify the element
	q[10] = newItem

	// Resort the queue because of the manual change
	q.Fix()

	assert.Len(t, q, numItems)
	assert.True(t, isSorted(q, true))
	assert.Equal(t, newItem, q[0])
}

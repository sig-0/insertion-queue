package queue

import "fmt"

type lessDelegate func(Item) bool

// mockItem is a mockable Item
type mockItem struct {
	value int

	lessFn lessDelegate
}

func (m *mockItem) Less(i Item) bool {
	if m.lessFn != nil {
		return m.lessFn(i)
	}

	other, ok := i.(*mockItem)
	if !ok {
		return false
	}

	return m.value < other.value
}

func (m *mockItem) String() string {
	return fmt.Sprintf("%d", m.value)
}

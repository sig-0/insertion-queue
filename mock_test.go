package queue

import "fmt"

type lessDelegate func(item *mockItem) bool

// mockItem is a mockable Item
type mockItem struct {
	lessFn lessDelegate

	value int
}

func (m *mockItem) Less(i *mockItem) bool {
	if m.lessFn != nil {
		return m.lessFn(i)
	}

	return m.value <= i.value
}

func (m *mockItem) String() string {
	return fmt.Sprintf("%d", m.value)
}

package array

import (
	"testing"
)

func TestNew(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	if a.Length() != 5 {
		t.Log("Array should have 5 items")
		t.Log("Expected", len(items), "\n Got", a.Length())
		t.Fail()
	}
}

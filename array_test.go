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

func TestSlice(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	b := a.Slice()

	if a.Length() != b.Length() {
		t.Log("Array slice without range (start or end) should have same length as original")
		t.Log("Expected", len(items), "\n Got", b.Length())
		t.Fail()
	}

	c := a.Slice(20)

	if c.Length() != 0 {
		t.Log("Array slice with out of range start index should have a length 0")
		t.Log("Expected", 0, "\n Got", c.Length())
		t.Fail()
	}

	d := a.Slice(2)

	if d.Length() != 3 {
		t.Log("Array slice created incorrect items")
		t.Log("Expected", 3, "\n Got", d.Length())
		t.Fail()
	}

	e := a.Slice(2, 3)

	if e.Length() != 1 {
		t.Log("Array slice created incorrect items")
		t.Log("Expected", 1, "\n Got", e.Length())
		t.Fail()
	}

	f := a.Slice(7, 29)

	if f.Length() != 0 {
		t.Log("Array slice should be empty")
		t.Log("Expected", 0, "\n Got", f.Length())
		t.Fail()
	}

}

func TestFind(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	value := a.Find(func(value interface{}, index int) bool {
		return value.(int) > 20
	})

	firstItem, _ := a.Get(0)

	if value != firstItem {
		t.Log("Array find should return first value that passes search criteria")
		t.Log("Expected", firstItem, "\n Got", value)
		t.Fail()
	}

	value2 := a.Find(func(value interface{}, index int) bool {
		return value.(int) > 30
	})

	if value2 != nil {
		t.Log("Array find should return nil when no value is found")
		t.Log("Expected", nil, "\n Got", value2)
		t.Fail()
	}

}

func TestFindIndex(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	index := a.FindIndex(func(value interface{}, index int) bool {
		return value.(int) > 20
	})

	if index != 0 {
		t.Log("Array find index should return first index position that passes search criteria")
		t.Log("Expected", 0, "\n Got", index)
		t.Fail()
	}

	index2 := a.FindIndex(func(value interface{}, index int) bool {
		return value.(int) > 30
	})

	if index2 != -1 {
		t.Log("Array find should return nil when no value is found")
		t.Log("Expected", -1, "\n Got", index2)
		t.Fail()
	}

}

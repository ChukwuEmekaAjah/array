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

func TestIncludes(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	included := a.Includes(2)

	if !included {
		t.Log("Array should return true for value in array")
		t.Log("Expected", true, "\n Got", included)
		t.Fail()
	}

	included2 := a.Includes(30)

	if included2 {
		t.Log("Array should return false for value not in array")
		t.Log("Expected", false, "\n Got", included2)
		t.Fail()
	}

}

func TestJoin(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	str := a.Join()
	expectedValue := "23,24,2,5,10"

	if str != expectedValue {
		t.Log("Array should return string values of its contents")
		t.Log("Expected", expectedValue, "\n Got", str)
		t.Fail()
	}

	str2 := a.Join(":")
	expectedValue2 := "23:24:2:5:10"

	if str2 != expectedValue2 {
		t.Log("Array should container joiner and return string values of its contents")
		t.Log("Expected", expectedValue2, "\n Got", str2)
		t.Fail()
	}

}

func TestString(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	str := a.String()
	expectedValue := "Array [23 24 2 5 10]"

	if str != expectedValue {
		t.Log("Array should return string values of its contents")
		t.Log("Expected", expectedValue, "\n Got", str)
		t.Fail()
	}

}

func TestLength(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	if len(items) != a.Length() {
		t.Log("Array should have same length as origin slice")
		t.Log("Expected", len(items), "\n Got", a.Length())
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	secondItem, _ := a.Get(1)

	if secondItem.(int) != items[1] {
		t.Log("Item at array index position should be same on slice")
		t.Log("Expected", items[1], "\n Got", secondItem)
		t.Fail()
	}

	_, err := a.Get(-1)

	if err == nil || err.Error() != "index out of range error" {
		t.Log("Array should return error for out of range item")
		t.Log("Expected", "index out of range error", "\n Got", err.Error())
		t.Fail()
	}

	_, err = a.Get(20)

	if err == nil || err.Error() != "index out of range error" {
		t.Log("Array should return error for out of range item")
		t.Log("Expected", "index out of range error", "\n Got", err.Error())
		t.Fail()
	}
}

func double(num interface{}, index int) interface{} {
	return num.(int) * 2
}

func TestMap(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	doubleArray := a.Map(double)

	if a.Length() != doubleArray.Length() {
		t.Log("Mapped array should have same length as original array")
		t.Log("Expected", a.Length(), "\n Got", doubleArray.Length())
		t.Fail()
	}

	i := 0
	for ; i < a.Length(); i++ {
		originalValue, _ := a.Get(i)
		mapValue, _ := doubleArray.Get(i)

		if (mapValue.(int) % originalValue.(int)) != 0 {
			t.Log("New array values should be double original array values")
			t.Log("Expected", 0, "\n Got", (mapValue.(int) % originalValue.(int)))
			t.Fail()
		}
	}
}

func isEven(value interface{}, index int) bool {
	return value.(int)%2 == 0
}
func TestFilter(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	evenArray := a.Filter(isEven)

	if evenArray.Length() != 3 {
		t.Log("Filtered array should only contain even items")
		t.Log("Expected", 3, "\n Got", evenArray.Length())
		t.Fail()
	}

}

func TestSome(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	someAreEven := a.Some(isEven)

	if !someAreEven {
		t.Log("Some items in the array are even")
		t.Log("Expected", true, "\n Got", false)
		t.Fail()
	}

	a.Fill(3, 0)
	someAreEven = a.Some(isEven)

	if someAreEven {
		t.Log("None of the items in the array is even")
		t.Log("Expected", true, "\n Got", false)
		t.Fail()
	}

}

func TestEvery(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	allAreEven := a.Every(isEven)

	if allAreEven {
		t.Log("Every item in the array is not even")
		t.Log("Expected", true, "\n Got", false)
		t.Fail()
	}

}

func TestPush(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	newValue := 7
	result := a.Push(newValue)

	if result != a.Length() {
		t.Log("Push should return the new length of the array")
		t.Log("Expected", a.Length(), "\n Got", result)
		t.Fail()
	}

	lastItem, _ := a.Get(a.Length() - 1)

	if lastItem.(int) != newValue {
		t.Log("Last item in array after push should equal newly inserted item")
		t.Log("Expected", newValue, "\n Got", lastItem)
		t.Fail()
	}

}

func TestPop(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	result := a.Pop()

	if a.Length() != len(items)-1 {
		t.Log("Pop should reduce array length by 1")
		t.Log("Expected", len(items)-1, "\n Got", a.Length())
		t.Fail()
	}

	if result != items[len(items)-1] {
		t.Log("Pop should remove and return last item in the array")
		t.Log("Expected", items[len(items)-1], "\n Got", result)
		t.Fail()
	}

}

func TestShift(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	result := a.Shift()

	if a.Length() != len(items)-1 {
		t.Log("Shift should reduce array length by 1")
		t.Log("Expected", len(items)-1, "\n Got", a.Length())
		t.Fail()
	}

	if result != items[0] {
		t.Log("Shift should return remove and return item in the array")
		t.Log("Expected", items[0], "\n Got", result)
		t.Fail()
	}

}

func TestUnShift(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	newValue := 7
	result := a.UnShift(newValue)

	if result != a.Length() {
		t.Log("UnShift should return the new length of the array")
		t.Log("Expected", a.Length(), "\n Got", result)
		t.Fail()
	}

	firstItem, _ := a.Get(0)

	if firstItem.(int) != newValue {
		t.Log("First item in array after unshift should equal newly inserted item")
		t.Log("Expected", newValue, "\n Got", firstItem)
		t.Fail()
	}

}

func TestIndexOf(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	result := a.IndexOf(5)

	if result != 3 {
		t.Log("IndexOf should return the position of specified item")
		t.Log("Expected", 3, "\n Got", result)
		t.Fail()
	}

	nonExistentIndex := a.IndexOf(0)

	if nonExistentIndex != -1 {
		t.Log("IndexOf should return -1 if item does not exist in array")
		t.Log("Expected", -1, "\n Got", nonExistentIndex)
		t.Fail()
	}

}

func TestConcat(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)
	b := New(interfaceItems)

	c := a.Concat(b)

	if c.Length() != (a.Length() + b.Length()) {
		t.Log("Concat should have length of both arrays")
		t.Log("Expected", (a.Length() + b.Length()), "\n Got", c.Length())
		t.Fail()
	}

}

func TestFill(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)
	fillValue := 2
	a.Fill(fillValue)

	for i := 0; i < a.Length(); i++ {
		value, _ := a.Get(i)

		if value.(int) != items[i] {
			t.Log("Fill should not replace any value in array")
			t.Log("Expected", items[i], "\n Got", value)
			t.Fail()
		}
	}

	start, end := 0, 3
	fillValue = 3
	a.Fill(3, start)

	for i := 0; i < a.Length(); i++ {
		value, _ := a.Get(i)

		if value.(int) != fillValue {
			t.Log("Fill should replace all values in the array with specified value")
			t.Log("Expected", fillValue, "\n Got", value)
			t.Fail()
		}
	}

	fillValue = 7
	a.Fill(fillValue, start, end)

	for i := start; i < end; i++ {
		value, _ := a.Get(i)

		if value.(int) != fillValue {
			t.Log("Fill should replace all values in the array with specified value")
			t.Log("Expected", fillValue, "\n Got", value)
			t.Fail()
		}
	}

	fillValue = 8
	start = -1
	a.Fill(fillValue, start)

	for i := 0; i < a.Length(); i++ {
		value, _ := a.Get(i)

		if value.(int) == fillValue {
			t.Log("Fill should not replace any value in array when range out of bounds")
			t.Log("Expected", items[i], "\n Got", value)
			t.Fail()
		}
	}
}

func TestValues(t *testing.T) {
	items := []int{23, 24, 2, 5, 10}
	interfaceItems := make([]interface{}, len(items))

	for i, v := range items {
		interfaceItems[i] = v
	}

	a := New(interfaceItems)

	c := a.Values()

	if len(c) != len(items) {
		t.Log("Values should return underlying slice")
		t.Log("Expected", len(items), "\n Got", len(c))
		t.Fail()
	}

	for i := 0; i < len(c); i++ {

		if c[i] != items[i] {
			t.Log("Values should return underlying slice")
			t.Log("Expected", items[i], "\n Got", c[i])
			t.Fail()
		}
	}

}

package array

import (
	"errors"
	"fmt"
)

// Array is a struct wrapper over slices to enable usage of methods indirectly on the slice
type Array struct {
	items []interface{}
}

// New creates a new array struct
func New(items []interface{}) *Array {
	newArray := Array{items: items}
	return &newArray
}

// Get returns the array element at the specified index
func (a *Array) Get(index int) (interface{}, error) {
	if index > len(a.items)-1 {
		return nil, errors.New("index out of range error")
	}

	if index < 0 {
		return nil, errors.New("index out of range error")
	}

	return a.items[index], nil

}

// Length returns the total number of elements in the array
func (a *Array) Length() int {
	return len(a.items)

}

// Map creates a new array from the return values of the map function
func (a *Array) Map(mapFunction func(interface{}, int) interface{}) *Array {
	newArray := make([]interface{}, a.Length())

	for index, value := range a.items {
		newValue := mapFunction(value, index)
		newArray[index] = newValue
	}

	return New(newArray)
}

// Filter creates a new array using the values that pass the filter function test
func (a *Array) Filter(filterFunction func(interface{}, int) bool) *Array {
	newArray := make([]interface{}, 0)

	for index, value := range a.items {

		if filterFunction(value, index) {
			newArray = append(newArray, value)
		}

	}

	return New(newArray)
}

// ForEach executes a function for each element in the array
func (a *Array) ForEach(function func(interface{}, int)) {

	for index, value := range a.items {
		function(value, index)
	}
}

// Some returns true if any of the elements in the array returns true after being tested by the check function
func (a *Array) Some(checkFunction func(interface{}, int) bool) bool {

	for index, value := range a.items {

		if checkFunction(value, index) {
			return true
		}

	}

	return false
}

// Every returns true if all the elements in the array return true after being tested by the check function
func (a *Array) Every(checkFunction func(interface{}, int) bool) bool {

	for index, value := range a.items {

		if !checkFunction(value, index) {
			return false
		}

	}

	return true
}

// Push adds a new element to the array and returns the new array length
func (a *Array) Push(item interface{}) int {

	a.items = append(a.items, item)

	return a.Length()
}

// Pop removes the last element in the array and returns it
func (a *Array) Pop() interface{} {

	value := a.items[a.Length()-1]

	a.items = a.items[0 : a.Length()-1]

	return value
}

// Shift removes the first element in the array and returns it
func (a *Array) Shift() interface{} {

	value := a.items[0]

	a.items = a.items[1:]

	return value
}

// Unshift adds an element to the beginning of the array
func (a *Array) UnShift(item interface{}) int {

	itemsList := make([]interface{}, a.Length()+1)

	itemsList[0] = item

	for i, v := range a.items {
		itemsList[i+1] = v
	}

	a.items = itemsList

	return a.Length()
}

// IndexOf returns the index of the first appearance of an element in the array
func (a *Array) IndexOf(item interface{}) int {

	for i, v := range a.items {
		if v == item {
			return i
		}
	}

	return -1
}

// Concat combines multiple arrays and returns a new array
func (a *Array) Concat(a2 *Array) *Array {

	itemsList := make([]interface{}, a.Length()+a2.Length())

	for i, v := range a.items {
		itemsList[i] = v
	}

	for i, v := range a2.items {
		itemsList[a.Length()-1+i] = v
	}

	a3 := New(itemsList)

	return a3
}

// Slice returns a new array with elements within the slice range
func (a *Array) Slice(pos ...int) *Array {

	if len(pos) == 0 {
		return New(a.items)
	}

	if len(pos) == 1 {
		if pos[0] < 0 || pos[0] > (a.Length()-1) {
			return New(make([]interface{}, 0))
		} else {
			return New(a.items[pos[0]:])
		}
	}

	if pos[0] < 0 || pos[0] > (a.Length()-1) {
		return New(make([]interface{}, 0))
	}

	if pos[1] > a.Length() {
		return New(a.items[pos[0]:])
	} else {
		return New(a.items[pos[0]:pos[1]])
	}
}

// Fill replaces the values in the array within the specified positiosn with the provided value
func (a *Array) Fill(value interface{}, pos ...int) *Array {

	if len(pos) == 0 {
		return a
	}

	if len(pos) == 1 {
		if pos[0] < 0 || pos[0] > (a.Length()-1) {
			return a
		} else {
			for counter := pos[0]; counter < a.Length(); counter++ {
				a.items[counter] = value
			}
			return a
		}
	}

	if pos[0] < 0 || pos[0] > (a.Length()-1) {
		return a
	}

	upperBound := pos[1]

	if a.Length() < upperBound {
		upperBound = a.Length()
	}

	for counter := pos[0]; counter < upperBound; counter++ {
		a.items[counter] = value
	}
	return a
}

// Find returns the first element that passes find function test
func (a *Array) Find(findFunction func(interface{}, int) bool) interface{} {

	for index, value := range a.items {

		if findFunction(value, index) {
			return value
		}

	}

	return nil
}

// Find returns the index of the first element that pass the find function test
func (a *Array) FindIndex(findFunction func(interface{}, int) bool) int {

	for index, value := range a.items {

		if findFunction(value, index) {
			return index
		}

	}

	return -1
}

// Includes returns true if the array has an element with the specified item value
func (a *Array) Includes(item interface{}) bool {

	return a.IndexOf(item) > -1
}

func (a *Array) Join(joiner ...string) string {

	returnString := ""
	joinerString := ","

	if len(joiner) > 0 {
		joinerString = joiner[0]
	}

	for _, value := range a.items {
		returnString += fmt.Sprintf("%v", value) + joinerString
	}

	if len(returnString) == 0 {
		return returnString
	}

	return returnString[0 : len(returnString)-1]
}

// String returns the string representation of the array
func (a *Array) String() string {

	return fmt.Sprintf("Array %v", a.items)
}

// Values returns the underlying array slice
func (a *Array) Values() []interface{} {

	return a.items
}

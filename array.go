package array

import (
	"errors"
)

type Array struct {
	items []interface{}
}

func (a *Array) Get(index int) (interface{}, error) {
	if index > len(a.items)-1 {
		return nil, errors.New("index out of range error")
	}

	if index < 0 {
		return nil, errors.New("index out of range error")
	}

	return a.items[index], nil

}

func (a *Array) Length() int {
	return len(a.items)

}

func (a *Array) Map(mapFunction func(interface{}, int) interface{}) []interface{} {
	newArray := make([]interface{}, a.Length())

	for index, value := range a.items {
		newValue := mapFunction(value, index)
		newArray[index] = newValue
	}

	return newArray
}

func New(items []interface{}) *Array {
	newArray := Array{items: items}
	return &newArray
}

func (a *Array) Filter(filterFunction func(interface{}, int) bool) []interface{} {
	newArray := make([]interface{}, a.Length())
	counter := 0

	for index, value := range a.items {

		if filterFunction(value, index) {
			newArray[counter] = value
			counter += 1
		}

	}

	return newArray
}

func (a *Array) ForEach(function func(interface{}, int)) {

	for index, value := range a.items {
		function(value, index)
	}
}

func (a *Array) Some(checkFunction func(interface{}, int) bool) bool {

	for index, value := range a.items {

		if checkFunction(value, index) {
			return true
		}

	}

	return false
}

func (a *Array) Every(checkFunction func(interface{}, int) bool) bool {

	for index, value := range a.items {

		if !checkFunction(value, index) {
			return false
		}

	}

	return true
}

func (a *Array) Push(item interface{}) int {

	a.items = append(a.items, item)

	return a.Length()
}

func (a *Array) Pop() interface{} {

	value := a.items[a.Length()-1]

	a.items = a.items[0 : a.Length()-1]

	return value
}

func (a *Array) Shift() interface{} {

	value := a.items[0]

	a.items = a.items[1:]

	return value
}

func (a *Array) UnShift(item interface{}) int {

	itemsList := make([]interface{}, a.Length()+1)

	itemsList[0] = item

	for i, v := range a.items {
		itemsList[i+1] = v
	}

	a.items = itemsList

	return a.Length()
}

func (a *Array) IndexOf(item interface{}) int {

	for i, v := range a.items {
		if v == item {
			return i
		}
	}

	return -1
}

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

func (a *Array) Fill(value interface{}, pos ...int) *Array {

	if len(pos) == 0 {
		return New(a.items)
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

func (a *Array) Find(findFunction func(interface{}, int) bool) interface{} {

	for index, value := range a.items {

		if findFunction(value, index) {
			return value
		}

	}

	return nil
}

func (a *Array) FindIndex(findFunction func(interface{}, int) bool) int {

	for index, value := range a.items {

		if findFunction(value, index) {
			return index
		}

	}

	return -1
}

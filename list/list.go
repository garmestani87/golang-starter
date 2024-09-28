package main

import (
	"fmt"
)

type ArrayList[T any] struct {
	items []T
	equal func(T, T) bool
}

func main() {
	list := ArrayList[int]{}
	list.add(1)
	list.add(2)
	list.add(3)
	list.add(4)

	list.equal = func(i1, i2 int) bool { return i1 == i2 }

	list.insertAt(2, 5)
	list.removeAt(2)
	list.remove(2)

	for _, item := range list.items {
		fmt.Println(item)
	}
}

func (arr *ArrayList[T]) add(item T) {
	arr.items = append(arr.items, item)
}

func (arr *ArrayList[T]) insertAt(idx int, item T) {
	arr.items = append(arr.items, item)
	copy(arr.items[idx+1:], arr.items[idx:])
	arr.items[idx] = item
}

func (arr *ArrayList[T]) removeAt(idx int) {
	arr.items = append(arr.items[:idx], arr.items[idx+1:]...)
}

func (arr *ArrayList[T]) remove(item T) {
	idx := arr.find(item)
	if idx != -1 {
		arr.removeAt(idx)
	}
}

func (arr *ArrayList[T]) find(item T) int {
	for idx, value := range arr.items {
		if arr.equal(item, value) {
			return idx
		}
	}
	return -1
}

package main

import "fmt"

type Node[T any] struct {
	value       T
	left, right *Node[T]
}

type BTree[T any] struct {
	root *Node[T]
}

func main() {
	tree := BTree[int]{}
	less := func(a, b int) bool {
		return a < b
	}
	equal := func(a, b int) bool {
		return a == b
	}
	values := []int{5, 3, 8, 1, 4, 7, 9}
	for _, v := range values {
		tree.Insert(v, less)
	}
	searchValue := 7
	node := tree.Search(searchValue, less, equal)
	if node != nil {
		fmt.Printf("Value %d found in the tree.\n", searchValue)
	} else {
		fmt.Printf("Value %d not found in the tree.\n", searchValue)
	}
}

func (bt *BTree[T]) Insert(val T, less func(T, T) bool) {
	newNode := &Node[T]{value: val}
	if bt.root == nil {
		bt.root = newNode
	} else {
		insertNode(bt.root, newNode, less)
	}
}

func insertNode[T any](root *Node[T], newNode *Node[T], less func(T, T) bool) {
	if less(newNode.value, root.value) {
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode, less)
		}
	} else {
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode, less)
		}
	}
}

func (bt *BTree[T]) Search(val T, less func(T, T) bool, equal func(T, T) bool) *Node[T] {
	return searchNode(bt.root, val, less, equal)
}

func searchNode[T any](root *Node[T], val T, less func(T, T) bool, equal func(T, T) bool) *Node[T] {
	if root == nil || equal(root.value, val) {
		return root
	}
	if less(val, root.value) {
		return searchNode(root.left, val, less, equal)
	}
	return searchNode(root.right, val, less, equal)
}

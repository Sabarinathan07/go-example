package main

import "fmt"

// FindIndex returns the index of value in the slice, or -1 if not found.
func FindIndex[Slice ~[]Elem, Elem comparable](slice Slice, value Elem) int {
    for i := range slice {
        if value == slice[i] {
            return i
        }
    }
    return -1
}

// LinkedList is a generic singly linked list.
type LinkedList[T any] struct {
    head, tail *Node[T]
}

// Node represents an element in the linked list.
type Node[T any] struct {
    next *Node[T]
    val  T
}

// Add appends a value to the end of the linked list.
func (list *LinkedList[T]) Add(value T) {
    if list.tail == nil {
        list.head = &Node[T]{val: value}
        list.tail = list.head
    } else {
        list.tail.next = &Node[T]{val: value}
        list.tail = list.tail.next
    }
}

// Elements returns all values in the linked list as a slice.
func (list *LinkedList[T]) Elements() []T {
    var values []T
    for node := list.head; node != nil; node = node.next {
        values = append(values, node.val)
    }
    return values
}

func main() {
    var fruits = []string{"food", "bar", "zoo"}

    fmt.Println("index of zoo:", FindIndex(fruits, "zoo"))

    _ = FindIndex[[]string, string](fruits, "zoo")

    numbers := LinkedList[int]{}
    numbers.Add(10)
    numbers.Add(13)
    numbers.Add(23)
    fmt.Println("list:", numbers.Elements())
}
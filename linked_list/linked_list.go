package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *LinkedList) String() string {
	if l.Head == nil {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Len: %d - LinkedList: ", l.Length))
	sb.WriteString("[")

	current := l.Head
	for current != nil {
		sb.WriteString(strconv.Itoa(current.Value))
		current = current.Next
		if current != nil {
			sb.WriteString(" -> ")
		}
	}

	sb.WriteString("]")
	return sb.String()
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("Val: %d", n.Value)
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
	}
}

func NewLinkedList(value int) *LinkedList {
	node := NewNode(value)
	return &LinkedList{
		Head:   node,
		Tail:   node,
		Length: 1,
	}
}

func (l *LinkedList) Push(value int) bool {
	newNode := NewNode(value)
	if l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
		l.Length = 1
		return true
	}

	l.Length++
	l.Tail.Next = newNode
	l.Tail = newNode
	return true
}

func (l *LinkedList) Pop() *int {
	if l.Length == 0 {
		return nil
	}

	if l.Length == 1 {
		value := l.Head.Value
		l.Head = nil
		l.Tail = nil
		l.Length = 0
		return &value
	}

	pre := l.Head
	post := l.Head.Next
	for post.Next != nil {
		pre = post
		post = post.Next
	}

	l.Tail = pre
	l.Tail.Next = nil
	l.Length--
	return &post.Value
}

func (l *LinkedList) Shift() *Node {
	if l.Head == nil || l.Length == 0 {
		return nil
	}

	if l.Head == l.Tail {
		value := l.Head
		l.Head = nil
		l.Tail = nil
		l.Length = 0
		return value
	}

	value := l.Head
	l.Head = l.Head.Next
	l.Length--
	return value
}

func (l *LinkedList) Unshift(value int) bool {
	newNode := NewNode(value)
	newNode.Next = l.Head
	l.Head = newNode
	l.Length++
	return true
}

func (l *LinkedList) Get(index int) *Node {
	if index > l.Length-1 || index < 0 {
		return nil
	}
	curr := l.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	return curr
}

func (l *LinkedList) Set(index int, value int) {
	node := l.Get(index)
	if node != nil {
		node.Value = value
	}
}

func (l *LinkedList) Insert(index int, value int) bool {
	if index == 0 {
		return l.Unshift(value)
	}
	if index == l.Length {
		return l.Push(value)
	}

	temp := l.Get(index - 1)
	if temp != nil {
		newNode := NewNode(value)
		newNode.Next = temp.Next
		temp.Next = newNode
		l.Length++
		return true
	}
	return false
}

func (l *LinkedList) Remove(index int) *int {
	if index == 0 {
		if res := l.Shift(); res != nil {
			return &res.Value
		}
		return nil
	}

	if index == l.Length-1 {
		return l.Pop()
	}

	if index < 0 || index >= l.Length {
		return nil
	}

	temp := l.Get(l.Length - 1)
	value := temp.Next.Value
	temp.Next = temp.Next.Next
	l.Length--
	return &value
}

func (l *LinkedList) Reverse() {
	var prev *Node
	current := l.Head
	var next *Node

	l.Tail = current

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}

func main() {
	ll := NewLinkedList(0)
	// ll := &LinkedList{}
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)
	fmt.Println(ll)
	ll.Reverse()
	fmt.Println(ll)
	fmt.Println(ll.Head.Next)
	fmt.Println(ll.Tail.Next)
}

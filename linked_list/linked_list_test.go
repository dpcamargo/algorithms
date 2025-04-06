package main

import "testing"

func TestNewNode(t *testing.T) {
	tests := []struct {
		value, expected int
	}{
		{2, 2},
		{3, 3},
		{0, 0},
	}
	for _, test := range tests {
		result := NewNode(test.value)
		expected := Node{
			Value: test.expected}

		if result.Value != expected.Value {
			t.Errorf("NewNode(%d) = %v; want %v", test.value, result, expected)
		}
		if result.Next != nil {
			t.Errorf("NewNode(%d) = %v; want %v", test.value, result, expected)
		}
	}
}

func TestNewLinkedList(t *testing.T) {
	tests := []struct {
		value, expected, length int
	}{
		{1, 1, 1},
		{3, 3, 1},
		{0, 0, 1},
	}
	for _, test := range tests {
		result := NewLinkedList(test.value)
		node := NewNode(test.expected)
		expected := LinkedList{
			Head:   node,
			Tail:   node,
			Length: test.length,
		}

		if result.Head.Value != expected.Head.Value {
			t.Errorf("NewLinkedList(%d) = %v; want %v", test.value, result.Head, expected.Head)
		}
		if result.Tail.Value != expected.Tail.Value {
			t.Errorf("NewLinkedList(%d) = %v; want %v", test.value, result.Tail, expected.Tail)
		}
		if result.Length != expected.Length {
			t.Errorf("NewLinkedList(%d) = %v; want %v", test.length, result.Length, expected.Length)
		}
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		initial  *LinkedList
		value    int
		expected int
		length   int
	}{
		{NewLinkedList(12), 1, 1, 2},
		{&LinkedList{}, 1, 1, 1},
	}
	for _, test := range tests {
		test.initial.Push(test.value)
		if test.initial.Tail.Value != test.value {
			t.Errorf("Push(%d).Tail.Value == %d; want %d", test.value, test.initial.Tail.Value, test.expected)
		}
		if test.initial.Length != test.length {
			t.Errorf("Push(%d).Length == %d; want %d", test.value, test.initial.Length, test.length)
		}
	}
}

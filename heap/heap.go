package main

import "fmt"

type Heap []int

func leftChild(index int) int {
	return 2*index + 1
}

func rightChild(index int) int {
	return 2*index + 2
}

func parent(index int) int {
	return (index - 1) / 2
}

func (h Heap) swap(index1, index2 int) {
	h[index1], h[index2] = h[index2], h[index1]
}

func (h *Heap) insert(value int) {
	*h = append(*h, value)
	current := len(*h) - 1
	for current > 0 {
		p := parent(current)
		if (*h)[current] < (*h)[p] {
			break
		}
		h.swap(current, p)
		current = p
	}
}

func newHeap(value int) *Heap {
	return &Heap{value}
}

func (h *Heap) remove() *int {
	if len(*h) == 0 {
		return nil
	}

	if len(*h) == 1 {
		return h.remove()
	}

	maxValue := (*h)[0]
	lastIndex := len(*h) - 1
	(*h)[0] = h.remove(lastIndex)

}

func sink()

func main() {
	h := newHeap(99)
	h.insert(72)
	h.insert(61)
	h.insert(58)
	fmt.Println(h)
	h.insert(100)
	fmt.Println(h)
	h.insert(75)
	fmt.Println(h)
}

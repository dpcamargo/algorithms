package main

import "fmt"

func bubbleSort(unsorted []int) {
	for i := len(unsorted) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if unsorted[j] > unsorted[j+1] {
				unsorted[j], unsorted[j+1] = unsorted[j+1], unsorted[j]
			}
		}
	}
}

func main() {
	arr := []int{5, 3, 7, 1, 2, 8, 6}
	bubbleSort(arr)
	fmt.Println(arr)
}

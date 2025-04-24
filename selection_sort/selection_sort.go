package main

import "fmt"

func selectionSort(unsorted []int) {
	for i := range unsorted {
		min := i
		for j, v := range unsorted[i:] {
			if v < unsorted[min] {
				min = j + i
			}
		}
		unsorted[min], unsorted[i] = unsorted[i], unsorted[min]
	}
}

func main() {
	arr := []int{5, 3, 7, 9, 1, 4, 2, 8, 6, 0}
	selectionSort(arr)
	fmt.Println(arr)
}

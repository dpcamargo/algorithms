package main

import "fmt"

func insertionSort(unsorted []int) {
	var temp int
	var j int
	for i := 1; i < len(unsorted); i++ {
		temp = unsorted[i]
		j = i - 1
		for j >= 0 && unsorted[j] > temp {
			unsorted[j+1] = unsorted[j]
			j--
		}
		unsorted[j+1] = temp
	}
}

func main() {
	arr := []int{5, 3, 7, 9, 1, 4, 2, 8, 6, 0}
	insertionSort(arr)
	fmt.Println(arr)
}

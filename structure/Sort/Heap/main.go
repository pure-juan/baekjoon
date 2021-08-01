package main

import "fmt"

func main() {
	arr := []int{3, 7, 8, 1, 5, 9, 10, 2, 4}
	sort(&arr, 0, len(arr))

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func sort(arr *[]int, parent int, last int) {
	fmt.Printf("arr: %v\n", arr)
	left := parent*2 + 1
	right := parent*2 + 2
	largest := parent

	if left < last && (*arr)[largest] < (*arr)[left] {
		largest = left
	}

	if right < last && (*arr)[largest] < (*arr)[right] {
		largest = right
	}

	if parent != largest {
		temp := (*arr)[parent]
		(*arr)[parent] = (*arr)[largest]
		(*arr)[largest] = temp
		sort(arr, largest, last)
	}
}

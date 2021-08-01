package main

import "fmt"

func main() {
	arr := []int{3, 7, 8, 1, 5, 9, 10, 2, 4}

	for i := len(arr)/2 - 1; i >= 0; i-- {
		fmt.Printf("%d\n", i)
		Heapify(&arr, i, len(arr))
	}
	Sort(&arr)

	for _, v := range arr {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}

func Heapify(arr *[]int, start int, length int) {
	fmt.Printf("arr: %v, start :%d\n", arr, start)
	// index의 왼쪽 오른쪽 자식 index
	left, right := start*2+1, start*2+2
	// 기본적으로 왼쪽 자식이 큰걸로
	largest := start
	if left < length {
		fmt.Printf("left: %d\n", (*arr)[left])
	}
	if right < length {
		fmt.Printf("right: %d\n", (*arr)[right])
	}
	// 오른쪽 자식이 더 크면 그걸로 선택
	if left < length && (*arr)[left] > (*arr)[largest] {
		largest = left
	}
	if right <= length && (*arr)[right] > (*arr)[largest] {
		largest = right
	}

	fmt.Printf("largest: %d\n", largest)

	if largest != start {
		temp := (*arr)[start]
		(*arr)[start] = (*arr)[largest]
		(*arr)[largest] = temp
		Heapify(arr, largest, length)
	}
}

func Sort(arr *[]int) {
	for length := len(*arr) - 1; length >= 0; length-- {
		temp := (*arr)[0]
		(*arr)[0] = (*arr)[length]
		(*arr)[length] = temp
		Heapify(arr, 0, length)
	}
}

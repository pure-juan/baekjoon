package main

import (
	"bufio"
	"fmt"
	"os"
)

type Heap struct {
	arr  []int
	size int
}

func (heap *Heap) Insert(value int) {
	heap.size = heap.size + 1
	index := heap.size

	for index != 1 && value > heap.arr[index/2] {
		heap.arr[index] = heap.arr[index/2]
		index /= 2
	}

	heap.arr[index] = value
}

func (heap *Heap) RemoveRoot() int {
	parent, child := 1, 2
	item, temp := heap.arr[1], heap.arr[heap.size]

	heap.size -= 1

	for child <= heap.size {
		if child < heap.size && heap.arr[child] < heap.arr[child+1] {
			child += 1
		}
		if temp >= heap.arr[child] {
			break
		} else {
			heap.arr[parent] = heap.arr[child]
			parent = child
			child = child * 2
		}
	}

	heap.arr[parent] = temp
	heap.arr = heap.arr[:len(heap.arr)-1]
	return item
}

func (heap *Heap) Print(writer *bufio.Writer) {
	fmt.Fprintf(writer, "%v\n", heap.arr)
}

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	arr := []int{3, 7, 8, 1, 5, 9, 10, 2, 4}
	length := len(arr)
	heap := Heap{
		arr:  make([]int, length+1),
		size: 0,
	}

	for _, v := range arr {
		heap.Insert(v)
	}
	heap.Print(writer)

	heap.RemoveRoot()
	heap.Print(writer)
}

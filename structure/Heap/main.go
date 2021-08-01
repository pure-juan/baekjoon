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

func (heap *Heap) Insert(value int) { // 추가
	heap.size = heap.size + 1 // 값 추가여서 사이즈 하나 늘림
	index := heap.size        //힙의 가장 마지막부터 시작함

	// index가 루트가 아닐때까지 (index 1번이 Root임, 0번 비워둠) && 추가하려는 값이 부모 값보다 크면 자리 변경
	for index != 1 && value > heap.arr[index/2] {
		heap.arr[index] = heap.arr[index/2] // 부모를 아래로 내리기
		index /= 2                          // 부모 노드로 이동
	}

	heap.arr[index] = value // 자리 변경
}

func (heap *Heap) RemoveRoot() int { // Root 삭제
	parent, child := 1, 2
	// item = Root Node를 할당, temp = 마지막 노드를 할당
	item, temp := heap.arr[1], heap.arr[heap.size]

	// size 하나 줄이기
	heap.size -= 1

	for child <= heap.size {
		// parent의 왼쪽 오른쪽 자식 중 더 큰 자식을 찾음
		// 1. 자식이 두개인지 확인, 2. 왼쪽자식과 오른쪽 자식을 비교
		if child < heap.size && heap.arr[child] < heap.arr[child+1] {
			// 오른쪽 자식이 더 크면 오른쪽 자식으로 변경
			child += 1
		}
		// 위 둘중 더 큰 자식 보다 마지막 노드가 크면 종료
		if temp >= heap.arr[child] {
			break
		} else {
			// 위 둘중 더 큰 자식 보다 마지막 노드가 작으면 자리 변경
			heap.arr[parent] = heap.arr[child]
			parent = child
			// 한단계 아래로 이동
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

	for i := 0; i < length; i++ {
		fmt.Fprintf(writer, "%d ", heap.RemoveRoot())
	}
	fmt.Fprintln(writer)
}

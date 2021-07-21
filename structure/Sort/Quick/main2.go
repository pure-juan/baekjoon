package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"sync"
)

func main2() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	writer := bufio.NewWriter(os.Stdout)
	g := &sync.WaitGroup{}
	defer writer.Flush()

	// arr := []int{3, 7, 8, 1, 5, 9, 6, 10, 2, 4}
	arr := []int{26, 32, 76, 48, 38, 96, 12, 68, 50, 94, 7, 22, 34, 10, 58, 65, 87, 97, 2, 6}
	g.Add(1)
	go sort2(&arr, 0, len(arr)-1, g)
	g.Wait()

	for _, v := range arr {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer, "")
}

func sort2(arr *[]int, start int, end int, g *sync.WaitGroup) {
	defer g.Done()
	if start == end {
		return
	}

	pivot := start
	left := start
	right := end

	for left < right {
		for {
			if (*arr)[left] > (*arr)[pivot] {
				break
			}

			if left >= len(*arr)-1 {
				break
			}

			left += 1
		}

		for {
			if (*arr)[right] < (*arr)[pivot] {
				break
			}

			if right < left {
				break
			}

			right -= 1
		}

		var temp int
		if left < right {
			temp = (*arr)[left]
			(*arr)[left] = (*arr)[right]
			(*arr)[right] = temp
		} else {
			temp = (*arr)[right]
			(*arr)[right] = (*arr)[pivot]
			(*arr)[pivot] = temp
			break
		}
	}

	g.Add(1)
	go sort2(arr, start, right, g)
	g.Add(1)
	go sort2(arr, left, end, g)
}

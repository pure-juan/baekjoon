package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// arr := []int{3, 7, 8, 1, 5, 9, 6, 10, 2, 4}
	arr := []int{26, 32, 76, 48, 38, 96, 12, 68, 50, 94, 7, 22, 34, 10, 58, 65, 87, 97, 2, 6}
	sort(&arr, 0, len(arr)-1)

	for _, v := range arr {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer, "")
}

func sort(arr *[]int, start int, end int) {
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

	sort(arr, start, right)
	sort(arr, left, end)
}

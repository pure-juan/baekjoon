package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	arr := []int{3, 7, 8, 1, 5, 9, 6, 10, 2, 4}
	sort(&arr, 0, len(arr)-1)

	for _, v := range arr {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer, "")
}

func sort(arr *[]int, start int, end int) {
	fmt.Printf("%d %d\n", start, end)
	if start+1 == end {
		return
	}

	pivot := start
	left := start + 1
	right := end

	for left < right {
		for {
			if (*arr)[left] > (*arr)[pivot] {
				break
			}

			if left >= len(*arr)-1 {
				break
			}

			left++
		}

		for {
			if (*arr)[right] < (*arr)[pivot] {
				break
			}

			if right <= start {
				break
			}

			right--
		}

		var temp int
		if left < right {
			temp = (*arr)[left]
			(*arr)[left] = (*arr)[right]
			(*arr)[right] = temp
		} else {
			temp = (*arr)[right]
			(*arr)[pivot] = (*arr)[right]
			(*arr)[right] = temp
		}

		fmt.Printf("%v left: %d, right: %d\n", arr, left, right)
	}

	sort(arr, start, right)
	sort(arr, right+1, end)
}

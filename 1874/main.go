package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var size int
	fmt.Fscanln(scanner, &size)
	var original = []int{}
	for i := 0; i < size; i++ {
		var input int
		fmt.Fscanln(scanner, &input)
		original = append(original, input)
	}

	var stack = []int{}
	var result = []int{}
	var index = 1
	for len(original) > 0 {
		if len(stack) == 0 || original[0] != stack[len(stack)-1] {
			if Contain(&stack, original[0]) {
				fmt.Fprintln(writer, "NO")
				return
			}
			stack = append(stack, index)
			result = append(result, 1)
			index++
		} else if original[0] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			original = original[1:]
			result = append(result, -1)
		}
	}

	for _, v := range result {
		if v == 1 {
			fmt.Fprintln(writer, "+")
		} else {
			fmt.Fprintln(writer, "-")
		}
	}
}

func Contain(arr *[]int, value int) bool {
	for _, v := range *arr {
		if v == value {
			return true
		}
	}

	return false
}

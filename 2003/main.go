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

	var numOfNumber, target int
	fmt.Fscanf(scanner, "%d %d\n", &numOfNumber, &target)

	var list []int = make([]int, numOfNumber)
	for i := 0; i < numOfNumber; i++ {
		var number int
		fmt.Fscanf(scanner, "%d", &number)
		list[i] = number
	}

	index := 0
	result := 0
	for index < numOfNumber {
		temp := list[index]
		if temp == target {
			result++
			index++
			continue
		}
		for i := index + 1; i < numOfNumber; i++ {
			temp += list[i]
			if temp == target {
				result++
				break
			} else if temp > target {
				break
			}
		}
		index++
	}

	fmt.Fprintf(writer, "%d\n", result)
}

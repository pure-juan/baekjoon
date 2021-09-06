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

	list := make([]int, numOfNumber)
	for i := 0; i < numOfNumber; i++ {
		var temp int
		fmt.Fscanf(scanner, "%d ", &temp)
		list[i] = temp
	}

	result := 0
Loop:
	for k := 1; k <= numOfNumber; k++ {
		for i := 0; i < numOfNumber; i++ {
			accumulate := list[i]
			if accumulate >= target {
				result = k
				break Loop
			}
			for j := 1; j < k; j++ {
				if i+j >= numOfNumber {
					break
				}
				accumulate += list[i+j]
				if accumulate >= target {
					result = k
					break Loop
				}
			}
		}
	}

	fmt.Fprintln(writer, result)
}

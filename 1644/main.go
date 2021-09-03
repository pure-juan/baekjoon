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

	var target int
	fmt.Fscanln(scanner, &target)

	list := MakePrime(target)

	result := 0
	for index, value := range list {
		var accumulate = value
		if accumulate == target {
			result++
			break
		}
		for _, j := range list[index+1:] {
			accumulate += j
			if accumulate == target {
				result++
				break
			} else if accumulate > target {
				break
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", result)
}

func MakePrime(target int) []int {
	list := make([]int, target+1)
	for i := 2; i <= target; i++ {
		list[i] = i
	}
	for i := 2; i <= target; i++ {
		if list[i] == 0 {
			continue
		}
		for j := i * 2; j <= target; j += i {
			list[j] = 0
		}
	}

	newArr := make([]int, 0)
	for _, value := range list {
		if value != 0 {
			newArr = append(newArr, value)
		}
	}

	return newArr
}

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
	var originalSequence = make([]int, size)
	for i := 0; i < size; i++ {
		var input int
		fmt.Fscanln(scanner, &input)
		originalSequence[i] = input
	}

	var sequence []int = []int{}
	var result []string = []string{}
	var current = 0
	var fail = false
	for len(originalSequence) > 0 {
		if len(sequence) > 0 && sequence[len(sequence)-1] == originalSequence[0] {
			if !Pop(&originalSequence, &sequence, &current, &result) {
				fail = true
				break
			}
		} else {
			if !Push(&originalSequence, &sequence, &current, &result) {
				fail = true
				break
			}
		}
	}

	if fail {
		fmt.Fprintf(writer, "NO")
	} else {
		for _, v := range result {
			fmt.Fprintln(writer, v)
		}
	}
}

func Push(originalSequence *[]int, sequence *[]int, current *int, result *[]string) bool {
	(*sequence) = append((*sequence), (*current)+1)
	(*result) = append((*result), "+")

	(*current)++
	return true
}

func Pop(originalSequence *[]int, sequence *[]int, currenet *int, result *[]string) bool {
	if len(*sequence) > 0 {
		(*sequence) = (*sequence)[:len(*sequence)-1]
	}
	(*originalSequence) = (*originalSequence)[1:]

	(*result) = append((*result), "-")

	return true
}

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
	var deck []int = make([]int, size)
	for i := 0; i < size; i++ {
		deck[i] = i + 1
	}

	flag := true
	for len(deck) != 1 {
		if flag {
			deck = deck[1:]
		} else {
			firstItem := deck[0]
			deck = deck[1:]
			deck = append(deck, firstItem)
		}

		flag = !flag
	}

	fmt.Fprintf(writer, "%d\n", deck[0])
}

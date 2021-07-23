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

	for repeat := size - 1; repeat >= 0; repeat-- {
		for i := 0; i < repeat; i++ {
			fmt.Fprint(writer, " ")
		}
		for i := 0; i < size-repeat; i++ {
			fmt.Fprint(writer, "*")
		}
		fmt.Fprintln(writer)
	}
}

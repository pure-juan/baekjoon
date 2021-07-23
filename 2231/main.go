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

	var number int
	var result int
	fmt.Fscanln(scanner, &number)

	for i := 0; i < number; i++ {
		result = Check(i)
		if result == number {
			fmt.Fprintf(writer, "%d\n", i)
			return
		}
	}

	fmt.Fprintf(writer, "0\n")
}

func Check(num int) int {
	result := num
	for {
		val := num % 10
		num /= 10
		result += val
		if num <= 0 {
			break
		}
	}

	return result
}

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

	var repeat int
	fmt.Fscanln(scanner, &repeat)
	var result []string = make([]string, repeat)
	for i := 0; i < repeat; i++ {
		var str string
		fmt.Fscanln(scanner, &str)
		var length = len(str)
		if length%2 != 0 {
			result[i] = "NO"
			continue
		}
		result[i] = Check(str)
	}

	for _, v := range result {
		fmt.Fprintln(writer, v)
	}
}

func Check(str string) string {
	var stack []string = []string{}
	for i := 0; i < len(str); i++ {
		var cur = string(str[i])
		if cur == "(" {
			stack = append(stack, "")
		} else {
			if len(stack) <= 0 {
				return "NO"
			} else {
				stack[len(stack)-1] = ""
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) > 0 {
		return "NO"
	}

	return "YES"
}

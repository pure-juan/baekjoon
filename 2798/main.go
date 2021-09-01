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
	var numOfCards int
	var number int
	fmt.Fscanln(scanner, &numOfCards, &number)
	listOfCard := make([]int, numOfCards)
	for i := 0; i < numOfCards; i++ {
		fmt.Fscanf(scanner, "%d", &listOfCard[i])
	}
	var length = len(listOfCard)
	var best int = 0
Loop:
	for i := 0; i < length; i++ {
		first := listOfCard[i]
		for j := i + 1; j < length; j++ {
			second := listOfCard[j]
			for k := j + 1; k < length; k++ {
				third := listOfCard[k]
				if first+second+third == number {
					best = first + second + third
					break Loop
				} else if first+second+third <= number && first+second+third > best {
					best = first + second + third
				}
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", best)
}

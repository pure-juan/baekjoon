package main

import (
	"bufio"
	"fmt"
	"os"
)

type People struct {
	weight int
	height int
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var numOfPeople int
	fmt.Fscanln(scanner, &numOfPeople)

	peoples := make([]People, numOfPeople)
	for i := 0; i < numOfPeople; i++ {
		var weight, height int
		fmt.Fscanf(scanner, "%d %d\n", &weight, &height)
		peoples[i] = People{
			weight: weight,
			height: height,
		}
	}

	rank := make([]int, numOfPeople)

	for i := 0; i < numOfPeople; i++ {
		numOfBigger := 0
		for j := 0; j < numOfPeople; j++ {
			if i == j {
				continue
			}

			if peoples[i].weight < peoples[j].weight && peoples[i].height < peoples[j].height {
				numOfBigger++
			}
		}
		rank[i] = numOfBigger + 1
	}

	for _, value := range rank {
		fmt.Fprintf(writer, "%d ", value)
	}
	fmt.Fprintln(writer)
}

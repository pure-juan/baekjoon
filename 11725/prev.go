// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	scanner := bufio.NewReader(os.Stdin)
// 	writer := bufio.NewWriter(os.Stdout)
// 	defer writer.Flush()
// 	var numOfNodes int
// 	fmt.Fscanln(scanner, &numOfNodes)
// 	var tree [][]bool = make([][]bool, numOfNodes)
// 	for i := 0; i < numOfNodes; i++ {
// 		tree[i] = make([]bool, 0)
// 	}
// 	for i := 0; i < numOfNodes-1; i++ {
// 		var firstNode, secondNode int
// 		fmt.Fscanln(scanner, &firstNode, &secondNode)
// 		tree[firstNode-1][secondNode-1] = true
// 		tree[secondNode-1][firstNode-1] = true
// 	}

// 	var visited = make([]bool, numOfNodes)
// 	var result = make([]int, numOfNodes)

// 	Search(tree, 0, &visited, &result)
// 	for _, v := range result[1:] {
// 		fmt.Fprintf(writer, "%d\n", v+1)
// 	}
// }

// func Search(tree [][]bool, current int, visited *[]bool, result *[]int) {
// 	(*visited)[current] = true

// 	for i, v := range tree[current] {
// 		if v && !(*visited)[i] {
// 			(*result)[i] = current
// 			Search(tree, i, visited, result)
// 		}
// 	}
// }

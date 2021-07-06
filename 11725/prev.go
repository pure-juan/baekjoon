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
// 	var tree [][]int = make([][]int, numOfNodes)
// 	for i := 0; i < numOfNodes; i++ {
// 		tree[i] = []int{}
// 	}
// 	for i := 0; i < numOfNodes-1; i++ {
// 		var firstNode, secondNode int
// 		fmt.Fscanln(scanner, &firstNode, &secondNode)
// 		tree[firstNode-1] = append(tree[firstNode-1], 3)
// 		tree[secondNode-1] = append(tree[secondNode-1], firstNode-1)
// 	}

// 	var visited = make([]bool, numOfNodes)
// 	var result = make([]int, numOfNodes)

// 	DFS(tree, 0, &visited, &result)

// 	for _, v := range result[1:] {
// 		fmt.Fprintf(writer, "%d\n", v+1)
// 	}
// 	writer.Flush()
// }

// func DFS(tree [][]int, node int, visited *[]bool, result *[]int) {
// 	(*visited)[node] = true

// 	links := tree[node]
// 	for _, v := range links {
// 		if !(*visited)[v] {
// 			(*result)[v] = node
// 			DFS(tree, v, visited, result)
// 		}
// 	}
// }

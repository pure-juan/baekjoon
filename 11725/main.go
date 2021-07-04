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
	var numOfNodes int
	fmt.Fscanln(scanner, &numOfNodes)
	var tree [][]int = make([][]int, numOfNodes)
	for i := 0; i < numOfNodes; i++ {
		tree[i] = make([]int, 0)
	}
	for i := 0; i < numOfNodes-1; i++ {
		var firstNode, secondNode int
		fmt.Fscanln(scanner, &firstNode, &secondNode)
		tree[firstNode-1] = append(tree[firstNode-1], secondNode-1)
		tree[secondNode-1] = append(tree[secondNode-1], firstNode-1)
	}

	var queue = []int{0}
	var visited = make([]bool, numOfNodes)
	var result = make([]int, numOfNodes)

	for len(queue) > 0 {
		BFS(tree, &queue, &visited, &result)
	}

	for _, v := range result[1:] {
		fmt.Fprintf(writer, "%d\n", v+1)
	}
	writer.Flush()
}

func BFS(tree [][]int, queue *[]int, visited *[]bool, result *[]int) {
	node := (*queue)[0]
	(*queue) = (*queue)[1:]
	(*visited)[node] = true

	for _, val := range tree[node] {
		if !(*visited)[val] {
			(*result)[val] = node
			(*queue) = append((*queue), val)
		}
	}
}

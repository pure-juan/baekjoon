package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Graph struct {
	links []int
}

func main() {
	var numOfEdges, numOfLinks, startFrom int
	fmt.Scanln(&numOfEdges, &numOfLinks, &startFrom)
	var DFSVisited = make([]bool, numOfEdges)
	var BFSVisited = make([]bool, numOfEdges)
	var graphs = make([]Graph, numOfEdges)
	var DFSArrived []string = []string{}
	var BFSArrived []string = []string{}

	for i := 0; i < numOfLinks; i++ {
		var firstEdge, secondEdge int
		fmt.Scanln(&firstEdge, &secondEdge)
		graphs[firstEdge-1].links = append(graphs[firstEdge-1].links, secondEdge)
		graphs[secondEdge-1].links = append(graphs[secondEdge-1].links, firstEdge)
	}

	for _, v := range graphs {
		sort.Ints(v.links)
	}

	DFS(graphs, startFrom, &DFSVisited, &DFSArrived)

	var queue []int = []int{}
	queue = append(queue, startFrom)
	BFS(graphs, startFrom, &BFSVisited, &BFSArrived, &queue)

	fmt.Println(strings.Join(DFSArrived, " "))
	fmt.Println(strings.Join(BFSArrived, " "))
}

func DFS(graphs []Graph, edge int, visited *[]bool, arrived *[]string) {
	(*visited)[edge-1] = true
	(*arrived) = append(*arrived, strconv.Itoa(edge))

	var links = graphs[edge-1].links
	for _, v := range links {
		if !(*visited)[v-1] {
			DFS(graphs, v, visited, arrived)
		}
	}
}

func BFS(graphs []Graph, edge int, visited *[]bool, arrived *[]string, queue *[]int) {
	(*visited)[edge-1] = true
	(*arrived) = append(*arrived, strconv.Itoa(edge))
	(*queue) = (*queue)[1:len(*queue)]

	var links = graphs[edge-1].links
	for _, v := range links {
		if !(*visited)[v-1] && !Contain(queue, v) {
			(*queue) = append(*queue, v)
		}
	}

	if len(*queue) > 0 {
		BFS(graphs, (*queue)[0], visited, arrived, queue)
	}
}

func Contain(arr *[]int, value int) bool {
	for _, v := range *arr {
		if v == value {
			return true
		}
	}

	return false
}

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
	var numOfEdges, numOfLinks, startFrom int        // 변수 선언
	fmt.Scanln(&numOfEdges, &numOfLinks, &startFrom) // 읽어 오기
	var DFSVisited = make([]bool, numOfEdges)        // DFS용 방문 확인
	var BFSVisited = make([]bool, numOfEdges)        // BFS용 방문 확인
	var graphs = make([]Graph, numOfEdges)           // 그래프
	var DFSArrived []string = []string{}             // DFS용 방문 기록
	var BFSArrived []string = []string{}             // BFS용 방문 기록

	for i := 0; i < numOfLinks; i++ { // Link 수 만큼 데이터 읽어오기
		var firstEdge, secondEdge int
		fmt.Scanln(&firstEdge, &secondEdge)
		graphs[firstEdge-1].links = append(graphs[firstEdge-1].links, secondEdge)  // 양방향
		graphs[secondEdge-1].links = append(graphs[secondEdge-1].links, firstEdge) // 양방향
	}

	for _, v := range graphs { // 그래프 정렬
		sort.Ints(v.links)
	}

	DFS(graphs, startFrom, &DFSVisited, &DFSArrived) // DFS

	var queue []int = []int{}                                // 큐 생성
	queue = append(queue, startFrom)                         // 큐에 기본 시작점 추가
	BFS(graphs, startFrom, &BFSVisited, &BFSArrived, &queue) // BFS

	fmt.Println(strings.Join(DFSArrived, " ")) // 출력
	fmt.Println(strings.Join(BFSArrived, " ")) // 출력
}

func DFS(graphs []Graph, edge int, visited *[]bool, arrived *[]string) { // DFS
	(*visited)[edge-1] = true
	(*arrived) = append(*arrived, strconv.Itoa(edge))

	var links = graphs[edge-1].links
	for _, v := range links {
		if !(*visited)[v-1] {
			DFS(graphs, v, visited, arrived)
		}
	}
}

func BFS(graphs []Graph, edge int, visited *[]bool, arrived *[]string, queue *[]int) { // BFS
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

func Contain(arr *[]int, value int) bool { // BFS용 유틸 함수
	for _, v := range *arr {
		if v == value {
			return true
		}
	}

	return false
}

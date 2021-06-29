/*
 * https://www.acmicpc.net/problem/2606
 * 바이러스
 */

package main

import "fmt"

type Graph struct {
	links []int
}

func main() {
	var numOfEdges, numOfLinks int
	fmt.Scanln(&numOfEdges) // 컴퓨터의 갯수
	fmt.Scanln(&numOfLinks) // 컴퓨터 쌍의 수
	var graphs []Graph = make([]Graph, numOfEdges)
	for i := 0; i < numOfLinks; i++ { // Link의 갯수 많큼 읽어오기
		var firstEdge, secondEdge int       // 양방향이라 두곳 다 넣어주기
		fmt.Scanln(&firstEdge, &secondEdge) // Link
		graphs[firstEdge-1].links = append(graphs[firstEdge-1].links, secondEdge)
		graphs[secondEdge-1].links = append(graphs[secondEdge-1].links, firstEdge)
	}

	var visited []bool = make([]bool, numOfEdges) // 방문 체크
	var count int = 0                             // 갯수
	DFSMain(graphs, &visited, &count)             // DFS
	// BFSMain(graphs, &visited, &count) // BFS
	fmt.Println(count - 1) // 출력
}

func DFSMain(graphs []Graph, visited *[]bool, count *int) { // DFS
	DFS(graphs, 1, visited, count)
}

func DFS(graphs []Graph, edge int, visited *[]bool, count *int) { // DFS
	(*visited)[edge-1] = true
	*count += 1

	var links = graphs[edge-1].links
	for _, v := range links {
		if !(*visited)[v-1] {
			DFS(graphs, v, visited, count)
		}
	}
}

func BFSMain(graphs []Graph, visited *[]bool, count *int) { // BFS
	var queue []int = []int{1}          // 큐
	BFS(graphs, visited, &queue, count) // BFS
}

func BFS(graphs []Graph, visited *[]bool, queue *[]int, count *int) { // BFS
	edge := (*queue)[0] - 1            // 큐에서 하나 빼오기
	(*visited)[edge] = true            // 방문 처리
	(*queue) = (*queue)[1:len(*queue)] // 큐에서 맨 앞 데이터 제거
	*count += 1                        // 카운팅

	var links = graphs[edge].links // 연결되어 있는 컴퓨터 가져오기
	for _, v := range links {
		if !(*visited)[v-1] && !Contain(queue, v) { // 방문 했는지와 큐에 들어가 있는지
			(*queue) = append(*queue, v) // 없으면 큐에 추가
		}
	}

	if len(*queue) > 0 { // 큐가 비어있지 않다면 계속 돌아감
		BFS(graphs, visited, queue, count)
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

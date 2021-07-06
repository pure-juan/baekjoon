package main

import (
	"fmt"
	"strings"
)

type Node struct {
	left  *Node
	right *Node
	data  string
}

func main() {
	var numOfNodes int
	var tree = Node{
		left:  nil,
		right: nil,
		data:  "A",
	}
	fmt.Scanln(&numOfNodes)
	for i := 0; i < numOfNodes; i++ {
		var target, left, right string
		fmt.Scanln(&target, &left, &right)
		node := Search(&tree, target)
		if left != "." {
			node.left = &Node{
				left:  nil,
				right: nil,
				data:  left,
			}
		}
		if right != "." {
			node.right = &Node{
				left:  nil,
				right: nil,
				data:  right,
			}
		}
	}

	preOrderVisited := []string{}
	inOrderVisited := []string{}
	postOrderVisited := []string{}
	PreOrderSearch(&tree, &preOrderVisited)
	InOrderSearch(&tree, &inOrderVisited)
	PostOrderSearch(&tree, &postOrderVisited)

	fmt.Printf("%s\n%s\n%s\n", strings.Join(preOrderVisited, ""), strings.Join(inOrderVisited, ""), strings.Join(postOrderVisited, ""))
}

func Search(node *Node, data string) *Node {
	if node != nil {
		if node.data == data {
			return node
		} else {
			nd := Search(node.left, data)
			if nd != nil {
				return nd
			}
			return Search(node.right, data)
		}
	}
	return nil
}

func PreOrderSearch(node *Node, visited *[]string) {
	if node != nil {
		*visited = append(*visited, node.data)
		PreOrderSearch(node.left, visited)
		PreOrderSearch(node.right, visited)
	}
}

func InOrderSearch(node *Node, visited *[]string) {
	if node != nil {
		InOrderSearch(node.left, visited)
		*visited = append(*visited, node.data)
		InOrderSearch(node.right, visited)
	}
}

func PostOrderSearch(node *Node, visited *[]string) {
	if node != nil {
		PostOrderSearch(node.left, visited)
		PostOrderSearch(node.right, visited)
		*visited = append(*visited, node.data)
	}
}

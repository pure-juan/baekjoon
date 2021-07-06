package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	data string
	next *Node
}

var head Node
var tail Node

func main() {
	for i := 0; i < 10; i++ {
		Append(strconv.Itoa(i))
	}
	PrintAll()
	RemoveFirst()
	Explain()
	PrintAll()
	Add(1, "Hi")
	Explain()
	PrintAll()
	RemoveLast()
	Explain()
	PrintAll()
	Remove(1)
	Explain()
	PrintAll()
}

func Get(index int) string {
	cur := head.next
	i := 0
	for cur != nil {
		if i == index {
			return cur.data
		}
	}

	return ""
}

func Add(index int, data string) {
	var newData = Node{
		data: data,
		next: nil,
	}
	cur := &head
	i := 0

	for cur != nil {
		if i == index {
			temp := cur.next
			cur.next = &newData
			newData.next = temp
			if tail.next == temp {
				tail.next = &newData
			}

			break
		}

		cur = cur.next
		i++
	}
}

func Append(str string) {
	var next = Node{
		data: str,
		next: nil,
	}

	if head.next == nil {
		head.next = &next
		tail.next = &next
	} else {
		tail.next.next = &next
		tail.next = &next
	}
}

func RemoveFirst() bool {
	next := head.next

	if next == nil {
		return false
	} else {
		if next.next != nil {
			head.next = next.next
		} else {
			head.next = nil
		}
	}

	return true
}

func RemoveLast() bool {
	cur := head.next

	for cur != nil {
		if cur.next == tail.next {
			cur.next = nil
			tail.next = cur
			break
		}
		cur = cur.next
	}

	return true
}

func Remove(index int) bool {
	if index == 0 {
		return RemoveFirst()
	}

	cur := head.next
	var pre *Node = &head
	var i = 0
	for cur != nil {
		if i == index {
			if cur.next == tail.next {
				tail.next = pre
			}
			pre.next = cur.next

			return true
		}

		pre = cur
		cur = cur.next
		i++
	}

	return true
}

func PrintAll() {
	var cur = head.next

	for cur != nil && cur.next != nil {
		fmt.Printf("%s\n", cur.data)
		cur = cur.next
	}

	if cur != nil {
		fmt.Printf("%s\n", cur.data)
	}

	fmt.Println("---")
}

func Explain() {
	fmt.Printf("First %v\n", head.next)
	fmt.Printf("Last %v\n", tail.next)
}

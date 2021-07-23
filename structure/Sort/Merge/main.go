package main

import (
	"bufio"
	"fmt"
	"os"
)

var temp []int

func main() {
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	// arr := []int{3, 7, 8, 1, 5, 9, 10, 2, 4}
	arr := []int{26, 32, 76, 48, 38, 96, 12, 68, 50, 94, 7, 22, 34, 10, 58, 65, 87, 97, 2, 6}

	splitedArr := Split(&arr) // 2차원 배열로 자르기
	fmt.Fprintf(writer, "split: %v\n", splitedArr)
	result := Merge(&splitedArr) // 시작!
	for _, v := range result {
		fmt.Fprintf(writer, "%d ", v)
	}
	fmt.Fprintln(writer)
}

func Split(arr *[]int) [][]int {
	newArr := [][]int{}

	for _, v := range *arr {
		newArr = append(newArr, []int{v})
	}

	return newArr
}

func Merge(arr *[][]int) []int {
	if len(*arr) == 1 { // 하나 남으면 끝난거
		return (*arr)[0]
	}

	originalLength := len(*arr)

	for i := 0; i < originalLength; i += 2 { // 2개씩 잡아서 정렬
		temp = []int{}
		if i+1 == originalLength { // 하나 남았으면 그대로 넘기기
			temp = append(temp, (*arr)[i]...)
			*arr = append(*arr, temp)
			break
		}
		first := &(*arr)[i]    // 첫번째
		second := &(*arr)[i+1] // 두번째
		for {
			if len(*first) == 0 && len(*second) == 0 { // 배열에 더이상 없을때
				break
			}
			if len(*first) == 0 && len(*second) != 0 { // 첫번째 배열이 비었고 두번째 배열만 있을때
				temp = append(temp, (*second)[0])
				*second = (*second)[1:]
			} else if len(*second) == 0 && len(*first) != 0 { // 위에 반대
				temp = append(temp, (*first)[0])
				*first = (*first)[1:]
			} else if (*first)[0] < (*second)[0] { // 첫번째 배열에 첫번째 아이템이 두번째 배열에 첫번째 아이템보다 작을때
				temp = append(temp, (*first)[0])
				*first = (*first)[1:]
			} else { // 위에 반대
				temp = append(temp, (*second)[0])
				*second = (*second)[1:]
			}
		}
		*arr = append(*arr, temp)
	}

	*arr = (*arr)[originalLength:]

	fmt.Printf("arr: %v\n", arr)

	return Merge(arr)
}

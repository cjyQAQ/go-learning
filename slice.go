package main

import "fmt"

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func printSliceWhithIndex(x []int, start int, end int) {
	var s []int
	fmt.Println("------------")
	printSlice(s)
	if end == -1 {
		s = x[start:]
	} else if start == -1 {
		s = x[:end]
	} else {
		s = x[start:end]
	}
	printSlice(s)
}

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(s)
	printSliceWhithIndex(s, 0, 3)
	printSliceWhithIndex(s, -1, 3)
	printSliceWhithIndex(s, 5, -1)
	s = append(s, 11, 12, 13)
	printSlice(s)
}

package main

import "fmt"

func main() {
	arr := make([]int, 0, 10)
	arr = append(arr, 1)
	arr = append(arr, 2)
	arr = append(arr, 3)
	fmt.Println(moveToEnd(arr, 1))
	arr = append(arr, 4)
	arr = append(arr, 5)
	fmt.Println(arr)
	fmt.Println(moveToEnd(arr, 4))
	fmt.Println(moveToEnd(arr, 4))
}

func moveToEnd(arr []int, v int) []int {
	// Write your code here
	n := 0
	for _, i := range arr {
		if i != v {
			arr[n] = i
			n++
		}
	}
	arr = arr[:n]
	return append(arr, v)
}

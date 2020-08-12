package main

import "fmt"

func main() {
	arr := []int{2, 1, 2, 3, 9, 7, 12, 0}
	SearchSort(arr)
	fmt.Println(arr)
}

func SearchSort(arr []int) {

	for i := len(arr) - 1; i >= 0; i-- {
		loca := 0
		for j := 0; j <= i; j++ {
			if arr[j] > arr[loca] {
				loca = j
			}
		}
		arr[i], arr[loca] = arr[loca], arr[i]
	}
}

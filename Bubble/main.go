package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 4, 5, 6}
	Bubble(arr)
	fmt.Println(arr)
}

func Bubble(arr []int) {
	l := len(arr)

	for i := l - 1; i > 0; i-- {
		swapped := false
		for j := 0; j < i; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if swapped {
			break
		}
	}
}

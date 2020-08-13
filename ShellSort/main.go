package main

import "fmt"

func main() {
	arr := []int{13, 14, 94, 33, 82, 25, 59, 94, 65, 23, 45, 27, 73, 25, 39, 10}
	ShellSort(arr, len(arr)/2)
	fmt.Println(arr)
}

func ShellSort(arr []int, l int) {

	if l < 1 {
		return
	}

	for k := 0; k < l; k++ {
		for j := 1; l*j+k < len(arr); j++ {
			for i := j; i > 0; i-- {
				pre := l*(i-1) + k
				next := l*i + k
				if arr[pre] < arr[next] {
					arr[pre], arr[next] = arr[next], arr[pre]
				}
			}
		}
	}

	ShellSort(arr, l/2)
}

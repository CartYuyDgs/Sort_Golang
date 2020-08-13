package main

import "fmt"

func main() {
	var arr = []int{-9, 2, 1, 4, 22, 0, -3}
	fmt.Println(arr)
	QuickSort(0, len(arr)-1, arr)
	fmt.Println(arr)
}

func QuickSort(left, right int, arr []int) {
	l := left
	r := right

	pivot := arr[(left+right)/2]
	fmt.Println(pivot)
	for l < r {
		for arr[l] < pivot {
			l++
		}

		for arr[r] > pivot {
			r--
		}

		if l >= r {
			break
		}

		arr[r], arr[l] = arr[l], arr[r]

		if arr[l] == pivot {
			r--
		}
		if arr[r] == pivot {
			l++
		}

	}
	if l == r {
		l++
		r--
	}

	if left < r {
		QuickSort(left, r, arr)
	}
	if right > l {
		QuickSort(l, right, arr)
	}
}

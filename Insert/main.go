package main

import "fmt"

func main() {
	var a = []int{1, 3, 3, 2, 34, 23, 22, 11, 9, 100}
	InsertSort(a)
}

func InsertSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		inserVal := arr[i]
		insertIndex := i - 1

		for insertIndex >= 0 && arr[insertIndex] < inserVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
			fmt.Println(arr)
		}

		if insertIndex+1 != i {
			arr[insertIndex+1] = inserVal
		}
		fmt.Println(arr)
	}

	//inserVal := arr[1]
	//insertIndex := 1 - 1
	//
	//for insertIndex >=0 && arr[insertIndex] < inserVal {
	//	arr[insertIndex+1] = arr[insertIndex]
	//	insertIndex --
	//}
	//
	//if insertIndex + 1 != 1 {
	//	arr[insertIndex+1] = inserVal
	//}
	//fmt.Println(arr)

	//inserVal = arr[2]
	//insertIndex = 2 - 1
	//
	//for insertIndex >=0 && arr[insertIndex] < inserVal {
	//	arr[insertIndex+1] = arr[insertIndex]
	//	insertIndex --
	//}
	//
	//if insertIndex + 1 != 2 {
	//	arr[insertIndex+1] = inserVal
	//}
	//fmt.Println(arr)
	//
	//inserVal = arr[3]
	//insertIndex = 3 - 1
	//
	//for insertIndex >=0 && arr[insertIndex] < inserVal {
	//	arr[insertIndex+1] = arr[insertIndex]
	//	insertIndex --
	//}
	//
	//if insertIndex + 1 != 3 {
	//	arr[insertIndex+1] = inserVal
	//}
	//fmt.Println(arr)
}

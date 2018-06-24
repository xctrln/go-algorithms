package main

import "fmt"

import "github.com/0xAX/go-algorithms"

func merge(arr1 []int, arr2 []int) []int {
	i := 0
	j := 0
	n := 0
	nlen := len(arr1) + len(arr2)
	narr := make([]int, 0, nlen)
	for j+i < nlen {

		if i < len(arr1) && j < len(arr2) {
			if arr1[i] > arr2[j] {
				narr = append(narr, arr2[j])
				j++
			} else {
				narr = append(narr, arr1[i])
				i++
			}
		} else if i >= len(arr1) {
			return append(narr, arr2[j:]...)
		} else if j >= len(arr2) {
			return append(narr, arr1[i:]...)
		}
		n++
	}
	return narr

}

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	} else {
		arr1 := arr[:len(arr)/2]
		arr2 := arr[len(arr)/2:]
		left := MergeSort(arr1)
		right := MergeSort(arr2)
		return merge(left, right)
	}
}

func main() {
	arr := utils.RandArray(10)
	narr := MergeSort(arr)
	fmt.Println(narr)
}

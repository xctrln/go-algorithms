package main

import "fmt"


func merge(arr1 []int, arr2 []int) []int  {
	i := 0
	j := 0
	n := 0
	nlen := len(arr1) + len(arr2)
	narr := make([]int, nlen) 
	for ; j + i < nlen;  {
		if arr1[i] > arr2[j] {
			narr[n] = arr2[j]
			j++
		}else{
			narr[n] = arr1[i]
			i++
		}
		n++
	}
	return narr
}

func MergeSort(arr []int) []int {
	if len(arr) == 1{
		return arr
	}else{
		arr1 := arr[0:len(arr)/2]
		arr2 := arr[len(arr)/2 + 1 : len(arr) - 1]
		left := MergeSort(arr1)
		right := MergeSort(arr2)
		return merge(left, right)
	}
}

func main(){
	arr := utils.RandArray(10)
	narr := MergeSort(arr)
	fmt.Println(narr)
}
package main

import "fmt"

import "github.com/0xAX/go-algrithms"

func merge(arr1, arr2)  {
	i := 0
	j := 0
	narr := []
	nlen = len(arr1) + len(arr2)
	for ; j + i < nlen;  {
		if arr1[i] > arr2[j] {
			narr.append(arr2[j])
			j++
		}else{
			narr.append(arr1[i])
			i++
		}
	}
	return narr
}

func MergeSort(arr){
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
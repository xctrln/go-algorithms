package main

import "fmt"

import "github.com/0xAX/go-algorithms"

func main(){
	arr := util.RandArray(10)
	alen = len(arr)

	for i := 0; i < alen; i++ {
		for j := 0; j < alen - 1 - i; j++ {
			if(arr[j] > arr[j+1]){
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)
}
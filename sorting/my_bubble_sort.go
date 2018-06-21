package main

import "fmt"

import "github.com/0xAX/go-algorithms"

func main(){
	arr := utils.RandArray(10)
	alen := len(arr)
	flag := 0

	for i := 0; i < alen; i++ {
		for j := 0; j < alen - 1 - i; j++ {
			if(arr[j] > arr[j+1]){
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = 1
			}
		}
		if flag == 0{
			break
		}
	}

	fmt.Println(arr)
}
package main

import "fmt"
import "github.com/0xAX/go-algorithms"

func main()  {
	arr := utils.RandArray(10)
	alen := len(arr)
	for i := 0; i < alen; i++{
		for j := i; j < alen; j++{
			if arr[i] > arr[j]{
				arr[i], arr[j] = arr[j], arr[i]
			}	
		}
	}
	fmt.Println(arr)
}
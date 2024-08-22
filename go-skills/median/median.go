package main

import "fmt"

func Median(data []int) int {
	data = bubblesort(data)
	n := len(data)
	median := 0

	if n%2 == 0{
		m1 := data[n/2]
		m2 := data[n/2] - 1
		median = (m1+m2)/2
	} else {
		median = data[n/2]
	}
	return median
}

func bubblesort(data []int) []int{
	for i:=0; i < len(data) -1 ; i++{
		for j := 0; j < len(data)-1-i;j++{
			if data[j] > data[j+1]{
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

func main(){
	data := []int{7,8,4,2,9,1,3, 6}
	fmt.Println(bubblesort(data))
	fmt.Println(Median(data))
}
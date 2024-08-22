package main

import (
	"fmt"
)

func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
}


func RevConcatAlternate(slice1,slice2 []int) []int {
	result := []int{}

	i,j := len(slice1)-1, len(slice2)-2

	for i >= 0 || j >= 0{
		if i == j{
			if i >= 0{
				result = append(result, slice1[i])
				i--
			}
			if j >= 0{
				result = append(result, slice2[j])
				j--
			}
		}else if i > j{
			result = append(result, slice1[i])
			i--
		} else {
			result = append(result, slice2[j])
			j--
		}
	}
	return result
}

package main

import (
	"fmt"
)

func main() {
	//ways to declare arrays
	var a1 []int = []int{1,2,3,4,5}
	// a2 := []int{1,2,3,4,5}
	// a3 := make([]int, 5)
	 arr := [8]int{02,13,42,53,64,75,56,37}

	 var s []int = arr[5:]
	fmt.Println(arr)
	fmt.Println("slice", a1)
	s = append(s, 56)
	fmt.Println("slice", s)
	arr2D := [2][3]int{{12,33, 34}, {1,2,3}}
	fmt.Println(arr2D)

	for i, element := range arr {
		fmt.Println(i, element)
	}

}
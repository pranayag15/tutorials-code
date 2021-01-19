package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//one way of declaring variable
	var x int = 5
	var z = 554
	// another way of declaring variable
	y := 7
	sum := y + x
	fmt.Println(sum, z)

	//if statemenents
	if y < 6 {
		fmt.Println("Less than 6")
	} else if y>6 && y<8 {
		fmt.Println("value is 7")
	}

	//* arrays
	a := []int{1,2,3,4,5}
	a = append(a, 90)
	fmt.Println(a)

	//map
	vertices := make(map[string]int)
	vertices["pranay"] = 21
	vertices["agarwal"] = 45
	fmt.Println(vertices)
	delete(vertices, "agarwal")
	fmt.Println(vertices)
	val, ifPresent := vertices["pranay"] // val waill contain thw value if thr key is present and ifPresent will set to true else vice-versa
	// val, ifPresent := vertices["chopra"]
	fmt.Println(val, ifPresent)
}
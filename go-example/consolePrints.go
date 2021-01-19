package main

import (
	"fmt"
)

func main() {
	var a = 56.0
	a = a+ 2.5
	b := "pranay"
	fmt.Println(a)
	fmt.Printf("Hello %T %v %v %v huhuhhuu \n", 10, 10, 52, b) // %T is for datatype
	var x = fmt.Sprintf("Hello %T %v %v", 10, 10, 52)  //store the sting format in variabke
	fmt.Println(x)
}
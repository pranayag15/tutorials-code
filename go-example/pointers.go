package main

import (
	"fmt"
)

func changeValue(str *string){
	*str = "Changed by function1!"
}


func changeValue2(str string) {
	str = "Chnaged by 2"
} 

func main() {
	x := 7
	y := &x
	fmt.Println(x, y, *y)
	*y  = 8
	fmt.Println(x, y, *y)
	var str string = "original value!"
	fmt.Println(str)
	changeValue(&str)
	fmt.Println(str)
	changeValue2(str)
	fmt.Println(str)
}
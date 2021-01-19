package main

import (
	"fmt"
)

func add(x, y int, z string) (int, string) { //x and y both are int and z string
	return x+y, "returning a string"
}

func anotherReturningWay(x, y int) (z1 int, z2 int) {
	defer fmt.Println("hello")
	z1 = x + y
	z2 = x - y
	fmt.Println("before return")
	return
}

func test() {
	fmt.Println("test")
}

func main() {
	add, st := add(5,6,"z is string")
	fmt.Println(add, st)
	z1, z2 := anotherReturningWay(56,7)
	fmt.Println(z1, z2)
	test := func (x int) int{
		return x*-1
	}(8)
	fmt.Println("calling test", test)
}
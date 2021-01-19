package main

import (
	"fmt"
)

type Point struct{
	x int
	y int
}

type Circle struct{
	radius float64
	// center *Point
	*Point
}

func changeValueStruct(p *Point){
	p.x = 100
}

type Student struct{
	name string
	grades []int
	age int
}

func (s *Student) setAge(age int){
	s.age = age
}

func main(){
	p1 := &Point{y:3}
	fmt.Println(p1)
	changeValueStruct(p1)
	fmt.Println(p1)

	c1 := Circle{4.56, p1}
	// fmt.Println(c1, c1.center, c1.center.x, "circle")
	fmt.Println(c1, c1.x, "circle")
	

	s1 := Student{"Pranay", []int{25,15,69}, 21}
	s1.setAge(5)
	fmt.Println(s1)
}
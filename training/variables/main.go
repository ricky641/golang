package main

import "fmt"

func main() {
	//Different possibilities of initialising and declaring variables
	//	var a int   //just declaration as nil value
	//	var b string
	//	var c float64
	//a = 5 //if var a is initially declared
	//
	//var a,b,c int - declare many at once
	//var a,b,c [int] = 1,2,3  --- Init many at once - type is optional
	//following can only be done inside function
	//var a,b,c = 1, false, "Hello"
	a := 5 //declared and assigned value
	b := "Hello"
	c := 3.4

	fmt.Printf("%v %T", a, a)
	fmt.Printf("%v %T", b, b)
	fmt.Printf("%v %T", c, c)
}

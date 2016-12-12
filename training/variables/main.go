package main

import "fmt"

//constants
const firstConst string = "Hello"

//multiple constants
const (
	name = "Rick"
	age  = 24
)

//kind of like enums
const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

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

	fmt.Printf("%v %T %v \t", a, a, name)
	fmt.Printf("%v %T\t", b, b)
	fmt.Printf("%v %T\t", c, c)
	//anonymous functions
	// increment := func(myvar int) int {
	// 	myvar++
	// 	return myvar
	// }
	increment2 := wrapper(10)     //just a assignment and function will be not be called at this stage
	fmt.Println(increment2(), &b) //& memory address

	var meters float64

	fmt.Println("Enter something")
	fmt.Scan(&meters) //taking input from user

	result := meters * 1.6

	fmt.Printf("%f", result)
}

//anonymous functions as return type
func wrapper(a int) func() int {
	return func() int {
		a++
		return a
	}
}

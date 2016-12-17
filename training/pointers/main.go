package main

import "fmt"

func main() {

	//basic pointers
	a := 20

	fmt.Println(a)
	fmt.Println(&a)

	b := &a //storing memory address

	fmt.Println(b)
	fmt.Println(*b) //retrieving value from that memory address

	*b = 42 //changing value at root memory address

	fmt.Println(a)

	//using pointers in functions
	z := 0
	fmt.Println(z)
	reference(&z)
	fmt.Print(z)
}

func reference(z *int) {
	*z = 6
}

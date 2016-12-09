package main

import "fmt"

func main() {

	//for fmt reference click here https://godoc.org/fmt
	fmt.Print("Hello from the other side")                           //without new line
	fmt.Println("Hello from the other side")                         //with new line
	fmt.Printf("%d \t %b \t %x \t %#x \t %q\n", 'A', 23, 23, 23, 65) //with formatted text

}

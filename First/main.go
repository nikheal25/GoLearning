package main

import "fmt"

func main()  {
	fmt.Println("Hello there")

	var nameOne string = "General"
	var nameTwo = "Kenobi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameThree = "!"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "nick"
	// the above logic only works in the function. Not the outside of the function.

	fmt.Println(nameFour)

	//* Integers
	var intOne int = 20
	var intTwo = 30
	var intThree = 40

	fmt.Println(intOne, intTwo, intThree)

	//* Bits and memory
	var numOne int8 = 20
	var numTwo int8 = -128
	var numThree uint16 = 256

	fmt.Println(numOne, numTwo, numThree)

}
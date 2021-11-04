package main

import "fmt"

func main() {
	star := "Polaris"
	starAddress := &star
	fmt.Println("The value of star is", star)

	*starAddress = "Sirius"
	fmt.Println("The actual value of star is", star)
}

/*
We can use our pointer to access the address and change its value!
This action is called dereferencing or indirecting.
We need to use the * operator again on a pointer and then assign a new value.
*/

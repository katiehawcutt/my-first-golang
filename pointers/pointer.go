package main

import "fmt"

func main() {
	star := "Polaris"
	starAddress := &star
	fmt.Println("The address of the star is", starAddress)
}

/*
Pointers are variables that specifically store addresses.
The * operator signifies that this variable will store an address
*/

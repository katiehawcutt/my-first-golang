package main

import "fmt"

// the function has a parameter of a pointer for a string
func brainwash(saying *string) {
	// Dereference saying below:
	*saying = "Beep Boop."
}

func main() {
	greeting := "Hello there!"
	fmt.Println("greeting is:", greeting)

	// To change a value in a different scope we must pass the pointer to a function
	brainwash(&greeting)
	fmt.Println("greeting is now:", greeting)
}

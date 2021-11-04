package main

import "fmt"

func main() {
	treasure := "The friends we make along the way."
	fmt.Println(&treasure)
}

/*
Go is a pass-by-value language.
A computer sets aside some space in its memory to store a value.
The space that the computer allocates is called an address.
An address is where a value is stored.
Each address is marked as a unique numerical value.
Every time we use a variable, we are retrieving the value stored at the variable’s address.
To find a variable’s address we use the & operator followed by the variable name.
When we see the 0x prefix, this means that the number is in formatted in hexadecimal,
which is a way of representing 16 digit numbers.
*/

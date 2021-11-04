package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	amountLeft := rand.Intn(1000)

	fmt.Println("amountLeft is: £", amountLeft)

	if amountLeft > 300 {
		fmt.Println("What should I spend this on?")
	} else {
		fmt.Println("Where did all my money go?")
	}
}

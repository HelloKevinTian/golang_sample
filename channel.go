package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Generate two random numbers.
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	x := rand.Intn(100)
	y := rand.Intn(100)

	// Create channels for sum and difference.
	sum := make(chan int)
	dif := make(chan int)
	ret1 := make(chan bool)
	ret2 := make(chan bool)

	// Start a goroutine to calculate the sum.
	go func() {
		sum <- x + y // Send the result to channel.
		ret1 <- true
	}()

	// Start a goroutine to calculate the difference.
	go func() {
		dif <- x - y // Send the result to channel.
		ret2 <- false
	}()

	// Receive the results from the channels and print.
	fmt.Printf("sum=%v dif=%v\n", <-sum, <-dif)
	fmt.Println(<-ret1, <-ret2)
}

//output
// sum=168 dif=-6
// true false

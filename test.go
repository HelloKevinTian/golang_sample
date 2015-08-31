package main

var a string

var c = make(chan int)

func f() {

	a = "hello, world"

	// <-c

	c <- 0

}

func main() {

	go f()

	// c <- 0

	<-c

	print(a)

}

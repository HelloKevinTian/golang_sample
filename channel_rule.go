package main

/*
	规则一：channel带缓冲：先 c <- 0，再 <- c
*/
// var c = make(chan int, 10) //带缓冲的channel
// var a string

// func f() {
// 	a = "hello, world"
// 	c <- 0
// }

// func main() {
// 	go f()
// 	<-c
// 	print(a)
// }

/*
	规则二：channel不带缓冲：先 <- c，再 c <- 0
*/
var c = make(chan int) //不带缓冲的channel
var a string

func f() {
	a = "hello, world"
	<-c
}
func main() {
	go f()
	c <- 0
	print(a)
}

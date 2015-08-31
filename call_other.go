package main

import (
	"fmt"
	"test"
)

func main() {
	fmt.Println("print from self")
	test.Test()
	test.Call()
}

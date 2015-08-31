package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("cpu num:", runtime.NumCPU())
	var chs = make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Add(i, i, chs[i])
	}

	for _, ch := range chs {
		i := <-ch
		fmt.Println("main", i)
	}
}

func Add(x, y int, ch chan int) {
	time.Sleep(1e9) //wait for 1 seconds
	z := x + y
	fmt.Println("goroutine", z)
	ch <- z
}

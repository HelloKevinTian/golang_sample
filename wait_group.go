// 一个 WaitGroup 等待一组协程的结束
// 主协程调用Add函数设置需要等待的协程的数量
// 然后每个协程运行并在结束时调用Done函数
// 与此同时，Wait函数用于阻塞，直到所有协程都执行结束
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.baiyuxiong.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			http.Get(url)
			fmt.Println(url)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
	fmt.Println("over")
}

//output
// http://www.baiyuxiong.com/
// http://www.google.com/
// http://www.golang.org/
// over

// 1、golang中有2种方式同步程序，一种使用channel，另一种使用锁机制。
// 这里要涉及的是锁机制，更具体的是sync.WaitGroup，一种较为简单的同步方法集。

// 2、sync.WaitGroup只有3个方法，Add()，Done()，Wait()。其中Done()是Add(-1)的别名。
// 简单的来说，使用Add()添加计数，Done()减掉一个计数，计数不为0, 阻塞Wait()的运行。

// 3、要注意的有一点。sync文档已经说明了的，The main goroutine calls Add to set the number of goroutines to wait for.
// Then each of the goroutines runs and calls Done when finished.也就是说，在运行main函数的goroutine里运行Add()函数，
// 在其他的goroutine里面运行Done()函数。这个我是踩过雷了的。

package main

import (
	"fmt"
)

func main() {
	if []byte("aaa")[0] == []byte("ccc")[0] {
		fmt.Println([]byte("aaa")[0])
		fmt.Println([]byte("ccc")[0])
		fmt.Println("ok")
	}
}

// output
// 97
// 99
// ok

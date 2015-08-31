package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
	//必须放在随机函数之前，否则每次随机的结果会是一样的
    timens := int64(time.Now().Nanosecond())
    rand.Seed(timens)

    for i := 0; i < 10; i++ {
        a := rand.Int()
        fmt.Printf("%d /\n ", a)
    }

    fmt.Println()

    for i := 0; i < 5; i++ {
        r := rand.Intn(8)
        fmt.Printf("%d /\n ", r)
    }

    fmt.Println()
    
    for i := 0; i < 10; i++ {
        fmt.Printf("%2.2f /\n ", 100*rand.Float32())
    }

    fmt.Println()
}
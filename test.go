package main

import (
	"fmt"
)

func main() {
	m := make(map[interface{}]interface{})
	// slice := []int{1, 2}
	// slice1 := make([]string, 5)
	// arr := [3]string{"a", "b", "c"}
	// aaa("4")
	// aaa("str")
	aaa(m)
	// aaa(slice1)
	// aaa(slice1)
	// aaa(arr)
}

func aaa(a interface{}) {
	fmt.Println(a.([]string))
}

package main

import (
	"fmt"
)

func main() {
	// //----------------------------------
	// var arr [5]int

	// for i := 0; i < len(arr); i++ {
	// 	arr[i] = i + 100
	// }

	// for i, v := range arr {
	// 	fmt.Println("arr item ", i, "is", v)
	// }

	// //----------------------------------
	// a := [...]string{"a", "b", "c", "d"}
	// for i := range a {
	// 	fmt.Println("Array item", i, "is", a[i])
	// }

	// //----------------------------------
	// var arr1 = new([5]int) //引用类型
	// // var arr1 [5]int            //值类型

	// for i := 0; i < len(arr1); i++ {
	// 	arr1[i] = i + 10000
	// }

	// arr2 := arr1
	// arr2[2] = 888
	// arr3 := arr2
	// arr3[2] = 777

	// for i, v := range arr1 {
	// 	fmt.Println("arr1 item", i, "is", v)
	// }
	// for i, v := range arr2 {
	// 	fmt.Println("arr2 item", i, "is", v)
	// }
	// for i, v := range arr3 {
	// 	fmt.Println("arr3 item", i, "is", v)
	// }

	// //----------------------------------
	var arrAge = []int{18, 20, 15, 22, 16}
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	slice := arrAge
	slice[0] = 111111

	for i, v := range arrAge {
		fmt.Println("arrAge item", i, "is", v)
	}
	for i, v := range slice {
		fmt.Println("slice item", i, "is", v)
	}

	//----------------------------------
	for i := 0; i < 3; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}
}

func fp(a *[3]int) { fmt.Println(a) }

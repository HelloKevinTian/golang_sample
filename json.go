package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

type Animal struct {
	Name  string
	Order string
}

func main() {
	//encode-------------------------
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)

	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(b)

	fmt.Println("")
	fmt.Println(string(b))

	//decode-----------------------------
	var jsonBlob = []byte(`[  
        {"Name": "Platypus", "Order": "Monotremata"},  
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}  
    ]`)

	var animals []Animal

	err1 := json.Unmarshal(jsonBlob, &animals)

	if err1 != nil {
		fmt.Println("error:", err1)
	}

	fmt.Println(animals)
}

// Encode Output:
// {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
// {"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}

// Decode Output:
// [{Platypus Monotremata} {Quoll Dasyuromorphia}]

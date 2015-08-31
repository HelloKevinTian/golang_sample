// package main

// import (
// 	"fmt"
// 	"github.com/codegangsta/inject"
// )

// type SpecialString interface{}

// func Say(name string, gender SpecialString, age int) {
// 	fmt.Printf("My name is %s, gender is %s, age is %d!\n", name, gender, age)
// }

// func main() {
// 	inj := inject.New()
// 	inj.Map("陈新")
// 	inj.MapTo("男", (*SpecialString)(nil))
// 	inj.Map(20)
// 	inj.Invoke(Say)
// }

//output
// My name is 陈新, gender is 男, age is 20!

//--------------------------------------------
// package main

// import (
// 	"fmt"
// 	"github.com/codegangsta/inject"
// )

// type SpecialString interface{}

// func main() {
// 	fmt.Println(inject.InterfaceOf((*interface{})(nil)))
// 	fmt.Println(inject.InterfaceOf((*SpecialString)(nil)))
// }

//output
// interface {}
// main.SpecialString

//------------------------
// package main

// import (
// 	"fmt"
// 	"github.com/codegangsta/inject"
// 	"reflect"
// )

// type SpecialString interface{}

// func main() {
// 	inj := inject.New()
// 	inj.Map("陈一回")
// 	inj.MapTo("男", (*SpecialString)(nil))
// 	inj.Map(20)
// 	fmt.Println("string is valid?", inj.Get(reflect.TypeOf("姓陈名一回")).IsValid())
// 	fmt.Println("SpecialString is valid?", inj.Get(inject.InterfaceOf((*SpecialString)(nil))).IsValid())
// 	fmt.Println("int is valid?", inj.Get(reflect.TypeOf(18)).IsValid())
// 	fmt.Println("[]byte is valid?", inj.Get(reflect.TypeOf([]byte("Golang"))).IsValid())
// 	inj2 := inject.New()
// 	inj2.Map([]byte("test"))
// 	inj.SetParent(inj2)
// 	fmt.Println("[]byte is valid?", inj.Get(reflect.TypeOf([]byte("Golang"))).IsValid())
// }

//output
// string is valid? true
// SpecialString is valid? true
// int is valid? true
// []byte is valid? false
// []byte is valid? true

//---------------------------------
// package main

// import (
// 	"fmt"
// 	"github.com/codegangsta/inject"
// )

// type SpecialString interface{}

// func Say(name string, gender SpecialString, age int) {
// 	fmt.Printf("My name is %s, gender is %s, age is %d!\n", name, gender, age)
// }

// func main() {
// 	inj := inject.New()
// 	inj.Map("陈一回")
// 	inj.MapTo("男", (*SpecialString)(nil))
// 	inj2 := inject.New()
// 	inj2.Map(20)
// 	inj.SetParent(inj2)
// 	inj.Invoke(Say)
// }

//output
// My name is 陈一回, gender is 男, age is 20!

//---------------------------
package main

import (
	"fmt"
	"github.com/codegangsta/inject"
)

type SpecialString interface{}
type TestStruct struct {
	Name   string `inject`
	Nick   []byte
	Gender SpecialString `inject`
	uid    int           `inject`
	Age    int           `inject`
}

//结论：Apply方法是用于对struct的字段进行注入，参数为指向底层类型为结构体的指针。
//可注入的前提是：字段必须是导出的(也即字段名以大写字母开头)，并且此字段的tag设置为`inject`
func main() {
	s := TestStruct{}
	inj := inject.New()
	inj.Map("陈一回")
	inj.MapTo("男", (*SpecialString)(nil))
	inj2 := inject.New()
	inj2.Map(20)
	inj.SetParent(inj2)
	inj.Apply(&s)
	fmt.Println("s.Name =", s.Name)
	fmt.Println("s.Gender =", s.Gender)
	fmt.Println("s.Age =", s.Age)
	fmt.Println("s.Nick =", s.Nick)
	fmt.Println("s.uid =", s.uid)
}

//output
// s.Name = 陈一回
// s.Gender = 男
// s.Age = 20
// s.Nick = []
// s.uid = 0

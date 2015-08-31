package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	TestString := "Hello wolrd"

	Md5Inst := md5.New()
	Md5Inst.Write([]byte(TestString))
	Md5Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n", Md5Result)

	Sha1Inst := sha1.New()
	Sha1Inst.Write([]byte(TestString))
	Sha1Result := Sha1Inst.Sum([]byte(""))
	fmt.Printf("%x\n", Sha1Result)
}

//output
// 5304e24786491b6c0457f9ba32311215
// 613f9a889d27d04b42ef3423bce364b5729140a9

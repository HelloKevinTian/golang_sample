package main

import (
	"log"
	"os"
)

func main() {
	fileName := "logger_test.log"
	logFile, err := os.Create(fileName)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	logger := log.New(logFile, "[Debug]", log.Llongfile)
	logger.Println("A debug message here")
	logger.SetPrefix("[Info]")
	logger.Println("A Info Message here ")
	logger.SetFlags(logger.Flags() | log.LstdFlags)
	logger.Println("A different prefix")
}

//output
// [Debug]/share/golang/logger.go:16: A debug message here
// [Info]/share/golang/logger.go:18: A Info Message here
// [Info]2015/08/26 15:08:50 /share/golang/logger.go:20: A different prefix

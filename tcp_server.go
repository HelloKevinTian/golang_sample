package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()

		log.Println("有一个客户连接\n")

		if err != nil {
			log.Println(err)
			continue
		}

		go echoFunc(conn)
	}
}

func echoFunc(c net.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("收到客户端数据1：", string(buf[:n]))
		log.Printf("收到客户端数据2： %s", buf[:n])
		c.Write(buf[:n])
	}
}

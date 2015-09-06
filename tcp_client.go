package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err == nil {
		conn.Write([]byte("hello world!"))

		addr := conn.LocalAddr()
		addr1 := conn.RemoteAddr()
		log.Println("客户端 & 服务器地址：", addr, addr1)

		for {
			buf := make([]byte, 1024)
			if length, err := conn.Read(buf); err == nil {
				if length > 0 {
					buf[length] = 0
					log.Printf("收到服务端反馈数据：%s", string(buf[:length]))
				}
			}
		}
	} else {
		log.Println("conn err")
		return
	}
}

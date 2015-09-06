package tcp_game_server

import (
	"log"
	"net"
)

func Start(host string) {
	ln, err := net.Listen("tcp", host)
	checkError(err)

	addr := ln.Addr()
	log.Println("服务器正在监听：", addr)

	for {
		conn, err := ln.Accept()
		checkError(err)

		client := conn.RemoteAddr()
		log.Println("有一个客户端连接：", client)

		go run(conn)
	}
}

func run(c net.Conn) {
	buf := make([]byte, 1024)
	defer c.Close()

	for {
		n, err := c.Read(buf)
		checkError(err)

		log.Println("收到客户端数据：", string(buf[:n]))
		c.Write(buf[:n])
	}
}

func checkError(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}

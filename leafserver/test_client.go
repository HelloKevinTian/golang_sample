package main

import (
	"encoding/binary"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	data := []byte(`{
        "Hello": {
            "Name": "kevin"
        }
    }`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// 发送消息
	log.Println("发送数据：", string(m))
	conn.Write(m)

	for {
		buf := make([]byte, 1024)
		if length, err := conn.Read(buf); err == nil {
			if length > 0 {
				buf[length] = 0
				log.Printf("收到服务端反馈数据：", string(buf[:length]))
			}
		}
	}
}

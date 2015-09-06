package main

import (
	"log"
	"tcp_game_server"
)

func main() {
	log.Println("==================tcp game server start=============")
	tcp_game_server.Start(":8080")
}

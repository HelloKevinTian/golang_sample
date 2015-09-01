/**
 * @ Author Kevin
 * @ Email  tianwen@chukong-inc.com
 * @ 2015/9/1
 */

package main

import (
	"game_web_server"
	"log"
)

func main() {
	log.Println("===============game web server===================")
	log.Println("               游戏后台管理系统")
	log.Println("===============game web server===================")
	game_web_server.Start(":8888", ":6379")
}

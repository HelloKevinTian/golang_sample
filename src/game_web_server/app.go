package main

import (
	"github.com/go-martini/martini"
	// "github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	// "html/template"
	"github.com/garyburd/redigo/redis"
	"log"
	// "net/http"
	"time"
)

func main() {
	// init db
	//args: 网络类型“tcp”	地址和端口	连接超时	读超时	写超时时间
	conn, err := redis.DialTimeout("tcp", ":6379", 0, 1*time.Second, 1*time.Second)
	// conn, err := redis.Dial("tcp", ":6379")
	checkErr(err, "redis conn error")

	size, _ := redis.Int(conn.Do("DBSIZE"))
	log.Println("db size is:", size)

	mail, _ := redis.Strings(conn.Do("HGETALL", "h_mail"))
	log.Println(mail)
	log.Println(mail[0], mail[3])

	// lets start martini and the real code
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello", mail)
	})

	m.RunOnAddr(":8888")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

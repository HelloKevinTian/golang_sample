package game_web_server

import (
	"github.com/go-martini/martini"
	// "github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	// "html/template"
	"github.com/garyburd/redigo/redis"
	"log"
)

var pool *redis.Pool

func Start(server, dbserver string) {
	// db
	pool = newPool(dbserver)
	conn := pool.Get()
	defer conn.Close()

	size, _ := redis.Int(conn.Do("DBSIZE"))
	log.Println("db size is:", size)

	// server
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "login", "game web server by go")
	})

	m.Post("/login", loginHandle)

	m.RunOnAddr(server)
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

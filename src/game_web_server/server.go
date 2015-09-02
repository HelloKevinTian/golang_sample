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

/*
 * 正式环境设置：export MARTINI_ENV=production
 */
func Start(server, dbserver string) {
	// db
	pool = newPool(dbserver)
	conn := pool.Get()
	defer conn.Close()

	size, _ := redis.Int(conn.Do("DBSIZE"))
	log.Println("db size is:", size)

	// server
	ch := make(chan bool, 1)
	go func() {
		m := martini.Classic()

		//==============================中间件=============================
		m.Use(render.Renderer())

		// m.Use(func(res http.ResponseWriter, req *http.Request) {
		// 	if req.Header.Get("X-API-KEY") != "secret123" {
		// 		res.WriteHeader(http.StatusUnauthorized)
		// 	}
		// })

		// log 记录请求完成前后  (*译者注: 很巧妙，掌声鼓励.)
		m.Use(func(c martini.Context, log *log.Logger) {
			log.Println("before a request")

			c.Next()

			log.Println("after a request")
		})

		//==============================路由===============================
		m.Get("/", func(r render.Render) {
			r.HTML(200, "login", "game web server by go")
		})

		// m.Get("/hello/:name", func(params martini.Params) string {
		// 	return "Hello " + params["name"]
		// })
		// m.Get("/hello/**", func(params martini.Params) string {
		// 	return "Hello " + params["_1"]
		// })

		m.Post("/login", loginHandle)

		m.NotFound(func(r render.Render) {
			// 处理 404
			r.HTML(200, "hello", " --from ck")
		})

		//==============================运行===============================
		m.RunOnAddr(server)

		ch <- true
	}()

	<-ch
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

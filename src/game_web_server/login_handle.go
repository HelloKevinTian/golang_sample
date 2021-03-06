package game_web_server

import (
	// "encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/martini-contrib/render"
	"net/http"
)

// type BaseJsonBean struct {
// 	Code    int         `json:"code"`
// 	Data    interface{} `json:"data"`
// 	Message string      `json:"message"`
// }

// func NewBaseJsonBean() *BaseJsonBean {
// 	return &BaseJsonBean{}
// }

func loginHandle(w http.ResponseWriter, req *http.Request, r render.Render) {
	conn := pool.Get()
	defer conn.Close()
	//获取客户端通过GET/POST方式传递的参数
	req.ParseForm()
	param_username, found1 := req.Form["username"]
	param_password, found2 := req.Form["password"]

	if !(found1 && found2) {
		fmt.Fprint(w, "请勿非法访问")
		return
	}

	// result := NewBaseJsonBean()
	username := param_username[0]
	password := param_password[0]

	s := "username:" + username + ",password:" + password
	fmt.Println(s)

	pass, _ := redis.String(conn.Do("hget", "h_user", username))

	if pass == password {
		// result.Code = 100
		// result.Message = "登录成功"

		arr := []string{"baoma", "aodi", "benchi", "cc"}
		ma1 := map[string]string{"name": "kevin", "age": "25"}
		ma2 := map[string]string{"name": "joe", "age": "33"}
		ma3 := map[string]string{"name": "lucy", "age": "21"}
		var map_arr []map[string]string
		map_slice := append(map_arr, ma1, ma2, ma3)
		grid := map[string]interface{}{"arr": arr, "map_arr": map_slice}
		r.HTML(200, "index", grid)
		// r.Redirect("hello", 200)
	} else {
		// result.Code = 101
		// result.Message = "用户名或密码不正确"
		r.HTML(200, "login", "game web server by go")
	}

	//向客户端返回JSON数据
	// bytes, _ := json.Marshal(result)
	// fmt.Fprint(w, string(bytes))
}

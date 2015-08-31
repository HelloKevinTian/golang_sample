// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	response, _ := http.Get("http://211.151.21.70:20003/test_status")
// 	defer response.Body.Close()
// 	body, _ := ioutil.ReadAll(response.Body)
// 	fmt.Println(string(body))
// }

//另一种实现方式
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", "http://211.151.21.70:20003/test_status", nil)

	reqest.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	reqest.Header.Set("Accept-Charset", "GBK,utf-8;q=0.7,*;q=0.3")
	reqest.Header.Set("Accept-Encoding", "gzip,deflate,sdch")
	reqest.Header.Set("Accept-Language", "zh-CN,zh;q=0.8")
	reqest.Header.Set("Cache-Control", "max-age=0")
	reqest.Header.Set("Connection", "keep-alive")

	response, _ := client.Do(reqest)
	if response.StatusCode == 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodystr := string(body)
		fmt.Println(bodystr)
	}
}

//第三种方式
// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// )

// func main() {
// 	resp, err := http.Get("http://211.151.21.70:20003/test_status")
// 	if err != nil {
// 		// 处理错误 ...
// 		return
// 	}
// 	fmt.Println(resp)
// 	fmt.Println(resp.Body)
// 	defer resp.Body.Close()
// 	io.Copy(os.Stdout, resp.Body)
// }

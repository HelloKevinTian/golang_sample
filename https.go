package main

import (
	"fmt"
	"net/http"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	// w.Write([]byte("Hello"))

	if r.Method == "GET" {
		w.Header().Set("Content-Type", "text/html")

		io.WriteString(w, "<form method=\"POST\" action=\"/upload\" "+
			" enctype=\"multipart/form-data\">"+
			"Choose an image to upload: <input name=\"image\" type=\"file\" />"+
			"<input type=\"submit\" value=\"Upload\" />"+
			"</form>")
		return
	}

func main() {
	http.HandleFunc("/hello", SayHello)
	fmt.Println("start server @ :8001")
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

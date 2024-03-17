package main

import (
	"fmt"
	"net/http"
)

func IndexPathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World") // 웹 핸들러 등록
}

func main() {
	http.Handle("/static", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/hello", IndexPathHandler)
	http.ListenAndServe(":3000", nil) // 웹 서버 시작
}

//stripPrefix

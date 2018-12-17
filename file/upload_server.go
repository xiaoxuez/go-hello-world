package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/upload", upload)
	http.HandleFunc("/hi", hi)
	err := http.ListenAndServe("0.0.0.0:8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
func hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func upload(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method) //POST

	//因为上传文件的类型是multipart/form-data 所以不能使用 r.ParseForm(), 这个只能获得普通post
	r.ParseMultipartForm(32 << 20) //上传最大文件限制32M

	user := r.Form.Get("username")
	password := r.Form.Get("projectName")

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err, "--------1------------") //上传错误
	}
	defer file.Close()

	fmt.Println(user, password, handler.Filename) //test 123456 json.zip
}

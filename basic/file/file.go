package main

import (
	"fmt"
	"path"
)

func main() {
	//f, err := os.Open("test.txt")
	//if err != nil {
	//	panic(err.Error())
	//}
	//_, err = f.Write([]byte("xxx"))
	//if err != nil {
	//	panic(err)
	//}
	dir := "/Users/xiaoxuez/go/src/github.com/xiaoxuez/go-hello-world/basic/file"
	fmt.Println(path.Base(dir))
}

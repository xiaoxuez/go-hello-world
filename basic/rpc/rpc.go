package rpc

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

/**
rpc的实现有http,tcp等实现方式
*/

//方法输入参数，需要都为exported
type Args struct {
	A, B int
}

//方法返回参数，需要都为exported
type Reply struct {
	S int
}

//方法接收者，int只是一个辅助作用
type Cal int

//api方法要求，方法接收者和参数都需要exported, 参数为两个，第一个为调用参数，第二个为返回参数，如果error不为nil,则返回error
func (c *Cal) Add(args *Args, result *int) error {
	*result = args.A + args.B
	return nil
}

//上述为rpc中的方法，以下则为如何与conn绑定使用
//最简单的方式，首先是Register，然后需要跟一个conn绑定，HandleHTTP就会跟http包中默认的server进行绑定，然后直接启动http请求就好了
func MakeHttp() {
	rpc.Register(new(Cal))
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)
}

//那么，http包中只有一个默认的server,如果要多个端口的话，就gg了，需要自定义server
func MakeServerHttp() {
	handler := rpc.NewServer()
	handler.Register(new(Cal))
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
	}))
}

//前两种都是选择http作为rpc的载体，下面尝试使用tcp作为rpc的载体
func MakeTcp() {
	rpc.Register(new(Cal))
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	l, e := net.ListenTCP("tcp", tcpAddr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go func(l *net.TCPListener) {
		for {
			conn, err := l.Accept()
			if err != nil {
				continue
			}
			rpc.ServeConn(conn)
		}
	}(l)
}

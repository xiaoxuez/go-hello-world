/*
@Time : 2018/12/3 下午2:22
@Author : xiaoxuez

go-stack通过封装runtime包，提供简单的api。
实现了捕获，操作和格式化调用堆栈的实用程序。

该实现负责解释由runtime.Callers返回的解释程序计数器（pc）值的细节和特殊情况。

*/

package main

import (
	"fmt"
	"github.com/go-stack/stack"
)

//基本使用，获取当前调用
func doTheThing() {
	c := stack.Caller(0)
	fmt.Println(c)         // "stack.go:21"
	fmt.Printf("%+v\n", c) // "stack/stack.go:21"
	fmt.Printf("%n\n", c)  // "doTheThing"
	s := stack.Trace().TrimRuntime()
	fmt.Println(s) // "[stack.go:25 stack.go:31]"
}

//stack.Caller(0), 这个0参数...
//Caller returns a Call from the stack of the current goroutine. The argument
//skip is the number of stack frames to ascend, with 0 identifying the calling function
//大概理解应该是到调用 上面有很多层(最下是自己 -> 调用方 -> 调用方.. )，
// skip就是跳过的层。如本示例中0的话 就是自己(第21行)， 1的话 就是调用方36行，再往上就再有调用方了..

func main() {
	doTheThing()
}

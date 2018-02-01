package main

import "fmt"
import "time"

func serve(ch <-chan int) { //分发函数
	for val := range ch {
		go handle(val)
	}
}
func handle(x int) { //处理函数
	defer func() {
		if err := recover(); err != nil {
			//err = 3,有panic传递。若仍需继续异常，则再次使用panic即可
			fmt.Println("work failed:", err)
		}
	}()
	if x == 3 {
		panic(x)
	}
	fmt.Println("work succeeded:", x)
}

func main() {
	var ch = make(chan int, 6)
	for i := 1; i <= 6; i++ {
		ch <- i
	}
	close(ch)
	go serve(ch)
	time.Sleep(time.Millisecond * 100)
	//执行结果为，
	/**
	work succeeded: 6
	work succeeded: 1
	work succeeded: 2
	work failed: 3 //捕获到了异常，只是结束当前线程，不会影响整个应用的行为
	work succeeded: 4
	work succeeded: 5
	*/

}

/**
Panic用法，
	感觉上来说有点像异常，调用panic之后，程序将异常退出。包括自身线程和主线程。疑问是，go的线程
本来就不好管理，某个线程异常了，会影响其他兄弟线程吗？..
	异常退出的顺序是，当前方法立即停止，然后调用当前方法的defer方法，然后异常返回给当前方法的调用者，
然后同样一直停止到所有方法调用完(感觉上有点像方法栈依次推掉方法)

Recover用法，
	感觉上有点像捕获过程，为了达到捕获效果，需要在defer方法中进行捕获，不然肯定就调用不到啊。
    捕获的异常只能是当前方法中的异常。


就以往浅薄经验上来说话，还是一样，异常尽量向上抛，在最顶层进行异常捕获。在go里就很简单了，在顶层方法中使用defer + recover就好啦
*/

package main

import (
	"log"
	"strconv"
	"time"
)

func main() {
	withoutBuffer()
}

/**
没有缓冲的通道。
输出为
 					start wait ~
2018/08/09 16:23:49 after wait.  1
2018/08/09 16:23:49  start wait ~
2018/08/09 16:23:49    main-loop:   1
因为line := waitSometime(<-scheduler)是先等的，case里的是后等的，所以每次都是waitSometime中的<-scheduler先执行，再试case中的
*/
func withoutBuffer() {
	scheduler := make(chan string)

	go func() {
		for {
			log.Println(" start wait ~ ")
			line := waitSometime(<-scheduler)
			log.Println("after wait. ", line)
			scheduler <- line
		}
	}()

	for {
		scheduler <- "a"
		select {
		case line, ok := <-scheduler:
			if ok {
				log.Println("   main-loop:  ", line)
			}
		}
	}
}

var increment = 0

func waitSometime(s2 string) string {
	time.Sleep(1 * time.Second)
	increment++
	return strconv.Itoa(increment)
}

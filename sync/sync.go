/*
@Time : 2018/12/3 下午3:02
@Author : xiaoxuez

sync包下的一些使用

1. atomic包的基本使用，提供了几种简单类型的低级原子操作
包括增减、比较并交换、载入、存储和交换。
2. Cond 条件变量，
3. Map: 协程安全的concurrent Map
4. mutex: 互斥锁
5. once: 只会执行一次的对象
6. pool: 池的目的是缓存已分配但未使用的项目以供以后重用，从而减轻对垃圾收集器的压力
7. rwmutex: 读写锁
8. waitgroup: 控制等待一组协程执行完毕
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

/**
文档中记载，除了使用特殊的或者低级操作，同步最好还是采用channels或者sync包下的功能来实现。
通过沟通来进行共享内存，不要通过共享内存来沟通。
提供的原子性操作的类型包括，int32\int64\uint32\uint64\uintptr\unsafe.Pointer\Value
基本操作包括，Add\CompareAndSwap\Load\Store\Swap\
*/
func AtomicExample() {
	//这个例子展示了使用复制习惯如何维护一个频繁读、但更新不频繁的数据结构。
	//m.Store(make(Map))
	//m.Load()
	type Map map[string]string
	var m atomic.Value //interface类型
	m.Store(make(Map))
	var mu sync.Mutex // used only by writers
	// read function can be used to read the data without further synchronization
	read := func(key string) (val string) {
		m1 := m.Load().(Map)
		return m1[key]
	}
	// insert function can be used to update the data without further synchronization
	insert := func(key, val string) {
		mu.Lock() // synchronize with other potential writers
		defer mu.Unlock()
		m1 := m.Load().(Map) // load current value of the data structure
		m2 := make(Map)      // create a new value
		for k, v := range m1 {
			m2[k] = v // copy all data from the current object to the new one
		}
		m2[key] = val // do the update that we need
		m.Store(m2)   // atomically replace the current object with the new one
		// At this point all new readers start working with the new version.
		// The old version will be garbage collected once the existing readers
		// (if any) are done with it.
	}
	_, _ = read, insert
	//Value的类型每次更新都应该为新的值。

	var locationEnabled uint32
	atomic.StoreUint32(&locationEnabled, 1)
	value := atomic.LoadUint32(&locationEnabled)
	fmt.Println(value, locationEnabled)
}

/**
条件变量:

与互斥量不同，条件变量的作用并不是保证在同一时刻仅有一个线程访问某一个共享数据，
而是在对应的共享数据的状态发生变化时，通知其他因此而被阻塞的线程。
条件变量总是与互斥量组合使用。互斥量为共享数据的访问提供互斥支持，
而条件变量可以就共享数据的状态的变化向相关线程发出通知。

*/
func CondExample() {
	cond := sync.NewCond(new(sync.Mutex))
	condition := 0

	// Consumer
	go func() {
		for {
			cond.L.Lock() //好尴尬  不知道这个锁是什么意思..这个没解锁下方又能锁..
			//这个锁应该跟wait和notify有挂钩...
			for condition == 0 {
				cond.Wait() //等待挂起，收到Signal会唤醒，重新判断条件
			}
			condition--
			fmt.Printf("Consumer: %d\n", condition)
			cond.Signal()
			cond.L.Unlock()
		}
	}()

	// Producer
	for {
		time.Sleep(time.Second)
		cond.L.Lock()
		for condition == 3 {
			cond.Wait()
		}
		condition++
		fmt.Printf("Producer: %d\n", condition)
		cond.Signal() //单发通知， 广播通知为broadcast方法
		cond.L.Unlock()
	}
	//output: Producer: 1\Consumer: 0\Producer: 1\Consumer: 0\Producer: 1\Consumer: 0
}

func OnceExample() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody) //只会执行一次,之后就直接跳过
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
		fmt.Println(i)
	}
	//output: Only once/0/1/2/3/4/...
}

/**
注意sync.Map的Range依旧是无序的，我们需要传给Range方法一个函数对象，
这个闭包函数接受key和value，返回bool，返回true则直接进行下一个循环，类似continue。返回false则直接终止Range，类似break。
*/
func MapExample() {
	var syncMap sync.Map
	//store
	syncMap.Store("key", "value")
	//load, load出来是interface, 需要进行类型转换
	value, ok := syncMap.Load("key")
	_, _ = value, ok
	//delete
	syncMap.Delete("key")
	//range
	syncMap.Range(func(key, value interface{}) bool {
		_, _ = key, value
		return true // return true to range next, false to interrupt
	})
}

/**
对象池的一个很好的例子是fmt例子，它维护一个动态大小的临时输出缓冲区存储。商店在负载下（当许多goroutine正在积极打印时）进行缩放，并在静止时收缩。
*/
func PoolExample() {
	Log(os.Stdout, "path", "/search?q=flowers")
}

var bufPool = sync.Pool{
	New: func() interface{} {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return new(bytes.Buffer)
	},
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(time.UnixDate)
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func WaitGroupExample() {
	var wg sync.WaitGroup
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("finish")
	//wait 10s -> finish
}

func main() {
	AtomicExample()
}

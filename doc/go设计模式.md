## 设计模式

参考自[...](https://books.studygolang.com/go-patterns/)

从设计模式的书(book文件夹中有设计模式的书)中，可以看到设计模式按照**目的**可分为创建型、结构型、行为型，其中，再按照范围来分，可分为用于类和对象。这里，先引用书中的图，列出一些常见的设计模式。

| 范围\目的  |                   创建型                    |                   结构型                    |             行为型             |
| :----: | :--------------------------------------: | :--------------------------------------: | :-------------------------: |
| **类**  |              Factory Method              |                Adapter(类)                | Interpreter、Template Method |
| **对象** | Abstract Factory、Builder、Prototype、Singleton、Object Pool | Adapter(对象)、Bridge、Composite、Decorator、... |    Iterator、Observer、...    |

上面列出了一些常见的设计模式，下面来看看go中的一些设计模式。

#### Builder

将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示。

在Go中，通常会有配置结构来实现相同的行为，但是构成结构体的builder方法中会使用一些检查的模板，例如`if cfg.Field != nil {...}`



##### 实现

```
package car

type Speed float64

const (
    MPH Speed = 1
    KPH       = 1.60934
)

type Color string

const (
    BlueColor  Color = "blue"
    GreenColor       = "green"
    RedColor         = "red"
)

type Wheels string

const (
    SportsWheels Wheels = "sports"
    SteelWheels         = "steel"
)

type Builder interface {
    Color(Color) Builder
    Wheels(Wheels) Builder
    TopSpeed(Speed) Builder
    Build() Interface
}

type Interface interface {
    Drive() error
    Stop() error
}
```

##### 使用

```
assembly := car.NewBuilder().Color(car.RedColor)

familyCar := assembly.Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
familyCar.Drive()

sportsCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
sportsCar.Drive()
```



#### Factory Method

定义一个用于创建对象的接口，让子类决定将哪一个类实例化，Factory Method使一个类的实例化延迟到其子类。

##### 实现

类型定义

```
package data

import "io"

type Store interface {
    Open(string) (io.ReadWriteCloser, error)
}
```

不同创建

```
package data

type StorageType int

const (
    DiskStorage StorageType = 1 << iota
    TempStorage
    MemoryStorage
)

func NewStore(t StorageType) Store {
    switch t {
    case MemoryStorage:
        return newMemoryStorage( /*...*/ )
    case DiskStorage:
        return newDiskStorage( /*...*/ )
    default:
        return newTempStorage( /*...*/ )
    }
}
```

##### 使用

```
s, _ := data.NewStore(data.MemoryStorage)
f, _ := s.Open("file")
```



#### Object Pool

对象池，用于准备和保持许多实例的需求下

##### 实现

```
s, _ := data.NewStore(data.MemoryStorage)
f, _ := s.Open("file")

n, _ := f.Write([]byte("data"))
defer f.Close()
```

##### 使用

```
p := pool.New(2)

select {
case obj := <-p:
    obj.Do( /*...*/ )

    p <- obj
default:
    // No more objects left — retry later or fail
    return
}
```

##### 小规则

- Object pool pattern is useful in cases where object initialization is more expensive than the object maintenance.
- If there are spikes in demand as opposed to a steady demand, the maintenance overhead might overweigh the benefits of an object pool.
- It has positive effects on performance due to objects being initialized beforehand.



#### Singleton 

单例

##### 实现

```
package singleton

type singleton map[string]string

var (
    once sync.Once //sync.Once可保证只执行一次
    instance singleton
)

func New() singleton {
    once.Do(func() {
        instance = make(singleton)
    })

    return instance
}
```

#### 使用

```
s := singleton.New()

s["this"] = "that"

s2 := singleton.New()
```



#### Decorator

包装

##### 实现

```
type Object func(int) int

func LogDecorate(fn Object) Object {
    return func(n int) int {
        log.Println("Starting the execution with the integer", n)

        result := fn(n)

        log.Println("Execution is completed with the result", result)

        return result
    }
}
```

使用

```
func Double(n int) int {
    return n * 2
}

f := LogDecorate(Double)

f(5)
```



#### Proxy

代理，提供了一个object可以控制进入另一个object，截断所有的调用。

例如：

##### 实现

```
// To use proxy and to object they must implement same methods
    type IObject interface {
        ObjDo(action string)
    }

    // Object represents real objects which proxy will delegate data
    type Object struct {
        action string
    }

    // ObjDo implements IObject interface and handel's all logic
    func (obj *Object) ObjDo(action string) {
        // Action behavior
        fmt.Printf("I can, %s", action)
    }

    // ProxyObject represents proxy object with intercepts actions
    type ProxyObject struct {
        object *Object
    }

    // ObjDo are implemented IObject and intercept action before send in real Object
    func (p *ProxyObject) ObjDo(action string) {
        if p.object == nil {
            p.object = new(Object)
        }
        if action == "Run" {
            p.object.ObjDo(action) // Prints: I can, Run
        }
    }
```

ProxyObject 和 Object都是实现了IObject的方法，ProxyObject在内部包装了Object对象，截断所有调用，再决定给不给到Object对象



#### Observer

观察者模式。

```
// Package main serves as an example application that makes use of the observer pattern.
package main

import (
	"fmt"
	"time"
)

type (
	Event struct {
		Data int64
	}
   //观察者，具有推送消息的特性
	Observer interface {
		OnNotify(Event)
	}

	//观察，组织观察者们，在适当的时候推送(Notify)event给观察者们
	Notifier interface {
		// Register allows an instance to register itself to listen/observe
		// events.
		Register(Observer)
		// Deregister allows an instance to remove itself from the collection
		// of observers/listeners.
		Deregister(Observer)
		// Notify publishes new events to listeners. The method is not
		// absolutely necessary, as each implementation could define this itself
		// without losing functionality.
		Notify(Event)
	}
)

type (
	//实现的观察者
	eventObserver struct{
		id int
	}
	//实现的被观察者们
	eventNotifier struct{
		// Using a map with an empty struct allows us to keep the observers
		// unique while still keeping memory usage relatively low.
		observers map[Observer]struct{}
	}
)

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Deregister(l Observer) {
	delete(o.observers, l)
}

func (p *eventNotifier) Notify(e Event) {
	for o := range p.observers {
		o.OnNotify(e)
	}
}

func main() {
	// Initialize a new Notifier.
	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

	// Register a couple of observers.
	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})

	// A simple loop publishing the current Unix timestamp to observers.
	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <- stop:
			return
		case t := <-tick:
			n.Notify(Event{Data: t.UnixNano()})
		}
	}
}
```



#### Strategy

实现算法行为在runtime的时候才被选择。包括定义算法、封装、交互的使用。

##### 实现

```
type Operator interface {
    Apply(int, int) int
}

type Operation struct {
    Operator Operator //将算法封装，由创建Operation的使用者决定Apply算法
}

func (o *Operation) Operate(leftValue, rightValue int) int {
    return o.Operator.Apply(leftValue, rightValue)
}
```

##### 使用

```
type Addition struct{}

func (Addition) Apply(lval, rval int) int {
    return lval + rval
}
```



#### Functional Options

```
package file

type Options struct {
    UID         int
    GID         int
    Flags       int
    Contents    string
    Permissions os.FileMode
}

type Option func(*Options)

func UID(userID int) Option {
    return func(args *Options) {
        args.UID = userID
    }
}

func GID(groupID int) Option {
    return func(args *Options) {
        args.GID = groupID
    }
}

func Contents(c string) Option {
    return func(args *Options) {
        args.Contents = c
    }
}

func Permissions(perms os.FileMode) Option {
    return func(args *Options) {
        args.Permissions = perms
    }
}

//定义实现，巧妙的地方在于，先生成对象args，再修改具体的 setter(args)，Options定义的类型是方法，UID、GID等返回的也都是方法，可直接调用。
package file

func New(filepath string, setters ...Option) error {
    // Default Options
    args := &Options{
        UID:         os.Getuid(),
        GID:         os.Getgid(),
        Contents:    "",
        Permissions: 0666,
        Flags:       os.O_CREATE | os.O_EXCL | os.O_WRONLY,
    }

    for _, setter := range setters {
        setter(args)
    }

    f, err := os.OpenFile(filepath, args.Flags, args.Permissions)
    if err != nil {
        return err
    } else {
        defer f.Close()
    }

    if _, err := f.WriteString(args.Contents); err != nil {
        return err
    }

    return f.Chown(args.UID, args.GID)
}
```

##### 实现

```
emptyFile, err := file.New("/tmp/empty.txt")
if err != nil {
    panic(err)
}

fillerFile, err := file.New("/tmp/file.txt", file.UID(1000), file.Contents("Lorem Ipsum Dolor Amet"))
if err != nil {
    panic(err)
}
```


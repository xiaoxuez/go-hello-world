package main

import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"
)

type MyLogger struct {
	*log.Logger
}

func (l MyLogger) Println(v ...interface{}) {

}

func P(l *log.Logger) {
	l.Println("1111")
}

var mu sync.Mutex

//var cancel context.CancelFunc

func LaLaLa() {

	fmt.Println("001")

	//running:
	for {
		fmt.Println("002")
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("003")
			break
		}
		fmt.Println("004")
	}
	fmt.Println("005")
}

type A interface {
	AA()
}

type a struct {
	name string
}

func (a1 *a) String() string {
	return a1.name
}

func (a1 *a) AA() {

}

func logA(a A) {
	fmt.Println(reflect.TypeOf(a))
	//fmt.Println(reflect.TypeOf(&a))
	b := &a
	fmt.Println(reflect.ValueOf(b))
	fmt.Println(reflect.ValueOf(*b))

	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(*b))
}

type registerStructMaps map[string]reflect.Type

//根据name初始化结构
//在这里根据结构的成员注解进行DI注入，这里没有实现，只是简单都初始化
func (rsm registerStructMaps) New(name string) (c interface{}, err error) {
	if v, ok := rsm[name]; ok {
		c = reflect.New(v).Interface()
	} else {
		err = fmt.Errorf("not found %s struct", name)
	}
	return
}

//根据名字注册实例
func (rsm registerStructMaps) Register(name string, c interface{}) {
	rsm[name] = reflect.TypeOf(c).Elem()
}

type Test struct {
	value string
}

func (test *Test) SetValue(value string) {
	test.value = value
}
func (test *Test) Print() {
	log.Println(test.value)
}
func main() {
	//rsm := registerStructMaps{}
	//	////注册test
	//	//rsm.Register("test", &Test{})
	//	////获取新的test的interface
	//	//test11, _ := rsm.New("test")
	//	//test22, _ := rsm.New("test")
	//	////因为 test11 和 test22都是interface{},必须转换为*Test
	//	//test1 := test11.(*Test)
	//	//test2 := test22.(*Test)
	//	//test1.SetValue("aaa")
	//	//test2.SetValue("bbb")
	//	//test1.Print()
	//	//test2.Print()

	//j := `{a: 1, b: 2}`
	//var jjj struct{}
	//json.Unmarshal([]byte(j), &jjj)
	//
	//v := reflect.ValueOf(&jjj).Elem()
	//fmt.Println(reflect.TypeOf(v))
	//fmt.Println(reflect.Indirect(v.FieldByName("a")))
	//a := new(test1)
	//a.name = "123"
	//fmt.Println(reflect.Indirect(reflect.ValueOf(a).Elem().FieldByName("name")))

	//var aaa = make(map[string]interface{})
	//var aa string
	//aa = "xxxxxxxx"
	//aaa["ww"] = &aa
	//fmt.Println(*aaa["ww"].(*string))
	//testttt(&aaa)
	//fmt.Println(*aaa["ww"].(*string))

	//j := `{"a": "sd", "b": 2.4, "c": true}`
	//aaa := make(map[string]interface{})
	//json.Unmarshal([]byte(j), &aaa)
	//fmt.Println(len(aaa))
	//a := aaa["b"]
	//fmt.Println(int(a.(float64)))

	//ctx, cancel := context.WithCancel(context.Background())
	//ch := make(chan bool)
	//go func() {
	//	select {
	//	case <-ctx.Done():
	//		fmt.Println("1111111")
	//		return
	//	case <-ch:
	//		fmt.Println("2222222")
	//		return
	//	}
	//}()
	//
	//time.Sleep(1 * time.Second)
	//cancel()
	//
	//time.Sleep(1 * time.Second)
	//ch <- true
	//a11 := &a{}
	//fmt.Println(reflect.ValueOf(a11).Type())
	//fmt.Println(reflect.ValueOf(a11).Elem().Type())

}

func testttt(aaa interface{}) {
	v := reflect.ValueOf(aaa).Elem()
	kind := v.Kind()
	if kind == reflect.Map {
		if value := v.MapIndex(reflect.ValueOf("ww")); value.IsValid() {
			fmt.Println(value.Elem().Type())
			value.Elem().Elem().Set(reflect.ValueOf("vv"))
		} else {
			fmt.Println("222")
		}
	}

}

type test1 struct {
	name string
}

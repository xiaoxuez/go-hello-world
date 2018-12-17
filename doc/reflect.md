## Reflect

最近写的代码中需要用到reflect包。

#### Api

+ ```
  TypeOf(i interface{}) Type
  ```

  返回为Type类型，为xx的类型，Type.Kind方法返回类型的种类(struct/ptr/)等

  e.g.

  ```
  a11 := &a{}
  reflect.TypeOf(a11)
  //*main.a          
  reflect.TypeOf(a11).Kind())
  //ptr
  ```

  ​

+ ```
  ValueOf(i interface{}) Value
  ```

  返回为Value类型，为xx的值。如xx为指针类型，Value.Elem可逆指针

  e.g.

  ```
  a11 := &a{}
  fmt.Println(reflect.ValueOf(a11).Type())
  //*main.a
  fmt.Println(reflect.ValueOf(a11).Elem().Type())
  //main.a
  ```

   

+ ```
  (v Value) Set(x Value)
  ```

  重置值的操作。但v必须要是可寻址变量。故原值要是个指针类型的，才能是可寻址的。

  关于可寻址的，如果某个值确实被存储在了计算机内存中，并且有一个内存地址可以代表这个值在内存中存储的起始位置，那么我们就说这个值以及代表它的变量是可寻址的。

  e.g.

  ```
  x := 2
  d := reflect.ValueOf(&x).Elem() 
  d.Set(reflect.ValueOf(4))
  fmt.Println(x) // "4"
  ```

  ​

+ ```
  reflect.New(argType.Type).Interface()
  ```

  ​

  可返回某Type的指针类型实例

#### e.g.

使用反射的包，常用到的就是json包。emmm...  没看过源码…这里就随便贴一段以太坊中的代码做示例吧...

是设置值的操作

```
func (arguments Arguments) unpackTuple(v interface{}, marshalledValues []interface{}) error {

	var (
		value = reflect.ValueOf(v).Elem()
		typ   = value.Type()
		kind  = value.Kind()
	)
	...
	for i, arg := range arguments.NonIndexed() {

		reflectValue := reflect.ValueOf(marshalledValues[i])

		switch kind {
		case reflect.Struct:
			if structField, ok := abi2struct[arg.Name]; ok {
				if err := set(value.FieldByName(structField), reflectValue, arg); err != nil {
					return err
				}
			}
		case reflect.Slice, reflect.Array:
			if value.Len() < i {
				return fmt.Errorf("abi: insufficient number of arguments for unpack, want %d, got %d", len(arguments), value.Len())
			}
			v := value.Index(i)
			if err := requireAssignable(v, reflectValue); err != nil {
				return err
			}

			if err := set(v.Elem(), reflectValue, arg); err != nil {
				return err
			}
		case reflect.Map:
			//这段是我自己加上去的...尬
			if dst := value.MapIndex(reflect.ValueOf(arg.Name)); dst.IsValid() {
				dst = dst.Elem()
				if err := set(dst, reflectValue, arg); err != nil {
					return err
				}
			}
		default:
			return fmt.Errorf("abi:[2] cannot unmarshal tuple in to %v", typ)
		}
	}
	return nil


}
```


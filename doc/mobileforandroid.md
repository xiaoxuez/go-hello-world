## Golang通过gomobile编译成android sdk



#### 编译

```
  1. $ go get golang.org/x/mobile/cmd/gomobile
     $ go get golang.org/x/mobile/cmd/gobind
  2. you must have ANDROID_HOME、ANDROID_NDK in env
  3. $ gomobile init  (--ndk $ANDROID_NDK)
  4. $ gomobile bind -target=android  ${go_src_project_path}
```



编译生成aar文件，将aar添加为android项目lib[方式](https://developer.android.com/studio/projects/android-library)



####  编程笔记

+ 编译的时候是指定文件夹，如，`gomobile bind -target=android github.com/blockchain/mobile`则只会将mobile下的*.go中的结构体/变量/方法等转换成可调用的api，例如，在go代码中引用别的包中的结构体，在转换好的sdk中是不可见的。
+ 方法参数和返回值，只支持**指针类型**和简单数据类型，不直接支持数组类型(可将数组封装在结构体内，通过size方法和get方法进行遍历和获取)
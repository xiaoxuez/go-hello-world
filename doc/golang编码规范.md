## Golang编码规范

参考自[官方文档](https://golang.org/doc/effective_go.html)。



#### formate

格式化。

官方提供有fmt工具以自动格式化代码。

建议使用官方提供的另外一个格式化工具goimports。

> Command goimports updates your Go import lines, adding missing ones and removing unreferenced ones.In addition to fixing imports, goimports also formats your code in the same style as gofmt so it can be used as a replacement for your editor's gofmt-on-save hook.

goimports具有fmt的功能，增加的功能是对import行的格式化。fmt默认对import格式化是按首字母排序，goimports是分组 + fmt的排序。分组排序结果如下。

```
import (
	"context"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)
```



#### Comment

注释。

```
// this is a comment
/*
this is a comment
*/
```

文件头部注释最好应包含作者和时间，(文件文档型说明也可包含其中)



#### 命名

+ 包名

  包名应由小写字母组成，不包含下划线或大写字母。如一个单词无法概全，应当以子包形式实现。如`encoding/base64`。避免使用`encoding_base64`和`encodingBase64`类似的写法

+ 采用驼峰式命名。如MixedCaps或mixedCaps(首字母大小写据导出条件而定)。避免使用下划线式命名。常量的命名也采用同样的方式。

+ 接口

  照惯例，一个方法接口由该方法name加上er后缀等类似的方式构造，如接口Reader、Writer的接口方法为Read、Write...



#### 文档

对于交互性强的代码，需配以注释或文档说明。



#### git

+ 为了避免git分叉。提交代码前先同步远端最新的代码，即`git pull`。如果出现冲突，应先保存本地修改 `git stash` 然后再次更新代码 `git pull`。然后恢复本地修改 `git stash pop` （`git stash list` 可以查看历史修改缓存。最后是提交代码 `git commit -am “xx”; git push `。
+ 避免提非文本文件（如二进制压缩包，word等） 即所有提交都应该是便于阅读的，或者说可以直接在浏览器上查看的
+ 避免提交大文件(超过过500k的就得慎重了)，内网还好，要是在网络条件差时会很难受的！因为一旦提交，git会永久保存改文件的历史记录，即使删除，除非删除整个项目。
+ 如果有修改，建议每天还是提一下代码，即使没做完，因为等你几天做完后，可能远端已经物是人非了，这也可以尽量减少冲突



#### test

能写测试的尽量加上测试，推荐个[写测试的工具](https://github.com/cweill/gotests)



#### 
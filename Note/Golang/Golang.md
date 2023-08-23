# golang简介

## go语言简介

Go（又称 Golang）是 Google 的 Robert Griesemer，Rob Pike 及 Ken Thompson 开发的一种静态强类型、编译型语言。Go 语言语法与 C 相近，但功能上有：内存安全，GC（垃圾回收），结构形态及 CSP-style 并发计算。

## go语言特点

1. 背靠大厂，google背书，可靠
2. 天生支持并发（最显著特点）
3. 语法简单，容易上手
4. 内置runtime，支持垃圾回收
5. 可直接编译成机器码，不依赖其他库
6. 丰富的标准库
7. 跨平台编译

## go语言的应用领域

1. 服务器编程
2. 开发云平台
3. 区块链
4. 分布式系统
5. 网络编程

# go语言开发环境搭建

## windows平台

**<font size = 5px>下载安装配置环境变量</font>**

下载地址：<font color = red>` https://golang.google.cn/dl/`</font>，这里提供了不同平台的go版本，根据自己的平台选择下载。

这里注意，安装路径选择一个比较好找的路径，例如：<font color =red>` c:/go`</font>，其他安装都选择“下一步”即可。

安装完成后，把<font color =red>` C:\go\bin`</font>目录添加到环境变量，这里就可以使用go了，在命令行输入<font color =red>` go version`</font>查看版本，输出结果显示如下所示。

```go
go version go1.16.6 windows/amd64
```

**<font size = 5px>配置go环境</font>**

```go
$env:GO111MODULE = "on"
$env:GOPROXY = "http://goproxy.cn"
```

> 使用go mod 管理库，需要科学上网

**<font size = 5px>安装配置git</font>**

```go
https://www.git-scm.com/download
```

配置环境变量，命令行输入git测试

**<font size = 5px>gotoot和gopath</font>**

goroot就是go安装的根目录，gopath就是go项目所在的路径，高版本go项目已经不再依赖gopath来管理项目，使用go mod来管理项目。

## Linux平台

如果不是要在Linux平台敲go代码就不需要在Linux平台安装Go，我们开发机上写好的go代码只需要跨平台编译（详见文章末尾的跨平台编译）好之后就可以拷贝到Linux服务器上运行了，这也是go程序跨平台易部署的优势。

我们在版本选择页面选择并下载好<font color =red>` go1.14.1.linux-amd64.tar.gz`</font>文件：

```go
wgbt https://dl.google.com/go/go1/14/1/linux-amd64.tar.gz
```

将下载好的文件解压<font color = red>` /usr/local`</font>目录下：

```go
tar -zxvf go1.14.1.linux-amd64.tar.gz -C /usr/local  # 解压
```

如果提示没有权限，加上<font color =red>` sudo`</font>以root用户身份再运行。执行完就可以在<font color =red>` /usr/local`</font>下看到<font color =red>` go`</font>目录了。

配置环境变量：Linux下有两个文件可以配置环境变量，其中<font color =red>` /etc/profile`</font>是对所有用户生效的；<font color =red>` $HOME/.profile`</font>是对当前用户生效的，根据自己情况自行选择一个文件打开，添加如下两行代码，保存退出。

```go
exprot GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
```

修改<font color =red>` /etc/profile`</font>后要重启生效，修改<font color =red>` $HOME/.profile`</font>后使用source命令加载<font color =red>` $HOME/.profile`</font>文件即可生效。检查：

```go
~ go version
go version go1.16.6 linux/amd64
```

# 开发工具简介

golang的开发工具有很多，例如：

1. vim
2. sublime
3. atom
4. LiteIDE
5. eclipse
6. goland
7. vscode

## 使用goland开发go应用

goland是一款付费的、功能强大的golang集成开发 环境，是Jetbrain公司的产品。下载地址：<font color =red>`https://www.jetbrains.com/go/` </font>goland非常智能，几乎不需要配置，安装即用。

## 使用vscode开发go应用

1. 下载安装vscode <font color =red>`https://code.visualstudio.com/`</font>
2. 下载安装插件，参考文档：<font color =red>`https://marketplace.visualstudio.com/items?itemName=golang.go`</font>
3. 安装code runner运行脚本

# go常用命令

## 查看可用命令

直接在终端输入<font color =red>` go help`</font> 即可显示所有的 go 命令以及相应命令功能简介，主要有下面这些：

+ build：编译包和依赖
+ clean：移除对象文件
+ doc：显示包或者符号的文档
+ env：打印go的环境信息
+ bug：启动错误报告
+ fix：运行go tool fix
+ fmt：运行gofmt进行格式化
+ generate：从processing source生成 go 文件
+ get：下载并安转包和依赖
+ install：编译并安转包和依赖
+ list：列出包
+ run：编译并运行go程序
+ test：运行测试
+ tool：运行go提供的工具
+ version：显示go的版本
+ vet：运行go tool vet

## 参考文档

<font color =red>` https://golang.org/doc/cmd`</font>

# golang开发 vscode快捷键

## 快捷键位置

```go
File→Preferences→Keyboard Shortcuts
```

## 常用快捷键

```go
1. 行注释 ctrl + /
2. 块注释 shift+alt+a（可以修改为ctrl+shift+/）
3. ctrl+shift+k 删除行
4. ctrl+e 查找文件
5. ctrl+shift+p 打开设置命令行
```

## 修改快捷键

1. 打开快捷键<font color = red>` File->Preference->Keyboard Shortcuts`</font>
2. 双击要修改的快捷键
3. 键盘输入想要使用的快捷键

## 快速生成代码片段

```go
pkgm main包+main主函数
ff fmt.Printf("", var)
for for i := 0;i < count; i++ {}
forr for _, v := range v {}
fmain func main(){}
a.print! fmt.Printf("a: %v\n", a)
```

# 如何编写golang代码

## 代码组织

go应用使用**包**和**模块**来组织代码，包对应到文件系统就是文件夹，模块就是** .go **的go源文件。一个包中会有多个模块，或者多个子包。

## go项目管理工具

早期的go项目使用gopath来管理项目，不方便而且容易出错，从 golang 1.11 开始使用gomod管理项目，当然还有第三方模块例如govender，我们给大家介绍gomod的使用

## 实现步骤

1. 创建项目
2. 初始化项目
3. 创建包
4. 创建模块
5. 相互调用

# golang标识符、关键字、命名规则

## 标识符

标识符的英文是<font color = orange>` identifier`</font>，通俗的讲，就是给变量、常量、函数、方法、结构体、数组、切片、接口起名字。

## 标识符的组成

1. 标识符有数字、字母和下划线（<font color = orange>` _`</font>）组成。
2. 只能以字母和下划线（<font color = orange>` _`</font>）开头。
3. 标识符区分大小写。

## 举例说明标识符的命名

### **正确的命名**

```go
package main

import "fmt"

func main() {
    var name string
    var age int
    var _sys int
}
```

### **错误的标识符**

```go
package main

import "fmt"

func main() {
    var 1name string
    var &age int
    var !sys int
}
```

## go语言关键字

| break    | default     | func   | interface | select |
| -------- | ----------- | ------ | --------- | ------ |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

除了以上介绍的关键字，Go 语言还有36个预定义标识符其中包括了基本数据类型的名称和一些基本的内置函数，见下表：

| append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
| ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
| copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
| int32  | int64   | iota    | len     | make   | new     | nil       | panic      | uint64  |
| print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |

## go语言命名规范

### Go是一门区分大小写的语言

命名规则涉及变量、常量、全局函数、结构、接口、方法等的命名。Go语言从语法层面进行了以下限定：任何需要对外暴露的名字必须以大写字母开头，不需要对外暴露的则应该以小写字母开头。

当命名（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：<font color =orange>` GetUserName`</font>，那么使用这种标识符的对象就**可以被外包的代码所使用**（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 <font color =orange>` public`</font>）；**命名如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的**（像面向对象语言中的 <font color =orange>` private`</font>）

### 包名称

保持<font color =orange>` package`</font> 的名字和目录保持一致，尽量采取有意义的包名，简短，有意义，尽量和标准库不要冲突。包名应该为**小写**单词，不要使用下划线或者混合大小写。

```go
package dao
package service
```

### 文件命名

尽量采取有意义的文件名，简短，有意义，应该为**小写**单词，使用**下划线**分隔各个单词。

```go
customer_dao.go
```

### 结构体命名

采用**驼峰命名法**，首字母根据访问控制大写或者小写

<font color =orange>` struct`</font>申明和初始化格式采用多行，例如下面：

```go
type CustomerOrder struct {
    Name string
    Address string
}
order := CustomerOrder{"tom","北京"}
```

### 接口命名

命名规则基本和上面的结构体类型相似

单个函数的结构名以“er”作为后缀，例如<font color =orange>` Reader`</font>，<font color =orange>` Writer`</font>。

``` go
type Reader interface {
    Read(p []byte) (n int,err error)
}
```

### 变量命名

和结构体类似，变量名称一般遵循**驼峰法**，首字母根据访问控制原则大写或者小写，但遇到特有名词时，需要遵循以下规则：

如果变量为私有，且特有名词为首个单词，则使用小写，如 APPService 若变量类型为 bool 类型，则名称应为 Has，is，Can 或 Allow 开头

```go
var isExist bool
var hasConflict bool
var canManage bool
var allowGitHook bool
```

### 常量命名

常量均需要使用全部**大写**字母组成，并使用**下划线分词**

```go
const APP_URL = "https://www.duoke360.com"
```

如果是枚举类型的常量，需要先创建相关类型：

```go
type Scheme string

const {
    HTTP  Scheme = "http"
    HTTPS Scheme = "https"
}
```

### 错误处理

错误处理的原则就是不能丢弃任何有返回err的调用，不要使用_丢弃，必须全部处理。接收到错误，要么返回err，或者使用log记录下来尽早return：一旦有错误发生，马上返回，尽量不要怕使用panic，除非你知道你在做什么，错误描述如果是英文必须为小写，不需要标点结尾，采用独立的错误流进行处理

```go
// 错误写法
if err != nil {
    // 错误处理
} else {
    // 正常代码
}

// 正确写法
if err != nil {
    // 错误处理
    return // 或者继续
}
// 正常代码
```

### 单元测试

单元测试文件名命名规范为<font color = orange>` example_test.go`</font> 测试用例的函数名称必须以<font color = orange>` Test`</font> 开头，例如<font color = orange>` TestExample`</font>每个重要函数都要首先编写测试用例，测试用例和正规代码一起提交方便进行回归测试。

# golang变量

变量是计算机语言能**储存**计算结果或能表示值得抽象概念。不同的变量保存的数据类型可能会不一样。

## 声明变量

Go语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且Go语言的变量声明后必须使用。

**声明变量的语法**

``` go
var identifier type
```

<font color = orange>` var`</font>：声明变量关键字

<font color = orange>` identifier`</font>：变量名称

<font color = orange>` type`</font>：变量类型

**例如**

```go
package main

import "fmt"

func main(){
    var name string
    var age int
    var ok bool
}
```

### 批量声明

使用一个<font color = orange>` var`</font>关键字，把一下变量写到一个<font color = orange>` ()`</font> 里

```go
package main
func main(){
    var(
    	name string
        age int 
        ok bool
    )
}
```

## 变量的初始化

Go语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，例如：整型和浮点型变量的默认值为<font color = orange>` 0`</font>。字符串变量的默认值为空字符串<font color = orange>` ""`</font>。布尔型变量默认为<font color = orange>` false`</font>。切片、函数、指针变量的默认为<font color = orange>` nil`</font>。

**变量初始化语法**

```go
var 变量名 类型 = 表达式
```

**例如**

```go
package main
func main(){
    var name string = "老郭"
    var site string = "www.duoke360.com"
    var age int = 30
}
```

**类型推导**

我们在声明变量时，可以根据初始化值进行类型推导，从而省略类型。

```go
package main
func main(){
    var name = "老郭"
    var site = "www.duoke360.com"
    var age = 30
}
```

**初始化多个变量**

可以一次性初始化多个变量，中间用逗号分隔。

```go
package main
func main(){
    var name,site,age = "老郭","www.duoke360.com",30
}
```

## 短变量声明

在函数内部，可以使用<font color = orange>` :=`</font> 运算符对变量进行声明和初始化。

```go
package main

func main(){
    name := "老郭"
    site := "www.duoke360.com"
    age := 20
}
```

> 注意：这种方法只适合在函数内部，函数外面不能使用。

## 匿名变量

如果我们接收多个变量，有一些变量使用不到，可以使用下划线<font color = orange>` _`</font>表示变量名称，这种变量叫做匿名变量。例如：

```go
package main

import "fmt"

func getNameAndAge()(string,int){
    return "老郭",30
}

func main(){
    name, _ := getNameAndAge()
    fmt.Printf("name: %v\n", name)
}
```

# go语言常量

常量，就是在程序**编译阶段**就确定下来的值，而程序在**运行时**则无法改变该值。在Go程序中，常量可以是数值类型（包括整型、浮点型和复数类型）、布尔类型、字符串类型等。

## 定义常量的语法

定义一个常量使用const关键字，语法格式如下：

```go
const constantName [type] = value
```

<font color = orange>` const`</font>：定义常量关键字

<font color = orange>` constantName`</font>：常量名称

<font color = orange>` type`</font>：常量类型

<font color = orange>` value`</font>：常量的值

**实例**

``` go
package main

func main(){
    const PI float64 = 3.14
    const PI2 = 3.1415 // 可以省略类型
    
    const(
    	width = 100
        height = 200
    )
    
    const i,j = 1,2 // 多重赋值
    const a,b,c = 1,2,"foo"
}
```

**<font color = orange>` const`</font> 同时声明多个常量时，如果省略了值则表示和上面一行的值相同**

```go
package main

import "fmt"

func main(){
    const(
    	a1 = 100
        a2
        a3
    )
    
    fmt.Printf("a1: %v\n", a1)
    fmt.Printf("a2: %v\n", a2)
    fmt.Printf("a3: %v\n", a3)
}
```

运行结果

```go
a1: 100
a2: 100
a3: 100
```

## <font color = orange>` iota`</font>

iota 比较特殊，可以被认定为是一个可被编译器修改的常量，它默认开始值是<font color = orange>` 0`</font>，每调用一次加<font color = orange>` 1`</font>。遇到<font color = orange>` const`</font> 关键字时被重置为<font color = orange>` 0`</font>。

**实例**

``` go
package main

import "fmt"

func main(){
    const(
    	a1 = iota
        a2 = iota
        a3 = iota
    )
    
    fmt.Printf("a1: %v\n", a1)
    fmt.Printf("a2: %v\n", a2)
    fmt.Printf("a3: %v\n", a3)
}
```

运行结果

```go
a1: 0
a2: 1
a3: 2
```

**使用<font color = orange>` _`</font>跳过某些值**

```go
package main

import "fmt"

func main(){
    const(
    	a1 = iota
        -
        a2 = iota
    )
    
    fmt.Printf("a1: %v\n", a1)
    fmt.Printf("a2: %v\n", a2)
}
```

运行结果

```go
a1: 0
a2: 2
```

**<font color = orange>` iota`</font>声明中间插队**

```go
package main

import "fmt"

func main(){
    const(
    	a1 = iota
        a2 = 100
        a3 = iota
    )
    
    fmt.Printf("a1: %v\n", a1)
    fmt.Printf("a2: %v\n", a2)
    fmt.Printf("a3: %v\n", a3)
}
```

运行结果

```go
a1: 0
a2: 100
a3: 2
```

# go语言数据类型

在 Go 编程语言中，数据类型用于声明函数和变量。

数据类型的出现是为了把数据分为所需**内存大小**不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存。

Go 语言按类别有以下几种数据类型：

| 序号 | 类型和描述                                                   |
| ---- | ------------------------------------------------------------ |
| 1    | **布尔型**：布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。 |
| 2    | **数字类型**：整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位运算采用补码。 |
| 3    | **字符串类型**：字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是有单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。 |
| 4    | **派生类型**：包括：(a) 指针类型（Pointer）(b) 数组类型 (c) 结构化类型（struct）(d) Channel 类型 (e) 函数类型 (f) 切片类型 (g) 接口类型（interface）(h) Map 类型 |

## 数字类型

Go 也有基于架构的类型，例如：int、uint 和 uintptr。

| 序号 | 类型和描述                                                   |
| ---- | ------------------------------------------------------------ |
| 1    | **uint8**：无符号 8 位整型（0 到 255）                       |
| 2    | **uint16**：无符号 16 位整型（0 到 65535）                   |
| 3    | **uint32**：无符号 32 位整型（0 到 4294967295）              |
| 4    | **uint64**：无符号 64 位整型（0 到 184467440737）            |
| 5    | **int8**：有符号 8 位整型（-128 到 127）                     |
| 6    | **int16**：有符号 16 位整型（-32768 到 32767）               |
| 7    | **int32**：有符号 32 位整型（-2147483647 到 214748367）      |
| 8    | **int64**：有符号 64 位整型（-9223372036854775808 到 9223372036854775807） |

## 浮点型

| 序号 | 类型和描述                         |
| ---- | ---------------------------------- |
| 1    | **float32**：IEEE-754 32位浮点型数 |
| 2    | **float64**：IEEE-754 64位浮点型数 |
| 3    | **complex64**：32位实数和虚数      |
| 4    | **complex128**：64位实数和虚数     |

## 其他数字类型

一下列出了其他更多的数字类型：

| 序号 | 类型和描述                               |
| ---- | ---------------------------------------- |
| 1    | **byte** 类似 uint8                      |
| 2    | **rune** 类似 int32                      |
| 3    | **uint** 32 或 64 位                     |
| 4    | **int** 与 uint 一样大小                 |
| 5    | **uintptr** 无符号整型，用于存放一个指针 |

# go语言布尔类型

go语言中的布尔类型有两个常量值：<font color = orange>` true`</font>和<font color = orange>` false`</font>。布尔类型经常用在**条件判断**语句，或者**循环语句**。也可以用在**逻辑表达式**中。

## 布尔类型

```go
package main

func main(){
    var b1 bool = true
    var b2 bool = false
    var b3 = true
    var b4 = false
    
    b5 := true
    b6 := false
    
    fmt.Printf("b1: %v\n" ,b1)
    fmt.Printf("b2: %v\n" ,b2)
    fmt.Printf("b3: %v\n" ,b3)
    fmt.Printf("b4: %v\n" ,b4)
    fmt.Printf("b5: %v\n" ,b5)
    fmt.Printf("b6: %v\n" ,b6)
}
```

运行结果

```go
b1: true
b2: false
b3: true
b4: false
b5: true
b6: false
```

## 用在条件判断中

``` go
package main

import "fmt"

func main(){
    age := 18
    ok := age >= 18
    if ok {
        fmt.Printf("你已经成年")
    } else {
        fmt.Printf("你还未成年")
    }
}
```

运行结果

```go
你已经成年
```

## 用在循环语句中

``` go
package main

import "fmt"

func main(){
    count := 10
    for i := 0; i < count; i++{
        fmt.Printf("i: %v\n", i)
    }
}
```

## 用在逻辑表达式中

```go
package main

import "fmt"

func main() {
    age := 18
    gender := "男"
    
    if age >= 18 && gender == "男"{
        fmt.Println("你是成年男子")
    } 
}
```

> 注意：不能使用 <font color = orange>` 0`</font> 和非 <font color = orange>` 0`</font> 表示真假

```go
package main

func main(){
    i := 1
    if i {
        // 编译错误
    }
}
```

# go语言数组类型

Go 语言支持整型和浮点型数据，并且原生支持复数，其中位的运算采用补码。

Go 也有基于架构的类型，例如：<font color = orange>` int`</font>、<font color = orange>` uint`</font> 和 <font color = orange>` uintptr`</font>。

这些类型的长度都是根据运行程序所在的操作系统类型所决定的：

+ <font color = orange>` int`</font> 和<font color = orange>` uint`</font> 在 32 位操作系统上，它们均使用 32 位（4个字节），在 64 位操作系统上，它们均使用 64 位（8个字节）
+ <font color = orange>` uintptr`</font> 的长度被设定为足够存放一个指针即可。

Go 语言中没有 <font color = orange>` float`</font> 类型。（Go语言只有 <font color = orange>` float32`</font> 和 <font color = orange>` float64`</font>）没有double类型。

与操作系统架构无关的类型都有固定的大小，并在类型的名称中就可以看出来：

整数：

+ int8（-128 -> 127）
+ int16（-32768 -> 32767）
+ int32（-2147483647 -> 214748367）
+ int64（-9223372036854775808 -> 9223372036854775807）

无符号整数：

+ uint8（0 -> 255）
+ uint16（0 -> 65535）
+ uint32（0 -> 4294967295）
+ uint64（0 -> 18446744073709551615）

浮点型（IEEE-754 标准）：

+ float32（+- 1e-45 -> +- 3.4 * 1e38）
+ float64（+- 5 * 1e-324 -> 107 * 1e308）

int 型是计算最快的一种类型。

整型的零值为0，浮点型的零值为0.0。

## 实例

下面实例演示了，各个数据类型的长度和取值范围

```go
package main

import (
	"fmt"
    "math"
    "unsafe"
)

func main(){
    var i8 int8
    var i16 int16
    var i32 int32
    var i64 int64
    var ui8 uint8
    var ui16 uint16
    var ui32 uint32
    var ui64 uint64
    
    fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
    fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
    fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
    fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)
    
    fmt.Printf("%T %dB %v~%v\n", ui8, unsafe.Sizeof(ui8), 0, math.MaxUint8)
    fmt.Printf("%T %dB %v~%v\n", ui16, unsafe.Sizeof(ui16), 0, math.MaxUint16)
    fmt.Printf("%T %dB %v~%v\n", ui32, unsafe.Sizeof(ui32), 0, math.MaxUint32)
    fmt.Printf("%T %dB %v~%v\n", ui64, unsafe.Sizeof(ui64), 0, uint64(math.MaxUint64))
    
    var f32 float32
    var f64 float64
    
    fmt.Printf("%T %dB %v~%v\n", f32, unsafe.Sizeof(f32), -math.MaxFloat32, math.MaxFloat32)
    fmt.Printf("%T %dB %v~%v\n", f64, unsafe.Sizeof(f64), -math.MaxFloat64, math.MaxFloat64)
    
    var ui uint
    ui = uint(math.MaxUint64) // 再+1会导致overflows错误
    fmt.Printf("%T %dB %v~%v\n", ui, unsafe.Sizeof(ui), 0, ui)
    
    var imax, imin int
    imax = int(math.MaxInt64) //再+1会导致overflow错误
    imin = int(math.MinInt64) //再-1会导致overflow错误
    
    fmt.Printf("%T %dB %v~%v\n", imax, unsafe.Sizeof(imax), imin, imax)
}
```

运行结果

```go
int8 1B -128~127
int16 2B -32768~32767
int32 4B -2147483648~2147483647
int64 8B -9223372036854775808~9223372036854775807
uint8 1B 0~255
uint16 2B 0~65535
uint32 4B 0~4294967295
uint64 8B 0~18446744073709551615
float32 4B -3.4028234663852886e+38~3.4028234663852886e+38
float64 8B -1.7976931348623157e+308~1.7976931348623157e+308
uint 8B 0~18446744073709551615
int 8B -9223372036854775808~9223372036854775807
```

## 以二进制、八进制或十六进制浮点数的格式定义数字

```go
package main

import "fmt"

func main(){
    // 十进制
    var a int = 10
    fmt.Printf("%d \n",a) // 10
    fmt.Printf("%b \n",a) // 1010  占位符%b表示二进制
    
    // 八进制  以0开头
    var b int = 077
    fmt.Printf("%o \n",b) // 77
    
    // 十六进制  以0x开头
    var c int = 0xff
    fmt.Printf("%x \n",c) // ff
    fmt.Printf("%X \n",c) // FF
}
```

运行结果

```go
10
1010
77
ff
FF
```

## 浮点型

Go语言支持两种浮点型数：<font color = orange>` float32`</font> 和 <font color = orange>` float64`</font>。这两种浮点型数据格式遵循<font color = orange>` IEEE 754`</font> 标准：<font color = orange>` float32`</font> 的浮点数的最大范围约为 <font color = orange>` 3.4e38`</font>，可以使用常量定义：<font color = orange>` math.MaxFloat32`</font>。<font color = orange>` float64`</font> 的浮点数的最大范围约为<font color = orange>` 1.8e308`</font>，可以使用一个常量定义：<font color = orange>` math.MaxFloat64`</font>。

打印浮点数时，可以使用<font color = orange>` fmt`</font>包配合动词<font color = orange>` %f`</font>，代码如下：

```go
package main
import (
	"fmt"
    "math"
)
func main(){
    fmt.Printf("%f\n", math.Pi)
    fmt.Printf("%.2f\n", math.Pi)  // 有两位小数点
}
```

 ## 复数

complex64和complex128

```go
var c1 complex64
c1 = 1 + 2i
var c2 complex128
c2 = 2 + 3i
fmt.Printf(c1)
fmt.Printf(c2)
```

复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。

# golang字符串

一个Go语言字符串是一个任意**字节的常量序列**。

## go语言字符串字面量

在Go语言中，字符串字面量使用双引号<font color = orange>` ""`</font>或者反引号<font color = orange>` </font>来创建。双引号用来创建可解析的字符串，支持转义，但不能用来引用多行；反引号用来创建原生的字符串字面量，可能由多行组成，但不支持转义，并且可以包含除了反引号外其他所有字符。双引号创建可解析字符串应用最广泛，反引号用来创建原生的字符串则多用于书写多行消息，HTML以及正则表达式。

**实例**

``` go
package main

func main(){
    var str1 string = "hello world"
    var html string = `
    	<html>
    		<head><title>hello golang</title>
    	</html>
    `
    fmt.Printf("str1: %v\n", str1)
    fmt.Printf("html: %v\n", html)
}
```

运行结果

``` go
str1: hello world
html: 
    	<html>
    		<head><title>hello golang</title>
    	</html>
```

## go语言字符串连接

**使用加号**

虽然Go语言中的字符串是不可变的，但字符串支持<font color = orange>` +`</font>级联操作和<font color = orange>` +=`</font>追加操作，例如：

``` go
package main

import "fmt"

func main(){
    name := "tom"
    age := "20"
    msg := name + " " + age
    fmt.Printf("msg: %v\n", msg)
    fmt.Printf("--------------")
    msg = ""
    msg += name
    msg += " "
    msg += age
    fmt.Printf("msg: %v\n", msg)
}
```

> golang 里面的字符串都是不可变的，每次运算都会产生一个新的字符串，所以会产生很多临时的无用的字符串，不仅没有用，还会给 gc带来额外的负担，所以性能比较差

**使用<font color = orange>` fmt.Sprintf()`</font>函数**

``` go
package main

import "fmt"

func main(){
    name := "tom"
    age := "20"
    msg := fmt.Sprintf("%s,%s",name,age)
    fmt.Printf("msg: %v/n",msg)
}
```

运行结果

``` go
msg: tom,20
```

> 内部使用<font color = orange>` []byte`</font>实现，不像直接运算符这种会产生很多临时字符串，但是内部的逻辑比较复杂，有很多额外的判断，还用到了<font color = orange>` interface`</font>，所以性能也不是很好

**<font color = orange>` strings.Join()`</font>**

```go
package main

import(
	"fmt"
    "strings"
)

func main(){
    name := "tom"
    age := "20"
    msg := strings.Join([]string{name,age},",")
    fmt.Printf("msg: %v\n",msg)
}
```

运行结果

```go
msg: tom,20
```

> join会先根据字符串数据的内容，计算出一个拼接之后的长度，然后申请对应大小的内存，一个一个字符串填入，在已有一个数组的情况下，这种效率会很高，但是本来没有，去构造这个数据的代价也不小

**<font color = orange>` buffer.WriteString()`</font>**

```go
package main

import(
	"bytes"
    "fmt"
)

func main(){
    var buffer bytes.Buffer
    buffer.WriteString("tom")
    buffer.WriteString(",")
    buffer.WriteString("20")
    fmt.Printf("buffer.String(): %v\n", buffer.String())
}
```

> 这个比较理想，可以当成可变字符使用，对内存的增长也有优化，如果能预估字符串的长度，还可以用<font color = orange>` buffer.Grow()`</font>接口来设置 capacity（容量）

## go语言字符串转义字符

Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

| 转义符                            | 含义                               |
| --------------------------------- | ---------------------------------- |
| <font color = orange>` \r`</font> | 回车符（返回行首）                 |
| <font color = orange>` \n`</font> | 换行符（直接跳到下一行的同列位置） |
| <font color = orange>` \t`</font> | 制表符                             |
| <font color = orange>` \'`</font> | 单引号                             |
| <font color = orange>` \"`</font> | 双引号                             |
| <font color = orange>` \\`</font> | 反斜杠                             |

**实例**

``` go
package main

import(
	"fmt"
)

func main(){
    fmt.Print("hello\tworld\n")
    fmt.Print("\"c:\\test\\\")
}
```

运行结果

``` go
hello	world
"c:\test\"
```

## go语言字符串切片操作

``` go
package main

import (
	"fmt"
)

func main(){
    str := "hello world"
    n := 3
    m := 5
    fmt.Println(str[n])	   // 获取字符串索引位置为n的原始字节
    fmt.Println(str[n:m])  // 截取得字符串索引位置为 n 到 m-1 的字符串
    fmt.Println(str[n:])   // 截取得字符串索引位置为 n 到 len(s)-1 的字符串
    fmt.Println(str[:m])   // 截取得字符串索引位置为 0 到 m-1 的字符串	
}
```

运行结果

``` go
108
lo
lo world
hello
```

## go语言字符串常用方法

| 方法                                | 介绍           |
| ----------------------------------- | -------------- |
| len(str)                            | 求长度         |
| +或fmt.Sprintf                      | 拼接字符串     |
| strings.Split                       | 分割           |
| strings.Contains                    | 判断是否包含   |
| strings.HasPrefix,strings.HasSuffix | 前缀/后缀判断  |
| strings.Index(),strings.LastIndex() | 子串出现的位置 |
| strings.Join(a[]string,sep string)  | join操作       |

**实例**

```go
package main

import(
	"fmt"
    "strings"
)

func main(){
    s := "hello world!"
    fmt.Printf("len(s): %v/n",len(s))
    fmt.Printf("strings.Split(s,\"hello\"): %v\n",strings.Split(s," "))
    fmt.Printf("strings.Contains(s,\"hello\"): %v\n",strings.Contains(s,"hello"))
    fmt.Printf("strings.HasPrefix(s,\"hello\"): %v\n",strings.HasPrefix(s,"hello"))
    fmt.Printf("strings.HasSuffix(s,\"hello\"): %v\n",strings.HasSuffix(s,"world!"))
    fmt.Printf("strings.Index(s.\"l\"): %v\n",strings.Index(s,"l"))
    fmt.Printf("strings.LastIndex(s.\"l\"): %v\n",strings.LastIndex(s,"l"))
}
```

运行结果

``` go
len(s): 12/nstrings.Split(s,"hello"): [hello world!]
strings.Contains(s,"hello"): true
strings.HasPrefix(s,"hello"): true
strings.HasSuffix(s,"hello"): true
strings.Index(s."l"): 2
strings.LastIndex(s."l"): 9
```

# golang格式化输出

下面实例使用到的结构体

``` go
type Website struct {
    Name string
}
// 定义结构体变量
var site = Website{Name:"duoke360"}
```

## 占位符

**普通占位符**

| 占位符 | 说明                                                         | 举例                                 | 输出                          |
| ------ | ------------------------------------------------------------ | ------------------------------------ | ----------------------------- |
| %v     | 相对应的默认格式，在打印结构体时，”加号“标记（%+v）会添加字段名 | Printf("%v",site),Printf("%+v",site) | {duoke360},{Name:duoke360}    |
| %#v    | 相应值的Go语法表示                                           | Printf("%#v",site)                   | main.Website{Name:"duoke360"} |
| %T     | 相应值的类型的Go语法表示                                     | Printf("%T",site)                    | main.Website                  |
| %%     | 字面上的百分号，并非值的占位符                               | Printf("%%")                         | %                             |

**布尔占位符**

| 占位符 | 说明                 | 举例              | 输出 |
| ------ | -------------------- | ----------------- | ---- |
| %t     | 单词 true 或 false。 | Printf("%t",true) | true |

**整数占位符**

| 占位符 | 说明                                       | 举例                | 输出   |
| ------ | ------------------------------------------ | ------------------- | ------ |
| %b     | 二进制表示                                 | Printf("%b",5)      | 101    |
| %c     | 相对应Unicode码点所表示的字符              | Printf("%c",0x4E2D) | 中     |
| %d     | 十进制表示                                 | Printf("%d",0x12)   | 18     |
| %o     | 八进制表示                                 | Printf("%o",10)     | 12     |
| %q     | 单引号围绕的字符字面值，由Go语法安全地转义 | Printf("%q",0x4E2D) | '中'   |
| %x     | 十六进制表示，字面形式为小写 a-f           | Printf("%x",13)     | d      |
| %X     | 十六进制表示，字面形式为大写 A-f           | Printf("%X",13)     | D      |
| %U     | Unicode格式：U+1234，等用于“U+%04X”        | Printf("%U",0x4E2D) | U+4E2D |

**浮点数和复数的组成部分（实部和虚部）**

| 占位符 | 说明                                                         | 举例                  | 输出         |
| ------ | ------------------------------------------------------------ | --------------------- | ------------ |
| %b     | 无小数部分，指数为二的幂的科学计数法，与 strconv.FormatFloat 的 'b' 转换格式一致。例如 -123456p-78 |                       |              |
| %e     | 科学计数法，例如 -1234.456e+78                               | Printf("%e",10.2)     | 1.020000e+01 |
| %E     | 科学计数法，例如 -1234.456e+78                               | Printf("%E",10.2)     | 1.020000E+01 |
| %f     | 有小数点而无指数，例如123.456                                | Printf("%f",10.2)     | 10.200000    |
| %g     | 根据情况选择 %e 或 %f 以产生更紧凑的（无末尾的0）输出        | Printf("%g",10.20)    | 10.2         |
| %G     | 根据情况选择 %E 或 %f 以产生更紧凑的（无末尾的0）输出        | Printf("%G",10.20+2i) | 10.2+2i      |

**字符串与字节切片**

| 占位符 | 说明                                   | 举例                          | 输出         |
| ------ | -------------------------------------- | ----------------------------- | ------------ |
| %s     | 输出字符串表示（string类型或[]byte）   | Printf("%s",[]byte("多课网")) | 多课网       |
| %q     | 双引号围绕的字符串，由Go语法安全地转义 | Printf("%q","多课网")         | "多课网"     |
| %x     | 十六进制，小写字母，每字节两个字符     | Printf("%x","golang)          | 676f6c616e67 |
| %X     | 十六进制，大写字母，每字节两个字符     | Printf("%X","golang)          | 646F6C616E67 |

**指针**

| 占位符 | 说明                  | 举例               | 输出     |
| ------ | --------------------- | ------------------ | -------- |
| %p     | 十六进制表示，前缀 0x | Printf("%p",&site) | 0x4f57f0 |

# golang运算符

Go 语言内置运算符有：

1. 算术运算符
2. 关系运算符
3. 逻辑运算符
4. 位运算符
5. 赋值运算符

## 算术运算符

| 运算符 | 描述 |
| ------ | ---- |
| +      | 相加 |
| -      | 相减 |
| *      | 相乘 |
| /      | 相除 |
| %      | 求余 |

**注意：**<font color = orange>` ++`</font>（自增）和 <font color = orange>` --`</font>（自减）在Go语言中是单独的语句，并不是运算符。

**实例**

``` go
package main

import "fmt"

func main(){
    a := 100
    b := 10
    
    fmt.Printf("(a + b): %v\n", (a + b))
    fmt.Printf("(a - b): %v\n", (a - b))
    fmt.Printf("(a * b): %v\n", (a * b))
    fmt.Printf("(a / b): %v\n", (a / b))
    fmt.Printf("(a %% b): %v\n", (a % b))
    
    a++
    fmt.Printf("a: %v\n", a)
    b--
    fmt.Printf("b: %v\n", b)
}
```

## 关系运算符

| 运算符 | 描述                                                         |
| ------ | ------------------------------------------------------------ |
| ==     | 检查两个值是否相等，如果相等返回 True 否则返回 False         |
| !=     | 检查两个值是否不相等，如果不相等返回 True 否则返回 False     |
| >      | 检查左边值是否大于右边值，如果是返回 True 否则返回 False     |
| >=     | 检查左边值是否大于右边值，如果是返回 True 否则返回 False     |
| <      | 检查左边值是否小于右边值，如果是返回 True 否则返回 False     |
| <=     | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False |

**实例**

``` go
package main

import "fmt"

func main(){
    a := 1
    b := 2
    
    fmt.Printf("a > b: %v\n",(a > b))
    fmt.Printf("a < b: %v\n",(a < b))
    fmt.Printf("a >= b: %v\n",(a >= b))
    fmt.Printf("a <= b: %v\n",(a <= b))
    fmt.Printf("a == b: %v\n",(a == b))
    fmt.Printf("a != b: %v\n",(a != b))
}
```

## 逻辑运算符

| 运算符 | 描述                                                         |
| ------ | ------------------------------------------------------------ |
| &&     | 逻辑 AND 运算符。如果两边的操作数都是 True，则为 True，否则为 False。 |
| \|\|   | 逻辑 OR 运算符。如果两边的操作数有一个为 True ，则为 True，否则为 False。 |
| ！     | 逻辑 NOT 运算符。如果条件为 True，则为 False，否则为 True。  |

**实例**

``` go
package main

import "fmt"

func main(){
    a := true
    b := false
    
    fmt.Printf("(a && b): %v/n",(a && b))
    fmt.Printf("(a || b): %v/n",(a || b))
    fmt.Printf("(!a): %v/n",(!a))
    fmt.Printf("(!b): %v/n",(!b))
}
```

## 位运算符

位运算符对整数在内存中的二进制位进行操作。

| 运算符 | 描述                                                         |
| ------ | ------------------------------------------------------------ |
| &      | 参与运算的两数各对应的二进位相与。（两位均为1才为1）         |
| \|     | 参与运算的两数各对应的二进位相或。（两位有一个为1就为1）     |
| ^      | 参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。（两位不一样则为1） |
| <<     | 左移n位就是乘以2的n次方。“a<<b”是把a的各二进位全部左移b位，高位丢弃，低位补0。 |
| >>     | 右移n位就是除以2的n次方。“a>>b”是把a的各二进位全部右移b位。  |

**实例**

``` go
package main

import "fmt"

func main(){
    a := 4 // 二进制 100
    fmt.Printf("a: %b\n", a)
    b := 8 // 二进制 1000
    fmt.Printf("b: %b\n", b)
    
    fmt.Printf("(a & b): %v, %b \n", (a & b),(a & b))
    fmt.Printf("(a | b): %v, %b \n", (a | b),(a | b))
    fmt.Printf("(a ^ b): %v, %b \n", (a ^ b),(a ^ b))
    fmt.Printf("(a << 2): %v, %b \n", (a << 2),(a << 2))
    fmt.Printf("(a >> 2): %v, %b \n", (b >> 2),(b >> 2))
}
```

## 赋值运算符

| 运算符 | 描述                                           |
| ------ | ---------------------------------------------- |
| =      | 简单的赋值运算符，将一个表达式的值赋给一个左值 |
| +=     | 相加后再赋值                                   |
| -=     | 相减后再赋值                                   |
| *=     | 相乘后再赋值                                   |
| /=     | 相除后再赋值                                   |
| %=     | 求余后再赋值                                   |
| <<=    | 左移后赋值                                     |
| >>=    | 右移后赋值                                     |
| &=     | 按位与后赋值                                   |
| \|=    | 按位或后赋值                                   |
| ^=     | 按位异或后赋值                                 |

**实例**

``` go
package main

import "fmt"

func main(){
    var a int
    a = 100
    fmt.Printf("a: %v\n", a)
    a += 1 // a = a + 1
    fmt.Printf("a: %v\n", a)
    a -= 1 // a = a - 1
    fmt.Printf("a: %v\n", a)
    a *= 2 // a = a * 2
    fmt.Printf("a: %v\n", a)
    a /= 2 // a = a / 2
    fmt.Printf("a: %v\n", a)
}
```

# go语言中的流程控制

## go语言中的条件

**条件语句**是用来判断给定的条件是否满足（表达式值是否为<font color = orange>` true`</font>或者<font color = orange>` false`</font>），并根据判断的结果（真或假）决定执行的语句，go语言中的条件语句也是这样的。

## go语言中的条件语句包含如下几种情况

1. **if 语句**：<font color = orange>` if`</font>语句由一个布尔表达式后紧跟一个或多个语句组成。
2. **if...else 语句**：<font color = orange>` if`</font>语句后可以使用可选的<font color = orange>` else`</font>语句，<font color = orange>` else`</font>语句中的表达式在布尔表达式为<font color = orange>` false`</font>时执行。
3. **if 嵌套语句**：你可以在<font color = orange>` if`</font>或<font color = orange>` else if`</font>语句中嵌套一个或多个<font color = orange>` if`</font>或<font color = orange>` else if`</font>语句
4. **switch 语句**：<font color = orange>` switch`</font>语句用于基于不同条件执行不同动作。
5. **select 语句**：<font color = orange>` select`</font>语句类似于<font color = orange>` switch`</font>语句，但是<font color = orange>` select`</font>会随机生成一个可运行的<font color = orange>` case`</font>。如果没有<font color = orange>` case`</font>可运行，它将阻塞，直到有<font color = orange>` case`</font>可运行。

## go语言中的循环语句

go语言中的循环只有for循环，去除了<font color = orange>` while、do while `</font>循环，使用起来更加简介。

1. for循环。
2. for range循环。

## go语言中的流程控制关键字

1. break
2. continue
3. goto

# golang中的if语句

**go语言**中的**if语句**和其他语言中的类似，都是根据给定的条件表达式运算结果来，判断执行流程。

## go语言的if语句语法

``` go
if 布尔表达式 {
    /* 在布尔表达式为 true 时执行 */
}
```

> 注意：在go语言中布尔表达式不用使用括号

## go语言if语句实例演示

**根据布尔值flag判断**

```go
import "fmt"

func test1(){
    var flag = true
    if flag {
        fmt.Println("flag is true")
    }
    fmt.Println("程序运行结束")
}

func main(){
    test1()
}
```

程序运行结果

``` go
flag is true
程序运行结束
```

**根据年龄判断是否成年**

``` go
package main

import "fmt"

func test2(){
    var age = 20
    if age >= 18 {
        fmt.Println("你是成年人")
    }
    fmt.Printf("程序运行结束")
}

func main(){
    test2()
}
```

程序运行结果

``` go
你是成年人
程序运行结束
```

> **初始变量可以声明在布尔表达式里面，注意它的作用域**

``` go
func test3(){
    if age := 20; age > 18 {
        fmt.Println("你是成年人")
    }
    fmt.Printf("程序运行结束")
}

func main(){
    // test1()
    // test2()
    test3()
}
```

程序运行结果

``` go
你是成年人
程序运行结束
```

> **注意：不能使用0或非0表示真假**

``` go
func test4(){
    var i = 1
    if i { // 编译失败
        fmt.Println("here")
    }
    fmt.Printf("程序运行结果")
}
```

**go语言if语句使用提示：**

1. 不需使用括号将**条件包含起来**
2. 大括号<font color = orange>` {}`</font>必须存在，即使只有一行语句
3. 左括号必须在<font color = orange>` if`</font>或<font color = orange>` else`</font>的同一行
4. 在<font color = orange>` if`</font>之后，条件语句之前，可以添加变量初始化语句，使用<font color = orange>` :`</font>进行分隔

# golang中的if else语句

go语言中的if else语句可以根据给定条件二选一。

## go语言的if else语句语法

``` go
if 布尔表达式 {
    /* 在布尔表达式为 true 时执行 */
} else {
    /* 在布尔表达式为 false 时执行 */
}
```

## go语言if else语句实例

**比较两个数的大小**

``` go
package main

import "fmt"

func f1(){
    a := 1
    b := 2
    if a > b {
        fmt.Printf("\"a>b\": %v\n", "a>b")
    } else {
        fmt.Printf("\"a<=b\": %v\n", "a<=b")
    }
}

func main() {
    f1()
}
```

**判断一个数是奇数还是偶数**

``` go
func f2() {
    var s int
    fmt.Println("请输入一个数字：")
    fmt.Scan(&s)
    
    if s%2 == 0{
        fmt.Println("s 是偶数\n")
    } else {
        fmt.Println("s 不是偶数\n")
    }
    fmt.Println("s 的值是："，s)
}
```

**判断一个人是否成年**

```go
func f3(){
    age := 20
    if age >= 18 {
        fmt.Println("你是成年人")
    } else {
        fmt.Println("你还未成年")
    }
}
```

**特殊写法，在if前面添加执行语句**

``` go
func f4() {
    if age := 20; age >= 18 {
        fmt.Println("你是成年人")
    } else {
        fmt.Println("你还未成年")
    }
}
```

**go语言if语句使用提示：**

1. 不需使用括号将**条件包含起来**
2. 大括号<font color = orange>` {}`</font>必须存在，即使只有一行语句
3. 左括号必须在<font color = orange>` if`</font>或<font color = orange>` else`</font>的同一行
4. 在<font color = orange>` if`</font>之后，条件语句之前，可以添加变量初始化语句，使用<font color = orange>` :`</font>进行分隔
5. 在有返回值的函数中，最终的<font color = orange>` return`</font>不能在条件语句中

# golang中的if else if语句

go语言if语句可以进行多重嵌套使用，使用多重判断。

## go语言中的if else if语法

``` go
if 布尔表达式1 {
    // do something
} else if 布尔表达式2 {
    // do something else
} else {
    // catch-all or default
}
```

## go语言中的if else if语法实例

**根据分数判断等级**

``` go
func f5() {
    score := 80
    if score >= 60 && score <= 70 {
        fmt.Println("C")
    }else if score > 70 && score <= 90 {
        fmt.Println("B")
    } else {
        fmt.Println("A")
    }
}
```

程序运行结果

``` go
B
```

**同样也可以写成这样**

``` go
func f5() {
    if score := 80; score >= 60 && score <= 70 {
        fmt.Println("C")
    } else if score > 70 && score <= 90 {
        fmt.Println("B")
    } else {
        fmt.Println("A")
    }
}
```

**输入星期几第一个字母来判断一下是星期几，如果第一个字母一样，则继续判断第二个字母**

``` go
func f6() {
    //  Monday Tuesday Wednesday Thursday Friday Saturday Sunday
    var c string
    fmt.Println("请输入一个字符：")
    fmt.Scan(&c)
    
    if c == "S" {
        fmt.Println("请输入第二个字符：")
        fmt.Scan(&c)
        if c == "a" {
            fmt.Println("Saturday")
        } else if c== "u" {
            fmt.Println("Sunday")
        } else {
            fmt.Println("输入错误")
        }
    } else if c == "F" {
        fmt.Println("Friday")
    } else if c == "M" {
        fmt.Println("Monday")
    } else if c == "T" {
        fmt.Println("输入第二个字符：")
        fmt.Scan(&c)
        if c == "u" {
            fmt.Println("Tuesday")
        } else if c == "h" {
            fmt.Println("Thursday")
        } else {
            fmt.Println("输入错误")
        }
    } else if c == "W" {
        fmt.Println("Wednesday")
    } else {
        fmt.Println("输入错误")
    }
}
```

# golang中嵌套if语句

go语言if语句可以嵌套多级进行判断。

## go语言if嵌套语法

``` go
if 布尔表达式 1 {
    /* 在布尔表达式 1 为 true 时执行 */
    if 布尔表达式 2 {
        /* 在布尔表达式 2 为 true 时执行 */
    }
}
```

## go语言if嵌套实例

**判断三个数的大小**

``` go
package main

import "fmt"

func f1() {
    a := 100
    b := 200
    c := 3
    if a > b {
        if a > c {
            fmt.Println("a最大")
        }
    } else {
        if b > c {
            fmt.Println("b最大")
        } else {
            fmt.Println("c最大")
        }
    }
}

func main() {
    f1()
}
```

**判断男生还是女生，还有是否成年**

``` go
func f2() {
    // 判断男女生及年龄
    gender := "女生"
    age := 16 
    if gender == "男生"{
        fmt.Println("男生")
        if age > 18 {
            fmt.Println("成年")
        } else {
            fmt.Println("未成年")
        }
    } else {
        fmt.Println("女生")
        if age > 18 {
            fmt.Println("成年")
        } else {
            fmt.Println("未成年")
        }
    }
}
```

# golang switch语句

go语言中的<font color = orange>` switch`</font>语句，可以非常容易判断多个值的情况。

## go语言中switch语句的语法

``` go
switch var1 {
    case var1:
    	...
    case var2:
    	...
    default:
    	...
}
```

## go语言中<font color = orange>` switch`</font>语句实例

**判断成绩**

``` go
func f() {
    grade := "A"
    switch grade {
        case "A":
       	 	fmt.Println("优秀")
        case "B":
        	fmt.Println("良好")
        default:
        	fmt.Println("一般")
    }
}
```

运行结果

``` go
优秀
```

**多条件匹配**

go语言<font color = orange>` switch`</font>语句，可以同时匹配多个条件，中间用逗号分隔，有其中一个匹配成功即可。

``` go
func f() {
    day := 3
    switch day {
        case 1,2,3,4,5:
        	fmt.Println("工作日")
        case 6,7:
        	fmt.Println("休息日")
    }
}
```

运行结果

``` go
工作日
```

**case可以是条件表达式**

``` go
func f() {
    score := 90
    switch {
        case score >= 90:
        fmt.Println("享受假期")
        case score < 90 && score >= 80:
        fmt.Println("好好学习吧！")
        default:
        fmt.Println("玩命学习！")
    }
}
```

运行结果

``` go
享受假期
```

**<font color = orange>` fallthrough`</font>可以执行满足条件的下一个<font color = orange>` case`</font>**

``` go
fun f3() {
    a := 100
    switch a {
    	case 100:
        	fmt.Println("100")
        	fallthrough
        case 200:
        	fmt.Println("200")
        case 300:
        	fmt.Println("300")
        default:
        	fmt.Println("other")
    }
}
```

运行结果

``` go
100
200
```

## go语言中<font color = orange>` switch`</font>语句的注意事项

1. 支持多条件匹配
2. 不同的<font color = orange>` case`</font>之间不使用<font color = orange>` break`</font>分隔，默认只会执行一个<font color = orange>` case`</font>
3. 如果想要执行多个<font color = orange>` case`</font>，需要使用<font color = orange>` fallthrough`</font>关键字，也可用<font color = orange>` break`</font>终止
4. 分支还可以使用**表达式**，例如：<font color = orange>` a > 10`</font>

# golang for循环语句

go语言中的<font color =orange>` for`</font>循环，只有<font color = orange>` for`</font>关键字，去除了像其他语言中的<font color = orange>` while`</font>和<font color = orange>` do while`</font>

## go语言for循环语法

``` go
for 初始语句;条件表达式;结束语句{
    循环体语句
}
```

> 注意：for表达式不用加括号

## go语言for循环案例

**循环输出1到10**

```go
func f() {
    for i := 1; i <= 10; i++ {
        fmt.Printf("i: %v\n", i)
    }
}
```

运行结果

```go
i: 1
i: 2
i: 3
i: 4
i: 5
i: 6
i: 7
i: 8
i: 9
i: 10
```

**初始条件，可以写到外面**

```go
func f() {
    i := 1
    for ; i <= 10; i++ {
        fmt.Printf("i: %v\n", i)
    }
}
```

运行结果

```go
i: 1
i: 2
i: 3
i: 4
i: 5
i: 6
i: 7
i: 8
i: 9
i: 10
```

**初始条件和结束条件都可以省略**

``` go
func f() {
    i := 1 // 初始条件
    for i <= 10 {
        fmt.Printf("i: %v\n", i)
        i++ // 结束条件
    }
}
```

运行结果

```go
i: 1
i: 2
i: 3
i: 4
i: 5
i: 6
i: 7
i: 8
i: 9
i: 10
```

这种情况类似其他语言中的<font color = orange>` while`</font>循环

**永真循环**

``` go
func f() {
    for i {
        fmt.Println("我一直在执行~")
    }
}

func main(){
    f()
}
```

运行结果

``` go
我一直在执行~
我一直在执行~
我一直在执行~
我一直在执行~
我一直在执行~
我一直在执行~
我一直在执行~
我一直在执行~
我一直在执行~
......
```

for循环可以通过<font color = orange>` break`</font>、<font color = orange>` goto`</font>、<font color = orange>` return`</font>、<font color = orange>` panic`</font>语句强制退出循环。

# golang for range循环

Go语言中可以使用<font color = orange>` for range`</font>遍历数组、切片、字符串、map 及通道（channel）。通过<font color = orange>` for range`</font>遍历的返回值有以下规律：

1. 数组、切片、字符串返回**索引和值**。
2. map返回**键和值**。
3. 通道（channel）只返回通道内的值。

## go语言for range实例

**循环数组**

``` go
func f() {
    var a = [5]int{1,2,3,4,5}
    for i, v := range a {
        fmt.Printf("i: %d, v: %v\n", i, v)
    }
}

func main() {
    f()
}
```

运行结果

``` go
i: 0, v: 1
i: 1, v: 2
i: 2, v: 3
i: 3, v: 4
i: 4, v: 5
```

**循环切片**

``` go
func f() {
    var s = []int{1,2,3,4,5}
    for _, v := range s {
        fmt.Printf("v: %v\n", v)
    }
}

func main() {
    f()
}
```

运行结果

``` go
v: 1
v: 2
v: 3
v: 4
v: 5
```

**循环map**

``` go
func f() {
    // key : value
    m := make(map[string]string, 0)
    m["name"] = "tom"
    m["age"] = "20"
    m["email"] = "ghz@gmail.com"
    for key, value := range m {
        fmt.Printf("%v:%v", key, value)
    }
}

func main() {
    f()
}
```

运行结果

``` go
name:tomage:20email:ghz@gmail.com
```

**循环字符串**

``` go
func f() {
    var s = "多课网，go教程"
    for i, v := range s {
        fmt.Printf("i: %d, v: %c\n",i, v)
    }
    // %c 按照字符输出
}

func main() {
    f()
}
```

运行结果

``` go
i: 0, v: 多
i: 3, v: 课
i: 6, v: 网
i: 9, v: ，
i: 12, v: g
i: 13, v: o
i: 14, v: 教
i: 17, v: 程
```

# golang流程控制关键字break

<font color = orange>` break`</font>语句可以结束<font color = orange>` for`</font>、<font color = orange>` switch`</font>和<font color = orange>` select`</font>的代码块。

## go语言使用break注意事项

1. 单独在<font color = orange>` select`</font>中使用<font color = orange>` break`</font>和不使用<font color = orange>` break`</font>没有啥区别。
2. 单独在表达式<font color = orange>` switch`</font>语句，并且没有<font color = orange>` fallthough`</font>，使用<font color = orange>` break`</font>和不使用<font color = orange>` break`</font>没有啥区别。
3. 单独在表达式<font color = orange>` switch`</font>语句，并且有<font color = orange>` fallthough`</font>，使用<font color = orange>` break`</font>能够终止<font color = orange>` fallthough`</font>后面的<font color = orange>` case`</font>语句的执行。
4. 带标签的<font color = orange>` break`</font>，可以跳出多层<font color = orange>` select/ switch`</font> 作用域。让break更加灵活，写法更加简单灵活，不需要使用控制变量一层一层跳出循环，没有带<font color = orange>` break`</font>的只能跳出当前语句块。

## go语言break关键字实例

**跳出for循环**

``` go
func f() {
    for i := 0; i < 10; i++ {
        if i == 5 {
            break //退出循环
        }
        fmt.Printf("i: %v\n", i)
    }
}

func main() {
    f()
}
```

运行结果

``` go
i: 0
i: 1
i: 2
i: 3
i: 4
```

**跳出switch**

``` go
func f() {
    i := 2
    switch i {
    case 1:
        fmt.Println("等于1")
        break
    case 2:
        fmt.Println("等于2")
        break
        fallthrough
    case 3:
        fmt.Println("等于3")
        break
    default :
        fmt.Println("不关心")
        break
    }
}

func main() {
    f()
}
```

运行结果

```go
等于2
```

注释掉<font color= orange>` fallthrough`</font>上面的<font color= orange>` break`</font>，运行结果

```go
等于2
等于3
```

**跳转到标签处**

```go
func f() {
    MY_LABEL:    
    for i := 0; i < 10; i++ {
        if i == 5 {
            break MY_LABEL
        }
        fmt.Printf("i: %v\n", i)
    }
    fmt.Println("end...")
}

func main() {
    f()
}
```

运行结果

```go
i: 0
i: 1
i: 2
i: 3
i: 4
end...
```

# golang流程控制关键字continue

<font color= orange>` continue`</font>只能用在循环中，在go中只能用在<font color= orange>` for`</font>循环中，它可以终止本次循环，进行下一次循环。在<font color= orange>` continue`</font>语句后添加标签时，表示开始标签对应的循环。

## go语言<font color = orange>` continue`</font>实例

**输出1到10之间的偶数**

```go
func f() {
    for i := 0; i < 10; i++ {
        if i%2 == 0 {
            fmt.Printf("i: %v\n",i)
        } else {
            continue
        }
    }
}
func main() {
    f()
}
```

运行结果

``` go
i: 0
i: 2
i: 4
i: 6
i: 8
```

**跳转到指定label**

``` go
func f() {
    for i := 0; i < 10; i++ {
        MYLABEL:
        for j := 0; j < 10; j++ {
            if i == 2 && j == 2 {
                continue MYLABEL
            }
            fmt.Printf("%v, %v\n",i, j)
        }
    } 
}

func main() {
    f()
}
```

# golang流程控制关键字goto

<font color = orange>goto</font>语句通过标签进行代码间的**无条件转换**。<font color = orange>goto</font>语句可以在快速跳出循环、避免重复退出上有一定帮助。Go语言中使用<font color = orange>goto</font>语句能简化一些代码的实现过程。例如双层嵌套的for循环要退出时：

## go语言关键字goto实例

**跳转到指定标签**

``` go
func f() {
    a := 0
    if a == 0 {
        goto LABEL1
    } else {
        fmt.Println("other")
    }
    
    LABEL1:
    fmt.Printf("next...")
}

func main() {
    f()
}
```

运行结果

```
next...
```

**跳出双重循环**

``` go
func f() {
    for i:= 0; i < 5; i++ {
        for j :=0; j < 5; j++ {
            if i == 2 && j == 2 {
                goto LABEL1
            }
        }
    }
    LABEL1:
    fmt.Println("label1")
}

func main() {
    f()
}
```

运行结果

```go
label1
```

# golang数组

数组是**相同数据类型**的一组数据的集合，数组一旦定义**长度不能修改**，数组可以通过**下标（或者叫索引）**来访问元素。

## go语言数组的定义

数组定义的语法如下：

``` go
var variable_name [SIZE] variable_type
```

<font color = orange>` variable_name`</font>：数组名称

<font color = orange>` SIZE`</font>：数组长度，必须是常量

<font color = orange>` variable_type`</font>：数组保存元素的类型

**实例**

``` go
package main

import "fmt"

func main() {
    var a [3]int    // 定义一个int类型的数组a，长度是3
    var s [2]string // 定义一个字符串类型的数组s，长度是2
    
    fmt.Printf("a: %T\n, a")
    fmt.Printf("s: %T\n, s")
}
```

运行结果

``` go
a: [3]int
s: [2]string
```

从上面运行结果，我们可以看出，数组和长度和元素类型共同组成了数组的类型。

## go语言数组的初始化

初始化，就是给数组的元素赋值，没有初始化的数组，默认元素都是**零值**，布尔类型是<font color = orange>` false`</font>，字符串是空字符串。

### 没有初始化的数组

**实例**

``` go
package main

import "fmt"

func main() {
    var a [3]int     // 定义一个int类型的数组a，长度是3
    var s [2]string  // 定义一个字符串类型的数组s，长度是2
    var b [2]bool
    
    fmt.Printf("a: %v\n, a")
    fmt.Printf("s: %v\n, s")
    fmt.Printf("b: %v\n, b")
}
```

运行结果

``` go
a: [0 0 0]
b: ["" ""]
c: [false false]
```

### 使用初始化列表

**实例**

```go
package main

import "fmt"

func main() {
    var a = [3]int{1, 2, 3}
    var s = [2]string{"tom", "kite"}
    var b = [2]bool{true, false}
    
    a1 := [2]int{1,2}  // 类型推断
    
    fmt.Printf("a: %v\n", a)
    fmt.Printf("s: %v\n", s)
    fmt.Printf("b: %v\n", b)
    fmt.Printf("a1: %v\n", a1)
}
```

运行结果

``` go
a: [1 2 3]
s: [tom kite]
b: [true false]
a1: [1 2]
```

使用初始化列表，就是将值写在**大括号**里面。

### 省略数组长度

数组长度可以省略，使用<font color = orange>` ...`</font>代替，更加初始化值得数量自动推断，例如：

``` go 
package main

import "fmt"

func main() {
    var a = [...]int{1, 2, 3}
    var s = [...]string{"tom", "kite"}
    var b = [...]bool{true, false}
    
    a1 := [...]int{1,2}  // 类型推断
    
    fmt.Printf("a: %v\n", a)
    fmt.Printf("s: %v\n", s)
    fmt.Printf("b: %v\n", b)
    fmt.Printf("a1: %v\n", a1)
}
```

运行结果

``` go
a: [1 2 3]
s: [tom kite]
b: [true false]
a1: [1 2]
```

### 指定索引值的方式来初始化

可以通过指定所有的方式来初始化，未指定所有的默认为零值。

**实例**

``` go
package main

import "fmt"

func main() {
    var a = [...]int{0: 1, 2: 2}
    var s = [...]string{1:"tom", 2:"kite"}
    var b = [...]bool{2:true, 5:false}
    
    a1 := [...]int{1,2}  // 类型推断
    
    fmt.Printf("a: %v\n", a)
    fmt.Printf("s: %v\n", s)
    fmt.Printf("b: %v\n", b)
    fmt.Printf("a1: %v\n", a1)
}
```

运行结果

```go
a: [1 0 2]
s: [ tom kite]
b: [false false true false false false]
a1: [1 2]
```

# go语言访问数组元素

可以通过下标的方式，来访问数组元素。数组的最大下标为数组长度-1，大于这个下标会发生数组越界。

## 访问数组元素

**实例**

``` go
package main

import "fmt"

func main() {
    var a [2]int
    a[0] = 100
    a[1] = 200
    
    fmt.Printf("a[0]: %v\n", a[0])
    fmt.Printf("a[0]: %v\n", a[1])
    
    // 修改 a[0] a[1]
    a[0] = 1
    a[1] = 2
    
    fmt.Println("-------------")
    
    fmt.Printf("a[0]: %v\n", a[0])
    fmt.Printf("a[1]: %v\n", a[1])
}
```

运行结果

``` go
a[0]: 100
a[1]: 200
-------------
a[0]: 1
a[1]: 2
```

## 根据数组长度遍历数组

可以根据数组长度，通过<font color = orange>` for`</font>循环的方式来遍历数组，数组的长度可以使用<font color = orange>` len`</font>函数获得。

**实例**

``` go
package main

import "fmt"

func main() {
    a := [...]int{1, 2, 3, 4, 5, 6}
    
    for i := 0; i < len(a); i++ {
        fmt.Printf("a[%d]: %v\n", i, a[i])
    }
}
```

运行结果

``` go
a[0]: 1
a[1]: 2
a[2]: 3
a[3]: 4
a[4]: 5
a[5]: 6
```

## 使用<font color = orange>` for range`</font>数组

还可以使用<font color = orange>` for range`</font>循环来遍历数组，range返回数组下标和对应的值

**实例**

``` go
package main

import "fmt"

func main() {
    a := [...]int{1,2,3,4,5,6}
    
    for i, v := range a {
        fmt.Printf("a[%d]: %v\n", i, v)
    }
}
```

运行结果

``` go
a[0]: 1
a[1]: 2
a[2]: 3
a[3]: 4
a[4]: 5
a[5]: 6
```

# golang切片

前面我们学习了数组，数组是固定长度，可以容纳相同数据类型的元素的集合。当长度固定时，使用还是带来一下限制，比如：我们申请的长度太大浪费内存，太小又不够用。

鉴于上述原因，我们有了go语言的切片，可以把切片理解为，可变长度的数组，其实它底层就是使用数组实现的，增加了自动扩容功能。切片（Slice）是一个拥有相同类型元素的可变长度的序列。

## go语言切片的语法

声明一个切片和声明一个数组类似，只要不添加长度就可以了

``` go
var identifier []type
```

切片是引用类型，可以使用<font color = orange>` make`</font>函数来创建切片：

``` go
var slice1 [] type = make([]type, len)

也可以简写为

slice1 := make([]type, len)
```

也可以指定容量，其中capacity（容量）为可选参数。

``` go
make([]T, length, capacity)
```

这里的len是数组的长度并且也是切片的初始长度。

## go语言切片实例

``` go
package main

import "fmt"

func main() {
    var names []string
    var numbers []int
    fmt.Printf("names: %v\n", names)
    fmt.Printf("numbers: %v\n", numbers)
    fmt.Println(name == nil)
    fmt.PrintLn(number == nil)
}
```

运行结果

``` go
name: []
number: []
true
true
```

## go语言切片的长度和容量

切片拥有自己的长度和容量，我们可以通过使用内置的<font color = orange>` len()`</font>函数求长度，使用内置的<font color = orange>` cap()`</font>函数求切片的容量。

**实例**

``` go
package main

import "fmt"

func main() {
    var names = []string{"tom","kite"}
    var number = []int{1,2,3}
    
    fmt.Printf("len: %d cap: %d\n", len(names), cap(names))
    fmt.Printf("len: %d cap: %d\n", len(numbers), cap(numbers))
    
    var s1 = make([]string,2,3)
    fmt.Printf("len: %d cap: %d\n", len(s1), cap(s1))
}
```

运行结果

``` go
len: 2 cap: 2
len: 3 cap: 3
len: 2 cap: 3
```

# golang切片的初始化

切片的初始化方法很多，可以直接初始化，也可以使用数组初始化等。

## 直接初始化

``` go
package main

import "fmt"

func main() {
    s := []int{1,2,3}
    fmt.Printf("s: %v\n",s)
}
```

运行结果

``` go
s: [1 2 3]
```

## 使用数组初始化

``` go
package main

import "fmt"

func main() {
    arr := [...]int{1,2,3}
    s1 := arr[:]
    fmt.Printf("s1: %v\n", s1)
}
```

运行结果

``` go
s1: [1 2 3]
```

## 使用数组的部分元素初始化（切片表达式）

切片的底层就是一个**数组**，所以我们可以**基于数组通过切片表达式得到切片**。切片表达式中的low和high表示一个索引范围（**左包含，右不包含**），得到的切片**长度**=high-low，容量等于得到的切片的底层数组的容量。

``` go
package main

import "fmt"

func main() {
    arr := [...]int{1,2,3,4,5,6}
    s1 := arr[2:5]
    fmt.Printf("s1: %v\n", s1)
    s2 := arr[2:]
    fmt.Printf("s2: %v\n", s2)
    s3 := arr[:3]
    fmt.Printf("s3: %v\n", s3)
}
```

运行结果

``` go
s1: [3 4 5]
s2: [3 4 5 6]
s3: [1 2 3]
```

## 空（nil）切片

一个切片在未初始化之前默认为nil，长度为0，容量为0。

``` go
package main

import "fmt"

func main() {
    var s1 []int
    fmt.Println(s1 == nil)
    fmt.Printf("len: %d, cap: %d\n",len(s1),cap(s1))
}
```

运行结果

``` go
true
len: 0, cap: 0
```

# go语言切片的遍历

切片的遍历和数组的遍历非常类似，可以使用for循环索引遍历，或者for range循环。

## for循环索引遍历

```go
package main

import "fmt"

func main() {
    s1 := []int{1,2,3,4,5}
    for i := 0; i < len(s1); i++ {
        fmt.Printf("s1[%d]: %v/n",i,s1[i])
    }
}
```

运行结果

``` go
s1[0]: 1
s1[1]: 2
s1[2]: 3
s1[3]: 4
s1[4]: 5
```

## for range循环索引遍历

``` go
package main

import "fmt"

func main() {
    var s1 = []int{1,2,3,4,5}
    for _,v := range s1 {
        fmt.Printf("v: %v\n",v)
    }
}
```

运行结果

``` go
v: 1
v: 2
v: 3
v: 4
v: 5
```

# go语言切片元素的添加删除和拷贝

切片是一个动态数组，可以使用<font color = orange>` append()`</font>函数添加元素，go语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。由于，切片是引用类型，通过赋值的方式，会修改原有内容，go提供了<font color = orange>` copy()`</font>函数来拷贝切片。

## 添加元素

``` go
package main

import "fmt"

func main() {
    s1 := []int{}
    s1 = append(s1,1)
    s1 = append(s1,2)
    s1 = append(s1,3,4,5) // 添加多个元素
    fmt.Printf("s1: %v\n", s1)
    
    s3 := []int{3,4,5}
    s4 := []int{1,2}
    s4 = append(s4,s3...) // 添加另外一个切片
    fmt.Printf("s4: %v\n", s4)
}
```

运行结果

``` go
s1: [1 2 3 4 5]
s4: [1 2 3 4 5]
```

## 删除元素

``` go
package main

import "fmt"

func main() {
    s1 := []int{1,2,3,4,5}
    // 删除索引为2的元素
    s1 = append(s1[:2],s1[3:]...)
    fmt.Printf("s1: %v\n",s1)
}
```

运行结果

``` go
s1: [1 2 4 5]
```

公式：要从切片a中删除索引为<font color = orange>` index`</font>的元素，操作方法是<font color = orange>` a = append(a[:index], a[index+1:]...)`</font>

## 查找元素

``` go
package main

import "fmt"

func main() {
    var s1 = []int{1,2,3,4}
    var key = 2
    for i,v := range s1 {
        if v == key {
            fmt.Printf("i: %v\n",i)
        }
    }
}
```

运行结果

``` go
i: 1
```

## 拷贝切片

``` go
package main

import "fmt"

func main() {
    s1 := []int{1,2,3}
    s2 := s1
    s1[0] = 100
    fmt.Printf("s1: %v\n",s1)
    fmt.Printf("s2: %v\n",s2)
    fmt.Println("-----------")
    
    s3 := make([]int,3)
    copy(s3,s1)
    s1[0] = 1
    fmt.Printf("s1: %v\n",s1)
    fmt.Printf("s3: %v\n",s3)
}
```

运行结果

``` go
s1: [100 2 3]
s2: [100 2 3]
-----------
s1: [1 2 3]
s3: [100 2 3]
```

# golang map

map是一种<font color = orange>` key:value`</font>键值对的数据结构容器。map内部实现是哈希表（<font color = orange>` hash`</font>）。

map最重要的一点是通过key来快速检索数据，key类似于索引，指向数据的值。

map是引用类型的。

## map的语法格式

可以使用内建函数make也可以使用map关键字来定义map

``` go
/* 声明变量，默认 map 是 nil */
var map_variable map[key_data_type]value_data_type
/* 使用 make 函数 */
map_variable = make(map[key_data_type]value_data_type)
```

<font color = orange>` map_variable`</font>：map变量名称

<font color = orange>` key_data_type`</font>：key的数据类型

<font color = orange>` value_data_type`</font>：值得数据类型

**实例**

下面声明一个保存个人信息的map

``` go
package main

import "fmt"

func main() {
    m1 := make(map[string]string)
    m1["name"] = "tom"
    m1["age"] = "20"
    m1["email"] = "tom@gmail.com"
    
    fmt.Printf("m1: %v\n",m1)
    
    m2 := map[string]string{
        "name":"kite",
        "age":"20",
        "email":"kite@gmail.com",
    }
    
    fmt.Printf("m2: %v\n",m2)
}
```

运行结果

``` go
m1: map[age:20 email:tom@gmail.com name:tom]
m2: map[age:20 email:kite@gmail.com name:kite]
```

## 访问map

可以通过下标key获得其值，例如：

``` go
package main

import "fmt"

func main() {
    m1 := make(map[string]string)
    m1["name"] = "tom"
    m1["age"] = "20"
    m1["email"] = "tom@gmail.com"
    
    name := m1["name"]
    age := m1["age"]
    email := m1["email"]
    fmt.Printf("name: %v\n",name)
    fmt.Printf("age: %v\n",age)
    fmt.Printf("email: %v\n",email)
}
```

运行结果

``` go
name: tom
age: 20
email: tom@gmail.com
```

## 判断某个键是否存在

go语言中有个判断map中键是否存在的特殊写法，格式如下：

``` go
value, ok := map[key]
```

如果ok为<font color = orange>` true`</font>，存在；否则，不存在。

**实例**

```  go
var m1 = map[string]string{"name":"tom","age":"20","email":"tom@gmail.com"}
var k1 = "name"
var k2 = "age1"
v,ok := m1[k1]
fmt.Printf("v: %v\n",v)
fmt.Printf("ok: %v\n",ok)
fmt.Println("----------")
v,ok = m1[k2]
fmt.Printf("v: %v\n",v)
fmt.Printf("ok: %v\n",ok)
```

运行结果

``` go
v: tom
ok: true
----------
v: 
ok: false
```

# go语言遍历map

可以使用<font color = orange>` for range`</font>循环进行map遍历，得到key和value值。

## 遍历key

``` go
package main

import "fmt"

func main() {
    m := make(map[string]string)
    m["name"] = "tom"
    m["age"] = "20"
    m["email"] = "tom@gmail.com"
    
    for key := range m {
        fmt.Println(key)
    }
}
```

运行结果

``` go
name
age
email
```

## 遍历key和value

``` go
package main

import "fmt"

func main() {
    m := make(map[string]string)
    m["name"] = "tom"
    m["age"] = "20"
    m["email"] = "tom@gmail.com"
    
    for key,value := range m {
        fmt.Println(key + ":" + value)
    }
}
```

运行结果

``` go
name:tom
age:20
email:tom@gmail.com
```

# golang函数

## golang函数简介

函数是go语言中的**一级公民**，我们把所有的功能单元都定义在函数中，可以重复使用。函数包含函数的名称、参数列表和返回值类型，这些构成了函数的签名（signature）。

## go语言中函数特性

1. go语言中有3种函数：普通函数、匿名函数（没有名称的函数）、方法（定义在struct上的函数）。
2. go语言中不允许函数重载（overload），也就是说不允许函数同名。
3. go语言中的函数不能嵌套函数，但可以嵌套匿名函数。
4. 函数是一个值，可以将函数赋值给变量，使得这个变量也成为函数。
5. 函数可以作为参数传递给另一个函数。
6. 函数的返回值可以是一个函数。
7. 函数调用的时候，如果有参数传递给函数，则先拷贝参数的副本，再将副本传递给函数。
8. 函数参数可以没有名称。

## go语言中函数的定义和调用

函数在使用之前必须先定义，可以调用函数来完成某个任务。函数可以重复调用，从而达到代码重用。

## go语言函数定义语法

``` go
func function_name( [parameter list]) (return_type){
    函数体
}
```

**语法解析：**

+ <font color = orange>` func`</font>：函数由<font color = orange>` func`</font>开始声明。
+ <font color = orange>` function_name`</font>：函数名称，函数名和参数列表一起构成函数签名。
+ <font color = orange>` [parameter list]`</font>：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
+ <font color = orange>` return_types`</font>：返回类型，函数返回一列值。<font color = orange>` return_types`</font>是该列值的数据类型。有些功能不需要返回值，这种情况下<font color = orange>` return_types`</font>不是必须的。
+ 函数体：函数定义的代码集合。

## go语言函数定义实例

**定义一个求和函数**

``` go
func sum(a int, b int)(ret int) {
    ret = a + b
    return ret
}
```

**定义一个比较两个数大小的函数**

``` go
func compare(a int, b int)(max int) {
    if a > b {
        max = a
    } else {
        max = b
    }
    return max
}
```

## go语言函数调用

当我们要完成某个任务时，可以调用函数来完成。调用函数要传递参数，如果有返回值可以获得返回值。

``` go
func main() {
    s := sum(1,2)
    fmt.Printf("s: %v\n", s)
    
    max := compare(1,2)
    fmt.Printf("max: %v\n",max)
}
```

运行结果

``` go
s: 3
max: 2
```

# golang函数的返回值

函数可以有0或多个返回值，返回值需要指定数据类型，返回值通过<font color = orange>` return`</font>关键字来指定。

<font color = orange>` return`</font>可以有参数，也可以没有参数，这些返回值可以有名称，也可以没有名称。go中的函数可以有多个返回值。

1. <font color = orange>` return`</font>关键中指定了参数时，返回值可以不用名称。如果<font color = orange>` return`</font>省略参数，则返回值部分必须带名称。
2. 当返回值有名称时，必须使用括号包围，逗号分隔，即使只有一个返回值。
3. 但即使返回值命名了，<font color = orange>` return`</font>中也可以强制指定其它返回值的名称，也就是说<font color = orange>` return`</font>的优先级更高。
4. 命名的返回值是预先声明好的，在函数内部可以直接使用，无需再次声明。命名返回值的名称不能和函数参数名称相同，否则报错提示变量重复定义。
5. <font color = orange>` return`</font>中可以有表达式，但不能出现赋值表达式，这和其它语言可能有所不同。例如<font color = orange>` return a+b`</font>是正确的，但<font color = orange>` return c=a+b`</font>是错误的。

## go语言函数返回值实例

**没有返回值**

``` go
func f1() {
    fmt.Printf("我没有返回值，只是进行一些计算")
}
```

**有一个返回值**

``` go
func sum(a int, b int)(ret int) {
    ret = a + b
    return ret
}
```

``` go
func sum(a int, b int)(int) {
    ret := a + b
    return ret
}
```

**多个返回值，且在return中指定返回的内容**

``` go
func f2() (name string, age int) {
    name = "老郭"
    age = 30
    return name,age
}
```

**多个返回值，返回值名称没有被使用**

``` go
func f3() (name string, age int) {
    name = "老郭"
    age = 30
    return // 等价于return name,age
}
```

**return覆盖命名返回值，返回值名称没有被调用**

``` go
func f4() (name string,age int) {
    n := "老郭"
    a := 30
    return n,a
}
```

``` go
func f4() (string,int) {
    n := "老郭"
    a := 30
    return n,a
}
```

> Go中经常会使用其中一个返回值作为函数是否执行成功、是否有错误信息的判断条件。例如<font color = orange>` return value,exists`</font>、<font color = orange>` return value,ok`</font>、<font color = orange>` return value,err`</font>等。

> 当函数的返回值过多时，例如有4个以上的返回值，应该将这些返回值收集到容器中，然后以返回容器的方式去返回。例如，同类型的返回值可以放进slice中，不同类型的返回值可以放进map中。

> 但函数有多个返回值时，如果其中某个或某几个返回值不想使用，可以通过下划线<font color = orange>` _`</font>来丢弃这些返回值。例如下面的<font color = orange>` f1`</font>函数两个返回值，调用该函数时，丢弃第二个返回值b，只保留了第一个返回值a赋值给了变量<font color = orange>` a`</font>

``` go
package main

import "fmt"

func f1() (int,int){
    return 1,2
}
func main() {
    _,x := f1()
    fmt.Printf("x: %v\n",x)
}
```

运行结果

``` go
x: 2
```

# golang函数的参数

go语言函数可以有0或多个参数，参数需要指定**数据类型**。

声明函数时的参数列表叫做形参，调用时传递的参数叫做实参。

go语言是通过**传值的方式传参**的，意味着传递给函数的是拷贝后的副本，所以函数内部访问、修改的也是这个副本。

go语言可以使用**变长参数**，有时候并不能确定参数的个数，可以使用变长参数，可以在函数定义语句的参数部分使用<font color = orange>` ARGS...TYPE`</font>的方式。这时会将<font color = orange>` ...`</font>代表的参数保存到一个名为ARGS的slice中，注意这些参数的数据类型都是TYPE。

## go语言函数的参数实例

**go语言传参**

``` go
// 形参列表
func f1(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func main() {
    // 实参列表
    r := f1(1,2)
    fmt.Printf("r: %v\n", r)
}
```

**演示参数传递，按值传递**

``` go
func f1(a int) {
    a = 200
    fmt.Printf("a1: %v\n",a)
}

func main() {
    a := 100
    f1(a)
    fmt.Printf("a: %v\n", a)
}
```

运行结果

``` go
a1: 200
a: 100
```

从运行结果可以看到，调用函数f1后，a的值并没有被改变，说明参数传递是拷贝了一个副本，也就是拷贝了一份新的内容进行运算。

> <font color = orange>` map`</font>、<font color = orange>` slice`</font>、<font color = orange>` interface`</font>、<font color = orange>` channel`</font>这些数据类型本身就是**指针**类型的，所以就算是拷贝传值也是拷贝的指针，拷贝后的参数仍然指向底层数据结构，所以修改它们**可能**会影响外部数据结构的值。

``` go
package main

import "fmt"

func f1(a []int) {
    a[0] = 100
}

func main() {
    a := []int{1,2}
    f1(a)
    fmt.Printf("a: %v\n",a)
}
```

运行结果

``` go
a: [1 2]
a: [100 2]
```

> 从运行结果发现，调用函数后，slice内容被改变了。

**变长参数**

``` go
package main

import "fmt"

func f1(args ...int){
    for _,v := range args {
        fmt.Printf("v: %v\n", v)
    }
}
func f2(name string,age int, args ...int) {
    fmt.Printf("name: %v\n",name)
    fmt.Printf("age: %v\n",age)
    for _,v := range args {
        fmt.Printf("v: %v\n", v)
    }
}
func main() {
    f1(1,2,3)
    fmt.Println("--------")
    f1(1,2,3,4,5,6)
    fmt.Println("--------")
    f2("tom",20,1,2,3)
}
```

运行结果

``` go
v: 1
v: 2     
v: 3     
-------- 
v: 1     
v: 2     
v: 3     
v: 4     
v: 5     
v: 6     
-------- 
name: tom
age: 20
v: 1     
v: 2     
v: 3   
```

# golang函数类型与函数变量

可以使用<font color = orange>` type`</font>关键字来定义一个函数类型，语法格式如下：

``` go
type fun func(int, int) int
```

上面语句定义了一个<font color = orange>` fun`</font>函数类型，它是一种函数类型，这种函数接收两个<font color = orange>` int`</font>类型的参数并且返回一个<font color = orange>` int`</font>类型的返回值。

下面我们定义两个这样结构的两个函数，一个求和，一个比较大小：

``` go
func sum(a int,b int) int {
    return a + b
}

func max(a int,b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}
```

下面定义一个<font color = orange>` fun`</font>，把<font color = orange>` sum`</font>和<font color = orange>` max`</font>赋值给它

``` go
package main

import "fmt"

type fun func(int,int) int

func sum(a int, b int) int {
    return a + b
}

func max(a int, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func main() {
    var f fun
    f = sum
    s := f(1,2)
    fmt.Printf("s: %v\n", s)
    f = max
    m := f(3,4)
    fmt.Prinf("m: %v\n", m)
}
```

运行结果

``` go
s: 3
m: 4
```

# golang高阶函数

go语言的函数，可以作为函数的参数，传递给另外一个函数，可以作为，另外一个函数返回值返回。

## go语言函数作为参数

``` go
package main

import "fmt"

func sayHello(name string) {
    fmt.Printf("Hello,%s",name)
}

func f1(name string,f func(string)){
    f(name)
}

func main() {
    f1("tom",sayHello)
}
```

运行结果

``` go
Hello,tom
```

## go语言函数作为返回值

``` go
package main 

import "fmt"

func add(x,y int) int {
    return x + y
}

func sub(x,y int) int {
    return x - y
}

func cal(s string) func(int int) int {
    switch s {
        case "+":
        return add
        case "-":
        return sub
        default:
        return nil
    }
}

func main() {
    add := cal("+")
    r := add(1,2)
    fmt.Printf("r:%v\n",r)
    
    fmt.Println("---------")
    
    sub := cal("-")
    r = sub(100,50)
    fmt.Printf("r:%v \n",r)
}
```

运行结果

``` go
r: 3
--------
r: 50
```

# golang匿名函数

go语言函数不能嵌套，但是在函数内部可以定义匿名函数，实现一下简单功能调用。

所谓匿名函数就是，没有名称的函数。

语法格式如下：

``` go
func (参数列表)(返回值)
```

> 当然可以既没有参数，可以没有返回值

## 匿名函数实例

``` go
func main() {
    max := func(a int,b int) int {
        if a > b{
            return a
        } else {
            return b
        }
    }
    
    i := max(1,2)
    fmt.Printf("i: %v\n",i)
}
```

运行结果

``` go
i: 2
```

自己执行

``` go
func main() {
    // 自己执行
    func(a int,b int){
        max := 0
        if a > b {
            max = a
        } else {
            max = b
        }
        fmt.Printf("max:%v\n",max)
    }(1,2)
}
```

运行结果

``` go
max: 2
```

# golang闭包

闭包可以理解成**定义在一个函数内部的函数**。在本质上，闭包是将函数内部和函数外连接起来的桥梁。或者说是函数和其引用环境的组合体。

闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，<font color = orange>` 闭包=函数+引用环境`</font>。首先我们来看一个例子：

``` go
package main

import "fmt"

// 返回一个函数
func add() func(int) int {
    var x int
    return func(y int) int{
        x += y
        return x
    }
}

func main() {
    var f = add()
    fmt.Println(f(10))
    fmt.Println(f(20))
    fmt.Println(f(30))
    fmt.Println("-------")
    f1 := add()
    fmt.Println(f1(40))
    fmt.Println(f1(50))
}
```

运行结果

``` go
10
30
60
-------
40
90
```

变量<font color = orange>` f`</font>是一个函数并且它引用了其外部作用域中的<font color = orange>` x`</font>变量，此时<font color = orange>` f`</font>就是一个闭包。在<font color = orange>` f`</font>的生命周期内，**变量<font color = orange>` x`</font>也一直有效**。闭包进阶实例1：

``` go
package main

import "fmt"

func add(x int) func(int) int {
     return func(y int) int {
         x += y
         return x
     }
}

func main() {
    var f = add(10)
    fmt.Println(f(10))
    fmt.Println(f(20))
    fmt.Println(f(30))
    
    fmt.Println("--------")
    
    f1 := add(20)
    fmt.Println(f1(40))
    fmt.Println(f1(50))
}
```

运行结果

``` go
20
40
70
--------
60
110
```

闭包进阶示例2：

``` go
func makeSuffixFunc(suffix string) func(string) string {
    return func(name string) string {
        if !string.HasSuffix(name,suffix) {
            return name + suffix
        }
        return name
    }
}

func main() {
    jpgFunc := makeSuffixFunc(".jpg")
    txtFunc := makeSuffixFunc(".txt")
    fmt.Println(jpgFunc("test"))
    fmt.Println(txtFunc("test"))
}
```

运行结果

``` go
test.jpg
test.txt
```

闭包进阶示例3：

``` go
func calc(base int) (func(int) int, func(int) int) {
    add := func(i int) int {
        base += i
        return base
    }
    
    sub := func(i int) int {
        base -= i
        return base
    }
    return add,base
}

func main() {
    f1,f2 := calc(10)
    fmt.Println(f1(1),f2(2))
    fmt.Println(f1(3),f2(4))
    fmt.Println(f1(5),f2(6))
}
```

运行结果

``` go
11 9
12 8
13 7
```

闭包其实并不复杂，只要牢记<font color = orange>` 闭包=函数+引用环境`</font>。

# golang递归

函数内部调用函数自身的函数称为递归函数。

使用递归函数最重要的三点：

1. 递归就是自己调用自己。
2. 必须先定义函数的退出条件，没有退出条件，递归将成为死循环。
3. go语言递归函数很可能会产生一大堆goroutine，也很可能会出现栈空间内存溢出问题。

## go语言递归实例

**阶乘**

``` go
package main

import "fmt"

fun a(n int) int {
    // 返回条件
    if n == 1 {
        return 1
    } else {
        return n * a(n - 1)
    }
}

func main() {
    n := 5
    r := a(n)
    fmt.Printf("r: %v\n",r)
}
```

运行结果

``` go
r: 120
```

**斐波那契数列**

它的计算公式为<font color = orange>` f(n)=f(n-1)+f(n-2)`</font>且<font color = orange>` f(2)=f(1)=1`</font>

``` go
package main

func f(n int) int {
    // 退出点判断
    if n == 1 || n == 2 {
        return 1
    }
    // 递归表达式
    return f(n-1) + f(n-2)
}

func main() {
    r := f(5)
    fmt.Printf("r: %v\n",r)
}
```

运行结果

``` go
r: 5
```

# golang defer语句

go语言中的<font color = orange>` defer`</font>语句会将其后面跟随的语句进行**延迟**处理。在<font color = orange>` defer`</font>归属的函数即将返回时，将延迟处理的语句按<font color = orange>` defer`</font>定义的**逆序**进行执行，也就是说，先被<font color = orange>` defer`</font>的语句最后被执行，最后被<font color = orange>` defer`</font>的语句，最先被执行。

## defer特性

1. 关键字 <font color = orange>` defer`</font> 用于注册延迟调用。
2. 这些调用直到<font color = orange>` return`</font>前才被执行。因此，可以用来做资源清理。
3. 多个<font color = orange>` defer`</font>语句，按先进后出的方式执行。
4. <font color = orange>` defer`</font>语句中的变量，在<font color = orange>` defer`</font>声明时就决定了。

## defer用途

1. 关闭文件句柄
2. 锁资源释放
3. 数据库连接释放

## go语言defer语句实例

**查看执行顺序**

``` go
func main() {
    fmt.Println("start")
    defer fmt.Println("step1")
    defer fmt.Println("step2")
    defer fmt.Println("step3")
    fmt.Println("end")
}
```

运行结果

``` go
start
end
step3
step2
step1
```

# golang init函数

golang有一个特殊的函数<font color = orange>` init`</font>函数，先于<font color = orange>` main`</font>函数执行，实现包级别的一些初始化操作。

## init函数的主要特点

+ init函数先于main函数**自动执行**，不能被其他函数调用；
+ init函数没有输入函数、返回值；
+ **包的每个源文件也可以有多个init函数**，这点比较特殊；
+ 同一个包的init执行顺序，golang没有明确定义，编程时要注意程序不要依赖这个执行顺序。
+ 不同包的init函数按照包导入的依赖关系决定执行顺序。

## golang初始化顺序

初始化顺序：**变量初始化->init()->main()**

**实例**

``` go
package main

import "fmt"

var a int = initVar()

func init() {
    fmt.Println("init2")
}

func init() {
    fmt.Println("init")
}

func initVar() int {
    fmt.Println("init var...")
    return 100
}

func main() {
    fmt.Println("main")
}
```

**运行结果**

``` go
init var...
init2 
init
main...
```

# golang指针

Go语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。传递数据使用指针，而无须拷贝数据。

类型指针不能进行偏移和运算。

Go语言中的指针操作非常简单，只需要记住两个符号：<font color = orange>` &`</font>（取地址）和<font color = orange>` *`</font>（根据地址取值）。

**<font size = 5px>指针地址和指针类型</font>**

每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go语言中使用<font color = orange>` &`</font>字符放在变量前面对变量进行**取地址**操作。

Go语言中的值类型<font color = orange>` （int、float、bool、string、array、struct）`</font>都有对应的指针类型，如：<font color = orange>` *int、*int64、*string`</font>等。

## 指针语法

一个指针变量指向了一个值的内存地址。（也就是我们声明了一个指针之后，可以像变量赋值一样，把一个值的内存地址放入到指针当中。）

类似于变量和常量，在使用指针前你需要声明指针。指针格式如下：

``` go
var var_name *var-type
```

<font color = orange>` var-type`</font>：为指针类型

<font color = orange>` var_name`</font>：为指针变量名

<font color = orange>` *`</font>：用于指定变量是作为一个指针。

## 指针声明实例

``` go
var ip *int 	/* 指向整型 */
var fp *float32 /* 指向浮点型 */
```

## 指针使用实例

``` go
package main

import "fmt"

func main() {
    var a int = 20  /* 声明实际变量 */
    var ip *int		/* 声明指针变量 */
    ip = &a			/* 指针变量的存储地址 */
    fmt.Printf("a 变量的地址是：%x\n",&a)
    /* 指针变量的存储地址 */
    fmt.Printf("ip 变量储存的指针地址：%x\n",ip)
    /* 使用指针访问值 */
    fmt.Printf("*ip 变量的值： %d\n",*ip)
}
```

运行结果

``` go
a 变量的地址是：c000016098
ip 变量储存的指针地址：c000016098
*ip 变量的值： 20
```

# golang指向数组的指针

## 定义语法

``` go
var ptr [MAX]*int  /* 表示数组里面的元素的类型是指针类型 */
```

## 实例演示

``` go
package main

import "fmt"

const MAX int = 3

func main() {
    a := []int{1,3,5}
    var i int
    var ptr [MAX]*int;
    fmt.Println(ptr)  //这个打印出来是[<nil> <nil> <nil>]
    for i = 0; i < MAX; i++ {
        ptr[i] = &a[i]  /* 整数地址赋值给指针数组 */
    }
    for i = 0; i < MAX; i++ {
        fmt.Printf("a[%d] = %d\n",i,*ptr[i]) //*ptr[i]就是打印出相关指针的值了。
    }
}
```

运行结果

``` go
[<nil> <nil> <nil>]
a[0] = 1
a[1] = 3
a[2] = 5
```

# golang类型定义和类型别名

在介绍**结构体**之前，我们先来看看什么是类型定义和类型别名。

## go语言类型定义

**类型定义语法**

``` go
type NewType Type
```

**实例**

``` go
package main

import "fmt"

func main() {
    // 类型定义
    type MyInt int
    // i 为MyInt类型
    var i MyInt
    i = 100
    fmt.Printf("i: %v i: %T\n",i,i)
}
```

运行结果

``` go
i: 100 i: main.MyInt
```

## go语言类型别名

**类型别名的语法**

``` go
type NewType = Type
```

**实例**

``` go
package main

import "fmt"

func main() {
    // 类型别名定义
    type MyInt2 = int
    // i 其实还是int类型
    var i MyInt2
    i = 100
    fmt.Printf("i: %v i: %T\n",i,i)
}
```

运行结果

``` go
i: 100 i: int
```

## go语言类型定义和类型别名的区别

1. 类型定义相当于定义了一个**全新的类型**，与之前的类型不同；但是类型别名并没有定义一个新的类型，而是使用了一个别名来替换之前的类型
2. 类型别名只会在**代码**中存在，在**编译完成**之后并不会存在该别名
3. 因为类型别名和原来的类型是一致的，所以原来类型所拥有的方法，类型别名也**也可以调用**，但是如果是重新定义的一个类型，那么**不可以**调用之前的任何方法

# golang结构体

go语言没有面向对象的概念了，但是可以使用结构体来实现，面向对象编程的一些特性，例如：继承、组合等特性。

## go语言结构体的定义

上一节我们介绍了类型定义，结构体的定义和类型定义类似，只不过多了一个<font color = orange>` struct`</font>关键字，语法结构如下：

``` go
type struct_variable_type struct {
    member definition;
    member definition;
    ...
    menber definition;
}
```

<font color = orange>` type`</font>：结构体关键字

<font color = orange>` struct_variable_type`</font>：结构体类型名称

<font color = orange>` struct`</font>：结构体定义关键字

<font color = orange>` member definition`</font>：成员定义

**实例**

下面我们定义一个人的结构体Person

``` go
type Person struct {
    id int
    name string
    age int
    email string
}
```

以上我们定义了一个Person结构体，有四个成员，来描述一个Person的信息。

形同类型可以**合并到一行**，例如：

``` go
type Person struct {
    id,age int
    name,email string
}
```

## 声明一个结构体变量

声明一个结构体变量和声明一个普通变量相同，例如：

``` go
var tom Person
fmt.Printf("tom: %v\n",tom)
kite := Person{}
fmt.Printf("kite: %v\n",kite)
```

运行结果

``` go
tom: {0 0  }
kite: {0 0  }
```

## 访问结构体成员

可以使用点运算符（<font color = orange>` .`</font>）,来访问结构体成员，例如：

``` go
package main

import "fmt"

func main() {
    type Persion struct {
        id,age int
        name,email string
    }
    
    var tom Person
    tom.id = 1
    tom.name = "tom"
    tom.age = 20
    tom.email = "tom@gmail.com"
    fmt.Printf("tom: %v\n",tom)
}
```

运行结果如下

``` go
tom: {1 20 tom tom@gmail.com}
```

## 匿名结构体

如果结构体是临时使用，可以不同起名字，直接使用，例如：

``` go
package main

import "fmt"

func main() {
    var dog struct {
        id int
        name string
    }
    dog.id = 1
    dog.name = "花花"
    fmt.Printf("dog: %v\n",dog)
}
```

# golang结构体的初始化

**未初始化的结构体，成员都是零值 int 0，float 0.0，bool false，string nil**

**实例**

``` go
package main 

import "fmt"

func main() {
    type Person struct {
        id,age int
        name,email string
    }
    
    var tom Person
    fmt.Printf("tom: %v\n",tom)
}
```

运行结果

``` go
tom: {0 0 "" ""}
```

## 使用键值对对结构体进行初始化

**实例**

``` go
package main 

import "fmt"

func main() {
    type Person struct {
        id,age int
        name,email string
    }
    
    kite := Person {
        id: 1,
        name: "kite",
        age: 20,
        email: "kite@gmail.com",
    }
    fmt.Printf("tom: %v\n",tom)
}
```

运行结果

``` go
kite:{1 20 kite kite@gmail.com}
```

> 注意：
>
> 1. 必须初始化结构体的所有字段。
> 2. 初始值的填充顺序必须与字段在结构体中的声明顺序一致。
> 3. 该方式不能和键值初始化方式混用。

## 部分成员初始化

用不到的成员，可以不进行初始化

``` go
package main

import "fmt"

func main() {
    type Person struct {
        id,age int
        name,email string
    }
    
    kite := Person {
        id: 1,
        name: "kite",
    }
    fmt.Printf("kite: %v\n",kite)
}
```

运行结果

``` go
kite: {1 0 kite ""}
```

# golang结构体指针

结构体指针和普通的变量指针相同，我先来回顾一下普通变量的指针，例如：

``` go
package main

import "fmt"

func main() {
    var name string
    name = "tom"
    // p_name 指针类型
    var p_name *string
    // &name 取name地址
    p_name = &name
    fmt.Printf("name: %v\n",name)
    // 输出指针地址
    fmt.Printf("p_name: %v\n", p_name)
    // 输出指针指向的内容值
    fmt.Printf("*p_name: %v\n",*p_name)
}
```

运行结果

``` go
name: tom
p_name: 0xc00010e120
*p_name: tom
```

## go结构体指针

**实例**

``` go
package main

import "fmt"

func main() {
    type Person struct {
        id int
        name string
    }
    
    var tom = Person{1,"tom"}
    
    var p_person *Person
    p_person = &tom
    fmt.Printf("tom: %v\n", tom)
    fmt.Printf("p_person: %p\n",p_person)
     fmt.Printf("*p_person: %p\n",*p_person)
}
```

运行结果

``` go
tom: {1 tom}
p_person: 0xc000004078
*p_person: {1 tom}
```

## 使用<font color = orange>` new`</font>关键字创建结构体指针

我们还可以通过使用<font color = orange>` new`</font>关键字对结构体进行实例化，得到的是结构体的地址，例如：

``` go
package main

import "fmt"

func main() {
    type Person struct {
        id int,
        name string,
    }
    
    var p_person = new(Person)
    fmt.Printf("p_person: &T\n",p_person)
}
```

运行结果

``` go
p_person: *main.Person
```

从运行结果，我们发现p_person为指针类型

## 访问结构体指针成员

访问结构体指针成员，也使用点运算符（<font color = orange>` .`</font>），例如：

``` go
package main

import "fmt"

func main() {
    type Person struct {
        id int
        name string
    }
    
    var p_person = new(Person)
    fmt.Printf("p_person: %T\n",p_person)
    
    p_person.id = 1
    p_person.name = "tom"
    fmt.Printf("*p_person: %v\n,*p_person")
}
```

运行结果

``` go
p_person: *main.Person
*p_person: {1 tom}
```

# golang结构体作为函数参数

go结构体可以像变量一样，作为函数的参数，传递给函数，这里分为两种情况：

1. 直接传递结构体，这里是一个副本（拷贝），在函数内部不会改变外面结构体内容。
2. 传递结构体指针，这时在函数内部，能够改变外部结构体内容。

## 直接传递结构体

**实例**

``` go
package main

import "fmt"

type Person struct {
    id int
    name string
}

func showPerson(person Person) {
    person.id = 1
    person.name = "kite"
    fmt.Printf("person: %v\n",person)
}

func main() {
    person := Person{1, "tom"}
    fmt.Printf("person: %v\n",person)
    fmt.Println("-------------")
    showPerson(person)
    fmt.Println("-------------")
    fmt.Printf("person: %v\n",person)
}
```

运行结果

``` go
person: {1 tom}
---------------
person: {1 kite}
---------------
person: {1 tom}
```

从运行结果可以看出，函数内部改变了结构体内容，函数外面并没有被改变。

## 传递结构体指针

**实例**

``` go
package main

import "fmt"

type Person struct {
    id int
    name string
}

func showPerson(person *Person) {
    person.id = 1
    person.name = "kite"
    fmt.Printf("person: %v\n",person)
}

func main() {
    person := Person{1,"tom"}
    fmt.Printf("person: %v\n",person)
    fmt.Println("-------------------")
    showPerson(&person)
    fmt.Println("-------------------")
    fmt.Printf("person: %v\n", person)
}
```

运行结果

``` go
person: {1 tom}
---------------
person: &{1 kite}
---------------
person: {1 kite}
```

从运行结果，我们可以看到，调用参数后，参数被改变了。

# golang嵌套结构体

go语言没有面对对象编程思想，也没有继承关系，但是可以通过结构体嵌套来实现这种效果。

下面通过实例演示如何实现结构体嵌套，加入有一个人<font color = orange>` Person`</font>结构体，这个人还养了一个宠物<font color = orange>` Dog`</font>结构体。

下面我们来看一下：

**Dog结构体**

``` go
type Dog struct {
    name string
    color string
    age int
}
```

**Person结构体**

``` go
type Person struct {
    dog Dog
    name string
    age int
}
```

访问它们

``` go
package main

import "fmt"

type Dog struct {
    name string
    color string
    age int
}

type Person struct {
    dog Dog
    name string
    age int
}

func main() {
    var tom Person
    tom.dog.name = "花花"
    tom.dog.color = "黑白花"
    tom.dog.age = 2
    
    tom.name = "tom"
    tom.age = 20
    
    fmt.Printf("tom: %v\n",tom)
}
```

# golang方法

go语言没有面向对象的特性，也没有类对象的概念。但是，可以使用结构体来模拟这些特性，我们都知道面向对象有类方法等概念。我们也可以声明一些方法，属于某个结构体。

## go语言方法的语法

Go中的方法，是一种特殊的函数，定义与struct之上（与struct关联、绑定），被称为struct的接受者（receiver）。

通俗的讲，方法就是有接收者的函数。

语法格式如下：

``` go
type mytype struct{}

func (recv mytype) my_method(pare) return_type {}
func (recv *mytype) my_method(pare) return_type {}
```

<font color = orange>` mytype`</font>：定义一个结构体

<font color = orange>` recv`</font>：接受该方法的结构体 （receiver）

<font color = orange>` my_method`</font>：方法名称

<font color = orange>` pare`</font>：参数列表

<font color = orange>` return_type`</font>：返回值类型

从语法格式可以看出，一个方法和一个函数非常相似，多了一个接受类型。

**实例**

``` go
package main

import "fmt"

type Person struct {
    name string
}

func (per Person) eat() {
    fmt.Println(per.name + "eating...")
   
}

func (per Person) sleep() {
    fmt.Println(per.name + "sleep...")
    
}

func main() {
    var per Person
    per.name = "tom"
    per.eat()
    per.sleep()
}
```

运行结果

``` go
tom eating...
tom sleep...
```

## go语言方法的注意事项

1. 方法的receiver type并非一定要是struct类型，type定义的类型别名、slice、map、channel、func类型等都可以。
2. struct结合它的方法就等价于面向对象中的类，只不过struct可以和它的方法分开，并非一定要属于同一个文件，但必须属于一个包。
3. 方法有两种接收类型：<font color = orange>` T Type`</font>和<font color = orange>` T *Type`</font>，它们之间有区别。
4. 方法就是函数，所有Go中没有方法重载（overload）的说法，也就是说同一个类型的所有方法名必须都唯一。
5. 如果receiver是一个指针类型，则会自动解除引用。
6. 方法和type是分开的，意味着实际的行为（behavior）和数据存储（field）是分开的，但是它们通过receiver建立起关联关系

# golang方法接收者类型

结构体实例，有值类型和指针类型，那么方法的接收者是结构体，那么也有值类型和指针类型。区别就是接收者是否赋值结构体副本、值类型复制，指针类型不复制。

## 值类型结构体和指针类型结构体

**实例**

``` go
package main

import "fmt"

type Person struct {
    name string
}

func main() {
    p1 := Person(name: "tom")
    fmt.Printf("p1: %T\n",p1)
    p2 := &Person(name: "tom")
    fmt.Printf("p2: %T\n",p2)
}
```

运行结果

``` go
p1: main.Person
p2: *main.Person
```

从运行结果，我们可以看出p1是值类型，p2是指针类型。

下面看一个传参结构体的例子：

``` go
package main

import "fmt"

type Person struct {
    name string
}

func showPerson(per Person) {
    fmt.Printf("per: &p\n", &per)
    per.name = "kite"
    fmt.Printf("per: %v\n", per)
}

func showPerson2(per *Person) {
    fmt.Printf("per: %p\n",per)
    per.name = "kite"
    fmt.Printf("per: %v\n",per)
}

func main() {
    p1 := Person{name: "tom"}
    fmt.Printf("p1: %p\n", &p1)
    showPerson(p1)
    fmt.Printf("p1: %v\n", p1)
    fmt.Println("------------")
    p2 := &Person{name: "tom"}
    fmt.Printf("p2: %p\n", p2)
    showPerson2(p2)
    fmt.Printf("p2: %v\n", p2)
}
```

## 方法的值类型和指针类型接收者

值类型的指针类型接收者，本质上和函数传参道理相同。

**实例**

``` go
package main

import "fmt"

type Person struct {
    name string
}

func (per Person) showPerson() {
    fmt.Printf("per: &p\n", &per)
    per.name = "kite"
    fmt.Printf("per: %v\n", per)
}

func (per *Person) showPerson2() {
    fmt.Printf("per: %p\n",per)
    per.name = "kite"
    fmt.Printf("per: %v\n",per)
}

func main() {
    p1 := Person{name: "tom"}
    fmt.Printf("p1: %p\n", &p1)
    p1.showPerson()
    fmt.Printf("p1: %v\n", p1)
    fmt.Println("------------")
    p2 := &Person{name: "tom"}
    fmt.Printf("p2: %p\n", p2)
    p2.showPerson2()
    fmt.Printf("p2: %v\n", p2)
}
```

# golang接口

接口像是一个公司的领导，他会定义一些通用规范，只设计规范，而不实现规范。

go语言的接口，是一种新的**类型定义**，它把所有的**具有共性的方法**定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

语法格式和方法非常类型。

## 接口的语法格式

``` go
/* 定义接口 */
type interface_name interface {
    method_name1 [return_type]
    method_name2 [return_type]
    method_name3 [return_type]
    ...
    method_namen [return_type]
}

/* 定义结构体 */
type struct_name struct {
    /* variables */
}

/* 实现接口方法 */
func (struct_name_variable struct_name) method_name1() [retrun_type] {
    /* 方法实现 */
}
...
func (struct_name_variable struct_name) method_namen() [retrun_type] {
    /* 方法实现 */
}
```

在接口定义中定义，若干个空方法。这些方法都具有通用性。

## 接口实例

下面我定义一个usb接口，有读read和写write两个方法，在定义一个电脑Computer和一个手机Mobile来实现这个接口。

**USB接口**

``` go
type USB interface {
    read()
    write()
}
```

**Computer结构体**

``` go
type Computer struct {
}
```

**Mobile结构体**

``` go
type Mobile struct {
}
```

**Computer实现USB接口方法**

``` go
func (c Computer) read() {
    fmt.Println("computer read...")
}

func (c Computer) write() {
    fmt.Println("computer write=...")
}
```

**Mobile实现USB接口方法**

``` go
func (m Mobile) read() {
    fmt.Println("mobile read...")
}

func (m Mobile) write() {
    fmt.Println("mobile write=...")
}
```

**测试**

``` go
func main() {
    c := Computer{}
    m := Mobole{}
    
    c.read()
    c.write()
    m.read()
    m.write()
}
```

运行结果

``` go
func main() {
    c := Computer{}
    m := Mobile{}
    
    c.read()
    c.write()
    m.read()
    m.write()
}
```

## 实现接口必须实现接口中的所有方法

下面我们定义一个OpenClose接口，里面有两个方法open和close，定义个Door结构体，实现其中一个方法。

``` go
package main

import "fmt"

type OpenClose interface {
    open()
    close()
}

type Door struct { 
}

func (d Door) open() {
    fmt.Println("open door...")
}

func main() {
    var oc Openclose
    oc = Door{} // 这里编译错误，提示只实现了一个接口
}
```

# golang接口值类型接收者和指针类型接收者

这个话题，本质上和方法的值类型接收者和指针类型接收者的思考方法是一样的，值接收者是一个拷贝，是一个副本，而指针接收者，传递的是指针。

## 实例演示

**定义一个Pet接口**

``` go
type Pet interface {
    eat()
}
```

**定义一个Dog结构体**

``` go
type Dog struct {
    name string
}
```

**实现Pet接收（接收者是值类型）**

``` go
func (dog Dog) eat() {
    fmt.Printf("dog: %p\n",&dog)
    fmt.Println("dog eat...")
    dog.name = "黑黑"
}
```

**测试**

``` go
func main() {
    dog := Dog{name: "花花"}
    fmt.Printf("dog: %p\n"，&dog)
    dog.eat()
    fmt.Printf("dog: %v\n"，dog)
}
```

运行结果

``` go
dog: 0xc000046240
dog: 0xc000046250
dog eat...
dog: {花花}
```

从运行结果，我们看出dog的地址变了，说明是复制了一份，dog的name没有变说明，外面的dog变量没有被改变。

**将Pet接口改成指针接收者**

``` go
func (dog *Dog) eat() {
    fmt.Printf("dog: %p\n"，&dog)
    fmt.Println("dog eat...")
    dog.name = "黑黑"
}
```

**再测试**

``` go
func main() {
    dog := &Dog{name: "花花"}
    fmt.Printf("dog: %p\n"，&dog)
    dog.eat()
    fmt.Printf("dog: %v\n"，dog)
}
```

运行结果

``` go
dog: 0xc00008c230
dog: 0xc00008c230
dog eat...
dog: &{黑黑}
```

# golang接口和类型的关系

1. 一个类型可以实现多个接口
2. 多个类型可以实现同一个接口（多态）

## 一个类型实现多个接口

一个类型实现多个接口，例如：有一个Player接口可以播放音乐，在有一个Video接口可以播放视频，一个手机Mobile实现这两个接口，既可以播放音乐，又可以播放视频。

**定义一个Player接口**

``` go
type Player interface {
    PlayMusic()
}
```

**定义一个Video接口**

``` go
type Video interface {
    playVideo()
}
```

**定义Mobile结构体**

``` go
type Mobile struct {   
}
```

**实现两个接口**

``` go
func (m Mobile) playMusic() {
    fmt.Println("播放音乐")
}

func (m Mobile) playVideo() {
    fmt.Println("播放视频")
}
```

**测试**

``` go
func main() {
    m := Mobile{}
    m.playMusic()
    m.playVideo()
}
```

运行结果

``` go
播放音乐
播放视频
```

## 多个类型实现同一个接口

比如，一个宠物接口Pet，猫类型Cat和狗类型Dog都可以实现该接口，都可以把猫和狗当宠物类型对待，这在其他语言中叫多态。

**定义一个Pet接口**

``` go
type Pet interface {
    eat()
}
```

**定义一个Dog结构体**

``` go
type Dog struct {
}
```

**定义一个Cat结构体**

``` go
type Cat struct {
}
```

**实现接口**

```go
func (cat Cat) eat() {
    fmt.Println("cat eat...")
}

func (dog Dog) eat() {
    fmt.Println("dog eat...")
}
```

**测试**

``` go
func main() {
    var p Pet
    p = Cat{}
    p.eat()
    p = Dog{}
    p.eat()
}
```

运行结果

``` go
cat eat...
dog eat...
```

# golang接口嵌套

接口可以通过嵌套，创建新的接口。例如：飞鱼，既可以飞，又可以游泳。我们创建一个飞Fly接口，创建一个游泳接口Swim，飞鱼接口有着两个接口组成。

**飞Flyer接口**

``` go
type Flyer interface {
    fly()
}
```

**创建Swimmer接口**

``` go
type Swimmer interface {
    swim()
}
```

**组成一个接口FlyFish**

``` go
// 接口的组合
type FlyFish interface {
    Flyer
    Swimmer
}
```

**创建一个结构体Fish**

``` go
type Fish struct {
}
```

**实现这个组合接口**

``` go
func (fish Fish) fly() {
    fmt.Println("fly")
}

func (fish Fish) swim() {
    fmt.Println("swim...")
}
```

**测试**

``` go
func main() {
    var ff FlyFish
    ff = Fish{}
    ff.fly()
    ff.swim()
}
```

运行结果

``` go
fly...
swim...
```

# golang通过接口实现OCP设计原则

面向**对象的可复用设计**的第一块基石，便是所谓的“开-闭”原则（Open-Closed Principle，常缩写为OCP）。虽然，go不是面向对象语言，但是也可以模拟实现这个原则，对**扩展**是开放的，对**修改**的关闭的。

## OCP设计原则实例

下面通过一个养宠物的例子，来解释OCP设计原则。

**定义一个宠物接口Pet**

``` go
type Pet interface {
    eat()
    sleep()
}
```

该接口有吃和睡两种方法。

**定义个Dog结构体**

``` go
type Dog struct {
    name string 
    age int
}
```

**Dog实现接口方法**

``` go
func (dog Dog) eat() {
    fmt.Println("dog eat...")
}

func (dog Dog) sleep() {
    fmt.Println("dog sleep...")
}
```

**定义一个Cat结构体**

``` go
type Cat struct {
    name string
    age int
}
```

**Cat实现接口方法**

``` go
func (cat Cat) eat() {
    fmt.Println("cat eat...")
}

func (cat Cat) sleep() {
    fmt.Println("cat sleep...")
}
```

**定义个Person结构体**

``` go
type Person struct {
    name string
}
```

**为Person添加一个养宠物方法**

``` go
// pet 既可以传递Dog也可以传递Cat
func (per Person) care(pet Pet) {
    pet.eat()
    pet.sleep()
}
```

**最后测试一下**

``` go
func main() {
    dog := Dog{}
    cat := Cat{}
    per := Person{}
    
    per.care(dog)
    per.care(cat)
}
```

运行结果

``` go
dog eat...
dog sleep...
cat eat...
cat sleep...
```

# golang模拟OOP的属性和方法

golang没有面向对象的概念，也没有封装的概念，但是可以通过结构体<font color = orange>` struct`</font>和函数**绑定**来实现OOP属性和方法等特性。接收者 receiver **方法**。

**例如**，想要定义一个Person类，有name和age属性，有eat/sleep/work方法。

``` go
package main

import "fmt"

type Person struct {
    name string 
    age int
}

func (per Person) eat() {
    fmt.Println("eat...")
}

func (per Person) sleep() {
    fmt.Println("sleep...")
}

func (per Person) work() {
    fmt.Println("work...")
}

func main() {
    per := Person{
        name : "tom"
        age : 20
    }
    
    fmt.Printf("per: %v\n",per)
    per.eat()
    per.sleep()
    per.work()
}
```

运行结果

``` go
per: {tom 20}
eat...
sleep...
work...
```

# golang继承

golang本质上没有oop的概念也没有继承的概念，但是可以通过结构体嵌套实现这个特性。

**例如**

``` go
package main

import "fmt"


type Animal struct {
    name string
    age int
}

func (a Animal) sleep() {
    fmt.Println("eat...")
}

func (a Animal) sleep() {
    fmt.Println("sleep...")
}

type Dog struct {
    Animal
}

type Cat struct {
    Animal
}

func main() {
    dog := Dog{
        Animal{
            name: "cat"
            age 2
        }
    }
    
    cat := Cat {
        Animal{
            name: "dog"
            age: 3
        }
    }
    
    dog.eat()
    dog.sleep()
    
    cat.eat()
    cat.sleep()
}
```

# golang构造函数

golang没有构造函数概念，可以使用函数来模拟构造函数的功能。

**例如**

``` go
package main

import "fmt"

type Person struct {
    name string
    age int
}

func NewPerson(name string , age int) (*Person,error) {
    if name == "" {
        return nil,fmt.Errof("name 不能为空")
    }
    if age < 0 {
        return nil,fmt.Errof("age 不能小于0")
    }
    return &Person{name: name,age: age},nil
}

func main() {
    person,err := NewPerson("tom",20)
    if err == nil {
        fmt.Printf("person: %v\n", *person)
    }
}
```

# golang包

包可以区分命令空间（一个文件夹中不能有两个同名文件），也可以更好的管理项目。go中创建一个包，一般是创建一个文件夹，在该文件夹里面的go文件中，使用package关键字声明包名称，通常，文件夹名称和包名称相同，并且，同一个文件下面只有一个包。

## 创建包

1. 创建一个名为dao的文件夹。
2. 创建一个dao.go文件。
3. 在该文件中声明包。

``` go
package dao

import "fmt"

func Test1() {
    fmt.Println("test package")
}
```

## 导入包

要使用某个包下面的变量或者方法，需要导入该包，导入包时，要导入从<font color = orange>` GOPATH`</font>开始的包路径，例如，在<font color= orange>` service.go`</font>中导入<font color = orange>` dao`</font>包

``` go
package main

import "dao"

func main() {
    dao.Test1()
}
```

## 包注意事项

+ 一个文件夹下只能有一个package
  + <font color = orange>` import`</font>后面的其实是<font color = orange>` GOPATH`</font>开始的相对路径，包括最后一段，但由于一个目录下只能有一个package，所以<font color = orange>` import`</font>一个路径就等于是<font color = orange>` import`</font>了这个路径下的包。
  + 注意，这里指的是“直接包含”的go文件。如果有子目录，那么子目录的父目录是完全两个包。
+ 比如你实现了一个计算器package，名叫<font color = orange>` calc`</font>，位于<font color = orange>` calc`</font>目录下；但又想给别人一个使用范围，于是在calc下可以建个example子目录（calc/example/），这个子目录有个example.go（calc/example/example.go）。此时example.go可以是main包，里面还可以有个main函数。
+ 一个package的文件不能在多个文件夹下
  + 如果多个文件夹下有重名的package，它们其实是彼此无关的package。
  + 如果一个go文件需要同时使用不同目录下的同名package，需要在<font color = orange>` import`</font>这些目录时为每个目录指定一个package别名。

# golang包管理公共go module

## go module简介

go modules 是 golang1.11 新家的特性，用来管理模块中的**包依赖关系**。

## go mod 使用方法

+ 初始化模块

  <font color = orange>` go mod init <项目模块名称>`</font>

+ 依赖关系处理，根据go.mod文件

  <font color = orange>` go mod tidy`</font>

+ 将依赖包复制到项目下的 vendor 目录

  <font color = orange>` go mod vendor`</font>

  **如果包被屏蔽（墙），可以使用这个命令，随后使用go buil -mod=vendor编译**

+ 显示依赖关系

  <font color = orange>` go list -m all`</font>

+ 显示详细依赖关系

  <font color = orange>` go list -m -json all`</font>

+ 下载依赖

  <font color = orange>` go mod download [path@version]`</font>

  **[path@version]是非必写的**

# golang并发编程之协程

Golang中的并发是**函数**相互独立运行的能力。**Goroutine**是并发运行的函数。Golang提供了Goroutines作为并发处理的一种方式。

创建一个协程非常简单，就是在一个任务函数前面添加一个go关键字：

``` go
go task()
```

## 实例1

``` go
package main

import (
	"fmt"
    "time"
)

func show(msg string) {
    for i := 1; i < 5; i++ {
        fmt.Printf("msg: %v\n", msg)
        time.Sleep(time.Millisecond * 100)
    }
}

func main() {
    go show("java")	// 启动了一个协程
    show("golang") // 在main协程中执行，如果它前面也添加go，程序没有输出
    fmt.Println("end...") // 主函数退出，程序就结束了
}
```

> 查看go关键字去掉的运行效果

## 实例2

``` go
package main

import (
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
    "time"
)

func responseSize(url string) {
    fmt.Println("Step1:",url)
    response,err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("Step2:",url)
    defer response.Body.Close()
    
    fmt.Println("Step3:",url)
    body,err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Step4",len(body))
}

func main() {
    go responseSize("https://www.duoke360.com")
    go responseSize("https://www.baidu.com")
    go responseSize("https://www.jd.com")
    time.Sleep(10 * time.Second)
}
```

# golang并发编程之通道channel

Go 提供了一种称为通道的机制，用于在goroutine之间**共享数据**。当您作为goroutine执行并发活动时，需要在goroutine之间共享资源或数据，通道充当goroutine之间的管道，并提供了一种机制来保证同步交换。

需要在声明通道时指定**数据类型**，我们可以共享内置、命名、结构和引用类型的值和指针。数据在通道上传递：在任何给定时间只有一个goroutine可以访问数据项：因此安装设计不会产生数据竞争。

根据数据交换的行为，有两种类型的通道：无缓冲通道和缓冲通道。无缓冲通道用于执行goroutine之间的同步通信，而缓冲通道用于异步通信。无缓冲通道保证在发送和接收发送的瞬间执行两个goroutine之间的交换。缓冲通道没有这样的保证。

**通道由 make 函数创建，该函数指定 chan 关键字和通道元素类型。**

**<font size = 5px>这是创建无缓冲和缓冲通道的代码块</font>**

**语法**

``` go
Unbuffered := make(chan int) // 整型无缓冲通道
buffered := make(chan int, 10) // 整型有缓冲通道
```

使用内置函数<font color = orange>` make`</font>创建无缓冲和缓冲通道。<font color = orange>` make`</font>的第一个参数需要关键字<font color = orange>` chan`</font>，然后是通道运行交换的数据类型。

**<font size = 5px>这是将值发送到通道的代码块需要使用<font color = orange>` <-`</font>运算符</font>**

**语法**

``` go
goroutine1 := make(chan string, 5) // 字符串缓冲通道 
goroutine1 <- "hello" // 通过通道发送字符串
```

一个包含5个值得缓冲区的字符串类型的goroutine1通道。然后我们通过通道发送字符串“hello”。

**<font size = 5px>这是从通道接收值的代码块</font>**

**语法**

``` go
data := <-goroutine1 // 从通道接收字符串
```

<font color = orange>` <-`</font>运算符添加到通道变量（goroutine1）的左侧，以接收来自通道的值。

## 无缓冲通道

在无缓冲通道中，在接收到任何值之前没有能力保护它。在这种类型的通道中，发送和接收goroutine在任何发送或接收操作完成之前的同一时刻准备就绪。如果两个goroutine没有在同一时刻准备好，则通道会让执行其各种发送或接收操作的goroutine首先等待。同步是通道上发送和接收之间交互的基础，没有另一个就不可能发生。

## 缓冲通道

在缓冲通道中，有能力在接收到一个或多个值之前保存它们。在这种类型的通道中，不要强制goroutine在同一时刻准备好执行发送和接收。当发送或接收阻塞时也有不同的条件。只有当通道中没有要接收的值时，接收才会阻塞。仅当没有可用缓冲区来放置正在发送的值时，发送才会阻塞。

## 通道的发送和接收特性

1. 对于同一个通道，发送操作之间是互斥的，接收操作之间也是互斥的。
2. 发送操作和接收操作中对元素值的处理都是不可分割的。
3. 发送操作在完全完成之前会被阻塞。接收操作也是如此。

**实例**

``` go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

// 创建int类型通道，只能传入int类型值
var values = make(chan int)

func send() {
    rand.Seed(time.Now().UnixNano())
    value := rand.Intn(10)
    fmt.Printf("send: %v\n", value)
    // time.Sleep(time.Second * 5)
    values <- value
}

func main() {
    // 从通道收值
    defer close(values)
    go send()
    fmt.Println("wait...")
    value := <- values
    fmt.Printf("receive: %v\n",value)
    fmt.Println("end...")
}
```

# golang并发编程之WaitGroup实现同步

**实例演示**

查看添加<font color =orange>` WaitGroup`</font>和不添加<font color = orange>` WaitGroup`</font>的区别

``` go
package main

import(
	"fmt"
    "sync"
)

var wg sync.WaitGroup

func hello(i int) {
    defer wg.Done() // goroutine结束就登记-1
    fmt.Println("Hello Goroutine!",i)
}

func main() {
    for i := 0;i < 10;i++ {
        wg.Add(1) // 启动一个goroutine就登记+1
        go hello(i)
    }
    wg.Wait() //等待所有登记的goroutine都结束
}
```

# golang并发编程之runtime包

runtime包里面定义了一些协程管理相关的api

## runtime.Gosched()

让出CPU时间片，重新等待安排任务

``` go
package main

import (
	"fmt"
    "runtime"
)

func show(s string) {
    for i := 0; i < 2; i++ {
        fmt.Println(s)
    }
}

func main() {
    go show("java")
    // 主协程
    for i := 0 i < 2; i++ {
        // 切一下，再次分配任务
        runtime.Gosched() // 注释掉查看结果,让给其他子协程来执行
        fmt.Println("golang")
    }
}
```

## runtime.Goexit()

退出当前协程

``` go
package main

import(
	"fmt"
    "runtime"
    "time"
)

func show() {
    for i := 0; i < 10; i++ {
        if i >= 5 {
            runtime.Goexit()
        }
        fmt.Printf("i: %v\n",i)
    }
}

func main() {
    go show()
    time.Sleep(time.Second)
}
```

## runtime.GOMAXPROCS

设置CPU运行的核心数

``` go
package main

import (
	"fmt"
    "runtime"
    "time"
)

func a() {
    for i := 1; i < 10; i++ {
        fmt.Println("A:",i)
    }
}

func b() {
    for i := 1; i < 10; i++ {
        fmt.Println("B:",i)
    } 
}

func main() {
    fmt.Printf("runtime.NumCPU(): %v\n",runtime.NumCPU())
    runtime.GOMAXPROCS(2) // 修改为1查看效果
    go a()
    go b()
    time.Sleep(time.Second)
}
```

# golang并发编程之Mutex互斥锁实现同步

除了使用channel实现同步之外，还可以使用Mutex互斥锁的方式实现同步，一个协程运行结束前，另一个协程不会进入。

``` go
package main

import (
	"fmt"
    "sync"
    "time"
)

var m int = 100

var lock sync.Mutex

var wt sync.WaitGroup

func add() {
    defer wt.Done()
    lock.Lock()
    m += 1
    time.Sleep(time.Millisecond * 10)
    lock.Unlock()
}

func sub() {
    defer wt.Done()
    lock.Lock()
    time.Sleep(time.Millisecond * 2)
    m -= 1
    lock.Unlock()
}

func main() {
    
    for i:= 0; i < 100; i++ {
        go add()
        wt.Add(1)
        go sub()
        wt.Add(1)
    }
    
    wt.Wait()
    fmt.Printf("m: %v/n",m)
}
```

# golang并发编程之channel的遍历

## 方法1：for循环+if判断

``` go
package main

import "fmt"

func main() {
    c := make(chan int)
    go func() {
        for i :=0; i < 10; i++ {
            c <- i
        }
        close(c)
    }()
    
    for {
        if data, ok := <- c;ok {
            fmt.Printf("data:%v\n",data)
        } else {
            break
        }
    }
}
```

## 方法2：for range

```go
package main

import "fmt"

fun main {
    
    c := make(chan int)
    
    go func() {
        for i := 0; i < 10; i++ {
            c <- i
        }
        close(c)
    }()
    
    for v := range c {
        fmt.Printf("v:%v\n",v)
    }
}
```

> 注意：如果通道关闭，读多写少，没有了就是默认值，例如，int就是0，如果没有关闭就会死锁。

# golang并发编程之select switch

1. select是Go中的一个控制结构，类似于<font color = orange>` switch`</font>语句，用于处理异步IO操作。<font color = orange>` select`</font>会监听case语句中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。

   > select中的case语句必须是一个channel操作
   >
   > select中的default子句总是可运行的。

2. 如果有多个<font color = orange>` case`</font>语句都可以运行，<font color = orange>` select`</font>会随机公平的选出一个执行，其他不会执行。

3. 如果没有可运行的<font color = orange>` case`</font>语句，且有<font color = orange>` default`</font>语句，那么就会执行<font color = orange>` default`</font>的动作。

4. 如果没有可运行的<font color = orange>` case`</font>语句，且没有<font color = orange>` default`</font>语句，<font color = orange>` select`</font>将阻塞，直到某个<font color = orange>` case`</font>通信可以执行。

**实例**

``` go
package main

import (
	"fmt"
    "time"
)

var chanInt = make(chan int)
var chanStr = make(chan string)

func main() {
    go func() {
        chanInt <- 100
        chanStr <- "hello"
        close(chanInt)
        close(chanStr)
    }() // 自己调用自己
    
    for {
        select {
            case r := <- chanInt:
            	fmt.Printf("chanInt: %v/n",r)
            case r := <- chanStr:
            	fmt.Printf("chanStr: %v\n",r)
            default:
            	fmt.Printf("default...")
        }
        time.Sleep(time.Second)
    }
}
```

# golang并发编程之Timer

Timer顾名思义，就是定时器的意思，可以实现一些定时操作，内部也是通过channel来实现的。

**实例**

``` go
package main

import (
	"fmt"
    "time"
)

func main() {
    timer1 := time.NewTimer(time.Second * 2)
    t1 := time.Now()
    fmt.Printf("t1:%v")
    
    t2 := <-time1.C  // 类型是channel，阻塞的，直到指定时间到了
    fmt.Printf("t2:%v\n", t2)
    
    // 如果只是单纯的等待的话，可以使用 time.Sleep 来实现
    timer2 := time.NewTimer(time.Second * 2)
    <-timer2.c
    fmt.Println("2s后")
    
    time.Sleep(time.Second * 2)
    fmt.Println("再一次2s后")
    
    <-time.After(time.Second * 2) // time.After函数的返回值是 chan Time
    fmt.Println("再再一次2s后")
    
    time3 := time.NewTime(time.Second)
    go func() {
        <-timer3.C
        fmt.Println("Timer 3 expired")
    }()
    
    stop := timer3.Stop // 停止定时器
    // 阻止timer事件发生，当该函数执行后，timer计时器停止，相应的时间不再执行
    if stop {
        fmt.Println("Timer 3 stopped")
    }
    
    fmt.Println("before")
    timer4 := time.NewTimer(time.Second * 5) // 原来设置5s
    timer4.Reset(time.Second * 1) // 重新设置时间，即修改NewTimer的时间
    <-timer4.C
    fmt.Println("after")
}
```

# golang并发编程之Ticker

Timer只执行一次，Ticker可以周期的执行。

**实例**

``` go
func main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(time.Second)
    counter := 1
    for _ = range ticker.C {
        fmt.Println("ticker 1")
        counter++
        if counter >= 5 {
            break
        }
    }
    ticker.Stop()
}
```

``` go
func main

import (
    "fmt"
    "time"
)

func main() {
    chanInt := make(chan int)
    
    ticker := time.NewTicker(time.Second)
    
    go func {
        for _ = range ticker.C {
            select {
                case chanInt <- 1:
                case chanInt <- 2:
                case chanInt <- 3:
            }
        }
    }()
    
    sum := 0
    for v := range chanInt {
        fmt.Printf("接收：%v\n",v)
        sum += v
        if sum >= 10 {
            fmt.Printf("sum: %v\n",sum)
            break
        }
    }
}
```

# golang并发编程之原子变量的引入

**先看一个实例**

``` go
func main

import (
    "fmt"
    "sync"
    "time"
)

var i = 100

var lock sync.Mutex

func add() {
    lock.Lock()
    i++
    lock.Unlock()
}

func sub() {
    lock.Lock()
    i--
    lock.Unlock()
}

func main() {
    for i := 0; i < 100; i++ {
        go add()
        go sub()
    } 
    
    time.Sleep(time.Second * 3)
    fmt.Printf("i: %v\n",i)
}
```

> 这是我们之前的写法，使用锁实现协程的同步

**下面使用原子变量**

``` go
func main

import (
    "fmt"
    "sync/atomic"
    "time"
)

var i int32 = 100


func add() {
    // cas compare and swap   old new
    atomic.AddInt32(&i, 1)
}

func sub() {
     atomic.AddInt32(&i, -1)
}

func main() {
    for i := 0; i < 100; i++ {
        go add()
        go sub()
    } 
    
    time.Sleep(time.Second * 3)
    fmt.Printf("i: %v\n",i)
}
```

# golang并发编程之原子操作详解

atomic提供的原子操作能够确保任一时刻只有一个goroutine对变量进行操作，善用atomic能够避免程序中出现大量的锁操作。

atomic常见的操作有：

+ 增减
+ 载入  read
+ 比较并交换cas
+ 交换
+ 存储 write

下面将分别介绍这些操作。

## 增减操作

atomic包中提供了如下以Add为前缀的增减操作：

``` go
- func AddInt32(addr *int32, delta int32)(new int32)
- func AddInt64(addr *int32, delta int32)(new int32)
- func AddUint32(addr *uint32, delta uint32)(new uint32)
- func AddUint32(addr *uint64, delta uint64)(new uint64)
- func AddUintptr(addr *uintptr, delta uintptr)(new uintptr)
```

## 载入操作 

atomic包中提供了如下以Load为前缀的载入操作：

``` go
- func LoadInt32(addr *int32)(val int32)
- func LoadInt64(addr *int64)(val int64)
- func LoadPointer(addr *unsafe.Pointer)(val unsafe.Pointer)
- func LoadUint32(addr *uint32)(val uint32)
- func LoadUint64(addr *uint64)(val uint64)
- func LoadUintptr(addr *uintptr)(val uintptr)
```

> 载入操作能够保证原子的读变量的值，当读取的时候，任何其他CPU操作都无法对该变量进行读写，其实现机制收到底层硬件的支持。

## 比较并交换

该操作简称CAS（Compare And Swap）。这类操作的前缀为<font color = orange>` CompareAndSwap`</font>：

``` go
- func CompareAndSwapInt32(addr *int32,old,new int32)(swapped bool)
- func CompareAndSwapInt64(addr *int64,old,new int64)(swapped bool)
- func CompareAndSwapPointer(addr *unsafe.Pointer,old,new unsafe.Pointer)(swapped bool)
- func CompareAndSwapUint32(addr *uint32,old,new uint32)(swapped bool)
- func CompareAndSwapUint64(addr *uint64,old,new uint64)(swapped bool)
- func CompareAndSwapUintptr(addr *uintptr,old,new uintptr)(swapped bool)
```

> 该操作在进行交换前首先确保变量的值未被更改，及仍然保持参数<font color = orange>` old`</font>所记录的值，满足此前提下才进行交换操作。CAS的做法类似操作数据库时常见的乐观锁机制。

## 交换

此类操作的前缀为：<font color = orange>` Swap`</font>

``` go
- func SwapInt32(addr *int32,new int32)(old int32)
- func SwapInt64(addr *int64,new int64)(old int64)
- func SwapPointer(addr *unsafe.Pointer,new unsafe.Pointer)(old unsafe.Pointer)
- func SwapUint32(addr *uint32,new uint32)(old uint32)
- func SwapUint32(addr *uint64,new uint64)(old uint64)
- func SwapUint32(addr *uintptr,new uintptr)(old uintptr)
```

> 相对于CAS，明显此类操作更为暴力直接，并不管变量的旧值是否被改变，直接赋予新值然后返回被替换的值。

## 存储

此类操作的前缀为<font color = orange> ` Store`</font>：

``` go
- func StoreInt32(addr *int32,val int32)
- func StoreInt64(addr *int64,val int64)
- func StorePointer(addr *unsafe.Pointer,val unsafe.Pointer)
- func StoreUint32(addr *uint32,val uint32)
- func StoreUint64(addr *uint64,val uint64)
- func StoreUintptr(addr *uintptr,val uintptr)
```

> 此类操作确保了写变量的原子性，避免其他操作读到了修改变量过程中的脏数据

# golang标准库os模块

## 文件目录相关

os标准库实现了平台（操作系统）无关的编程接口。

``` go
package main

import(
	"fmt"
    "os"
)

// 创建文件
func createFile() {
    f,err := os.Create("test.txt")
    if err != nil {
        fmt.Printf("err: %v\n", err)
    } else {
        fmt.Printf("f: %v\n",f)
    }
}

// 创建目录
func createDir() {
    // 创建单个目录
    /* err := os.Mkdir("test", os.ModePerm)
    if err != nil {
    	fmt.Printf("err: %v\n",err)
    } */
    err := os.MkdirAll("test/a/b", os.ModePerm)
    if err != nil {
        fmt.Printf("err； %v\n", err)
    }
}

// 删除目录
func removeDir() {
    // 删除单个目录
    /* err := os.Remove("test.txt")
    if err != nil {
    	fmt.Printf("err: %v\n",err)
    } */
    
    err := os.RemoveAll("test")
    if err != nil {
        fmt.Printf("err: %v\n",err)
    }
}

// 获取工作目录
func getWd() {
    dir,err := os.Getwd()
    if err != nil {
        fmt.Printf("err: %v\n", err)
    } else {
        fmt.Printf("dir: %v\n", dir)
    }
}

// 修改工作目录
func chWd() {
    err := os.Chdir("d:/")
    if err != nil {
        fmt.Printf("err: %v\n", err)
    }
}

// 获得临时目录
func getTemp() {
    s := os.TempDir()
    fmt.Printf("s: %v\n", s)
}

// 重命名文件
func renameFile() {
    err := os.Rename("test.txt","test2.txt")
    if err != nil {
        fmt.Printf("err: %v\n", err)
    }
}

// 读文件
func readFile() {
    b, err := os.ReadFile("test2.txt")
    if err != nil {
        fmt.Printf("err: %v\n", err)
    } else {
        fmt.Printf("b: %v\n", string(b[:])) // 取切片里面所有内容
    }
}

// 写文件
func writeFile() {
    s := "hello world"
    os.WriteFile("test2.txt", []byte(s), os.ModePerm)
}

func main() {
    // createFile()
    // createDir()
    // removeDir()
    // getWd()
    // chWd()
    // renameFile()
    // readFile()
    // writeFile()
    // getTemp()
}
```

## File文件读操作

这里结束和<font color = orange>` File`</font>结构体相关的文件读操作

``` go
package main

import (
	"fmt"
    "os"
)

// 打开关闭文件
func openCloseFile() {
    // 只能读
    f,_ := os.Open("a.txt")
    fmt.Printf("f.Name(): %v\n", f.Name())
    // 根据第二个参数 可以读写或者创建
    f2, _ := os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 0755)
    fmt.Printf("f2.Name(): %v\n", f2.Name())
    
    err := f.Close()
    fmt.Printf("err: %v\n", err)
    err2 := f2.Close()
    fmt.Printf("err2: %v\n", err2)
}

// 创建文件
func creatFile() {
    // 等价于：OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
    f, _ := os.Create("a2.txt")
    fmt.Printf("f.Name(): %v\n", f.Name())
    // 第一个参数 目录默认：Temp 第二个参数 文件名前缀
    f2, _ := os.CreateTemp("", "temp")
    fmt.Printf("f2.Name(): %v\n", f2.Name())
}

// 读操作
func readOps() {
    // 循环读取
    /* f, _ := os.Open(a.txt)
    for {
    	buf := make([]byte, 6)
    	n,err := f.Read(buf)
    	fmt.Printf("n: %v\n", n)
    	if err == io.EOF {
    		break
    	}
    }
    f.Close()
	*/
    /* buf := make([]byte, 10)
    f2, _ := os.Open("a.txt")
    // 从5开始读10个字节
    n, _ := f2.ReadAt(buf, 5)
    fmt.Printf("n: %v\n", n)
    fmt.Printf("string(buf): %v\n", string(buf))
    f2.Close() */
    
    // 测试 a目录下面有b和c目录
    /* f, _ := os.Open("a")
    de, _ := f.ReadDir(-1)
    for _, v := range de {
    	fmt.Printf("v.IsDir(): %v\n", v.IsDir())
    	fmt.Printf("v.Name(): %v\n", v.Name())
    } */
    
    // 定位
    f, _ := os.Open("a.txt")
    f.Seek(3, 0)
    buf := make([]byte ,10)
    n, _ := f.Read(buf)
    fmt.Printf("n: %v\n", n)
    fmt.Printf("string(buf): %v\n", string(buf))
    f.Close()
}

func main() {
    // openCloseFile()
    // createFile()
    readOps()
}
```

##  File文件写操作

这里结束和<font color = orange>` File`</font>结构体相关的文件写操作

``` go
package main

import "os"

// 写字节数组
func write() {
    f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0775)
    f.Write([]byte("hello golang"))
    f.Close()
}

// 写字符串
func writeString() {
    f, _ := os.OpenFile("a.txt", os.O_PDWR|os.O_TRUNC, 0775)
    f.WriteString("hello world....")
    f.close()
}

// 随机写
func writeAt() {
    f, _ := os.OpenFile("a.txt", os.O_RDWR, 0775)
    f.WriteAt([]byte("aaa"), 3)   //(待写入数据,偏移量) 通常与Seek()配合使用
    f.Close()
}

func main() {
    // writeString()
    // write
    writeAt
}
```

## 进程相关操作

``` go
package main

import (
	"fmt"
    "os"
    "time"
)

func main() {
    // 获得当前正在运行的进程id
    fmt.Printf("os.Getpid(): %v\n", os.Getpid())
    // 父id
    fmt.Printf("os.Getppid(): %v\n", os.Getppid())
    
    // 设置新进程的属性
    attr := &os.ProcAttr{
        // files指定新进程继承的活动文件对线
        // 前三个分别，标准输入、标准输出、标准错误输出
        Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
        // 新进程的环境变量
        Env: os.Environ(),
    }
    
    // 开始第一个新进程
    p, err := os.Start.Process("C:\\Windows\\System32\\notepad.exe", []string{"C:\\Windows\\System32\\notepad.exe","D:\\a.txt"}, attr)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(p)
    fmt.Println("进程ID：", p.Pid)
    
    // 通过进程ID查找进程
    p2, _ := os.FindProcess(p.Pid)
    fmt.Println(p2)
    
    // 等待10秒，执行函数
    time.AfterFunc(time.Second * 10, func(){
        // 向p进程发送退出信号
        p.Signal(os.Kill)
    })
    
    // 等待进程p的退出，返回进程状态
    ps, _ := p.Wait()
    fmt.Println(ps.String())
}
```

# golang标准库os包和环境相关的方法

``` go
package main

import (
    "fmt"
    "os"
)

func main() {
    // 获得所有环境变量
    s := os.Environ()
    fmt.Printf("s: %v\n", s)
    // 获得某个环境变量
    s2 := os.Getenv("GOPATH")
    fmt.Printf("s2: %v\n", s2)
    // 设置环境变量
    os.Setenv("env1", "env1")
    s2 = os.Getenv("aaa")
    fmt.Printf("s2: %v\n", s2)
    fmt.Println("-----------")
    
    // 查找
    s3, b := os.LookupEnv("env1")
    fmt.Printf("b: %v\n", b)
    fmt.Printf("s3: %v\n", s3)
    
    // 替换
    os.Setenv("NAME", "gopher")
    os.Setenv("BURROW", "/usr/gopher")
    
    fmt.Println(os.ExpandEnv("$NAME lives in ${BURRIW}.")) //$ 后面替换方法的第二个参数
    
    // 清空环境变量
    // os.Clearenv
}
```

# golang标准库io包 input output

Go 语言中，为了方便开发者使用，将IO操作封装在了如下几个包中：

+ io为IO原语（I/O primitives）提供了基本的接口    os File Reader Writer
+ io/ioutil 封装一些实用I/O函数
+ fmt实现格式化I/O，类似C语言的 printf 和 scanf format fmt
+ bufio 实现带缓冲I/O

## io — 基本的IO接口

在 io 包中最重要的是两个接口：Reader 和 Writer 接口。本章所提到的各种 IO 包，都跟这两个接口有关，也就是说，只要实现了这两个接口，它就有了 IO 的功能

### Reader 接口

``` go
type Reader interface {
    Read(p []byte)(n int, err error)
}
```

### Writer 接口

``` go
type Writer interface {
    Write(p []byte)(n int, err error)
}
```

## 实现了Reader和Writer接口的类型

``` go
os.File 同时实现了 io.Reader 和 io.Writer
string.Reader 实现了 io.Reader
bufio.Reader/Writer 分别实现了 io.Reader 和 io.Writer
bytes.Buffer 同时实现了 io.Reader 和 io.Writer
bytes.Reader 实现了 io.Reader
comparess/gzip.Reader/Writer 分别实现了 io.Reader 和 io.Writer
crypto/cipher.StreamReader/Writer 分别实现了 io.Reader 和 io.Writer
crypto/tls.Conn 同时实现了 io.Reader 和 io.Writer
encoding/csv.Reader/Writer 分别实现了 io.Reader 和 io.Writer
```

# golang标准库ioutil包

封装一些实用的I/O函数

| 名称      | 作用                                                       |
| --------- | ---------------------------------------------------------- |
| ReadAll   | 读取数据，返回读到的字节 slice                             |
| ReadDir   | 读取一个目录，返回目录入口数组[]os.FileInfo                |
| ReadFile  | 读一个文件，返回文件内容（字节slice）                      |
| WriteFile | 根据文件路径，写入字节slice                                |
| TempDir   | 在一个目录中创建指定前缀名的临时目录，返回新临时目录的路径 |
| TempFile  | 在一个目录中创建指定前缀的临时文件，返回os.File            |

# golang标准库bufio

**bufio**

> bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对线。

**常量**

``` go
const (
	defaultBufSize = 4096
)
```

**变量**

``` go
var (
    ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
    ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
    ErrBufferFull = errors.New("bufio: buffer full")
    ErrNegativeCount = errors.New("bufio: negative count")
)
```

``` go
var (
    ErrTooLong = errors.New("bufio.Scanner: token too long")
    ErrNegativeAdvance = error.New("bufio.Scanner: SplitFunc returns negative advance count")
    ErrAdvanceTooFar = error.New("bufio.Scanner: SplitFunc returns advance count beyond input")
)
```

> 会被Scanner类型返回的错误。

**type Reader**

``` go
type Reader struct {
    buf []byte
    rd io.Reader // reader provided by the client
    r,w int      // buf read and write positions
    err error
    lastByte	 // last byte read for UnreadByte: -1 means invalid
    lastRuneSize // size of last rune read for UnreadRune; -1 means invalid
}
```

> Reader实现了给一个io.Reader接口对象附加缓冲。

**func NewReader**

``` go
func NewReader(rd io.Reader) *Reader
```

> NewReader创建一个具有默认大小的缓冲、从r读取的*Reader。NewReader 相当于 NewReaderSize(rz,4096)

**func NewReaderSize**

``` go
func NewReaderSize(rd io.Reader, size int) *Reader
```

> NewReaderSize创建一个具有最少有size尺寸的缓冲、从r读取的Reader。如果参数r已经是一个具有足够大的缓冲的Reader类型值，会返回r。

**func (*Reader) Result (r io.Reader)**

``` go
func (b *Reader) Reset(r io.Reader)
```

> Reset丢弃缓冲中的数据，清除任何错误，将b重设为其下层从r读取数据

``` go
func main() {
    s := string.NewReader("ABCDEFG")
    str := string.NewReader("123456")
    br := bufio.NewReader(s)
    b, _ := br.ReadString('\n')
    fmt.Println(b)
    br.Reset(str)
    b, _ = br.ReadString('\n')
    fmt.Println(b)
}
```

**func  (*Reader) Read**

``` go
func (b *Reader) Read(p []byte) (n int, err error)
```

> Read读取数据写入p，本方法返回写入p的字节数。本方法一次调用最多会调用下层Reader接口一次Read方法，因此返回值可能小于len(p)。读取到达结尾时，返回值n将为0而err将为io.EOF。

``` go
func main() {
    s := string.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
    br := bufio.NewReader(s)
    p := make([]byte, 10)
    for {
        n,err := br.Read(p)
        if err == io.EOF {
            break
        } else {
            fmt.Printf("string(p): %v\n", string(p[0:n]))
        }
    }
}
```

**func (*Reader) ReadByte **

``` go
func (b *Reader) ReadByte() (c byte, err error)
```

> ReadByte读取并返回一个字节。如果没有可用的数据，会返回错误。

**func (*Reader) UnreadByte**

``` go
func (b *Reader) UnreadByte() error
```

> UnreadByte吐出最近一次读取操作读取的最后一个字节。（只能吐出最后一次，多次调用会出问题）

``` go
func main() {
    s := strings.NewReader("ABCDEFG")
    br := bufio.NewReader(s)
    
    c, _ := br.ReadByte()
    fmt.Printf("%c\n",c)
    
    c, _ = br.ReadByte()
    fmt.Printf("%c\n",c)
    
    br.UnreadByte()
    c, _ = br.ReadByte()
    fmt.Printf("%c\n",c)
}
```

**func (*Reader) ReadRune**

``` go
func (b *Reader) ReadRune() (r rune, size int, err error)
```

> ReadRune读取一个utf-8编码的unicode码值，返回该码值、其编码长度和可能的错误。如果utf-8编码非法，读取位置只移动1字节，返回U+FFFD，返回值size为1而err为nil。如果没有可用的数据，会返回数据。

**func (*Reader) UnreadRune**

``` go
func (b *Reader) UnreadRune() error
```

> UnreadRune吐出最后一次ReadRune调用读取的unicode码值。如果最近一次读取不是调用的ReadRune，会返回错误。（从这点看，UnreadRun比UnreadByte严格很多）

``` go
func main() {
    s := strings.NewReader("你好，世界！")
    br := bufio.NewReader(s)
    
    c, size, _ := br.ReadRune()
    fmt.Prinf("%c %v\n",c,size)
    
    c, size, _ := br.ReadRune()
    fmt.Prinf("%c %v\n",c,size)
    
    br.UnreadRune()
    c, size, _ := br.ReadRune()
    fmt.Prinf("%c %v\n",c,size)
}
```

**func (*Reader) ReadLine**

``` go
func (b *Reader) ReadLine() (line []byte,isPrefix bool,err error)
```

> ReadLine是一个低水平的行数据读取原语。大多数调用者应使用ReadBytes('\n')或ReadString('\n')代替，或者使用Scanner。
>
> ReadLine尝试返回一行数据，不包括行尾标志的字节。如果行太长超过了缓冲，返回值isPrefix会被设为true，并返回行的前面一部分，该行剩下的部分将在之后的调用中返回。返回值isPrefix会在返回该行最后一个片段时才设为false。返回切片是缓冲的子切片，只在下一次读取操作之前有效。ReadLine要么返回一个非nil的line，要么返回一个非nil的err，两个返回值至少一个非nil。
>
> 返回的文本不包含行尾的标志字节 （"\r\n"或"\n"）。如果输入流结束时没有行尾标志字节，方法不会出错，也不会指出这一情况。在调用ReadLine之后调用的UnreadByte会总是吐出最后一个读取的字节（很可能是该行的行尾标志字节），即使该字节不是ReadLine返回值的一部分。

``` go
func main() {
    s := strings.NewReader("ABC\nDEF\r\nGHI\nGUI")
    br := bufio.NewReader(s)
    
    w, isPrefix, _ := br.ReadLine()
    fmt.Printf("%q %v\n", w, isPrefix)
    
    w, isPrefix, _ := br.ReadLine()
    fmt.Printf("%q %v\n", w, isPrefix)
    
    w, isPrefix, _ := br.ReadLine()
    fmt.Printf("%q %v\n", w, isPrefix)
    
    w, isPrefix, _ := br.ReadLine()
    fmt.Printf("%q %v\n", w, isPrefix)
}
```

**func (*Reader) ReadSlice**

``` go
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
```

> ReadSlice读取直到第一次遇到delim字节，返回缓冲里的包含已读取的数据和delim字节的切片。该返回值只要下一次读取操作之前合法。如果ReadSlice放在读取到delim之前遇到错误，它会返回在错误之前读取的数据在缓冲中的切片以及该错误（一般是io.EOF）。如果在读取到delim之前缓冲就被写满了，ReadSlice失败并返回ErrBufferFull。因为ReadSlice的返回值会被下一次I/O操作重写，调用者应尽量使用ReadBytes或者ReadString替代本方法。当且仅当ReadSlice方法返回的切片不以delim结尾时，会返回一个非nil的错误。

``` go
func main() {
    s := strings.NewReader("ABC,DEG,GHI,JKL")
    br := bufio.NewReader(s)
    
    w, _ := br.ReadSlice(",")
    fmt.Printf("%q\n")
    
    w, _ := br.ReadSlice(",")
    fmt.Printf("%q\n")
    
    w, _ := br.ReadSlice(",")
    fmt.Printf("%q\n")
}
```

**func (*Reader) ReadBytes**

``` go
func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
```

> ReadBytes读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的切片。如果ReadBytes放在读取到delim之前遇到错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。当且仅当ReadBytes方法返回的切片不以delim结尾时，会返回一个非nil的错误。

``` go
func main() {
    s := strings.NewReader("ABC,DEG,GHI,JKL")
    br := bufio.NewReader(s)
    
    w, _ := br.ReadBytes(",")
    fmt.Printf("%q\n")
    
    w, _ := br.ReadBytes(",")
    fmt.Printf("%q\n")
    
    w, _ := br.ReadBytes(",")
    fmt.Printf("%q\n")
}
```

**func (*Reader) ReadString**

``` go
func (b *Reader) ReadString(delim byte)(line String, err error)
```

> ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的切片。如果ReadString放在读取到delim之前遇到错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）。当且仅当ReadString方法返回的切片不以delim结尾时，会返回一个非nil的错误。

``` go
func main() {
    s := strings.NewReader("ABC DEG GHI JKL")
    br := bufio.NewReader(s)
    
    w, _ := br.ReadString(" ")
    fmt.Printf("%q\n")
    
    w, _ := br.ReadString(" ")
    fmt.Printf("%q\n")
    
    w, _ := br.ReadString(" ")
    fmt.Printf("%q\n")
}
```

____

**func (b *Reader) WriteTo**

> WriteTo方法实现了io.WriterTo接口。

``` go
func main() {
    s := string.NewReader("ABCDEFGHIJKLMN")
    br := bufio.NewReader(s)
    b := bytes.NewBuffer(make([]byte,0)
    
    br.WriteTo(b)
    fmt.Printf("%s\n",b)
}
```

**type.Writer**

``` go
type Writer struct {
    err error
    buf []byte
    n int
    wr io.Writer
}
```

> Writer实现了为io.Writer接口对象提供缓冲。如果在向一个Writer类型值写入时遇到了错误，该对象不再接受任何数据，且所有写操作都会返回该错误。在说有数据都写入后，调用者有义务调用Flush方法以保证所有数据都交给下层的io.Writer。

**func NewWriter**

``` go
func NewWriter(w io.Writer) *Writer
```

> NewWriter创建了一个具有默认大小缓冲、写入w的*Writer。NewWriter相当于NewWriterSize(wr.4096)

**func NewWriterSize**

``` go
func NewWriterSize(w io.Writer, size int) *Writer
```

> NewWriterSize创建了一个具有最少有size尺寸的缓冲，写入w的Wirter。如果参数w已经输一个具有足够大缓冲的Writer类型值，会返回w。

**func (*Writer) Reset**

``` go
func (b *Writer) Reset(w io.Writer)
```

> Reset丢弃缓冲中的数据，清除任何错误，将b重设为将其输出写入w。

``` go
func main() {
    b := bytes.NewBuffer(make([]byte,0))
    bw := bufio.NewWriter(b)
    bw.WriteString("123456789")
    c := bytes.NewBuffer(make([]byte,0))
    bw.Reset(c)
    bw.WriteString("456")
    bw.Flush()
    fmt.Println(b)
    fmt.Println(c)
}
```

**func (*Writer) Bufferd**

``` go
func (b *Writer) Buffered() int
```

> Buffered返回缓冲中已使用的字节数。

**func (*Writer) Available**

``` go
func (b *Writer) Available() int
```

> Available返回缓冲中还有多少字节未使用。

**func (*Writer) Write**

``` go
func (b *Writer) Write(p []byte) (nn int, err error)
```

> Write将p的内容写入缓冲。返回写入的字节数。如果返回值nn < len(p)，还会返回一个错误说明原因。

**func (*Writer) WriterString**

``` go
func (b *Writer) WriteString(s string)(nn int, err error)
```

> WriteString写入一个字符串，返回写入的字节数。如果返回值nn < len(p)，还会返回一个错误说明原因。

**func (*Writer) WriterByte**

``` go
func (b *Writer) WriterByte(c byte)(err error)
```

> WriterByte写入单个字节

**func (*Writer) WriterRune**

``` go
func (b *Writer) WriterRune(r rune)(size int, err error)
```

> WriterRune写入一个unicode码值（的utf-8编码）返回写入的字节数和可能的错误。

**func (*Writer) Flush**

``` go
func (b *Writer) Flush() err error
```

> Flush方法将缓冲中的数据写入下层的io.Writer接口。

**func (*Writer) ReadFrom**

> ReadFrom实现了io.Readerfrom接口。

``` go
func main() {
    b := bytes.NewBuffer(make([]byte,0))
    bw := bufio.NewWriter(b)
    fmt.Println(bw.Available()) // 4096
    fmt.Println(bw.Buffered()) // 0
    
    bw.WriteString("ABCDEFGHIJKLMN")
    fmt.Println(bw.Available())
    fmt.Println(bw.Buffered())
    fmt.Printf("%q\n",b)
    
    bw.Flush()
    fmt.Println(bw.Available())
    fmt.Println(bw.Buffered())
    fmt.Printf("%q\n",b)
}
```

``` go
func main() {
	b := bytes.NewBuffer(make([]byte,0))
    bw := bufio.NewWriter(b)
    // 写入缓存
    // byte等同于 int8
    bw.WriteByte('H')
    bw.WriteByte('e')
    bw.WriteByte('l')
    bw.WriteByte('l')
    bw.WriteByte('0')
    bw.WriteByte(' ')
    // rune等同于int32
    bw.WriteRune('世')
    bw.WriteRune('界')
    bw.WriteRune('!')
    // 写入b
    bw.Flush()
    fmt.Println(b)
}
```

``` go
func main() {
    b := bytes.NewBuffer(make([]byte,0))
    s := strings.NewReader("Hello 世界！")
    bw := bufio.NewWriter(b)
    bw.ReadFrom(s)
    // bw.Flush()           ReadFrom无需使用Flush，其自己已经写入
    fmt.Println(b) // Hello 世界！
}
```

____

**type ReadWriter**

``` go
type ReadWriter struct {
    *Reader
    *Writer
}
```

> ReadWriter类型保管了指向Reader和Writer类型的指针，（因此）实现了io.ReadWriter接口。

**func NewReadWriter**

``` go
func NewReadWriter(r *Reader, w *Writer) *ReadWriter
```

> NewReadWriter申请创建一个新的、将读写操作分派给r和w的ReadWriter

``` go
func main() {
    b := bytes.NewBuffer(make([]byte,0))
    bw := bufio.NewWriter(b)
    s := strings.NewReader("123")
    br := bufio.NewReader(s)
    rw := bufio.NewReadWriter(br,bw)
    p, _ := rw.ReadString('\n')
    fmt.Println(string(p))     // 123
    rw.WriteString("asdf") 
    rw.Flush()
    fmt.Println(b) // asdf
}
```

____

**type SplitFunc**

``` go
type SplitFunc func(data []byte,atEOF bool)(advance int, token []byte, err error)
```

> SplitFunc类型代表用于对输出作词法分析的分割函数。
>
> 参数data是未处理的数据的一个开始部分的切片，参数atEOF表示是否Reader接口不能提供更多的数据。返回值是解析位置前进的字节数，将要返回给调用者的token切片，以及可能遇到的错误。如果数据不足以（保证）生成一个完整的token，例如需要一整行数据但data里没有换行符，SplitFunc可以返回（0，nil，nil）来告诉Scanner读取更多的数据写入切片然后用从一个位置起始，长度更长的切片再试一次（调用SplitFunc类型函数）。
>
> 如果返回值err非nil，扫描将终止并将该错误返回给Scanner的调用者。
>
> 除非atEOF为真，永远不会使用空切片data调用SplitFunc类型函数。然而，如果atEOF为真，data却可能是非空的、且包含未处理的文本。
>
> SplitFunc的作用很简单，从data中找出你感兴趣的数据，然后返回并告诉调用者，data中有多少数据你已经处理过了。

**func ScanBytes**

``` go
func ScanBytes(data []byte,atEOF bool) (advance int ,token []byte, err error)
```

> ScanBytes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将每个字节作为一个token返回。

**func ScanRunes**

``` go
func ScanRunes(data []byte, atEOF bool) (advance int ,token []byte, err error)
```

> ScanRunes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将每个utf-8编码的unicode码值作为一个token返回。本函数返回的rune序列和range一个字符串的输出rune序列相同。错误的utf-8编码会翻译为U+FFFD="\xef\xbf\xbd"，但只会消耗一个字节，调用者无法区分正确编码rune和错误编码的rune。

**func SacnWords**

``` go
func ScanWords(data []byte, atEOF bool) (advance int, token []byte , err error)
```

> ScanRunes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将空白（参见unicode.IsSpace）分隔得片段（去掉前后空白后）作为一个token返回。本函数永远不会返回空字符串。用来找出data中单行数据并返回（包括空行）

**func ScanLines**

``` go
func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)
```

> ScanLines是用于Scanner类型得分割函数（符合SplitFunc），本函数会将每一行文本去掉末尾得换行标记作为一个token返回。返回得行可以是空字符串。换行标记为一个可选的得回车后跟一个必选得换行符。最后一行即使没有换行也会作为一个token返回。

**type Scanner**

``` go
type Scanner struct {
    r            io.Reader // The reader by provided by the client
    split        SplitFunc // The function to split the tokens
    maxTokenSize int       // Maximum size of a token; modified by tests
    token        []byte    // Last token retruned by split
    buf          []byte    // Buffer used as argument to split
    start        int       // First non-processed byte in buf
    end          int       // End of data in buf
    err          error     // Sticky err
}
```

> Scanner类型提供了方便的读写数据的接口，如从换行符分隔的文本里读取每一行。成功调用的Scan方法会逐步提供文件的token，跳过token之间的字节。token由SplitFunc类型的分割函数指定；默认的分割函数会将输入分割为多个行，并去掉行尾的换行标志。本报预定义的分割函数可以将文件分割为行、字节、unicode码值、空白分割的word。调用者可以定制自己的分割函数。扫描会在抵达输入流结尾、遇到的第一个I/O错误、token过大不能保存进缓冲时，不可恢复的停止。当扫描停止后，当前读取位置可能会远在最后一个获得token后面。需要更多对错误管理的控制或token很大，或必须从reader连续扫描的程序，应使用bufio.Reader代替。

**func NewScanner**

``` go
func NewSanner(r io.Reader) *Scanner
```

> NewScanner创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines。

**func (*Scanner) Split**

``` go
func (s *Scanner) Split(split SplitFunc)
```

> Split设置该Scanner的分割函数。本方法必须在Scan之前调用。

``` go
func main() {
    s := strings.NewReader("ABC DEF GHI JKL")
    bs := bufio.NewScanner(s)
    bs.Split(bufio.ScanWords)
    for bs.Scan() {
        fmt.Println(bs.Text())
    }
}
```

**func (*Scanner) Scan**

``` go
func (*Sanner) Scan() bool
```

> Scan方法获取当前位置的token（该token可以Bytes或Text方法获得），并让Scanner的扫描位置移动到下一个token。当扫描因为抵达输入流结尾或者遇到错误而停止时，本方法会返回false。在Scan方法返回false后，Err方法将返回扫描时遇到的任何错误；除非是io.EOF，此时Err会返回nil。

``` go
func main() {
    s := strings.NewReader("Hello 世界！")
    bs := bufio.NewScanner(s)
    bs.Split(bufio.ScanBytes)
    for bs.Scan() {
        fmt.Printf("%s", bs.Text())
    }
}
```

**func (*Scanner) Bytes**

``` go
func (s *Scanner) Bytes() []byte
```

> Bytes方法返回最近一次Scan调用生成的token。底层数组指向的数据可能会被下一次Scan的调用重写。

``` go
func main() {
    s := strings.NewReader("Hello 世界！")
    bs := bufio.NewScanner(s)
    bs.Split(bufio.ScanRunes)
    for bs.Scan() {
        fmt.Printg("%s",bs.Text())
    }
}
```

**func (*Scanner) Text**

``` go
func (s *Scanner) Text() string
```

> Bytes方法返回最近一次Scan调用生成的token，会申请创建一个字符串保存token并返回该字符串。

**func (*Scanner) Err**

``` go
func (s *Scanner) Err() error
```

> Err返回Scanner遇到第一个非EOF的错误。

# golang标准库log

## log简介

golang内置了<font color = orange>` log`</font>包，实现简单的日志服务。通过调用<font color = orange>` log`</font>包的函数，可以实现简单的日志打印功能。

## log的使用

<font color = orange>` log`</font>包中有3个系列的日志打印函数，分别<font color = orange>` print`</font>系列、<font color = orange>` panic`</font>系列、<font color = orange>` fatal`</font>系列。

| 函数系类 | 作用                                                         |
| -------- | ------------------------------------------------------------ |
| print    | 单纯打印日志                                                 |
| panic    | 打印日志，抛出panic日常                                      |
| fatal    | 打印日志，强制结束程序（os.Exit(1)），<font color = orange>` defer`</font>函数不会执行 |

**实例**

``` go
func main

import(
	"fmt"
    "log"
)

func main() {
    defer fmt.Println("发生了 panic错误！")
    log.Print("my log")
    log.Printf("my log %d", 100)
    name := "tom"
    age := 20
    log.Println(name," ",age)
    log.Panic("致命错误！")
    // log.Fatal("致命错误！")
    fmt.Println("end...")
}
```

## log配置

**标准log配置**

默认情况下log只会打印出时间，但是实际情况下我们可能还需要获取文件名，行号等信息，<font color = orange>` log`</font>包提供给我们定制的接口。

<font color = orange>` log`</font>包提供两个标准log配置的相关方法：

``` go
func Flag() int        // 返回标准log输出配置
func SetFlags(flag int)// 设置标准log输出配置  
```

**flag参数**

``` go
const {
    // 控制输出日志信息的细节，不能控制输出的顺序和格式
    // 输出的日志在每一项后会有一个冒号分隔，例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: 	        message
    Ldate           = 1 << iota       // 日期：2009/01/23
    Ltime                             // 时间：01:23:23
    Lmiscroseconds                    // 微秒级别的时间：01:23:23.123123 (用于增强Ltime位)
    Llongfile                         // 文件全路径名+行号：/a/b/c/d.go:23
    Lshortfile                        // 文件名+行号：d.go.23 (会覆盖掉Llongfile)
    LUTC                              // 使用UTC时间
    LstFlags        = Ldate | Ltime   // 标准logger的初始值
}
```

**标准日志配置示例**

``` go
package main

import (
	"fmt"
    "log"
)

func main() {
    i := log.Flags()
    fmt.Printf("i: %v\n", i)
    log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
    log.Print("my log")
}
```

输出结果：

``` go
i: 3
2021//11/29 16:49:53 c:/Users/52406/Desktop/golangprojects/duoke360.com/pro01/demo.go:12:my log
```

**日志前缀配置**

<font color = orange>` log`</font>包提供两个日志前缀配置的相关函数：

``` go
func Prefix() string          // 返回日志的前缀配置
func SetPrefix(prefix string) // 设置日志前缀
```

**日志前缀配置示例**

``` go
package main

import (
	"fmt"
    "log"
)

func main() {
    s := log.Prefix()
    fmt.Printf("s: %v\n", s)
    log.SetPrefix("MyLog: ")
    s = log.Prefix()
    fmt.Printf("s: %v\n", s)
    log.Print("my log...")
}
```

输出结果：

``` go
s :
s : MyLog:
MyLog: 2021//11/29 16:51:55 my log...
```

**日志输出位置配置**

前面介绍的都是将日志输出到控制台上，golang的<font color = orange>`log`</font>包还支持将日志输出到文件中。<font color = orange>`log`</font>包提供了<font color = orange>`func SetOutput(w io.Writer)`</font>函数，将日志输出到文件中。

**日志输出位置配置**

``` go
package main

import (
	"log"
    "os"
)

func main() {
    f, err := os.OpenFile("a.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        log.Panic("打开日志文件异常")
    }
    log.SetOutput(f)
    log.Print("my log...")
}
```

**结果**：日志输出到当前目录下a.log文件中

``` go
2021//11/29 16:57:13 my log...
```

**自定义logger**

<font color = orange>`log`</font>包为我们提供了内置函数，让我们能自定义logger。从效果上来看，就是将标题3中的标准日志配置、日志前缀配置、日志输出位置配置整合到一个函数中，是日志配置不在那么繁琐。

<font color = orange>`log`</font>包中提供了<font color = orange>`func New(out io.Writer,perfix string,flag int) *Logger`</font>函数来实现自定义logger。

**示例**

``` go
var logger *log.Logger

func init() {
    logFile, err := os.OpenFile("a.log",os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        log.Panic("打开日志文件异常")
    }
    logger = log.New(logFile,"success",log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
    logger.Println("自定义logger")
}
```

# golang标准库builtin

这些包提供了一些类型声明、变量和常量声明，还有一些便利函数，这个包不需要导入，这些变量和函数就可以直接使用。

## 常用函数

**append**

``` go
func append(slice []Type, elems ...Type) []Type
slice = append(slice, elem1, elem2)       // 直接在slice后面添加单个元素，添加元素类型可以和slice相同，也可以不同
slice = append(slice,anotherSlice...)     // 直接将另外一个slice添加到slice后面，但其本质还是将anotherSlice中的元素一个一个添加到slice中，和第一种方法类似
```

**实例**

``` go
package main

import "fmt"

func main() {
    
    s1 := []int{1, 2, 3}
    i := append(s1, 4)
    fmt.Printf("i: %v\n", i)
    
    s2 := []int{7, 8, 9}
    i2 := append(s1, s2...)
    fmt.Printf("i2: %v\n", i2)
}
```

## len

返回，数组、切片、字符串、通道的长度

**实例**

``` go
package main

import "fmt"

func main() {
    
    s1 := "hello world"
    i := len(s1)
    fmt.Printf("i: %v\n", i)
    
    s2 := []int{1, 2, 3}
    fmt.Printf("len(s2): %v\n", len(s2))
}
```

## print、println

打印输出到控制台。

**实例**

``` go
package main

import "fmt"

func main() {
    name := "tom"
    age := 20
    print(name, "", age, "\n")
    fmt.Println("-----------")
    println(name,"",age)
}
```

## panic

抛出一个panic异常

**实例**

``` go
package main

import "fmt"

func main() {
    defer fmt.Println("panic 异常后执行...")
    panic("panic 错误...")
    fmt.Println("end...")
}
```

## new和make

<font color = orange>` new`</font>和<font color = orange>` make`</font>区别：

1. <font color = orange>` make`</font>只能用来分配及初始化类型为<font color = orange>` slice`</font>，<font color = orange>` map`</font>，<font color = orange>` chan`</font>的数据；<font color = orange>` new`</font>可以分配任意类型的数据
2. <font color = orange>` new`</font>分配返回的是指针，即类型<font color = orange>` *T`</font>；<font color = orange>` make`</font>返回引用，即<font color = orange>` T`</font>；
3. <font color = orange>` new`</font>分配的空间被清零，<font color = orange>` make`</font>分配后，会进行**初始化**、

**实例**

**new**

``` go
func testNew() {
    b := new(bool)
    fmt.Println(*b)
    i := new(int)
    fmt.Println(*i)
    s := new(string)
    fmt.Println(*s)
}

func main() {
    testNew()
}


运行结果：
false
0
""
```

**make**

> 内建函数make(T,args)与new(T)的用途不一样。它只用来创建slice，map和channel，并且返回一个初始化的（而不是置零），类型为T的值（而不是*T）。之所以有所不同，是因为这三个类型的背后引用了使用前必须初始化的数据结构。例如，slice是一个三元描述符，包含了一个指向数据（在数组中）的指针，长度以及容量，在这些项被初始化之前，slice都是nil的。对于slice，map和channal，make初始化这些内部数据结构，并准备好可用的值。

``` go
make([]int , 10 , 100)
```

分配一个有100个int的数组，然后创建一个长度为10，容量为100的slice结构，该slice引用包含前10个元素的数组。对应的，new([]int)返回一个指向新分配的，被置零的slice结构体指针，即指向值为nil的slice的指针。

``` go
var p *[]int = new([]int)      // allocates slice structure; *p == nil; rarelu userful

var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic: 习惯的做法
v := make([]int, 100)

运行结果：
p: &[]
v: [0 0 0 0 0 0 0 0 0 0]
```

# golang标准库bytes

bytes包提供了对**字节切片**进行读写操作的一系列函数，字节切片处理的函数比较分为基本处理函数、比较函数、后缀检查函数、索引函数、分割函数、大小写处理函数和子切片函数等。

## 常用函数

``` go
package main

import (
	"bytes"
    "fmt"
)

func main() {
    var i int = 1
    var j byte = 2
    j = byte(i)                  // 将int类型转为byte类型
    fmt.Printf("j: %v\n", j)
    
    // Contains
    b := []byte("duoke360.com")  // 字符串强转为byte切片
    sublice1 := []byte("duoke360")
    sublice2 := []byte("DuoKe360")
    fmt.Println(bytes.Contains(b, sublice1)) // true
    fmt.Println(bytes.Contains(b, sublice2)) // false
    
    // Count
    s := []byte("hellooooooooo")
    sep1 := []byte("h")
    sep2 := []byte("l")
    sep3 := []byte("o")
    fmt.Println(bytes.Count(s, sep1)) // 1
    fmt.Println(bytes.Count(s, sep2)) // 2
    fmt.Println(bytes.Count(s, sep3)) // 9
    
    // Repeat
    b = []byte("hi")
    fmt.Println(string(bytes.Repeat(b, 1))) // hi
    fmt.Println(string(bytes.Repeat(b, 3))) // hihihi
    
    // Replace
    s = []byte("hello,world")
    old := []byte("o")
    news := []byte("ee")
    fmt.Println(string(bytes.Replace(s, old, news, 0)))  // hello,world
    fmt.Println(string(bytes.Replace(s, old, news, 1)))  // hellee,world
    fmt.Println(string(bytes.Replace(s, old, news, 2)))  // hellee,weerld
    fmt.Println(string(bytes.Replace(s, old, news, -1))) // hellee,weerld,-1代表不限次数
    
    // Runes
    s = []byte("你好世界")
    r := bytes.Runes(s)
    fmt.Println("转换前字符串的长度：", len(s)) // 12
    fmt.Println("转换后字符串的长度：", len(r)) // 4
    
    // Join
    s2 := [][]byte{[]byte("你好")，[]byte("世界")}
    sep4 := []byte(",")
    fmt.Println(string(bytes.Join(s2, sep4)))  // 你好，世界
    sep5 := []byte("#")
    fmt.Println(string(bytes.Join(s2, sep5)))  // 你好#世界
}
```

## Reader类型

Reader实现了<font color = orange>` io.Reader,io.ReaderAt,io.WriteTo,io.Seeker,io.ByteScanner,io.RuneScanner`</font>接口，Reader是只读的，可以seek。

``` go
func testReader() {
    data := "123456789"
    // 通过[]byte创建Reader
    re := bytes.NewReader([]byte(data))
    // 返回未读取部分的长度
    fmt.Println("re len: ", re.Len())   // 9
    // 返回底层数据总长度
    fmt.Println("re size: ", re.Size()) // 9
    fmt.Println("------------")
    
    buf = make([]byte 2)
    for {
        // 读取数据
        n,err := re.Read(buf)
        if err != nil {
            break
        }
        fmt.Println(string(buf[:n])) // 12
        							 // 34
        				             // 56
                                     // 78
                                     // 9
    }
    fmt.Println("------------")
    
    // 设置偏移量，因为上面的操作已经修改了读取位置等信息
    re.Seek(0, 0)
    for {
        // 一个字节一个字节的读
        b, err := re.ReadByte()
        if err != nil {
            break
        }
        fmt.Println(string(b))
    }
    fmt.Println("------------")
    
    re.Seek(0, 0)
    off := int64(0)
    for {
        // 指定偏移量读取
        n, err := re.ReadAt(buf, off)
        if err != nil {
            break
        }
        off += int64(n)
        fmt.Println(off, string(buf[:n]))  // 2 12
                                           // 4 34
                                           // 6 56
                                           // 8 78
    }
}
```

## Buffer类型

缓冲区是具有读取和写入方法的可变大小的字节缓冲区。Buffer的零值是准备使用的空缓冲区。

声明一个Buffer的四种方法：

``` go
var b byte.Buffer                    // 直接定义一个Buffer变量，不再初始化，可以直接使用
b := new(bytes.Buffer)               // 使用New返回Buffer变量
b := bytes.NewBuffer(s []byte)       // 从一个[]byte切片，构造一个Buffer
b := bytes.NewBufferString(s string) // 从一个string变量，构造一个Buffer
```

**往buffer中写入数据**

``` go
b.Write(d []byte)              // 将切片d写入Buffer尾部
b.WriteString(s string)        // 将字符串s写入Buffer尾部
b.WriteByte(c byte)            // 将字符c写入Buffer尾部
b.WriteRune(r rune)            // 将一个rune类型的数据放到缓冲器的尾部
b.WriteTo(w io.Writer)         // 将Buffer中的内容输出到实现io.Writer接口的可写入对象中
```

> 注：将文件中的内容写入Buffer，则使用ReadForm(i io.Reader)

**从Buffer中读取数据到指定容器**

``` go
c := make([]byte, 8)
b.Read(c) // 一次读取8个byte到c容器中，每次读取新的8个byte覆盖c中的原来内容
b.ReadByte() // 读取第一个byte，b的第1个byte被拿掉，赋值给 a => a,_ := b.ReadByte()
b.ReadRune() // 读取第一个rune，b的第1个rune被拿掉，赋值给 r => r,_ := b.ReadRune()
b.ReadBytes(delimiter byte) // 需要一个byte作为分隔符，读的时候从缓冲器里面找第一个出现的分隔符(delim),找到后，把从缓冲区头部到分隔符之间的所有byte进行返回，作为byte类型的slice，返回后缓冲器也会空掉一部分
b.ReadString(delimioter byte) // 需要一个byte作为分隔符，读的时候从缓冲器里面找第一个出现的分隔符(delim),找到后，把从缓冲区头部到分隔符之间的所有byte进行返回，作为字符串返回，返回后，缓冲器也会空掉一部分b.ReadForm(i io.Reader):从一个实现io.Reader接口r，把r里面的内容读到缓冲器里，n返回度的数量

file,_ = os.Open("text.txt")
buf := bytes.NewBufferString("Hello world")
buf.ReadFrom(file)
// 将text.txt内容追加到缓冲器的尾部
fmt.Println(buf.String())
// 清空数据
b.Reset()
// 转为字符串
b.String()
```

# golang标准库errors

errors包实现了操作错误的函数。语言使用error类型来返回函数执行过程中遇到的错误，如果返回的error值为nil，则表示未遇到错误，否则error会返回一个字符串，用于说明遇到了什么错误。

## error结构

``` go
type error interface {
    Error() string
}
```

你可以用任何类型去实现它（只要添加一个Error()方法即可），也就是说，error可以是任何类型，这意味着，函数返回的error值实际可以包含任意信息，不一定是字符串。

error不一定表示一个错误，它可以表示任何信息，比如io包中就用error类型的<font color = orange>` io.EOF`</font>表示数据读取结束，而不是遇到了什么错误。

error包实现了一个最简单的error类型，只包含一个字符串，它可以记录大多数情况下遇到的错误信息。errors包的用法也很简单，只有一个<font color = orange>` New`</font>函数，用于生成一个最简单的error对象：

``` go
func New(text string) error
```

``` go
package main

import (
    "error"
	"fmt"
)

func check(s string) error {
    if s == "" {
        return error.New("字符串不能为空")
    } else {
        return nil
    }
}

func main() {
    check("hello")
    err := check("")
    fmt.Printf("err: %v\n" , err.Error())
}
```

## 自定义错误

``` go
package main

import (
	"fmt"
    "time"
)

// MyError is an error implementation that includes a time and message
type MyError struct {
    When time.Time
    What string
}

func (e MyError) Error() string {  // 实现错误接口
    return fmt.Sprintf("%v: %v", e.When,e.What)
}

func oops() error {
    return MyError {
        time.Date(1989,3,15,22,30,0,0,time.UTC)
        "the file system has gone away"
    }
}

func main() {
    if err := oops(); err != nil {
        fmt.Println(err)
    }
}
```

# golang标准库中的sort包

## sort包的内容，以及使用

sort包提供了排序切片和用户自定义数据集以及相关功能的函数。

sort包主要针对<font color= orange>` []int`</font>、<font color= orange>` []float64`</font>、<font color= orange>` []string`</font>、以及其他**自定义切片**的排序。

## 结构体

三种结构体的方法都是一样的，只是分别针对int切片、float64切片、strings切片这三种不同的类型。然后三种结果都有五个公开的方法

``` go
type IntSlice struct ([]int)
type Float64Slice
type StringSlice
```

``` go
func (p xxxSlice) Len() int  //切片长度
func (p xxxSlice) Less(i,j int) bool
func (p xxxSlice) Swap(i,j int)
func (p xxxSlice) Search(x xxx) int
// 这个和后面那个功能一样
func (p xxxSlice) Sort()
```

## 函数

``` go
func Ints(a []int)
func IntsAreSorted(a []int) bool         // 判断是否排序过
func SearchInts(a []int, x int) int      // 通过二分查找找到指定元素
func Float64s(a []float64)
func Float64sAreSorted(a []float64) bool
func SearchFloat64s(a []float64, x float64) bool
func Strings(a []string)
func StringsAreSorted(a []string) bool
func SearchStrings(a []string, x string) int
func Sort(data Interface)
func Stable(data Interface)
func Reverse(data Interface) Interface
func ISSorted(data Interface) bool
func Search(n int, f func(int) bool) int
```

**接口 type interface**

``` go
type Interface interface {
    Len() int           // Len方法返回集合中的元素个数
    Less(i,j int) bool  // i>j, 该方法返回索引i的元素是否比索引j的元素小。
    Swap(i,j int)       // 交换i,j的值
}
```

**实例**

``` go
package main

import (
	"fmt"
    "sort"
)

type NewInts []uint

func (n NewInts) Len() int {
    return len(n)
}

func (n NewInts) Less(i,j int) bool {
    fmt.Println(i,j,n[i]<n[j],n)
    return n[i] < n[j]
}

func (n NewInts) Swap(i,j int) {
    n[i],n[j] = n[j],n[i]
}

func main() {
    n := []uint{1,3,2}
    sort.Sort(NewInts(n))
    fmt.Println(n)
}
```

## 综合实例

**<font color =orange>` []float64`</font>**：

``` go
f := []float64(1.1, 4.4, 5.5, 3.3, 2.2)
sort.Float64s(f)
fmt.Printf("f: %v\n", f)
// f:[1.1 2.2 3.3 4.4 5.5]
```

**<font color =orange>` []int`</font>**：

``` go
f := []int{3,5,1,2,4}
sort.Ints(f)
fmt.Printf("f: %v\n",f)
// f: [1 2 3 4 5]
```

**string**：

``` go
// 字符串排序，先比较高位，相同的在比较低位
// [] string
ls := sort.StringSlice {
    "100",
    "42",
    "41",
    "3",
    "2",
}
fmt.Println(ls)  // [100 42 41 3 2]
sort.Strings(ls)
fmt.Println(ls)  // [100 2 3 41 42]


// 字符串排序，先比较高位，相同的在比较低位
ls := sort.StringSlice {
    "d",
    "ac",
    "c",
    "ab",
    "e",
}
fmt.Println(ls)  // [d ac c ab e]
sort.Strings(ls)
fmt.Println(ls)  // [ab ac c d e]

// 汉字排序，依次比较byte大小
ls := sort.StringSlice {
    "啊",
    "博",
    "次",
    "得",
    "饿",
    "周",
}
fmt.Println(ls)  // [啊 博 次 得 饿 周]
sort.Strings(ls)
fmt.Println(ls)  // [博 周 啊 得 次 饿]

for _, v := range ls {
    fmt.Println(v, []byte(v))
}

// 博 [229 141 154]
// 周 [229 145 168]
// 啊 [229 149 138]
// 得 [229 190 151]
// 次 [230 172 161]
// 饿 [233 165 191]
```

**复杂结构：<font color = orange>` [][]int`</font>:**

``` go
type testSlice [][]int

func (l testSlice) Len() int {
    return len(l)
}

func (l testSlice) Swap(i, j, int) {
    l[i],l[j] = l[j],l[i]
}

func (l testSlice) Less(i, j int) bool {
    retrun l[i][1] < l[j][1]  // 比较第二个位置
}

func main() {
    ls := testSlice {
        {1,4}
        {9,3}
        {7,5}
    }
    
    fmt.Println(ls)  // [[1 4] [9 3] [7 5]]
	sort.Strings(ls)
	fmt.Println(ls)  // [[9 3] [1 4] [7 5]]
}
```

**复杂结构体：<font color = orange>` []map[string]int    [{"k":0},{"k1":1},{"k2":2}]`</font>:**

``` go
type testSlice []map[string]float64

func (l testSlice) Len() int {
    return len(l)
}

func (l testSlice) Swap(i, j, int) {
    l[i],l[j] = l[j],l[i]
}

func (l testSlice) Less(i, j int) bool {
    retrun l[i]["a"] < l[j]["a"]  // 按照"a"对应得值排序
}

func main() {
    ls := testSlice {
        {"a": 4, "b": 12},
        {"a": 3, "b": 11},
        {"a": 5, "b": 10},
    }
    
    fmt.Println(ls)  // [map[a:4 b:12] map[a:3 b:11] map[a:5 b:10]]
	sort.Strings(ls)
	fmt.Println(ls)  // [map[a:3 b:11] map[a:4 b:12] map[a:5 b:10]]
}
```

**复杂结构体：<font color = orange>` []struct`</font>:**

``` go
type People struct {
    Name string
    Age int
}

type testSlice []People

func (l testSlice) Len() int {
    return len(l)
}

func (l testSlice) Swap(i, j, int) {
    l[i],l[j] = l[j],l[i]
}

func (l testSlice) Less(i, j int) bool {
    retrun l[i].Age < l[j].Age  // 按照年龄对应得值排序
}

func main() {
    ls := testSlice {
        {Name: "n1", Age: 12},
        {Name: "n2", Age: 11},
        {Name: "n3", Age: 10},
    }
    
    fmt.Println(ls)  // [{n1 12} {n2 11} {n3 10}]
	sort.Strings(ls)
	fmt.Println(ls)  // [{n3 10} {n2 11} {n1 12}]
}
```

# golang标准库time

time包提供测量和显示时间得功能。

## 基本使用

打印显示出现在的时间，基本示例如下。

其中now为<font color = orange>` time.Time`</font>类型，Month为<font color = orange>` time.Month`</font>类型

``` go
func test1() {
    now := time.Now() // 获取当前时间
    // current time:2020-12-01 22:24:30.85736 +0800 CST m=+0.000096031
    fmt.Printf("current time:%v\n", now)
    year := now.Year()      // 年
    month := now.Month()    // 月
    day := now.Day()        // 日
    hour := now.Hour()      // 小时
    minute := now.Minute    // 分钟
    second := now.Second    // 秒
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,minute,second)
    fmt.Printf("%T,%T,%T,%T,%T,%T,%T\n",now,year,month,day,hour,minute,second)
}
```

## 时间戳

在编程中对于时间戳的应用也尤为广泛，例如在Web开发中做cookies有效期，接口加密，Redis中的key有效期等等，大部分都是使用了时间戳。

时间戳是自1970年1月1日（08:00:00GMT）至当前时间的总毫秒数。它也被称为Unix时间戳（UnixTimestamp）。在Golang中，获取时间戳的操作如下：

``` go
func test2() {
    now := time.Now()
    // 当前时间戳 TimeStamp type:int64, TimeStamp:1606832965
    fmt.Printf("TimeStamp type:%T,TimeStamp:%v",now.Unix(),now.Unix())
}
```

除此之外还有纳秒时间戳，我们可以使用<font color = orange>` time.Now().UnixNano()来获取它`</font>

``` go
func test3() {
    now := time.Now()
    // 纳秒级时间戳TimeStamp type:int64,TimeStamp:1606833059999670000
    fmt.Printf("TimeStamp type:%T,TimeStamp:%v",now.UnixNano(),now.UnixNano())
}
```

**时间戳转化为普通的时间格式**

在<font color = orange>` go`</font>语言中可以<font color = orange>` time.Unix`</font>来直接将时间戳转化为当前时间格式，实现瞬间替换。

``` go
func timeStampToTime() {
    timestamp := time.Now().Unix()
    timeObj := time.Unix(timestamp, 0) // 将时间戳转为时间格式
    fmt.Println(timeObj)
    year := timeObj.Year()      // 年
    month := timeObj.Month()    // 月
    day := timeObj.Day()        // 日
    hour := timeObj.Hour()      // 小时
    minute := timeObj.Minute    // 分钟
    second := timeObj.Second    // 秒
    fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,minute,second)
}
```

## 操作时间

**ADD**

``` go
package main

import (
	"fmt"
    "time"
)

func add(h,m,s,mls,ns time.Duration) {
    now := time.Now()
    fmt.Println(now.Add(time.Hour*h + time.Minute*m + time.Second*s + time.Millisecond*mls + time.Microsecond*msc + time.Nanosecond*ns))
}

func main() {
    test4(3,4,5,6,7,8)
}
```

> 注意在这里不能增加年\月\日，仅能增加时分秒，也就是以下的才被允许

``` go
const {
    Nanosecond Duration = 1
    Microsecond         = 1000 * Nanosecond
    Millisecond         = 1000 * Microsecond
    Second              = 1000 * Millisecond
    Minute              = 60 * Second
    Hour                = 60 * Minute
}
```

**Sub**

``` go
func sub() {
    now := time.Now()
    targetTime := now.Add(time.Hour)
    // 目标时间于此时相比相差1h0m0s
    fmt.Println(targetTime.Sub(now))
}
```

> 谁的sub谁为参照时间

**Equal**

``` go
func (t Time) Equal(u Time) bool
```

判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确的比较。

**Before**

``` go
func (t Time) Before(u Time) bool
```

如果t代表的时间点在u之前，返回真；否则返回假。

**After**

``` go
func (t Time) After(u Time) bool
```

如果t代表的时间点在u之后，返回真；否则返回假。

## 定时器

使用<font color = orange>` time.Tick(时间间隔)`</font>来设置定时器，定时器的本质上是一个通道（channel）。

``` go
func tick() {
    ticker := time.Tick(time.Second) // 定义一个1秒间隔的定时器
    for i := range ticker {
        fmt.Println(i) //每秒都会执行的任务
    }
}
```

## 时间格式化

时间类型有一个自带的方法<font color = orange>` Format`</font>进行格式化，需要注意的是Go语言中格式化时间模板不是常见的<font color = orange>` Y-m-d H:M:S`</font>而是使用Go诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）

``` go
func format() {
    now := time.Now()
    // 格式化的模板为Go的出生时间2006年1月2日15点04分 Mon Jan
    // 24小时制
    fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
    // 12小时制
    fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
    fmt.Println(now.Format("15:04 2006/01/02"))
    fmt.Println(now.Format("2006/01/02"))
}
```

> 补充：如果像格式化为12小时方式，需指定<font color = orange>` PM`</font>

## 解析字符串格式的时间

``` go
now := time.Now()
fmt.Println(now)
// 加载时区
loc,err := time.Location("Asia/Shanghai")
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(timeObj)
fmt.Println(timeObj.Sub(now))
```

# golang标准库encoding/json

这个包可以实现json的编码和解码，就是将json字符串转为<font color =orange>` struct`</font>，或者将<font color =orange>` struct`</font>转为json。

## 核心的两个函数

``` go
func Marshal(v interface{}) ([]byte, error)
```

> 将struct编码成json，可以接收任意类型

``` go
func Unmarshal(data []byte, v interface{}) error
```

> 将json转码成struct结构体

## 两个核心结构体

``` go
type Decoder struct {
    // contains filtered or unexported fields
}
```

> 从输入流读取并解析json

``` go
type Encoder struct {
    // contains filtered or unexported fields
}
```

> 写json到输出流

## 实例演示

**结构体转换为json**

``` go
type Person struct {
    Name string
    Age int
    Email string
}

func Marshal() {
    p := Person{
        Name: "tom",
        Age: 20,
        Email: "tom@gmail.com",
    }
    b, _ := json.Marshal(p)
    fmt.Printf("b: %v/n", string(b))
}
```

**json转换为结构体**

``` go
func Unmarshal() {
    b1 := []byte(`{"Name":"tom","Age":20,"Email":"tom@qq.com"}`)
    var m Person
    json.Unmarshal(b1, &m)
    fmt.Printf("m: %v\n", m)
}
```

**解析嵌套类型**

``` go
// 解析嵌套类型
func test3() {
    b := []byte(`{"Name":"tom","Age":20,"Email":"tom@qq.com","Parents":["tom","kite"]}`)
    var f interface{}
    json.Unmarshal(b, &f)
    fmt.Printf("f: %v\n", f)
}
```

**解析嵌套引用类型**

``` go
func test4() {
    type Person struct {
        Name string
        Age int
        Email string
        Parent []string
    }
    
    p := Person {
        Name: "tom",
        Age: 20,
        Email: "tom@gmail.com",
        Parent: []string{"big tom","big kite"},
    }
    
    b, _ := json.Marshal(p)
    fmt.Printf("b: %v/n", string(b))
}
```

**io流Reader Writer可以扩展到http websocket等场景**

``` go
func test5() {
    // dec := json.NewDecoder(os.Stdin)
    // a.json : {{"Name":"tom","Age":20,"Email":"tom@qq.com","Parents":["tom","kite"]}}
    f, _ = os.Open("a.json")
    dec := json.NewDecoder(f)
    enc := json.NewEncoder(os.Stdout)
    for {
        var v map[string]interface{}
        if err := dec.Decode(&v); err != nil {
            log.Println(err)
            return
        }
        fmt.Printf("v: %v\n", v)
        if err := enc.Encode(&v); err != nil {
             log.Println(err)
        }
    }
}
```

# golang标准库encoding/xml

xml包实现xml解析

## 核心的两个函数

``` go
func Marshal(v interface{}) ([]byte, error)
```

> 将struct编码成xml，可以接收任意类型

``` go
func Unmarshal(data []byte, v interface{}) error
```

> 将xml转码成struct结构体

## 两个核心结构体

``` go
type Decoder struct {
    ...
}
```

> 从输入流读取并解析xml

``` go
type Encoder struct {
    // contains filtered or unexported fields
}
```

> 写xml到输出流

## 实例演示

**结构体转为xml**

``` go
type Person struct {
    XMLName xml.Name `xml:"person"`   
    Name    string   `xml:"game"`
    Age     int      `xml:"age"`
    Email   string   `xml:"email"`
}

func Marshal() {
    p := Person {
        Name: "tom",
        Age: 20,
        Email: "tom@gmail.com",
    }
    // b, _ := xml.Marshal(p)
    b, _ := xml.MarshalIndent(p," ","  ")
    fmt.Printf("%v\n", string(b))
}
```

**也可以读写文件**

``` go
func read() {
    /* 
    	<?xml version="1.0" encoding="UTF-8"?>
    	<person>
    		<name>tom</name>
    		<age>20</age>
    		<email>tom@gmail.com</email>
    	</person>
    */
    b, _ := ioutil.ReadFile("a.xml")
    var p Person
    xml.Unmarshal(b, &p)
    fmt.Printf("p: %p\n", p)
}

func write() {
    type Person struct {
        XMLName xml.Name `xml:"person"`   
    	Name    string   `xml:"game"`
    	Age     int      `xml:"age"`
    	Email   string   `xml:"email"`
    }
    
    p := Person {
        Name: "tom",
        Age: 20,
        Email: "tom@gmail.com",
    }
    
    f, _ := os.OpenFile("a.xml", os.OWRONLY, 0777)
    defer f.Close()
    e := xml.NewEncoder(f)
    e.Encode(p)
}
```

# golang标准库math

该包包含了一些常量和一些有用的数学计算函数，例如：三角函数、随机数、绝对值、平方根等。

## 常量

``` go
fmt.Printf("float64的最大值是：%.f\n", math.MaxFloat64)
fmt.Printf("float64的最小值是：%.f\n", math.SmallestNonzeroFloat64)
fmt.Printf("float32的最大值是：%.f\n", math.MaxFloat32)
fmt.Printf("float32的最小值是：%.f\n", math.SmallestNonzeroFloat32)
fmt.Printf("Int8的最大值是：%d\n", math.MaxInt8)
fmt.Printf("Int8的最小值是：%d\n", math.MinInt8)
fmt.Printf("UInt8的最大值是：%d\n", math.MaxUInt8)
fmt.Printf("Int16的最大值是：%d\n", math.MaxInt8)
fmt.Printf("Int16的最小值是：%d\n", math.MinInt16)
fmt.Printf("Int32的最大值是：%d\n", math.MaxInt32)
fmt.Printf("Int32的最小值是：%d\n", math.MinInt32)
fmt.Printf("UInt32的最大值是：%d\n", math.MaxUInt32)
fmt.Printf("Int64的最大值是：%d\n", math.MaxInt64)
fmt.Printf("Int64的最小值是：%d\n", math.MinInt64)
fmt.Printf("圆周率默认为.200f\n", math.Pi)
```

## 常用函数

``` go
package main

import(
	"fmt"
    "math"
)

func main() {
    /* 
    	取绝对值，函数签名如下:
    	func Abs(x float64) float64
    */
    fmt.Printf("[-3.14]的绝对值为:[%.2f]\n",math.Abs(-3.14))
    
    /* 
    	取x的y次方，函数签名如下:
    	func Pow(x, y float64) float64
    */
    fmt.Printf("[2]的16次方为:[%.f]\n",math.Pow(2, 16))
    
    /* 
    	10的x次方，函数签名如下:
    	func Pow10(n int) float64
    */
    fmt.Printf("10的[3]次方为:[%.f]\n",math.Pow10(3))
    
    /* 
    	取x开平方，函数签名如下:
    	func Sqrt(x float64) float64
    */
    fmt.Printf("[64]的开平方为:[%.f]\n",math.Sqrt(64))
    
    /* 
    	取x开立方，函数签名如下:
    	func Cbrt(x float64) float64
    */
    fmt.Printf("[27]的开立方为:[%.f]\n",math.Cbrt(27))
    
     /* 
    	向上取整，函数签名如下:
    	func Ceil(x float64) float64
    */
    fmt.Printf("[3.14]向上取整为:[%.f]\n",math.Ceil(3.14))
    
    /* 
    	向下取整，函数签名如下:
    	func Floor(x float64) float64
    */
    fmt.Printf("[8.75]向下取整为:[%.f]\n",math.Floor(8.75))
    
     /* 
    	取余数，函数签名如下:
    	func Mod(x float64) float64
    */
    fmt.Printf("[10/3]的余数为:[%.f]\n",math.Mod(10,3))
    
     /* 
    	分别取整数和小数部分，函数签名如下:
    	func Modf(f, float64) (int float64, frac float64)
    */
    Integer, Decimal := math.Modf(3.13159265358979)
    fmt.Printf("[3.13159265358979]的整数部分为:[%.f],小数部分为:[%.14f]\n",Integer,Decimal)
}
```

## 随机数

``` go
package main

import(
	"fmt"
    "math/rand"
    "time"
)

func init() {
    // 以时间作为初始化种子
    rand.Seed(time.Now().UnixNano()) 
}

func main() {
    for i := 0; i < 10; i++ {
        a := rand.Int()
        fmt.Println(a)
    } 
    fmt.Println("---------")
    for i := 0; i < 10; i++ {
        a := rand.Intn(100)
        fmt.Println(a)
    } 
    fmt.Println("---------")
    for i := 0; i < 10; i++ {
        a := rand.Float32()
        fmt.Println(a)
    } 
}
```

# golang操作mysql数据库

## 准备数据库和表

**下载安装Mysql**

``` go
https://dev.mysql.com/downloads/mysql/
```

**创建一个go_db数据库**

``` mysql
create database go_db
```

**打开数据库**

``` mysql
use go_db
```

**创建表**

``` mysql
CREATE TABLE user_tbl {
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(20),
    PASSWORD VARCHAR(20)
}
```

**添加模拟数据**

``` mysql
INSERT INTO user_tbl(username,PASSWORD) VALUES ("tom","123")
INSERT INTO user_tbl(username,PASSWORD) VALUES ("kite","456")
```

## 安装配置mysql驱动

**安装驱动**

``` go
go get -u github.com/go-sql-driver/mysql
```

**初始化模块**

``` go
go mod init m
```

**执行go mod tidy**

``` go
go mod tidy
```

**导入驱动**

``` go
package main

import (
	"fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    
    db, err := sql.Open("mysql", "user:password@/dbname")
    if err != nil {
        panic(err)
    }
    // See "Important settings" section
    db.SetConnMaxLifetime(time.Minute * 3)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(10)
}
```

## 获得数据库连接

**导入包**

``` go
package main

import (
	"fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)
```

**获得连接**

``` go
package main

import (
	"time"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    
    db, err := sql.Open("mysql", "root:123456@/go_dbs")
    if err != nil {
        panic(err)
    }
    print(db)
    // 最大连接时长
    db.SetConnMaxLifetime(time.Minute * 3)
    // 最大连接数
    db.SetMaxOpenConns(10)
    // 空闲连接数
    db.SetMaxIdleConns(10)
}
```

**初始化连接**

Open函数可能只是验证器参数格式是否正确，实际上并不创建于数据库的连接。如果要检查数据源的名称是否真实有效，应该调用Ping方法。

返回的DB对象可以安全的被多个goroutine并发使用，并且维护其自己的空闲连接池。因此，Open函数应该仅被调用一次，很少需要关闭这个DB对象。

``` go
package main

import (
	"time"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

// 定义一个全局对象db
var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error){
    
    dsn := "root:root@tcp(127.0.0.1:3306)/golang_login?					 charset=utf8mb4&parseTime=True&loc=Local"
    // 不会校验账号密码是否正确
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        return err
    }
    // 尝试和数据库建立连接(校验dsn是否正确)
    err = db.Ping()
    if err != nil {
        return err
    }
    return nil
}

func main() {
    err := initDB() // 调用输出化数据库的函数
    if err != nil {
        fmt.Printf("初始化失败!err: %v\n", err)
        return
    } else {
        fmt.Printf("初始化成功!)
    }
}
```

## 插入数据

插入、更新和删除操作都使用<font color =orange>` Exec`</font>方法。

``` go
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
```

``` go
// 插入数据
func insertData() {
    sqlStr := "insert into user_tbl(username,password) values(?,?)"
    ret,err := db.Exec(sql,"张三","zs123")
    if err != nil {
        fmt.Printf("insert failed, err:%v\n",err)
        return
    }
    theID,err := ret.LastInsertId()  // 新插入数据的id
    if err != nil {
        fmt.Printf("get lastinsert ID failed, err:%v\n",err)
        return
    }
    fmt.Printf("insert success, the id is %d\n",TheID)
}
```

## 查询操作

### 单行查询

单行查询<font color = orange>` db.QueryRow()`</font>执行一次查询，并期望返回最多一行结果（即Row）。QueryRow总是返回非nil的值，知道返回值的Scan方法被调用时，才会返回被延迟的错误。

**定义一个结构体**

``` go
type user struct {
    id int
    username string
    password string
}
```

**查询**

``` go
// 查询一条用户数据
func queryRowDemo() {
    sqlStr = "select id,username,password from user_tbl where id=?"
    var u user
    // 确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
    err := db.QueryRow(sqlStr, 1).Scan(&u.id,&u.username,&u.password)
    if err != nil {
        fmt.Printf("scan failed, err:%v\n", err)
        return
    }
    fmt.Printf("id: %d name: %s age: %s\n", u.id,u.username,u.password)
}
```

### 多行查询

多行查询<font color = orange>` db.Query()`</font>执行一次查询，返回多行结果（即Rows），一般用于执行select命令。参数args表示query中的占位参数。

``` go
// 查询多条数据示例
func queryMultiRow() {
    sqlStr := "select id,username,password from user_tbl where id > ?"
    rows,err := db.Query(sqlStr, 0)
    if err != nil {
        fmt.Printf("query failed, err:%v\n",err)
        return
    }
    // 非常重要，关闭rows释放持有的数据库链接
    defer rows.Close()
    
    // 循环读取结果集中的数据
    for rows.Next() {
        var u user
        err := rows.Scan(&u.id,&u.username,&u.password)
        if err != nil {
            fmt.Printf("scan failed, err:%v\n",err)
            return
        }
        fmt.Printf("id:%d username:%s password:%s\n",u.id,u.username,u.password)
    }
}
```

## 更新数据

插入、更新和删除操作都使用<font color = orange>` Exec`</font>方法。

``` go
func (d *DB) Exec(query string, args ...interface{}) (Result error)
```

**更新**

``` go
func updateData() {
    sql := "update user_tbl set username=?,password=? where id=?"
    ret,err := db.Exec(sql,"kite2","kite123","2")
    if err != nil {
        fmt.Printf("更新失败, err: %v\n", err)
        return
    }
    rows,err := ret.RowsAffected()
    if err != nil {
         fmt.Printf("更新行失败, err: %v\n", err)
         return
    }
    fmt.Printf("更新成功,更新的行数: %d\n", rows)
}
```

## 删除数据

插入、更新和删除操作都使用<font color = orange>` Exec`</font>方法。

``` go
func (d *DB) Exec(query string, args ...interface{}) (Result error)
```

**删除**

``` go
func updateData() {
    sql := "delete from user_tbl where id = ?"
    ret,err := db.Exec(sql,"1")
    if err != nil {
        fmt.Printf("删除失败, err: %v\n", err)
        return
    }
    rows,err := ret.RowsAffected()
    if err != nil {
         fmt.Printf("删除行失败, err: %v\n", err)
         return
    }
    fmt.Printf("删除成功,删除的行数: %d\n", rows)
}
```

# golang泛型

从1.18版本开始，Go添加了对泛型的支持，即类型参数

## 泛型函数

没有泛型的痛点

1. 对于相近类型的操作，可能会写多个一模一样的函数
2. 不停的类型转换

``` go
package main

import "fmt"
// 遍历int切片
func PrintIntSlice(slice []int) {
	for _, v := range slice {
		fmt.Printf("%T  %v\n", v, v)
	}
}
// 遍历int64切片
func PrintInt64Slice(slice []int64) {
	for _, v := range slice {
		fmt.Printf("%T  %v\n", v, v)
	}
}
// int64切片转int切片
func Int64SliceToIntSlice(Int64Slice []int64) (IntSlice []int) {
	for _, v := range Int64Slice {
		IntSlice = append(IntSlice, int(v))
	}
	return IntSlice
}

func main() {
	PrintIntSlice([]int{1, 2, 3, 4, 5})
	in64Slice := []int64{4, 5, 7}
	println(in64Slice)
	intSlice := Int64SliceToIntSlice(in64Slice)
	fmt.Println(intSlice)
}
```

有了泛型函数，就很方便

``` go
package main

import "fmt"

// 泛型函数
func PrintSlice[T int | int64 | string](slice []T) {
	fmt.Printf("%T\n", slice)
	for _, v := range slice {
		fmt.Printf("%T  %v\n", v, v)
	}
}

func main() {
    
	PrintSlice([]int{1, 2, 3, 4, 5})
	PrintSlice([]int64{1, 2, 3})
	PrintSlice([]string{"hello world"})
    
    // 标准写法
    PrintSlice[int]([]int{1, 2, 3, 4, 5})
	PrintSlice[int64]([]int64{1, 2, 3})
	PrintSlice[string]([]string{"hello world"})
}
```

## 泛型切片

``` go
package main

import "fmt"

type MySlice[T int | string] []T

func PrintSlice[T int | int64 | string](slice []T) {
	fmt.Printf("%T\n", slice)
	for _, v := range slice {
		fmt.Printf("%T  %v\n", v, v)
	}
}

func main() {
	v1 := MySlice[int]{1, 2, 3, 4}
	PrintSlice(v1)
}

```

## 泛型Map

``` go
package main

import "fmt"

type MyMap[k int | string, v any] map[k]v

type User struct {
	Name string
}

func main() {
	v1 := MyMap[string, string]{
		"key": "humanising",
	}
	fmt.Println(v1)

	v2 := MyMap[int, User]{
		1: User{
			Name: "你好",
		},
	}
	fmt.Println(v2)
}

```

## 泛型约束

### 约束参数

``` go
package main

import "fmt"

type NumStr interface {
	Num | Str
}

// ~的意思就是底层数据类型
type Num interface {
	~int | ~int32 | ~int64
}

type Str interface {
	string
}

type mySlice[T NumStr] []T

func main() {
	m1 := mySlice[int]{1, 2, 3}
	fmt.Println(m1)
	m2 := mySlice[int64]{1, 2, 3}
	fmt.Println(m2)
	m3 := mySlice[string]{"hello"}
	fmt.Println(m3)
}

```

### 约束方法

``` go
package main

import (
	"fmt"
	"strconv"
)

type Price int

func (p Price) String() string {
	// int转数字
	return strconv.Itoa(int(p))
}

type Price2 string

func (p Price2) String() string {
	// int转数字
	return string(p)
}

type showPrice interface {
	~int | ~string
	String() string
}

func showPriceFunc[T showPrice](p T) {
	fmt.Println(p.String())
}

func main() {
	var p1 Price = 12
	showPriceFunc(p1)
	var p2 Price2 = "56"
	showPriceFunc(p2)
}
```
































































































































































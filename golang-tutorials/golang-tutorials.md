# Go入门教程

[Go 入门教程](https://laravelacademy.org/books/golang-tutorials)



## 1 开篇：为什么学习 Go 语言

 2009

静态的、强类型的、编译型编程语言，为并发而生

### 与C/C++对比

Go 设计的初衷是替代 C，所以二者有很多相似之处，但 Go 做的更多：

- 提供了==自动管理线程和垃圾回收的运行时==，在 C/C++ 中，需要自行管理线程和内存
- 更快的编译速度

适用场景不同：

- C/C++可用于高性能嵌入式系统、大型云应用以及桌面程序开发
- Go适用于系统和云平台开发

Go**不适用于高性能嵌入式系统**，因为嵌入式系统资源有限，而Go运行时调度线程和垃圾回收需要额外的开销。至今没有提供 GUI SDK，所以也**不适用于桌面程序开发**。

### 与Java对比

![](images/image-15829504024192.jpg)

Java 程序编译之后需要安装额外的 Java runtime 运行，Java 程序的可移植性依赖 Java runtime，Go 不需要，Go运行时已经包含在这个编译的二进制文件中了，这体现在部署上的区别就是需要在服务器安装 Java runtime，而 Go 只需要部署单文件即可。

另外就是程序具体执行的时候，Go 被编译成二进制文件被所在操作系统执行，而 Java 通常是在包含了 JIT 编译器的 JVM 中执行，JIT 会对代码进行优化。

### 对比Python/PHP

Python/PHP 都是动态语言，而 Go 是静态语言，会做类型检查，可靠性更高。

开发 Web 应用时，Python/PHP 通常躲在 Nginx/Apache后面作为后台进程，Go则提供了内置的Web服务器，完全可以直接在生产环境使用。

![](images/image-15829509535807.jpg)

Python/PHP 之所以要借助额外的 Web 服务器是因为**对并发请求的处理**，Python 有一个全局锁同时只允许运行一个线程，PHP 本身就没有多线程多进程机制，一次请求从头到位都是一个独立的进程，为了让基于 Python/PHP 的 Web 应用支持并发请求，必须借助外部 Web 服务器。

而 Go 内置的 Web 服务器充分利用了 `goroutine`，对并发连接有很好的支持。此外，由于协程的本质是**在同一个进程中调度不同线程**，所以还支持共享资源。

另外就是 Python/PHP 作为动态语言，性能不如 Go，如果要提升 Python/PHP 性能，必须通过 C 语言编写扩展，复杂度和学习成本太高。

### 对比JavaScript

这里的 JavaScript，主要是 Node.js。

JavaScript 是单线程模型，尽管异步 IO 机制可以使用不同的线程，主程序还是以单线程模式运行的，主程序代码耗时会阻塞其他代码的执行。

而 Go 语言的多线程模型可以通过运行时管理调度协程在多个处理器的不同线程中运行，可以充分利用系统硬件。

Node.js 使用 Google Chrome 的 V8 引擎，其中包含了带有 JIT 编译器的虚拟机，可以对 JavaScript 代码进行优化来提升性能，而 Go 代码直接被编译成机器码执行，没有类似的东西，也无此必要。

### 学习路线图

[官方文档](https://golang.org/doc/)

![](images/a6ca8872092aa01127dfd345a587dbf2.png)

## 2 入门

### 第一个Go程序

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
```

和 Java 类似，Go 使用包作为基本单位来管理代码（可以类比为 PHP 中的命名空间），每个 Go 源代码文件的开头都是一个 `package` 声明，表示该文件中 Go 代码所属的包。包是 Go 语言里最基本的分发单位，也是工程管理中依赖关系的体现。

要生成 Go 可执行程序，==必须==声明一个名为 `main` 的包，并且在该包中包含一个名为 `main()` 的主函数，该函数是 Go 可执行程序的==执行起点==，这一点和 C 语言和 Java 语言很像，后续编译 Go 项目程序的时候也要从包含 `main` 包的文件开始。Go 语言的 `main()` 函数不能带参数，也不能定义返回值。

在包声明之后，是一系列的 `import` 语句，用于导入该程序所依赖的包（可类比为 PHP 中通过 `use` 引入其它命名空间的类来理解）。由于本示例程序用到了 `Println()` 函数，所以需要导入该函数所属的 `fmt` 包。

有一点需要注意，与 Java 和 PHP 不同，在 Go 语言中，**不得包含在源代码文件中没有用到的包**，否则 Go 编译器会报编译错误。这与下面的强制函数左花括号 `{` 的放置位置以及之后会提到的函数名的大小写规则，均体现了 Go 语言**在语言层面解决软件工程问题的设计哲学**。

所有 Go 函数（包括在面向对象编程中会提到的类型成员函数）都以关键字 `func` 开头（这一点与 PHP、Java、JavaScript 等语言通过 `function` 定义函数不同）。另外在 Go 函数中，左花括号 `{` 必须在函数定义行的末尾，不能另起一行，否则 Go 编译器会报编译错误。

#### 编译&运行程序

 `go build` 命令对 Go 程序进行编译，然后直接运行编译后的可执行文件执行 Go 程序代码：

```sh
➜ go build hello.go 
➜ ./hello 
Hello, world!
```

 `go run` 命令将编译和执行指令合二为一:

```sh
go run hello.go 
Hello, world!
```



### Go项目工程管理示例 🔖



### 单元测试、问题定位及代码调试

#### 单元测试

单元测试文件默认以同一目录下文件名后缀 `_test` 作为标识。

```go
package simplemath

import "testing"

func TestAdd(t *testing.T) {
    r := Add(1, 2)
    if r != 3 {
        t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
    }
}
```

引入 [testing](https://golang.org/pkg/testing/) 包，类似PHP的PHPUnit或Java的JUnit。

#### 问题定位与调试

- 打印

```go
fval := 110.48 
ival := 200 
sval := "This is a string. " 
fmt.Println("The value of fval is", fval) 
fmt.Printf("fval=%f, ival=%d, sval=%s\n", fval, ival, sval) 
fmt.Printf("fval=%v, ival=%v, sval=%v\n", fval, ival, sval)
```

- 日志

 [log](https://golang.org/pkg/log/) 包



- IDE调试



- GDB调试





## 3 数据类型篇

### 3.1 变量、作用域、常量和枚举

#### 变量使用入门

##### 变量声明和命名规则

```go
var v1 int            // 整型
var v2 string         // 字符串
var (
    v1 int 
    v2 string
）

var v3 bool           // 布尔型
var v4 [10]int        // 数组，数组元素类型为整型
var v5 struct {       // 结构体，成员变量 f 的类型为64位浮点型
    f float64
} 
var v6 *int           // 指针，指向整型
var v7 map[string]int   // map（字典），key为字符串类型，value为整型
var v8 func(a int) int  // 函数，参数类型为整型，返回值类型为整型
```

变量在声明之后，系统会自动将变量值初始化为对应类型的零值。

**驼峰命名法**

##### 变量初始化

```go
var v1 int = 10   // 方式一，常规的初始化操作
var v2 = 10       // 方式二，此时变量类型会被编译器自动推导出来
v3 := 10          // 方式三，可以省略 var，编译器可以自动推导出v3的类型
```

纯粹的变量声明时可不能省略类型，那样会编译器会报错。

推导类型是在编译期做的，而不是运行时。

在 `:=` 运算符左侧的变量应该是未声明过的。

##### 变量赋值与多重赋值



```go
i, j = j, i
```

##### 匿名变量

在使用传统的强类型语言编程时，经常会出现这种情况，即在调用函数时为了获取一个值，却因为该函数返回多个值而不得不定义一堆没用的变量。

GO中使用多重赋值和匿名变量来避免这种丑陋的写法：

```go
_, nickName := GetName()
```



#### 变量的作用域

如果一个变量在函数体外声明，则被认为是全局变量，可以在整个包甚至外部包（变量名以大写字母开头）使用，不管你声明在哪个源文件里或在哪个源文件里调用该变量。

在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，函数的参数和返回值变量也是局部变量。



#### 常量使用入门

##### 常量定义

```go
const Pi float64 = 3.14159265358979323846 
const zero = 0.0 // 无类型浮点常量 
const (          // 通过一个 const 关键字定义多个常量，和 var 类似
    size int64 = 1024
    eof = -1  // 无类型整型常量 
) 
const u, v float32 = 0, 3  // u = 0.0, v = 3.0，常量的多重赋值 
const a, b, c = 3, 4, "foo" // a = 3, b = 4, c = "foo", 无类型整型和字符串常量
```

##### 预定义常量

`true`

`false` 

`iota` 一个可被编译器修改的常量，在每一个 `const` 关键字出现时被重置为 0，然后在下一个 `const` 出现之前，每出现一次 `iota`，其所代表的数字会自动增 1。

```go
package main

const (    // iota 被重置为 0
    c0 = iota   // c0 = 0
    c1 = iota   // c1 = 1
    c2 = iota   // c2 = 2
)

const (
    u = iota * 2;  // u = 0
    v = iota * 2;  // v = 2
    w = iota * 2;  // w = 4
)

const x = iota;  // x = 0
const y = iota;  // y = 0
```

如果两个 `const` 的赋值语句的表达式是一样的，那么还可以省略后一个赋值表达式。因此，上面的前两个 `const` 语句可简写为：

```go
const ( 
    c0 = iota 
    c1 
    c2 
)

const ( 
    u = iota * 2 
    v 
    w 
)
```

##### 枚举

Go 语言并不支持其他语言用于表示枚举的 `enum` 关键字，而是通过在 `const` 后跟一对圆括号定义一组常量的方式来实现枚举。

```go
const (
    Sunday = iota 
    Monday 
    Tuesday 
    Wednesday 
    Thursday 
    Friday 
    Saturday 
    numberOfDays
)
```

#### 常量的作用域

和函数体外声明的变量一样，以大写字母开头的常量在包外可见（类似于 `public` 修饰的类属性）；

以小写字母开头的常量只能在包内访问（类似于通过 `protected` 修饰的类属性）；

函数体内声明的常量只能在函数体内生效。



### 3.2 Go支持的数据类型概述及布尔类型

Go 语言内置对以下这些基本数据类型的支持：

- 布尔类型：bool
- 整型：int8、byte、int16、int、uint、uintptr 等
- 浮点类型：float32、float64
- 复数类型：complex64、complex128
- 字符串：string
- 字符类型：rune
- 错误类型：error

此外，Go 语言还支持以下这些复合类型：

- 指针（pointer）
- 数组（array）
- 切片（slice）
- 字典（map）
- 通道（chan） 【主要用于并发编程时不同协程之间的通信】
- 结构体（struct）
- 接口（interface）

结构体类似于面向对象编程语言中的类（class），Go 沿用了 C 语言的这一复合类型，而没有像传统面向对象编程那样引入单独的类概念。



### 3.3 整型及运算符

#### 整型

| 类型      | 长度（单位：字节） | 说明                                 | 值范围                                   | 默认值 |
| --------- | ------------------ | ------------------------------------ | ---------------------------------------- | ------ |
| `int8`    | 1                  | 带符号8位整型                        | -128~127                                 | 0      |
| `uint8`   | 1                  | 无符号8位整型，与 `byte` 类型等价    | 0~255                                    | 0      |
| `int16`   | 2                  | 带符号16位整型                       | -32768~32767                             | 0      |
| `uint16`  | 2                  | 无符号16位整型                       | 0~65535                                  | 0      |
| `int32`   | 4                  | 带符号32位整型，与 `rune` 类型等价   | -2147483648~2147483647                   | 0      |
| `uint32`  | 4                  | 无符号32位整型                       | 0~4294967295                             | 0      |
| `int64`   | 8                  | 带符号64位整型                       | -9223372036854775808~9223372036854775807 | 0      |
| `uint64`  | 8                  | 无符号64位整型                       | 0~18446744073709551615                   | 0      |
| `int`     | 32位或64位         | 与具体平台相关                       | 与具体平台相关                           | 0      |
| `uint`    | 32位或64位         | 与具体平台相关                       | 与具体平台相关                           | 0      |
| `uintptr` | 与对应指针相同     | 无符号整型，足以存储指针值的未解释位 | 32位平台下为4字节，64位平台下为8字节     | 0      |

和其他编程语言一样，可以通过增加前缀 `0` 来表示八进制数（如：077），增加前缀 `0x` 来表示十六进制数（如：0xFF），以及使用 `E` 来表示 10 的连乘（如：`1E3 = 1000`）。





#### 运算符

`+`、`-`、`*`、`/` 和 `%`（取余运算只能用于整数）

在 Go 语言中，也支持自增/自减运算符，即 `++`/`--`，但是只能作为语句，不能作为表达式，且只能用作后缀，不能放到变量前面。

支持 `+=`、`-=`、`*=`、`/=`、`%=` 



 `>`、`<`、`==`、`>=`、`<=` 和 `!=`，比较运算符运行的结果是布尔值。



| 位运算符 | 含义     | 结果                                           |
| -------- | -------- | ---------------------------------------------- |
| `x & y`  | 按位与   | 把 x 和 y 都为 1 的位设为 1                    |
| `x | y`  | 按位或   | 把 x 或 y 为 1 的位设为 1                      |
| `x ^ y`  | 按位异或 | 把 x 和 y 一个为 1 一个为 0 的位设为 1         |
| `^x`     | 按位取反 | 把 x 中为 0 的位设为 1，为 1 的位设为 0        |
| `x << y` | 左移     | 把 x 中的位向左移动 y 次，每次移动相当于乘以 2 |
| `x >> y` | 右移     | 把 x 中的位向右移动 y 次，每次移动相当于除以 2 |



| 逻辑运算符 | 含义                | 结果                                                   |
| ---------- | ------------------- | ------------------------------------------------------ |
| `x && y`   | 逻辑与运算符（AND） | 如果 x 和 y 都是 true，则结果为 true，否则结果为 false |
| `x || y`   | 逻辑或运算符（OR）  | 如果 x 或 y 是 true，则结果为 true，否则结果为 false   |
| `!x`       | 逻辑非运算符（NOT） | 如果 x 为 true，则结果为 false，否则结果为 true        |



##### 运算符优先级

由上到下表示优先级从高到低，或者数字越大，优先级越高：

```go
6      ^（按位取反） !
5      *  /  %  <<  >>  &  &^
4      +  -  |  ^（按位异或）
3      ==  !=  <  <=  >  >=
2      &&
1      ||
```



### 3.4 浮点型与复数类型

#### 浮点型

Go 语言中的浮点数采用[IEEE-754](https://zh.wikipedia.org/zh-hans/IEEE_754) 标准的表达方式，定义了两个类型：`float32` 和 `float64`，其中 `float32` 是单精度浮点数，可以精确到小数点后 7 位（类似 PHP、Java 等语言的 `float` 类型），`float64` 是双精度浮点数，可以精确到小数点后 15 位（类似 PHP、Java 等语言的 `double` 类型）。

```go
var floatValue1 float32

floatValue1 = 10
floatValue2 := 10.0 // 如果不加小数点，floatValue2 会被推导为整型而不是浮点型
floatValue3 := 1.1E-10
```

对于浮点类型需要被自动推导的变量，其类型将被自动设置为 `float64`，而不管赋值给它的数字是否是用 32 位长度表示的。

```go
// 编译报错 ：cannot use floatValue2 (type float64) as type float32 in assignment
floatValue1 = floatValue2  // floatValue2 是 float64 类型
```

必须强制转换：

```go
floatValue1 = float32(floatValue2)
```

> 在实际开发中，应该尽可能地使用 `float64` 类型，因为 [math](https://golang.org/pkg/math/) 包中所有有关数学运算的函数都会要求接收这个类型。

##### 浮点数的比较

浮点数支持通过算术运算符进行四则运算，也支持通过比较运算符进行比较（前提是运算符两边的操作数类型一致），但是涉及到相等的比较除外，因为我们上面提到，看起来相等的两个十进制浮点数，**在底层转化为二进制时会丢失精度**，因此不能被表象蒙蔽。

替代的解决方案：

```go
p := 0.00001
// 判断 floatValue1 与 floatValue2 是否相等
if math.Dim(float64(floatValue1), floatValue2) < p {
    fmt.Println("floatValue1 和 floatValue2 相等")
} 
```

一种近似判断，通过一个可以接受的最小误差值 `p`，约定如果两个浮点数的差值在此精度的误差范围之内，则判定这两个浮点数相等。这个解决方案**也是其他语言判断浮点数相等所采用的通用方案**。

#### 复数类型

把整型和浮点型这种日常比较常见的数字称为实数，复数是实数的延伸，可以通过两个实数（在计算机中用浮点数表示）构成，一个表示实部（real），一个表示虚部（imag），常见的表达形式如下：

```go
z = a + bi
```

其中 a、b 均为实数，i 称为虚数单位，当 b = 0 时，z 就是常见的实数，当 a = 0 而 b ≠ 0 时，将 z 称之为**纯虚数**。



在 Go 语言中，复数支持两种类型：`complex64`（32 位实部和虚部） 和 `complex128`（64 位实部与虚部）。



### 3.5 字符串及底层字符类型

#### 字符串

##### 1️⃣基本使用

在 Go 语言中，字符串是一种基本类型，默认是通过 UTF-8编码的字符序列，当字符为 ASCII 码时则占用 1 个字节，其它字符根据需要占用 2-4 个字节，比如中文编码通常需要 3 个字节。

###### 声明和初始化

###### 格式化输出

###### 转义字符

###### 多行字符串

```go
results := `Search results for "Golang":
- Go
- Golang
Golang Programming
`
fmt.Printf("%s", results)
```



##### 2️⃣不可变值类型

在 Go 语言中，字符串是一种不可变值类型

```go
str := "Hello world"
str[0] = 'X' // 编译错误
```



##### 3️⃣字符编码

如果你的 Go 代码需要包含非 ANSI 字符，保存源文件时请注意编码格式必须选择 UTF-8。特别是在 Windows 下一般编辑器都默认保存为本地编码，比如中国地区可能是 GBK 编码而不是 UTF-8。

字符串的编码转换是处理文本文档（比如 TXT、XML、HTML 等）时非常常见的需求，不过 Go 语言默认仅支持 UTF-8 和 Unicode 编码，对于其他编码，Go 语言标准库并没有内置的编码转换支持。



##### 4️⃣字符串操作



###### 字符串连接



###### 字符串切片

**==左闭右开==**



> 此外 Go 字符串也支持字符串比较、是否包含指定字符/子串、获取指定子串索引位置、字符串替换、大小写转换、trim 等操作，请参考标准库 [strings](https://golang.org/pkg/strings/) 包。

##### 5️⃣字符串遍历

1. 以**字节数组**的方式遍历：

```go
str := "Hello, 世界" 
n := len(str) 
for i := 0; i < n; i++ {
    ch := str[i]    // 依据下标取字符串中的字符，ch 类型为 byte
    fmt.Println(i, ch) 
}
```

结果长度为13，是字节`byte`

2. 以 **Unicode 字符**遍历：

```go
str := "Hello, 世界" 
for i, ch := range str { 
    fmt.Println(i, ch)    // ch 的类型为 rune 
}
```

结果长度为9，是字符`rune`



#### 底层字符类型

Go 语言对字符串中的单个字符进行了单独的类型支持，在 Go 语言中支持两种字符类型：

- 一种是 `byte`，代表 UTF-8 编码中单个字节的值（它也是 `uint8` 类型的别名，两者是等价的，因为正好占据 1 个字节的内存空间）；
- 另一种是 `rune`，代表单个 Unicode 字符（它也是 `int32` 类型的别名，因为正好占据 4 个字节的内存空间。关于 `rune` 相关的操作，可查阅 Go 标准库的 [unicode](https://golang.org/pkg/unicode/) 包）。

##### UTF-8 和 Unicode 的区别



##### 将 Unicode 编码转化为可打印字符



### 3.6 基本数据类型之间的转化

由于 Go 是强类型语言，所以不支持动态语言那种自动转化，而是要对变量进行强制类型转化。

#### 数值类型之间的转化

##### 整型之间的转化

```go
v1 := uint(16)   // 初始化 v1 类型为 unit
v2 := int8(v1)   // 将 v1 转化为 int8 类型并赋值给 v2
v3 := uint16(v2) // 将 v2 转化为 uint16 类型并赋值给 v3
```



##### 原码、反码和补码 🔖



##### 整型与浮点型之间的转化

```go
v1 := 99.99
v2 := int(v1)  // v2 = 99  小数位被丢弃
```



#### 字符串和其他基本类型之间的转化

##### 将整型转化为字符串

整型数据可以通过 Unicode 字符集转化为对应的 UTF-8 编码的字符串：

```go
v1 := 65
v2 := string(v1)  // v2 = A

v3 := 30028
v4 := string(v3)  // v4 = 界
```

还可以将 `byte` 数组或者 `rune` 数组转化为字符串，因为字符串底层就是通过这两个基本字符类型构建的：

```go
v1 := []byte{'h', 'e', 'l', 'l', 'o'}
v2 := string(v1)  // v2 = hello

v3 := []rune{0x5b66, 0x9662, 0x541b}
v4 := string(v3)  // v4 = 学院君
```

`byte` 是 `uint8` 的别名，`rune` 是 `int32` 的别名，所以也可以看做是整型数组和字符串之间的转化。

##### strconv包

Go 语言默认不支持将字符串类型强制转化为数值类型，即使字符串中包含数字也不行。

如果要实现更强大的基本数据类型与字符串之间的转化，可以使用 Go 官方 `strconv` 包提供的函数：

```go
v1 := "100"
v2, _ := strconv.Atoi(v1)  // 将字符串转化为整型，v2 = 100

v3 := 100
v4 := strconv.Itoa(v3)   // 将整型转化为字符串, v4 = "100"

v5 := "true"
v6, _ := strconv.ParseBool(v5)  // 将字符串转化为布尔型
v5 = strconv.FormatBool(v6)  // 将布尔值转化为字符串

v7 := "100"
v8, _ := strconv.ParseInt(v7, 10, 64)   // 将字符串转化为整型，第二个参数表示进制，第三个参数表示最大位数
v7 = strconv.FormatInt(v8, 10)   // 将整型转化为字符串，第二个参数表示进制

v9, _ := strconv.ParseUint(v7, 10, 64)   // 将字符串转化为无符号整型，参数含义同 ParseInt
v7 = strconv.FormatUint(v9, 10)  // 将无符号整数型转化为字符串，参数含义同 FormatInt

v10 := "99.99"
v11, _ := strconv.ParseFloat(v10, 64)   // 将字符串转化为浮点型，第二个参数表示精度
v10 = strconv.FormatFloat(v11, 'E', -1, 64)

q := strconv.Quote("Hello, 世界")    // 为字符串加引号
q = strconv.QuoteToASCII("Hello, 世界")  // 将字符串转化为 ASCII 编码
```



### 3.7 数组使用入门及其不足

#### 数组的声明和初始化

在 Go 语言中，数组是固定长度的、同一类型的数据集合。数组中包含的每个数据项被称为数组元素，一个数组包含的元素个数被称为数组的长度。

```go
[capacity]data_type{element_values}
```



```go
var a [8]byte // 长度为8的数组，每个元素为一个字节
var b [3][3]int // 二维数组（9宫格）
var c [3][3][3]float64 // 三维数组（立体的9宫格）
var d = [3]int{1, 2, 3}  // 声明时初始化
var e = new([3]string)   // 通过 new 初始化

a := [5]int{1,2,3,4,5}
a := [...]int{1, 2, 3}  // 语法糖，省略数组长度的声明
a := [5]int{1, 2, 3}  // 如果没有填满，则空位会通过对应的元素类型零值填充
a := [5]int{1: 3, 3: 7} // 还可以初始化指定下标位置的元素值，未设置的位置也会以对应元素类型的零值填充
```

数组长度在声明后就不可更改，在声明时可以指定数组长度为一个常量或者一个常量表达式（常量表达式是指在编译期即可计算结果的表达式）。数组的长度是该数组类型的一个内置常量，可以用 Go 语言的内置函数 `len()` 来获取

#### 数组元素的访问和设置



#### 遍历数组



Go 语言还提供了一个关键字 `range`，用于以更优雅的方式遍历数组中的元素：

```go
for i, v := range arr { 
    fmt.Println("Element", i, "of arr is", v) 
}
```

```go
for _, v := range arr {
   // ...
}
for i := range arr {
   // ...
}
```

#### 多维数组

```go
// 通过二维数组生成九九乘法表
var multi [9][9]string
for j := 0; j < 9; j++ {
    for i := 0; i < 9; i++ {
            n1 := i + 1
            n2 := j + 1
            if n1 < n2 {  // 摒除重复的记录
                continue
            }
            multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1 * n2)
        }
  }

// 打印九九乘法表
for _, v1 := range multi {
    for _, v2 := range v1 {
        fmt.Printf("%-8s", v2)  // 位宽为8，左对齐
    }
    fmt.Println()
}
```



#### 数组类型的不足

- 长度固定

- 值类型

切片类型，一个引用类型的、支持动态添加元素的新「数组」类型。

>  在Go语言中很少使用数组，大多数时候会使用切片取代它。

### 3.8 切片

#### 1️⃣切片的定义

在 Go 语言中，切片是一个新的数据类型，与数组最大的不同在于，切片的类型字面量中只有元素的类型，没有长度：

```go
var slice []string = []string{"a", "b", "c"}
```

![](images/image-20.jpeg)

切片是一个可变长度的、同一类型元素集合，切片的长度可以随着元素数量的增长而增长（但**不会随着元素数量的减少而减少**），不过切片从底层管理上来看依然使用数组来管理元素，可以看作是**对数组做了一层简单的封装**。基于数组，切片添加了一系列管理功能，可以随时动态扩充存储空间。



#### 2️⃣创建切片

##### 基于数组

切片可以只使用数组的一部分元素或者整个数组来创建，甚至可以创建一个比所基于的数组还要大的切片：

```go
// 先定义一个数组
months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

// 基于数组创建切片
q2 := months[3:6]    // 第二季度
summer := months[5:8]  // 夏季

fmt.Println(q2)
fmt.Println(summer)  
```

Go 语言支持通过 `array[start:end]` 这样的方式基于数组生成一个切片，`start` 表示切片在数组中的下标起点，`end` 表示切片在数组中的下标终点，两者之间的元素就是切片初始化后的元素集合。和字符串切片一样，这也是个**左闭右开**的集合。

切片底层引用了一个数组，由三个部分构成 —— **指针、长度和容量**，指针指向数组起始下标，长度对应切片中元素的个数，容量则是切片起始位置到底层数组结尾的位置：

![](images/16118499891425.jpg)

切片长度不能超过容量。内置函数 `len` 获取切片的长度， `cap` 函数获取切片容量

##### 基于切片

```go
firsthalf := months[:6]
q1 := firsthalf[:3] // 基于 firsthalf 的前 3 个元素构建新切片
```



```go
q1 := firsthalf[:9]
// q1结果 [January February March April May June July August September]
```

因为 `firsthalf` 的容量是 12，只要选择的范围不超过 `firsthalf` 的容量，那么这个创建操作就是合法的，所以虽然是基于切片创建切片，但本质上还是基于数组。



##### 直接创建

内置函数 `make()` 可以用于灵活地创建切片。

```go
mySlice1 := make([]int, 5)  // 初始长度为 5 
mySlice2 := make([]int, 5, 10)  // 初始长度为 5、容量为 10
mySlice3 := []int{1, 2, 3, 4, 5}  // 长度和容量均为5
```

#### 3️⃣遍历切片



#### 4️⃣动态增加元素

一个切片的容量初始值根据创建方式的不同而不同：

- 对于基于数组和切片创建的切片而言，默认容量是从切片起始索引到对应底层数组的结尾索引；
- 对于通过内置 `make` 函数创建的切片而言，在没有指定容量参数的情况下，默认容量和切片长度一致。

```go
var oldSlice = make([]int, 5, 10)  // [0 0 0 0 0]

newSlice := append(oldSlice, 1, 2, 3)  // [0 0 0 0 0 1 2 3]
```

```go
appendSlice := []int{1, 2, 3, 4, 5}
newSlice := append(oldSlice, appendSlice...)  // 注意末尾的 ... 不能省略
```



##### 自动扩容

如果追加的元素个数超出 `oldSlice` 的默认容量，则底层会自动进行扩容。



##### 内容复制

内置函数 `copy()`，用于将元素从一个切片复制到另一个切片。如果两个切片不一样大，就会按其中较小的那个切片的元素个数进行复制。

```go
slice1 := []int{1, 2, 3, 4, 5} 
slice2 := []int{5, 4, 3}

// 复制 slice1 到 slice 2
copy(slice2, slice1) // 只会复制 slice1 的前3个元素到 slice2 中
// slice2 结果: [1, 2, 3]
// 复制 slice2 到 slice 1
copy(slice1, slice2) // 只会复制 slice2 的 3 个元素到 slice1 的前 3 个位置
// slice1 结果：[5, 4, 3, 4, 5]
```





#### 5️⃣动态删除元素

```go
slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
slice3 = slice3[:len(slice3) - 5]  // 删除 slice3 尾部 5 个元素
slice3 = slice3[5:]  // 删除 slice3 头部 5 个元素
```



```go
slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
slice4 := append(slice3[:0], slice3[3:]...)  // 删除开头三个元素
slice5 := append(slice3[:1], slice3[4:]...)  // 删除中间三个元素
slice6 := append(slice3[:0], slice3[:7]...)  // 删除最后三个元素
slice7 := slice3[:copy(slice3, slice3[3:])]  // 删除开头前三个元素
```



> 关于 Go 切片元素的动态插入、新增、删除操作，[go切片图示](https://ueokande.github.io/go-slice-tricks/)



#### 6️⃣数据共享问题 🔖

切片底层是基于数组实现的，对应的结构体对象：

```go
type slice struct {
    array unsafe.Pointer //指向存放数据的数组指针
    len   int            //长度有多大
    cap   int            //容量有多大
}
```



解决方案



### 3.9 字典

#### 字典声明和初始化

需要在声明时指定键和值的类型。

Go字典是个无序集合，底层不会按照元素添加顺序维护元素的存储顺序。

```go
var testMap map[string]int
testMap = map[string]int{
  "one": 1,
  "two": 2,
  "three": 3,
}

k := "two"
v, ok := testMap[k]
if ok {
  fmt.Printf("The element of key %q: %d\n", k, v)
} else {
  fmt.Println("Not found!")
}
```

还可以像切片一样，通过`make()`初始化：

```go
var testMap = make(map[string]int)

testMap = make(map[string]int, 100)   // 指定该字典的初始存储容量（超出会自动扩容）
```

#### 使用入门

```go
// 赋值
testMap["four"] = 4

// 查找元素，第二个返回布尔值用于判断查找是否成功
value, ok := testMap["one"] 
if ok { // 找到了
  // 处理找到的value 
}

//删除元素
delete(testMap, "four")

// 遍历字典
for key, value := range testMap {
    fmt.Println(key, value)
}

// 键值对调
invMap := make(map[int] string, 3)
for k, v := range testMap {
    invMap[v] = k
}

```



> 🔖
>
> Go 语言中的字典和 Redis 一样，底层也是通过哈希表实现的，添加键值对到字典时，实际是将键转化为哈希值进行存储，在查找时，也是先将键转化为哈希值去哈希表中查询，从而提高性能。
>
> 但是哈希表存在哈希冲突问题，即不同的键可能会计算出同样的哈希值，这个时候 Go 底层还会判断原始键的值是否相等。也正因如此，我们在声明字典的键类型时，要求数据类型必须是支持通过 == 或 != 进行判等操作的类型，比如数字类型、字符串类型、数组类型、结构体类型等，不过为了提高字典查询性能，类型长度越短约好，通常，我们会将其设置为整型或者长度较短的字符串类型。

#### 字典排序

字典是一个无序集合，如果你想要对字典进行排序，可以通过分别为字典的键和值创建切片，然后通过对切片进行排序来实现。

##### 按照键进行排序

```go
keys := make([]string, 0)
for k, _ := range testMap {
    keys = append(keys, k)
}

sort.Strings(keys)  // 对键进行排序

fmt.Println("Sorted map by key:")
for _, k := range keys {
    fmt.Println(k, testMap[k])
}
```

##### 按照值进行排序

```go
values := make([]int, 0)
for _, v := range testMap {
    values = append(values, v)
}

sort.Ints(values)   // 对值进行排序

fmt.Println("Sorted map by value:")
for _, v := range values  {
    fmt.Println(invMap[v], v)
}
```



### 3.10 指针 🔖



**变量**的本质对一块内存空间的命名，我们可以通过引用变量名来使用这块内存空间存储的值，而指针则是用来指向这些变量值所在**内存地址的值**。



#### unsafe.Pointer





## 4 流程控制篇

- 条件语句：用于条件判断，对应的关键字有 `if`、`else` 和 `else if`；
- 分支语句：用于分支选择，对应的关键字有 `switch`、`case`、`fallthrough`和 `select`（用于通道，后面介绍协程时会提到）；
- 循环语句：用于循环迭代，对应的关键字有 `for` 和 `range`；
- 跳转语句：用于代码跳转，对应的关键字有 `goto`。

### 条件语句

```go
// if
if condition { 
    // do something 
}

// if...else...
if condition { 
    // do something 
} else {
    // do something 
}

// if...else if...else...
if condition1 { 
    // do something 
} else if condition2 {
    // do something else 
} else {
    // catch-all or default 
}
```

- 条件语句不需要使用圆括号将条件包含起来 `()`；
- 无论语句体内有几条语句，花括号 `{}` 都是必须存在的；
- 左花括号 `{` 必须与 `if` 或者 `else` 处于同一行；
- 在 `if` 之后，条件语句之前，可以添加变量初始化语句，使用 `;` 间隔，比如上述代码可以这么写 `if score := 100; score > 90 {`

### 分支语句

```go
switch var1 {
    case val1:
        ... 
    case val2:
        ... 
    default:
        ...
}
```



```go
score := 100
switch {
case score >= 90:
    fmt.Println("Grade: A")
case score >= 80 && score < 90:
    fmt.Println("Grade: B")
case score >= 70 && score < 80:
    fmt.Println("Grade: C")
case score >= 60 && score < 70:
    fmt.Println("Grade: D")
default:
    fmt.Println("Grade: F")
}
```

只有在与 `case` 分支值判等的时候，才能把变量 `score` 放到 `switch` 关键字后面：

```go
score := 100
switch score {
case 90, 100:
    fmt.Println("Grade: A")
case 80:
    fmt.Println("Grade: B")
case 70:
    fmt.Println("Grade: C")
case 60:
case 65:
    fmt.Println("Grade: D")
default:
    fmt.Println("Grade: F")
}
```

合并分支：

```go
score := 60
switch score {
...
case 60:
    fallthrough
case 65:
    fmt.Println("Grade: D")
...
}
```



- 和条件语句一样，左花括号 `{` 必须与 `switch` 处于同一行；
- 单个 `case` 中，可以出现多个结果选项（通过逗号分隔）；
- 与其它语言不同，Go 语言不需要用 `break` 来明确退出一个 `case`；
- 只有在 `case` 中明确添加 `fallthrough` 关键字，才会继续执行紧跟的下一个 `case`；
- 可以不设定 `switch` 之后的条件表达式，在这种情况下，整个 switch 结构与多个 `if...else...` 的逻辑作用等同。



### 循环语句

```go
sum := 0 
for i := 1; i <= 100; i++ { 
    sum += i 
}
fmt.Println(sum)
```



无限循环：

```go
sum := 0
i := 0
for {
    i++
    if i > 100 {
        break
    }
    sum += i
}
fmt.Println(sum)
```



`for` 循环条件表达式中也支持多重赋值:

```go
a := []int{1, 2, 3, 4, 5, 6} 
for i, j := 0, len(a) – 1; i < j; i, j = i + 1, j – 1 { 
    a[i], a[j] = a[j], a[i] 
}
fmt.Println(a)  // [6 5 4 3 2 1]
```



for-range 结构:

```go
for k, v := range a {
    fmt.Println(k, v)
}
```

可用于可迭代的集合（数组、切片、字典）。





基于条件判断进行循环

```go
sum := 0
i := 0
for i < 100 {
    i++
    sum += i
}
fmt.Println(sum)
```



- 和条件语句、分支语句一样，左花括号 `{` 必须与 `for` 处于同一行；
- 不支持 `whie` 和 `do-while` 结构的循环语句；
- 可以通过 `for-range` 结构对可迭代集合进行遍历；
- 支持基于条件判断进行循环迭代；
- 允许在循环条件中定义和初始化变量，且支持多重赋值；
- Go 语言的 `for` 循环同样支持 `continue` 和 `break` 来控制循环，但是它提供了一个更高级的 `break`，可以选择中断哪一个循环



### 跳转语句





## 5 函数式编程篇

### 5.1 函数使用入门和常用内置函数

在 Go 语言中，函数主要有三种类型：

- 普通函数
- 匿名函数（闭包）
- 类方法

#### 函数定义

关键字 `func`、函数名、参数列表、返回值、函数体和返回语句

![](images/image.png)

如果函数的参数列表中包含若干个类型相同的参数，则可以在参数列表中省略前面变量的类型声明，只保留最后一个:

```go
func add(a, b int) int { 
    // ...
}
```



#### 函数调用





**在调用其他包定义的函数时，只有函数名首字母大写的函数才可以被访问**。

Go 语言中没有 `public`、`protected`、`private` 之类的关键字，它是通过首字母的大小写来区分可见性的：首字母小写的函数只能在同一个包中访问，首字母大写的函数才可以在其他包中调用，Go 文件中定义的全局变量也是如此。



#### 系统内置函数

[builtin package - builtin - Go Packages](https://pkg.go.dev/builtin)

| 名称                      | 说明                                                         |
| ------------------------- | ------------------------------------------------------------ |
| `close`                   | 用于在管道通信中关闭一个管道                                 |
| `len`、`cap`              | `len` 用于返回某个类型的长度（字符串、数组、切片、字典和管道），`cap` 则是容量的意思，用于返回某个类型的最大容量（只能用于数组、切片和管道） |
| `new`、`make`             | `new` 和 `make` 均用于分配内存，`new` 用于值类型和用户自定义的类型（类），`make` 用于内置引用类型（切片、字典和管道）。它们在使用时将类型作为参数：`new(type)`、`make(type)`。`new(T)` 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针，可以用于基本类型：`v := new(int)`。`make(T)` 返回类型 T 的初始化之后的值，所以 `make` 不仅分配内存地址还会初始化对应类型。 |
| `copy`、`append`          | 分别用于切片的复制和动态添加元素                             |
| `panic`、`recover`        | 两者均用于错误处理机制                                       |
| `print`、`println`        | 打印函数，在实际开发中建议使用 [fmt](https://golang.org/pkg/fmt/) 包 |
| `complex`、`real`、`imag` | 用于复数类型的创建和操作                                     |



### 5.2 参数传递、变长参数与多返回值



#### 1️⃣传递参数

按值传参

引用传参

#### 2️⃣变长参数

所谓变长参数指的是函数参数的数量不确定，可以按照需要传递任意数量的参数到指定函数，合适地使用变长参数，可以让代码更简洁，尤其是输入输出类函数，比如打印函数 `fmt.Printf` 的参数就是典型的变长参数。

##### 基本定义和传值

只需要在参数类型前加上 `...` 前缀，就可以将该参数声明为变长参数：

```go
func myfunc(numbers ...int) {
    for _, number := range numbers {
        fmt.Println(number)
    }
}
```

这种变长参数还支持传递一个 `[]int` 类型的切片，传递切片时需要在末尾加上 `...` 作为标识，表示对应的参数类型是变长参数：

```go
slice := []int{1, 2, 3, 4, 5}
myfunc(slice...)
myfunc(slice[1:3]...)
```

> 注：形如 `...type` 格式的类型只能作为函数的参数类型存在，并且必须是函数的最后一个参数。

之所以支持传入切片，是因为从底层实现原理上看，类型 `...type` 本质上是一个切片，也就是 `[]type`，这也是为什么上面的参数 `numbers` 可以用 `for` 循环来获取每个传入的参数值。



##### 任意类型的变长参数（泛型）

指定变长参数类型为 `interface{}`：

```go
func Printf(format string, args ...interface{}) { 
    // ...
}
```

自定义一个支持任意类型的变长参数函数：

```go
func myPrintf(args ...interface{}) {
    for _, arg := range args {
        switch reflect.TypeOf(arg).Kind() {
        case reflect.Int:
            fmt.Println(arg, "is an int value.")
        case reflect.String:
            fmt.Printf("\"%s\" is a string value.\n", arg)
        case reflect.Array:
            fmt.Println(arg, "is an array type.")
        default:
            fmt.Println(arg, "is an unknown type.")
        }
    }
}

func main() {
    myPrintf(1, "1", [1]int{1}, true)
}
```

其实这里就是一个**泛型**功能，Go 语言并没有在语法层面提供对泛型的支持，所以目前只能自己这样通过反射和 `interface{}` 类型实现。

`interface{}` 是一个**空接口**，可以用于表示任意类型。



#### 3️⃣多返回值

```go
func add(a, b *int) (int, error) {
    if *a < 0 || *b < 0 {
        err := errors.New("只支持非负整数相加")
        return 0, err
    }
    *a *= 2
    *b *= 3
    return *a + *b, nil
}

func main()  {
    x, y := -1, 2
    z, err := add(&x, &y)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Printf("add(%d, %d) = %d\n", x, y, z)
}
```



##### 命名返回值

在设置多返回值时，还可以对返回值进行变量命名，这样，我们就可以在函数中直接对返回值变量进行赋值，而不必每次都按照指定的返回值格式返回多个变量了：

```go
func add(a, b *int) (c int, err error) {
    if *a < 0 || *b < 0 {
        err = errors.New("只支持非负整数相加")
        return
    }
    *a *= 2
    *b *= 3
    c = *a + *b
    return
}
```



### 5.3 匿名函数与闭包

#### 1️⃣匿名函数的定义和使用

==匿名函数==是一种没有指定函数名的函数声明方式（与之相对的，有名字的函数被称为==具名函数==）。

```go
// 1、将匿名函数赋值给变量
add := func(a, b int) int {
    return a + b
}

// 调用匿名函数 add
fmt.Println(add(1, 2))  

// 2、定义时直接调用匿名函数
func(a, b int) {
    fmt.Println(a + b)
} (1, 2) 
```

#### 2️⃣匿名函数与闭包

所谓闭包指的是引用了自由变量（未绑定到特定对象的变量，通常在函数外定义）的函数，被引用的自由变量将和这个函数一同存在，即使已经离开了创造它的上下文环境也不会被释放（比如传递到其他函数或对象中）。简单来说，「闭」的意思是「封闭外部状态」，**即使外部状态已经失效，闭包内部依然保留了一份从外部引用的变量**。

闭包只能通过匿名函数实现，我们可以把闭包看作是**有状态的匿名函数**，反过来，如果匿名函数引用了外部变量，就形成了一个闭包（Closure）。

支持闭包的语言都将函数作为**第一类对象**（firt-class object，有的地方也译作第一级对象、一等公民等，都是一个意思），Go语言也不例外，这意味 Go 函数和普通Go数据类型（整型、字符串、数组、切片、字典、结构体等）具有同等的地位，可以赋值给变量，也可以作为参数传递给其他函数，还能够被函数动态创建和返回。

> 注：所谓第一类对象指的是运行期可以被创建并作为参数传递给其他函数或赋值给变量的实体，在绝大多数语言中，数值和基本类型都是第一类对象，在支持闭包的编程语言中（比如 Go、PHP、JavaScript、Python 等），函数也是第一类对象，而像 C、C++ 等不支持匿名函数的语言中，函数不能在运行期创建，所以在这些语言中，函数不是第一类对象。

#### 3️⃣匿名函数的常见使用场景

##### 保证局部变量的安全性



##### 将匿名函数作为函数参数

```go
add := func(a, b int) int {
    return a + b
}

// 将函数类型作为参数
func(call func(int, int) int) {
    fmt.Println(call(1, 2))
}(add)
```



##### 将匿名函数作为函数返回值

```go
// 将函数作为返回值类型
func deferAdd(a, b int) func() int {
    return func() int {
        return a + b
    }
}

func main() {
    ...

    // 此时返回的是匿名函数
    addFunc := deferAdd(1, 2)
    // 这里才会真正执行加法操作
    fmt.Println(addFunc())
}
```





### 5.4 通过高阶函数实现装饰器模式 

#### 高阶函数

高阶函数，就是接收其他函数作为参数传入，或者把其他函数作为结果返回的函数。

#### 装饰器模式

装饰器模式（Decorator）是一种软件设计模式，其应用场景是为某个已经存在的功能模块（类或者函数）添加一些「装饰」功能，而又不会侵入和修改原有的功能模块。就好比我们给房间做节日装饰一样，它不会调整这个房间原有的任何固有框架，而又让房间充满节日气氛。

有过 Python、Java 编程经验的同学应该对这个模式很熟悉，在 Python、Java 中，我们可以通过注解非常优雅地实现装饰器模式，比如给某个功能模块添加日志功能、或者为路由处理器添加中间件功能，这些都可以通过装饰器实现。

不过 Go 语言的设计哲学就是简单，没有提供「注解」之类的语法糖，在函数式编程中，要实现装饰器模式，可以借助高阶函数来实现。

#### 通过高阶函数实现装饰器模式

##### 

```go
package main

import (
	"fmt"
	"time"
)
/*
	通过高阶函数实现装饰器模式
*/

// 为函数类型设置别名提高代码可读性
type MultiPlyFunc func(int, int) int

// 乘法运算函数
func multiply(a, b int) int {
	return a * b
}

// 通过高阶函数在不侵入原有函数实现的前提下计算乘法函数执行时间
func execTime(f MultiPlyFunc) MultiPlyFunc {
	return func(a, b int) int {
		start := time.Now()      // 起始时间
		c := f(a, b)             // 执行乘法运算函数
		end := time.Since(start) // 函数执行完毕耗时
		fmt.Printf("--- 执行耗时: %v ---\n", end)
		return c // 返回计算结果
	}
}

func main() {
	a := 2
	b := 8
	// 通过修饰器调用乘法函数，返回的是一个匿名函数
	decorator := execTime(multiply)
	// 执行修饰器返回函数
	c := decorator(a, b)
	fmt.Printf("%d x %d = %d\n", a, b, c)
}
```

结果：

```sh
$ go run decorator.go 
--- 执行耗时: 166ns ---
2 x 8 = 16

```

核心思路就是在被修饰的功能模块（这里是外部传入的乘法函数 `f`）执行前后加上一些额外的业务逻辑，而又不影响原有功能模块的执行。显然，装饰器模式是遵循 SOLID 设计原则中的开放封闭原则的 —— 对代码扩展开放，对代码修改关闭。

### 5.5 递归函数及性能调优  🔖

#### 递归函数及编写思路

在实际开发过程中，某个问题满足以下条件就可以通过递归函数来解决：

- 一个问题的解可以被拆分成多个子问题的解
- 拆分前的原问题与拆分后的子问题除了数据规模不同，求解思路完全一样
- 子问题存在递归终止条件

#### 通过斐波那契数列求解演示



#### 递归函数性能优化

##### 递归函数执行耗时对比

```go
package main

import (
    "fmt"
    "time"
)

type FibonacciFunc func(int) int

// 通过递归函数实现斐波那契数列
func fibonacci(n int) int {
    // 终止条件
    if n == 1 {
        return 0
    }
    if n == 2 {
        return 1
    }
    // 递归公式
    return fibonacci(n-1) + fibonacci(n-2)
}

// 斐波那契函数执行耗时计算
func fibonacciExecTime(f FibonacciFunc) FibonacciFunc {
    return func(n int) int {
        start := time.Now() // 起始时间
        num := f(n)  // 执行斐波那契函数
        end := time.Since(start) // 函数执行完毕耗时
        fmt.Printf("--- 执行耗时: %v ---\n", end)
        return num  // 返回计算结果
    }
}

func main() {
    n1 := 5
    f := fibonacciExecTime(fibonacci)
    r1 := f(n1)
    fmt.Printf("The %dth number of fibonacci sequence is %d\n", n1, r1)

    n2 := 50
    r2 := f(n2)
    fmt.Printf("The %dth number of fibonacci sequence is %d\n", n2, r2)
}
```

```sh
$ go run recursive.go
-----执行耗时：334ns ------
The 5th number of fibonacci sequence is 3
-----执行耗时：34.0056395s ------
The 50th number of fibonacci sequence is 7778742049

```

> 1s = 10亿ns
>
> 1µs=1000ns



究其原因，一方面是因为递归函数调用产生的额外开销，另一方面是因为目前这种实现存在着重复计算，比如我在计算 `fibonacci(50)` 时，会转化为计算 `fibonacci(49)` 与 `fibonacci(48)` 之和，然后我在计算 `fibonacci(49)` 时，又会转化为调用 `fibonacci(48)` 与 `fibonacci(47)` 之和，这样一来 `fibonacci(48)` 就会两次重复计算，这一重复计算就是一次新的递归（从序号 48 递归到序号 1），以此类推，大量的重复递归计算堆积，最终导致程序执行缓慢。

##### 通过内存缓存技术优化递归函数性能

先对后一个原因进行优化，通过缓存中间计算结果来避免重复计算，从而提升递归函数的性能。

```go
const MAX = 50
var fibs [MAX]int

// 缓存中间结果的递归函数优化版
func fibonacci2(n int) int {
    if n == 1 {
        return 0
    }

    if n == 2 {
        return 1
    }

    index := n - 1
    if fibs[index] != 0 {
        return fibs[index]
    }

    num := fibonacci2(n-1) + fibonacci2(n-2)
    fibs[index] = num
    return num
}
```

```go
func main() {
    n1 := 5
    f1 := fibonacciExecTime(fibonacci)
    r1 := f1(n1)
    fmt.Printf("The %dth number of fibonacci sequence is %d\n", n1, r1)

    n2 := 50
    r2 := f1(n2)
    fmt.Printf("The %dth number of fibonacci sequence is %d\n", n2, r2)

    f2 := fibonacciExecTime(fibonacci2)
    r3 := f2(n2)
    fmt.Printf("The %dth number of fibonacci sequence is %d\n", n2, r3)
}
```

```sh
$ go run recursive.go
-----执行耗时：209ns ------
The 5th number of fibonacci sequence is 3
-----执行耗时：34.569491583s ------
The 50th number of fibonacci sequence is 7778742049
-----执行耗时：584ns ------
The 50th number of fibonacci sequence is 7778742049
```



这种优化是在内存中保存中间结果，所以称之为**==内存缓存技术==**（memoization），这种内存缓存技术在优化计算成本相对昂贵的函数调用时非常有用。

##### 通过尾递归优化递归函数性能

对造成上述递归函数性能低下的第一个原因进行优化。

在计算机科学里，**==尾调用==**是指一个函数的最后一个动作是调用一个函数（只能是一个函数调用，不能有其他操作，比如函数相加、乘以常量等）：

```go
func f(x int) int {
  ...
  return g(x);
}
```

这种情况下称该调用位置为尾位置，若这个函数在尾位置调用自身，则称这种情况为**==尾递归==**，它是尾调用的一种特殊情形。尾调用的一个重要特性是它不是在函数调用栈上添加一个新的堆栈帧 —— 而是更新它，尾递归自然也继承了这一特性，这就使得原来层层递进的调用栈变成了线性结构，因而可以极大优化内存占用，提升程序性能，这就是尾递归优化技术。



🔖

### 5.5 引入 Map-Reduce-Filter 模式处理集合元素



### 5.6 基于管道技术实现函数的流式调用







## 6 面向对象篇

### 6.1 类型系统概述

>  对于面向对象编程的支持，Go 语言的实现可以说是完全颠覆了以往我们对面向对象编程的认知。

简洁之处在于，Go 语言并没有沿袭传统面向对象编程中的诸多概念，比如类的继承、接口的实现、构造函数和析构函数、隐藏的 this 指针等，也没有 `public`、`protected`、`private` 之类的访问修饰符。

优雅之处在于，Go 语言对面向对象编程的支持是语言类型系统中的天然组成部分，整个类型系统通过接口串联，浑然一体。

#### 类型系统概述

类型系统是指一个语言的类型体系结构。一个典型的类型系统通常包含如下基本内容：

- 基本类型，如 `byte`、`int`、`bool`、`float`、`string` 等；
- 复合类型，如数组、切片、字典、指针、结构体等；
- 可以指向任意对象的类型（Any 类型）；
- 值语义和引用语义；
- 面向对象，即所有具备面向对象特征（比如成员方法）的类型；
- 接口。

类型系统描述的是这些内容在一个语言中如何被关联。因为 Java 语言自诞生以来被称为最纯正的面向对象语言，所以我们就先以 Java 语言为例讲一讲类型系统。

#### Java vs Go 类型系统设计



### 6.2 类的定义、初始化和成员方法





### 6.3 通过组合实现类的继承和方法重写





### 6.4 类属性和成员方法的可见性 🔖



### 6.5 接口定义及实现🔖

**如果说 goroutine 和 channel 是支撑起 Go 语言并发模型的基石，那么接口就是 Go 语言整个类型系统的基石**。

#### 传统侵入式接口实现



#### Go 语言的接口实现



#### 通过组合实现接口继承



### 6.6 接口赋值



### 6.7 类型断言





### 6.8 空接口、反射和泛型





## 7 错误处理篇

### 7.1 error 类型及其使用



### 7.2 defer 语句及其使用



### 7.3 panic 和 recover



## 8 网络编程篇

### 8.1 从多进程、多线程到协程



### 8.2 协程实现原理及使用入门



### 8.3 基于共享内存实现协程通信



### 8.4 基于锁和原子操作实现并发安全



### 通道类型



### 利用多核 CPU 实现并行计算





### 通过 context 包提供的函数实现多协程之间的协作



### 临时对象池 sync.Pool





## 9 网络编程篇







## 数据结构和算法篇



---



# Go Web编程

[Go Web 编程](https://laravelacademy.org/books/go-web-programming)



# Gin使用教程

[Gin 使用教程](https://laravelacademy.org/books/gin-tutorial)

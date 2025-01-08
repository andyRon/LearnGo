Go语言更多
------



参考：

[《GO语言从入门到项目实践（超值版）》余建熙 2022](http://www.tup.tsinghua.edu.cn/bookscenter/book_09281301.html)

[《Go语言开发实战（慕课版）》千锋教育高教产品研发部 2020](https://book.douban.com/subject/34956558/)

[《Go语言从入门到项目实战（视频版）》](https://book.douban.com/subject/36049170/) 刘瑜 2022

## 47 Go标准库

参考：

https://www.topgoer.com/%E5%B8%B8%E7%94%A8%E6%A0%87%E5%87%86%E5%BA%93/

### 47.1 fmt包

#### 向外输出

##### Print系列

区别在于Print函数直接输出内容，Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。

```go
func Print(a ...interface{}) (n int, err error)
func Printf(format string, a ...interface{}) (n int, err error)
func Println(a ...interface{}) (n int, err error)
```

##### Fprint系列

Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。

```go
func Fprint(w io.Writer, a ...interface{}) (n int, err error)
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
```

```go
// 向标准输出写入内容
fmt.Fprintln(os.Stdout, "向标准输出写入内容")
fileObj, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
if err != nil {
    fmt.Println("打开文件出错，err:", err)
    return
}
name := "枯藤"
// 向打开的文件句柄中写入内容
fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)
```

注意，只要满足io.Writer接口的类型都支持写入。

##### Sprint系列

Sprint系列函数会把传入的数据生成并返回一个字符串。

```go
func Sprint(a ...interface{}) string
func Sprintf(format string, a ...interface{}) string
func Sprintln(a ...interface{}) string
```

##### Errorf

Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。

```go
func Errorf(format string, a ...interface{}) error
```

通常使用这种方式来自定义错误类型，例如：

```go
err := fmt.Errorf("这是一个错误")
```

#### 格式化占位符

`*printf`系列函数都支持format格式化参数。

按照占位符将被替换的变量类型划分：

| 通用占位符               | 说明                                                         |
| ------------------------ | ------------------------------------------------------------ |
| %v                       | 值的默认格式表示                                             |
| %+v                      | 类似%v，但输出结构体时会添加字段名                           |
| %#v                      | 值的Go语法表示                                               |
| %T                       | 打印值的类型                                                 |
| %%                       | 百分号                                                       |
| **布尔占位符**           |                                                              |
| %t                       | true或false                                                  |
| **整型占位符**           |                                                              |
| %b                       | 表示为二进制                                                 |
| %c                       | 该值对应的unicode码值                                        |
| %d                       | 表示为十进制                                                 |
| %o                       | 表示为八进制                                                 |
| %x                       | 表示为十六进制，使用a-f                                      |
| %X                       | 表示为十六进制，使用A-F                                      |
| %U                       | 表示为Unicode格式：U+1234，等价于”U+%04X”                    |
| %q                       | 该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示 |
| **浮点数与复数占位符**   |                                                              |
| %b                       | 无小数部分、二进制指数的科学计数法，如-123456p-78            |
| %e                       | 科学计数法，如-1234.456e+78                                  |
| %E                       | 科学计数法，如-1234.456E+78                                  |
| %f                       | 有小数部分但无指数部分，如123.456                            |
| %F                       | 等价于%f                                                     |
| %g                       | 根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）       |
| %G                       | 根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）       |
| **字符串和[]byte占位符** |                                                              |
| %s                       | 直接输出字符串或者[]byte                                     |
| %q                       | 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示 |
| %x                       | 每个字节用两字符十六进制数表示（使用a-f                      |
| %X                       | 每个字节用两字符十六进制数表示（使用A-F）                    |
| **指针**                 |                                                              |
| %p                       | 表示为十六进制，并加上前导的0x                               |
|                          |                                                              |
|                          |                                                              |
|                          |                                                              |

##### 宽度标识符

| 占位符 | 说明               |
| ------ | ------------------ |
| %f     | 默认宽度，默认精度 |
| %9f    | 宽度9，默认精度    |
| %.2f   | 默认宽度，精度2    |
| %9.2f  | 宽度9，精度2       |
| %9.f   | 宽度9，精度0       |

示例代码如下：

```go
n := 88.88
fmt.Printf("%f\n", n)
fmt.Printf("%9f\n", n)
fmt.Printf("%.2f\n", n)
fmt.Printf("%9.2f\n", n)
fmt.Printf("%9.f\n", n)
```

输出结果如下：

```
    88.880000
    88.880000
    88.88
        88.88
           89
```

##### 其他falg

| 占位符 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| ’+’    | 总是输出数值的正负号；对%q（%+q）会生成全部是ASCII字符的输出（通过转义）； |
| ’ ‘    | 对数值，正数前加空格而负数前加负号；对字符串采用%x或%X时（% x或% X）会给各打印的字节之间加空格 |
| ’-’    | 在输出右边填充空白而不是默认的左边（即从默认的右对齐切换为左对齐）； |
| ’#’    | 八进制数前加0（%#o），十六进制数前加0x（%#x）或0X（%#X），指针去掉前面的0x（%#p）对%q（%#q），对%U（%#U）会输出空格和单引号括起来的go字面值； |
| ‘0’    | 使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面； |

举个例子：

```go
s := "枯藤"
fmt.Printf("%s\n", s)
fmt.Printf("%5s\n", s)
fmt.Printf("%-5s\n", s)
fmt.Printf("%5.7s\n", s)
fmt.Printf("%-5.7s\n", s)
fmt.Printf("%5.2s\n", s)
fmt.Printf("%05s\n", s)
```

输出结果如下：

```
    枯藤
       枯藤
    枯藤
       枯藤
    枯藤
       枯藤
    000枯藤
```

#### 获取输入

fmt包下有fmt.Scan、fmt.Scanf、fmt.Scanln三个函数，可以在程序运行过程中从标准输入获取用户的输入。

##### fmt.Scan

```go
func Scan(a ...interface{}) (n int, err error)
```

```go
func main() {
    var (
        name    string
        age     int
        married bool
    )
    fmt.Scan(&name, &age, &married)
    fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
```

fmt.Scan从标准输入中扫描用户输入的数据，将以空白符分隔的数据分别存入指定的参数。

##### fmt.Scanf

```go
func Scanf(format string, a ...interface{}) (n int, err error)
```

Scanf从标准输入扫描文本，根据format参数指定的格式去读取由空白符分隔的值保存到传递给本函数的参数中。

```go
func main() {
    var (
        name    string
        age     int
        married bool
    )
    fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &married)
    fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
```

```
$ ./fmt3    
1:andy 2:19 3:true 
扫描结果 name:andy age:19 married:true 

```

fmt.Scanf不同于fmt.Scan简单的以空格作为输入数据的分隔符，fmt.Scanf为输入数据指定了具体的输入内容格式，只有按照格式输入数据才会被扫描并存入对应变量。

##### fmt.Scanln

```go
func Scanln(a ...interface{}) (n int, err error)
```

Scanln类似Scan，它在遇到换行时才停止扫描。最后一个数据后面必须有换行或者到达结束位置。

```go
func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Scanln(&name, &age, &married)
	fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)
}
```

```
./fmt4
李德胜 36 true
扫描结果 name:李德胜 age:36 married:true 
```

##### bufio.NewReader

想完整获取输入的内容，而输入的内容可能包含空格，这种情况下可以使用bufio包来实现。



##### Fscan系列

类似之前三个函数，只不过它们不是从标准输入中读取数据而是从io.Reader中读取数据。

```go
func Fscan(r io.Reader, a ...interface{}) (n int, err error)
func Fscanln(r io.Reader, a ...interface{}) (n int, err error)
func Fscanf(r io.Reader, format string, a ...interface{}) (n int, err error)
```

##### Sscan系列

不是从标准输入中读取数据而是从指定字符串中读取数据。

```go
func Sscan(str string, a ...interface{}) (n int, err error)
func Sscanln(str string, a ...interface{}) (n int, err error)
func Sscanf(str string, format string, a ...interface{}) (n int, err error)
```







### 47.2 time包

![](images/time.png)

#### 时间类型



#### 时间戳



##### 时间间隔



##### 时间操作

```
Add
Sub
Equal
Before
After
```

##### 定时器

使用time.Tick(时间间隔)来设置定时器，定时器的本质上是一个通道（channel）。

```go
func tickDemo() {
    ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
    for i := range ticker {
        fmt.Println(i)//每秒都会执行的任务
    }
}
```

##### 时间格式化

Go语言中格式化时间模板不是常见的Y-m-d H:M:S而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。



### 47.3 flag包

```go
func main() {
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}
```

```
./flag1 a b c d
args[0]=./flag1
args[1]=a
args[2]=b
args[3]=c
args[4]=d
```

#### flag参数类型

flag包支持的命令行参数类型有bool、int、int64、uint、uint64、float float64、string、duration。

| flag参数     | 有效值                                                       |
| ------------ | ------------------------------------------------------------ |
| 字符串flag   | 合法字符串                                                   |
| 整数flag     | 1234、0664、0x1234等类型，也可以是负数。                     |
| 浮点数flag   | 合法浮点数                                                   |
| bool类型flag | 1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False。    |
| 时间段flag   | 任何合法的时间段字符串。如”300ms”、”-1.5h”、”2h45m”。 合法的单位有”ns”、”us” /“µs”、”ms”、”s”、”m”、”h”。 |

#### 两种定义命令行flag参数的方法

##### flag.Type()

`flag.Type(flag名, 默认值, 帮助信息)*Type`

```go
name := flag.String("name", "张三", "姓名")
age := flag.Int("age", 18, "年龄")
married := flag.Bool("married", false, "婚否")
delay := flag.Duration("d", 0, "时间间隔")
```

此时name、age、married、delay均为对应类型的指针。

##### flag.TypeVar()

`flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)`

```go
var name string
var age int
var married bool
var delay time.Duration
flag.StringVar(&name, "name", "张三", "姓名")
flag.IntVar(&age, "age", 18, "年龄")
flag.BoolVar(&married, "married", false, "婚否")
flag.DurationVar(&delay, "d", 0, "时间间隔")
```



#### 其他函数

##### flag.Parse()

定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。

支持的命令行参数格式有以下几种：

- -flag xxx （使用空格，一个-符号）
- --flag xxx （使用空格，两个-符号）
- -flag=xxx （使用等号，一个-符号）
- --flag=xxx （使用等号，两个-符号）

其中，布尔类型的参数必须使用等号的方式指定。

Flag解析在第一个非flag参数（单个”-“不是flag参数）之前停止，或者在终止符”–“之后停止。



- flag.Args() ////返回命令行参数后的其他参数，以[]string类型
- flag.NArg() //返回命令行参数后的其他参数个数
- flag.NFlag() //返回使用的命令行参数个数

#### 使用案例

```go
func main() {
	// //定义命令行参数方式
	var name string
	var age int
	var married bool
	var delay time.Duration
	flag.StringVar(&name, "name", "andyron", "姓名")
	flag.IntVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "married", false, "婚否")
	flag.DurationVar(&delay, "d", 0, "延迟的时间间隔")

	// 解析命令行参数
	flag.Parse()
	fmt.Println(name, age, married, delay)
	// 返回命令行参数后的其他参数
	fmt.Println(flag.Args())
	// 返回命令行参数后的其他参数个数
	fmt.Println(flag.NArg())
	// 返回使用的命令行参数个数
	fmt.Println(flag.NFlag())
}
```



命令行参数使用提示：

```shell
./flag_demo -help
Usage of ./flag_demo:
  -age int
        年龄 (default 18)
  -d duration
        延迟的时间间隔
  -married
        婚否
  -name string
        姓名 (default "andyron")
```

正常使用命令行flag参数：

```shell
$ ./flag_demo -name pprof -age 19 -married=false -d=24s
pprof 19 false 24s
[]
0
4
```

使用非flag命令行参数：

```shell
$ ./flag_demo a c d
andyron 18 false 0s
[a c d]
3
0
```



### 47.4 strings包——字符串处理



#### 检索字符串

![](images/image-20250102222004211.png)

#### 分割字符串

![](images/image-20250102222415128.png)



#### 大小写转换

![](images/image-20250102222439824.png)

#### 修剪字符串

![](images/image-20250102222522187.png)

#### 比较字符串

![](images/image-20250102222556733.png)





### 47.5 strconv包

字符串与其他基本数据类型之间的类型转换。

#### Parse类函数

Parse类函数主要的功能是将字符串转换为其他类型。

![](images/image-20250102222737543.png)



#### Format类函数

Format类函数主要的功能是将其他类型格式化成字符串。

![](images/image-20250102224530130.png)





### 47.6 regexp正则表达式包

#### 正则表达式中主要元字符

![](images/epub_28438052_403.jpeg)

#### regexp包中核心函数及方法

🔖



### 47.7 math包

![](images/image-20250102224620500.png)

### 47.8 随机数

“math/rand”包实现了伪随机数生成器，能够生成整型和浮点型的随机数。使用随机数生成器需要放入种子。可以使用Seed()函数生成一个不确定的种子放入随机数生成器，这样每次运行随机数生成器都会生成不同的序列。如果没有在随机数生成器中放入种子，则默认使用具有确定性状态的种子，此时可以理解为种子的值是一个常数1，即Seed(1)。

#### rand包的核心方法

![](images/image-20250102225648425.png)

#### 获取随机数的几种方式



### 47.9 log



其它

`log/slog`

第三方日志库

[zap](https://github.com/uber-go/zap)

[logrus](https://github.com/appleboy/gorush)

https://github.com/rs/zerolog



### net包

net/http



### 文件操作相关

os包

io/ioutil 包



### 并发相关

**`sync` 包**

**`context` 包**





### 数据结构与算法相关

**`sort` 包**

**`container` 目录下的相关包（如 `list`、`map`、`heap` 等）**





### template



### encoding

数据格式：JSON、XML、MSGPack等



### reflect反射



## 48 反射

参考：《Go语言从入门到项目实战》-10

反射是指在程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入可执行部分。在运行程序时，程序无法获取自身的信息。

支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。

Go语言程序在运行期使用reflect包访问程序的反射信息。

### 48.1　使用反射访问变量

`reflect`包

#### 获取变量的类型

#### 获取变量的值

#### 反射值的非空和有效性判定



### 48.2 反射值的非空和有效性判定



### 48.3 使用反射访问结构体



### 48.4 使用反射修改值



### 48.5 使用反射调用函数



### 48.6 使用反射创建变量



## 49 命令行工具

### 编译命令go build

go build命令常用参数：

```
-v 编译时显示包名
-p n 指定编译时并发的数量（使用n表示），该值默认为CPU 的逻
-a 强制进行重新构建
-n 仅输出编译时执行的所有命令
-x 执行编译并输出编译时执行的所有命令
-race 开启竞态检测
```

### 清理命令go clean

清理所有编译生成的文件，具体包括：

1. 当前目录下生成的与包名或Go源码文件名相同的可执行文件，以及当前目录中的_obj和_test目录中名称为_testmain.go、test.out、build.out、a.out及后缀为.5、.6、.8、.a、.o和.so的文件。这些文件通常是执行go build命令后生成的。
2. 以当前目录下生成的包名加“.test”后缀为名的文件。这些文件通常是执行go test命令后生成的。
3. 工作区中pkg和bin目录的相应归档文件和可执行文件。这些文件通常是执行go install命令后生成的。

go clean命令通常用于使用VCS（版本控制系统，如Git）的团队，在提交代码前运行，以免将编译时生成的临时文件及编译后生成的可执行文件等错误地提交到代码仓库中。

```
-i	清除关联的安装包和可运行文件，这些文件通常是执行 gQ install 命令后生成的
-n	仅输出清理时执行的所有命令
-r	递归清除在 import 中引入的包
-x	执行清理并输出清理时执行的所有命令
-cache	清理缓存，这些缓存文件通常是执行go build命令后生成的
-testcache	清理测试结果
```

### 运行命令go run

执行go run命令时也会编译Go源码文件，但生成的可执行文件被存放在临时目录中，并自动运行这个可执行文件。

```go
func main() {
	fmt.Println(os.Args)
}
```

```sh
go run main.go -color blue
[/var/folders/8k/ntbhdf615p34cflx1_qwv38r0000gn/T/go-build692689547/b001/exe/main -color blue]
```

### 代码格式化命令gofmt



### 编译并安装命令go install



### 获取包命令go get



## 50 文件处理

### 50.1 文件操作

大多数文件操作的函数都是在os包中的，几个目录操作例子：

```go
func Mkdir(name string, perm FileMode) error  // 创建名称为name的目录，权限设置是perm，如0555
func MkdirAll(path string, perm FileMode) error // 根据path创建多级子目录，如zuolan/test1/test2
func Remove(name string) error  // 删除名称为name的目录，当目录下有文件或者其他目录时会出错
func RemoveAll(path string) error  // 根据path删除多级子目录，如果path是单个名称，那么该目录下的子目录全部删除
```

#### 创建文件与查看状态

```go
// 1 新建文件
func Create(name string) (file *File, err Erroe)  // 根据提供的文件名称创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的
func NewFile(fd uintptr, name string) *File  // 根据文件描述符创建相应的文件，返回一个文件对象
// 2 新建文件夹
func MkdirAll(path sring, perm FileMode) eror
// 3 文件/文件夹状态
func Stat(name string) (FileInfo, error)
```

在创建文件夹或者文件时，权限是一次性指定的，后续若要修改文件权限，需要使用其他函数。

判断文件是否存在，可以使用函数os.IsNotExit(err)，这个函数可以通过传入的ert参数判断文件是否存在并返回一个布尔值。



#### 重命名与移动

```go
func Rename(oldpath, newpath string) error
```



#### 打开与关闭

```go
// 只读方式，内部调用OpenFile()
func Open(name string) (file *File, err Error)
// flag是打开的方式，包括只读、读写等
func OpenFile(name string, flag int, perm uint32) (file *File, err Error)
```

flag属性可以单独使用，也可以组合使用:

```go
os.O_CREATE | os. O_APPEND
os.O_CREATE | os. O_TRUNC | os. O_WRONLY
//os. O_RDONLY      //只读
//os. O_WRONLY       //只写
//os. O_RDWR         //读写
//os. O_APPEND       //往文件中添加（Append）
//os. O_CREATE       //如果文件不存在则先创建
//os. O_TRUNC        //文件打开时裁剪文件
//os. O_EXCL         //和O_CREATE一起使用,文件不能存在
//os. O_SYNC         //以同步I/O的方式打开
```



#### 删除与截断

```go
err := os. Remove("test.txt")
if err !=nil{
  log. Fatal(err)
}
```

```go
err := os.Truncate("test.txt", 100)
if err !=nil{
  log. Fatal(err)
}
```

裁剪一个文件到100B，如果文件本来就少于100B，则文件中原始内容得以保留，剩余的字节以null字节填充。

如果文件本来超过100B，则超过的字节会被抛弃，这样总是得到精确的100B的文件。而如果传入0，则会清空文件。

#### 读写文件

读写文件中最常见的操作有复制文件、编辑、跳转、替换等。

##### 1复制文件

`io.Copy(dst Writer, src Reader) (written int64, err error)`

注意：

Create函数执行之后需要Close()函数关闭回收资源。

调用io包中的复制函数之后文件内容并没有真正保存在文件中，而是使用Sync()函数同步之后才真正保存到硬盘中。

##### 2跳转函数

Seek()函数的特点类似于鼠标光标的定位，指定位置之后可以执行复制、剪切、粘贴等操作。

##### 3写入函数

```go
func (file *File) Write(b []byte) (n int, err Error)             //写入byte类型的信息到文件
func (file *File) WriteAt(b []byte, off int64) (n int, err Error)//在指定位置开始写入byte类型的信息
func (file *File) WriteString(s string) (ret int, err Error)     //写入string信息到文件
```



#### 权限控制





#### 文件链接

在Linux系统中肯定会经常遇到硬链接或者软链接之类的文件，对于一个普通文件，它实际上指向了硬盘的一个索引地址。硬链接会创建一个新的指针并且指向同一个地方，硬链接会保持与原文件双向同步，其中一个文件改动，另一个文件也会改动，但只有所有的链接被删除后文件才会被删除（即移动和重命名都不会影响硬链接）。硬链接只在相同的文件系统中才能工作。

软链接和硬链接不一样，它不直接指向硬盘中相同的地方，而是通过名字引用其他文件，它们可以指向不同的文件系统中的不同文件。Windows操作系统不支持软链接。

```go
// 创建一个硬链接
	// 创建后同一个文件内容会有两个文件名,改变一个文件的内容会影响另一个
	// 删除和重命名不会影响另一个
	hardLink := filePath + "_hl"
	err := os.Link(filePath, hardLink)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("创建硬链接")
	// 创建一个软链接
	softLink := filePath + "_sl"
	err = os.Symlink(fileName, softLink)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("创建软链接")

	// Lstat返回一个文件的信息,但是当文件是一个软链接时,它返回软链接的信息,而不是引用的文件的信息
	// Symlink在Windows中不工作
	fileInfo, err := os.Lstat(softLink)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("链接信息: %+v", fileInfo)
	// 改变软链接的拥有者不会影响原始文件
	err = os.Lchown(softLink, os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}
```

os.Lstat()的函数名中可以看出这是一个针对软链接的函数，用于查看软链接自己的属性，使用os.Stat()函数会获取软链接指向的原文件信息。

需要注意软链接和硬链接实现的异同，从上面这两个函数的第一个参数来看，虽然都是oldname，但实际例子中传递给函数的并不是同一个函数，硬链接是filePath，而软链接是fileName，因为硬链接是从项目根目录开始创建硬链接的，而软链接是根据目标文件的相对位置创建软链接的。

### 50.2 XML文件处理

```xml
<?xml version="1.0" encoding="utf-8"?>
<servers version="1">
    <server>
        <serverName>Local_Web</serverName>
        <serverIP>172.0.0.1</serverIP>
    </server>
    <server>
        <serverName> Local_DB</serverName>
        <serverIP>172.0.0.2</serverIP>
    </server>
</servers>
```

#### 解析XML

```go
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}
type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func main() {
	file, err := os.Open("servers.xml")
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Println(v)
}
```

结果：

```sh
$ go run xml.go
{{ servers} 1 [{{ server} Local_Web 172.0.0.1} {{ server}  Local_DB 172.0.0.2}] 
    <server>
        <serverName>Local_Web</serverName>
        <serverIP>172.0.0.1</serverIP>
    </server>
    <server>
        <serverName> Local_DB</serverName>
        <serverIP>172.0.0.2</serverIP>
    </server>
}
```

XML文件解析成对应的struct对象是通过xml包的Unmarshal函数来解析XML文件：

```go
func Unmarshal(data []byte, v interface{}) error
```

data接收的是XML数据流，v是需要输出的结构，定义为interface，目前支持struct、slice和string，xml包内部采用了反射进行数据的映射，所以v中的字段必须是导出的。Unmarshal解析时XML元素和字段是怎样对应起来的呢？这是有一个优先级读取流程的，首先会读取struct tag，如果没有，那么就会读取对应字段名。必须注意的一点是，解析的时候，tag、字段名、XML元素都是大小写敏感的，所以，字段必须一一对应。

Go语言的反射机制，可以利用这些tag信息将来自XML文件中的数据反射成对应的struct对象。

解析XML到struct时需要遵循以下规则：

1. 如果struct的一个字段是string或者[]byte类型，且它的tag含有",innerxml"，Unmarshal会将此字段所对应的元素内所有内嵌的原始XML累加到此字段上，如上面例子中的Description定义，最后的输出如下：

   ```xml
   		<server>
           <serverName>Local_Web</serverName>
           <serverIP>172.0.0.1</serverIP>
       </server>
       <server>
           <serverName> Local_DB</serverName>
           <serverIP>172.0.0.2</serverIP>
       </server>
   ```

2. 如果struct中有一个名为XMLName，且类型为xml.Name的字段，那么在解析时就会保存这个element的名字到该字段，如上面例子中的servers。

3. 如果某个struct字段的tag定义中含有XML结构中element的名称，那么解析时就会把相应的element值赋给该字段，如上面例子中的serverName和serverIP定义。

4. 如果某个struct字段的tag定义了含有",attr"，那么解析时就会将该结构所对应的element的与字段同名的属性的值赋给该字段，如上version定义。

5. 如果某个struct字段的tag定义形如"a>b>c"，则解析时，会将XML结构a下面的b下面的c元素的值赋给该字段。

6. 如果某个struct字段的tag定义了"-"，那么不会为该字段解析匹配任何XML数据。

7. 如果struct字段后面的tag定义了",any"，当它的子元素在不满足其他规则时就会匹配到这个字段。

8. 如果某个XML元素包含一条或者多条注释，那么这些注释将被累加到第一个tag含有",comments"的字段上，这个字段的类型可能是[]byte或string，如果没有这样的字段存在，那么注释将会被抛弃。

#### 生成XML

```go
func Marshal(v interface{}) ([]byte, error)
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
```



```go
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}
type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func main() {
	v := &servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Local_Web", "172.0.0.1"})
	v.Svs = append(v.Svs, server{"Local_DB", "172.0.0.2"})
	output, err := xml.MarshalIndent(v, "  ", "  ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
```

结果：

```xml
<?xml version="1.0" encoding="UTF-8"?>
  <servers version="1">
    <server>
      <serverName>Local_Web</serverName>
      <serverIP>172.0.0.1</serverIP>
    </server>
    <server>
      <serverName>Local_DB</serverName>
      <serverIP>172.0.0.2</serverIP>
    </server>
  </servers>
```



Marshal函数接收的参数v是interface{}类型，即它可以接收任意类型的参数，xml包会根据下面的规则来生成相应的XML文件：

1. 如果v是array或者slice，那么便输出每一个元素，类似于value。
2. 如果v是指针，那么会输出Marshal指针指向的内容，如果指针为空，什么都不输出。
3. 如果v是interface，那么就处理interface所包含的数据。
4. 如果v是其他数据类型，就会输出这个数据类型所拥有的字段信息。

生成的XML文件中的element的名字根据如下优先级从struct中获取：

1. 如果v是struct，XMLName的tag中定义的名称。
2. 类型为xml.Name的名叫XMLName的字段的值。
3. 通过struct中字段的tag来获取。
4. 通过struct的字段名来获取。
5. marshall的类型名称。

设置struct中字段的tag信息以控制最终XML文件的生成：

1. XMLName不会被输出。

2. tag中含有"-"的字段不会输出。

3. tag中含有"name,ttr"，会以name作为属性名，字段值作为值输出为这个XML元素的属性，如上version字段所描述。

4. tag中含有",attr"，会以这个struct的字段名作为属性名输出为XML元素的属性，类似于上一条，只是这个name默认是字段名。

5. tag中含有",chardata"，输出为XML的character data而非element。

6. tag中含有",innerxml"，将会被原样输出，而不会进行常规的编码过程。

7. tag中含有",comment"，将被当作XML注释来输出，而不会进行常规的编码过程，字段值中不能含有"--"字符串。

8. tag中含有"omitempty"，如果该字段的值为空值，那么该字段就不会被输出到XML，空值包括false、0、nil指针或nil接口，以及任何长度为0的array、slice、map或者string。

9. tag中含有"a>b>c"，那么就会循环输出3个元素，a包含b，b包含c，例如：

   ```
       FirstName string `xml: "name>first"`
       LastName string `xml :"name>last"`
       <name>
           <first>Asta</first>
           <1ast>Xie</last>
       </name>
   ```

   

#### XML文件的读写操作



### 50.3 JSON文件处理

#### 解析JSON

解析JSON解析JSON有两种方法，一种是解析到结构体，另一种是解析到接口，前者是在知晓被解析的JSON数据结构的前提下采取的方案，如果不知道被解析的数据的格式，则应该采用解析到接口的方案。

json包有对应的函数：

```go
func Unmarshal(data []byte, v interface{}) error
```





#### 生成JSON





#### JSON文件的读写操作



### 50.4 日志

```sh
go get -u github.com/sirupsen/logrus
```



### 50.5 压缩

#### 打包与解包

`archive/zip`

打包

解包



#### 压缩与解压

`compress/gzip`

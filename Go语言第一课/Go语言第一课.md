Go语言第一课
---

[Go 语言第一课](https://time.geekbang.org/column/intro/100093501)

官方网站：https://golang.google.cn/ or https://go.dev/

发布时间：2021-2022

## 0 这样入门Go，才能少走弯路

### 入坑Go的三大理由

#### 1 对初学者足够友善，能够快速上手

#### 2 生产力与性能的最佳结合

Go已经成为了云基础架构语言，它在**云原生基础设施、中间件与云服务**领域大放异彩。同时，GO在**DevOps/SRE、区块链、命令行交互程序（CLI）、Web服务，还有数据处理**等方面也有大量拥趸，我们甚至可以看到Go在**微控制器（TinyGo）、机器人（Gobot）、游戏领域（游戏服务器和工具链）**也有广泛应用。

#### 3 快乐又有“钱景”

简单的语法、得心应手的工具链、丰富和健壮的标准库，还有生产力与性能的完美结合、免除内存管理的心智负担，对并发设计的原生支持

## 怎样学才能少走弯路？

Go独特的编程思维方式

“设计意识”。编程语言学习的最终目的是写出具有现实实用意义的程序，所以你要培养自己**选择适当的语言元素构造程序骨架的能力**。



# 前置篇：心定之旅

## 1 前世今生：Go的历史和现状

了解一门编程语言的历史和现状，以及未来的走向，可以建立起**学习的“安全感”**，相信它能够给你带来足够的价值和收益，更加坚定地学习下去。

### 诞生

Go语言的创始人有三位：

- 图灵奖获得者、C语法联合发明人、Unix之父**肯·汤普森（Ken Thompson）**
- Plan 9操作系统领导者、UTF-8编码的最初设计者**罗伯·派克（Rob Pike）**
- Java的HotSpot虚拟机和Chrome浏览器的JavaScript V8引擎的设计者之一**罗伯特·格瑞史莫（Robert Griesemer）**。

![](images/image-20241109142046263.png)

开始：

- 2007年9月20日 讨论

不便：C++的巨大复杂性、编译构建速度慢以及在编写服务端程序时对并发支持的不足

思路：在C语言的基础上，修正一些明显的缺陷，删除一些被诟病较多的特性，增加一些缺失的功能，比如，使用import替代include、去掉宏、增加垃圾回收、支持接口等。

9月25日，罗伯·派克 命名为“go”。

“Golang”仅应用于命名Go语言官方网站。

### 从“三人行”到“众人拾柴”

- 2008年初，肯·汤普森实现了第一版Go编译器，用于验证之前的设计。这个编译器先将Go代码转换为C代码，再由C编译器编译成二进制文件。

- 2008年年中，Go的第一版设计就基本结束了。伊恩·泰勒（Ian Lance Taylor）为Go语言实现了一个gcc的前端，这也是Go语言的第二个编译器。

之后他成为了Go语言，以及其工具设计和实现的核心人物之一。

- 2008年，**罗斯·考克斯（Russ Cox）**利用函数类型是“一等公民”，而且它也可以拥有自己的方法这个特性巧妙设计出了http包的`HandlerFunc`类型。这样，我们通过显式转型就可以让一个普通函数成为满足http.Handler接口的类型了。

- 2009年10月30日，Go语言第一次公之于众。

- 2009年11月10日，谷歌官方宣布Go语言项目开源，之后这一天也被Go官方确定为Go语言的诞生日。

“吉祥物”，是一只由罗伯·派克夫人芮妮·弗伦奇（Renee French）设计的地鼠，从此地鼠（gopher）也就成为了世界各地Go程序员的象征，Go程序员也被昵称为Gopher。

![](images/image-20240707112359627.png)

- 2012年3月28日，Go 1.0版本正式发布

![](images/image-20240701100957741.png)

- 2024-9-1  **Austin Clements**将接替Russ Cox成为Go的技术负责人

### Go是否值得我们学习？

现代**==云计算基础设施软件==**的大部分流行和可靠的作品，都是用Go编写的，比如：Docker、Kubernetes、Prometheus、Ethereum（以太坊）、Istio、CockroachDB、InfluxDB、Terraform、Etcd、Consul等等。

![](images/image-20241109142217711.png)

Go在经历了一个漫长的技术萌芽期后，从实现自举的Go 1.5版本开始逐步进入“期望膨胀期”，在经历从Go 1.6到Go 1.9版本的发布后，业界对Go的期望达到了峰值。

但随后“泡沫破裂”，在Go 1.11发布前跌到了“泡沫破裂期”的谷底，Go 1.11版本引入了Go module，给社区解决Go包依赖问题注射了一支强心剂，于是Go又开始了缓慢爬升。

从TIOBE提供的曲线来看，Go 1.12到Go 1.15版本的发布让我们有信心认为Go已经走出“泡沫破裂谷底期”，进入到“稳步爬升的光明期”。

[TIOBE Index - Go](https://www.tiobe.com/tiobe-index/go/)



### 思考题

>  相较于传统的静态编译型编程语言（如C、C++），Go做出了哪些改进？

1. **内存管理：从手动到自动的飞跃**

C/C++需手动调用`malloc/free`或`new/delete`，易因遗漏或重复释放引发崩溃；而Go的GC通过三色标记法实现高效回收，延迟控制在毫秒级。

2. **并发模型：原生支持高并发编程**

Go通过goroutine（协程）和channel（通道）实现并发编程范式革新：

- 协程开销：单个goroutine初始仅需2KB栈内存，可同时启动数百万个，远优于C/C++的线程（MB级开销）。
- 通信机制：Channel提供线程安全的通信方式，避免C/C++中复杂的锁竞争问题。
- 开发效率：通过go关键字即可启动协程，相比C/C++依赖第三方库（如POSIX线程）大幅简化代码量。

3. **语法设计与开发效率**

极简语法与强制代码规范

- 关键字数量：仅25个（C++有80+），去除类继承、构造函数等复杂概念，改用组合和接口实现扩展。
- 代码格式化：内置gofmt工具统一代码风格，消除团队协作中的格式争议。
- 编译速度：增量编译速度比C/C++快5-10倍，支持“编码即运行”的快速迭代模式

4. **工具链与标准库的完备性**

5. **性能与资源消耗的平衡**

执行速度：编译为本地机器码，性能为C的1.2-1.5倍，远超Java/Python。

容器适配：生成静态二进制文件，无需外部依赖，基础镜像仅5-20MB（C/C++需携带运行时库）。

内存占用：微服务实例常驻内存可低至数十MB，适合云原生弹性伸缩场景。

6. **安全性与跨平台支持**

7. **应用场景扩展**

云原生与分布式系统的首选语言

总结：Go语言通过**自动内存管理、原生并发支持、极简语法**三大支柱，解决了C/C++在开发效率、安全性和现代应用适配上的痛点，同时保持了接近原生代码的性能。

## 2 Go语言的设计哲学

编程语言的设计哲学，就是指决定这门**语言==演化进程==的高级原则和依据**。

> 设计哲学之于编程语言，就好比一个人的价值观之于这个人的行为。

Go语言的设计哲学总结为五点：简单、显式、组合、并发和面向工程。

### 简单

Go语言的设计者们在语言设计之初，就拒绝了走**语言特性融合**的道路，选择了“做减法”并致力于打造一门简单的编程语言。

其实，Go语言也没它看起来那么简单，自身实现起来并不容易，但这些**复杂性被Go语言的设计者们“隐藏”了**，所以Go语法层面上呈现了这样的状态：

- 仅有25个关键字，主流编程语言最少；

- **内置垃圾收集**，降低开发人员内存管理的心智负担；

- **首字母大小写决定可见性**，无需通过额外关键字修饰；

- **变量初始为类型零值**，避免以随机值作为初值的问题；

- **内置数组边界检查**，极大减少越界访问带来的安全隐患；

- **内置并发**支持，简化并发程序设计；

- **内置接口**类型，为组合的设计哲学奠定基础；

- 原生提供完善的**工具链**，开箱即用；

- … …

> 任何的设计都存在着**==权衡与折中==**。

> 简单意味着可以使用**==更少的代码==**实现相同的功能；简单意味着代码具有更好的**==可读性==**，而可读性好的代码通常意味着更好的**==可维护性以及可靠性==**。

简单的设计哲学是Go生产力的源泉。

### 显式

Go不允许不同类型的整型变量进行混合计算，它同样也**不会对其进行隐式的自动转换**。

Go希望开发人员明确知道自己在做什么，这与C语言的“信任程序员”原则完全不同。

除此之外，Go设计者所崇尚的显式哲学还直接决定了Go语言错误处理的形态：Go语言采用了显式的基于值比较的错误处理方案，**函数/方法中的错误都会通过return语句显式地返回**，并且通常调用者不能忽略对返回的错误的处理。

### 组合

Go语言不像C++、Java等主流面向对象语言，在Go中是找不到经典的==面向对象语法元素==、==类型体系==和==继承机制==的，Go推崇的是==组合==的设计哲学。

在Go语言设计层面，Go设计者为开发者们提供了**==正交==的语法元素**，以供后续组合使用，包括：

- Go语言**无类型层次体系**，各类型之间是相互独立的，**没有子类型**的概念；
- 每个类型都可以有自己的方法集合，类型定义与方法实现是**正交独立**的；
- 实现某个接口时，无需像Java那样采用特定关键字修饰；
- 包之间是相对独立的，**没有子包**的概念。

Go语言为支撑组合的设计提供了==类型嵌入（Type Embedding）==。通过类型嵌入，我们可以将已经实现的功能嵌入到新类型中，以快速满足新类型的功能需求，这种方式有些类似经典面向对象语言中的“继承”机制，但在原理上却与面向对象中的继承完全不同，这是一种Go设计者们精心设计的“语法糖”。

被嵌入的类型和新类型两者之间**没有任何关系**，甚至相互完全**不知道对方的存在**，更没有经典面向对象语言中的那种父类、子类的关系，以及==向上、向下转型（Type Casting）==。通过新类型实例调用方法时，方法的匹配主要取决于**方法名字**，而不是类型。这种组合方式，我称之为==垂直组合==，即通过类型嵌入，快速让一个新类型“复用”其他类型已经实现的能力，实现功能的==垂直扩展==。

```go
// $GOROOT/src/sync/pool.go
type poolLocal struct {
    private interface{}   
    shared  []interface{}
    Mutex               
    pad     [128]byte  
}
```

在poolLocal这个结构体类型中嵌入了类型Mutex，这就使得poolLocal这个类型具有了互斥同步的能力，可以通过poolLocal类型的变量，直接调用Mutex类型的方法Lock或Unlock。

```go
// $GOROOT/src/io/io.go
type ReadWriter interface {
    Reader
    Writer
}
```

通过嵌入接口类型的方式来实现接口行为的聚合，组成==大接口==，这种方式在标准库中尤为常用，并且已经成为了Go语言的一种惯用法。

垂直组合本质上是一种“==能力继承==”，采用嵌入方式定义的新类型继承了嵌入类型的能力。Go还有一种常见的组合方式，叫==水平组合==。和垂直组合的能力继承不同，水平组合是一种==能力委托（Delegate）==，通常使用接口类型来实现水平组合。

Go语言中的接口只是**方法集合**，并且它与实现者之间的关系无需通过显式关键字修饰，它让程序内部各部分之间的耦合降至最低，同时它也是连接程序各个部分之间“纽带”。

水平组合的模式1️⃣，通过接受接口类型参数的普通函数进行组合：

```go
// $GOROOT/src/io/ioutil/ioutil.go
func ReadAll(r io.Reader)([]byte, error)

// $GOROOT/src/io/io.go
func Copy(dst Writer, src Reader)(written int64, err error)
```

函数ReadAll通过io.Reader这个接口，将io.Reader的实现与ReadAll所在的包低耦合地水平组合在一起了，从而达到从任意实现io.Reader的数据源读取所有数据的目的。类似的水平组合“模式”还有**点缀器、中间件**等。

2️⃣将Go语言内置的并发能力进行灵活组合以实现，比如，通过goroutine+channel的组合，可以实现类似Unix Pipe的能力。

总之，组合原则的应用实质上是塑造了Go程序的骨架结构。类型嵌入为类型提供了垂直扩展能力，而接口是水平组合的关键。

### 并发

> “并发”出现的背景
> 
> CPU都是靠提高**==主频==**来改进性能的，但是现在这个做法已经遇到了瓶颈。主频提高导致CPU的功耗和发热量剧增，反过来制约了CPU性能的进一步提高。2007年开始，处理器厂商的竞争焦点从主频转向了**==多核==**。

Go放弃了传统的基于**操作系统线程**的并发模型，而采用了**用户层轻量级线程**，Go将之称为==goroutine==。

goroutine占用的资源非常小，Go运行时默认为每个goroutine分配的栈空间仅==2KB==。goroutine调度的切换也不用**陷入（trap）**操作系统内核层完成，代价很低。因此，一个Go程序中可以创建**成千上万**个并发的goroutine。而且，所有的Go代码都在goroutine中执行，哪怕是go运行时的代码也不例外。

Go还在语言层面内置了辅助并发设计的原语：`channel`和`select`。开发者可以通过语言内置的channel**传递消息或实现同步**，并通过select实现**多路channel的并发控制**。

并发与组合的哲学是一脉相承的，**并发是一个更大的组合的概念**，它在程序设计的全局层面对程序进行拆解组合，再映射到程序执行层面上：**goroutines各自执行特定的工作，通过channel+select将goroutines组合连接起来**。并发的存在鼓励程序员在程序设计时进行**独立计算的分解**，而对并发的原生支持让Go语言也更适应现代计算环境。

### 面向工程

Go语言设计的初衷，就是面向解决真实世界中Google内部大规模软件开发存在的各种问题，为这些问题提供答案，这些问题包括：**程序构建慢、依赖管理失控、代码难于理解、跨语言构建难**等。

Go在语法设计细节上做了精心的打磨。比如：

- 重新设计**编译单元和目标文件格式**，实现Go源码快速构建，让大工程的构建时间缩短到类似动态语言的交互式解释的编译速度；
- 如果源文件导入它不使用的包，则程序将无法编译。这可以充分保证任何Go程序的**依赖树是精确的**。这也可以保证在构建程序时不会编译额外的代码，从而最大限度地缩短编译时间；
- **去除包的循环依赖**，循环依赖会在大规模的代码中引发问题，因为它们要求编译器同时处理更大的源文件集，这会减慢增量构建；
- **包路径是唯一的，而包名不必唯一的**。导入路径必须唯一标识要导入的包，而名称只是包的使用者如何引用其内容的约定。“包名称不必是唯一的”这个约定，大大降低了开发人员给包起唯一名字的心智负担；
- 故意**不支持默认函数参数**。因为在规模工程中，很多开发者利用默认函数参数机制，向函数添加过多的参数以弥补函数API的设计缺陷，这会导致函数拥有太多的参数，降低清晰度和可读性；
- 增加**类型别名（type alias）**，支持大规模代码库的重构。

Go**标准库功能丰富**，多数功能不需要依赖外部的第三方包或库。

Go语言就提供了足以让所有其它主流语言开发人员羡慕的**工具链**，涵盖了**编译构建、代码格式化、包依赖管理、静态代码检查、测试、文档生成与查看、性能剖析、语言服务器、运行时程序跟踪**等方方面面。其中`gofmt`统一了Go语言的代码风格

Go在标准库中提供了官方的**词法分析器、语法解析器和类型检查器**相关包，开发者可以基于这些包快速构建并扩展Go工具链。



# 入门篇：勤加练手

## 3 配好环境：选择一种最适合你的Go安装方法

### 3.1 选择Go版本

Go语言的版本发布策略:

- 每年发布两次大版本，一般是在二月份和八月份发布

- 对最新的两个Go稳定大版本提供支持

- 支持的范围主要包括**修复版本中存在的重大问题、文档变更以及安全问题更新**等。

### 3.2 安装Go

- Linux

```sh
$ wget -c https://golang.google.cn/dl/go1.16.5.linux-amd64.tar.gz
$ tar -C /usr/local -xzf go1.16.5.linux-amd64.tar.gz

$ ls -F /usr/local/go/
CONTRIBUTING.md  README.md        api/             doc/             misc/            test/
LICENSE          SECURITY.md      bin/             go.env           pkg/
PATENTS          VERSION          codereview.cfg   lib/             src/
```

```
export PATH=$PATH:/usr/local/go/bin
```

```
source ~/.profile
```





### 3.3 安装多个Go版本

#### 方法一：重新设置PATH环境变量



#### 方法二：go get命令



#### 方法三：go get命令安装非稳定版本





### 3.4 配置Go

`go env`

| 名称          | 作用                                                         |                                                              |
| ------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `GOARCH`      | 用于指示Go编译器生成代码 所针对的平台CPU架构                 | 主要值是AMD64、Arm等，默认值是本机的CPU架构                  |
| `GOOS`        | 用于指示Go编译器生成代码 所针对的搡作系统                    | 主要值是Linux、Danwin、Windows等，默认值是本机的操作系統     |
| `GO111MODULE` | 它的值决定了当前使用的构建模式是传统的GOPATH模式还是新引入的Go Module模式 | 在Go 1.16版本Go Module构建模式默认开启，该变量值默认为on     |
| `GOCACHE`     | 用于指示存储构建结果缓存的路径，这些缓存可能会被后续的构建所使用 | 在不同的操作系統上，GOCACHE有不同的默认值。以Linux为例，我们使用`go env GOCACHE`查看其值为：`$HOME/.cache/go-build` |
| `GOMODCACHE`  | 用于指示存放GoModule的路径                                   | 在不同的操作系統上，GOMODCACHE有不同的默认值。以Linux为例，我们使用`go env GOMODCACHE`查看其值为： `$HOME/go/pkg/mod` |
| `GOPROXY`     | 用来配置Go Module proxy服 务                                 | 其默认值为“https://proxy.golang.org,direct”d。中国大陆地区，一股设置为大陆地区提供的module proxy服务以加速Go Module的获取速度。 值可设置为“https://goproxy.cn,direct” |
| `GOPATH`      | 在传统的GOPATH构建模式下，用于指示Go包搜索路径的环境变量，在Go module机制启用之前是Go核心配置项。Go1.8版本之前需要手工配置，G0 1.8版本引入了默认的GOPATH （`$HOME/go`）。在Go Module 模式正式上位后，很可能会被移除 |                                                              |
| `GOROOT`      | 指示Go安装路径。Go1.10版本引入了默认GOROOT，开发者置无需显示设置，Go程序会自动根据自己所在路径推导出GOROOT的路径 |                                                              |

`go help environment`  查看更多关于Go配置项的说明



## 4 初窥门径：一个Go程序的结构是怎样的？

Go源文件总是用**==全小写==**字母形式的短小单词命名，并且以.go扩展名结尾。多个单词就直接连接起来，**不要用下划线连接**（下划线在在Go源文件命名中有特殊作用）。

```go
package main

import "fmt"

func main() {
    fmt.Println("hello, world")
} 
```

```shell
$ go build main.go
$ ./main
hello, world  
```

### 示例程序的结构

`package main`定义了Go中的一个包package。包是Go语言的基本组成单元，通常使用单个的**==小写==**单词命名，一个Go程序本质上就是**一组包的集合**。

所有Go代码都有自己隶属的包。

**整个Go程序中仅允许存在一个名为main的包。**

**main函数**比较特殊：当你运行一个可执行的Go程序的时候，所有的代码都会从这个入口函数开始运行。

`Gofmt`是Go语言在解决规模化（scale）问题上的一个最佳实践。这个工具可以将代码自动格式化为约定的风格。



```go
fmt.Println("hello, world")
```

将字符串输出到终端的标准输出（stdout）上。

三个注意点：

**注意点1**：标准Go代码风格使用**Tab**而不是空格来实现缩进的。

**注意点2**：我们调用一个位于Go标准库的fmt包中名为Println的函数，其实做了两步操作。

第一步是在源文件的开始处通过import声明导入fmt包的包路径。

第二步则是在main函数体中，通过fmt这个**限定标识符**（Qualified Identifier）调用Println函数。

这两个fmt的含义是不同的：

- `import “fmt”` 一行中“`fmt`”代表的是包的导入路径（Import），它表示的是标准库下的fmt目录，整个import声明语句的含义是导入标准库fmt目录下的包；
- `fmt.Println`函数调用一行中的“`fmt`”代表的则是包名。两者是不一样的。

在Go语言中，只有首字母为大写的标识符才是导出的（Exported），才能对包外的代码可见；main包是不可以像标准库fmt包那样被导入（Import）的，否则报错：`import “xx/main” is a program, not an importable package`。

**注意点3**：传入的字符串也就是执行程序后在终端的标准输出上看到的字符串（中文等也是如此）。

因为Go源码文件本身采用的是Unicode字符集，而且用的是UTF-8标准的字符编码方式，这与编译后的程序所运行的环境所使用的字符集和字符编码方式是一致的。

### Go语言中程序是怎么编译的？

```shell
$ go build main.go  
```

`go run`这类命令更多用于开发调试阶段，真正的交付成果还是需要使用`go build`命令构建的。

### 复杂项目下Go程序的编译是怎样的

==Go module==构建模式是在Go 1.11版本正式引入的，为的是**彻底解决Go项目复杂版本依赖的问题**，在Go 1.16版本中，Go module已经成为了Go默认的**包依赖管理机制和Go源码构建机制**。

Go Module的核心是一个名为go.mod的文件，在这个文件中存储了这个module对第三方依赖的全部信息。

```shell
$ go mod init github.com/andyron/hellomodule
go: creating new go.mod: module github.com/bigwhite/hellomodule
go: to add module requirements and sums:
	go mod tidy
```

`go mod init `命令创建`go.mod`文件：

```
module github.com/andyron/hellomodule

go 1.22.1
```

一个module就是一个包的集合，这些包和module一起**打版本、发布和分发**。go.mod所在的目录被称为它声明的**module的根目录**。

第一行内容是用于声明==module路径（module path）==的。而且，module隐含了一个==命名空间==的概念，module下每个包的导入路径都是由**==module path==和==包所在子目录的名字==**结合在一起构成。比如，如果hellomodule下有子目录pkg/pkg1，那么pkg1下面的包的导入路径就是由module path（`github.com/andyron/hellomodule`）和包所在子目录的名字（pkg/pkg1）结合而成，也就是`github.com/andyron/hellomodule/pkg/pkg1`。

`go 1.22.1`是一个Go版本指示符，用于表示这个module是在某个特定的Go版本的module语义的基础上编写的。

> `go mod tidy`，用于清理和管理项目的依赖关系，可以确保你的 `go.mod` 和 `go.sum` 文件是最新的，它会执行下面的操作：
>
> - **添加缺失的依赖**
> - **移除未使用的依赖**
> - **更新依赖的版本**
>
> ```shell
> $ go mod tidy       
> go: downloading go.uber.org/zap v1.18.1
> go: downloading github.com/valyala/fasthttp v1.28.0
> go: downloading github.com/andybalholm/brotli v1.0.2
> ... ...
> ```

`go.sum`文件记录了hellomodule的**直接依赖和间接依赖包的相关版本的hash值，用来校验本地包的真实性**。在构建的时候，如果本地依赖包的hash值与go.sum文件中记录的不一致，就会被拒绝构建。



## 5 标准先行：Go项目的布局标准是什么？

### 5.1 Go语言“创世项目”结构是怎样的？🔖

“Go语言的创世项目”就是Go语言项目自身。



```sh
$ cd go // 进入Go语言项目根目录
$ git checkout go1.3 // 切换到go 1.3版本
$ tree -LF 1 ./src/
./src//
├── Make.dist
├── all.bash*
├── all.bat
├── all.rc*
├── archive/
├── clean.bash*
├── clean.bat
├── clean.rc*
├── cmd/
├── compress/
├── container/
├── internal/
├── io/
├── lib9/
├── libbio/
├── liblink/
├── log/
├── make.bash*
├── make.bat
├── make.rc*
├── math/
├── nacltest.bash*
├── path/
├── pkg/
├── race.bash*
├── race.bat
├── run.bash*
├── run.bat
├── run.rc*
├── runtime/
├── sudo.bash*
├── testing/
├── text/
├── time/
└── unicode/
```



三个比较重要的演进:

- **演进一：Go 1.4版本删除pkg这一中间层目录并引入internal目录**



- **演进二：Go1.6版本增加vendor目录**



- **演进三：Go 1.13版本引入go.mod和go.sum**

### 5.2 现在的Go项目的典型结构布局是怎样的？

#### 1️⃣可执行程序项目

```shell
$tree -F exe-layout 
exe-layout
├── cmd/
│   ├── app1/
│   │   └── main.go
│   └── app2/
│       └── main.go
├── go.mod
├── go.sum
├── internal/
│   ├── pkga/
│   │   └── pkg_a.go
│   └── pkgb/
│       └── pkg_b.go
├── pkg1/
│   └── pkg1.go
├── pkg2/
│   └── pkg2.go
└── vendor/
```



典型五个部分：

- 放在项目顶层的Go Module相关文件，包括==go.mod和go.sum==；

- **cmd目录**：存放项目要编译构建的可执行文件所对应的main包的源码文件；如果项目中有多个可执行文件需要构建，每个可执行文件的main包单独放在一个子目录中，比如图中的app1、app2，cmd目录下的各app的main包将整个项目的依赖连接在一起。

  通常main包应该很简洁，它会做一些**命令行参数解析、资源初始化、日志设施初始化、数据库连接初始化**等工作，之后就会将程序的执行权限交给更高级的执行控制对象。

  也有一些Go项目将cmd这个名字改为app或其他名字，但它的功能其实并没有变。

- 项目包目录：每个项目下的非main包都“平铺”在项目的根目录下，每个目录对应一个Go包；

- **pkgN目录**，这是一个存放项目自身要使用、同样也是可执行文件对应main包所要依赖的库文件，同时这些目录下的包还可以被外部项目引用。

- **internal目录**：存放仅项目内部引用的Go包，这些包无法被项目之外引用；

- **vendor目录**：这是一个可选目录，为了兼容Go 1.5引入的vendor构建模式而存在的。这个目录下的内容均由Go命令自动维护，不需要开发者手工干预。

- 有些开发者喜欢借助一些第三方的构建工具辅助构建，比如：make、bazel等。你可以将这类外部辅助构建工具涉及的诸多脚本文件（比如Makefile）放置在项目的顶层目录下，就像Go创世项目中的all.bash那样。



##### 多个module

Go支持在一个项目/仓库中存在多个module，但这种管理方式可能要比一定比例的代码重复引入更多的复杂性。

建议将项目拆分为多个项目（仓库），每个项目单独作为一个module进行单独的版本管理和演进。



#### 2️⃣库项目

Go库项目仅对外暴露Go包，典型布局形式：

```shell
$tree -F lib-layout 
lib-layout
├── go.mod
├── internal/
│   ├── pkga/
│   │   └── pkg_a.go
│   └── pkgb/
│       └── pkg_b.go
├── pkg1/
│   └── pkg1.go
└── pkg2/
    └── pkg2.go
```

库类型项目相比于Go可执行程序项目的布局，**不需要构建可执行程序**，所以去除了cmd目录。

库类型项目不推荐在项目中放置vendor目录去缓存库自身的第三方依赖，库项目仅通过go.mod文件明确表述出该项目依赖的module或包以及版本要求就可以了。

Go库项目的初衷是为了**对外部（开源或组织内部公开）暴露API，对于仅限项目内部使用而不想暴露到外部的包，可以放在项目顶层的internal目录下面**。当然internal也可以有多个并存在于项目结构中的任一目录层级中，关键是项目结构设计人员要明确各级internal包的应用层次和范围。



### 5.3 注意早期Go可执行程序项目的典型布局

```shell
$tree -L 3 -F early-project-layout
early-project-layout
└── exe-layout/
    ├── cmd/
    │   ├── app1/
    │   └── app2/
    ├── go.mod
    ├── internal/
    │   ├── pkga/
    │   └── pkgb/
    ├── pkg/
    │   ├── pkg1/
    │   └── pkg2/
    └── vendor/
```



思考：

> 考虑Go项目结构的最小标准布局中都应该包含哪些东西呢？

```
# 最小标准布局（Russ Cox 推荐）
project-root/
├── go.mod
├── LICENSE
├── main.go          # 主程序入口
└── utils/           # 工具包
    └── math.go

# 扩展后的典型结构（社区实践）
project-root/
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   └── auth/
├── pkg/
│   └── db/
├── go.mod
└── LICENSE
```



## 6 构建模式：Go是怎么解决包依赖管理问题的？

### 6.1 Go构建模式是怎么演化的？

Go程序由Go包组合而成的，Go程序的==构建过程==就是**确定包版本、编译包以及将编译后得到的目标文件链接在一起**的过程。

Go语言的构建模式历经了三个迭代和演化过程：

#### 1️⃣最初期的GOPATH



```shell
$go get github.com/sirupsen/logrus
```



**在GOPATH构建模式下，Go编译器实质上并没有关注Go项目所依赖的第三方包的版本。**引入了Vendor机制试图解决这个问题。

#### 2️⃣1.5版本的Vendor机制

vendor机制本质上就是在Go项目的某个特定目录下，将项目的所有依赖包缓存起来，这个特定目录名就是vendor。



#### 3️⃣现在的Go Module

一个Go Module是一个Go包的集合。module是有版本的，所以module下的包也就有了**版本属性**。这个module与这些包会组成一个**独立的版本单元**，它们一起打版本、发布和分发。

在Go Module模式下，通常一个代码仓库对应一个Go Module。一个Go Module的顶层目录下会放置一个go.mod文件，每个go.mod文件会定义唯一一个module，也就是说**Go Module与go.mod是一一对应的**。

go.mod文件所在的顶层目录也被称为**module的根目录**，module根目录以及它子目录下的所有Go包均归属于这个Go Module，这个module也被称为**main module**。



### 6.2 创建一个Go Module

步骤：

1. 第一步，通过`go mod init`创建go.mod文件，将当前项目变为一个Go Module；
2. 第二步，通过`go mod tidy`命令自动更新当前module的依赖信息；
3. 第三步，执行`go build`，执行新module的构建。

由`go mod tidy`下载的依赖module会被放置在本地的module缓存路径下，默认值为`$GOPATH[0]/pkg/mod`，Go 1.15及以后版本可以通过`GOMODCACHE`环境变量，自定义本地module的缓存路径。

> 推荐把go.mod和go.sum两个文件与源码，一并提交到代码版本控制服务器上。

go build命令会读取go.mod中的依赖及版本信息，并在本地module缓存路径下找到对应版本的依赖module，执行编译和链接。



项目所依赖的包有很多版本，Go Module是如何选出最适合的那个版本的呢？

### 6.3 深入Go Module构建模式

Go语言设计者在设计Go Module构建模式，来解决“包依赖管理”的问题时，进行了创新:

#### 语义导入版本(Semantic Import Versioning)

版本号，都符合`vX.Y.Z`的格式，由==前缀v==和一个==满足语义版本规范的版本号==组成。语义版本号分成3部分：主版本号(major)、次版本号(minor)和补丁版本号(patch)。

![](images/image-20240702171216789.png)

借助于语义版本规范，Go命令可以确定同一module的两个版本发布的先后次序，而且可以确定它们是否兼容。

按照语义版本规范，**主版本号不同的两个版本是相互不兼容的**。而且，在主版本号相同的情况下，**次版本号大都是向后兼容次版本号小的版本。补丁版本号也不影响兼容性**。

而且，Go Module规定：**如果同一个包的新旧版本是兼容的，那么它们的包导入路径应该是相同的**。以logrus为例，选出两个版本v1.7.0和v1.8.1.。按照上面的语义版本规则，这两个版本的主版本号相同，新版本v1.8.1是兼容老版本v1.7.0的。那么，我们就可以知道，如果一个项目依赖logrus，无论它使用的是v1.7.0版本还是v1.8.1版本，它都可以使用下面的包导入语句导入logrus包：

```go
import "github.com/sirupsen/logrus"
```

> 新问题：
> 
> 假如在未来的某一天，logrus的作者发布了logrus v2.0.0版本。那么根据语义版本规则，该版本的主版本号为2，已经与v1.7.0、v1.8.1的主版本号不同了，那么v2.0.0与v1.7.0、v1.8.1就是不兼容的包版本。然后我们再按照Go Module的规定，如果一个项目依赖logrus v2.0.0版本，那么它的包导入路径就不能再与上面的导入方式相同了。那我们应该使用什么方式导入logrus v2.0.0版本呢？

Go Module创新性地给出了一个方法：**将包主版本号引入到包导入路径中**：

```go
import "github.com/sirupsen/logrus/v2"
```

甚至可以同时依赖一个包的两个不兼容版本：

```go
import (
  "github.com/sirupsen/logrus"
  logv2 "github.com/sirupsen/logrus/v2"
)
```

> v0.y.z版本应该使用哪种导入路径呢？
> 
> v0.y.z这样的版本号是用于项目初始开发阶段的版本号。在这个阶段任何事情都有可能发生，其API也不应该被认为是稳定的。Go Module将这样的版本(v0)与主版本号v1做同等对待，也就是采用不带主版本号的包导入路径，这样一定程度降低了Go开发人员使用这样版本号包时的心智负担。

#### 最小版本选择(Minimal Version Selection)

![](images/image-20240702171932253.png)

> myproject有两个直接依赖A和B，A和B有一个共同的依赖包C，但A依赖C的v1.1.0版本，而B依赖的是C的v1.3.0版本，并且此时C包的最新发布版为C v1.7.0。这个时候，Go命令是如何为myproject选出间接依赖包C的版本呢？选出的究竟是v1.7.0、v1.1.0还是v1.3.0呢？

当前存在的主流编程语言，以及Go Module出现之前的很多Go包依赖管理工具都会选择依赖项的“**最新最大(Latest Greatest)版本**”，也就是v1.7.0。

Go设计者另辟蹊径，在诸多兼容性版本间，他们不光要考虑最新最大的稳定与安全，还要尊重各个module的述求：A明明说只要求C v1.1.0，B明明说只要求C v1.3.0。所以Go会在该项目依赖项的所有版本中，选出符合项目整体要求的“最小版本”。

这个例子中，C v1.3.0是符合项目整体要求的版本集合中的版本最小的那个，于是Go命令选择了C v1.3.0，而不是最新最大的C v1.7.0。Go团队认为**“最小版本选择”为Go程序实现持久的和可重现的构建提供了最佳的方案**。

### 6.4 Go各版本构建模式机制和切换

![](images/image-20240702172417325.png)

> 未来，Go Module构建模式将成为Go语言唯一的标准构建模式。

思考：

> 如何将基于GOPATH构建模式的现有项目迁移为使用Go Module构建模式？🔖



## 7 构建模式：GoModule的6类常规操作

维护Go Module就是对Go Module依赖包的管理。

### 1️⃣为当前module添加一个依赖

```go
package main

import "github.com/sirupsen/logrus"
import "github.com/google/uuid"

func main() {
    logrus.Println("hello, go module mode.")
    logrus.Println(uuid.NewString())
}
```

可以`go get github.com/google/uuid`，也可以使用`go mod tidy`命令，在执行构建前自动分析源码中的依赖变化，识别新增依赖项并下载它们。

但对于复杂的项目变更而言，逐一手工添加依赖项显然很没有效率，go mod tidy是更佳的选择。

### 2️⃣升级/降级依赖的版本

```sh
$ go list -m -versions github.com/sirupsen/logrus
github.com/sirupsen/logrus v0.1.0 v0.1.1 v0.2.0 v0.3.0 v0.4.0 v0.4.1 v0.5.0 v0.5.1 v0.6.0 v0.6.1 v0.6.2 v0.6.3 v0.6.4 v0.6.5 v0.6.6 v0.7.0 v0.7.1 v0.7.2 v0.7.3 v0.8.0 v0.8.1 v0.8.2 v0.8.3 v0.8.4 v0.8.5 v0.8.6 v0.8.7 v0.9.0 v0.10.0 v0.11.0 v0.11.1 v0.11.2 v0.11.3 v0.11.4 v0.11.5 v1.0.0 v1.0.1 v1.0.3 v1.0.4 v1.0.5 v1.0.6 v1.1.0 v1.1.1 v1.2.0 v1.3.0 v1.4.0 v1.4.1 v1.4.2 v1.5.0 v1.6.0 v1.7.0 v1.7.1 v1.8.0 v1.8.1 v1.8.2 v1.8.3 v1.9.0 v1.9.1 v1.9.2 v1.9.3
```

降级：

```sh
$ go get github.com/google/uuid@v1.7.0
```

或

```sh
$ go mod edit -require=github.com/sirupsen/logrus@v1.7.0
$ go mod tidy       
go: downloading github.com/sirupsen/logrus v1.7.0
```

升级：

```sh
$ go get github.com/google/uuid@v1.7.1
```

在Go Module构建模式下，当依赖的主版本号为0或1的时候，我们在Go源码中导入依赖包，**不需要在包的导入路径上增加版本号**，也就是：

```go
import github.com/user/repo/v0 等价于 import github.com/user/repo
import github.com/user/repo/v1 等价于 import github.com/user/repo
```

### 3️⃣添加一个主版本号大于1的依赖

语义导入版本机制有一个原则：**如果新旧版本的包使用相同的导入路径，那么新包与旧包是兼容的**。也就是说，如果新旧两个包不兼容，那么我们就应该采用不同的导入路径。

```go
import github.com/user/repo/v2/xxx
```

主版本号大于1的依赖，**在声明它的导入路径的基础上，加上版本号信息**。

```go
package main

import (
	_ "github.com/go-redis/redis/v7" // “_”为空导入
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("hello, go module mode")
	logrus.Println(uuid.NewString())
}
```

```shell
$go get github.com/go-redis/redis/v7
go: downloading github.com/go-redis/redis/v7 v7.4.1
go: downloading github.com/go-redis/redis v6.15.9+incompatible
go get: added github.com/go-redis/redis/v7 v7.4.1
```

go get选择了go-redis v7版本下当前的最新版本v7.4.1。

> ==“空导入”（blank import）==指的是导入了一个包，但在代码中并不使用该包中定义的任何函数、变量、类型等元素，仅仅是为了触发该包内部可能存在的初始化逻辑。例如，有些包在被导入时会自动执行一些设置、注册等初始化操作，这时就可以使用空导入的方式。

### 4️⃣升级依赖版本到一个不兼容版本🔖



### 5️⃣移除一个依赖

列出当前module的所有依赖：

```sh
$ go list -m all
```

删除代码中对包依赖，然后`go build`是不会从当前module中移除相关依赖的，需要使用`go mod tidy`命令。

go mod tidy会自动分析源码依赖，而且将不再使用的依赖从go.mod和go.sum中移除。

### 6️⃣特殊情况：使用vendor

vendor机制虽然诞生于GOPATH构建模式主导的年代，但在Go Module构建模式下，它依旧被保留了下来，并且成为了Go Module构建机制的一个很好的**补充**。特别是在一些<u>不方便访问外部网络</u>，并且对Go应用构建性能敏感的环境，比如在一些内部的持续集成或持续交付环境（CI/CD）中，使用vendor机制可以实现与Go Module等价的构建。

和GOPATH构建模式不同，Go Module构建模式下，再也**无需手动维护vendor目录下的依赖包**了，Go提供了可以快速建立和更新vendor的命令:

```sh
$ go mod vendor
$ tree -LF 2 vendor
vendor
├── github.com/
│   ├── google/
│   ├── magefile/
│   └── sirupsen/
├── golang.org/
│   └── x/
└── modules.txt
```

`go mod vendor`命令在vendor目录下，创建了一份这个项目的依赖包的副本，并且通过`vendor/modules.txt`记录了vendor下的module以及版本。

如果我们要基于vendor构建，而不是基于本地缓存的Go Module构建，我们需要在go build后面加上`-mod=vendor`参数。

在Go 1.14及以后版本中，如果Go项目的顶层目录下存在vendor目录，那么go build**默认也会优先基于vendor构建**，除非你给go build传入`-mod=mod`的参数。

### 思考题

> 如果你是一个公共Go包的作者，在发布你的Go包时，有哪些需要注意的地方？



## 8 入口函数与包初始化：搞清Go程序的执行次序

Go程序由一系列Go包组成，代码的执行也是在各个包之间跳转。入口是**main包中的main函数**。

### 1️⃣main.main函数：Go应用的入口函数

Go语言要求：**可执行程序的main包必须定义main函数，否则Go编译器会报错**。在启动了多个Goroutine的Go应用中，main.main函数将在Go应用的主Goroutine中执行。

不过很有意思的是，在多Goroutine的Go应用中，相较于main.main作为Go应用的==入口==，**main.main函数==返回==的意义其实更大**，因为main函数返回就意味着**整个Go程序的终结**，而且你也不用管这个时候是否还有其他子Goroutine正在执行。

除了main包外，其他包也可以拥有自己的名为main的函数或方法。但按照Go的**可见性规则**（小写字母开头的标识符为非导出标识符），非main包中自定义的main函数仅限于包内使用。

> 对于main包的main函数来说，虽然是用户层逻辑的入口函数，但它却**不一定是用户层第一个被执行的函数**。

### 2️⃣init函数：Go包的初始化函数

如果main包依赖的包中定义了init函数，或者是main包自身定义了init函数，那么Go程序在这个包初始化的时候，就会自动调用它的init函数，因此这些init函数的执行就都会发生在main函数之前。

每个组成Go包的Go源文件中，可以定义多个init函数。

在初始化Go包时，Go会按照一定的次序，**逐一、顺序地**调用这个包的init函数。<u>一般来说，先传递给Go编译器的源文件中的init函数，会先被执行；而同一个源文件中的多个init函数，会按**声明顺序**依次执行。</u>

### 3️⃣Go包的初始化次序

从程序逻辑结构角度来看，Go包是程序逻辑封装的**基本单元**，每个包都可以理解为是一个“自治”的、封装良好的、对外部暴露**有限接口**的**基本单元**。**一个Go程序就是由一组包组成的，程序的初始化就是这些包的初始化**。每个Go包还会有自己的依赖包、常量、变量、init函数（其中main包有main函数）等。

> 注意📢：我们在阅读和理解代码的时候，需要知道这些元素在在程序初始化过程中的初始化顺序，这样便于我们确定在某一行代码处这些元素的当前状态。

Go包的初始化次序：

![](images/image-20240702182932741.png)

1. 首先，main包依赖pkg1和pkg4两个包，所以第一步，Go会根据包导入的顺序，先去初始化main包的第一个依赖包pkg1。
2. 第二步，Go在进行包初始化的过程中，会采用“**==深度优先==**”的原则，递归初始化各个包的依赖包。在上图里，pkg1包依赖pkg2包，pkg2包依赖pkg3包，pkg3没有依赖包，于是Go在pkg3包中按照“==常量 -> 变量 -> init函数==”的顺序先对pkg3包进行初始化；
3. 紧接着，在pkg3包初始化完毕后，Go会回到pkg2包并对pkg2包进行初始化，接下来再回到pkg1包并对pkg1包进行初始化。在调用完pkg1包的init函数后，Go就完成了main包的第一个依赖包pkg1的初始化。
4. 接下来，Go会初始化main包的第二个依赖包pkg4，pkg4包的初始化过程与pkg1包类似，也是先初始化它的依赖包pkg5，然后再初始化自身；然后，当Go初始化完pkg4包后也就完成了对main包所有依赖包的初始化，接下来初始化main包自身。
5. 最后，在main包中，Go同样会按照“常量 -> 变量 -> init函数”的顺序进行初始化，执行完这些初始化工作后才正式进入程序的入口函数main函数。

🔖  包引入错误？变量和常量的执行顺序为什么反了？  [Go 1.22引入的包级变量初始化次序问题 | Tony Bai](https://tonybai.com/2024/03/29/the-issue-in-pkg-level-var-init-order-in-go-1-22/)

Go包的初始化次序，三点：

- 依赖包按“深度优先”的次序进行初始化；
- 每个包内按以“常量 -> 变量 -> init函数”的顺序进行初始化；
- 包内的多个init函数按出现次序进行自动调用。

### 4️⃣init函数的用途

Go包初始化时，init函数的初始化次序在变量之后，这给了开发人员在init函数中**对包级变量进行进一步检查与操作**的机会。

#### 用途1：重置包级变量值

负责对包内部以及暴露到外部的包级数据（主要是包级变量）的**初始状态进行检查**。

例如，标准库flag包：(`flag/flag.go`)❤️

flag包定义了一个导出的包级变量`CommandLine`，如果用户没有通过flag.NewFlagSet创建新的代表命令行标志集合的实例，那么CommandLine就会作为flag包各种导出函数背后，默认的代表命令行标志集合的实例。

```go
var CommandLine = NewFlagSet(os.Args[0], ExitOnError)

func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
    f := &FlagSet{
        name:          name,
        errorHandling: errorHandling,
    }
    f.Usage = f.defaultUsage
    return f
}

func (f *FlagSet) defaultUsage() {
    if f.name == "" {
        fmt.Fprintf(f.Output(), "Usage:\n")
    } else {
        fmt.Fprintf(f.Output(), "Usage of %s:\n", f.name)
    }
    f.PrintDefaults()
}
```

在通过NewFlagSet创建CommandLine变量绑定的FlagSet类型实例时，CommandLine的Usage字段被赋值为defaultUsage。也就是说，如果保持现状，那么使用flag包默认CommandLine的用户就无法自定义usage的输出了。于是，flag包在init函数中重置了CommandLine的Usage字段：

```go
func init() {
    CommandLine.Usage = commandLineUsage // 重置CommandLine的Usage字段
}

func commandLineUsage() {
    Usage()
}

var Usage = func() {
    fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
    PrintDefaults()
}
```

CommandLine的Usage字段，设置为了一个flag包内的未导出函数commandLineUsage，后者则直接使用了flag包的另外一个导出包变量Usage。这样，就可以通过init函数，将CommandLine与包变量Usage关联在一起了。

然后，当用户将自定义的usage赋值给了flag.Usage后，就相当于改变了默认代表命令行标志集合的CommandLine变量的Usage。这样当flag包完成初始化后，CommandLine变量便处于一个合理可用的状态了。

#### 用途2：实现对包级变量的复杂初始化

有些包级变量需要一个比较复杂的初始化过程，有些时候，使用它的**类型零值**（每个Go类型都具有一个零值定义）或通过简单初始化表达式不能满足业务逻辑要求，而init函数则非常适合完成此项工作，标准库http包中就有这样一个典型示例：

```go
// net/http/h2_bundle.go
var (
    http2VerboseLogs    bool // 初始化时默认值为false
    http2logFrameWrites bool // 初始化时默认值为false
    http2logFrameReads  bool // 初始化时默认值为false
    http2inTests        bool // 初始化时默认值为false
)

func init() {
    e := os.Getenv("GODEBUG")
    if strings.Contains(e, "http2debug=1") {
        http2VerboseLogs = true // 在init中对http2VerboseLogs的值进行重置
    }
    if strings.Contains(e, "http2debug=2") {
        http2VerboseLogs = true // 在init中对http2VerboseLogs的值进行重置
        http2logFrameWrites = true // 在init中对http2logFrameWrites的值进行重置
        http2logFrameReads = true // 在init中对http2logFrameReads的值进行重置
    }
}
```

http包定义了一系列布尔类型的特性开关变量，可以通过GODEBUG环境变量的值，开启相关特性开关。

#### 用途3：在init函数中实现“注册模式”

lib/pq包访问PostgreSQL数据库的代码示例：

```go
import (
    "database/sql"
    _ "github.com/lib/pq"
)

func main() {
    db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
    if err != nil {
        log.Fatal(err)
    }

    age := 21
    rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)
    ...
}
```

以空导入的方式导入[lib/pq](https://github.com/lib/pq)包的，main函数中没有使用pq包的任何变量、函数或方法，这样就实现了对PostgreSQL数据库的访问。而这一切的奥秘，全在pq包的init函数中：

```go
func init() {
    sql.Register("postgres", &Driver{})
}
```

利用了用空导入的方式导入lib/pq包时产生的一个“副作用”，也就是lib/pq包作为main包的依赖包，它的init函数会在pq包初始化的时候得以执行。

init函数中，pq包将自己实现的sql驱动注册到了sql包中。这样只要应用层代码在Open数据库的时候，传入驱动的名字（这里是“postgres”)，那么通过sql.Open函数，返回的数据库实例句柄对数据库进行的操作，实际上调用的都是pq包中相应的驱动实现。

实际上，这种**通过在init函数中注册自己的实现的模式，就有效降低了Go包对外的直接暴露，尤其是包级变量的暴露**，从而避免了外部通过包级变量对包状态的改动。

另外，从标准库database/sql包的角度来看，这种“注册模式”实质是一种**工厂设计模式**的实现，sql.Open函数就是这个模式中的工厂方法，它根据外部传入的驱动名称“生产”出不同类别的数据库实例句柄。🔖

这种“注册模式”在标准库的其他包中也有广泛应用，比如说，使用标准库image包获取各种格式图片的宽和高：

```go
package main

import (
    "fmt"
    "image"
    _ "image/gif"  // 以空导入方式注入gif图片格式驱动
    _ "image/jpeg" // 以空导入方式注入jpeg图片格式驱动
    _ "image/png"  // 以空导入方式注入png图片格式驱动
    "os"
)

func main() {
    width, height, err := imageSize(os.Args[1])
    if err != nil {
        fmt.Println("获取图片大小错误：", err)
        return
    }
    fmt.Printf("图片大小：[%d, %d]\n", width, height)
}

func imageSize(imageFile string) (int, int, error) {
    f, _ := os.Open(imageFile) // 打开图文文件
    defer f.Close()

    img, _, err := image.Decode(f) // 对文件进行解码，得到图片实例
    if err != nil {
        return 0, 0, err
    }

    b := img.Bounds() // 返回图片区域
    return b.Max.X, b.Max.Y, nil
}
```

上面这个示例程序支持png、jpeg、gif三种格式的图片，而达成这一目标的原因，正是image/png、image/jpeg和image/gif包都在各自的init函数中，将自己“注册”到image的支持格式列表中了：

```go
// $GOROOT/src/image/png/reader.go
func init() {
    image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
}

// $GOROOT/src/image/jpeg/reader.go
func init() {
    image.RegisterFormat("jpeg", "\xff\xd8", Decode, DecodeConfig)
}

// $GOROOT/src/image/gif/reader.go
func init() {
    image.RegisterFormat("gif", "GIF8?a", Decode, DecodeConfig)
}  
```

```sh
$ go run main.go go.png
图片大小：[276, 348]
```

### 思考题

> 当init函数在检查包数据初始状态时遇到失败或错误的情况，该如何处理呢？

需根据错误性质和项目需求采取不同的策略。

#### 1️⃣直接终止程序（适用于关键初始化失败）

当初始化错误导致程序无法正常运行时，**立即终止进程**是最直接的方案：

```go
func init() {
    if err := checkConfig(); err != nil {
        log.Fatalf("初始化失败: %v", err)  // 记录错误日志并退出
        // 或直接 panic(err)
    }
}
```

- 适用场景：
  - 数据库连接、配置文件加载等**关键依赖初始化失败**
  - 无法通过重试或降级恢复的致命错误（如加密密钥缺失）
- **优点**：快速暴露问题，避免后续逻辑因数据无效导致更严重错误
- **缺点**：缺乏容错能力，可能导致服务整体不可用（需结合监控告警）

#### 2️⃣延迟初始化（降低耦合性）

将可能失败的逻辑从 `init` 函数中剥离，改为**显式调用+错误返回**的延迟初始化模式：

```go
// 包级变量改为延迟初始化
var dbConn *sql.DB

func InitDB() error {
    var err error
    dbConn, err = sql.Open("mysql", "user:pass@/dbname")
    return err
}

// 调用方在 main 中处理错误
func main() {
    if err := mypkg.InitDB(); err != nil {
        log.Fatal(err)
    }
}
```

- 适用场景：
  - 非关键路径的初始化（如第三方服务连接、缓存预热）
  - 需要灵活控制初始化时机的场景（如测试环境跳过某些步骤）
- 优点：
  - 错误处理可控，支持重试、降级策略
  - 提升代码可测试性（Mock初始化逻辑更简单）

#### 3️⃣默认值降级（保证基本功能可用）

对于非致命错误，可设置**安全默认值**并记录告警，允许程序以有限功能运行：

```go
var config Config

func init() {
    if err := loadConfig(); err != nil {
        log.Printf("使用默认配置: %v", err)
        config = DefaultConfig()  // 降级为内置默认值
    }
}
```

- 适用场景：
  - 可选配置缺失（如日志级别未指定时默认为 INFO）
  - 外部依赖不可用时的本地回退（如 CDN 资源加载失败时使用本地文件）
- **注意事项**：需在文档中明确默认值行为，避免隐蔽的逻辑错误

#### 4️⃣异步重试机制（提升鲁棒性）

针对网络抖动等临时性错误，可在 `init` 中启动**异步重试**：

```go
func init() {
    go func() {
        for retries := 0; retries < 3; retries++ {
            if err := connectService(); err == nil {
                return
            }
            time.Sleep(2 * time.Second)
        }
        log.Fatal("服务连接重试失败")
    }()
}
```

- 适用场景：
  - 依赖服务启动延迟（如容器化环境中服务发现）
  - 间歇性网络问题（如云服务 API 调用）
- **风险**：需确保异步逻辑不会阻塞主流程或引入竞态条件

#### 总结

|   **策略**   |   **适用阶段**    | **错误严重性** | **恢复能力** |
| :----------: | :---------------: | :------------: | :----------: |
| 直接终止程序 |  关键路径初始化   |    致命错误    |      无      |
|  延迟初始化  | 非关键/可控初始化 |   可恢复错误   |      高      |
|  默认值降级  | 非核心功能初始化  |   非致命错误   |     部分     |
|   异步重试   |  网络依赖初始化   |    临时错误    |     中等     |

建议优先采用**==延迟初始化模式==**，仅在极端情况下使用 `init` 中的 `panic`。对于大型项目，可通过依赖注入框架（如 Wire）管理初始化流程，彻底规避 `init` 的隐式调用问题。

#### `error` 与 `panic` 的对比与选择

|    **维度**    |           **`error`**            |           **`panic`**            |
| :------------: | :------------------------------: | :------------------------------: |
|  **适用场景**  | 可预见的常规错误（如文件不存在） | 不可恢复的致命错误（如内存溢出） |
|  **处理方式**  |  通过返回值传递，调用方显式检查  |  终止当前函数，触发 `defer` 链   |
|  **恢复能力**  |       错误处理后流程可继续       |    若无 `recover` 则程序崩溃     |
| **代码侵入性** |         高（需逐层处理）         |        低（但破坏控制流）        |



## 9 即学即练：构建一个Web服务就是这么简单

### 9.1 最简单的HTTP服务

[Go Developer Survey 2024 H1 Results](https://go.dev/blog/survey2024-h1-results) Go应用最广泛的领域调查结果图

![](images/image-20240708111713853.png)

两个web服务相关，API/RPC服务和Web服务（返回html页面）。

```go
package main

import "net/http"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hello, world!"))
    })
    http.ListenAndServe(":8080", nil)
}
```

`ListenAndServe`

`HandleFunc`

第二个参数r代表来自客户端的HTTP请求，第一个参数w则是用来操作返回给客户端的应答的，基于http包实现的HTTP服务的处理函数都要符合这一原型。

将请求中的URI路径与设置的模式字符串进行==最长前缀匹配==，并执行匹配到的模式字符串所对应的处理函数。

### 9.2 图书管理API服务 ❤️ 🔖🔖

小项目模拟的是真实世界的一个书店的图书管理后端服务。这个服务为平台前端以及其他客户端，提供针对图书的CRUD（创建、检索、更新与删除）的基于HTTP协议的API。

![](images/image-20240703103915037.png)

#### 项目建立与布局设计

bookstore

服务大体拆分为两大部分：

- 一部分是HTTP服务器，用来对外提供API服务；
- 另一部分是图书数据的存储模块，所有的图书数据均存储在这里。

Go项目布局标准：

```
├── cmd/
│   └── bookstore/         // 放置bookstore main包源码
│       └── main.go
├── go.mod                 // module bookstore的go.mod
├── go.sum
├── internal/              // 存放项目内部包的目录
│   └── store/
│       └── memstore.go     
├── server/                // HTTP服务器模块
│   ├── middleware/
│   │   └── middleware.go
│   └── server.go          
└── store/                 // 图书数据存储模块
    ├── factory/
    │   └── factory.go
    └── store.go
```

#### 项目main包

![](images/image-20240703104236493.png)

```go
// cmd/bookstore/main.go
package main

import (
    _ "bookstore/internal/store"
    "bookstore/server"
    "bookstore/store/factory"
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    s, err := factory.New("mem") // 1️⃣创建图书数据存储模块实例
    if err != nil {
        panic(err)
    }

    srv := server.NewBookStoreServer(":8080", s) // 2️⃣创建http服务实例

    errChan, err := srv.ListenAndServe() // 运行http服务
    if err != nil {
        log.Println("web server start failed: ", err)
        return
    }
    log.Println("web server start ok")

    // 3️⃣通过监视系统信号实现了http服务实例的优雅退出
    c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGINT, syscall.SIGTERM) // 捕获系统信号SIGINT、SIGTERM

    select { // 监视来自errChan以及c的事件
    case err = <-errChan:
        log.Println("web server run failed:", err)
        return
    case <-c:
        log.Println("bookstore program is exiting...")
        ctx, cf := context.WithTimeout(context.Background(), time.Second)
        defer cf()
        err = srv.Shutdown(ctx) // 优雅关闭http服务实例
    }

    if err != nil {
        log.Println("bookstore program exit error: ", err)
        return
    }
    log.Println("bookstore program exit ok")
}
```

> 在Go中，main包是整个程序的入口，还是整个程序中==主要模块初始化与组装的场所==。

**==优雅退出==**，指的就是**程序有机会等待其他的事情处理完再退出**。比如**尚未完成的事务处理、清理资源（比如关闭文件描述符、关闭socket）、保存必要中间状态、内存数据持久化落盘**等等。

http服务实例内部的退出清理工作，包括：**立即关闭所有listener、关闭所有空闲的连接、等待处于活动状态的连接处理完毕**等等。

通过`signal`包的`Notify`捕获了SIGINT、SIGTERM这两个系统信号。

#### 图书数据存储模块（store)

用来**存储整个bookstore的图书数据**的。

图书数据存储有很多种实现方式，最简单的方式莫过于在内存中创建一个map，以图书id作为key，来保存图书信息。生产环境，需要通过Nosql数据库或关系型数据库。

考虑到对多种存储实现方式的支持，将针对图书的有限种存储操作，放置在一个接口类型Store中：

```go
// store/store.go
 type Book struct {
     Id      string   `json:"id"`      // 图书ISBN ID
     Name    string   `json:"name"`    // 图书名称
     Authors []string `json:"authors"` // 图书作者
     Press   string   `json:"press"`   // 出版社
 }

 type Store interface {
     Create(*Book) error        // 创建一个新图书条目
     Update(*Book) error        // 更新某图书条目
     Get(string) (Book, error)  // 获取某图书信息
     GetAll() ([]Book, error)   // 获取所有图书信息
     Delete(string) error       // 删除某图书条目
 }
```

一个对应图书条目的抽象数据类型Book，以及针对Book存取的接口类型Store。这样，对于想要进行图书数据操作的一方来说，他只需要得到一个满足Store接口的实例，就可以实现对图书数据的存储操作了，不用再关心图书数据究竟采用了何种存储方式。这就实现了**图书存储操作与底层图书数据存储方式的解耦**。而且，这种==面向接口编程==也是Go组合设计哲学的一个重要体现。

> 如何创建一个满足Store接口的实例呢？

参考《设计模式》提供的多种创建型模式，选择一种Go风格的工厂模式（创建型模式的一种）来实现满足Store接口实例的创建。`store/factory`包:

```go
// store/factory/factory.go
package factory

import (
    "bookstore/store"
    "fmt"
    "sync"
)

var (
    providersMu sync.RWMutex
    providers   = make(map[string]store.Store) // 使用map类型对工厂可以“生产”的、满足Store接口的实例类型进行管理
)

// Register 让各个实现Store接口的类型可以把自己“注册”到工厂中来
func Register(name string, p store.Store) {
    providersMu.Lock()
    defer providersMu.Unlock()
    if p == nil {
        panic("store: Register provider is nil")
    }

    if _, dup := providers[name]; dup {
        panic("store: Register called twice for provider " + name)
    }
    providers[name] = p
}

// New 传入期望使用的图书存储实现的名称，得到对应的类型实例
func New(providerName string) (store.Store, error) {
    providersMu.RLock()
    p, ok := providers[providerName]
    providersMu.RUnlock()
    if !ok {
        return nil, fmt.Errorf("store: unknown provider %s", providerName)
    }
    return p, nil
}
```

效仿了Go标准库的database/sql包采用的方式，factory包采用了一个map类型数据，对工厂可以“生产”的、满足Store接口的实例类型进行管理。factory包还提供了Register函数，**让各个实现Store接口的类型可以把自己“注册”到工厂中来**。



一个基于内存map的Store接口的实现:

```go
// internal/store/memstore.go
package store

import (
    mystore "bookstore/store"
    factory "bookstore/store/factory"
    "sync"
)

func init() {
    factory.Register("mem", &MemStore{
        books: make(map[string]*mystore.Book),
    })
}

// MemStore 是一个基于内存map的Store接口的实现
type MemStore struct {
    sync.RWMutex
    books map[string]*mystore.Book
}

// ...具体实现方法
```

init函数中调用factory包提供的Register函数，把自己的实例以“mem”的名称注册到factory中的。这样做有一个好处，依赖Store接口进行图书数据管理的一方，只要导入internal/store这个包，就可以自动完成注册动作了。

```go
import (
  ... ...
  _ "bookstore/internal/store" // internal/store将自身注册到factory中
)

func main() {
    s, err := factory.New("mem") // 创建名为"mem"的图书数据存储模块实例
    if err != nil {
        panic(err)
    }
    ... ...
}   
```

#### HTTP服务模块（server）

HTTP服务模块的职责是**对外提供HTTP API服务，处理来自客户端的各种请求，并通过Store接口实例执行针对图书数据的相关操作**。

#### 编译、运行与验证

```sh
$ curl -X POST -H "Content-Type:application/json" -d '{"id": "978-7-111-55842-2", "name": "The Go Programming Language", "authors":["Alan A.A.Donovan", "Brian W. Kergnighan"],"press": "Pearson Education"}' localhost:8080/book



$ curl -X GET -H "Content-Type:application/json" localhost:8080/book/978-7-111-55842-2
{"id":"978-7-111-55842-2","name":"The Go Programming Language","authors":["Alan A.A.Donovan","Brian W. Kergnighan"],"press":"Pearson Education"}
```

### 思考题 🔖

> 基于nosql数据库，怎么实现一个新store.Store接口的实现？



# 基础篇：“脑勤”多理解

## 10 变量声明：静态语言有别于动态语言的重要特征

在编程语言中，为了方便操作内存特定位置的数据，**用一个特定的名字与位于特定位置的内存块绑定在一起**，这个名字被称为==变量==。

变量所绑定的内存区域是要有一个明确的**边界**的。

> 编程语言的编译器或解释器是如何知道一个变量所能引用的内存区域边界呢？

动态语言的解释器可以在运行时通过**对变量赋值的分析**，自动确定变量的边界。

静态语言通过==变量声明==，语言使用者可以显式告知编译器一个变量的边界信息。

### 10.1 Go语言的变量声明方法

![](images/image-20240703175402113.png)

变量声明分为四个部分：

- var是修饰变量声明的关键字；
- a为变量名；
- int为该变量的类型；
- 10是变量的初值。

如果没有显式为变量赋予初值，Go编译器会为变量赋予这个类型的**零值**。

![](images/image-20240703175543137.png)

==变量声明块（block）==：

```go
var (
  a int = 128
  b int8 = 6
  s string = "hello"
  c rune = 'A'
  t bool = true
)
```

```go
var a, b, c int = 5, 6, 7


var (
  a, b, c int = 5, 6, 7
  c, d, e rune = 'C', 'D', 'E'
) 
```

两种变量声明的“语法糖”:

1️⃣省略类型信息的声明

```go
var b = 13
```

Go编译器会根据右侧变量初值自动推导出变量的类型，并给这个变量赋予初值所对应的默认类型。

```go
var a, b, c = 12, 'A', "hello"
```

2️⃣短变量声明

```go
a := 12
b := 'A'
c := "hello"
```

Go语言的变量可以分为两类：

1. 一类称为==包级变量(package varible)==，也就是在包级别可见的变量。如果是导出变量（大写字母开头），那么这个包级变量也可以被视为==全局变量==；
2. 另一类则是==局部变量(local varible)==，也就是Go函数或方法体内声明的变量，仅在函数或方法体内可见。

### 10.2 包级变量的声明形式

包级变量只能使用带有var关键字的变量声明形式，不能使用短变量声明形式，但在形式细节上可以有一定灵活度。

可以从“变量声明时是否延迟初始化”这个角度，对包级变量的声明形式进行一次分类。

#### 第一类：声明并同时显式初始化

```go
// $GOROOT/src/io/io.go
var ErrShortWrite = errors.New("short write")
var ErrShortBuffer = errors.New("short buffer")
var EOF = errors.New("EOF")
```

种类可以使用省略类型信息的“语法糖”格式：

```go
var varName = initExpression
```

如果不接受默认类型，**要显式地为包级变量指定类型**。

```go
//第一种：
var a = 13 // 使用默认类型
var b int32 = 17  // 显式指定类型
var f float32 = 3.14 // 显式指定类型

//第二种：
var a = 13 // 使用默认类型
var b = int32(17) // 显式指定类型
var f = float32(3.14) // 显式指定类型
```

#### 第二类：声明但延迟初始化

**声明聚类**

通常将同一类的变量声明放在一个var变量声明块中，不同类的声明放在不同的var声明块中：

```go
// $GOROOT/src/net/net.go

var (
    netGo  bool 
    netCgo bool 
)

var (
    aLongTimeAgo = time.Unix(1, 0)
    noDeadline = time.Time{}
    noCancel   = (chan struct{})(nil)
)
```

通常也将延迟初始化的变量声明放在一个var声明块(比如上面的第一个var声明块)，然后将声明且显式初始化的变量放在另一个var块中（比如上面的第二个var声明块）。



**就近原则**:尽可能在靠近第一次使用变量的位置声明这个变量。

```go
// $GOROOT/src/net/http/request.go

var ErrNoCookie = errors.New("http: named cookie not present")
func (r *Request) Cookie(name string) (*Cookie, error) {
    for _, c := range readCookies(r.Header, name) {
        return c, nil
    }
    return nil, ErrNoCookie
}
```

如果一个包级变量在包内部被多处使用，那么这个变量还是放在源文件头部声明比较适合的。

### 10.3 局部变量的声明形式

1. 第一类：对于延迟初始化的局部变量声明，我们采用通用的变量声明形式

```go
var err error
```

2. 第二类：对于声明且显式初始化的局部变量，建议使用短变量声明形式

```go
a := 17
f := 3.14
s := "hello, gopher!"

a := int32(17)
f := float32(3.14)
s := []byte("hello, gopher!")
```

> **尽量在分支控制时使用短变量声明形式。**



### 小结

![](images/image-20240703180658029.png)

### 思考题

> Go语言变量声明中，类型是放在变量名的后面的，有什么好处？



## 11 代码块与作用域：如何保证变量不会被遮蔽？

==变量遮蔽（Variable Shadowing）==

代码块（Block）

作用域（Scope）

### 11.1 代码块与作用域

Go语言中的==代码块(Block)==是包裹在一对大括号内部的声明和语句序列。

如果一对大括号内部没有任何声明或其他语句，叫做==空代码块==。Go代码块支持嵌套。

```go
func foo() { //代码块1
    { // 代码块2
        { // 代码块3
            { // 代码块4

            }
        }
    }
}
```

1-4都是==显式代码块（Explicit Blocks）==

==隐式代码块（Implicit Block）==是没有显式代码块那样的肉眼可见的配对大括号包裹，无法通过大括号来识别隐式代码块。

![](images/image-20240708113446638.png)

隐式代码块从大到小：

- ==宇宙代码块（Universe Block）==，所有Go源码都在这个隐式代码块中

- ==包代码块（Package Block）==，每个Go包都对应一个隐式包代码块

- ==文件代码块（File Block）==，

- 控制语句层面（if、for与switch）。注意，这里的控制语句隐式代码块与控制语句使用大括号包裹的显式代码块并不是一个代码块。

- case/default子句

==作用域(Scope)==的概念是**针对标识符的，不局限于变量**。每个标识符都有自己的作用域，而**一个标识符的作用域就是指这个标识符在被声明后可以被有效使用的源码区域**。

作用域是一个编译期的概念，也就是说，编译器在编译过程中会对每个标识符的作用域进行检查，对于在标识符作用域外使用该标识符的行为会给出编译错误的报错。

**声明于外层代码块中的标识符，其作用域包括所有内层代码块**。

- 首先看看位于最外层的宇宙隐式代码块的标识符。

我们并不能声明这一块的标识符，因为这一区域是Go语言==预定义标识符==：

![](images/image-20240703182641189.png)

由于这些预定义标识符位于包代码块的外层，所以它们的作用域是范围最大的，对于开发者而言，它们的作用域就是源代码中的任何位置。不过，**这些预定义==标识符==不是关键字**，我们**同样可以在内层代码块中声明同名的标识符**。

- 第二个问题：既然宇宙代码块里存在预定义标识符，而且宇宙代码块的下一层是包代码块，那还有哪些标识符具有包代码块级作用域呢？

答案是，**包顶层声明**中的常量、类型、变量或函数（不包括方法）对应的标识符的作用域是包代码块。

不过，对于作用域为包代码块的标识符，一个特殊情况就是当一个包A导入另外一个包B后，包A仅可以使用被导入包包B中的==导出标识符（Exported Identifier）==。

> 什么是导出标识符呢？
> 
> 按照Go语言定义，一个标识符要成为导出标识符需同时具备两个条件：
> 
> - 一是这个标识符声明在包代码块中，或者它是一个字段名或方法名；
> - 二是它名字第一个字符是一个大写的Unicode字符。
> 
> 这两个条件缺一不可。

```go
func (t T) M1(x int) (err error) {
// 代码块1
    m := 13

    // 代码块1是包含m、t、x和err三个标识符的最内部代码块
    { // 代码块2

        // "代码块2"是包含类型bar标识符的最内部的那个包含代码块
        type bar struct {} // 类型标识符bar的作用域始于此
        { // 代码块3

            // "代码块3"是包含变量a标识符的最内部的那个包含代码块
            a := 5 // a作用域开始于此
            {  // 代码块4 
                //... ...
            }
            // a作用域终止于此
        }
        // 类型标识符bar的作用域终止于此
    }
    // m、t、x和err的作用域终止于此
}
```

上面示例中定义了类型T的一个方法M1，**方法接收器**(receiver)变量t、函数参数x，以及返回值变量err对应的标识符的作用域范围是M1函数体对应的显式代码块1。虽然t、x和err并没有被函数体的大括号所显式包裹，但它们属于函数定义的一部分，所以作用域依旧是代码块1。

**函数内部声明的常量或变量对应的标识符**的作用域范围开始于常量或变量声明语句的末尾，并终止于其最内部的那个包含块的末尾。在上述例子中，变量m、自定义类型bar以及在代码块3中声明的变量a均符合这个划分规则。

位于控制语句隐式代码块中的标识符的作用域划分

```go
func bar() {
  if a := 1; false {
  } else if b := 2; false {
  } else if c := 3; false {
  } else {
    println(a, b, c)
  }
} 
```

将上面示例中隐式代码块转换为显式代码块后:

```go
func bar() {
    { // 等价于第一个if的隐式代码块
        a := 1 // 变量a作用域始于此
        if false {

        } else {
            { // 等价于第一个else if的隐式代码块
                b := 2 // 变量b的作用域始于此
                if false {

                } else {
                    { // 等价于第二个else if的隐式代码块
                        c := 3 // 变量c作用域始于此
                        if false {

                        } else {
                            println(a, b, c)
                        }
                        // 变量c的作用域终止于此
                    }
                }
                // 变量b的作用域终止于此
            }
        }
        // 变量a作用域终止于此
    }
}
```

### 11.2 避免变量遮蔽的原则

**一个变量的作用域起始于其声明所在的代码块，并且可以一直扩展到嵌入到该代码块中的所有内层代码块**，而正是这样的作用域规则，成为了滋生“变量遮蔽问题”的土壤。

变量遮蔽问题的根本原因，就是**内层代码块中声明了一个与外层代码块同名且同类型的变量**。

```go
var a int = 2023

func checkYear() error {
	err := errors.New("wrong year")

	switch a, err := getYear(); a { // 遮蔽包代码块中的变量(a); 遮蔽外层显式代码块中的变量(err)
	case 2023:
		fmt.Println("it is", a, err)
	case 2024:
		fmt.Println("it is", a)
		err = nil
	}
	fmt.Println("after check, it is", a)
	return err
}

type new int // 遮蔽预定义标识符

func getYear() (new, error) {
	var b int16 = 2024
	return new(b), nil
}

func main() {
	err := checkYear()
	if err != nil {
		fmt.Println("call checkYear error:", err)
		return
	}
	fmt.Println("call checkYear ok")
}
```

- 遮蔽预定义标识符
- 遮蔽包代码块中的变量
- 遮蔽外层显式代码块中的变量

短变量声明与控制语句的结合十分容易导致变量遮蔽问题，并且很不容易识别。

### 11.3 利用工具检测变量遮蔽问题

`go vet`工具可以用于对Go源码做一系列静态检查，量遮蔽检查的插件需要单独安装：

```sh
$ sudo go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
```



```shell
$ go vet -vettool=$(which shadow) ./complex.go
# command-line-arguments
# [command-line-arguments]
./complex.go:15:12: declaration of "err" shadows declaration at line 13
```

go vet只给出了err变量被遮蔽的提示，变量a以及预定义标识符new被遮蔽的情况并没有给出提示。只能作为辅助作用。



## 12 基本数据类型：Go原生支持的数值类型有哪些？

类型不仅是静态语言编译器的要求，更是我们**对现实事物进行抽象的基础**。

Go语言的类型大体可分为三种：==基本数据类型==、==复合数据类型==和==接口类型==。

Go语言原生支持的==数值类型==包括**整型、浮点型以及复数类型**。

### 12.1 被广泛使用的整型

Go语言的整型，主要用来表示现实世界中整型数量，比如：人的年龄、班级人数等。

它可以分为==平台无关整型==和==平台相关整型==，区别是在**不同CPU架构或操作系统下长度是否是一致**。

#### 平台无关整型

![](images/image-20240703183721386.png)

![](images/image-20240703183955257.png)

Go采用**2的补码（Two’s Complement）**作为整型的比特位编码方法。因此，不能简单地将最高比特位看成负号，把其余比特位表示的值看成负号后面的数值。Go的补码是通过**原码逐位取反后再加1**得到的，比如，以-127这个值为例，它的补码转换过程就是这样的：

![](images/image-20240708154448556.png)

#### 平台相关整型

![](images/image-20240703183824577.png)

三个平台相关整型：`int`、`uint`与`uintptr`。

**在编写有移植性要求的代码时，千万不要强依赖这些类型的长度**。

```go
var a, b = int(5), uint(6)
var p uintptr = 0x12345678
fmt.Println("signed integer a's length is", unsafe.Sizeof(a))   // 8
fmt.Println("unsigned integer b's length is", unsafe.Sizeof(b)) // 8
fmt.Println("uintptr's length is", unsafe.Sizeof(p))            // 8
```

#### 整型的溢出问题

由于整型无法表示它溢出后的那个“结果”，所以出现溢出情况后，对应的整型变量的值依然会落到它的取值范围内，只是结果值与预期不符，导致程序逻辑出错。

```go
var s int8 = 127
s += 1 // 预期128，实际结果-128

var u uint8 = 1
u -= 2 // 预期-1，实际结果255
```

最容易发生在**循环语句的结束条件判断**中。

#### 字面值与格式化输出

Go语言在设计开始，继承了C语言的==数值字面值（Number Literal）==的语法形式：

```go
a := 53        // 十进制
b := 0700      // 八进制，以"0"为前缀
c1 := 0xaabbcc // 十六进制，以"0x"为前缀
c2 := 0Xddeeff // 十六进制，以"0X"为前缀
```

Go1.13版本又增加了对二进制字面值的支持和两种八进制字面值的形式:

```go
d1 := 0b10000001 // 二进制，以"0b"为前缀
d2 := 0B10000001 // 二进制，以"0B"为前缀
e1 := 0o700      // 八进制，以"0o"为前缀
e2 := 0O700      // 八进制，以"0O"为前缀 
```

为提升字面值的可读性，Go 1.13版本还支持在字面值中增加数字分隔符“`_`”:

```go
a := 5_3_7   // 十进制: 537
b := 0b_1000_0111  // 二进制位表示为10000111 
c1 := 0_700  // 八进制: 0700
c2 := 0o_700 // 八进制: 0700
d1 := 0x_5c_6d // 十六进制：0x5c6d
```

通过标准库fmt包的格式化输出函数，可将一个整型变量输出为不同进制的形式：

```go
var a int8 = 59
fmt.Printf("%b\n", a) //输出二进制：111011
fmt.Printf("%d\n", a) //输出十进制：59
fmt.Printf("%o\n", a) //输出八进制：73
fmt.Printf("%O\n", a) //输出八进制(带0o前缀)：0o73
fmt.Printf("%x\n", a) //输出十六进制(小写)：3b
fmt.Printf("%X\n", a) //输出十六进制(大写)：3B
```



### 12.2 浮点型

和使用广泛的整型相比，浮点型的使用场景就相对聚焦了，主要集中在**科学数值计算、图形图像处理和仿真、多媒体游戏以及人工智能**等领域。

#### 浮点型的二进制表示

[IEEE 754标准](https://zh.wikipedia.org/wiki/IEEE_754)是IEEE制定的二进制浮点数算术标准，它是20世纪80年代以来最广泛使用的浮点数运算标准，被许多CPU与浮点运算器采用。现存的大部分主流编程语言，包括Go语言，都提供了符合IEEE 754标准的浮点数格式与算术运算。

EE 754标准规定了四种表示浮点数值的方式：**单精度（32位）、双精度（64位）**、扩展单精度（43比特以上）与扩展双精度（79比特以上，通常以80位实现）。后两种其实很少使用

`float32`与`float64`（没有`float`），默认值都为`0.0`，占用的内存**空间大小**不同，可以表示的浮点数的**范围与精度**也不同。

IEEE 754规范给出了在内存中存储和表示一个浮点数的标准形式：

![](images/image-20240703185755005.png)

符号位、阶码（即经过换算的指数），以及尾数

![](images/image-20240703185844655.png)

当符号位为1时，浮点值为负值；当符号位为0时，浮点值为正值。公式中offset被称为==阶码偏移值==。

![](images/image-20240703185935111.png)

- 单精度浮点类型（float32）为符号位分配了1个bit，为阶码分配了8个bit，剩下的23个bit分给了尾数。
- 双精度浮点类型，除了符号位的长度与单精度一样之外，其余两个部分的长度都要远大于单精度浮点型，阶码可用的bit位数量为11，尾数则更是拥有了52个bit位。

> 如何将一个十进制形式的浮点值139.8125，转换为IEEE 754规定中的那种单精度二进制表示？

- 步骤一：把这个浮点数值的整数部分和小数部分，分别转换为二进制形式（后缀d表示十进制数，后缀b表示二进制数）：

整数部分：139d => 10001011b；

小数部分：0.8125d => 0.1101b（十进制小数转换为二进制可采用“乘2取整”的竖式计算）。

原浮点值`139.8125d`就变成`10001011.1101b`。

- 步骤二：移动小数点，直到整数部分仅有一个1，也就是`10001011.1101b` => `1.00010111101b`。

小数点向左移了7位，这样指数就为7，尾数为`00010111101b`。

- 步骤三：计算阶码。

IEEE754规定不能将小数点移动得到的指数，直接填到阶码部分，指数到阶码还需要一个转换过程。对于float32的单精度浮点数而言，==阶码 = 指数 + 偏移值==。

偏移值的计算公式为`2^(e-1)-1`，其中e为阶码部分的bit位数，这里为8，于是单精度浮点数的阶码偏移值就为`2^(8-1)-1 = 127`。

这样在这个例子中，阶码 = 7 + 127 = 134d =10000110b。float64的双精度浮点数的阶码计算也是这样的。

- 步骤四：将符号位、阶码和尾数填到各自位置，得到最终浮点数的二进制表示。尾数位数不足23位，可在后面补0。

![](images/image-20241203091305828.png)

最终浮点数`139.8125d`的二进制表示就为`0b_0_10000110_00010111101_000000000000`。

```go
var f float32 = 139.8125
bits := math.Float32bits(f)
fmt.Printf("%b\n", bits)  // 1000011000010111101000000000000
```

阶码和尾数的长度决定了浮点类型可以表示的浮点数范围与精度。

双精度浮点类型（float64）阶码与尾数使用的比特位数更多，它可以表示的精度要远超单精度浮点类型，所以在日常开发中，我们使用双精度浮点类型（float64）的情况更多，这也是Go语言中浮点常量或字面值的默认类型。

float32由于表示范围与精度有限，会出现一些问题，如：

```go
var f1 float32 = 16777216.0
var f2 float32 = 16777217.0
fmt.Println(f1 == f2) // true
```



#### 字面值与格式化输出

1. 一类是直白地用十进制表示的浮点值形式

```go
3.1415
.15  // 整数部分如果为0，整数部分可以省略不写
81.80
82. // 小数部分如果为0，小数点后的0可以省略不写
```

2. 另一类则是科学计数法形式

```go
6674.28e-2 // 6674.28 * 10^(-2) = 66.742800
.12345E+5  // 0.12345 * 10^5 = 12345.000000


0x2.p10  // 2.0 * 2^10 = 2048.000000
0x1.Fp+0 // 1.9375 * 2^0 = 1.937500
```

```go
var f float64 = 123.45678
fmt.Printf("%f\n", f) // 123.456780

fmt.Printf("%e\n", f) // 1.234568e+02
fmt.Printf("%x\n", f) // 0x1.edd3be22e5de1p+06
```

### 12.3 复数类型

🔖



### 12.4 延展：创建自定义的数值类型

- 可以通过type关键字基于原生数值类型来声明一个新类型:

```go
type MyInt int32
```

MyInt类型的底层类型是int32，它的数值性质与int32完全相同，但它们仍然是**完全不同的两种类型**。

```go
var m int = 5
var n int32 = 6
var a MyInt = m // 错误：在赋值中不能将m（int类型）作为MyInt类型使用
var a MyInt = n // 错误：在赋值中不能将n（int32类型）作为MyInt类型使用   
```

- 也可以通过Go提供的**类型别名（Type Alias）**语法来自定义数值类型。

和上面使用标准type语法的定义不同的是，通过类型别名语法定义的新类型与原类型别无二致，可以完全相互替代。

```go
type MyInt = int32

var n int32 = 6
var a MyInt = n // ok
```

### 思考题

> 下面例子中f1为何会与f2相等？
> 
> ```go
> var f1 float32 = 16777216.0
> var f2 float32 = 16777217.0
> f1 == f2 // true
> ```



## 13 基本数据类型：为什么Go要原生支持字符串类型？

### 13.1 原生支持字符串有什么好处？

C语言没有提供对字符串类型的原生支持，字符串是以字符串字面值或以’`\0`’结尾的字符类型数组来呈现的。

```c
#define GO_SLOGAN "less is more"
const char * s1 = "hello, gopher"
char s2[] = "I love go"   
```

这样定义的非原生字符串在使用过程中会有很多问题，比如：

- 不是原生类型，编译器不会对它进行**类型校验**，导致类型安全性差；
- 字符串操作时要时刻考虑结尾的’\0’，防止缓冲区溢出；
- 以字符数组形式定义的“字符串”，它的值是可变的，在并发场景中需要考虑同步问题；
- 获取一个字符串的长度代价较大，通常是O(n)时间复杂度；
- C语言没有内置对非ASCII字符（如中文字符）的支持。

Go语言通过string类型统一了对“字符串”的抽象。这样无论是**字符串常量、字符串变量**或是代码中出现的**字符串字面值**，它们的类型都被统一设置为`string`。

```go
const (
    GO_SLOGAN = "less is more" // GO_SLOGAN是string类型常量
    s1 = "hello, gopher"       // s1是string类型常量
)

var s2 = "I love go" // s2是string类型变量
```

这样的设计带来的好处：

- 第一点：string类型的数据是**==不可变==**的，提高了字符串的==并发安全性==和==存储利用率==。

Go字符串可以被多个Goroutine共享，开发者不用因为担心并发安全问题。

由于字符串的不可变性，针对同一个字符串值，无论它在程序的几个位置被使用，Go编译器只需要为它分配一块存储就好了，大大提高了存储利用率。

- 第二点：没有结尾’\0’，而且获取长度的时间复杂度是常数时间，消除了获取字符串长度的开销。 

在C语言中，获取一个字符串的长度可以调用标准库的strlen函数，这个函数的实现原理是遍历字符串中的每个字符并做计数，直到遇到字符串的结尾’\0’停止。显然这是一个线性时间复杂度的算法，执行时间与字符串中字符个数成正比。并且，它存在一个约束，那就是传入的字符串必须有结尾’\0’，结尾’\0’是字符串的结束标志。

Go字符串中没有结尾’\0’，获取字符串长度更不需要结尾’\0’作为结束标志。并且，Go**获取字符串长度是一个常数级时间复杂度**，无论字符串中字符个数有多少，都可以快速得到字符串的长度值。

- 第三点：原生支持“所见即所得”的原始字符串，大大降低构造多行字符串时的心智负担。

```go
var s string = `         ,_---~~~~~----._
    _,,_,*^____      _____*g*\"*,--,
   / __/ /'     ^.  /      \ ^@q   f
  [  @f | @))    |  | @))   l  0 _/
   \/   \~____ / __ \_____/     \
    |           _l__l_           I
    }          [______]           I
    ]            | | |            |
    ]             ~ ~             |
    |                            |
     |                           |`
fmt.Println(s)
```

- 第四点：对非ASCII字符提供原生支持，消除了源码在不同环境下显示乱码的可能。

Go字符串中的每个字符都是一个Unicode字符，并且这些Unicode字符是以UTF-8编码格式存储在内存当中的。

### 13.2 Go字符串的组成

Go语言在看待Go字符串组成这个问题上，有两种视角。

一种是**字节视角**，也就是和所有其它支持字符串的主流语言一样，Go语言中的字符串值也是一个可空的==字节序列==，字节序列中的字节个数称为该字符串的长度。一个个的字节只是孤立数据，不表意。

```go
var s = "中国人"
fmt.Printf("the length of s = %d\n", len(s)) // 9

for i := 0; i < len(s); i++ {
    fmt.Printf("0x%x ", s[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
}
fmt.Printf("\n")
```

如果要表意，就需要从字符串的另外一个视角来看：字符串是由一个可空的==字符序列==构成。

```go
var s = "中国人"
fmt.Println("the character count in s is", utf8.RuneCountInString(s)) // 3

for _, c := range s {
    fmt.Printf("0x%x ", c) // 0x4e2d 0x56fd 0x4eba
}
fmt.Printf("\n")
```

Go采用的是Unicode字符集，每个字符都是一个Unicode字符，那么这里输出的0x4e2d、0x56fd和0x4eba就应该是某种Unicode字符的表示了。以0x4e2d为例，它是汉字“中”在Unicode字符集表中的==码点（Code Point）==。

> Unicode字符集中的每个字符，都被分配了统一且唯一的字符编号。所谓**Unicode码点**，就是指将Unicode字符集中的所有字符“排成一队”，字符在这个“队伍”中的**位次**，就是它在Unicode字符集中的码点。也就说，一个码点唯一对应一个字符。

#### rune类型与字符字面值

Go使用`rune`这个类型来表示一个Unicode码点。rune本质上是int32类型的别名类型，它与int32类型是完全等价的：

```go
// $GOROOT/src/builtin.go
type rune = int32
```

**一个rune实例就是一个Unicode字符，一个Go字符串也可以被视为rune实例的集合。可以通过字符字面值来初始化一个rune变量。**

字符字面值有多种表示法:

- 单引号括起的字符字面值

```go
'a'  // ASCII字符
'中' // Unicode字符集中的中文字符
'\n' // 换行字符
'\'' // 单引号字符
```

- Unicode专用的转义字符\u或\U作为前缀

```go
'\u4e2d'     // 字符：中
'\U00004e2d' // 字符：中
'\u0027'     // 单引号字符
```

`\u`后面接两个十六进制数。如果是用两个十六进制数无法表示的Unicode字符，可以使用`\U`，\U后面可以接四个十六进制数来表示一个Unicode字符。

- 由于表示码点的rune本质上就是一个整型数，还可**用整型值来直接作为字符字面值给rune变量赋值**

```go
'\x27'  // 使用十六进制表示的单引号字符
'\047'  // 使用八进制表示的单引号字符
```

#### 字符串字面值

**把表示单个字符的单引号，换为表示多个字符组成的字符串的双引号**

```go
"abc\n"
"中国人"
"\u4e2d\u56fd\u4eba" // 中国人
"\U00004e2d\U000056fd\U00004eba" // 中国人
"中\u56fd\u4eba" // 中国人，不同字符字面值形式混合在一起
"\xe4\xb8\xad\xe5\x9b\xbd\xe4\xba\xba" // 十六进制表示的字符串字面值：中国人
```

> 最后一行使用的是十六进制形式的字符串字面值，但每个字节的值与前面几行的码点值完全对应不上啊，这是为什么呢？
> 
> utf8

#### UTF-8编码方案

UTF-8编码解决的是**Unicode码点值在计算机中如何存储和表示（位模式）的问题**。

> 码点唯一确定一个Unicode字符，直接用码点值不行么？

这的确是可以的，并且UTF-32编码标准就是采用的这个方案。UTF-32编码方案**==固定==使用4个字节表示每个==Unicode字符码点==**，这带来的好处就是编解码简单，但缺点也很明显，主要有下面几点：

- 这种编码方案使用4个字节存储和传输一个整型数的时候，需要考虑不同平台的字节序问题;
- 由于采用4字节的固定长度编码，与采用1字节编码的ASCII字符集无法兼容；
- 所有Unicode字符码点都用4字节编码，显然空间利用率很差。

针对这些问题，Go语言之父Rob Pike发明了UTF-8编码方案。和UTF-32方案不同，UTF-8方案使用**变长度字节**，对Unicode字符的码点进行编码。编码采用的字节数量与Unicode字符在码点表中的序号有关：**表示序号（码点）小的字符使用的字节数量少，表示序号（码点）大的字符使用的字节数多。**

UTF-8编码使用的字节数量从1个到4个不等。

- 前128个与ASCII字符重合的码点（U+0000~U+007F）使用1个字节表示；
- 带变音符号的拉丁文、希腊文、西里尔字母、阿拉伯文等使用2个字节来表示；
- 而东亚文字（包括汉字）使用3个字节表示；
- 其他极少使用的语言的字符则使用4个字节表示。

这样的编码方案是兼容ASCII字符内存表示的，这意味着采用UTF-8方案在内存中表示Unicode字符时，已有的ASCII字符可以被直接当成Unicode字符进行存储和传输，不用再做任何改变。

此外，UTF-8的编码单元为一个字节（也就是一次编解码一个字节），所以我们在处理UTF-8方案表示的Unicode字符的时候，就不需要像UTF-32方案那样考虑字节序问题了。相对于UTF-32方案，UTF-8方案的**空间利用率**也是最高的。

现在，<u>UTF-8编码方案已经成为Unicode字符编码方案的**事实标准**</u>，各个平台、浏览器等默认均使用UTF-8编码方案对Unicode字符进行编、解码。Go语言也不例外，采用了UTF-8编码方案存储Unicode字符，我们在前面按字节输出一个字符串值时看到的字节序列，就是对字符进行UTF-8编码后的值。

使用Go在标准库中提供的UTF-8包，对Unicode字符（rune）进行编解码:

```go
// rune -> []byte
func encodeRune() {
    var r rune = 0x4e2d
    fmt.Printf("这个unicode字符是：%c\n", r)
    buf := make([]byte, 3)
    _ = utf8.EncodeRune(buf, r) // 对rune进行utf8编码
    fmt.Printf("这个字符的utf8描述为：0x%X\n", buf)
}

// []byte -> rune
func decodeRune() {
    var buf = []byte{0xe4, 0xb8, 0xad}
    r, _ := utf8.DecodeRune(buf) // 对buf进行utf8解码
    fmt.Printf("字节序列解码后的unicode字符是：%s\n", string(r))
}
```

### 13.3 Go字符串类型的内部表示 ❤️

Go字符串类型的优秀的性质，与其在**编译器和运行时中的内部表示**是分不开的。

```go
// $GOROOT/src/reflect/value.go

// StringHeader是一个string的运行时表示
type StringHeader struct {
    Data uintptr
    Len  int
}
```

string类型其实是一个“描述符”，它本身并不真正存储字符串数据，而仅是由一个**指向底层存储的指针和字符串的长度**字段组成的。

![](images/image-20240708164110090.png)

Go编译器把源码中的string类型映射为运行时的一个**二元组（Data, Len）**，真实的字符串值数据就存储在一个被Data指向的底层数组中。

```go
// 字符串内部表示
func str4() {
    var s = "hello"
    hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // 将string类型变量地址显式转型为reflect.StringHeader
    fmt.Printf("0x%x\n", hdr.Data)                     // 0x10a30e0
    p := (*[5]byte)(unsafe.Pointer(hdr.Data))          // 获取Data字段所指向的数组的指针
    dumpBytesArray((*p)[:])                            // [h e l l o ]   // 输出底层数组的内容
}
func dumpBytesArray(arr []byte) {
    fmt.Printf("[")
    for _, b := range arr {
        fmt.Printf("%c ", b)
    }
    fmt.Printf("]\n")
}
```

🔖 新版本reflect.StringHeader过期

### 13.4 Go字符串类型的常见操作

#### 下标操作

#### 字符迭代

常规for迭代与for range迭代

> 注意，这两种形式的迭代对字符串进行操作得到的结果是不同的。

#### 字符串连接

`+`/`+=`

#### 字符串比较

`==`、`!=` 、`>=`、`<=`、`>` 和 `<`

#### 字符串转换

Go支持字符串与字节切片、字符串与rune切片的双向转换，并且这种转换无需调用任何函数，只需使用显式类型转换就可以了。

```go
var s string = "中国人"

// string -> []rune
rs := []rune(s) 
fmt.Printf("%x\n", rs) // [4e2d 56fd 4eba]

// string -> []byte
bs := []byte(s) 
fmt.Printf("%x\n", bs) // e4b8ade59bbde4baba

// []rune -> string
s1 := string(rs)
fmt.Println(s1) // 中国人

// []byte -> string
s2 := string(bs)
fmt.Println(s2) // 中国人
```

### 思考题

> Go提供多种字符串连接服务，包括基于+/+=的字符连接、基于strings.Builder、strings.Join、fmt.Sprintf等函数来进行字符串连接操作。那么，哪种连接方式是性能最高的呢？



## 14 常量：Go在“常量”设计上的创新有哪些？

Go语言在常量方面的创新：

- 支持无类型常量；
- 支持隐式自动转型；
- 可用于实现枚举。

### 14.1 常量以及Go原生支持常量的好处

Go语言的常量是一种**在源码编译期间被创建的语法元素**。

```go
const Pi float64 = 3.14159265358979323846 // 单行常量声明

// 以const代码块形式声明常量
const (
    size int64 = 4096
    i, j, s = 13, 14, "bar" // 单行声明多个常量
)
```

Go常量的类型只局限于基本数据类型，包括**数值类型、字符串类型，以及只有两个取值（true和false）的布尔类型**。

> 原生不支持常量的C语言
> 
> 在C语言中，**字面值担负着常量的角色**，可以使用数值型、字符串型字面值来应对不同场合对常量的需求。
> 
> 为了不让这些字面值以“魔数（Magic Number）”的形式分布于源码各处，早期C语言的常用实践是使用**宏（macro）**定义记号来指代这些字面值，这种定义方式被称为**宏定义常量**，比如：
> 
> ```c
> #define FILE_MAX_LEN 0x22334455
> #define PI 3.1415926
> #define GO_GREETING "Hello, Gopher"
> #define A_CHAR 'a'
> ```
> 
> **宏定义的常量会有很多问题**。比如，它是一种仅在预编译阶段进行替换的字面值，继承了宏替换的复杂性和易错性，而且还有类型不安全、无法在调试时通过宏名字输出常量的值，等等问题。
> 
> 后续C标准中提供的const关键字修饰的标识符，也依然不是一种圆满方案。因为const关键字修饰的标识符本质上依旧是变量，它甚至无法用作数组变量声明中的初始长度（除非用GNU扩展C）。
> 
> ```c
> const int size = 5;
> int a[size] = {1,2,3,4,5}; // size本质不是常量，这将导致编译器错误 
> ```

### 14.2 无类型常量

Go语言对类型安全是有严格要求的：**即便两个类型拥有着相同的底层类型，但它们仍然是不同的数据类型，不可以被相互比较或混在一个表达式中进行运算**。

这一要求不仅仅适用于变量，也同样适用于有类型常量（Typed Constant）。

```go
type myInt int
const n myInt = 13
const m int = n + 5 // 编译器报错：cannot use n + 5 (type myInt) as type int in const initializer

func main() {
    var a int = 5
    fmt.Println(a + n) // 编译器报错：invalid operation: a + n (mismatched types int and myInt)
}
```

而且，**有类型常量与变量混合在一起进行运算求值的时候，也必须遵守类型相同这一要求**，否则我们只能通过显式转型才能让上面代码正常工作：

```go
type myInt int
const n myInt = 13
const m int = int(n) + 5  // OK

func main() {
    var a int = 5
    fmt.Println(a + int(n))  // 输出：18
}
```

也可以使用Go中的==无类型常量（Untyped Constant）==来实现：

```go
type myInt int
const n = 13

func main() {
    var a myInt = 5
    fmt.Println(a + n)  // 输出：18
} 
```

常量n的默认类型int与myInt并不是同一个类型啊，为什么可以放在一个表达式中计算而没有报编译错误呢？

### 14.3 隐式转型

对于无类型常量参与的表达式求值，Go编译器会**根据上下文中的类型信息**，把无类型常量自动转换为相应的类型后，再参与求值计算，这一转型动作是隐式进行的。

但由于转型的对象是一个常量，所以这并不会引发类型安全问题，Go编译器会保证这一转型的安全性。

如果Go编译器在做隐式转型时，发现无法将常量转换为目标类型，Go编译器也会报错：

```go
const m = 1333333333

var k int8 = 1
j := k + m // 编译器报错：constant 1333333333 overflows int8 
```



### 14.4 实现枚举

Go语言其实并没有原生提供枚举类型。

Go语言中，可以使用**const代码块**定义的常量集合，来实现枚举。这是因为枚举类型本质上就是一个由**有限数量常量**所构成的集合。

Go将C语言枚举类型的这种基于前一个枚举值加1的特性，分解成了Go中的两个特性：**自动重复上一行，以及引入const块中的行==偏移量指示器==`iota`**，这样它们就可以分别独立使用了。

#### Go的const语法提供了“隐式重复前一个非空表达式”的机制

```go
const (
  Apple, Banana = 11, 22
  Strawberry, Grape 
  Pear, Watermelon 
)
```

这里常量定义的后两行并没有被显式地赋予初始值，所以Go编译器就为它们自动使用上一行的表达式，也就获得了下面这个等价的代码：

```go
const (
  Apple, Banana = 11, 22
  Strawberry, Grape  = 11, 22 // 使用上一行的初始化表达式
  Pear, Watermelon  = 11, 22 // 使用上一行的初始化表达式
) 
```

#### `iota`

iota是Go语言的一个==预定义标识符==，它表示的是const声明块（包括单行声明）中，每个常量所处位置在块中的**偏移值**（从零开始）。

同时，每一行中的iota自身也是一个无类型常量，可以像前面我们提到的无类型常量那样，自动参与到不同类型的求值过程中来，不需要我们再对它进行显式转型操作。

比如Go标准库中sync/mutex.go中的一段基于iota的枚举常量的定义：

```go
// $GOROOT/src/sync/mutex.go 
const ( 
    mutexLocked = 1 << iota
    mutexWoken
    mutexStarving
    mutexWaiterShift = iota
    starvationThresholdNs = 1e6
)
```

- 首先，这个const声明块的第一行是`mutexLocked = 1 << iota` ，iota的值是这行在const块中的偏移，因此iota的值为0，得到`mutexLocked`这个常量的值为`1 << 0`，也就是1。

- 接着，第二行：mutexWorken 。因为这个const声明块中并没有显式的常量初始化表达式，所以我们根据const声明块里“**隐式重复前一个非空表达式**”的机制，这一行就等价于mutexWorken = 1 << iota。而且，又因为这一行是const块中的第二行，所以它的偏移量iota的值为1，我们得到mutexWorken这个常量的值为1 << 1，也就是2。

- 然后是mutexStarving。这个常量同mutexWorken一样，这一行等价于mutexStarving = 1 << iota。而且，也因为这行的iota的值为2，我们可以得到mutexStarving这个常量的值为 1 << 2，也就是4;

- 再然后看mutexWaiterShift = iota 这一行，这一行为常量mutexWaiterShift做了显式初始化，这样就不用再重复前一行了。由于这一行是第四行，而且作为行偏移值的iota的值为3，因此mutexWaiterShift的值就为3。

- 最后一行，代码中直接用了一个具体值1e6给常量starvationThresholdNs进行了赋值，那么这个常量值就是1e6本身了。

> 提醒:位于同一行的iota即便出现多次，多个iota的值也是一样的.
> 
> ```go
> const (
>   Apple, Banana = iota, iota + 10 // 0, 10 (iota = 0)
>   Strawberry, Grape // 1, 11 (iota = 1)
>   Pear, Watermelon  // 2, 12 (iota = 2)
> )
> ```

```go
// $GOROOT/src/syscall/net_js.go
const (
    _ = iota
    IPV6_V6ONLY  // 1
    SOMAXCONN    // 2
    SO_ERROR     // 3
)
```

枚举常量值并不连续时，也可以借助空白标识符来实现：

```go
const (
  _ = iota // 0
  Pin1
  Pin2
  Pin3
  _
  Pin5    // 5   
) 
```

iota特性让**维护枚举常量列表变得更加容易**。

比如使用传统的枚举常量声明方式，来声明一组按首字母排序的“颜色”常量：

```go
const ( 
  Black  = 1 
  Red    = 2
  Yellow = 3
)
```

要增加一个新颜色Blue:

```go
const (
  Blue   = 1
  Black  = 2
  Red    = 3
  Yellow = 4
)
```

使用iota重新定义:

```go
const (
  _ = iota     
  Blue
  black
  Red 
  Yellow     
) 
```

这样，无论后期我们需要增加多少种颜色，只需将常量名插入到对应位置就可以，其他就不需要再做任何手工调整了。

如果一个Go源文件中有多个const代码块定义的不同枚举，每个const代码块中的iota也是独立变化的:

```go
const (
  a = iota + 1 // 1, iota = 0
  b            // 2, iota = 1
  c            // 3, iota = 2
)

const (
    i = iota << 1 // 0, iota = 0
    j             // 2, iota = 1
    k             // 4, iota = 2
)
```

> 内置包中iota使用例子：
>
> ```go
> // log/log.go
> const (
> 	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
> 	Ltime                         // the time in the local time zone: 01:23:23
> 	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
> 	Llongfile                     // full file name and line number: /a/b/c/d.go:23
> 	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
> 	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
> 	Lmsgprefix                    // move the "prefix" from the beginning of the line to before the message
> 	LstdFlags     = Ldate | Ltime // initial values for the standard logger
> )
> ```
>
> 

### 思考题

> 虽然iota带来了灵活性与便利，但是否存在一些场合，在定义枚举常量时使用显式字面值更为适合呢？



## 15 同构复合类型：从定长数组到变长切片

==复合类型==：由多个**同构类型（相同类型）**或**异构类型（不同类型）**的元素的值组合而成。

Go原生内置了多种复合数据类型，包括**数组、切片（slice）、map、结构体，以及channel**等。

### 15.1 数组有哪些基本特性？

Go语言的数组是一个长度固定的、由同构类型元素组成的连续序列。

两个重要属性：==元素的类型==和==数组长度==。

```go
var arr [N]T
```

数组元素的类型可以为**任意的Go原生类型或自定义类型**。

如果两个数组类型的元素类型T与数组长度N都是一样的，那么这两个**数组类型是等价**的，如果有一个属性不同，它们就是两个不同的数组类型。

**数组类型不仅是逻辑上的连续序列，而且在实际内存分配时也占据着一整块内存**。

![](images/image-20240708185652161.png)

预定义函数`len`可以用于获取一个数组类型变量的长度，通过unsafe包提供的`Sizeof`函数可以获得一个数组变量的总大小：

```go
var arr = [6]int{1, 2, 3, 4, 5, 6}
fmt.Println("数组长度：", len(arr))           // 6
fmt.Println("数组大小：", unsafe.Sizeof(arr)) // 48 
```

数组大小就是所有元素的大小之和。

声明一个数组类型变量，如果不进行显式初始化，那么数组中的元素值就是它类型的零值。

```go
var arr1 [6]int // [0 0 0 0 0 0]
```

显示初始化：

```go
var arr2 = [6]int {
  11, 12, 13, 14, 15, 16,
} // [11 12 13 14 15 16]

var arr3 = [...]int {
    21, 22, 23,
} // [21 22 23]
fmt.Printf("%T\n", arr3) // [3]int
```

可以忽略掉右值初始化表达式中数组类型的长度，用“`…`”替代，Go编译器会根据数组元素的个数，自动计算出数组长度。

对一个长度较大的稀疏数组进行显式初始化：

```go
var arr4 = [...]int{
  99: 39, // 将第100个元素(下标值为99)的值赋值为39，其余元素值均为0
}
fmt.Printf("%T\n", arr4) // [100]int
```

### 15.2 多维数组

```go
var mArr [2][3][4]int
```

![](images/image-20240704072318759.png)

数组类型变量是一个整体，这就意味着**一个数组变量表示的是整个数组**。这点与C语言完全不同，在C语言中，**数组变量可视为指向数组第一个元素的指针**。

这样一来，无论是参与迭代，还是作为实际参数传给一个函数/方法，Go传递数组的方式都是纯粹的**值拷贝**，这会带来较大的内存拷贝开销。解决办法：指针或切片。

### 15.3 切片❤️

数组两点不足：**固定的元素个数，以及传值机制下导致的开销较大**。

==切片（slice）==，来弥补数组的这两处不足。

切片的声明并初始化，相对于数组少了长度：

```go
var nums = []int{1,2,3,4,5,6}
```

```go
fmt.Println(len(nums)) // 6

nums = append(nums, 7) // 切片变为[1 2 3 4 5 6 7]
fmt.Println(len(nums)) // 7
```

#### Go是如何实现切片类型的？

Go切片在运行时其实是一个三元组结构:

```go
// runtime/slice.go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```

- array: 是指向底层数组的指针；
- len: 是切片的长度，即切片中当前元素的个数；
- cap: 是底层数组的长度，也是切片的最大容量，cap值永远大于等于len值。

![](images/image-20240704073320719.png)

**Go编译器会自动为每个新创建的切片**，建立一个底层数组，默认底层数组的长度与切片初始元素个数相同。

创建切片的其他方式：

- 方法一：通过`make`函数来创建切片，并指**定底层数组的长度**。

```go
sl := make([]byte, 6, 10) // 其中10为cap值，即底层数组长度，6为切片的初始长度

sl := make([]byte, 6) // cap = len = 6
```

- 方法二：采用`array[low : high : max]`语法基于一个已存在的数组创建切片。这种方式被称为==数组的切片化==

```go
arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
sl := arr[3:7:9]
```

![](images/image-20240704074015623.png)

基于数组创建的切片，它的起始元素从low所标识的下标值开始，切片的长度（len）是high - low，它的容量是max - low。而且，由于切片sl的底层数组就是数组arr，对切片sl中元素的修改将直接影响数组arr变量。比如，如果我们将切片的第一个元素加10，那么数组arr的第四个元素将变为14：

```go
sl[0] += 10
fmt.Println("arr[3] =", arr[3]) // 14
```

这样看来，**切片好比打开了一个访问与修改数组的“窗口”**，通过这个窗口，可以直接操作底层数组中的**部分元素**。这有些类似于操作文件之前打开的“**文件描述符**”（Windows上称为**句柄**），通过文件描述符我们可以对底层的真实文件进行相关操作。可以说，**切片之于数组就像是文件描述符之于文件**。

在Go语言中，数组更多是“退居幕后”，承担的是**底层存储空间**的角色。切片就是数组的“描述符”，也正是因为这一特性，切片才能在函数参数传递时**避免较大性能开销**。因为传递的并不是数组本身，而是数组的“描述符”，而这个**描述符的大小是固定的**（见上面的三元组结构），无论底层的数组有多大，切片打开的“窗口”长度有多长，它都是不变的。此外，在进行数组切片化的时候，通常省略max，而max的默认值为数组的长度。

另外，针对一个已存在的数组，我们还可以建立多个操作数组的切片，这些切片共享同一底层数组，切片对底层数组的操作也同样会**反映到其他切片中**。

![](images/image-20240704074617178.png)

- 方法三：基于切片创建切片。

这种切片的运行时表示原理与上面的是一样的。

#### 切片的动态扩容

“动态扩容”指当通过`append`操作向切片追加数据的时候，如果这时切片的len值和cap值是相等的，也就是说切片**底层数组已经没有空闲空间**再来存储追加的值了，Go运行时就会对这个切片做扩容操作，来保证切片始终能存储下追加的新值。也就是会重新分配了其底层数组。

```go
var s []int
s = append(s, 11) 
fmt.Println(len(s), cap(s)) //1 1
s = append(s, 12) 
fmt.Println(len(s), cap(s)) //2 2
s = append(s, 13) 
fmt.Println(len(s), cap(s)) //3 4
s = append(s, 14) 
fmt.Println(len(s), cap(s)) //4 4
s = append(s, 15) 
fmt.Println(len(s), cap(s)) //5 8
```

- 最开始，s初值为零值（nil），这个时候s没有“绑定”底层数组。先通过append操作向切片s添加一个元素11，这个时候，append会先分配底层数组u1（数组长度1），然后将s内部表示中的array指向u1，并设置len = 1, cap = 1;
- 接着，我们通过append操作向切片s再添加第二个元素12，这个时候len(s) = 1，cap(s) = 1，append判断底层数组剩余空间已经不能够满足添加新元素的要求了，于是它就创建了一个新的底层数组u2，长度为2（u1数组长度的2倍），并把u1中的元素拷贝到u2中，最后将s内部表示中的array指向u2，并设置len = 2, cap = 2；
- 然后，第三步，我们通过append操作向切片s添加了第三个元素13，这时len(s) = 2，cap(s) = 2，append判断底层数组剩余空间不能满足添加新元素的要求了，于是又创建了一个新的底层数组u3，长度为4（u2数组长度的2倍），并把u2中的元素拷贝到u3中，最后把s内部表示中的array指向u3，并设置len = 3, cap为u3数组长度，也就是4；
- 第四步，我们依然通过append操作向切片s添加第四个元素14，此时len(s) = 3,cap(s) = 4，append判断底层数组剩余空间可以满足添加新元素的要求，所以就把14放在下一个元素的位置(数组u3末尾），并把s内部表示中的len加1，变为4；
- 第五步又通过append操作，向切片s添加最后一个元素15，这时len(s) = 4，cap(s) = 4，append判断底层数组剩余空间又不够了，于是创建了一个新的底层数组u4，长度为8（u3数组长度的2倍），并将u3中的元素拷贝到u4中，最后将s内部表示中的array指向u4，并设置len = 5, cap为u4数组长度，也就是8。

总结，针对元素是int型的数组，新数组的容量是当前数组的**==2倍==**。新数组建立后，append会把旧数组中的数据拷贝到新数组中，之后新数组便成为了切片的底层数组，旧数组会被垃圾回收掉。

不过append操作的这种自动扩容行为，有些时候会给我们开发者带来一些困惑，比如基于一个已有数组建立的切片，一旦追加的数据操作触碰到切片的容量上限（实质上也是数组容量的上界)，切片就会和原数组==解除“绑定”==，后续对切片的任何修改都不会反映到原数组中了。

```go
u := [...]int{11, 12, 13, 14, 15}
fmt.Println("array:", u) // [11, 12, 13, 14, 15]
s := u[1:3]
fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
s = append(s, 24)
fmt.Println("after append 24, array:", u)
fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
s = append(s, 25)
fmt.Println("after append 25, array:", u)
fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
s = append(s, 26)
fmt.Println("after append 26, array:", u)
fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

s[0] = 22
fmt.Println("after reassign 1st elem of slice, array:", u)
fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
```

结果：

```go
array: [11 12 13 14 15]
slice(len=2, cap=4): [12 13]
after append 24, array: [11 12 13 24 15]
after append 24, slice(len=3, cap=4): [12 13 24]
after append 25, array: [11 12 13 24 25]
after append 25, slice(len=4, cap=4): [12 13 24 25]
after append 26, array: [11 12 13 24 25]
after append 26, slice(len=5, cap=8): [12 13 24 25 26]
after reassign 1st elem of slice, array: [11 12 13 24 25]
after reassign 1st elem of slice, slice(len=5, cap=8): [22 13 24 25 26]
```

### 思考题

下下面这两个切片变量sl1与sl2的差异:

```go
var sl1 []int
var sl2 = []int{}
```

`sl1` 是一个未初始化的切片，它指向一个空数组，而 `sl2` 是一个初始化的切片，它指向一个虽然为空但实际存在的数组。在使用切片时，需要确保其不为 `nil`，因为 `nil` 切片无法被索引，也不能进行其他操作。



## 16 复合数据类型：原生map类型的实现机制是怎样的？

map(映射、哈希表或字典)

map是切片之后，第二个由**Go编译器与运行时联合实现**的复合数据类型，它有着**复杂的内部实现**，但却提供了十分简单友好的开发者使用接口。

### 16.1 什么是map类型？

```go
map[key_type]value_type
```

```go
map[string]string // key与value元素的类型相同
map[int]string    // key与value元素的类型不同
```

如果两个map类型的**key元素类型相同，value元素类型也相同**，那么它们是**同一个map类型**，否则就是不同的map类型。

map类型对value的类型没有限制，但是对key的类型却有严格要求，因为map类型要保证**key的==唯一性==**，**key的类型必须支持“==”和“!=”两种比较操作符**。

在Go语言中，**函数类型、map类型自身，以及切片**只支持与nil的比较，而不支持同类型两个变量的比较，不能作为map的key类型的。

```go
s1 := make([]int, 1)
s2 := make([]int, 2)
f1 := func() {}
f2 := func() {}
m1 := make(map[int]string)
m2 := make(map[int]string)
println(s1 == s2) // 错误：invalid operation: s1 == s2 (slice can only be compared to nil)
println(f1 == f2) // 错误：invalid operation: f1 == f2 (func can only be compared to nil)
println(m1 == m2) // 错误：invalid operation: m1 == m2 (map can only be compared to nil)
```

### 16.2 map变量的声明和初始化

```go
var m map[string]int // 一个map[string]int类型的变量
```

有显式地赋予map变量初值，map类型变量的默认值为nil。

初值为零值nil的切片类型变量，可以借助内置的append的函数进行操作，这种在Go语言中被称为“==零值可用==”。定义“零值可用”的类型，可以提升开发者的使用体验，不用再担心变量的初始状态是否有效。

但map类型，因为它**内部实现的复杂性，无法“零值可用”**。所以，如果对处于零值状态的map变量直接进行操作，就会导致运行时异常（panic），从而导致程序进程异常退出：

```go
var m map[string]int // m = nil
m["key"] = 1         // 发生运行时异常：panic: assignment to entry in nil map
```

所以，**必须对map类型变量进行显式初始化后才能使用**。

为map类型变量显式赋值有两种方式:

#### 方法一：使用复合字面值初始化map类型变量

```go
m := map[int]string{}
```

此时m不是nil，对其进行键值对操作，不会引发运行时异常

```go
m1 := map[int][]string{
  1: []string{"val1_1", "val1_2"},
  3: []string{"val3_1", "val3_2", "val3_3"},
  7: []string{"val7_1"},
}
```

```go
type Position struct { 
  x float64 
  y float64
}

m2 := map[Position]string{
  Position{29.935523, 52.568915}: "school",
  Position{25.352594, 113.304361}: "shopping-mall",
  Position{73.224455, 111.804306}: "hospital",
}
```

“语法糖”，允许省略字面值中的元素类型，简写为：

```go
m2 := map[Position]string{
  {29.935523, 52.568915}: "school",
  {25.352594, 113.304361}: "shopping-mall",
  {73.224455, 111.804306}: "hospital",
}
```

#### 方法二：使用make为map类型变量进行显式初始化

```go
m1 := make(map[int]string) // 未指定初始容量
m2 := make(map[int]string, 8) // 指定初始容量为8
```

### 16.3 map的基本操作

#### 1️⃣插入新键值对

```go
m := make(map[int]string)
m[1] = "value1"
m[2] = "value2"
m[3] = "value3"


m := map[string]int {
    "key1" : 1,
    "key2" : 2,
}

m["key1"] = 11 // 11会覆盖掉"key1"对应的旧值1
m["key3"] = 3  // 此时m为map[key1:11 key2:2 key3:3]
```

#### 2️⃣获取键值对数量

```go
m := map[string]int {
    "key1" : 1,
    "key2" : 2,
}

fmt.Println(len(m)) // 2
m["key3"] = 3  
fmt.Println(len(m)) // 3 
```

> 不能对map类型变量调用cap，来获取当前容量

#### 3️⃣查找和数据读取

```go
m := make(map[string]int)
v := m["key1"]  // 如果这个键在map中并不存在，也会得到一个值：value元素类型的零值。
```

不能用上面的方法判断某个key是否在map中。

Go语言的map类型支持通过用一种名为“==comma ok==”的惯用法，进行对某个key的查询。

```go
m := make(map[string]int)
v, ok := m["key1"]
if !ok {
    // "key1"不在map中
}

// "key1"在map中，v将被赋予"key1"键对应的value
```

#### 4️⃣删除数据

```go
delete(m, "key2")
```

delete函数是从map中删除键的唯一方法。

#### 5️⃣遍历map中的键值数据

在Go中，遍历map的键值对只有一种方法，那就是像对待切片那样通过**for range**语句对map数据进行遍历。

```go
package main

import "fmt"

func main() {
    m := map[int]int{
        1: 11,
        2: 12,
        3: 13,
    }

    fmt.Printf("{ ")
    for k, v := range m {
        fmt.Printf("[%d, %d] ", k, v)
    }
    fmt.Printf("}\n")
} 
```

> **程序逻辑千万不要依赖遍历map所得到的的元素次序**。

### 16.4 map变量的传递开销

不用担心开销的问题。

和切片类型一样，map也是引用类型。这就意味着map类型变量作为参数被传递给函数或方法的时候，实质上传递的只是一个“描述符”，而不是整个map的数据拷贝，所以这个传递的开销是固定的，而且也很小。

并且，当map变量被传递到函数或方法内部后，我们在函数内部对map类型参数的修改在函数外部也是可见的。

```go
package main

import "fmt"

func foo(m map[string]int) {
    m["key1"] = 11
    m["key2"] = 12
}

func main() {
    m := map[string]int{
        "key1": 1,
        "key2": 2,
    }

    fmt.Println(m) // map[key1:1 key2:2]  
    foo(m)
    fmt.Println(m) // map[key1:11 key2:12] 
}
```

### 16.5 map的内部实现 ❤️ 

Go运行时使用一张哈希表来实现抽象的map类型。运行时实现了map类型操作的所有功能，包括查找、插入、删除等。在编译阶段，Go编译器会将Go语法层面的map操作，重写成运行时对应的函数调用。大致的对应关系是这样的：

```go
// 创建map类型变量实例
m := make(map[keyType]valType, capacityhint) → m := runtime.makemap(maptype, capacityhint, m)

// 插入新键值对或给键重新赋值
m["key"] = "value" → v := runtime.mapassign(maptype, m, "key") v是用于后续存储value的空间的地址

// 获取某键的值 
v := m["key"]      → v := runtime.mapaccess1(maptype, m, "key")
v, ok := m["key"]  → v, ok := runtime.mapaccess2(maptype, m, "key")

// 删除某键
delete(m, "key")   → runtime.mapdelete(maptype, m, “key”)
```

map类型在Go运行时层实现的示意图：

![](images/iShot_2025-03-26_12.16.45.png)



#### 1️⃣初始状态

与语法层面 map 类型变量（m）一一对应的是`*runtime.hmap` 的实例，即runtime.hmap类型的指针，也就是之前提到的**map 类型的描述符**。

hmap 类型是 map 类型的头部结构（header），它存储了后续 map 类型操作所需的所有信息，包括：

![](images/image-20250326122149813.png)

真正用来存储键值对数据的是桶，也就是bucket，每个bucket中存储的是Hash值低bit位数值相同的元素，默认的元素个数为 `BUCKETSIZE`（值为 8，Go 1.17版本中在`$GOROOT/src/cmd/compile/internal/reflectdata/reflect.go`中定义，与 `runtime/map.go` 中常量 bucketCnt 保持一致）。

> ```go
> // .../reflect.go
> const (
> 	BUCKETSIZE  = abi.MapBucketCount
> 	MAXKEYSIZE  = abi.MapMaxKeyBytes
> 	MAXELEMSIZE = abi.MapMaxElemBytes
> )
> 
> // runtime/map.go
> const (
> 	// Maximum number of key/elem pairs a bucket can hold.
> 	bucketCntBits = abi.MapBucketCountBits
> 	bucketCnt     = abi.MapBucketCount
>   ...
>   
> //  internal/abi/map.go
> const (
> 	MapBucketCountBits = 3 // log2 of number of elements in a bucket.
> 	MapBucketCount     = 1 << MapBucketCountBits
> 	MapMaxKeyBytes     = 128 // Must fit in a uint8.
> 	MapMaxElemBytes    = 128 // Must fit in a uint8.
> )
> ```



当某个bucket（比如buckets[0])的8个空槽slot）都填满了，且map尚未达到扩容的条件的情况下，运行时会建立overflow bucket，并将这个overflow bucket挂在上面bucket（如buckets[0]）末尾的overflow指针上，这样两个buckets形成了一个链表结构，直到下一次map扩容之前，这个结构都会一直存在。

每个bucket由三部分组成，从上到下分别是tophash区域、key存储区域和value存储区域。

##### tophash区域

当我们向map插入一条数据，或者是从map按key查询数据的时候，运行时都会使用哈希函数对key做哈希运算，并获得一个哈希值（hashcode）。这个hashcode非常关键，运行时会把hashcode“一分为二”来看待，其中低位区的值用于选定bucket，高位区的值用于在某个bucket中确定key的位置。示意图：

![](images/image-20250326123250461.png)

因此，每个bucket的tophash区域其实是用来快速定位key位置的，这样就避免了逐个key进行比较这种代价较大的操作。尤其是当key是size较大的字符串类型时，好处就更突出了。这是一种以空间换时间的思路。



##### key存储区域

tophash区域下面是一块连续的内存区域，存储的是这个bucket承载的所有key数据。运行时在分配bucket的时候需要知道key的Size。

运行时是如何知道key的size的呢？

当声明一个map类型变量，比如`var m map[string]int`时，Go运行时就会为这个变量对应的特定map类型，生成一个`runtime.maptype`实例。如果这个实例已经存在，就会直接复用。

```go
type maptype struct {
  typ        _type
  key        *_type
  elem       *_type
  bucket     *_type // internal type representing a hash bucket
  keysize    uint8  // size of key slot
  elemsize   uint8  // size of elem slot
  bucketsize uint16 // size of bucket
  flags      uint32
} 
```

这个实例包含了我们需要的map类型中的所有”元信息”。我们前面提到过，编译器会把语法层面的map操作重写成运行时对应的函数调用，这些运行时函数都有一个共同的特点，那就是第一个参数都是maptype指针类型的参数。🔖

**Go运行时就是利用maptype参数中的信息确定key的类型和大小的。**

##### value存储区域

key存储区域下方的另外一块连续的内存区域是key对应的value。和key一样，这个区域的创建也是得到了maptype中信息的帮助。

Go运行时采用了把key和value分开存储的方式，而不是采用一个kv接着一个kv的kv紧邻方式存储，这带来的其实是算法上的复杂性，但却减少了因内存对齐带来的内存浪费。

以map[int8]int64为例，看看下面的存储空间利用率对比图：

![](images/image-20250326124054630.png)

当前Go运行时使用的方案内存利用效率很高，而kv紧邻存储的方案在map[int8]int64这样的例子中内存浪费十分严重，它的内存利用率是72/128=56.25%，有近一半的空间都浪费掉了。

==注意==：如果key或value的数据长度大于一定数值，那么运行时不会在bucket中直接存储数据，而是会存储key或value数据的指针。目前Go运行时定义的最大key和value的长度是这样的：

```go
// $GOROOT/src/runtime/map.go
const (
    maxKeySize  = 128
    maxElemSize = 128
)
```



#### 2️⃣map扩容

map会对底层使用的内存进行自动管理。因此，在使用过程中，当插入元素个数超出一定数值后，map一定会存在自动扩容的问题，也就是怎么扩充bucket的数量，并重新在bucket间均衡分配数据的问题。

map在什么情况下会进行扩容呢？

Go运行时的map实现中引入了一个LoadFactor（负载因子），当**count > LoadFactor \* 2^B**或overflow bucket过多时，运行时会自动对map进行扩容。目前Go最新1.17版本LoadFactor设置为6.5（loadFactorNum/loadFactorDen）。

```go
// $GOROOT/src/runtime/map.go
const (
	... ...

	loadFactorNum = 13
	loadFactorDen = 2
	... ...
)

func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
	... ...
	if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
		hashGrow(t, h)
		goto again // Growing the table invalidates everything, so try again
	}
	... ...
}
```

这两方面原因导致的扩容，在运行时的操作其实是不一样的。如果是因为overflow bucket过多导致的“扩容”，实际上运行时会新建一个和现有规模一样的bucket数组，然后在assign和delete时做排空和迁移。

如果是因为当前数据数量超出LoadFactor指定水位而进行的扩容，那么运行时会建立一个**两倍于现有规模的bucket数组**，但真正的排空和迁移工作也是在assign和delete时逐步进行的。原bucket数组会挂在hmap的oldbuckets指针下面，直到原buckets数组中所有数据都迁移到新数组后，原buckets数组才会被释放。你可以结合下面的map扩容示意图来理解这个过程，这会让你理解得更深刻一些：



![](images/image-20240708190713060.png)

#### 3️⃣map与并发

从上面的实现原理来看，充当map描述符角色的hmap实例自身是有状态的（hmap.flags），而且对状态的读写是没有并发保护的。所以说map实例不是并发写安全的，也不支持并发读写。如果我们对map实例进行并发读写，程序运行时就会抛出异常。

🔖



### 总结

日常使用map的场合要把握住下面几个要点：

- 不要依赖map的元素遍历顺序；
- map不是线程安全的，不支持并发读写；
- 不要尝试获取map中元素（value）的地址。

### 思考题 🔖

> 对map类型进行遍历所得到的键的次序是随机的，实现一个方法，让能对map的进行稳定次序遍历？

一、核心实现方法

#### 1. **提取键并排序（适用于所有键类型）**
通过将 `map` 的键提取到切片中排序，再按顺序遍历，这是最通用的方法：
```go
package main
import (
    "fmt"
    "sort"
)

func main() {
    m := map[string]int{"banana": 2, "apple": 1, "cherry": 3}
    
    // 提取键到切片
    keys := make([]string, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    
    // 按字母升序排序
    sort.Strings(keys)
    
    // 按排序后的键遍历
    for _, k := range keys {
        fmt.Printf("%s: %d\n", k, m[k])
    }
}
```
**输出：**
```
apple: 1
banana: 2
cherry: 3
```
• **优点**：简单通用，支持任意可比较的键类型（如 `int`、`string` 等）。
• **缺点**：需额外内存存储键切片，排序可能带来性能开销（适用于小规模数据）。

#### 2. **自定义排序逻辑（复杂键类型）**
若键为结构体等复杂类型，需实现 `sort.Interface` 接口：
```go
type Person struct {
    Name string
    Age  int
}

func main() {
    m := map[Person]string{
        {"Alice", 30}: "Engineer",
        {"Bob", 25}:   "Designer",
    }
    
    // 提取键并自定义排序
    keys := make([]Person, 0, len(m))
    for k := range m {
        keys = append(keys, k)
    }
    
    sort.Slice(keys, func(i, j int) bool {
        return keys[i].Name < keys[j].Name // 按名字字母排序
    })
    
    // 遍历
    for _, k := range keys {
        fmt.Printf("%v: %s\n", k, m[k])
    }
}
```



进阶优化方案

#### 3. **预缓存排序后的键（频繁遍历场景）**
若需多次遍历同一 `map`，可预计算并缓存排序后的键切片，避免重复排序：
```go
var (
    m     = map[string]int{"banana": 2, "apple": 1}
    keys  []string
    once  sync.Once
)

func getSortedKeys() []string {
    once.Do(func() { // 保证仅初始化一次
        keys = make([]string, 0, len(m))
        for k := range m {
            keys = append(keys, k)
        }
        sort.Strings(keys)
    })
    return keys
}

// 使用时调用 getSortedKeys() 获取有序键列表
```



#### 4. **使用有序数据结构替代（高稳定性需求）**
若对顺序有极高要求，可改用以下数据结构代替原生 `map`：
• **有序映射库**：如 `github.com/iancoleman/orderedmap`，支持插入顺序遍历。
• **双向链表 + 哈希表**：自行实现类似 LRU 缓存的结构，维护键的顺序。

#### 注意事项
1. **键的动态变化**  
   若 `map` 的键频繁增删，需在每次修改后重新生成排序后的键列表，否则会导致顺序不一致。

2. **性能权衡**  
   • 对于大规模数据，排序可能成为性能瓶颈（时间复杂度为 O(n log n)）。
   • 推荐在初始化时排序，或使用预缓存策略减少开销。

3. **并发安全**  
   若在并发环境中操作 `map`，需通过 `sync.Mutex` 或 `sync.RWMutex` 保护键切片和 `map` 的读写。

---

#### 总结
| 方法                 | 适用场景                 | 性能影响 | 实现复杂度 |
| -------------------- | ------------------------ | -------- | ---------- |
| 提取键并排序         | 通用场景，小规模数据     | 中等     | 低         |
| 自定义排序逻辑       | 复杂键类型（如结构体）   | 中等     | 中         |
| 预缓存排序键         | 多次遍历同一 `map`       | 低       | 中         |
| 使用有序数据结构替代 | 高频遍历且顺序敏感的场景 | 低       | 高         |

通过上述方法，可有效实现 `map` 的稳定顺序遍历。若需完整代码示例或进一步优化思路，可参考 [Go 官方文档](https://pkg.go.dev/sort) 或第三方库如 `orderedmap`。



## 17 复合数据类型：用结构体建立对真实世界的抽象

> 编写程序的目的就是与真实世界交互，解决真实世界的问题，帮助真实世界提高运行效率与改善运行质量。

之前有基本数据类型和三个复合数据类型，还缺少一种**==通用的、对实体对象进行聚合抽象==**的能力。==结构体类型==（`struct`）

### 17.1 如何自定义一个新类型？

1. 第一种是==类型定义==（Type Definition）

```go
type T S // 定义一个新类型T
```

S可以是任何一个已定义的类型，包括Go原生类型或者自定义类型

```go
type T1 int 
type T2 T1  
```

如果一个新类型是基于某个Go原生类型定义的，那么这个Go原生类型为新类型的==底层类型（Underlying Type)==。上面的例子中，T1、T2的底层类型都是int。

> 为什么要提到底层类型这个概念呢？
> 
> 因为底层类型在Go语言中有重要作用，它被用来==判断两个类型本质上是否相同==（Identical）。

T1和T2的底层类型都是类型int，所以它们在本质上是相同的。而**本质上相同的两个类型，它们的变量可以通过显式转型进行相互赋值，相反，如果本质上是不同的两个类型，它们的变量间连显式转型都不可能，更不要说相互赋值了**。

除了基于已有类型定义新类型之外，还可以**==基于类型字面值==**来定义新类型，这种方式多用于自定义一个新的复合类型：

```go
type M map[int]string
type S []string
```

类型定义也支持通过**type块**的方式:

```go
type (
   T1 int
   T2 T1
   T3 string
)
```

2. 第二种是类型别名（Type Alias）。这种类型定义方式通常用在项目的**渐进式重构**，还有对已有包的**二次封装方面**。

```go
type T = S // type alias
```

类型别名并没有定义出新类型，T与S实际上就是**同一种类型**。

### 17.2 如何定义一个结构体类型？

复合类型的定义一般都是通过类型字面值的方式来进行的，作为复合类型之一的结构体类型也不例外:

```go
type T struct {
    Field1 T1
    Field2 T2
    ... ...
    FieldN Tn
}
```

还可以用**空标识符“`_`”作为结构体类型定义中的字段名称**。这样以空标识符为名称的字段，不能被外部包引用，甚至无法被结构体所在的包使用。

其他几种特殊情况：

#### 空结构体

```go
type Empty struct{} // Empty是一个不包含任何字段的空结构体类型
```

空结构体类型有什么用呢？

```go
var s Empty
println(unsafe.Sizeof(s)) // 0
```

空结构体类型变量的内存占用为0。基于**空结构体类型内存零开销**这样的特性，在日常Go开发中会经常使用空结构体类型元素，作为**一种“事件”信息**进行Goroutine之间的通信:

```go
var c = make(chan Empty) // 声明一个元素类型为Empty的channel
c<-Empty{}               // 向channel写入一个“事件”
```

这种以空结构体为元素类建立的channel，是目前能实现的、内存占用最小的Goroutine间通信方式。

#### 使用其他结构体作为自定义结构体中字段的类型

```go
type Person struct {
  Name string
  Phone string
  Addr string
}

type Book struct {
  Title string
  Author Person
  ... ...
}  
```

嵌入字段（Embedded Field） 

### 17.3 结构体变量的声明与初始化

#### 零值初始化

如果一种类型采用零值初始化得到的**零值变量**，是有意义的，而且是直接可用的，我称这种类型为**“零值可用”类型**。

在Go语言标准库和运行时的代码中，有很多践行“零值可用”理念的好例子，最典型的莫过于sync包的`Mutex`类型了。Mutex是Go标准库中提供的、用于**多个并发Goroutine之间进行同步的互斥锁**。

运用“零值可用”类型，给Go语言中的线程互斥锁带来了什么好处呢？横向对比一下C语言，要在C语言中使用线程互斥锁，通常需要：

```c
pthread_mutex_t mutex; 
pthread_mutex_init(&mutex, NULL);

pthread_mutex_lock(&mutex); 
... ...
pthread_mutex_unlock(&mutex); 
```

在C中使用互斥锁，需要首先声明一个mutex变量，还必须使用pthread_mutex_init函数对其进行专门的初始化操作后，它才能处于可用状态。

go语言：

```go
var mu sync.Mutex
mu.Lock()
mu.Unlock()  
```

另一个例子，标准库中的bytes.Buffer结构体类型：

```go
var b bytes.Buffer
b.Write([]byte("Hello, Go"))
fmt.Println(b.String()) // 输出：Hello, Go 
```

#### 使用复合字面值

最简单的对结构体变量进行显式初始化的方式，就是**按顺序依次给每个结构体字段进行赋值**。

```go
type Book struct {
    Title string              // 书名
    Pages int                 // 书的页数
    Indexes map[string]int    // 书的索引
}

var book = Book{"The Go Programming Language", 700, make(map[string]int)}
```

这样依次赋值的问题：

- 当结构体类型定义中的字段顺序发生变化，或者字段出现增删操作时，我们就需要手动调整该结构体类型变量的显式初始化代码，让赋值顺序与调整后的字段顺序一致。

- 当一个结构体的字段较多时，这种逐一字段赋值的方式实施起来就会比较困难

- 一旦结构体中包含非导出字段，那么这种逐一字段赋值的方式就不再被支持了，编译器会报错

> 事实上，Go语言并不推荐按字段顺序对一个结构体类型变量进行显式初始化，甚至Go官方还在提供的`go vet`工具中专门内置了一条检查规则：“`composites`”，用来静态检查代码中结构体变量初始化是否使用了这种方法，一旦发现，就会给出警告。

推荐使用**“`field:value`”形式的复合字面值**：

```go
var t = T{
  F2: "hello",
  F1: 11,
  F4: 14,
} 
```

#### 使用特定的构造函数

```go
// $GOROOT/src/time/sleep.go
type runtimeTimer struct {
    pp       uintptr
    when     int64
    period   int64
    f        func(interface{}, uintptr) 
    arg      interface{}
    seq      uintptr
    nextwhen int64
    status   uint32
}

type Timer struct {
    C <-chan Time
    r runtimeTimer
}

// Timer结构体专用的构造函数
func NewTimer(d Duration) *Timer {
    c := make(chan Time, 1)
    t := &Timer{
        C: c,
        r: runtimeTimer{
            when: when(d),
            f:    sendTime,
            arg:  c,
        },
    }
    startTimer(&t.r)
    return t
}
```

### 17.4 结构体类型的内存布局

Go结构体类型是既数组类型之后，第二个将它的元素（结构体字段）一个接着一个以“平铺”形式，存放在一个连**续内存块**中的。

![](images/image-20240704120037598.png)

在真实情况下，虽然Go编译器没有在结构体变量占用的内存空间中插入额外字段，但结构体字段实际上可能并不是紧密相连的，中间可能存在“缝隙”。这些“缝隙”同样是结构体变量占用的内存空间的一部分，它们是Go编译器插入的“==填充物（Padding）==”。

![](images/image-20240708193452381.png)

填充物是因为需要==内存对齐==，指的就是各种内存对象的内存地址不是随意确定的，必须满足特定要求。

对于各种**基本数据类型**来说，它的变量的内存地址值必须是**其类型本身大小的整数倍**，比如，一个int64类型的变量的内存地址，应该能被int64类型自身的大小，也就是8整除；一个uint16类型的变量的内存地址，应该能被uint16类型自身的大小，也就是2整除。

对于结构体而言，它的变量的内存地址，只要是**它最长字段长度与系统对齐系数两者之间较小的那个的整数倍**就可以了。但对于结构体类型来说，还要让它**每个字段的内存地址都严格满足内存对齐要求**。

```go
type T struct {
   b byte
   i int64
   u uint16
}
```

计算过程：

![](images/image-20240708194400125.png)

整个计算过程分为两个阶段。

第一个阶段是**对齐结构体的各个字段**。

- 首先，第一个字段b是长度1个字节的byte类型变量，这样字段b放在任意地址上都可以被1整除，所以它是天生对齐的。用一个sum来表示当前已经对齐的内存空间的大小，这个时候sum=1；
- 接下来，第二个字段i是一个长度为8个字节的int64类型变量。按照内存对齐要求，它应该被放在可以被8整除的地址上。但是，如果把i紧邻b进行分配，当i的地址可以被8整除时，b的地址就无法被8整除。这个时候，需要在b与i之间做一些填充，使得i的地址可以被8整除时，b的地址也始终可以被8整除，于是我们在i与b之间填充了7个字节，此时此刻sum=1+7+8；
- 再下来，第三个字段u是一个长度为2个字节的uint16类型变量，按照内存对其要求，它应该被放在可以被2整除的地址上。有了对其的i作为基础，现在知道将u与i相邻而放，是可以满足其地址的对齐要求的。i之后的那个字节的地址肯定可以被8整除，也一定可以被2整除。于是我们把u直接放在i的后面，中间不需要填充，此时此刻，sum=1+7+8+2。

现在结构体T的所有字段都已经对齐了，开始第二个阶段，也就是**对齐整个结构体**。

结构体的内存地址为min（结构体最长字段的长度，系统内存对齐系数）的整数倍，那么这里结构体T最长字段为i，它的长度为8，而64bit系统上的系统内存对齐系数一般为8，两者相同，我们取8就可以了。那么整个结构体的对齐系数就是8。

> 为什么上面的示意图还要在结构体的尾部填充了6个字节呢？

结构体T的对齐系数是8，那么就要保证每个结构体T的变量的内存地址，都能被8整除。如果只分配一个T类型变量，不再继续填充，也可能保证其内存地址为8的倍数。但如果考虑我们分配的是一个元素为T类型的数组，比如下面这行代码，我们虽然可以保证T[0]这个元素地址可以被8整除，但能保证T[1]的地址也可以被8整除吗？

```go
var array [10]T
```

数组是元素连续存储的一种类型，元素T[1]的地址为T[0]地址+T的大小(18)，显然无法被8整除，这将导致T[1]及后续元素的地址都无法对齐，这显然不能满足内存对齐的要求。加6变成24，就能被8整除了。

> 为什么会出现内存对齐的要求呢？

出于对处理器**存取数据效率**的考虑。在早期的一些处理器中，比如Sun公司的Sparc处理器仅支持内存对齐的地址，如果它遇到没有对齐的内存地址，会引发段错误，导致程序崩溃。我们常见的x86-64架构处理器虽然处理未对齐的内存地址不会出现段错误，但数据的存取性能也会受到影响。

从上可看出，Go语言中结构体类型的大小受内存对齐约束的影响。不**同的字段排列顺序也会影响到“填充字节”的多少，从而影响到整个结构体大小**。

```go
type T struct {
    b byte
    i int64
    u uint16
}

type S struct {
    b byte
    u uint16
    i int64
}

func main() {
    var t T
    println(unsafe.Sizeof(t)) // 24
    var s S
    println(unsafe.Sizeof(s)) // 16
}
```

所以，在日常定义结构体时，一定要注意**结构体中字段顺序**，尽量合理排序，降低结构体对内存空间的占用。

有些时候，为了保证某个字段的内存地址有更为严格的约束，也会做**主动填充**。比如runtime包中的mstats结构体定义就采用了主动填充：

```go
// $GOROOT/src/runtime/mstats.go
type mstats struct {
    ... ...
    // Add an uint32 for even number of size classes to align below fields
    // to 64 bits for atomic operations on 32 bit platforms.
    _ [1 - _NumSizeClasses%2]uint32 // 这里做了主动填充

    last_gc_nanotime uint64 // last gc (monotonic time)
    last_heap_inuse  uint64 // heap_inuse at mark termination of the previous GC
    ... ...
}
```

通过空标识符来进行主动填充。

### 思考题

> Go语言不支持在结构体类型定义中，递归地放入其自身类型字段，但却可以拥有自身类型的指针类型、以自身类型为元素类型的切片类型，以及以自身类型作为value类型的map类型的字段，你能思考一下其中的原因吗？



## 18 控制结构：if的“快乐路径”原则

Go中程序的分支结构：if和switch-case两种语句形式；循环结构：只有for。

![](images/image-20240709105248548.png)

#### if语句的自用变量

在if布尔表达式前声明的变量

这些变量只可以在if语句的代码块范围内使用

```go
func main() {
    if a, c := f(), h(); a > 0 {
        println(a)
    } else if b := f(); b > 0 {
        println(a, b)
    } else {
        println(a, b, c)
    }
}
```



> 如果一个if语句使用了多分支结构，那么if语句中的几个布尔表达式如何排列能达到最好的效果呢？



## 19 控制结构：Go的for循环，仅此一种

### 19.1 for语句的经典使用形式

![](images/image-20241225161246544.png)

1️⃣ **循环前置语句**。**仅会被执行一次**

2️⃣ **条件判断表达式**

3️⃣ **循环体**

4️⃣ **循环后置语句**

### 19.2 for range

Go为数组、切片、字符串、map、channel的for循环提供了语法糖：for range。

#### 切片

```go
for i, v := range sl {
    fmt.Printf("sl[%d] = %d\n", i, v)
}
```

这里的i和v对应的是经典for语句形式中循环前置语句的循环变量，它们的初值分别为切片sl的第一个元素的下标值和元素值。

三种变种：

```go
for i := range sl {
	// ... 
}

for _, v := range sl {
	// ... 
}

for _, _ = range sl {
	// ... 
}
```

#### string

```go
var s = "中国人"
for i, v := range s {
    fmt.Printf("%d %s 0x%x\n", i, string(v), v)
}
```

```
0 中 0x4e2d
3 国 0x56fd
6 人 0x4eba
```

for range对于string类型来说，**每次循环得到的v值是一个Unicode字符码点**，也就是rune类型值，而不是一个字节，返回的第一个值i为**该Unicode字符码点的内存编码（UTF-8）的第一个字节在字符串内存序列中的位置**。

> for与for range，对string类型进行循环操作的语义是不同的。



#### map

在Go语言中，**对map进行循环操作，for range是唯一的方法**。



#### channel

```go
var c = make(chan int)
for v := range c {
   // ... 
}
```

在这个例子中，for range每次从channel中读取一个元素后，会把它赋值给循环变量v，并进入循环体。当channel中没有数据可读的时候，for range循环会阻塞在对channel的读操作上。直到channel关闭时，for range循环才会结束，这也是for range循环与channel配合时隐含的循环判断条件。



#### 小结

| 数据类型        |      返回值      |  顺序性  | 底层传递方式 |
| :-------------- | :--------------: | :------: | :----------: |
| 数组            |     索引、值     | 固定顺序 |    值传递    |
| 切片            |     索引、值     | 固定顺序 |   引用传递   |
| 字符串          | 字节索引、`rune` |  按字符  |    不可变    |
| 映射（Map）     |      键、值      |   随机   |    哈希表    |
| 通道（Channel） |        值        | 先进先出 |     指针     |

### 19.3 带label的continue语句

label语句的作用，是标记跳转的目标。

```go
func main() {
    var sum int
    var sl = []int{1, 2, 3, 4, 5, 6}

loop:
    for i := 0; i < len(sl); i++ {
        if sl[i]%2 == 0 {
            // 忽略切片中值为偶数的元素
            continue loop
        }
        sum += sl[i]
    }
    println(sum) // 9
}

    
```





### 19.4 break语句的使用



### 19.5 for语句的常见“坑”与避坑方法 🔖

for range这个语法糖提高了Go的表达能力，同时也可能带入一些问题。

#### 1️⃣循环变量的重用



#### 2️⃣参与循环的是range表达式的副本



#### 3️⃣遍历map中元素的随机性



## 20 控制结构：Go中的switch语句有哪些变化？

switch语句是按照case出现的先后顺序对case表达式进行求值的，如果将匹配成功概率高的case表达式排在前面，就会有助于提升switch语句执行效率。这点对于case后面是表达式列表的语句同样有效，可以将匹配概率最高的表达式放在表达式列表的最左侧。

### switch语句的灵活性

**首先，switch语句各表达式的求值结果可以为各种类型值，只要它的类型支持==比较操作==就可以了。**



**第二点：switch语句支持声明临时变量。**



**第三点：case语句支持表达式列表。**



**第四点：取消了默认执行下一个case代码逻辑的语义。**

### type switch

```go
func main() {
    var x interface{} = 13
    switch x.(type) {
    case nil:
        println("x is nil")
    case int:
        println("the type of x is int")
    case string:
        println("the type of x is string")
    case bool:
        println("the type of x is string")
    default:
        println("don't support the type")
    }
}
```

switch关键字后面跟着的表达式为`x.(type)`，这种表达式形式是switch语句**专有的**，而且也只能在switch语句中使用。这个表达式中的**x必须是一个接口类型变量**，表达式的求值结果是这个接口类型变量对应的**动态类型**。

case关键字后面接的就不是普通意义上的表达式了，而是一个个**具体的类型**。

### 跳不出循环的break



## 21 函数：请叫我“一等公民”

函数是现代编程语言的基本语法元素，无论是在命令式语言、面向对象语言还是动态脚本语言中，函数都位列C位。

在Go语言中，**函数是唯一一种基于特定输入，实现特定任务并可返回任务执行结果的代码块**（Go语言中的方法本质上也是函数）。

如果忽略Go包在Go代码组织层面的作用，可以说**Go程序就是一组函数的集合**。

### 21.1 Go函数与函数声明

**Function**这个单词原本是**==功能、职责==**的意思。编程语言使用Function这个单词，表示将一个大问题分解后而形成的、若干具有特定功能或职责的小任务，可以说十分贴切。函数代表的小任务可以在一个程序中被多次使用，甚至可以在不同程序中被使用，因此**函数的出现也提升了整个程序界代码复用的水平**。

普通Go函数的声明：

![](images/image-20240704122156576.png)

==变长参数==，在类型前面增加了一个“`…`”符号。

**具名返回值**

把上面的函数声明等价转换为<u>变量声明的形式</u>：

![](images/image-20240704122531758.png)

**这不就是在声明一个类型为函数类型的变量吗**！

**函数声明中的函数名其实就是变量名**，函数声明中的**func关键字、参数列表和返回值列表**共同构成了**==函数类型==**。而参数列表与返回值列表的组合也被称为**==函数签名==**，它是决定两个函数类型是否相同的决定因素。因此，函数类型也可以看成是由func关键字与函数签名组合而成的。

通常，在表述函数类型时会省略函数签名参数列表中的参数名，以及返回值列表中的返回值变量名：

```go
func(io.Writer, string, ...interface{}) (int, error)
```

这样，如果两个函数类型的函数签名是相同的，即便参数列表中的参数名，以及返回值列表中的返回值变量名都是不同的，那么这两个函数类型也是相同类型。

结论：**每个函数声明所定义的函数，仅仅是对应的函数类型的一个实例**，就像var a int = 13这个变量声明语句中a是int类型的一个实例一样。

类似17章中，使用复合类型字面值对结构体类型变量进行显式初始化的内容

```go
s := T{}      // 使用复合类型字面值对结构体类型T的变量进行显式初始化
f := func(){} // 使用变量声明形式的函数声明
```

`T{}`被为==复合类型字面值==；`func(){}`就叫“==函数字面值（Function Literal）==”，它特别像一个没有函数名的函数声明，因此也叫它==匿名函数==。

#### 函数参数的那些事儿

由于函数分为**声明**与**使用**两个阶段，在不同阶段，参数的称谓也有不同。在函数声明阶段，把参数列表中的参数叫做**形式参数**（==Parameter==，简称==形参==），在函数体中，使用的都是形参；而在函数实际调用时传入的参数被称为**实际参数**（==Argument==，简称==实参==）。

![](images/image-20240704123547073.png)

当实际调用函数的时候，实参会传递给函数，并和形式参数逐一绑定，编译器会<u>根据各个形参的类型与数量，来检查传入的实参的类型与数量是否匹配</u>。只有匹配，程序才能继续执行函数调用，否则编译器就会报错。

Go语言中，函数参数传递采用是值传递的方式。所谓“**==值传递==**”，就是将实际参数在内存中的表示**==逐位拷贝==（Bitwise Copy）**到形式参数中。对于像**整型、数组、结构体**这类类型，它们的内存表示就是它们自身的数据内容，因此当这些类型作为实参类型时，值传递拷贝的就是它们自身，传递的开销也与它们自身的大小成正比。

但是像**string、切片、map**这些类型就不是了，它们的内存表示对应的是它们**数据内容的“==描述符==”**。当这些类型作为实参类型时，值传递拷贝的也是它们数据内容的“描述符”，不包括数据内容本身，所以这些类型传递的开销是**固定**的，与数据内容大小无关。这种只拷贝“描述符”，不拷贝实际数据内容的拷贝过程，也被称为“**==浅拷贝==**”。

不过函数参数的传递也有两个例外，当函数的形参为**接口类型**，或者形参是**变长参数**时，简单的值传递就不能满足要求了，这时Go编译器会介入：<u>对于类型为接口类型的形参，Go编译器会把传递的实参赋值给对应的接口类型形参；对于为变长参数的形参，Go编译器会将零个或多个实参按一定形式转换为对应的变长形参。</u>

那么这里，零个或多个传递给变长形式参数的实参，被Go编译器转换为何种形式了呢？

```go
func myAppend(sl []int, elems ...int) []int {
    fmt.Printf("%T\n", elems) // []int
    if len(elems) == 0 {
        println("no elems to append")
        return sl
    }

    sl = append(sl, elems...)
    return sl
}

func main() {
    sl := []int{1, 2, 3}
    sl = myAppend(sl) // no elems to append
    fmt.Println(sl) // [1 2 3]
    sl = myAppend(sl, 4, 5, 6)
    fmt.Println(sl) // [1 2 3 4 5 6]
}
```

在Go中，**变长参数实际上是通过切片来实现的**。所以，在函数体中，就可以**使用切片支持的所有操作来操作变长参数**，这会大大简化了变长参数的使用复杂度。比如myAppend中，我们使用len函数就可以获取到传给变长参数的实参个数。

#### 函数支持多返回值

```go
func foo()                       // 无返回值
func foo() error                 // 仅有一个返回值
func foo() (int, string, error)  // 有2或2个以上返回值
```

==具名返回值==（Named Return Value）变量可以像函数体中声明的局部变量一样在函数体内使用。

> 日常编码中，究竟该使用普通返回值形式，还是具名返回值形式呢？
> 
> **Go标准库以及大多数项目代码中的函数，都选择了使用普通的非具名返回值形式**。但在一些特定场景下，具名返回值也会得到应用。比如，<u>当函数使用defer，而且还在defer函数中修改外部函数返回值时</u>，具名返回值可以让代码显得更优雅清晰。
> 
> <u>当函数的返回值个数较多时</u>，每次显式使用return语句时都会接一长串返回值，这时，用具名返回值可以让函数实现的可读性更好一些:
> 
> ```go
> // $GOROOT/src/time/format.go
> func parseNanoseconds(value string, nbytes int) (ns int, rangeErrString string, err error) {
>     if !commaOrPeriod(value[0]) {
>         err = errBad
>         return
>     }
>     if ns, err = atoi(value[1:nbytes]); err != nil {
>         return
>     }
>     if ns < 0 || 1e9 <= ns {
>         rangeErrString = "fractional second"
>         return
>     }
> 
>     scaleDigits := 10 - nbytes
>     for i := 0; i < scaleDigits; i++ {
>         ns *= 10
>     }
>     return
> }  
> ```

### 21.2 函数是“一等公民”

> wiki发明人、C2站点作者[沃德·坎宁安(Ward Cunningham)](http://c2.com/)对“一等公民”的[解释](http://wiki.c2.com//?FirstClass)：
> 
> 如果一门编程语言对某种语言元素的**创建和使用没有限制**，我们可以**像对待值（value）一样对待**这种语法元素，那么我们就称这种语法元素是这门编程语言的“一等公民”。拥有“一等公民”待遇的语法元素可以存储在变量中，可以作为参数传递给函数，可以在函数内部创建并可以作为返回值从函数返回。

#### 特征一：Go函数可以存储在变量中

```go
var (
    myFprintf = func(w io.Writer, format string, a ...interface{}) (int, error) {
        return fmt.Fprintf(w, format, a...)
    }
)

func main() {
    fmt.Printf("%T\n", myFprintf) // func(io.Writer, string, ...interface {}) (int, error)
    myFprintf(os.Stdout, "%s\n", "Hello, Go") // 输出Hello，Go
}
```

#### 特征二：支持在函数内创建并通过返回值返回。 🔖

```go
func setup(task string) func() {
    println("do some setup stuff for", task)
    return func() {
        println("do some teardown stuff for", task)
    }
}

func main() {
    teardown := setup("demo")
    defer teardown()
    println("do some bussiness stuff")
}
```

这个例子，模拟了执行一些重要逻辑之前的上下文建立（==setup==），以及之后的上下文拆除（==teardown==）。在一些单元测试的代码中，我们也经常会在执行某些用例之前，建立此次执行的上下文（setup），并在这些用例执行后拆除上下文（teardown），避免这次执行对后续用例执行的干扰。

在这个例子中，在setup函数中创建了这次执行的上下文拆除函数，并通过返回值的形式，将这个拆除函数返回给了setup函数的调用者。setup函数的调用者，在执行完对应这次执行上下文的重要逻辑后，再调用setup函数返回的拆除函数，就可以完成对上下文的拆除了。

setup函数中创建的拆除函数也是一个匿名函数，但和前面看到的匿名函数有一个不同，这个不同就在于这个匿名函数**使用了定义它的函数setup的局部变量task**，这样的匿名函数在Go中也被称为==闭包（Closure）==。

闭包本质上就是一个**匿名函数或叫函数字面值**，它们可以引用它的包**裹函数**，也就是创建它们的函数中定义的**变**量。然后，这些变量在包裹函数和匿名函数之间**共享**，只要闭包可以被访问，这些共享的变量就会继续存在。

#### 特征三：作为参数传入函数。

```go
time.AfterFunc(time.Second*2, func() { println("timer fired") })
```

#### 特征四：拥有自己的类型。

每个函数都和整型值、字符串值等一等公民一样，拥有自己的类型，也就是==函数类型==。

也可以基于函数类型来自定义类型，就像基于整型、字符串类型等类型来自定义类型一样。

```go
// $GOROOT/src/net/http/server.go
type HandlerFunc func(ResponseWriter, *Request)

// $GOROOT/src/sort/genzfunc.go
type visitFunc func(ast.Node) ast.Visitor
```

### 21.3 函数“一等公民”特性的高效运用 🔖

#### 应用一：函数类型的妙用

函数也可以被显式转型。

```go
func greeting(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, Gopher!\n")
}                    

func main() {
    http.ListenAndServe(":8080", http.HandlerFunc(greeting))
}
```

```go
// $GOROOT/src/net/http/server.go
func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}
```

```go
// $GOROOT/src/net/http/server.go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```



#### 应用二：利用闭包简化函数调用。



### 思考题

> 还能列举出其他的高效运用函数“一等公民”特性的例子吗？



## 22 函数：怎么结合多返回值进行错误处理？

多返回值是Go语言函数，区别于其他主流静态编程语言中函数的一个重要特点。

### 22.1 Go语言是如何进行错误处理的？

C语言的错误处理机制，通常用一个类型为**整型的函数返回值作为错误状态标识**，函数调用者会基于值比较的方式，对这一代表错误状态的返回值进行检视。通常，这个返回值为0，就代表函数调用成功；如果这个返回值是其它值，那就代表函数调用出现错误。也就是说，**函数调用者需要根据这个返回值代表的错误状态，来决定后续执行哪条错误处理路径上的代码**。

C语言的这种简单的、基于错误值比较的错误处理机制有什么优点呢？

- 首先，它让每个开发人员必须**显式地去关注和处理每个错误**，经过显式错误处理的代码会更健壮，也会让开发人员对这些代码更有信心。
- 另外，这些错误就是**普通的值**，所以不需要用额外的语言机制去处理它们，只需利用已有的语言机制，像处理其他普通类型值一样的去处理错误就可以了，这也让代码更容易调试，更容易针对每个错误处理的决策分支进行测试覆盖。【Go继承了这种简单与显示结合的特征】

C这种错误处理机制的缺点：

- C的函数最多仅支持一个返回值，很多开发者会把这单一的返回值“一值多用”。一个返回值，不仅要承载函数要返回给**调用者的信息**，又要承载函数调用的**最终错误状态**。
  
  比如C标准库中的fprintf函数，在正常情况下，它的返回值表示输出到FILE流中的字符数量，但如果出现错误，这个返回值就变成了一个负数，代表具体的错误值。
  
  ```c
  // stdio.h
  int fprintf(FILE * restrict stream, const char * restrict format, ...);
  ```
  
  特别是当返回值为其他类型，比如字符串的时候，还很难将它与错误状态融合到一起。这个时候，很多C开发人员要么使用输出参数，承载要返回给调用者的信息，要么自定义一个包含返回信息与错误状态的结构体，作为返回值类型。大家做法不一，就很难形成统一的错误处理策略。

为了避免这种情况，Go函数增加了多返回值机制。

```go
// fmt包
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
```

Go语言惯用法，是使用error这个接口类型表示错误，并且按惯例，通常将error类型返回值放**在返回值列表的==末尾==**。

### 22.2 error类型与错误值构造

> `error`接口类型究竟如何表示错误？

```go
// $GOROOT/src/builtin/builtin.go
type interface error {
    Error() string
}
```

任何实现了error的Error方法的类型的实例，都可以作为错误值赋值给error接口变量。

两种构造错误值的方法： `errors.New`和`fmt.Errorf`:

```go
err := errors.New("your first demo error")
errWithCtx = fmt.Errorf("index %d is out of bounds", i)
```

都返回的是同一个实现了error接口的类型的实例，未导出的类型就是`errors.errorString`:

```go
// $GOROOT/src/errors/errors.go
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```

但在一些场景下，错误处理者需要从错误值中提取出**更多信息**，帮助他选择错误处理路径，显然这两种方法就不能满足了。这个时候，可以自定义错误类型来满足这一需求。比如：标准库中的net包就定义了一种携带额外错误上下文的错误类型：

```go
// $GOROOT/src/net/net.go
type OpError struct {
    Op string
    Net string
    Source Addr
    Addr Addr
    Err error
}
```

这样，错误处理者就可以根据这个类型的错误值提供的**额外上下文信息**，比如Op、Net、Source等，做出错误处理路径的选择，比如下面标准库中的代码：

```go
// $GOROOT/src/net/http/server.go
func isCommonNetReadError(err error) bool {
    if err == io.EOF {
        return true
    }
    if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
        return true
    }
    if oe, ok := err.(*net.OpError); ok && oe.Op == "read" {
        return true
    }
    return false
}
```

这段代码利用==类型断言（Type Assertion）==，判断error类型变量err的**动态类型**是否为`*net.OpError`或 `net.Error`。如果err的动态类型是 `*net.OpError`，那么类型断言就会返回这个动态类型的值（存储在oe中），代码就可以通过判断它的Op字段是否为"read"来判断它是否为CommonNetRead类型的错误。

使用error类型，而不是传统意义上的整型或其他类型作为错误类型，三点好处：

- 第一点：统一了错误类型。
- 第二点：错误是值。 可以像整型值那样对错误做“==”和“!=”的逻辑比较。
- 第三点：易扩展，支持自定义错误上下文。

error接口是错误值的提供者与错误值的检视者之间的**契约**。error接口的实现者负责提供**错误上下文**，供负责错误处理的代码使用。这种**错误具体上下文与作为错误值类型的error接口类型的解耦**，也体现了Go组合设计哲学中“==正交==”的理念。

Go语言的几种错误处理的惯用策略，学习这些策略将有助于我们提升函数错误处理设计的能力。

### 22.3 策略一：透明错误处理策略

简单来说，Go语言中的错误处理，就是**根据函数/方法返回的error类型变量中携带的错误值信息做决策，并选择后续代码执行路径的过程**。

最简单的错误策略莫过于**完全不关心返回错误值携带的具体上下文信息**，只要发生错误就进入唯一的错误处理执行路径：

```go
err := doSomething()
if err != nil {
    // 不关心err变量底层错误值所携带的具体上下文信息
    // 执行简单错误处理逻辑并返回
    ... ...
    return err
}
```

这也是Go语言中**最常见的错误处理策略**，80%以上的Go错误处理情形都可以归类到这种策略下。

此种情况，错误值的构造方（如上面的函数doSomething）可以直接使用Go标准库提供的两个基本错误值构造方法errors.New和fmt.Errorf来构造错误值：

```go
func doSomething(...) error {
    ... ...
    return errors.New("some error occurred")
}
```

这样构造出的错误值代表的上下文信息，对错误处理方是透明的，因此这种策略称为“**透明错误处理策略**”。

### 22.4 策略二：“哨兵”错误处理策略 🔖

当错误处理方不能只根据“透明的错误值”就做出错误处理路径选取的情况下，错误处理方会尝试对返回的错误值进行检视，于是就有可能出现下面代码中的**反模式**：

```go
data, err := b.Peek(1)
if err != nil {
    switch err.Error() {
    case "bufio: negative count":
        // ... ...
        return
    case "bufio: buffer full":
        // ... ...
        return
    case "bufio: invalid use of UnreadByte":
        // ... ...
        return
    default:
        // ... ...
        return
    }
}
```

简单来说，反模式就是，**错误处理方以透明错误值所能提供的唯一上下文信息（描述错误的字符串），作为错误处理路径选择的依据**。

但这种“反模式”会造成严重的**隐式耦合**。这也就意味着，错误值构造方不经意间的一次错误描述字符串的改动，都会造成错误处理方处理行为的变化，并且这种通过字符串比较的方式，对错误值进行检视的性能也很差。

那这有什么办法吗？

Go标准库采用了定义导出的（Exported）“哨兵”错误值的方式，来辅助错误处理方检视（inspect）错误值并做出错误处理分支的决策，比如下面的bufio包中定义的“哨兵错误”：

```go
// $GOROOT/src/bufio/bufio.go
var (
    ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
    ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
    ErrBufferFull        = errors.New("bufio: buffer full")
    ErrNegativeCount     = errors.New("bufio: negative count")
)
```

```go
data, err := b.Peek(1)
if err != nil {
    switch err {
    case bufio.ErrNegativeCount:
        // ... ...
        return
    case bufio.ErrBufferFull:
        // ... ...
        return
    case bufio.ErrInvalidUnreadByte:
        // ... ...
        return
    default:
        // ... ...
        return
    }
}
```

一般“哨兵”错误值变量以ErrXXX格式命名。和透明错误策略相比，“哨兵”策略让错误处理方在有检视错误值的需求时候，可以“有的放矢”。

不过，对于API的开发者而言，暴露“哨兵”错误值也意味着这些**错误值和包的公共函数/方法一起成为了API的一部分**。一旦发布出去，开发者就要对它进行很好的维护。而“哨兵”错误值也让使用这些值的错误处理方对它产生了依赖。

从Go 1.13版本开始，标准库errors包提供了`Is`函数用于错误处理方对错误值的检视。Is函数类似于把一个error类型变量与“哨兵”错误值进行比较:

```go
// 类似 if err == ErrOutOfBounds{ … }
if errors.Is(err, ErrOutOfBounds) {
    // 越界的错误处理
} 
```

不同的是，如果error类型变量的底层错误值是一个**包装错误（Wrapped Error）**，errors.Is方法会沿着该包装错误所在**错误链（Error Chain)**，与链上所有被包装的错误（Wrapped Error）进行比较，直至找到一个匹配的错误为止。

```go
var ErrSentinel = errors.New("the underlying sentinel error")

func main() {
    err1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
    err2 := fmt.Errorf("wrap err1: %w", err1)
  println(err2 == ErrSentinel) //false
    if errors.Is(err2, ErrSentinel) {
        println("err2 is ErrSentinel")
        return
    }

    println("err2 is not ErrSentinel")
}
```

这里通过fmt.Errorf函数，并且使用%w创建包装错误变量err1和err2，其中err1实现了对ErrSentinel这个“哨兵错误值”的包装，而err2又对err1进行了包装，这样就形成了一条错误链。位于错误链最上层的是err2，位于最底层的是ErrSentinel。之后，再分别通过值比较和errors.Is这两种方法，判断err2与ErrSentinel的关系。

### 22.5 策略三：错误值类型检视策略

如果遇到错误处理方需要错误值提供更多的“错误上下文”的情况，那么就通过**自定义错误类型的构造错误值**的方式来实现。

由于错误值都通过error接口变量统一呈现，要得到底层错误类型携带的错误上下文信息，错误处理方需要使用Go提供的==类型断言机制（Type Assertion）==或==类型选择机制（Type Switch）==，这种错误处理方式，可称之为**错误值类型检视策略**。

```go
// $GOROOT/src/encoding/json/decode.go
type UnmarshalTypeError struct {
    Value  string       
    Type   reflect.Type 
    Offset int64        
    Struct string       
    Field  string      
}

func (d *decodeState) addErrorContext(err error) error {
    if d.errorContext.Struct != nil || len(d.errorContext.FieldStack) > 0 {
        switch err := err.(type) {
        case *UnmarshalTypeError:
            err.Struct = d.errorContext.Struct.Name()
            err.Field = strings.Join(d.errorContext.FieldStack, ".")
            return err
        }
    }
    return err
} 
```

这段代码通过类型switch语句得到了err变量代表的动态类型和值，然后在匹配的case分支中利用错误上下文信息进行处理。

从Go 1.13版本开始，标准库errors包提供了`As`函数给错误处理方检视错误值。As函数类似于通过类型断言判断一个error类型变量是否为特定的自定义错误类型:

```go
// 类似 if e, ok := err.(*MyError); ok { … }
var e *MyError
if errors.As(err, &e) {
    // 如果err类型为*MyError，变量e将被设置为对应的错误值
}
```

不同的是，如果error类型变量的动态错误值是一个包装错误，errors.As函数会沿着该包装错误所在错误链，与链上所有被包装的错误的类型进行比较，直至找到一个匹配的错误类型，就像errors.Is函数那样。

```go
type MyError struct {
    e string
}

func (e *MyError) Error() string {
    return e.e
}

func main() {
    var err = &MyError{"MyError error demo"}
    err1 := fmt.Errorf("wrap err: %w", err)
    err2 := fmt.Errorf("wrap err1: %w", err1)
    var e *MyError
    if errors.As(err2, &e) {
        println("MyError is on the chain of err2")
        println(e == err)                  
        return                             
    }                                      
    println("MyError is not on the chain of err2")
} 
```

### 22.6 策略四：错误行为特征检视策略

除了第一种策略有效降低了错误的构造方与错误处理方两者之间的耦合。其它两种策略依然在错误的构造方与错误处理方两者之间建立了耦合。

错误行为特征检视策略：**将某个包中的错误类型归类，统一提取出一些公共的错误行为特征，并将这些错误行为特征放入一个公开的接口类型中**。

以标准库中的net包为例，它将包内的所有错误类型的公共行为特征抽象并放入net.Error这个接口中：

```go
// $GOROOT/src/net/net.go
type Error interface {
    error
    Timeout() bool  
    Temporary() bool
}
```

net.Error接口包含两个用于判断错误行为特征的方法：Timeout用来判断是否是超时（Timeout）错误，Temporary用于判断是否是临时（Temporary）错误。

而错误处理方只需要依赖这个公共接口，就可以检视具体错误值的错误行为特征信息，并根据这些信息做出后续错误处理分支选择的决策。

```go
// $GOROOT/src/net/http/server.go
func (srv *Server) Serve(l net.Listener) error {
    ... ...
    for {
        rw, e := l.Accept()
        if e != nil {
            select {
            case <-srv.getDoneChan():
                return ErrServerClosed
            default:
            }
            if ne, ok := e.(net.Error); ok && ne.Temporary() {
                // 注：这里对临时性(temporary)错误进行处理
                ... ...
                time.Sleep(tempDelay)
                continue
            }
            return e
        }
        ...
    }
    ... ...
}
```

Accept方法实际上返回的错误类型为*OpError，它是net包中的一个自定义错误类型，它实现了错误公共特征接口net.Error:

```go
// $GOROOT/src/net/net.go
type OpError struct {
    ... ...
    // Err is the error that occurred during the operation.
    Err error
}

type temporary interface {
    Temporary() bool
}

func (e *OpError) Temporary() bool {
  if ne, ok := e.Err.(*os.SyscallError); ok {
      t, ok := ne.Err.(temporary)
      return ok && t.Temporary()
  }
  t, ok := e.Err.(temporary)
  return ok && t.Temporary()
}
```

错误处理策略选择上的建议：

- 请尽量使用“透明错误”处理策略，降低错误处理方与错误值构造方之间的耦合；
- 如果可以通过错误值类型的特征进行错误检视，那么请尽量使用“错误行为特征检视策略”;
- 在上述两种策略无法实施的情况下，再使用“哨兵”策略和“错误值类型检视”策略；
- Go 1.13及后续版本中，尽量用errors.Is和errors.As函数替换原先的错误检视比较语句。



## 23 函数：怎么让函数更简洁健壮？

健壮的函数意味着，**无论调用者如何使用你的函数，它都能以合理的方式处理调用者的任何输入，并给调用者返回预设的、清晰的错误值**。即便你的函数发生内部异常，函数也会尽力从异常中恢复，尽可能地不让异常蔓延到整个程序。

而简洁优雅，则意味着函数的实现易读、易理解、更易维护，同时简洁也意味着统计意义上的更少的bug。

### 23.1 健壮性的“三不要”原则

- 原则一：**不要相信任何外部输入的参数。**
- 原则二：**不要忽略任何一个错误。**

对应函数实现中，对标准库或第三方包提供的函数或方法调用，不能假定它一定会成功，**一定要显式地检查这些调用返回的错误值**。一旦发现错误，要及时终止函数执行，防止错误继续传播。

- 原则三：**不要假定异常不会发生。**

**异常不是错误**。错误是可预期的，也是经常会发生的，我们有对应的公开错误码和错误处理预案，但异常却是少见的、意料之外的。通常意义上的异常，指的是硬件异常、操作系统异常、语言运行时异常，还有更大可能是代码中潜在bug导致的异常，比如代码中出现了以0作为分母，或者是数组越界访问等情况。

### 23.2 认识Go语言中的异常：panic

在Go中，panic主要有两类来源，一类是来自**Go运行时**，另一类则是**Go开发人员通过panic函数主动触发的**。

无论是哪种，一旦panic被触发，后续Go程序的执行过程都是一样的，这个过程被Go语言称为**==panicking==**。

当函数F调用panic函数时，函数F的执行将停止。不过，函数F中已进行求值的deferred函数都会得到正常执行，执行完这些deferred函数后，函数F才会把控制权返还给其调用者。

对于函数F的调用者而言，函数F之后的行为就如同调用者调用的函数是panic一样，该panicking过程将继续在栈上进行下去，直到当前Goroutine中的所有函数都返回为止，然后Go程序将崩溃退出。

```go
func foo() {
    println("call foo")
    bar()
    println("exit foo")
}

func bar() {
    println("call bar")
    panic("panic occurs in bar")  // 调用panic函数手动触发了panic
    zoo()
    println("exit bar")
}

func zoo() {
    println("call zoo")
    println("exit zoo")
}

func main() {
    println("call main")
    foo()
    println("exit main")
}
```

函数的调用次序依次为main -> foo ->bar -> zoo。在bar函数中，调用panic函数手动触发了panic。结果为：

```
call main
call foo
call bar
panic: panic occurs in bar
```

panicking过程:

- 程序从入口函数main开始依次调用了foo、bar函数，在bar函数中，代码在调用zoo函数之前调用了panic函数触发了异常。那示例的panicking过程就从这开始了。bar函数调用panic函数之后，它自身的执行就此停止了，所以我们也没有看到代码继续进入zoo函数执行。并且，bar函数没有捕捉这个panic，这样这个panic就会沿着函数调用栈向上走，来到了bar函数的调用者foo函数中。
- 从foo函数的视角来看，**这就好比将它对bar函数的调用，换成了对panic函数的调用一样**。这样一来，foo函数的执行也被停止了。由于foo函数也没有捕捉panic，于是panic继续沿着函数调用栈向上走，来到了foo函数的调用者main函数中。
- 同理，从main函数的视角来看，这就好比将它对foo函数的调用，换成了对panic函数的调用一样。结果就是，main函数的执行也被终止了，于是整个程序异常退出，日志"exit main"也没有得到输出的机会。

可以通过`recover`函数来实现捕捉panic并恢复程序正常执行秩序:

```go
func bar() {
    defer func() {
        if e := recover(); e != nil {
            fmt.Println("recover the panic:", e)
        }
    }()

    println("call bar")
    panic("panic occurs in bar")
    zoo()
    println("exit bar")
}
```

### 23.3 如何应对panic？

#### 第一点：评估程序对panic的忍受度

**不同应用对异常引起的程序崩溃退出的忍受度是不一样的。**

比如，一个单次运行于控制台窗口中的命令行交互类程序（CLI），和一个常驻内存的后端HTTP服务器程序，对异常崩溃的忍受度就是不同的。

前者即便因异常崩溃，对用户来说也仅仅是再重新运行一次而已。但后者一旦崩溃，就很可能导致整个网站停止服务。

所以，**针对各种应用对panic忍受度的差异，采取的应对panic的策略也应该有不同**。

像后端HTTP服务器程序这样的任务关键系统，就需要在特定位置捕捉并恢复panic，以保证服务器整体的健壮度。在这方面，Go标准库中的http server就是一个典型的代表。

Go标准库提供的http server采用的是，每个客户端连接都使用一个单独的Goroutine进行处理的并发处理模型。也就是说，客户端一旦与http server连接成功，http server就会为这个连接新创建一个Goroutine，并在这Goroutine中执行对应连接（conn）的serve方法，来处理这条连接上的客户端请求。

**无论在哪个Goroutine中发生未被恢复的panic，整个程序都将崩溃退出**。所以，为了保证处理某一个客户端连接的Goroutine出现panic时，不影响到http server主Goroutine的运行，Go标准库在serve方法中加入了对panic的捕捉与恢复，下面是serve方法的部分代码片段：

```go
// $GOROOT/src/net/http/server.go
// Serve a new connection.
func (c *conn) serve(ctx context.Context) {
    c.remoteAddr = c.rwc.RemoteAddr().String()
    ctx = context.WithValue(ctx, LocalAddrContextKey, c.rwc.LocalAddr())
    defer func() {
        if err := recover(); err != nil && err != ErrAbortHandler {
            const size = 64 << 10
            buf := make([]byte, size)
            buf = buf[:runtime.Stack(buf, false)]
            c.server.logf("http: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
        }
        if !c.hijacked() {
            c.close()
            c.setState(c.rwc, StateClosed, runHooks)
        }
    }()
    ... ...
}
```

serve方法在一开始处就设置了defer函数，并在该函数中捕捉并恢复了可能出现的panic。这样，即便处理某个客户端连接的Goroutine出现panic，处理其他连接Goroutine以及http server自身都不会受到影响。

这种**局部不要影响整体**的异常处理策略，在很多并发程序中都有应用。并且，捕捉和恢复panic的位置通常都在子Goroutine的起始处，这样设置可以捕捉到后面代码中可能出现的所有panic，就像serve方法中那样。

#### 第二点：提示潜在bug

有了对panic忍受度的评估，panic是不是也没有那么“恐怖”了呢？而且，甚至可以借助panic来帮助快速找到潜在bug。

C语言中有个很好用的辅助函数，断言（assert宏）。在使用C编写代码时，我们经常在一些代码执行路径上，使用断言来表达这段执行路径上某种条件一定为真的信心。断言为真，则程序处于正确运行状态，断言为否就是出现了意料之外的问题，而这个问题很可能就是一个潜在的bug，这时我们可以借助断言信息快速定位到问题所在。

不过，Go语言标准库中并没有提供断言之类的辅助函数，但我们可以使用panic，部分模拟断言对潜在bug的提示功能。比如，下面就是标准库`encoding/json`包使用panic指示潜在bug的一个例子：

```go
// $GOROOT/src/encoding/json/decode.go
... ...
//当一些本不该发生的事情导致我们结束处理时，phasePanicMsg将被用作panic消息
//它可以指示JSON解码器中的bug，或者
//在解码器执行时还有其他代码正在修改数据切片。

const phasePanicMsg = "JSON decoder out of sync - data changing underfoot?"

func (d *decodeState) init(data []byte) *decodeState {
    d.data = data
    d.off = 0
    d.savedError = nil
    if d.errorContext != nil {
        d.errorContext.Struct = nil
        // Reuse the allocated space for the FieldStack slice.
        d.errorContext.FieldStack = d.errorContext.FieldStack[:0]
    }
    return d
}

func (d *decodeState) valueQuoted() interface{} {
    switch d.opcode {
    default:
        panic(phasePanicMsg)

    case scanBeginArray, scanBeginObject:
        d.skip()
        d.scanNext()

    case scanBeginLiteral:
        v := d.literalInterface()
        switch v.(type) {
        case nil, string:
            return v
        }
    }
    return unquotedValue{}
}
```

在`valueQuoted`这个方法中，如果程序执行流进入了default分支，那这个方法就会引发panic，这个panic会提示开发人员：这里很可能是一个bug。

同样，在json包的encode.go中也有使用panic做潜在bug提示的例子：

```go
// $GOROOT/src/encoding/json/encode.go

func (w *reflectWithString) resolve() error {
    ... ...
    switch w.k.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        w.ks = strconv.FormatInt(w.k.Int(), 10)
        return nil
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        w.ks = strconv.FormatUint(w.k.Uint(), 10)
        return nil
    }
    panic("unexpected map key type")
}
```

`resolve`方法的最后一行代码就相当于一个“代码逻辑不会走到这里”的断言。一旦触发“断言”，这很可能就是一个潜在bug。

去掉这行代码并不会对`resolve`方法的逻辑造成任何影响，但真正出现问题时，开发人员就缺少了“断言”潜在bug提醒的辅助支持了。

在Go标准库中，**大多数panic的使用都是充当类似==断言==的作用的**。

#### 第三点：不要混淆异常与错误 

Java的“checked exception”处理的本质是**错误处理**，对应go中基于error值比较模型的错误处理。

而Go中的panic呢，更接近于Java的`RuntimeException`+`Error`。



### 23.4 使用defer简化函数实现

```go
func doSomething() error {
    var mu sync.Mutex
    mu.Lock()

    r1, err := OpenResource1()
    if err != nil {
        mu.Unlock()
        return err
    }

    r2, err := OpenResource2()
    if err != nil {
        r1.Close()
        mu.Unlock()
        return err
    }

    r3, err := OpenResource3()
    if err != nil {
        r2.Close()
        r1.Close()
        mu.Unlock()
        return err
    }

    // 使用r1，r2, r3
    err = doWithResources() 
    if err != nil {
        r3.Close()
        r2.Close()
        r1.Close()
        mu.Unlock()
        return err
    }

    r3.Close()
    r2.Close()
    r1.Close()
    mu.Unlock()
    return nil
}
```

类代码的特点就是在函数中会申请一些资源，并在函数退出前释放或关闭这些资源，比如这里的互斥锁mu以及资源r1~r3就是这样。

函数的实现需要确保，无论函数的执行流是按预期顺利进行，还是出现错误，这些资源在函数退出时都要被及时、正确地释放。为此，我们需要尤为关注函数中的错误处理，在错误处理时不能遗漏对资源的释放。

但这样的要求，就导致我们在进行资源释放，尤其是有多个资源需要释放的时候，比如上面示例那样，会大大增加开发人员的心智负担。同时当待释放的资源个数较多时，整个代码逻辑就会变得十分复杂，程序可读性、健壮性也会随之下降。但即便如此，如果函数实现中的某段代码逻辑抛出panic，传统的错误处理机制依然没有办法捕获它并尝试从panic恢复。

Go语言引入defer的初衷，就是解决这些问题。defer具体的运作机制是怎样的呢？

defer是Go语言提供的一种延迟调用机制，defer的运作离不开函数。

- 在Go中，只有在函数（和方法）内部才能使用defer；
- defer关键字后面只能接函数（或方法），这些函数被称为**deferred函数**。defer将它们注册到其所在Goroutine中，用于存放deferred函数的栈数据结构中，这些deferred函数将在执行defer的函数退出前，按后进先出（LIFO）的顺序被程序调度执行。

![](images/image-20240710100648693.png)

而且，无论是执行到函数体尾部返回，还是在某个错误处理分支显式return，又或是出现panic，已经存储到deferred函数栈中的函数，都会被调度执行。所以说，<u>deferred函数是一个可以在任何情况下为函数进行**收尾工作**的好“伙伴”</u>。

修改为：

```go
func doSomething() error {
    var mu sync.Mutex
    mu.Lock()
    defer mu.Unlock()

    r1, err := OpenResource1()
    if err != nil {
        return err
    }
    defer r1.Close()

    r2, err := OpenResource2()
    if err != nil {
        return err
    }
    defer r2.Close()

    r3, err := OpenResource3()
    if err != nil {
        return err
    }
    defer r3.Close()

    // 使用r1，r2, r3
    return doWithResources() 
}
```

资源释放函数的defer注册动作，紧邻着资源申请成功的动作，这样成对出现的惯例就极大降低了遗漏资源释放的可能性。

同时，代码的简化也意味代码可读性的提高，以及代码健壮度的增强。

### 23.5 defer使用的几个注意事项

defer不仅可以用来**捕捉和恢复panic**，还能让函数变得**更简洁和健壮**。

#### 第一点：明确哪些函数可以作为deferred函数

可以是自定义的函数或方法，但对于有返回值的自定义函数或方法，返回值会在deferred函数被调度执行的时候被自动丢弃。

> Go语言内置函数的完全列表：
>
> ```plain
> Functions:
> 	append cap close complex copy delete imag len
> 	make new panic print println real recover
> ```

，**内置函数能否作为 `deferred` 函数取决于其返回值类型和语法特性**。

- **可直接使用**的内置函数：`close`、`copy`、`delete`、`print`、`panic`、`recover` 等无返回值的函数。

- 不能直接使用的，**需包裹匿名函数**的内置函数：`append`、`len`、`cap`、`make`、`new` 等有返回值的函数。

  ```go
  // 示例：间接调用 append
  defer func() { _ = append(slice, 1) }()  // 合法，但返回值无实际意义[4,9](@ref)
  
  // 示例：间接调用 len
  defer func() { fmt.Println(len(slice)) }()  // 合法
  ```



#### 第二点：注意defer关键字后面表达式的求值时机

牢记：**defer关键字后面的表达式，是在将deferred函数注册到deferred函数栈的时候进行求值的**。



#### 第三点：知晓defer带来的性能损耗

Go核心团队对defer性能进行了多次优化，到Go 1.17版本之后，defer的开销已经足够小了。



### 思考题

> 除了捕捉panic、延迟释放资源外，日常编码中还有哪些使用defer的小技巧呢？

#### 1. **性能分析与调试**
利用 `defer` 记录函数执行时间，辅助优化代码性能。  **示例**：  
```go
func ExpensiveTask() {
    start := time.Now()
    defer func() {
        log.Printf("耗时: %v", time.Since(start))
    }()
    // 执行耗时操作...
}
```
• **原理**：`defer` 在函数退出时记录时间差，适用于定位性能瓶颈。
• **适用场景**：算法耗时分析、接口响应时间监控。

#### 2. **锁的延迟释放**
在并发场景中，确保锁的释放与加锁逻辑对应，避免死锁。  **标准库示例**（`sync.Mutex`）：  
```go
var mu sync.Mutex
func SafeWrite(data map[string]int, key string) {
    mu.Lock()
    defer mu.Unlock() // 确保解锁执行
    data[key] = 1
}
```
• **优势**：即使函数中途 `return` 或发生 `panic`，锁仍会被释放。
• **扩展**：读写锁（`sync.RWMutex`）的 `RLock()`/`RUnlock()` 同样适用。

#### 3. **修改函数返回值**
通过 `defer` 修改命名返回值，实现后置逻辑调整。  **示例**：  
```go
func ProcessData(input string) (result string, err error) {
    defer func() {
        if err == nil {
            result = "处理结果：" + result // 修改返回值
        }
    }()
    // 正常处理逻辑...
    return rawResult, nil
}
```
• **关键点**：仅适用于命名返回值（Named Return Values）。
• **陷阱**：若返回值被 `defer` 修改，需注意调用方是否依赖原始值。

#### 4. **上下文清理与状态重置**
管理全局状态或临时环境，例如重置日志输出格式。  **示例**：  
```go
func LogWithFormat(format string) {
    original := log.Flags()
    defer log.SetFlags(original) // 恢复原始日志格式
    log.SetFlags(log.Ldate | log.Lmicroseconds)
    // 使用临时格式记录日志...
}
```
• **应用场景**：临时修改环境变量、数据库连接参数等。

#### 5. **协程同步控制**
结合 `sync.WaitGroup` 确保协程任务完成。  **示例**：  
```go
func ProcessBatch(jobs []Job) {
    var wg sync.WaitGroup
    for _, job := range jobs {
        wg.Add(1)
        go func(j Job) {
            defer wg.Done() // 确保协程结束前调用 Done()
            // 处理任务...
        }(job)
    }
    wg.Wait()
}
```
• **作用**：防止因 `panic` 或提前 `return` 导致 `WaitGroup` 计数器未归零。

#### 6. **事务回滚与提交**
在数据库操作中，通过 `defer` 统一处理事务状态。  **伪代码示例**：  
```go
func UpdateOrder(db *sql.DB) (err error) {
    tx, err := db.Begin()
    if err != nil {
        return err
    }
    defer func() {
        if err != nil { // 根据错误决定回滚或提交
            tx.Rollback()
        } else {
            tx.Commit()
        }
    }()
    // 执行多个 SQL 操作...
    return nil
}
```
• **设计意义**：集中事务状态判断，避免遗漏 `Rollback()`。

#### 7. **HTTP 响应处理**
在 HTTP 服务中统一关闭响应体或处理错误。  **标准库示例**（`net/http`）：  
```go
resp, err := http.Get("https://example.com")
if err != nil {
    return err
}
defer resp.Body.Close() // 确保响应体关闭
body, err := io.ReadAll(resp.Body)
```
• **扩展技巧**：结合 `defer` 和 `recover()` 统一捕获 HTTP 处理函数的 `panic`。

#### 注意事项与陷阱
1. **循环中的 `defer`**  
   避免在循环内直接使用 `defer`（如文件关闭），可能导致资源释放延迟。应包裹为函数：  
   
   ```go
   for _, file := range files {
       func(f *os.File) {
           defer f.Close()
           // 处理文件...
       }(file)
   }
   ```
   
2. **参数立即求值**  
   `defer` 的参数在注册时求值，而非执行时。例如：  
   ```go
   x := 1
   defer fmt.Println(x) // 输出 1
   x = 2
   ```

3. **性能优化**  
   Go 1.14+ 的开放编码优化（Open-coded Defer）可减少 `defer` 的开销，但需满足条件（如函数内 `defer` 数量不超过 8 个）。



## 24 方法：理解“方法”的本质

函数是Go代码中的基本功能逻辑单元，它承载了**Go程序的所有执行逻**辑。可以说，Go程序的执行流本质上就是**在函数调用栈中上下流动，从一个函数到另一个函数**。

### 24.1 认识Go方法

Go引入方法这一元素，并不是要支持面向对象编程范式，而是Go践行==组合设计哲学==的一种实现层面的需要。

![](images/image-20240704132750133.png)

**这个`receiver`参数也是==方法与类型之间的纽带==，也是方法与函数的最大不同。**

Go中的方法必须是归属于一个类型的，而receiver参数的类型就是这个方法归属的类型，或者说这个方法就是这个类型的一个方法。

方法的一般声明形式：

```go
func (t *T或T) MethodName(参数列表) (返回值列表) {
    // 方法体
}
```

无论receiver参数的类型为`*T`还是T，都把一般声明形式中的T叫做receiver参数t的==基类型==。

如果t的类型为T，那么说这个方法是类型`T`的一个方法；如果t的类型为`*T`，那么就说这个方法是类型`*T`的一个方法。

每个方法只能有一个receiver参数。

> 方法接收器（receiver）参数、函数/方法参数，以及返回值变量对应的作用域范围，都是函数/方法体对应的显式代码块。

这就意味着，receiver部分的参数名不能与方法参数列表中的形参名，以及具名返回值中的变量名存在冲突，必须在这个方法的作用域中具有唯一性。

如果在方法体中，没有用到receiver参数，也可以省略receiver的参数名。

**receiver参数的<u>基类型本身不能为指针类型或接口类型</u>。**

```go
type MyInt *int
func (r MyInt) String() string { // r的基类型为MyInt，编译器报错：invalid receiver type MyInt (MyInt is a pointer type)
    return fmt.Sprintf("%d", *(*int)(r))
}

type MyReader io.Reader
func (r MyReader) Read(p []byte) (int, error) { // r的基类型为MyReader，编译器报错：invalid receiver type MyReader (MyReader is an interface type)
    return r.Read(p)
}
```

Go要求，**方法声明(的位置)要与receiver参数的基类型声明放在同一个包内**。基于这个约束可以得到两个推论：

1. 不能为原生类型（诸如int、float64、map等）添加方法。

```go
func (i int) Foo() string { // 编译器报错：cannot define new methods on non-local type int
    return fmt.Sprintf("%d", i) 
}
```

2. 不能跨越Go包为其他包的类型声明新方法。

```go
import "net/http"

func (s http.Server) Foo() { // 编译器报错：cannot define new methods on non-local type http.Server
} 
```



```go
type T struct{}

func (t T) M(n int) {
}

func main() {
	var t T
	t.M(1) // 通过类型T的变量实例调用方法M

	p := &T{}
	p.M(2) // 通过类型*T的变量实例调用方法M
}
```

### 24.2 方法的本质是什么？

```go
type T struct { 
    a int
}

func (t T) Get() int {  
    return t.a 
}

func (t *T) Set(a int) int { 
    t.a = a 
    return t.a 
}
```

> C++中的对象在调用方法时，编译器会自动传入指向对象自身的this指针作为方法的第一个参数。
>
> Go方法中的原理也是相似的，只不过是将receiver参数以第一个参数的身份并入到方法的参数列表中。上面的方法可以等价转换为普通函数：
>
> ```go
> // 类型T的方法Get的等价函数
> func Get(t T) int {  
>     return t.a 
> }
> 
> // 类型*T的方法Set的等价函数
> func Set(t *T, a int) int { 
>     t.a = a 
>     return t.a 
> }
> ```
>
> 这种等价转换是由Go编译器在编译和生成代码时自动完成的。



==方法表达式（Method Expression）== 🔖



Go语言中的方法的本质就是，**一个以方法的receiver参数作为第一个参数的普通函数**。

```go
func main() {
    var t T
    f1 := (*T).Set // f1的类型，也是*T类型Set方法的类型：func (t *T, int)int
    f2 := T.Get    // f2的类型，也是T类型Get方法的类型：func(t T)int
    fmt.Printf("the type of f1 is %T\n", f1) // the type of f1 is func(*main.T, int) int
    fmt.Printf("the type of f2 is %T\n", f2) // the type of f2 is func(main.T) int
    f1(&t, 3)
    fmt.Println(f2(t)) // 3
}    
```





### 巧解难题🔖



## 25 方法：方法集合与如何选择receiver类型？

由于在Go语言中，**方法本质上就是函数**，所以之前关于函数设计的内容对方法也同样适用，比如错误处理设计、针对异常的处理策略、使用defer提升简洁性，等等。

### 25.1 receiver参数类型对Go方法的影响

```go
func (t T) M1() <=> F1(t T)
func (t *T) M2() <=> F2(t *T)
```

- 首先，当receiver参数的类型为T时：值拷贝传递
- 第二，当receiver参数的类型为*T时：传递的是T类型实例的地址

### 25.2 选择receiver参数类型的第一个原则

**如果Go方法要把对receiver参数代表的类型实例的修改，反映到原类型实例上，那么应该选择*T作为receiver参数的类型**。

### 25.3 选择receiver参数类型的第二个原则

如果receiver参数类型的size较大，以值拷贝形式传入就会导致较大的性能开销，这时选择*T作为receiver类型可能更好些。

### 25.4 方法集合

==方法集合（Method Set）==

```go
      type Interface interface {
    M1()
    M2()
}

type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

func main() {
    var t T
    var pt *T
    var i Interface

    i = pt
    i = t // cannot use t (type T) as type Interface in assignment: T does not implement Interface (M2 method has pointer receiver)
}
```





**方法集合也是用来判断一个类型是否实现了某接口类型的唯一手段**，可以说，“**方法集合决定了接口实现**”。

Go中任何一个类型都有属于自己的方法集合:

- 以int、*int为代表的Go原生类型由于没有定义方法，所以它们的方法集合都是空的，称其**拥有空方法集合**。

- 接口类型相对特殊，它只会列出代表接口的方法列表，不会具体定义某个方法，它的方法集合就是它的**方法列表中的所有方法**。

```go
package main

import (
	"fmt"
	"reflect"
)

type T struct{}

func (T) M1() {}
func (T) M2() {}

func (*T) M3() {}
func (*T) M4() {}

// 查看一个非接口类型的方法集合
func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)

	if dynTyp == nil {
		fmt.Println("there is no dynamic type\n")
		return
	}

	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}

	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}

func main() {
	var n int
	dumpMethodSet(n)
	dumpMethodSet(&n)

	var t T
	dumpMethodSet(t)
	dumpMethodSet(&t)
}
```

结果：

```
int's method set is empty!
*int's method set is empty!
main.T's method set:
- M1
- M2

*main.T's method set:
- M1
- M2
- M3
- M4
```

- Go语言规定，`*T`类型的方法集合包含所有以`*T`为receiver参数类型的方法，以及所有以T为receiver参数类型的方法。

所谓**的方法集合决定接口实现**的含义就是：如果某类型T的方法集合与某接口类型的方法集合相同，或者类型T的方法集合是接口类型I方法集合的超集，那么我们就说这个类型T实现了接口I。

### 25.5 选择receiver参数类型的第三个原则🔖

第三个原则的选择依据就是**T类型是否需要实现某个接口**，也就是是否存在将T类型的变量赋值给某接口类型变量的情况。

如果**T类型需要实现某个接口**，那我们就要使用T作为receiver参数的类型，来满足接口类型方法集合中的所有方法。

如果T不需要实现某一接口，但`*T`需要实现该接口，那么根据方法集合概念，`*T`的方法集合是包含T的方法集合的，这样我们在确定Go方法的receiver的类型时，参考原则一和原则二就可以了。

如果说前面的两个原则更多聚焦于类型内部，从单个方法的实现层面考虑，那么这第三个原则则是更多从全局的设计层面考虑，聚焦于这个类型与接口类型间的耦合关系。



## 26 方法：如何用类型嵌入模拟实现“继承”？

==独立的自定义类型==就是这个类型的所有方法都是自己显式实现的。

Go语言不支持经典面向对象的编程范式与语法元素，所以我们这里只是借用了“继承”这个词汇而已，说是“继承”，实则依旧是一种**组合**的思想。

而这种“继承”，我们是通过Go语言的类型嵌入（Type Embedding）来实现的。

### 26.1 什么是类型嵌入

==类型嵌入（Type Embedding）==指的就是在一个类型的定义中嵌入了其他类型。Go语言支持两种类型嵌入，分别是==接口类型的类型嵌入==和==结构体类型的类型嵌入==。

#### 1️⃣接口类型的类型嵌入

```go
type E interface {
  M1()
  M2()
}
```

```go
type I interface {
  M1()
  M2()
  M3()
}
```

```go
type I interface {
  E
  M3()
}
```

接口类型嵌入的语义就是新接口类型（如接口类型I）将嵌入的接口类型（如接口类型E）的方法集合，并入到自己的方法集合中。

Go中的接口类型中只包含少量方法，并且常常只是一个方法。通过在接口类型中嵌入其他接口类型可以实现接口的组合，这也是**Go语言中基于已有接口类型构建新接口类型的惯用法。**

Go标准库有很多这种组合方式，比如，io包的ReadWriter、ReadWriteCloser等接口类型就是通过嵌入Reader、Writer或Closer三个基本的接口类型组合而成的：

```go
// $GOROOT/src/io/io.go

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}


type ReadWriter interface {
    Reader
    Writer
}

type ReadCloser interface {
    Reader
    Closer
}

type WriteCloser interface {
    Writer
    Closer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}    
```



#### 2️⃣结构体类型的类型嵌入

```go
type T1 int
type t2 struct{
    n int
    m int
}

type I interface {
    M1()
}

type S1 struct {
    T1
    *t2
    I            
    a int
    b string
}
```

结构体S1定义中有三个“非常规形式”的标识符，分别是`T1`、`t2`和`I`，它们**既代表字段的名字，也代表字段的类型**。

- 标识符T1表示字段名为T1，它的类型为自定义类型T1；
- 标识符t2表示字段名为t2，它的类型为自定义结构体类型t2的指针类型；
- 标识符I表示字段名为I，它的类型为接口类型I。

这种以某个==类型名==、==类型的指针类型名==或==接口类型名==，直接作为结构体字段的方式就叫做结**==构体的类型嵌入==**，这些字段也被叫做**==嵌入字段==**（Embedded Field）。

```go
type MyInt int

func (n *MyInt) Add(m int) {
    *n = *n + MyInt(m)
}

type t struct {
    a int
    b int
}

type S struct {
    *MyInt
    t
    io.Reader
    s string
    n int
}

func main() {
    m := MyInt(17)
    r := strings.NewReader("hello, go")
    s := S{
        MyInt: &m,
        t: t{
            a: 1,
            b: 2,
        },
        Reader: r,
        s:      "demo",
    }

    var sl = make([]byte, len("hello, go"))
    s.Reader.Read(sl)
    fmt.Println(string(sl)) // hello, go
    s.MyInt.Add(5)
    fmt.Println(*(s.MyInt)) // 22
}   
```

> 嵌入字段的可见性与嵌入字段的类型的可见性是一致的。如果嵌入类型的名字是首字母大写的，那么也就说明这个嵌入字段是可导出的。

> Go语言规定如果结构体使用从其他包导入的类型作为嵌入字段，比如pkg.T，那么这个嵌入字段的字段名就是T，代表的类型为pkg.T。所以上面的S类型中对应字段名称是`Reader`，代表的类型是`io.Reader`。

嵌入字段的用法与普通字段不同点：

- 和Go方法的receiver的基类型一样，嵌入字段类型的底层类型不能为指针类型。
- 嵌入字段的名字在结构体定义也必须是唯一的。



#### 3️⃣“实现继承”的原理

部分变化：

```go
var sl = make([]byte, len("hello, go"))
s.Read(sl) 
fmt.Println(string(sl))
s.Add(5) 
fmt.Println(*(s.MyInt))
```

像是：**Read方法与Add方法就是类型S方法集合中的方法**。

其实，这两个方法就来自结构体类型S的两个嵌入字段Reader和MyInt。**结构体类型S“继承”了Reader字段的方法Read的实现，也“继承”了*MyInt的Add方法的实现**。

- 当我们通过结构体类型S的变量s调用Read方法时，Go发现结构体类型S自身并没有定义Read方法，于是Go会查看S的嵌入字段对应的类型是否定义了Read方法。这个时候，Reader字段就被找了出来，之后s.Read的调用就被转换为s.Reader.Read调用。
- 这样一来，嵌入字段Reader的Read方法就被提升为S的方法，放入了类型S的方法集合。同理`*MyInt`的Add方法也被提升为S的方法而放入S的方法集合。从外部来看，这种嵌入字段的方法的提升就给了我们一种结构体类型S“继承”了io.Reader类型Read方法的实现，以及`*MyInt`类型Add方法的实现的错觉。

类型嵌入这种看似“继承”的机制，实际上是一种组合的思想。更具体点，它是一种组合中的**代理（delegate）模式**，如下图所示：

![](images/image-20241231115101611.png)

S只是一个代理（delegate），对外它提供了它可以代理的所有方法，如例子中的Read和Add方法。当外界发起对S的Read方法的调用后，S将该调用委派给它内部的Reader实例来实际执行Read方法。



### 26.2 类型嵌入与方法集合

#### 结构体类型中嵌入接口类型

```go
type I interface {
    M1()
    M2()
}

type T struct {
    I
}

func (T) M3() {}

func main() {
    var t T
    var p *T
    dumpMethodSet(t)
    dumpMethodSet(p)
}
```

```go
main.T's method set:
- M1
- M2
- M3

*main.T's method set:
- M1
- M2
- M3
```

**结构体类型的方法集合，包含嵌入的接口类型的方法集合。**

🔖

#### 结构体类型中嵌入结构体类型

```go
type T1 struct{}

func (T1) T1M1()   { println("T1's M1") }
func (*T1) PT1M2() { println("PT1's M2") }

type T2 struct{}

func (T2) T2M1()   { println("T2's M1") }
func (*T2) PT2M2() { println("PT2's M2") }

type T struct {
    T1
    *T2
}

func main() {
    t := T{
        T1: T1{},
        T2: &T2{},
    }

    dumpMethodSet(t)
    dumpMethodSet(&t)
}
```

输出结果：

```
main.T's method set:
- PT2M2
- T1M1
- T2M1

*main.T's method set:
- PT1M2
- PT2M2
- T1M1
- T2M1
```

- 类型T的方法集合 = T1的方法集合 + `*T2`的方法集合
- 类型`*T`的方法集合 = `*T1`的方法集合 + `*T2`的方法集合





### 26.3 defined类型与alias类型的方法集合

Go语言中，凡通过类型声明语法声明的类型都被称为==defined类型==。

```go
type I interface {
  M1()
  M2()
}
type T int
type NT T // 基于已存在的类型T创建新的defined类型NT
type NI I // 基于已存在的接口类型I创建新defined接口类型NI
```

🔖



> 无论原类型是接口类型还是非接口类型，类型别名都与原类型拥有完全相同的方法集合。

## 27 即学即练：跟踪函数调用链，理解代码更直观

### 27.1 引子

使用defer可以跟踪函数的执行过程。

```go
package main

func main() {
    defer Trace("main")()
    foo()
}

func Trace(name string) func() {
    println("enter: ", name)
    return func() {
        println("exit: ", name)
    }
}

func foo() {
    defer Trace("foo")()
    bar()
}

func bar() {
    defer Trace("bar")()
}
```

函数调用跟踪：

```
enter:  main
enter:  foo
enter:  bar
exit:  bar
exit:  foo
exit:  main
```

程序按main -> foo -> bar的函数调用次序执行，代码在函数的入口与出口处分别输出了跟踪日志。

不足之处：

- 调用Trace时需手动显式传入要跟踪的函数名；
- 如果是并发应用，不同Goroutine中函数链跟踪混在一起无法分辨；
- 输出的跟踪结果缺少层次感，调用关系不易识别；
- 对要跟踪的函数，需手动调用Trace函数。

> 目标：**实现一个自动注入跟踪代码，并输出有层次感的函数调用链跟踪命令行工具**。

### 27.2 自动获取所跟踪函数的函数名



### 27.3 增加Goroutine标识 🔖



### 27.4 让输出的跟踪信息更具层次感 🔖

对于程序员来说，缩进是最能体现出“层次感”的方法。



### 27.5 利用代码生成自动注入Trace函数 🔖

#### 将Trace函数放入一个独立的module中

#### 自动注入Trace函数

#### 利用instrument工具注入跟踪代码





# 核心篇：“脑勤+”洞彻核心

核心篇主要涵盖**接口类型语法与Go原生提供的三个并发原语（Goroutine、channel与select）**，之所以将它们放在核心语法的位置，是因为它们不仅代表了Go语言在**编程语言领域的创新**，更是影响Go**==应用骨架==（Application Skeleton）**设计的重要元素。

所谓应用骨架，就是指将应用代码中的**业务逻辑、算法实现逻辑、错误处理逻辑**等“皮肉”逐一揭去后所呈现出的应用结构。

![](images/image-20240711195856322.png)

从静态角度去看，我们能清晰地看到应用程序的组成部分以及各个部分之间的连接；从动态角度去看，我们能看到这幅骨架上可独立运动的几大机构。

前者我们可以将其理解为Go应用内部的耦合设计，而后者我们可以理解为应用的并发设计。

## 28 接口：接口即契约 ❤️

### 28.1 认识接口类型

**==接口类型==是由type和interface关键字定义的一组方法集合**，其中，方法集合唯一确定了这个接口类型所表示的接口。

```go
type MyInterface interface {
    M1(int) error
    M2(io.Writer, ...string)
}
```

接口类型MyInterface所表示的接口的方法集合，包含两个方法M1和M2。之所以称M1和M2为“方法”，更多是**从这个接口的实现者的角度考虑的**。但从上面接口类型声明中各个“方法”的形式上来看，这更像是**不带有func关键字的函数名+函数签名（参数列表+返回值列表）**的组合。

并且，接口类型的方法集合中声明的方法，它的**参数列表和返回值列表都不需要写出形参名字**。也就是说，**方法的参数列表中形参名字与返回值列表中的具名返回值，都不作为区分两个方法的凭据**。比如，下面两个等价：

```go
type MyInterface interface {
    M1(a int) error
    M2(w io.Writer, strs ...string)
}

type MyInterface interface {
    M1(n int) error
    M2(w io.Writer, args ...string)
} 
```

不过，Go语言要求接口类型声明中的**方法必须是具名的**，并且方法名字在这个接口类型的方法集合中是**唯一**的。

Go接口类型允许嵌入的不同接口类型的方法集合存在**交集**，但前提是交集中的方法不仅名字要一样，它的函数签名部分也要保持一致，也就是参数列表与返回值列表也要相同，否则Go编译器照样会报错。

```go
type Interface1 interface {
    M1()
}
type Interface2 interface {
    M1(string) 
    M2()
}

type Interface3 interface{
    Interface1
    Interface2 // 编译器报错：duplicate method M1
    M3()
}
```

在Go接口类型的方法集合中放入首字母小写的**非导出方法**也是合法的。

```go
// $GOROOT/src/context.go

// A canceler is a context type that can be canceled directly. The
// implementations are *cancelCtx and *timerCtx.
type canceler interface {
    cancel(removeFromParent bool, err error)
    Done() <-chan struct{}
}
```

如果接口类型的方法集合中包含非导出方法，那么这个接口类型自身通常也是非导出的，它的应用范围也仅局限于包内。【很少会用这种带有非导出方法的接口类型】

==空接口类型== `interface{}`

接口类型一旦被定义后，它就和其他Go类型一样可以用于声明变量:

```go
var err error   // err是一个error接口类型的实例变量
var r io.Reader // r是一个io.Reader接口类型的实例变量 
```

这些类型为接口类型的变量被称为**==接口类型变量==**，如果没有被显式赋予初值，接口类型变量的默认值为nil。如果要为接口类型变量显式赋予初值，我们就要为接口类型变量选择合法的右值。

Go规定：**如果一个类型T的方法集合是某接口类型I的方法集合的等价集合或超集，我们就说类型T实现了接口类型I，那么类型T的变量就可以作为合法的右值赋值给接口类型I的变量。**

可以将任何类型的值作为右值，赋值给空接口类型的变量。

```go
var i interface{} = 15 // ok
i = "hello, golang" // ok
type T struct{}
var t T
i = t  // ok
i = &t // ok
```

空接口类型的这一可接受任意类型变量值作为右值的特性，让他成为Go加入泛型语法之前**唯一一种具有“泛型”能力的语法元素**。

Go语言还支持接口类型变量赋值的“**逆操作**”，也就是通过接口类型变量“还原”它的右值的类型与值信息，这个过程被称为“**==类型断言==（Type Assertion）**”：

```go
v, ok := i.(T) 
```

其中i是某一个接口类型变量，如果T是一个非接口类型且T是想要还原的类型，那么这句代码的含义就是**断言存储在接口类型变量i中的值的类型为T**。

如果接口类型变量i之前被赋予的值确为T类型的值，那么这个语句执行后，左侧“comma, ok”语句中的变量ok的值将为true，变量v的类型为T，它值会是之前变量i的右值。如果i之前被赋予的值不是T类型的值，那么这个语句执行后，变量ok的值为false，**变量v的类型还是那个要还原的类型，但它的值是类型T的零值**。

```go
var a int64 = 13
var i interface{} = a
v1, ok := i.(int64) 
fmt.Printf("v1=%d, the type of v1 is %T, ok=%t\n", v1, v1, ok) // v1=13, the type of v1 is int64, ok=true
v2, ok := i.(string)
fmt.Printf("v2=%s, the type of v2 is %T, ok=%t\n", v2, v2, ok) // v2=, the type of v2 is string, ok=false
v3 := i.(int64) 
fmt.Printf("v3=%d, the type of v3 is %T\n", v3, v3) // v3=13, the type of v3 is int64
v4 := i.([]int) // panic: interface conversion: interface {} is int64, not []int
fmt.Printf("the type of v4 is %T\n", v4)
```

如果`v, ok := i.(T)`中的T是一个接口类型，那么类型断言的语义就会变成：**断言i的值实现了接口类型T**。如果断言成功，变量v的类型为i的值的类型，而并非接口类型T。如果断言失败，v的类型信息为接口类型T，它的值为nil。

```go
type MyInterface interface {
    M1()
}

type T int

func (T) M1() {
    println("T's M1")
}              

func main() {  
    var t T    
    var i interface{} = t
    v1, ok := i.(MyInterface)
    if !ok {   
        panic("the value of i is not MyInterface")
    }          
    v1.M1()    
    fmt.Printf("the type of v1 is %T\n", v1) // the type of v1 is main.T

    i = int64(13)
    v2, ok := i.(MyInterface)
    fmt.Printf("the type of v2 is %T\n", v2) // the type of v2 is <nil>
    // v2 = 13 //  cannot use 1 (type int) as type MyInterface in assignment: int does not implement MyInterface (missing M1   method) 
}
```



### 28.2 尽量定义“小接口”

- 隐式契约，无需签署，自动生效

Go语言中接口类型与它的实现者之间的关系是**隐式**的，不需要像其他语言（比如Java）那样要求实现者显式放置“implements”进行修饰，实现者只需要实现接口方法集合中的全部方法便算是遵守了契约，并立即生效了。

- 更倾向于“小契约”

Go选择了使用“小契约”，表现在代码上就是尽量定义小接口，即**方法个数在1~3个之间的接口**。

早期版本的Go标准库（Go 1.13版本）、Docker项目（Docker 19.03版本）以及Kubernetes项目（Kubernetes 1.17版本）中定义的接口类型方法集合中方法数量：

![](images/image-20240711203002095.png)

### 28.3 小接口有哪些优势？

#### 第一点：接口越小，抽象程度越高

计算机程序本身就是对真实世界的==抽象与再建构==。抽象就是**对同类事物去除它具体的、次要的方面，抽取它相同的、主要的方面**。不同的抽象程度，会导致抽象出的概念对应的事物的集合不同。**抽象程度越高，对应的集合空间就越大；抽象程度越低，也就是越具像化，更接近事物真实面貌，对应的集合空间越小。**

对生活中不同抽象程度的形象诠释：

![](images/image-20240711203233701.png)

这张图中我们分别建立了三个抽象：

- 会飞的。这个抽象对应的事物集合包括：蝴蝶、蜜蜂、麻雀、天鹅、鸳鸯、海鸥和信天翁；
- 会游泳的。鸭子、海豚、人类、天鹅、鸳鸯、海鸥和信天翁；
- 会飞且会游泳的。天鹅、鸳鸯、海鸥和信天翁。

“会飞的”、“会游泳的”这两个抽象对应的事物集合，要大于“会飞且会游泳的”所对应的事物集合空间，也就是说“会飞的”、“会游泳的”这两个抽象程度更高。

```go
// 会飞的
type Flyable interface {
    Fly()
}

// 会游泳的
type Swimable interface {
    Swim()
}

// 会飞且会游泳的
type FlySwimable interface {
    Flyable
    Swimable
}
```

![](images/image-20240711203516982.png)

Flyable只有一个Fly方法，FlySwimable则包含两个方法Fly和Swim。我们看到，具有更少方法的Flyable的抽象程度相对于FlySwimable要高，包含的事物集合（7种动物）也要比FlySwimable的事物集合（4种动物）大。也就是说，接口越小（接口方法少)，抽象程度越高，对应的事物集合越大。

#### 第二点：小接口易于实现和测试

#### 第三点：小接口表示的“契约”职责单一，易于复用组合

Go推崇通过组合的方式构建程序。Go开发人员一般会尝试通过嵌入其他已有接口类型的方式来构建新接口类型，就像通过嵌入io.Reader和io.Writer构建io.ReadWriter那样。

那构建时，如果有众多候选接口类型供我们选择，我们会怎么选择呢？

显然，我们会选择那些**新接口类型需要的契约职责，同时也要求不要引入我们不需要的契约职责**。在这样的情况下，拥有单一或少数方法的小接口便更有可能成为我们的目标，而那些拥有较多方法的大接口，可能会因引入了诸多不需要的契约职责而被放弃。由此可见，小接口更契合Go的组合思想，也更容易发挥出组合的威力。

### 28.4 定义小接口，你可以遵循的几点

- 首先，**别管接口大小，先抽象出接口**。

尽管接口不是Go独有的，但**专注于接口是编写强大而灵活的Go代码的关键**。因此，在定义小接口之前，我们需要先针对==问题领域==进行深入理解，聚焦抽象并发现接口，就像下图所展示的那样，先针对领域对象的行为进行抽象，形成一个接口集合：

![](images/image-20240704184513581.png)

**初期，我们先不要介意这个接口集合中方法的数量**，因为对问题域的理解是循序渐进的，在第一版代码中直接定义出小接口可能并不现实。而且，标准库中的io.Reader和io.Writer也不是在Go刚诞生时就有的，而是在发现对网络、文件、其他字节数据处理的实现十分相似之后才抽象出来的。并且越偏向业务层，抽象难度就越高，这或许也是前面图中Go标准库小接口（1~3个方法）占比略高于Docker和Kubernetes的原因。

- 第二，将大接口拆分为小接口。

![](images/image-20240704184559349.png)

- 最后，我们要注意接口的**单一契约职责**。



## 29 接口：为什么nil接口不等于nil？

接口是Go这门静态语言中唯一“==动静兼备==”的语法特性。

### 29.1 接口的静态特性与动态特性

接口的**静态特性**体现在**接口类型变量具有静态类型**，比如`var err error`中变量err的静态类型为error。拥有静态类型，那就意味着编译器会在编译阶段对所有接口类型变量的赋值操作进行类型检查，编译器会**检查右值的类型是否实现了该接口方法集合中的所有方法**。如果不满足，就会报错：

```go
var err error = 1 // cannot use 1 (type int) as type error in assignment: int does not implement error (missing Error method)
```

而接口的**动态特性**，就体现在**接口类型变量在运行时还存储了右值的==真实类型信息==**，这个右值的真实类型被称为接口类型变量的==动态类型==。

```go
var err error
err = errors.New("error1")
fmt.Printf("%T\n", err)  // *errors.errorString
```

这个示例通过`errros.New`构造了一个错误值，赋值给了error接口类型变量err，并通过fmt.Printf函数输出接口类型变量err的动态类型为`*errors.errorString`。

接口的这种“动静皆备”的特性的好处是：

- 首先，接口类型变量在程序**运行时可以被赋值为不同的动态类型变量**，每次赋值后，接口类型变量中存储的动态类型信息都会发生变化，这让Go语言可以像动态语言（比如Python）那样拥有使用Duck Typing（鸭子类型）的灵活性。所谓==鸭子类型==，就是指某类型所表现出的特性（比如是否可以作为某接口类型的右值），不是由其**基因**（比如C++中的父类）决定的，而是由类型所表现出来的**行为**（比如类型拥有的方法）决定的。

```go
type QuackableAnimal interface {
    Quack()
}

type Duck struct{}

func (Duck) Quack() {
    println("duck quack!")
}

type Dog struct{}

func (Dog) Quack() {
    println("dog quack!")
}

type Bird struct{}

func (Bird) Quack() {
    println("bird quack!")
}                         
                          
func AnimalQuackInForest(a QuackableAnimal) {
    a.Quack()             
}                         
                          
func main() {             
    animals := []QuackableAnimal{new(Duck), new(Dog), new(Bird)}
    for _, animal := range animals {
        AnimalQuackInForest(animal)
    }  
}
```

接口类型QuackableAnimal来代表具有“会叫”这一特征的动物，而Duck、Bird和Dog类型各自都具有这样的特征，于是我们可以将这三个类型的变量赋值给QuackableAnimal接口类型变量a。每次赋值，变量a中存储的动态类型信息都不同，Quack方法的执行结果将根据变量a中存储的动态类型信息而定。

这里的Duck、Bird、Dog都是“鸭子类型”，但它们之间并没有什么联系，之所以能作为右值赋值给QuackableAnimal类型变量，只是因为他们表现出了QuackableAnimal所要求的特征罢了。

不过，与动态语言不同的是，Go接口还可以**保证“动态特性”使用时的安全性**。比如，编译器在编译期就可以捕捉到将int类型变量传给QuackableAnimal接口类型变量这样的明显错误，决不会让这样的错误遗漏到运行时才被发现。

### 29.2 nil error值 != nil

```go
type MyError struct {
    error
}

var ErrBad = MyError{
    error: errors.New("bad things happened"),
}

func bad() bool {
    return false
}

func returnsError() error {
    var p *MyError = nil
    if bad() {
        p = &ErrBad
    }
    return p
}

func main() {
    err := returnsError()
    if err != nil {
        fmt.Printf("error occur: %+v\n", err)  // 结果
        return
    }
    fmt.Println("ok")
}
```



### 29.3 接口类型变量的内部表示 ❤️

接口类型“动静兼备”的特性也决定了它的变量的内部表示绝不像一个静态类型变量（如int、float64）那样简单。

在`$GOROOT/src/runtime/runtime2.go`中找到接口类型变量在运行时的表示：

```go
// $GOROOT/src/runtime/runtime2.go
type iface struct {
    tab  *itab
    data unsafe.Pointer
}

type eface struct {
    _type *_type
    data  unsafe.Pointer
} 
```

在运行时层面，接口类型变量有两种内部表示:

- `eface`用于表示没有方法的空接口（empty interface）类型变量，也就是`interface{}`类型的变量；
- `iface`用于表示其余拥有方法的接口interface类型变量。

这两个结构的共同点是它们**都有两个指针字段**，并且第二个指针字段的功能相同，都是指向**当前赋值给该接口类型变量的动态类型变量的值**。

不同点就在于eface表示的空接口类型并**没有方法列表**，因此它的第一个指针字段指向一个`_type`类型结构，这个结构为该接口类型变量的**==动态类型的信息==**，它的定义是这样的：

```go
// $GOROOT/src/runtime/type.go
type _type struct {
    size       uintptr
    ptrdata    uintptr // size of memory prefix holding all pointers
    hash       uint32
    tflag      tflag
    align      uint8
    fieldAlign uint8
    kind       uint8
    // function for comparing objects of this type
    // (ptr to object A, ptr to object B) -> ==?
    equal func(unsafe.Pointer, unsafe.Pointer) bool
    // gcdata stores the GC type data for the garbage collector.
    // If the KindGCProg bit is set in kind, gcdata is a GC program.
    // Otherwise it is a ptrmask bitmap. See mbitmap.go for details.
    gcdata    *byte
    str       nameOff
    ptrToThis typeOff
}
```

而iface除了要存储动态类型信息之外，还要存储**接口本身的信息**（**==接口的类型信息、方法列表信息==**等）以及**==动态类型所实现的方法的信息==**，因此iface的第一个字段指向一个itab类型结构。

```go
// $GOROOT/src/runtime/runtime2.go
type itab struct {
    inter *interfacetype
    _type *_type
    hash  uint32 // copy of _type.hash. Used for type switches.
    _     [4]byte
    fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
}
```

`interfacetype`结构，存储着这个接口类型自身的信息，由**类型信息（typ）、包路径名（pkgpath）和接口方法集合切片（mhdr）**组成。

```go
// $GOROOT/src/runtime/type.go
type interfacetype struct {
    typ     _type
    pkgpath name
    mhdr    []imethod
}
```

字段`_type`则存储着这个接口类型变量的动态类型的信息，字段`fun`则是动态类型**已实现的接口方法的调用地址数组**。

一个用eface表示的空接口类型变量的例子：

```go
type T struct {
    n int
    s string
}

func main() {
    var t = T {
        n: 17,
        s: "hello, interface",
    }
    
    var ei interface{} = t // Go运行时使用eface结构表示ei
}
```

这个例子中的空接口类型变量ei在Go运行时的表示是这样的：

![](images/image-20240711225025399.png)

空接口类型的表示较为简单，图中上半部分_type字段指向它的动态类型T的类型信息，下半部分的data则是指向一个T类型的实例值。

更复杂的用iface表示非空接口类型变量的例子：

```go
type T struct {
    n int
    s string
}

func (T) M1() {}
func (T) M2() {}

type NonEmptyInterface interface {
    M1()
    M2()
}

func main() {
    var t = T{
        n: 18,
        s: "hello, interface",
    }
    var i NonEmptyInterface = t
} 
```

NonEmptyInterface接口类型变量在Go运行时表示的示意图：

![](images/image-20240711225301398.png)

每个接口类型变量在运行时的表示都是由两部分组成的，针对不同接口类型我们可以简化记作：`eface(_type, data)`和`iface(tab, data)`。

eface和iface的第一个字段tab和_type可以统一看作是**动态类型的类型信息**。

Go语言中每种类型都会有唯一的`_type`信息，无论是内置原生类型，还是自定义类型都有。Go运行时会为程序内的全部类型建立只读的共享`_type`信息表，因此拥有相同动态类型的同类接口类型变量的`_type/tab`信息是相同的。

而接口类型变量的data部分则是指向一个**动态分配的内存空间**，这个内存空间存储的是赋值给接口类型变量的动态类型变量的值。未显式初始化的接口类型变量的值为nil，也就是这个变量的`_type/tab`和data都为nil。

也就是说，我们判断两个接口类型变量是否相同，只需要判断_type/tab是否相同，以及data指针指向的内存空间所存储的数据值是否相同就可以了。这里要注意不是data指针的值相同噢。

由于eface和iface是runtime包中的非导出结构体定义，不能直接在包外使用，所以也就无法直接访问到两个结构体中的数据。不过，Go语言提供了`println`预定义函数，可以用来输出eface或iface的两个指针字段的值。

在编译阶段，编译器会根据要输出的参数的类型将println替换为特定的函数:

```go
// $GOROOT/src/runtime/print.go
func printeface(e eface) {
    print("(", e._type, ",", e.data, ")")
}

func printiface(i iface) {
    print("(", i.tab, ",", i.data, ")")
}
```



使用println函数输出各类接口类型变量的内部表示信息，并结合输出结果，解析接口类型变量的等值比较操作:

#### 第一种：nil接口变量

```go
func printNilInterface() {
	// nil接口变量
	var i interface{} // 空接口类型
	var err error     // 非空接口类型
	println(i)
	println(err)
	println("i = nil:", i == nil)
	println("err = nil:", err == nil)
	println("i = err:", i == err)
}

/*
(0x0,0x0)
(0x0,0x0)
i = nil: true
err = nil: true
i = err: true
*/
```

无论是空接口类型还是非空接口类型变量，一旦变量值为nil，那么它们内部表示均为(0x0,0x0)，也就是类型信息、数据值信息均为空。

#### 第二种：空接口类型变量

```go
func printEmptyInterface() {
   var eif1 interface{} // 空接口类型
   var eif2 interface{} // 空接口类型
   var n, m int = 17, 18

   eif1 = n
   eif2 = m

   println("eif1:", eif1)
   println("eif2:", eif2)
   println("eif1 = eif2:", eif1 == eif2) // false

   eif2 = 17
   println("eif1:", eif1)
   println("eif2:", eif2)
   println("eif1 = eif2:", eif1 == eif2) // true

   eif2 = int64(17)
   println("eif1:", eif1)
   println("eif2:", eif2)
   println("eif1 = eif2:", eif1 == eif2) // false
}
```

对于空接口类型变量，只有_type和data所指数据内容一致的情况下，两个空接口类型变量之间才能划等号。

#### 第三种：非空接口类型变量

```go
type T int

func (t T) Error() string {
	return "bad error"
}
func printNonEmptyInterface() {
	var err1 error // 非空接口类型
	var err2 error // 非空接口类型
	err1 = (*T)(nil)
	println("err1:", err1)
	println("err1 = nil:", err1 == nil)

	err1 = T(5)
	err2 = T(6)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)

	err2 = fmt.Errorf("%d\n", 5)
	println("err1:", err1)
	println("err2:", err2)
	println("err1 = err2:", err1 == err2)
}
/**
err1: (0x102e30558,0x0)
err1 = nil: false
err1: (0x102e30578,0x102e0feb8)
err2: (0x102e30578,0x102e0fec0)
err1 = err2: false
err1: (0x102e30578,0x102e0feb8)
err2: (0x102e305d8,0x14000010040)
err1 = err2: false
 */
```

第一个输出err1是`(0x102e30558,0x0)`，也就是**非空接口类型变量的类型信息并不为空，数据指针为空**，因此它与`nil（0x0,0x0）`之间不能划等号。

这就解释了[29.2](#29.2 nil error值 != nil)中的问题。

#### 第四种：空接口类型变量与非空接口类型变量的等值比较

```go
func printEmptyInterfaceAndNonEmptyInterface() {
	var eif interface{} = T(5)
	var err error = T(5)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err) // true

	err = T(6)
	println("eif:", eif)
	println("err:", err)
	println("eif = err:", eif == err) // false
}
```

接口类型变量和非空接口类型变量内部表示的结构有所不同（第一个字段：`_type` vs. `tab`)，两者似乎一定不能相等。但Go在进行等值比较时，类型比较使用的是eface的`_type`和iface的`tab._type`。

### 29.4 输出接口类型变量内部表示的详细信息 ❤️

println输出的接口类型变量的内部表示信息，有些时候又显得过于简略。

通过“**复制代码**”的方式将它们拿到runtime包外面来



### 29.5 接口类型的装箱（boxing）原理

==装箱（boxing）==是编程语言领域的一个基础概念，一般是指**把一个值类型转换成引用类型**，比如在支持装箱概念的Java语言中，将一个int变量转换成Integer对象就是一个装箱操作。

在Go语言中，**将任意类型赋值给一个接口类型变量**也是装箱操作。

```go
// interface_internal.go
  type T struct {
      n int
      s string
  }
  
  func (T) M1() {}
  func (T) M2() {}
  
  type NonEmptyInterface interface {
      M1()
      M2()
  }
  
  func main() {
      var t = T{
          n: 17,
          s: "hello, interface",
      }
      var ei interface{}
      ei = t
 
      var i NonEmptyInterface
      i = t
      fmt.Println(ei)
      fmt.Println(i)
  }
```



```sh
go tool compile -S interface_internal.go > interface_internal.s
```

> 问题： 🔖
>
> `could not import fmt (file not found)`
>
> 可能原因：https://github.com/golang/go/issues/58629



## 30 接口：Go中最强大的魔法

### 30.1 一切皆组合

如果C++和Java是关于类型**层次**结构和类型**分类**的语言，那么Go则是关于**组合**的语言。

如果把Go应用程序比作是一台机器的话，那么组合关注的就是如何将散落在各个包中的“零件”关联并组装到一起。

正交性

> ==正交（Orthogonality）==是从几何学中借用的术语，说的是如果两条线以直角相交，那么这两条线就是正交的，比如我们在代数课程中经常用到的坐标轴就是这样。用向量术语说，**这两条直线互不依赖，沿着某一条直线移动，你投影到另一条直线上的位置不变。**

在计算机技术中，正交性用于表示某种**不相依赖性或是解耦性**。如果两个或更多事物中的一个发生变化，不会影响其他事物，那么这些事物就是正交的。比如，在设计良好的系统中，数据库代码与用户界面是正交的：你可以改动界面，而不影响数据库；更换数据库，而不用改动界面。

**编程语言的语法元素间和语言特性也存在着正交的情况，并且通过将这些正交的特性组合起来，我们可以实现更为高级的特性**。在语言设计层面，Go语言就为广大Gopher提供了诸多**正交的语法元素**供后续组合使用，包括：🔖

- Go语言无类型体系（Type Hierarchy），没有父子类的概念，类型定义是正交独立的；
- 方法和类型是正交的，每种类型都可以拥有自己的方法集合，方法本质上只是一个将receiver参数作为第一个参数的函数而已；
- 接口与它的实现者之间无“显式关联”，也就说接口与Go语言其他部分也是正交的。

在这些正交语法元素当中，接口作为**Go语言提供的具有天然正交性的语法元素**，在Go程序的静态结构搭建与耦合设计中扮演着至关重要的角色。

![](images/image-20240704185654091.png)

### 30.2 垂直组合

类型嵌入（Type Embedding）

#### 第一种：通过嵌入接口构建接口

```go
// $GOROOT/src/io/io.go
type ReadWriter interface {
    Reader
    Writer
}
```

#### 第二种：通过嵌入接口构建结构体类型

```go
type MyReader struct {
	io.Reader // underlying reader
	N int64   // max bytes remaining
}
```

#### 第三种：通过嵌入结构体类型构建新结构体类型

### 30.3 水平组合 🔖

接口分离原则（ISP原则，Interface Segregation Principle）



```go
func Save(f *os.File, data []byte) error
```

改成：

```go
func Save(w io.Writer, data []byte) error
```

io.Writer仅包含一个Write方法，而且这个方法恰恰是Save唯一需要的方法。

以io.Writer接口类型表示数据写入的目的地，既可以支持向磁盘写入，也可以支持向网络存储写入，并支持任何实现了Write方法的写入行为，这让Save函数的**扩展性**得到了质的提升。

对Save函数的测试也将变得十分容易:

```go
func TestSave(t *testing.T) {
    b := make([]byte, 0, 128)
    buf := bytes.NewBuffer(b)
    data := []byte("hello, golang")
    err := Save(buf, data)
    if err != nil {
        t.Errorf("want nil, actual %s", err.Error())
    }

    saved := buf.Bytes()
    if !reflect.DeepEqual(saved, data) {
        t.Errorf("want %s, actual %s", string(data), string(saved))
    }
}
```

由于bytes.Buffer实现了Write方法，进而实现了io.Writer接口，可以合法地将变量buf传递给Save函数。此后过程中，不需要创建任何磁盘文件或建立任何网络连接。

### 30.4 接口应用的几种模式

通过接口进行水平组合的基本模式就是：**使用接受接口类型参数的函数或方法**。在这个基本模式基础上，还有其他几种“衍生品”。

#### 基本模式

接受接口类型参数的函数或方法是水平组合的基本语法:

```go
func YourFuncName(param YourInterfaceType)
```

![](images/image-20240711233617283.png)

函数/方法参数中的接口类型作为“关节（连接点）”，支持将位于多个包中的多个类型与YourFuncName函数连接到一起，共同实现某一新特性。

同时，接口类型和它的实现者之间隐式的关系却在不经意间满足了：依赖抽象（DIP）、里氏替换原则（LSP）、接口隔离（ISP）等代码设计原则，这在其他语言中是需要很“刻意”地设计谋划的，但对Go接口来看，这一切却是自然而然的。

#### 创建模式

Go社区流传一个经验法则：“**接受接口，返回结构体（Accept interfaces, return structs）**”，这其实就是一种把接口作为“关节”的应用模式。我这里把它叫做**创建模式**，是因为这个经验法则多用于创建某一结构体类型的实例。

```go
// $GOROOT/src/sync/cond.go
type Cond struct {
    ... ...
    L Locker
}

func NewCond(l Locker) *Cond {
    return &Cond{L: l}
}

// $GOROOT/src/log/log.go
type Logger struct {
    mu     sync.Mutex 
    prefix string     
    flag   int        
    out    io.Writer  
    buf    []byte    
}

func New(out io.Writer, prefix string, flag int) *Logger {
    return &Logger{out: out, prefix: prefix, flag: flag}
}

// $GOROOT/src/log/log.go
type Writer struct {
    err error
    buf []byte
    n   int
    wr  io.Writer
}

func NewWriterSize(w io.Writer, size int) *Writer {
    // Is it already a Writer?
    b, ok := w.(*Writer)
    if ok && len(b.buf) >= size {
        return b
    }
    if size <= 0 {
        size = defaultBufSize
    }
    return &Writer{
        buf: make([]byte, size),
        wr:  w,
    }
} 
```



#### 包装器模式

在基本模式的基础上，当返回值的类型与参数类型相同时:

```go
func YourWrapperFunc(param YourInterfaceType) YourInterfaceType
```

通过这个函数，我们可以实现对输入参数的类型的包装，并在不改变被包装类型（输入参数类型）的定义的情况下，返回具备新功能特性的、实现相同接口类型的新类型。这种接口应用模式我们叫它**包装器模式**，也叫**装饰器模式**。包装器多用于**对输入数据的过滤、变换等**操作。

```go
// $GOROOT/src/io/io.go
func LimitReader(r Reader, n int64) Reader { return &LimitedReader{r, n} }

type LimitedReader struct {
    R Reader // underlying reader
    N int64  // max bytes remaining
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
    // ... ...
}
```



#### 适配器模式

适配器模式不是基本模式的直接衍生模式，但这种模式是中间件模式的前提。

适配器模式的核心是**适配器函数类型（Adapter Function Type）**，它是一个辅助水平组合实现的“工具”类型。



#### 中间件（Middleware）

中间件（Middleware）这个词的含义可大可小。在Go Web编程中，“中间件”常常指的是一个实现了http.Handler接口的http.HandlerFunc类型实例。实质上，这里的**中间件就是包装模式和适配器模式结合的产物**。



### 30.5 尽量避免使用空接口作为函数参数类型

> 空接口不提供任何信息（The empty interface says nothing）。



## 31 并发：Go的并发方案实现方案是怎样的？❤️

Go的设计者将面向多核、**原生支持并发**作为了Go语言的设计目标之一，并将面向并发作为Go的设计哲学。

### 31.1 什么是并发？

并发（concurrency） 

并行（parallelism）

很久以前，面向大众消费者的主流处理器（CPU）都是单核的，操作系统的**基本调度与执行单元**是==进程（process）==。这个时候，**用户层的应用**有两种设计方式，一种是==单进程应用==，也就是每次启动一个应用，操作系统都只启动一个进程来运行这个应用。

单进程应用的情况下，用户层应用、操作系统进程以及处理器之间的关系:

![](images/image-20240711234318877.png)

这个设计下，每个单进程应用对应一个操作系统进程，操作系统内的多个进程按时间片大小，被轮流调度到仅有的一颗单核处理器上执行。换句话说，这颗单核处理器在某个时刻只能执行一个进程对应的程序代码，两个进程不存在并行执行的可能。

**==并行（parallelism）==，指的就是在同一时刻，有两个或两个以上的任务（这里指进程）的代码在处理器上执行**。从这个概念也可以知道，多个处理器或多核处理器是并行执行的必要条件。

总的来说，单进程应用的设计比较简单，它的内部仅有一条代码执行流，代码从头执行到尾，**不存在==竞态==，无需考虑同步问题**。

用户层的另外一种设计方式，就是==多进程应用==，也就是应用通过fork等系统调用创建多个子进程，共同实现应用的功能。多进程应用的情况下，用户层应用、操作系统进程以及处理器之间的关系:

![](images/image-20240711234429579.png)

以图中的App1为例，这个应用设计者将应用内部划分为多个模块，每个模块用一个进程承载执行，每个模块都是一个单独的执行流，这样，App1内部就有了多个独立的代码执行流。

但限于当前仅有一颗单核处理器，这些进程（执行流）依旧无法并行执行，无论是App1内部的某个模块对应的进程，还是其他App对应的进程，都得逐个按时间片被操作系统调度到处理器上执行。

**粗略看起来，多进程应用与单进程应用相比并没有什么质的提升。那我们为什么还要将应用设计为多进程呢？**

这更多是从**应用的结构**角度去考虑的，多进程应用由于**将==功能职责==做了划分**，并指定专门的模块来负责，所以从结构上来看，要比单进程更为清晰简洁，可读性与可维护性也更好。**这种将程序分成多个可独立执行的部分的结构化程序的设计方法，就是==并发设计==**。采用了并发设计的应用也可以看成是**一组独立执行的模块的组合**。

不过，进程并不适合用于承载采用了并发设计的应用的模块执行流。因为进程是操作系统中资源拥有的基本单位，它不仅包含应用的代码和数据，还有**系统级的资源**，比如**文件描述符、内存地址空间**等等。进程的“包袱”太重，这导致它的创建、切换与撤销的代价都很大。

于是线程便走入了人们的视野，==线程==就是**运行于进程上下文中的更轻量级的执行流**。同时随着处理器技术的发展，多核处理器硬件成为了主流，这让真正的并行成为了可能，于是主流的应用设计模型变成了这样：

![](images/image-20240717152619145.png)

基于线程的应用通常采用**单进程多线程的模型**，一个应用对应一个进程，应用通过并发设计将自己**划分为多个模块，每个模块由一个线程独立承载执行**。多个线程共享这个进程所拥有的资源，但线程作为执行单元可被独立调度到处理器上运行。

线程的创建、切换与撤销的代价相对于进程是要小得多。当这个应用的多个线程同时被调度到不同的处理器核上执行时，我们就说这个应用是并行的。

> Rob Pike：**并发不是并行，并发关乎结构，并行关乎执行**。

总的来说，并发是**在==应用设计与实现阶段==要考虑的问题**。并发考虑的是如何将应用划分为多个互相配合的、可独立执行的模块的问题。采用并发设计的程序并不一定是并行执行的。

在不满足并行必要条件的情况下（也就是仅有一个单核CPU的情况下），即便是采用并发设计的程序，依旧不可以并行执行。而在满足并行必要条件的情况下，采用并发设计的程序是可以并行执行的。而那些没有采用并发设计的应用程序，除非是启动多个程序实例，否则是无法并行执行的。

在多核处理器成为主流的时代，即使采用并发设计的应用程序以单实例的方式运行，其中的每个内部模块也都是运行于一个单独的线程中的，多核资源也可以得到充分利用。而且，**==并发让并行变得更加容易==**，采用并发设计的应用可以将负载自然扩展到各个CPU核上，从而提升处理器的利用效率。

在传统编程语言（如C、C++等）中，基于**多线程模型**的应用设计就是一种典型的并发程序设计。但传统编程语言并非面向并发而生，没有对并发设计提供过多的帮助。并且，这些语言<u>多以操作系统线程作为承载分解后的代码片段（模块）的执行单元</u>，由操作系统执行调度。传统支持并发的方式的不足：

- 首先就是复杂。

  **创建容易==退出难==**。如果你做过C/C++编程，那你肯定知道，如果我们要利用libpthread库中提供的API创建一个线程，虽然要传入的参数个数不少，但好歹还是可以接受的。但一旦涉及线程的退出，就要考虑**新创建的线程是否要与主线程分离**（detach），还是需要**主线程等待子线程终止（join）并获取其终止状态**？又或者是否需要在**新线程中设置取消点（cancel point）来保证被主线程取消（cancel）的时候能顺利退出**。

  而且，并发执行单元间的==通信困难且易错==。多个线程之间的通信虽然有多种机制可选，但用起来也是相当复杂。并且一旦涉及共享内存，就会用到各种锁互斥机制，死锁便成为家常便饭。另外，线程栈大小也需要设定，开发人员需要选择使用默认的，还是自定义设置。

- 第二就是难于规模化（scale）。每个线程占用的资源不小，操作系统调度切换线程的代价也不小。

  对于很多网络服务程序来说，由于不能大量创建线程，只能选择在少量线程里做网络多路复用的方案，也就是使用epoll/kqueue/IoCompletionPort这套机制，即便有像[libevent](https://github.com/libevent/libevent)和[libev](http://software.schmorp.de/pkg/libev.html)这样的第三方库帮忙，写起这样的程序也是很不容易的，存在大量**钩子回调**，给开发人员带来不小的**心智负担**。



### 31.2 Go的并发方案：goroutine

Go并没有使用操作系统线程作为承载分解后的代码片段（模块）的基本执行单元，而是实现了`goroutine`这一**由Go运行时（runtime）负责调度的、轻量的用户级线程**，为并发程序设计提供原生支持。【Goroutine 可以被称为**协程**或**Go 协程**。】

相比传统操作系统线程来说，goroutine的优势主要是：

- 资源占用小，每个goroutine的**初始栈**大小仅为==2k==；
- 由Go运行时而不是操作系统调度，goroutine上下文切换在用户层完成，开销更小；
- 在**==语言层==**面而不是通过标准库提供。goroutine由`go`关键字创建，**一退出就会被回收或销毁**，开发体验更佳；
- 语言内置`channel`作为goroutine间通信原语，为并发设计提供了强大支撑。

通过并发设计的Go应用可以更好地、更自然地适应**规模化（scale）**。

比如，当应用被分配到更多计算资源，或者计算处理硬件增配后，Go应用不需要再进行结构调整，就可以充分利用新增的计算资源。而且，经过并发设计后的Go应用也会更加契合Gopher们的开发分工协作。

#### goroutine的基本用法

**并发**是一种能力，它让你的程序可以由若干个代码片段**组合**而成，并且每个片段都是独立运行的。

无论是==Go自身运行时代码==还是==用户层Go代码==，都无一例外地运行在goroutine中。

Go语言通过`go关键字+函数/方法`的方式创建一个goroutine。创建后，新goroutine将拥有独立的代码执行流，并与创建它的goroutine一起被Go运行时调度。

```go
go fmt.Println("I am a goroutine")

var c = make(chan int)
go func(a, b int) {
    c <- a + b
}(3,4)
 
// $GOROOT/src/net/http/server.go
c := srv.newConn(rw)
go c.serve(connCtx)
```

上面例子中，通过go关键字，可以基于已有的具名函数/方法创建goroutine，也可以基于匿名函数/闭包创建goroutine。

多数情况下，不需要考虑对goroutine的退出进行控制：**goroutine的执行函数的返回，就意味着goroutine退出。**

如果main goroutine退出了，那么也意味着整个应用程序的退出。此外，你还要注意的是，**goroutine执行的函数或方法即便有返回值，Go也会忽略这些返回值**。所以，如果你要获取goroutine执行后的返回值，你需要另行考虑其他方法，比如通过goroutine间的通信来实现。

#### goroutine间的通信

传统的编程语言（比如：C++、Java、Python等）并非==面向并发==而生的，所以他们面对并发的逻辑多是**基于操作系统的线程**。并发的执行单元（线程）之间的通信，利用的也是操作系统提供的**线程或进程间通信**的原语，比如：**共享内存、信号（signal）、管道（pipe）、消息队列、套接字（socket）**等。

在这些通信原语中，使用最多、最广泛的（也是最高效的）是结合了线程同步原语（比如：锁以及更为低级的原子操作）的共享内存方式，因此，我们可以说传统语言的并发模型是**基于对内存的共享的**。

不过，这种传统的基于共享内存的并发模型很**难用**，且**易错**，尤其是在大型或复杂程序中，开发人员在设计并发程序时，需要根据线程模型对程序进行建模，同时规划线程之间的通信方式。如果选择的是高效的基于共享内存的机制，那么他们还要花费大量心思设计**线程间的同步机制**，并且在设计同步机制的时候，还要考虑多线程间复杂的内存管理，以及如何防止死锁等情况。

这种情况下，开发人员承受着巨大的心智负担，并且基于这类传统并发模型的程序难于编写、阅读、理解和维护。一旦程序发生问题，查找Bug的过程更是漫长和艰辛。

Go语言从设计伊始，就将解决上面这个传统并发模型的问题作为Go的一个目标，并在新并发模型设计中借鉴了著名计算机科学家[Tony Hoare](https://en.wikipedia.org/wiki/Tony_Hoare)提出的**==CSP==（Communicating Sequential Processes，==通信顺序进程==）**并发模型。

CSP模型旨在简化并发程序的编写，让并发程序的编写与编写顺序程序一样简单。Tony Hoare认为输入输出应该是基本的编程原语，**数据处理逻辑**（也就是CSP中的P）<u>只需调用输入原语获取数据，顺序地处理数据，并将结果数据通过输出原语输出就可以了。</u>

因此，在Tony Hoare眼中，**一个符合CSP模型的并发程序应该是一组通过输入输出原语连接起来的P的集合**。从这个角度来看，CSP理论不仅是一个并发参考模型，也是一种**并发程序的程序组织方法**。它的组合思想与Go的设计哲学不谋而合。

Tony Hoare的CSP理论中的P，也就是“Process（进程）”，是一个抽象概念，它代表**任何顺序==处理逻辑==的封装**，它获取输入数据（或从其他P的输出获取），并生产出可以被其他P消费的输出数据。CSP通信模型的示意图：

![](images/image-20240711235319311.png)

这里的P并不一定与操作系统的进程或线程划等号。在Go中，与“Process”对应的是`goroutine`。为了实现CSP并发模型中的输入和输出原语，Go还引入了**goroutine（P）之间的通信原语**`channel`。goroutine可以从channel获取输入数据，再将处理后得到的结果数据通过channel输出。通过channel将goroutine（P）组合连接在一起，让设计和编写大型并发系统变得更加简单和清晰。

上面提到的获取goroutine的退出状态，就可以使用channel原语实现：

```go
func spawn(f func() error) <-chan error {
    c := make(chan error)

    go func() {
        c <- f()
    }()

    return c
}

func main() {
    c := spawn(func() error {
        time.Sleep(2 * time.Second)
        return errors.New("timeout")
    })
    fmt.Println(<-c)
}
```

<u>在main goroutine与子goroutine之间建立了一个元素类型为error的channel，子goroutine退出时，会将它执行的函数的错误返回值写入这个channel，main goroutine可以通过读取channel的值来获取子goroutine的退出状态。</u>

虽然CSP模型已经成为Go语言支持的主流并发模型，但Go也支持传统的、基于共享内存的并发模型，并提供了基本的低级别同步原语（主要是`sync`包中的互斥锁、条件变量、读写锁、原子操作等）。

**Go始终推荐以CSP并发模型风格构建并发程序**。不过，对于局部情况，比如涉及**性能敏感的区域**或**需要保护的结构体数据**时，我们可以使用更为高效的低级同步原语（如`mutex`），保证goroutine对数据的同步访问。

### 思考

> goroutine作为Go应用的基本执行单元，它的创建、退出以及goroutine间的通信都有很多常见的模式可循。日常开发中实用的goroutine使用模式有哪些？
>



## 32 并发：聊聊Goroutine调度器的原理

> Go运行时是如何将一个个Goroutine调度到CPU上执行的？

### 32.1 Goroutine调度器

将这些Goroutine按照一定算法放到==“CPU”==上执行的程序，就被称为==Goroutine调度器（Goroutine Scheduler）==。

一个Go程序对于操作系统来说只是一个**用户层程序**，操作系统眼中只有线程。

一个Go程序 = 用户层代码 + Go运行时。

Goroutine们要竞争的“CPU”资源就是**操作系统线程**。

Goroutine调度器的任务也就明确了：**将Goroutine按照一定算法放到不同的操作系统线程中去执行**。

### 32.2 Goroutine调度器模型与演化过程 🔖

Goroutine调度器的实现不是一蹴而就的，它的调度模型与算法也是几经演化，**从最初的G-M模型、到G-P-M模型，从不支持抢占，到支持协作式抢占，再到支持基于信号的异步抢占**，Goroutine调度器经历了不断地优化与打磨。

#### G-M模型

2012年3月28日，Go 1.0

在这个调度器中，每个Goroutine对应于运行时中的一个抽象结构：==G==(Goroutine)；而被视作“物理CPU”的操作系统线程被抽象为：==M==(machine)。

调度器的工作就是将G调度到M上去运行。

G-M模型的一个重要不足：**限制了Go并发程序的伸缩性，尤其是对那些有高吞吐或并行计算需求的服务程序**。这个问题主要体现在：

- 单一全局互斥锁`(Sched.Lock)` 和集中状态存储的存在，导致所有Goroutine相关操作，比如创建、重新调度等，都要上锁；
- Goroutine传递问题：M经常在M之间传递“可运行”的Goroutine，这导致调度延迟增大，也增加了额外的性能损耗；
- 每个M都做内存缓存，导致内存占用过高，数据局部性较差；
- 由于系统调用（syscall）而形成的频繁的工作线程阻塞和解除阻塞，导致额外的性能损耗。

#### G-P-M调度模型

==Go 1.1==

[work stealing算法](http://supertech.csail.mit.edu/papers/steal.pdf)

为了解决上面的一些问题，德米特里·维尤科夫通过向G-M模型中增加了一个P，让Go调度器具有很好的伸缩性。

==P==是一个“逻辑Proccessor”，每个G（Goroutine）要想真正运行起来，首先需要被分配一个P，也就是进入到P的本地运行队列（local runq）中。对于G来说，P就是运行它的“CPU”，可以说：**在G的眼里只有P**。但从Go调度器的视角来看，真正的“CPU”是M，只有将P和M绑定，才能让P的runq中的G真正运行起来。



![](images/b5d81e17e041461ea7e78ae159f6ea3c.jpg)



此时，调度器仍然有一个令人头疼的问题，那就是**不支持抢占式调度**，这导致一旦某个G中出现死循环的代码逻辑，那么G将永久占用分配给它的P和M，而位于同一个P中的其他G将得不到调度，出现“**饿死**”的情况。

更为严重的是，当只有一个P（GOMAXPROCS=1）时，整个Go程序中的其他G都将“饿死”。于是德米特里·维尤科夫又提出了《[Go Preemptive Scheduler Design](https://docs.google.com/document/d/1ETuA2IOmnaQ4j81AtTGT40Y4_Jr6_IDASEKg0t0dBR8/edit#!)》并在**Go 1.2中实现了基于协作的“抢占式”调度**。

#### “抢占式”调度

Go 1.2

这个抢占式调度的原理就是，Go编译器在每个函数或方法的入口处加上了一段额外的代码(`runtime.morestack_noctxt`)，让运行时有机会在这段代码中检查是否需要执行抢占调度。

这种解决方案只能说局部解决了“饿死”问题，只在有函数调用的地方才能插入“抢占”代码（**埋点**），对于没有函数调用而是纯算法循环计算的G，Go调度器依然无法抢占。

比如，死循环等并没有给编译器插入抢占代码的机会，这就会导致GC在等待所有Goroutine停止时的等待时间过长，从而[导致GC延迟](https://github.com/golang/go/issues/10958)，内存占用瞬间冲高；甚至在一些特殊情况下，导致在**STW**（stop the world）时死锁。

为了解决这些问题，在==Go 1.14==版本中接受了奥斯汀·克莱门茨（Austin Clements）的[提案](https://go.googlesource.com/proposal/+/master/design/24543-non-cooperative-preemption.md)，增加了**对非协作的抢占式调度的支持**，这种抢占式调度是基于**系统信号**的，也就是通过向线程发送信号的方式来抢占正在运行的Goroutine。

除了这些大的迭代外，Goroutine的调度器还有一些小的优化改动，比如**通过文件I/O poller减少M的阻塞等**。

Go运行时已经实现了`netpoller`，这使得即便G发起网络I/O操作，也不会导致M被阻塞（仅阻塞G），也就不会导致大量线程（M）被创建出来。

但是对于文件I/O操作来说，一旦阻塞，那么线程（M）将进入挂起状态，等待I/O返回后被唤醒。这种情况下P将与挂起的M分离，再选择一个处于空闲状态（idle）的M。如果此时没有空闲的M，就会新创建一个M（线程），所以，这种情况下，大量I/O操作仍然会导致大量线程被创建。

为了解决这个问题，Go开发团队的伊恩·兰斯·泰勒（Ian Lance Taylor）在==Go 1.9==中增加了一个[针对文件I/O的Poller](https://groups.google.com/forum/#!topic/golang-dev/tT8SoKfHty0)的功能，这个功能可以像netpoller那样，在G操作那些支持监听（`pollable`）的文件描述符时，仅会阻塞G，而不会阻塞M。不过这个功能依然不能对常规文件有效，常规文件是不支持监听的（pollable）。但对于Go调度器而言，这也算是一个不小的进步了。

从Go 1.2以后，Go调度器就一直稳定在G-P-M调度模型上，尽管有各种优化和改进，但也都是基于这个模型之上的。那未来的Go调度器会往哪方面发展呢？

德米特里·维尤科夫在2014年9月提出了一个新的设计草案文档：《[NUMA‐aware scheduler for Go](https://docs.google.com/document/u/0/d/1d3iI2QWURgDIsSR6G2275vMeQ_X7w-qxM2Vp7iGwwuM/pub)》，作为对未来Goroutine调度器演进方向的一个提议，不过至今似乎这个提议也没有列入开发计划。

### 32.3 深入G-P-M模型 🔖

Go语言中Goroutine的调度、GC、内存管理等是Go语言原理最复杂、最难懂的地方，并且这三方面的内容随着Go版本的演进也在不断更新。

基于Go 1.12.7版本（支持基于协作的抢占式调度）

#### G、P和M

`$GOROOT/src/runtime/runtime2.go`

- G: 代表Goroutine，存储了Goroutine的执行栈信息、Goroutine状态以及Goroutine的任务函数等，而且G对象是可以重用的；
- P: 代表逻辑processor，**P的数量决定了系统内最大可并行的G的数量**，P的最大作用还是其拥有的各种**G对象队列、链表、一些缓存和状态**；
- M: M代表着真正的执行计算资源。在绑定有效的P后，进入一个调度循环，而调度循环的机制大致是**从P的本地运行队列以及全局队列中获取G，切换到G的执行栈上并执行G的函数，调用goexit做清理工作并回到M，如此反复**。M并不保留G状态，这是G可以跨M调度的基础。

```go
//src/runtime/runtime2.go
type g struct {
    stack      stack   // offset known to runtime/cgo
    sched      gobuf
    goid       int64
    gopc       uintptr // pc of go statement that created this goroutine
    startpc    uintptr // pc of goroutine function
    ... ...
}

type p struct {
    lock mutex

    id          int32
    status      uint32 // one of pidle/prunning/...
  
    mcache      *mcache
    racectx     uintptr

    // Queue of runnable goroutines. Accessed without lock.
    runqhead uint32
    runqtail uint32
    runq     [256]guintptr

    runnext guintptr

    // Available G's (status == Gdead)
    gfree    *g
    gfreecnt int32

    ... ...
}



type m struct {
    g0            *g     // goroutine with scheduling stack
    mstartfn      func()
    curg          *g     // current running goroutine
    ... ...
}
```

Goroutine调度器的目标，就是公平合理地将各个G调度到P上“运行”。

下面重点看看**G是如何被调度**的。

#### G被抢占调度

如果某个G没有进行系统调用（syscall）、没有进行I/O操作、没有阻塞在一个channel操作上，**调度器是如何让G停下来并调度下一个可运行的G的呢**？

答案就是：**G是被抢占调度的**。

前面说过，除非极端的无限循环，否则只要G调用函数，Go运行时就有了抢占G的机会。Go程序启动时，运行时会去启动一个名为`sysmon`的M（一般称为**==监控线程==**），这个M的特殊之处在于它不需要绑定P就可以运行（以g0这个G的形式），这个M在整个Go程序的运行过程中至关重要，下面是对sysmon被创建的部分代码以及sysmon的执行逻辑摘录：

```go
//$GOROOT/src/runtime/proc.go

// The main goroutine.
func main() {
     ... ...
    systemstack(func() {
        newm(sysmon, nil)
    })
    .... ...
}

// Always runs without a P, so write barriers are not allowed.
//
//go:nowritebarrierrec
func sysmon() {
    // If a heap span goes unused for 5 minutes after a garbage collection,
    // we hand it back to the operating system.
    scavengelimit := int64(5 * 60 * 1e9)
    ... ...

    if  .... {
        ... ...
        // retake P's blocked in syscalls
        // and preempt long running G's
        if retake(now) != 0 {
            idle = 0
        } else {
            idle++
        }
       ... ...
    }
}
```

sysmon每`20us~10ms`启动一次，主要完成了这些工作：

- 释放闲置超过5分钟的**span内存**；
- 如果超过2分钟没有垃圾回收，强制执行；
- 将长时间未处理的`netpoll`结果添加到任务队列；
- 向长时间运行的G任务发出抢占调度；
- 收回因`syscall`长时间阻塞的P；

sysmon将“向长时间运行的G任务发出抢占调度”，这个事情由函数`retake`实施：

```go
// $GOROOT/src/runtime/proc.go

// forcePreemptNS is the time slice given to a G before it is
// preempted.
const forcePreemptNS = 10 * 1000 * 1000 // 10ms

func retake(now int64) uint32 {
          ... ...
           // Preempt G if it's running for too long.
            t := int64(_p_.schedtick)
            if int64(pd.schedtick) != t {
                pd.schedtick = uint32(t)
                pd.schedwhen = now
                continue
            }
            if pd.schedwhen+forcePreemptNS > now {
                continue
            }
            preemptone(_p_)
         ... ...
}

func preemptone(_p_ *p) bool {
    mp := _p_.m.ptr()
    if mp == nil || mp == getg().m {
        return false
    }
    gp := mp.curg
    if gp == nil || gp == mp.g0 {
        return false
    }

    gp.preempt = true //设置被抢占标志

    // Every call in a go routine checks for stack overflow by
    // comparing the current stack pointer to gp->stackguard0.
    // Setting gp->stackguard0 to StackPreempt folds
    // preemption into the normal stack overflow check.
    gp.stackguard0 = stackPreempt
    return true
}
```

上面的代码中显示，**如果一个G任务运行10ms，sysmon就会认为它的运行时间太久而发出抢占式调度的请求**。一旦G的抢占标志位被设为true，那么等到这个G下一次调用函数或方法时，运行时就可以将G抢占并移出运行状态，放入队列中，等待下一次被调度。

除了这个常规调度之外，还有两个特殊情况下G的调度方法:

- 第一种：channel阻塞或网络I/O情况下的调度。

  如果G被阻塞在某个channel操作或网络I/O操作上时，G会被放置到某个等待（wait）队列中，而M会尝试运行P的下一个可运行的G。如果这个时候P没有可运行的G供M运行，那么M将解绑P，并进入挂起状态。当I/O操作完成或channel操作完成，在等待队列中的G会被唤醒，标记为可运行（runnable），并被放入到某P的队列中，绑定一个M后继续执行。

- 第二种：系统调用阻塞情况下的调度。

  如果G被阻塞在某个系统调用（system call）上，那么不光G会阻塞，执行这个G的M也会解绑P，与G一起进入挂起状态。如果此时有空闲的M，那么P就会和它绑定，并继续执行其他G；如果没有空闲的M，但仍然有其他G要去执行，那么Go运行时就会创建一个新M（线程）。

当系统调用返回后，阻塞在这个系统调用上的G会尝试获取一个可用的P，如果没有可用的P，那么G会被标记为runnable，之前的那个挂起的M将再次进入挂起状态。

### 小结

![](images/image-20250103164942657.png)

基于Goroutine的并发设计离不开一个高效的生产级调度器。Goroutine调度器演进了十余年，先后经历了**G-M模型、G-P-M模型和work stealing算法、协作式的抢占调度以及基于信号的异步抢占**等改进与优化，目前Goroutine调度器相对稳定和成熟，可以适合绝大部分生产场合。

现在的G-P-M模型和最初的G-M模型相比，通过向G-M模型中增加了一个代表**逻辑处理器**的P，使得Goroutine调度器具有了更好的伸缩性。

M是Go**代码运行的真实载体**，包括Goroutine调度器自身的逻辑也是在M中运行的。

P在G-P-M模型中占据核心地位，它拥有待调度的G的队列，同时M要想运行G必须绑定一个P。一个G被调度执行的时间不能过长，超过特定长的时间后，G会被设置为**可抢占**，并在下一次执行函数或方法时被Go运行时移出运行状态。

如果G被阻塞在某个channel操作或网络I/O操作上时，M可以不被阻塞，这避免了大量创建M导致的开销。但如果G因慢系统调用而阻塞，那么M也会一起阻塞，但在阻塞前会与P解绑，P会尝试与其他M绑定继续运行其他G。但若没有现成的M，Go运行时会建立新的M，这也是系统调用可能导致系统线程数量增加的原因，你一定要注意这一点。

### 思考

```go
func deadloop() {
    for {
    } 
}

func main() {
    go deadloop()
    for {
        time.Sleep(time.Second * 1)
        fmt.Println("I got scheduled!")
    }
}
```

问题：

1. 在一个拥有多核处理器的主机上，使用 Go 1.13.x 版本运行这个示例代码，你在命令行终端上是否能看到“I got scheduled!”输出呢？也就是 main goroutine 在创建 deadloop goroutine 之后是否能继续得到调度呢？
2. 我们通过什么方法可以让上面示例中的 main goroutine，在创建 deadloop goroutine 之后无法继续得到调度？



答案：🔖

go1.13的话加上runtime.GOMAXPROCS(1)， main goroutine在创建 deadloop goroutine 之后就无法继续得到调度。

但如果是go1.14之后的话即使加上runtime.GOMAXPROCS(1)， main goroutine在创建 deadloop goroutine 之后还是可以得到调度，应该是因为增加了对非协作的抢占式调度的支持。



## 33 并发：小channel中蕴含大智慧

Go语言的CSP模型的实现包含两个主要组成部分：

- 一个是Goroutine，它是Go应用并发设计的基本构建与执行单元；
- 另一个就是channel，它在并发模型中扮演着重要的角色。

channel既可以用来实现Goroutine间的==通信==，还可以实现Goroutine间的==同步==。

### 33.1 作为一等公民的channel

意味着可以像使用普通变量那样使用channel。

#### 创建channel

和切片、结构体、map等一样，channel也是一种复合数据类型。也就是说，我们在声明一个channel类型变量时，必须给出其具体的元素类型。

```go
var ch chan int  // 声明了一个元素为int类型的channel类型变量ch
```

为channel类型变量赋初值的唯一方法就是`make`:

```go
ch1 := make(chan int)     // 无缓冲channel
ch2 := make(chan int, 5)   // 带缓冲channel
```

`make(chan T)`创建的、元素类型为T的channel类型，是**无缓冲channel**；

`make(chan T, capacity)`创建的元素类型为T、缓冲区长度为capacity的channel类型，是**带缓冲channel**。

这两种类型的变量关于发送（send）与接收（receive）的特性是不同的。

#### 发送与接收

`<-`操作符用于对channel类型变量进行发送与接收操作：

```go
ch1 <- 13    // 将整型字面值13发送到无缓冲channel类型变量ch1中
n := <- ch1  // 从无缓冲channel类型变量ch1中接收一个整型值存储到整型变量n中
ch2 <- 17    // 将整型字面值17发送到带缓冲channel类型变量ch2中
m := <- ch2  // 从带缓冲channel类型变量ch2中接收一个整型值存储到整型变量m中
```

1️⃣由于无缓冲channel的运行时层实现不带有缓冲区，所以Goroutine对无缓冲channel的接收和发送操作是**==同步的==**。也就是说，对同一个无缓冲channel，只有对它进行接收操作的Goroutine和对它进行发送操作的Goroutine都存在的情况下，通信才能得以进行，否则单方面的操作会让对应的Goroutine陷入挂起状态，如：

```go
func main() {
    ch1 := make(chan int)
    ch1 <- 13 // fatal error: all goroutines are asleep - deadlock!
    n := <-ch1
    println(n)
}
```

创建了一个无缓冲的channel类型变量ch1，对ch1的读写都放在了一个Goroutine中。

运行结果提示我们所有Goroutine都处于休眠状态，程序处于==死锁==状态。要想解除这种错误状态，只需要将接收操作，或者发送操作放到另外一个Goroutine中就可以了：

```go
func main() {
    ch1 := make(chan int)
    go func() {
        ch1 <- 13 // 将发送操作放入一个新goroutine中执行
    }()
    n := <-ch1
    println(n)
}
```

结论：**对无缓冲channel类型的发送与接收操作，一定要==放在两个不同的Goroutine==中进行，否则会导致deadlock**。

2️⃣和无缓冲channel相反，带缓冲channel的运行时层实现带有缓冲区，因此，对带缓冲channel的发送操作在缓冲区未满、接收操作在缓冲区非空的情况下是**==异步==**的（发送或接收不需要阻塞等待）。

也就是说，对一个带缓冲channel来说，在缓冲区未满的情况下，对它进行发送操作的Goroutine并不会阻塞挂起；在缓冲区有数据的情况下，对它进行接收操作的Goroutine也不会阻塞挂起。

但当缓冲区满了的情况下，对它进行发送操作的Goroutine就会阻塞挂起；当缓冲区为空的情况下，对它进行接收操作的Goroutine也会阻塞挂起。

几个关于带缓冲channel的操作的例子:

```go
ch2 := make(chan int, 1)
n := <-ch2 // 由于此时ch2的缓冲区中无数据，因此对其进行接收操作将导致goroutine挂起

ch3 := make(chan int, 1)
ch3 <- 17  // 向ch3发送一个整型数17
ch3 <- 27  // 由于此时ch3中缓冲区已满，再向ch3发送数据也将导致goroutine挂起
```

操作符`<-`还可以声明**只发送channel类型**（send-only）和**只接收channel类型**（recv-only）: 🔖

```go
ch1 := make(chan<- int, 1) // 只发送channel类型
ch2 := make(<-chan int, 1) // 只接收channel类型

<-ch1       // invalid operation: <-ch1 (receive from send-only type chan<- int)
ch2 <- 13   // invalid operation: ch2 <- 13 (send to receive-only type <-chan int)
```

**试图从一个只发送channel类型变量中接收数据，或者向一个只接收channel类型发送数据，都会导致编译错误**。

通常只发送channel类型和只接收channel类型，会被用作函数的参数类型或返回值，用于限制对channel内的操作，或者是明确可对channel进行的操作的类型，例子：

```go
func produce(ch chan<- int) {
    for i := 0; i < 10; i++ {
        ch <- i + 1
        time.Sleep(time.Second)
    }
    close(ch)
}

func consume(ch <-chan int) {
    for n := range ch {
        println(n)
    }
}

func main() {
    ch := make(chan int, 5)
    var wg sync.WaitGroup
    wg.Add(2)
    go func() {
        produce(ch)
        wg.Done()
    }()

    go func() {
        consume(ch)
        wg.Done()
    }()

    wg.Wait()
}
```

启动了两个Goroutine，分别代表生产者（produce）与消费者（consume）。生产者只能向channel中发送数据，我们使用`chan<- int`作为produce函数的参数类型；消费者只能从channel中接收数据，我们使用`<-chan int`作为consume函数的参数类型。

在消费者函数consume中，我们使用了for range循环语句来从channel中接收数据，for range会阻塞在对channel的接收操作上，直到channel中有数据可接收或channel被关闭循环，才会继续向下执行。channel被关闭后，for range循环也就结束了。

#### 关闭channel



```go
ch := make(chan int, 5)
close(ch)
ch <- 13 // panic: send on closed channel
```



#### select 

同时对多个channel进行操作时，会结合Go为CSP并发模型提供的另外一个原语**select**，一起使用。

通过select，可以同时在多个channel上进行发送/接收操作：

```go
select {
case x := <-ch1:     // 从channel ch1接收数据
	... ...

case y, ok := <-ch2: // 从channel ch2接收数据，并根据ok值判断ch2是否已经关闭
	... ...

case ch3 <- z:       // 将z值发送到channel ch3中:
	... ...

default:             // 当上面case中的channel通信均无法实施时，执行该默认分支
}
```

当select语句中没有default分支，而且所有case中的channel操作都阻塞了的时候，整个select语句都将被阻塞，直到某一个case上的channel变成可发送，或者某个case上的channel变成可接收，select语句才可以继续进行下去。



### 33.2 无缓冲channel的惯用法 🔖

无缓冲channel兼具通信和同步特性，在并发程序中应用颇为广泛。

#### 第一种用法：用作信号传递

- 1对1通知信号

```go
type signal struct{}

func worker() {
    println("worker is working...")
    time.Sleep(1 * time.Second)
}

func spawn(f func()) <-chan signal {
    c := make(chan signal)
    go func() {
        println("worker start to work...")
        f()
        c <- signal{}
    }()
    return c
}

func main() {
    println("start a worker...")
    c := spawn(worker)
    <-c
    fmt.Println("worker work done!")
}
```



- 1对n通知信号，常被用于协调多个Goroutine一起工作

```go
func worker(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(1 * time.Second)
	fmt.Printf("worker %d: works done\n", i)
}

type signal struct{}

func spawnGroup(f func(i int), num int, groupSignal <-chan signal) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			<-groupSignal
			fmt.Printf("worker %d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal{}
	}()
	return c
}

func main() {
	fmt.Println("start a group of workers...")
	groupSignal := make(chan signal)
	c := spawnGroup(worker, 5, groupSignal)
	time.Sleep(5 * time.Second)
	fmt.Println("the group of workers start to work...")
	close(groupSignal)
	<-c
	fmt.Println("the group of workers work done!")
}
```

这个例子中，main goroutine创建了一组5个worker goroutine，这些Goroutine启动后会阻塞在名为groupSignal的无缓冲channel上。main goroutine通过`close(groupSignal)`向所有worker goroutine广播“开始工作”的信号，收到groupSignal后，所有worker goroutine会**“同时”**开始工作，就像起跑线上的运动员听到了裁判员发出的起跑信号枪声。

运行结果：

```go
start a group of workers...
the group of workers start to work...
worker 3: start to work...
worker 3: is working...
worker 4: start to work...
worker 4: is working...
worker 1: start to work...
worker 1: is working...
worker 5: start to work...
worker 5: is working...
worker 2: start to work...
worker 2: is working...
worker 3: works done
worker 4: works done
worker 5: works done
worker 1: works done
worker 2: works done
the group of workers work done!
```

可以看到，关闭一个无缓冲channel会让所有阻塞在这个channel上的接收操作返回，从而实现了一种1对n的**“广播”**机制。

#### 第二种用法：用于替代锁机制

无缓冲channel具有同步特性，这让它在某些场合可以替代锁，让程序更加清晰，可读性也更好。对比：

- 一个传统的、基于“共享内存”+“互斥锁”的Goroutine安全的计数器

```go
type counter struct {
    sync.Mutex
    i int
}

var cter counter

func Increase() int {
    cter.Lock()
    defer cter.Unlock()
    cter.i++
    return cter.i
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            v := Increase()
            fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
            wg.Done()
        }(i)
    }

    wg.Wait()
}
```

这个示例中，使用了一个带有互斥锁保护的全局变量作为计数器，所有要操作计数器的Goroutine共享这个全局变量，并在互斥锁的同步下对计数器进行自增操作。

- 无缓冲channel替代锁

```go
type counter struct {
    c chan int
    i int
}

func NewCounter() *counter {
    cter := &counter{
        c: make(chan int),
    }
    go func() {
        for {
            cter.i++
            cter.c <- cter.i
        }
    }()
    return cter
}

func (cter *counter) Increase() int {
    return <-cter.c
}

func main() {
    cter := NewCounter()
    var wg sync.WaitGroup
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func(i int) {
            v := cter.Increase()
            fmt.Printf("goroutine-%d: current counter value is %d\n", i, v)
            wg.Done()
        }(i)
    }
    wg.Wait()
}
```

这个实现中将计数器操作全部交给一个独立的Goroutine去处理，并通过无缓冲channel的同步阻塞特性，实现了计数器的控制。这样其他Goroutine通过Increase函数试图增加计数器值的动作，实质上就转化为了一次无缓冲channel的接收动作。

这种并发设计逻辑更符合Go语言所倡导的**“不要通过共享内存来通信，而是通过通信来共享内存”**的原则。

### 33.3 带缓冲channel的惯用法 🔖🔖

带缓冲的channel与无缓冲的channel的最大不同之处，就在于它的**==异步性==**。也就是说，对一个带缓冲channel，在缓冲区未满的情况下，对它进行发送操作的Goroutine不会阻塞挂起；在缓冲区有数据的情况下，对它进行接收操作的Goroutine也不会阻塞挂起。

#### 第一种用法：用作消息队列

channel的原生特性与消息队列十分相似，包括Goroutine安全、有FIFO（first-in, first out）保证等。

其实，和无缓冲channel更多用于信号/事件管道相比，可自行设置容量、异步收发的带缓冲channel更适合被用作为消息队列，并且，带缓冲channel在数据收发的性能上要明显好于无缓冲channel。

以通过对channel读写的基本测试来印证这一点 🔖

- **单接收单发送性能的基准测试**



- **多接收多发送性能基准测试**





总结：

- 无论是1收1发还是多收多发，带缓冲channel的收发性能都要好于无缓冲channel；
- 对于带缓冲channel而言，发送与接收的Goroutine数量越多，收发性能会有所下降；
- 对于带缓冲channel而言，选择适当容量会在一定程度上提升收发性能。



#### 第二种用法：用作计数信号量（counting semaphore）

Go并发设计的一个惯用法，就是将带缓冲channel用作计数信号量（counting semaphore）。

带缓冲channel中的当前数据个数代表的是，当前同时处于活动状态（处理业务）的Goroutine的数量，而带缓冲channel的容量（capacity），就代表了允许同时处于活动状态的Goroutine的最大数量。向带缓冲channel的一个发送操作表示获取一个信号量，而从channel的一个接收操作则表示释放一个信号量。



#### len(channel)的应用

**len**是Go语言的一个内置函数，它支持接收数组、切片、map、字符串和channel类型的参数，并返回对应类型的“长度”，也就是一个整型值。

针对channel ch的类型不同，len(ch)有如下两种语义：

- 当ch为无缓冲channel时，len(ch)总是返回0；
- 当ch为带缓冲channel时，len(ch)返回当前channel ch中**尚未被读取的元素个数**。



### 33.4 nil channel的妙用

如果一个channel类型变量的值为nil，称它为**nil channel**。nil channel有一个特性，那就是**对nil channel的读写都会发生阻塞**。



### 33.5 与select结合使用的一些惯用法

#### 第一种用法：利用default分支避免阻塞



#### 第二种用法：实现超时机制

带超时机制的select，是Go中常见的一种select和channel的组合用法。通过超时事件，既可以避免长期陷入某种操作的等待中，也可以做一些异常处理工作。



#### 第三种用法：实现心跳机制

结合time包的Ticker，可以实现带有心跳机制的select。这种机制让我们可以在监听channel的同时，执行一些**周期性的任务**



### 小结

Go为了原生支持并发，把channel视作一等公民身份，这就大幅提升了开发人员使用channel进行并发设计和实现的体验。

通过预定义函数make可以创建两类channel：无缓冲channel与带缓冲的channel。这两类channel具有不同的收发特性，可以适用于不同的应用场合：无缓冲channel兼具通信与同步特性，常用于作为信号通知或替代同步锁；而带缓冲channel的异步性，让它更适合用来实现基于内存的消息队列、计数信号量等。

值为nil的channel的阻塞特性，有些时候它也能帮上大忙。而面对已关闭的channel也一定要小心，尤其要避免向已关闭的channel发送数据，那会导致panic。

select是Go为了支持同时操作多个channel，而引入的另外一个并发原语，select与channel有几种常用的固定搭配。

### 思考

> 日常开发中还见过哪些实用的channel使用模式呢？



## 34 并发：如何使用共享变量？ 🔖

> Rob Pike：“不要通过共享内存来通信，应该通过通信来共享内存（Don’t communicate by sharing memory, share memory by communicating）”

Go主流风格：**使用channel进行不同Goroutine间的通信**。

不过，<u>Go也并没有彻底放弃基于共享内存的并发模型，而是在提供CSP并发模型原语的同时，还通过标准库的sync包，提供了针对传统的、基于共享内存并发模型的低级同步原语</u>，包括：互斥锁（`sync.Mutex`）、读写锁（sync.RWMutex）、条件变量（`sync.Cond`）等，并通过atomic包提供了原子操作原语等等。

### 34.1 sync包低级同步原语可以用在哪？

一般建议优先考虑CSP并发模型进行并发程序设计。下面一些场景依然需要sync包提供的低级同步原语。

- 首先是**需要高性能的临界区（critical section）同步机制场景**。

sync.Mutex和channel各自实现的临界区同步机制，做个简单的性能基准测试对比：

`sync_test.go`

```go
package main

import (
	"sync"
	"testing"
)

var cs = 0 // 模拟临界区要保护的数据
var mu sync.Mutex
var c = make(chan struct{}, 1)

func criticalSectionSyncByMutex() {
	mu.Lock()
	cs++
	mu.Unlock()
}

func criticalSectionSyncByChan() {
	c <- struct{}{}
	cs++
	<-c
}

func BenchmarkCriticalSectionSyncByMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		criticalSectionSyncByMutex()
	}
}

func BenchmarkCriticalSectionSyncByMutexInParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			criticalSectionSyncByMutex()
		}
	})
}

func BenchmarkCriticalSectionSyncByChan(b *testing.B) {
	for n := 0; n < b.N; n++ {
		criticalSectionSyncByChan()
	}
}

func BenchmarkCriticalSectionSyncByChanInParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			criticalSectionSyncByChan()
		}
	})
}
```



```sh
$ go test -bench .
goos: darwin
goarch: arm64
pkg: gofirst/ch34
BenchmarkCriticalSectionSyncByMutex-8             	88083549	        13.64 ns/op
BenchmarkCriticalSectionSyncByMutexInParallel-8   	22337848	        55.29 ns/op
BenchmarkCriticalSectionSyncByChan-8              	28172056	        42.48 ns/op
BenchmarkCriticalSectionSyncByChanInParallel-8    	 5722972	       208.1 ns/op
PASS

```

结果：无论是在单Goroutine情况下，还是在并发测试情况下，`sync.Mutex`实现的同步机制的性能，都要比channel实现的高出三倍多。

- 第二种就是**在不想转移结构体对象所有权，但又要保证结构体内部状态数据的同步访问的场景**。

基于channel的并发设计，有一个特点：在Goroutine间通过channel转移数据对象的所有权。所以，只有拥有数据对象所有权（从channel接收到该数据）的Goroutine才可以对该数据对象进行状态变更。

如果你的设计中没有转移结构体对象所有权，但又要保证结构体内部状态数据在多个Goroutine之间同步访问，那么你可以使用sync包提供的低级同步原语来实现，比如最常用的`sync.Mutex`。



### 34.2 sync包中同步原语使用的注意事项 





### 34.3 互斥锁（Mutex）还是读写锁（RWMutex）？ 

使用互斥锁的两个原则：

- **尽量减少在锁中的操作**。这可以减少其他因Goroutine阻塞而带来的损耗与延迟。
- **一定要记得调用Unlock解锁**。忘记解锁会导致程序局部死锁，甚至是整个程序死锁，会导致严重的后果。同时，我们也可以结合第23讲学习到的defer，优雅地执行解锁操作。



通常，**互斥锁（Mutex）是临时区同步原语的首选**，它常被用来对结构体对象的内部状态、缓存等进行保护，是使用最为广泛的临界区同步原语。相比之下，读写锁的应用就没那么广泛了，只活跃于它擅长的场景下。



结论：

- 并发量较小的情况下，Mutex性能最好；随着并发量增大，Mutex的竞争激烈，导致加锁和解锁性能下降；
- RWMutex的读锁性能并没有随着并发量的增大，而发生较大变化，性能始终恒定在40ns左右；
- 在并发量较大的情况下，RWMutex的写锁性能和Mutex、RWMutex读锁相比，是最差的，并且随着并发量增大，RWMutex写锁性能有继续下降趋势。

**读写锁适合应用在具有一定并发量且读多写少的场合**。



### 34.4 条件变量

`sync.Cond`是传统的条件变量原语概念在Go语言中的实现。



### 34.5 原子操作（atomic operations）

`atomic`包

原子操作（atomic operations）是相对于普通指令操作而言的。

### 小结

如果考虑使用低级同步原语，一般都是因为低级同步原语可以提供**更佳的性能表现**，性能基准测试结果告诉我们，使用低级同步原语的性能可以高出channel许多倍。在性能敏感的场景下，我们依然离不开这些低级同步原语。

在使用sync包提供的同步原语之前，一定要牢记这些原语使用的注意事项：**不要复制首次使用后的Mutex/RWMutex/Cond等**。一旦复制，你将很大可能得到意料之外的运行结果。

sync包中的低级同步原语各有各的擅长领域：

- 在具有一定并发量且读多写少的场合使用RWMutex；
- 在需要“等待某个条件成立”的场景下使用Cond；
- 当你不确定使用什么原语时，那就使用Mutex吧。

如果你对同步的性能有极致要求，且并发量较大，读多写少，那么可以考虑一下atomic包提供的原子操作函数。



### 思考

> 使用基于共享内存的并发模型时，最令人头疼的可能就是“死锁”问题的存在了。
>
> 死锁的产生条件？编写一个程序模拟一下死锁的发生？



## 35 即学即练：如何实现一个轻量级线程池？🔖

### 35.1 为什么要用到Goroutine池？

**Goroutine的开销虽然“廉价”，但也不是免费的**。

最明显的，一旦规模化后，这种非零成本也会成为瓶颈。我们以一个Goroutine分配2KB执行栈为例，100w Goroutine就是2GB的内存消耗。

其次，Goroutine从[Go 1.4版本](https://go.dev/doc/go1.4)开始采用了连续栈的方案，也就是每个Goroutine的执行栈都是一块连续内存，如果空间不足，运行时会分配一个更大的连续内存空间作为这个Goroutine的执行栈，将原栈内容拷贝到新分配的空间中来。

连续栈的方案，虽然能避免Go 1.3采用的分段栈会导致的[“hot split”问题](https://docs.google.com/document/d/1wAaf1rYoM4S4gtnPh0zOlGzWtrZFQ5suE8qr2sD8uWQ/pub)，但连续栈的原理也决定了，一旦Goroutine的执行栈发生了grow，那么即便这个Goroutine不再需要那么大的栈空间，这个Goroutine的栈空间也不会被Shrink（收缩）了，这些空间可能会处于长时间闲置的状态，直到Goroutine退出。

另外，随着Goroutine数量的增加，Go运行时进行Goroutine调度的处理器消耗，也会随之增加，成为阻碍Go应用性能提升的重要因素。



面对这样的问题，Goroutine池就是一种常见的解决方案。这个方案的核心思想是对Goroutine的重用，也就是把M个计算任务调度到N个Goroutine上，而不是为每个计算任务分配一个独享的Goroutine，从而提高计算资源的利用率。



### 35.2 workerpool的实现原理

采用完全基于channel+select的实现方案，不使用其他数据结构，也不使用sync包提供的各种同步结构

workerpool的实现主要分为三个部分：

- pool的创建与销毁；
- pool中worker（Goroutine）的管理；
- task的提交与调度。

![](images/image-20240729115211470.png)

capacity是pool的一个属性，代表整个pool中worker的最大容量。我们使用一个带缓冲的channel：active，作为worker的“计数器”，这种channel使用模式就是[33]()中的**计数信号量**。

当active channel可写时，我们就创建一个worker，用于处理用户通过Schedule函数提交的待处理的请求。当active channel满了的时候，pool就会停止worker的创建，直到某个worker因故退出，active channel又空出一个位置时，pool才会创建新的worker填补那个空位。

这张图里，我们把用户要提交给workerpool执行的请求抽象为一个Task。Task的提交与调度也很简单：Task通过Schedule函数提交到一个task channel中，已经创建的worker将从这个task channel中读取task并执行。



### 35.3 workerpool的一个最小可行实现



### 35.4 添加功能选项机制









# 实战篇：打通“最后一公里”

## 36 打稳根基：怎么实现一个TCP服务器？（上）

“最后一公里”就是**面对一个实际问题的解决思路**。

遇到一个实际问题，通常使用这个思路：

![](images/83d0018200214fd78fd1a586717dae0a.jpg)

- 首先是要理解问题。==解决实际问题的过程起始于对问题的理解==。我们要搞清楚为什么会有这个问题，问题究竟是什么。对于技术人员来说，最终目的是识别出可能要用到的技术点。

- 然后我们要对识别出的技术点，做相应的==技术预研与储备==。怎么做技术预研呢？我们至少要了解**技术诞生的背景、技术的原理、技术能解决哪些问题以及不能解决哪些问题，还有技术的优点与不足**，等等。当然，如果没有新技术点，可以忽略这一步。

- 最后，我们要基于技术预研和储备的结果，进行解决方案的设计与实现，这个是技术人最擅长的。

> 那为什么这个解决实际问题的步骤是一个循环呢？
>
> 这是由问题的难易程度，以及人的认知能力有差别所决定的。如果问题简单或人的认知能力很强，我们可以一次性解决这个实际问题；如果问题复杂或人的认知能力稍弱，那么一个循环可能无法彻底解决这个问题，我们就会再一次进入该循环，直到问题得到完美解决。



![](images/image-20240704191921768.png)

### 36.1 什么是网络编程

![](images/image-20240729120048843.png)

通常来说，我们更多关注OSI网络模型中的传输层（四层）与应用层（七层），也就是TCP/IP网络模型中的最上面两层。

TCP/IP网络模型，实现了两种传输层协议：TCP和UDP。

- TCP是面向连接的流协议，为通信的两端提供稳定可靠的数据传输服务；
- 而UDP则提供了一种无需建立连接就可以发送数据包的方法。

两种协议各有擅长的应用场景。

日常开发中使用最多的是TCP协议。基于TCP协议，我们实现了各种各样的满足用户需求的应用层协议。比如，常用的HTTP协议就是应用层协议的一种，而且是使用得最广泛的一种。而基于HTTP的Web编程就是一种**针对应用层的网络编程**。还可以**基于传输层暴露给开发者的编程接口，实现应用层的自定义应用协议**。

目前各大主流操作系统平台中，最常用的传输层暴露给用户的**网络编程接口**，就是==套接字（socket）==。**直接基于socket编程实现应用层通信业务，也是最常见的一种网络编程形式。**

### 36.2 问题描述

> 实现一个基于TCP的**自定义**应用层协议的通信服务端。

基于TCP的自定义应用层协议通常有两种常见的定义模式：

- ==二进制模式==：采用长度字段标识独立数据包的边界。采用这种方式定义的常见协议包括**MQTT**（物联网最常用的应用层协议之一）、**SMPP**（短信网关点对点接口协议）等；
- ==文本模式==：采用特定分隔符标识流中的数据包的边界，常见的包括HTTP协议等。

相比之下，二进制模式要比文本模式编码更紧凑也更高效，所以我们这个问题中的自定义协议也采用了二进制模式，协议规范内容如下图：

![](images/image-20240704192638336.png)

这个协议的通信两端的通信流程：

![](images/image-20240704192738073.png)

现在任务，就是实现支持这个协议通信的服务端。

- 首先，socket是传输层给用户提供的编程接口，要进行的网络通信绕不开socket  ->  【了解socket编程模型】
- 其次，一旦通过socket将双方的连接建立后，剩下的就是通过网络I/O操作在两端收发数据了  ->  【学习基本网络I/O操作的方法与注意事项】
- 最后，任何一端准备发送数据或收到数据后都要对数据进行操作，由于TCP是流协议  ->  【了解针对字节的操作】

在Go中，字节操作基本上就是byte切片的操作。

### 36.3 TCP Socket编程模型

TCP Socket诞生以来，它的**编程模型**，也就是**网络I/O模型**已几经演化。

网络I/O模型定义的是**应用线程与操作系统内核之间的交互行为模式**。通常用**阻塞（Blocking）**/**非阻塞（Non-Blocking）**来描述网络I/O模型。

常用的网络I/O模型:

- 阻塞I/O(Blocking I/O)

![](images/image-20240704192907447.png)

在**阻塞I/O模型**下，当用户空间应用线程，向操作系统内核发起I/O请求后（一般为操作系统提供的I/O系列系统调用），内核会尝试执行这个I/O操作，并等所有数据就绪后，将数据从内核空间拷贝到用户空间，最后系统调用从内核空间返回。而在这个期间内，用户空间应用线程将阻塞在这个I/O系统调用上，无法进行后续处理，只能等待。

因此，在这样的模型下，**一个线程仅能处理一个网络连接上的数据通信**。即便连接上没有数据，线程也只能阻塞在对Socket的读操作上（以等待对端的数据）。

虽然这个模型对应用整体来说是低效的，但对开发人员来说，这个模型却是最容易实现和使用的，所以，各大平台在默认情况下都将Socket设置为阻塞的。

- 非阻塞I/O（Non-Blocking I/O）

![](images/image-20240704192928299.png)

和阻塞I/O模型正相反，在**非阻塞模型**下，当用户空间线程向操作系统内核发起I/O请求后，内核会执行这个I/O操作，如果这个时候数据尚未就绪，就会立即将“未就绪”的状态以错误码形式（比如：EAGAIN/EWOULDBLOCK），返回给这次I/O系统调用的发起者。而后者就会根据系统调用的返回状态来决定下一步该怎么做。

在非阻塞模型下，位于用户空间的I/O请求发起者通常会通过==轮询==的方式，去一次次发起I/O请求，直到读到所需的数据为止。不过，这样的轮询是**对CPU计算资源的极大浪费**，因此，非阻塞I/O模型单独应用于实际生产的比例并不高。

- I/O多路复用（I/O Multiplexing）

为了避免非阻塞I/O模型轮询对计算资源的浪费，同时也考虑到阻塞I/O模型的低效，**开发人员首选的网络I/O模型**，逐渐变成了建立在内核提供的多路复用函数`select`/`poll`等（以及性能更好的`epoll`等函数）基础上的**I/O多路复用模型**。

![](images/image-20240704192946825.png)

在这种模型下，应用线程首先将需要进行I/O操作的Socket，都添加到多路复用函数中（这里以`select`为例），然后阻塞，等待select系统调用返回。

当内核发现有数据到达时，对应的Socket具备了通信条件，这时select函数返回。然后用户线程会针对这个Socket再次发起网络I/O请求，比如一个`read`操作。由于数据已就绪，这次网络I/O操作将得到预期的操作结果。



相比于阻塞模型一个线程只能处理一个Socket的低效，I/O多路复用模型中，一个应用线程可以同时处理多个Socket。同时，I/O多路复用模型由内核实现可读/可写事件的通知，避免了非阻塞模型中轮询，带来的CPU计算资源浪费的问题。

目前，主流网络服务器采用的都是“I/O多路复用”模型，有的也结合了多线程。不过，**I/O多路复用**模型在支持更多连接、提升I/O操作效率的同时，也给使用者带来了不小的复杂度，以至于后面出现了许多**高性能的I/O多路复用框架**，比如：[libevent](http://libevent.org/)、[libev](http://software.schmorp.de/pkg/libev.html)、[libuv](https://github.com/libuv/libuv)等，以帮助开发者简化开发复杂性，降低心智负担。

### 36.4 Go语言socket编程模型

阻塞I/O模型是对开发人员最友好的，也是心智负担最低的模型，而**I/O多路复用**的这种**通过回调割裂执行流**的模型，对开发人员来说还是过于复杂了，于是Go选择了为开发人员提供**阻塞I/O模型**，Gopher只需在Goroutine中以最简单、最易用的**“阻塞I/O模型”**的方式，进行Socket操作就可以了。

再加上，Go没有使用基于线程的并发模型，而是使用了开销更小的Goroutine作为基本执行单元，这让每个Goroutine处理一个TCP连接成为可能，并且在高并发下依旧表现出色。

不过，网络I/O操作都是系统调用，Goroutine执行I/O操作的话，一旦阻塞在系统调用上，就会导致M也被阻塞，为了解决这个问题，Go设计者将这个“复杂性”隐藏在Go运行时中，他们在运行时中实现了==网络轮询器（netpoller)==，netpoller的作用，就是只阻塞执行网络I/O操作的Goroutine，但不阻塞执行Goroutine的线程（也就是M）。

这样一来，对于Go程序的用户层（相对于Go运行时层）来说，它眼中看到的goroutine采用了“阻塞I/O模型”进行网络I/O操作，Socket都是“阻塞”的。

但实际上，这样的“假象”，是通过Go运行时中的netpoller **I/O多路复用机制**，“模拟”出来的，对应的、真实的底层操作系统Socket，实际上是非阻塞的。只是运行时拦截了针对底层Socket的系统调用返回的错误码，并通过**netpoller**和Goroutine调度，让Goroutine“阻塞”在用户层所看到的Socket描述符上。

比如：当用户层针对某个Socket描述符发起`read`操作时，如果这个Socket对应的连接上还没有数据，运行时就会将这个Socket描述符加入到netpoller中监听，同时发起此次读操作的Goroutine会被挂起。

直到Go运行时收到这个Socket数据可读的通知，Go运行时才会重新唤醒等待在这个Socket上准备读数据的那个Goroutine。而这个过程，从Goroutine的视角来看，就像是read操作一直阻塞在那个Socket描述符上一样。

而且，Go语言在网络轮询器（netpoller）中采用了I/O多路复用的模型。考虑到最常见的多路复用系统调用select有比较多的限制，比如：**监听Socket的数量有上限（1024）、时间复杂度高，**等等，Go运行时选择了在不同操作系统上，使用操作系统各自实现的高性能多路复用函数，比如：Linux上的`epoll`、Windows上的`iocp`、FreeBSD/MacOS上的`kqueue`、Solaris上的`event port`等，这样可以最大程度提高netpoller的调度和执行性能。



### 36.5 socket监听（listen）与接收连接（accept）

**socket编程的核心在于服务端**，而服务端有着自己一套相对固定的套路：==Listen+Accept==。在这套固定套路的基础上，服务端程序通常采用一个Goroutine处理一个连接，它的大致结构如下：

```go
func handleConn(c net.Conn) {
     defer c.Close()
     for {
         // read from the connection
         // ... ...
         // write to the connection
         //... ...
     }
 }
 
 func main() {
     l, err := net.Listen("tcp", ":8888")
     if err != nil {
         fmt.Println("listen error:", err)
         return
     }
 
     for {
         c, err := l.Accept()
         if err != nil {
             fmt.Println("accept error:", err)
             break
         }
         // start a new goroutine to handle
         // the new connection.
         go handleConn(c)
     }
 }
```

在第12行使用了net包的Listen函数绑定（bind）服务器端口8888，并将它转换为监听状态，Listen返回成功后，这个服务会进入一个循环，并调用net.Listener的Accept方法接收新客户端连接。

在没有新连接的时候，这个服务会阻塞在Accept调用上，直到有客户端连接上来，Accept方法将返回一个net.Conn实例。通过这个`net.Conn`，我们可以和新连上的客户端进行通信。这个服务程序启动了一个新Goroutine，并将net.Conn传给这个Goroutine，这样这个Goroutine就专职负责处理与这个客户端的通信了。

而net.Listen函数很少报错，除非是监听的端口已经被占用。

```sh
netstat -an | grep 8888 
```



### 36.6 向服务端建立TCP连接

一旦服务端按照上面的`Listen + Accept`结构成功启动，客户端便可以使用`net.Dial`或`net.DialTimeout`向服务端发起连接建立的请求：

> ==Dial==**是一个源自传统通信术语的隐喻，意指**主动发起网络连接的行为，类似于电话拨号（Dial）的动作。这一术语在编程中广泛用于表示客户端主动向服务端建立通信通道的过程。

```go
conn, err := net.Dial("tcp", "localhost:8888")
conn, err := net.DialTimeout("tcp", "localhost:8888", 2 * time.Second)
```

Dial函数向服务端发起TCP连接，这个函数会一直阻塞，直到连接成功或失败后，才会返回。而DialTimeout带有超时机制，如果连接耗时大于超时时间，这个函数会返回超时错误。 对于客户端来说，连接的建立还可能会遇到几种特殊情形。

- 第一种情况：网络不可达或对方服务未启动。

如果传给`Dial`的服务端地址是网络不可达的，或者服务地址中端口对应的服务并没有启动，端口未被监听（Listen），`Dial`几乎会立即返回类似这样的错误：

```go
dial error: dial tcp :8888: getsockopt: connection refused
```

- 第二种情况：对方服务的listen backlog队列满。

当对方服务器很忙，瞬间有大量客户端尝试向服务端建立连接时，服务端可能会出现listen backlog队列满，接收连接（accept）不及时的情况，这就会导致客户端的`Dial`调用阻塞，直到服务端进行一次accept，从backlog队列中腾出一个槽位，客户端的Dial才会返回成功。

而且，不同操作系统下backlog队列的长度是不同的，在macOS下，这个默认值如下：

```sh
$ sysctl -a | grep kern.ipc.somaxconn
kern.ipc.somaxconn: 128
```

🔖

- 第三种情况：若网络延迟较大，Dial将阻塞并超时。

如果网络延迟较大，TCP连接的建立过程（三次握手）将更加艰难坎坷，会经历各种丢包，时间消耗自然也会更长，这种情况下，`Dial`函数会阻塞。如果经过长时间阻塞后依旧无法建立连接，那么`Dial`也会返回类似`getsockopt: operation timed out`的错误。

在连接建立阶段，多数情况下`Dial`是可以满足需求的，即便是阻塞一小会儿也没事。但对于那些需要有严格的连接时间限定的Go应用，如果一定时间内没能成功建立连接，程序可能会需要执行一段“错误”处理逻辑，所以，这种情况下，我们使用`DialTimeout`函数更适合。

### 36.7 全双工通信

一旦客户端调用Dial成功，就在客户端与服务端之间建立起了一条全双工的通信通道。通信双方通过各自获得的Socket，可以在向对方发送数据包的同时，接收来自对方的数据包。下图展示了系统层面对这条全双工通信通道的实现原理：

![](images/image-20240729154153246.png)

任何一方的操作系统，都会为已建立的连接分配一个**发送缓冲区**和一个**接收缓冲区**。

以客户端为例，客户端会通过成功连接服务端后得到的conn（封装了底层的socket）向服务端发送数据包。这些数据包会先进入到己方的发送缓冲区中，之后，这些数据会被操作系统内核通过网络设备和链路，发到服务端的接收缓冲区中，服务端程序再通过代表客户端连接的conn读取服务端接收缓冲区中的数据，并处理。

反之，服务端发向客户端的数据包也是先后经过服务端的发送缓冲区、客户端的接收缓冲区，最终到达客户端的应用的。

### 36.8 Socket读操作

连接建立起来后，就要在连接上进行读写以完成业务逻辑。

Go运行时隐藏了**I/O多路复用**的复杂性。Go语言使用者只需采用**Goroutine+阻塞I/O模型**，就可以满足大部分场景需求。Dial连接成功后，会返回一个net.Conn接口类型的变量值，这个接口变量的底层类型为一个`*TCPConn`：

```go
//$GOROOT/src/net/tcpsock.go
type TCPConn struct {
    conn
}
```

TCPConn内嵌了一个非导出类型：`conn`（封装了底层的socket），因此，TCPConn“继承”了`conn`类型的`Read`和`Write`方法，后续通过`Dial`函数返回值调用的`Read`和`Write`方法都是net.conn的方法，它们分别代表了对socket的读和写。

通过几个场景来总结Go中从socket读取数据的行为特点：

- 首先是Socket中无数据的场景。

  连接建立后，如果客户端未发送数据，服务端会阻塞在Socket的读操作上，这和前面提到的“阻塞I/O模型”的行为模式是一致的。执行该这个操作的Goroutine也会被挂起。Go运行时会监视这个Socket，直到它有数据读事件，才会重新调度这个Socket对应的Goroutine完成读操作。

- 第二种情况是Socket中有部分数据。

  如果Socket中有部分数据就绪，且数据数量小于一次读操作期望读出的数据长度，那么读操作将会成功读出这部分数据，并返回，而不是等待期望长度数据全部读取后，再返回。

  举个例子，服务端创建一个长度为10的切片作为接收数据的缓冲区，等待Read操作将读取的数据放入切片。当客户端在已经建立成功的连接上，成功写入两个字节的数据（比如：hi）后，服务端的Read方法将成功读取数据，并返回`n=2，err=nil`，而不是等收满10个字节后才返回。

- 第三种情况是Socket中有足够数据。

  如果连接上有数据，且数据长度大于等于一次`Read`操作期望读出的数据长度，那么`Read`将会成功读出这部分数据，并返回。这个情景是最符合我们对`Read`的期待的了。

  以上面的例子为例，当客户端在已经建立成功的连接上，成功写入15个字节的数据后，服务端进行第一次`Read`时，会用连接上的数据将我们传入的切片缓冲区（长度为10）填满后返回：`n = 10, err = nil`。这个时候，内核缓冲区中还剩5个字节数据，当服务端再次调用`Read`方法时，就会把剩余数据全部读出。

- 最后一种情况是设置读操作超时。

有些场合，对socket的读操作的阻塞时间有严格限制的，但由于Go使用的是阻塞I/O模型，如果没有可读数据，Read操作会一直阻塞在对Socket的读操作上。

这时，可以通过net.Conn提供的SetReadDeadline方法，设置读操作的超时时间，当超时后仍然没有数据可读的情况下，Read操作会解除阻塞并返回超时错误，这就给Read方法的调用者提供了进行其他业务处理逻辑的机会。

SetReadDeadline方法接受一个绝对时间作为超时的deadline。一旦通过这个方法设置了某个socket的Read deadline，当发生超时后，如果我们不重新设置Deadline，那么后面与这个socket有关的所有读操作，都会返回超时失败错误。

结合SetReadDeadline设置的服务端一般处理逻辑：

```go
func handleConn(c net.Conn) {
    defer c.Close()
    for {
        // read from the connection
        var buf = make([]byte, 128)
        c.SetReadDeadline(time.Now().Add(time.Second))
        n, err := c.Read(buf)
        if err != nil {
            log.Printf("conn read %d bytes,  error: %s", n, err)
            if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
                // 进行其他业务逻辑的处理
                continue
            }
            return
        }
        log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
    }
}
```

使用SetReadDeadline（time.Time{}）实现取消超时设置。



### 36.9 Socket写操作

通过net.Conn实例的Write方法，可以将数据写入Socket。当Write调用的返回值n的值，与预期要写入的数据长度相等，且err = nil时，就执行了一次成功的Socket写操作，这是在调用Write时遇到的最常见的情形。

和Socket的读操作一些特殊情形相比，Socket写操作遇到的特殊情形同样不少：

#### 第一种情况：写阻塞

TCP协议通信两方的操作系统内核，都会为这个连接保留数据缓冲区，调用Write向Socket写入数据，实际上是将数据写入到操作系统协议栈的数据缓冲区中。TCP是全双工通信，因此每个方向都有独立的数据缓冲。当发送方将对方的接收缓冲区，以及自身的发送缓冲区都写满后，再调用Write方法就会出现阻塞的情况。

例子，客户端：

```go
func main() {
    log.Println("begin dial...")
    conn, err := net.Dial("tcp", ":8888")
    if err != nil {
        log.Println("dial error:", err)
        return
    }
    defer conn.Close()
    log.Println("dial ok")

    data := make([]byte, 65536)
    var total int
    for {
        n, err := conn.Write(data)
        if err != nil {
            total += n
            log.Printf("write %d bytes, error:%s\n", n, err)
            break
        }
        total += n
        log.Printf("write %d bytes this time, %d bytes in total\n", n, total)
    }

    log.Printf("write %d bytes in total\n", total)
}
```

客户端每次调用Write方法向服务端写入65536个字节，并在Write方法返回后，输出此次Write的写入字节数和程序启动后写入的总字节数量。

服务端的处理程序逻辑摘录：

```go
... ...
func handleConn(c net.Conn) {
    defer c.Close()
    time.Sleep(time.Second * 10)
    for {
        // read from the connection
        time.Sleep(5 * time.Second)
        var buf = make([]byte, 60000)
        log.Println("start to read from conn")
        n, err := c.Read(buf)
        if err != nil {
            log.Printf("conn read %d bytes,  error: %s", n, err)
            if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
                continue
            }
        }

        log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
    }
}
... ...
```

务端在前10秒中并不读取数据，因此当客户端一直调用Write方法写入数据时，写到一定量后就会发生阻塞。

🔖

后续当服务端每隔5秒进行一次读操作后，内核socket缓冲区腾出了空间，客户端就又可以写入了:



#### 第二种情况：写入部分数据





#### 第三种情况：写入超时





### 36.10 并发Socket读写







### 36.11 Socket关闭

通常情况下，当客户端需要断开与服务端的连接时，客户端会调用net.Conn的Close方法关闭与服务端通信的Socket。

如果客户端主动关闭了Socket，那么服务端的`Read`调用将会读到什么呢？

分“有数据关闭”和“无数据关闭”两种情况。

- “==有数据关闭==”是指在客户端关闭连接（Socket）时，Socket中还有服务端尚未读取的数据。在这种情况下，服务端的Read会成功将剩余数据读取出来，最后一次Read操作将得到`io.EOF`错误码，表示客户端已经断开了连接。
- 如果是在“==无数据关闭==”情形下，服务端调用的Read方法将直接返回`io.EOF`。

不过因为Socket是全双工的，客户端关闭Socket后，如果服务端Socket尚未关闭，这个时候服务端向Socket的写入操作依然可能会成功，因为数据会成功写入己方的内核socket缓冲区中，即便最终发不到对方socket缓冲区也会这样。因此，当发现对方socket关闭后，己方应该正确合理处理自己的socket，再继续write已经没有任何意义了。



## 37 代码操练：怎么实现一个TCP服务器？（中）

### 37.1 建立对协议的抽象

程序是对现实世界的抽象。对于现实世界的自定义应用协议规范，我们需要在程序世界建立起对这份**协议的抽象**。

#### 深入协议字段

一个高度简化的、基于二进制模式定义的协议。

二进制模式定义的特点，就是采用长度字段标识独立数据包的边界。

在这个协议规范中：

请求包和应答包的第一个字段（totalLength）都是包的总长度，它就是用来标识包边界的那个字段，也是在应用层用于“分割包”的最重要字段。

请求包与应答包的第二个字段也一样，都是commandID，这个字段用于标识包类型

- 连接请求包（值为0x01）
- 消息请求包（值为0x02）
- 连接响应包（值为0x81）
- 消息响应包（值为0x82）

```go
const (
    CommandConn   = iota + 0x01 // 0x01，连接请求包
    CommandSubmit               // 0x02，消息请求包
)

const (
    CommandConnAck   = iota + 0x81 // 0x81，连接请求的响应包
    CommandSubmitAck               // 0x82，消息请求的响应包
)
```

请求包与应答包的第三个字段都是ID，ID是每个连接上请求包的消息流水号，顺序累加，步长为1，循环使用，多用来请求发送方后匹配响应包，所以要求一对请求与响应消息的流水号必须相同。

请求包与响应包唯一的不同之处，就在于最后一个字段：请求包定义了有效载荷（payload），这个字段承载了应用层需要的业务数据；而响应包则定义了请求包的响应状态字段（result），这里其实简化了响应状态字段的取值，成功的响应用0表示，如果是失败的响应，无论失败原因是什么，我们都用1来表示。

明确了应用层协议的各个字段定义之后，我们接下来就看看如何建立起对这个协议的抽象。

#### 建立Frame和Packet抽象

TCP连接上的数据是一个没有边界的字节流，但在业务层眼中，没有字节流，只有各种协议消息。因此，无论是从客户端到服务端，还是从服务端到客户端，业务层在连接上看到的都应该是一个挨着一个的协议消息流。

建立第一个抽象：**Frame**。每个Frame表示一个协议消息，这样在业务层眼中，连接上的字节流就是由一个接着一个Frame组成的，如下图所示：

![](images/image-20250121200532054.png)

自定义协议就封装在这一个个的Frame中。协议规定了将Frame分割开来的方法，那就是利用每个Frame开始处的totalLength，每个Frame由一个totalLength和Frame的负载（payload）构成。

![](images/image-20250121200632964.png)

这样通过Frame header: totalLength就可以将Frame之间隔离开来。

建立协议的第二个抽象：**Packet**。将Frame payload定义为一个Packet。上图右侧展示的就是Packet的结构。

Packet就是业务层真正需要的消息，每个Packet由Packet头和Packet Body部分组成。Packet头就是commandID，用于标识这个消息的类型；而ID和payload（packet payload）或result字段组成了Packet的Body部分，对业务层有价值的数据都包含在Packet Body部分。



### 37.2 协议的解包与打包

所谓协议的**解包（decode）**，就是指识别TCP连接上的字节流，将一组字节“转换”成一个特定类型的协议消息结构，然后这个消息结构会被业务处理逻辑使用。

而**打包（encode）**刚刚好相反，是指将一个特定类型的消息结构转换为一组字节，然后这组字节数据会被放在连接上发送出去。

具体到我们这个自定义协议上，解包就是指`字节流 -> Frame`，打包是指`Frame -> 字节流`。针对这个协议的服务端解包与打包的流程图：

![](images/image-20250121200929681.png)

TCP流数据先后经过frame decode和packet decode，得到应用层所需的packet数据，而业务层回复的响应，则先后经过packet的encode与frame的encode，写入TCP数据流中。

#### Frame的实现

协议部分最重要的两个抽象是Frame和Packet，建立frame包与packet包。

frame包的职责是提供识别TCP流边界的编解码器，可以很容易为这样的编解码器，定义出一个统一的接口类型StreamFrameCodec：

```go
// tcp-server-demo1/frame/frame.go

type FramePayload []byte

type StreamFrameCodec interface {
    Encode(io.Writer, FramePayload) error   // data -> frame，并写入io.Writer
    Decode(io.Reader) (FramePayload, error) // 从io.Reader中提取frame payload，并返回给上层
}
```



#### Packet的实现

和Frame不同，Packet有多种类型（这里只定义了Conn、submit、connack、submit ack)。

抽象这些类型需要遵循的共同接口：

```go
// tcp-server-demo1/packet/packet.go

type Packet interface {
    Decode([]byte) error     // []byte -> struct
    Encode() ([]byte, error) //  struct -> []byte
}
```

其中，Decode是将一段字节流数据解码为一个Packet类型，可能是conn，可能是submit等，具体要根据解码出来的commandID判断。

而Encode则是将一个Packet类型编码为一段字节流数据。

🔖

### 37.3 服务端的组装



### 37.4 验证测试





## 38 成果优化：怎么实现一个TCP服务器？（下）

### 38.1 Go程序优化的基本套路

Go程序的优化，也有着固定的套路可循:

![](images/image-20240704193459259.png)

- 首先我们要建立性能基准。要想对程序实施优化，我们首先要有一个初始“参照物”，这样我们才能在执行优化措施后，检验优化措施是否有效，所以这是优化循环的第一步。
- 第二步是性能剖析。要想优化程序，我们首先要找到可能影响程序性能的“瓶颈点”，这一步的任务，就是通过各种工具和方法找到这些“瓶颈点”。
- 第三步是代码优化。我们要针对上一步找到的“瓶颈点”进行分析，找出它们成为瓶颈的原因，并有针对性地实施优化。
- 第四步是与基准比较，确定优化效果。这一步，我们会采集优化后的程序的性能数据，与第一步的性能基准进行比较，看执行上述的优化措施后，是否提升了程序的性能。

### 38.2 建立性能基准

建立性能基准的方式大概有两种，

- 一种是通过**编写Go原生提供的性能基准测试（benchmark test）用例**来实现，这相当于对程序的局部热点建立性能基准，常用于一些算法或数据结构的实现，比如分布式全局唯一ID生成算法、树的插入/查找等。

- 另外一种是**基于度量指标为程序建立起图形化的性能基准**，这种方式适合针对程序的整体建立性能基准。而我们的自定义协议服务端程序就十分适合用这种方式，接下来我们就来看一下基于度量指标建立基准的一种可行方案。

#### 建立观测设施

这些年，基于**Web的可视化工具、开源监控系统以及时序数据库**的兴起，给我们建立性能基准带来了很大的便利，业界有比较多成熟的工具组合可以直接使用。但业界最常用的还是**Prometheus+Grafana**的组合。

以Docker为代表的轻量级容器（container）的兴起，让这些工具的部署、安装都变得十分简单。使用docker-compose工具，基于容器安装Prometheus+Grafana的组合。

🔖

#### 配置Grafana



#### 在服务端埋入度量数据采集点

哪些度量数据能反映出服务端的性能指标呢？

- 当前已连接的客户端数量（client_connected）；
- 每秒接收消息请求的数量（req_recv_rate）；
- 每秒发送消息响应的数量（rsp_send_rate）。



#### 第一版性能基准



### 38.3 尝试用pprof剖析

Go内置了对Go代码进行性能剖析的工具：**pprof**。



### 38.4 代码优化

#### 带缓存的网络I/O



#### 重用内存对象



# 泛型篇



## 39 Go泛型诞生过程

[Go 1.18 Beta2版本](https://go.dev/blog/go1.18beta2)

### 1 为什么要加入泛型？

[维基百科-泛型编程](https://en.wikipedia.org/wiki/Generic_programming)，最初泛型编程概念的文章中给了解释：“**泛型编程的中心思想是对具体的、高效的算法进行抽象，以获得通用的算法，然后这些算法可以与不同的数据表示法结合起来，产生各种各样有用的软件**”。

将**算法与类型解耦**



在没有泛型的情况下，需要针对不同类型重复实现相同的算法逻辑。

对于简单的、诸如上面这样的加法函数还可忍受，但对于复杂的算法，比如涉及复杂排序、查找、树、图等算法，以及一些容器类型（链表、栈、队列等）的实现时，缺少了泛型的支持还真是麻烦。

在没有泛型之前，Gopher们通常使用空接口类型`interface{}`，作为算法操作的对象的数据类型，不过这样做的不足之处也很明显：**一是无法进行类型安全检查，二是性能有损失**。



### 2 Go泛型设计的简史

[“泛型窘境”](https://research.swtch.com/generic) 2019 Russ Cox提出了Go泛型实现的三个可遵循的方法，以及每种方法的不足，也就是三个slow（拖慢）：

- **拖慢程序员**：不实现泛型，不会引入复杂性，但就像前面例子中那样，需要程序员花费精力重复实现AddInt、AddInt64等；
- **拖慢编译器**：就像C++的泛型实现方案那样，通过增加编译器负担为每个类型实例生成一份单独的泛型函数的实现，这种方案产生了大量的代码，其中大部分是多余的，有时候还需要一个好的链接器来消除重复的拷贝；
- **拖慢执行性能**：就像Java的泛型实现方案那样，通过隐式的装箱和拆箱操作消除类型差异，虽然节省了空间，但代码执行效率低。

在当时，三个slow之间需要取舍，就如同数据一致性的CAP原则一样，无法将三个slow同时消除。

[“Why Generics?”](https://go.dev/blog/why-generics)

[《Featherweight Go》](https://arxiv.org/abs/2005.11710)

[《The Next Step for Generics》](https://go.dev/blog/generics-next-step)

...

最后，在2021年12月14日，[Go 1.18 beta1版本发布](https://go.dev/blog/go1.18beta1)，这个版本包含了对Go泛型的正式支持。

### 3 Go泛型的基本语法

Go泛型是Go开源以来在语法层面的最大一次变动。

[Go泛型的最后一版技术提案](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)

#### 1️⃣ 类型参数（type parameter）

==类型参数==是在函数声明、方法声明的receiver部分或类型定义的类型参数列表中，声明的（非限定）类型名称。类型参数在声明中充当了一个**未知类型的占位符（placeholder）**，在泛型函数或泛型类型实例化时，类型参数会被一个类型实参替换。

普通函数的参数列表：

```go
func Foo(x, y aType, z anotherType)
```

这里，x, y, z是形参（parameter）的名字，也就是变量，而aType，anotherType是形参的类型，也就是类型。

泛型函数的类型参数（type parameter）列表：

```go
func GenericFoo[P aConstraint, Q anotherConstraint](x,y P, z Q)
```

这里，P、Q是类型形参的名字，也就是类型。aConstraint，anotherConstraint代表类型参数的==约束（constraint）==，可以理解为对类型参数可选值的一种限定。

相较而言多出一个组成部分：==类型参数列表==。

类型参数列表位于函数名与函数参数列表之间，通过一个**方括号**括起，不支持变长类型参数。而且，类型参数列表中声明的类型参数，可以作为函数普通参数列表中的形参类型。

P、Q的类型什么时候才能确定呢？等到泛型函数**具化（instantiation）**时才能确定。另外，按惯例，类型参数（type parameter）的名字都是首字母大写的，通常都是用单个大写字母命名。

#### 2️⃣ 约束（constraint）

约束（constraint）规定了一个类型实参（type argument）必须满足的条件要求。如果某个类型满足了某个约束规定的所有条件要求，那么它就是这个约束修饰的类型形参的一个合法的类型实参。

```go
type C1 interface {
	~int | ~int32
	M1()
}

type T struct{}

func (T) M1() {
}

type T1 int

func (T1) M1() {
}

func foo[P C1](t P) {

}

func main() {
	var t1 T1
	foo(t1)
	var t T
	foo(t) // 编译器报错： Cannot use T as the type C1. Type does not implement constraint 'C1' because type is not included in type set ('~int', '~int32')
}
```

C1是定义的约束，它声明了一个方法M1，以及两个可用作类型实参的类型(~int | ~int32)，类型列表中的多个类型实参类型用“|”分隔。

还定义了两个自定义类型T和T1，两个类型都实现了M1方法，但T类型的底层类型为struct{}，而T1类型的底层类型为int，这样就导致了虽然T类型满足了约束C1的方法集合，但类型T因为底层类型并不是int或int32而不满足约束C1，这也就会导致`foo(t)`调用在编译阶段报错。

建议：**做约束的接口类型与做传统接口的接口类型最好要分开定义**，除非约束类型真的既需要方法集合，也需要类型列表。

#### 3️⃣ 类型具化（instantiation）

```go
func Sort[Elem interface{ Less(y Elem) bool }](list []Elem) {
}

type book struct {
}

func (x book) Less(y book) bool {
	return true
}

func main() {
	var bookshelf []book
	Sort[book](bookshelf) // 泛型函数调用
}
```

上面的泛型函数调用`Sort[book](bookhelf)`会分成两个阶段：

第一个阶段就是==具化（instantiation）==。

形象点说，**具化（instantiation）就好比一家生产“排序机器”的工厂根据要排序的对象的类型，将这样的机器生产出来的过程**。我们继续举前面的例子来分析一下，整个具化过程如下：

1. 工厂接单：**Sort[book]**，发现要排序的对象类型为book；
2. 模具检查与匹配：检查book类型是否满足模具的约束要求（也就是是否实现了约束定义中的Less方法）。如果满足，就将其作为类型实参替换Sort函数中的类型形参，结果为**Sort[book]**，如果不满足，编译器就会报错；
3. 生产机器：将泛型函数Sort具化为一个**新函数**，这里我们把它起名为**booksort**，其函数原型为**func([]book)**。本质上**booksort := Sort[book]**。

第二阶段是==调用（invocation）==。

一旦“排序机器”被生产出来，那么它就可以对目标对象进行排序了，这和普通的函数调用没有区别。这里就相当于调用booksort（bookshelf），整个过程只需要检查传入的函数实参（bookshelf）的类型与booksort函数原型中的形参类型（[]book）是否匹配就可以了。

伪代码来表述上面两个过程：

```plain
Sort[book](bookshelf)

<=>

具化：booksort := Sort[book]
调用：booksort(bookshelf)
```

简化，调用Sort不需要传入类型实参book，和普通函数调用那样，Go编译器会根据传入的实参变量，进行实参类型参数的自动推导（Argument type inference）：

```go
Sort(bookshelf)
```

#### 4️⃣ 泛型类型

除了函数可以携带类型参数变身为“泛型函数”外，类型也可以拥有类型参数而化身为“泛型类型”，如定义一个向量泛型类型：

```go
type Vector[T any] []T
```

这是一个带有类型参数的类型定义，类型参数位于类型名的后面，同样用方括号括起。在类型定义体中可以引用类型参数列表中的参数名（比如T）。类型参数同样拥有自己的约束，如上面代码中的**any**。在Go 1.18中，any是interface{}的别名，也是一个预定义标识符，使用any作为类型参数的约束，代表没有任何约束。

使用泛型类型，我们也要遵循先具化，再使用的顺序，比如下面例子：

```go
type Vector[T any] []T

func (v Vector[T]) Dump() {
    fmt.Printf("%#v\n", v)
}

func main() {
    var iv = Vector[int]{1,2,3,4}
    var sv Vector[string]
    sv = []string{"a","b", "c", "d"}
    iv.Dump()
    sv.Dump()
}
```

在这段代码中，在使用Vector[T]之前都显式用类型实参对泛型类型进行了具化，从而得到具化后的类型Vector[int]和Vector[string]。 Vector[int]的底层类型为[]int，Vector[string]的底层类型为[]string。然后我们再对具化后的类型进行操作。



### 4 Go泛型的性能🔖



### 5 Go泛型的使用建议🔖

Go核心团队最担心的就是“泛型被滥用”，所以Go核心团队在各种演讲场合都在努力地告诉大家Go泛型的适用场景以及应该如何使用。

#### 什么情况适合使用泛型



#### 什么情况不宜使用泛型





## 40 类型参数

Go的泛型与其他主流编程语言的泛型不同点（[不支持的若干特性](https://github.com/golang/proposal/blob/master/design/43651-type-parameters.md#omissions)）：

- **不支持泛型特化（specialization）**，即不支持编写一个泛型函数针对某个具体类型的特殊版本；
- **不支持元编程（metaprogramming）**，即不支持编写在编译时执行的代码来生成在运行时执行的代码；
- **不支持操作符方法（operator method）**，即只能用普通的方法（method）操作类型实例（比如：getIndex(k)），而不能将操作符视为方法并自定义其实现，比如一个容器类型的下标访问c[k]；
- **不支持变长的类型参数（type parameters）**；
- …

### 40.1 例子：返回切片中值最大的元素

```go
import "fmt"

func maxAny(sl []any) any {
	if len(sl) == 0 {
		panic("slice is empty")
	}

	max := sl[0]
	for _, v := range sl[1:] {
		switch v.(type) {
		case int:
			if v.(int) > max.(int) {
				max = v
			}
		case string:
			if v.(string) > max.(string) {
				max = v
			}
		case float64:
			if v.(float64) > max.(float64) {
				max = v
			}
		}
	}
	return max
}

func main() {
	i := maxAny([]any{1, 2, -4, -6, 7, 0})
	m := i.(int)
	fmt.Println(m)                                                 // 输出：7
	fmt.Println(maxAny([]any{"11", "22", "44", "66", "77", "10"})) // 输出：77
	fmt.Println(maxAny([]any{1.01, 2.02, 3.03, 5.05, 7.07, 0.01})) // 输出：7.07
}
```

```go
// max_test.go
func BenchmarkMaxInt(b *testing.B) {
    sl := []int{1, 2, 3, 4, 7, 8, 9, 0}
    for i := 0; i < b.N; i++ {
        maxInt(sl)
    }
}

func BenchmarkMaxAny(b *testing.B) {
    sl := []any{1, 2, 3, 4, 7, 8, 9, 0}
    for i := 0; i < b.N; i++ {
        maxAny(sl)
    }
}
```

```shell
$ go test -v -bench . ./max_test.go max_any.go max_int.go
goos: darwin
goarch: arm64
BenchmarkMaxInt
BenchmarkMaxInt-8       325038272                3.453 ns/op
BenchmarkMaxAny
BenchmarkMaxAny-8       100000000               10.66 ns/op
PASS
ok      command-line-arguments  2.960s

```

看到，基于any(interface{})实现的maxAny其执行性能要比像maxInt这样的函数慢上数倍。



```go
// max_generics.go
type ordered interface {
	~int | ~int32 | ~int16 | ~int64 | ~int8 |
		~float32 | ~float64 |
		~string |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func maxGenerics[T ordered](sl []T) T {
	if len(sl) == 0 {
		panic("slice is empty")
	}

	max := sl[0]
	for _, v := range sl {
		if v > max {
			max = v
		}
	}
	return max
}

type myString string

func main() {
	var m int = maxGenerics([]int{1, 2, -4, -6, 7, 0})
	fmt.Println(m)                                                           // 输出：7
	fmt.Println(maxGenerics([]string{"11", "22", "44", "66", "77", "10"}))   // 输出：77
	fmt.Println(maxGenerics([]float64{1.01, 2.02, 3.03, 5.05, 7.07, 0.01}))  // 输出：7.07
	fmt.Println(maxGenerics([]int8{1, 2, -4, -6, 7, 0}))                     // 输出：7
	fmt.Println(maxGenerics([]myString{"11", "22", "44", "66", "77", "10"})) // 输出：77
}
```

```shell
$ go test -v -bench . ./max_test.go max_any.go max_int.go max_generics.go
goos: darwin
goarch: arm64
BenchmarkMaxInt
BenchmarkMaxInt-8               324042938                3.448 ns/op
BenchmarkMaxAny
BenchmarkMaxAny-8               100000000               10.72 ns/op
BenchmarkMaxGenerics
BenchmarkMaxGenerics-8          306493706                3.769 ns/op
PASS
ok      command-line-arguments  4.529s
```



### 40.2 类型参数（type parameters）

**Go泛型方案的实质是对类型参数（type parameter）的支持**，包括：

- 泛型函数（generic function）：带有类型参数的函数；
- 泛型类型（generic type）：带有类型参数的自定义类型；
- 泛型方法（generic method）：泛型类型的方法。

#### 1️⃣ 泛型函数

![](images/46bf5dc6778b44abacb7c3e9d3aa57d9.jpg)

##### 调用泛型函数

![](images/e19fcdaa14d3442da3aa8b5d77069c9e.jpg)

##### 泛型函数实例化（instantiation）

![](images/bf8571f4e2ce48a782411aaf0e5cf022.jpg)

#### 2️⃣ 泛型类型

![](images/bd4292c1f4ec4288ab8a385e51a7811b.jpg)

##### 使用泛型类型



##### 泛型方法

Go类型可以拥有自己的方法（method），泛型类型也不例外，为泛型类型定义的方法称为**泛型方法（generic method）**。







## 41 定义泛型约束

虽然泛型是开发人员表达“通用代码”的一种重要方式，但这并不意味着所有泛型代码对所有类型都适用。更多的时候，我们需要对泛型函数的类型参数以及泛型函数中的实现代码**设置限制**。泛型函数调用者只能传递满足限制条件的类型实参，泛型函数内部也只能以类型参数允许的方式使用这些类型实参值。在Go泛型语法中，我们使用**类型参数约束**（type parameter constraint）（以下简称**约束**）来表达这种限制条件。

约束之于类型参数就好比函数参数列表中的类型之于参数：

![](images/image-20240721151712110.png)

### 41.1 最宽松的约束：any



### 41.2 支持比较操作的内置约束：comparable



### 41.3 自定义约束

Go泛型最终决定使用interface语法来定义约束。这样一来，**凡是接口类型均可作为类型参数的约束**。



![](images/b31aba8af0c9439c8a530128ae976c3f.jpg)



```go
type Ia interface {
	int | string  // 仅代表int和string
}

type Ib interface {
	~int | ~string  // 代表以int和string为底层类型的所有类型
}
```



![](images/91186776da4d4c6dac2e12153ff7b082.jpg)



### 41.4 类型集合（type set）

![](images/image-20240721152001159.png)



![](images/image-20240721152035052.png)

### 41.5 简化版的约束形式

```go
type I interface { // 独立于泛型函数外面定义
    ~int | ~string
}

func doSomething1[T I](t T)
func doSomething2[T interface{~int | ~string}](t T) // 以接口类型字面值作为约束
```

```go
func doSomething2[T ~int | ~string](t T) // 简化版的约束形式
```



一般形式来表述：

```go
func doSomething[T interface {T1 | T2 | ... | Tn}](t T)

等价于下面简化版的约束形式：

func doSomething[T T1 | T2 | ... | Tn](t T) 
```



### 41.6 约束的类型推断



## 42 明确使用时机🔖

Go语言开发人员都有义务去正确、适当的使用泛型，而不是滥用或利用泛型炫技。

### 42.1 何时适合使用泛型？

#### 场景一：编写通用数据结构时



#### 场景二：函数操作的是Go原生的容器类型时



#### 场景三：不同类型实现一些方法的逻辑相同时



### 42.2 Go泛型实现原理简介

#### Stenciling方案

![](images/bf7591aebdbd4731aaa2097cc89efbb6.jpg)

#### Dictionaries方案

![](images/a19c9ff543a249d488a915c0ab056fab.jpg)



#### Go最终采用的方案：GC Shape Stenciling方案

![](images/cf7fb42e6c9a4b099ea40ebd4577f06d.jpg)

```go
// gcshape.go
func f[T any](t T) T {
    var zero T
    return zero
}

type MyInt int

func main() {
    f[int](5)
    f[MyInt](15)
    f[int64](6)
    f[uint64](7)
    f[int32](8)
    f[rune](18)
    f[uint32](9)
    f[float64](3.14)
    f[string]("golang")

    var a int = 5
    f[*int](&a)
    var b int32 = 15
    f[*int32](&b)
    var c float64 = 8.88
    f[*float64](&c)
    var s string = "hello"
    f[*string](&s)
}
```





![](images/1a0d0ac6ef0842668e5355742a69375c.jpg)

### 42.3 泛型对执行效率的影响



# 补充

## 43 如何拉取私有的Go Module？🔖

### 43.1 导入本地module

彻底抛弃Gopath构建模式，全面拥抱Go Module构建模式。

> **如果我们的项目依赖的是本地正在开发、尚未发布到公共站点上的Go Module，那么我们应该如何做呢？**

Go 1.18之前:**go.mod的replace指示符**



### 43.2 拉取私有module的需求与参考方案

![](images/06be396d9cce45cf8db37119c429cf35.jpg)







![](images/a6583e4e450b48678796b6f1a6b2bba5.jpg)

- 第一个方案是通过直连组织公司内部的私有Go Module服务器拉取。

![](images/75d9ee7bfd434a96be8d6ff3e034af77.jpg)



- 第二种方案，是将外部Go Module与私有Go Module都交给内部统一的GOPROXY服务去处理：

![](images/7d834252238e4101adc7f49be4ec8fb5.jpg)





### 43.3 统一Goproxy方案的实现思路与步骤

![](images/10de4af0a27c4df9b577542da3ebfdca.jpg)

#### 选择一个GOPROXY实现



#### 自定义包导入路径并将其映射到内部的vcs仓库

![](images/image-20241218112944766.png)

#### 开发机(客户端)的设置



#### 方案的“不足”

- 第一点：开发者还是需要额外配置GONOSUMDB变量。



- 第二点：新增私有Go Module，vanity.yaml需要手工同步更新。



- 第三点：无法划分权限。







## 44 作为Go Module的作者，你应该知道的几件事

从Go Module的作者或维护者的视角，看看在规划、发布和维护Go Module时需要考虑和注意什么事情，包括go项目**仓库布局、Go Module的发布、升级module主版本号、作废特定版本的module**，等等。

### 44.1 仓库布局：是单module还是多module

如果没有单一仓库（monorepo）的强约束，那么在默认情况下，选择**一个仓库管理一个module**是不会错的，这是管理Go Module的最简单的方式，也是最常用的标准方式。这种方式下，module维护者维护起来会很方便，module的使用者在引用module下面的包时，也可以很容易地确定包的导入路径。

🌰例子，在github.com/bigwhite/srsm这个仓库下管理着一个Go Module（==srsm==是single repo single module的缩写）。

通常情况下，module path与仓库地址保持一致，都是github.com/bigwhite/srsm，这点会体现在go.mod中：

```go
// go.mod
module github.com/bigwhite/srsm

go 1.22.1
```

然后对仓库打tag，这个tag也会成为Go Module的版本号，这样，对仓库的版本管理其实就是对Go Module的版本管理。

如果仓库布局如下：

```sh
./srsm
├── go.mod
├── go.sum
├── pkg1/
│   └── pkg1.go
└── pkg2/
    └── pkg2.go
```

那么这个module的使用者可以很轻松地确定pkg1和pkg2两个包的导入路径，一个是`github.com/bigwhite/srsm/pkg1`，另一个则是`github.com/bigwhite/srsm/pkg2`。

如果module演进到了v2.x.x版本，那么以pkg1包为例，它的包的导入路径就变成了`github.com/bigwhite/srsm/v2/pkg1`。



如果组织层面要求采用单一仓库（monorepo）模式，也就是**所有Go Module都必须放在一个repo下**，那只能使用单repo下管理多个Go Module的方法了。

> Go Module的设计者Russ Cox：“在单repo多module的布局下，添加module、删除module，以及对module进行版本管理，都需要相当谨慎和深思熟虑，因此，管理一个单module的版本库，几乎总是比管理现有版本库中的多个module要容易和简单”。

🌰单repo多module的例子，假设repo地址是`github.com/bigwhite/srmm`。这个repo下的结构布局如下（srmm是single repo multiple modules的缩写）：

```sh
./srmm
├── module1
│   ├── go.mod
│   └── pkg1
│       └── pkg1.go
└── module2
    ├── go.mod
    └── pkg2
        └── pkg2.go
```

srmm仓库下面有两个Go Module，分为位于子目录module1和module2的下面，这两个目录也是各自module的根目录（module root）。这种情况下，module的path也不能随意指定，必须包含子目录的名字。

以module1为例分析，它的path是`github.com/bigwhite/srmm/module1`，只有这样，Go命令才能根据用户导入包的路径，找到对应的仓库地址和在仓库中的相对位置。同理，module1下的包名同样是以module path为前缀的，比如：`github.com/bigwhite/srmm/module1/pkg1`。

在单仓库多module模式下，各个module的版本是独立维护的。因此，我们在通过打tag方式发布某个module版本时，tag的名字必须包含子目录名。比如：如果我们要发布module1的v1.0.0版本，我们不能通过给仓库打v1.0.0这个tag号来发布module1的v1.0.0版本，**正确的作法应该是打`module1/v1.0.0`这个tag号**。

现在可能觉得这样理解起来也没有多复杂，但当各个module的主版本号升级时，你就会感受到这种方式带来的繁琐了。



### 44.2 发布Go Module

发布的步骤也十分简单，就是**为repo打上tag并推送到代码服务器上**就好了。



### 44.3 作废特定版本的Go Module🔖

![](images/23931514c6ea70debaeb9e5709cec20a.jpg)

![](images/85fyy48dd62570758dd21666a4646631.jpg)

![](images/737eb3121d9fa70387742de128eba500.jpg)





#### 修复broken版本并重新发布

![](images/3aead734389a42e9b1faa6835a048121.jpg)

#### 发布module的新patch版本

![](images/1bbdb1fdad664081aa1be1ab3da1dff9.jpg)



### 44.4 升级module的major版本号🔖



- 第一种情况：repo下的所有module统一进行版本发布。



- 第二个情况：repo下的module各自独立进行版本发布。

![](images/image-20250104155812081.png)







## 45 Go指针

### 45.1 什么是指针类型

**如果我们拥有一个类型T，那么以T作为基类型的指针类型为*T**。

`unsafe.Pointer`类似于C语言中的`void*`，用于表示一个通用指针类型，也就是**任何指针类型都可以显式转换为一个unsafe.Pointer，而unsafe.Pointer也可以显式转换为任意指针类型**，如下面代码所示：

```go
var p *T
var p1 = unsafe.Pointer(p) // 任意指针类型显式转换为unsafe.Pointer
p = (*T)(p1)               // unsafe.Pointer也可以显式转换为任意指针类型
```

unsafe.Pointer是Go语言的高级特性，在**Go运行时与Go标准库**中unsafe.Pointer都有着广泛的应用。但unsafe.Pointer属于unsafe编程范畴，这里就不深入了。

如果指针类型变量没有被显式赋予初值，那么它的值为**nil**：

```go
var p *T
println(p == nil) // true
```

给一个指针类型变量赋值：

```go
var a int = 13
var p *int = &a  // 给整型指针变量p赋初值
```

`&a`作为*int指针类型变量p的初值，`&`符号称为**取地址符号**，这一行的含义就是将变量a的地址赋值给指针变量p。

只能使用基类型变量的地址给对应的指针类型变量赋值，如果类型不匹配，Go编译器是会报错的。

以最简单的整型变量为例，看看对应的内存单元存储的内容：

![](images/image-20241218142759552.png)

看到，**对于非指针类型变量，Go在对应的内存单元中放置的就是该变量的值**。我们对这些变量进行修改操作的结果，也会直接体现在这个内存单元上：

![](images/image-20241218142850668.png)



以*int类型指针变量为例，查看指针类型变量在对应的内存空间：

![](images/image-20241218143002693.png)

**Go为指针变量p分配的内存单元中存储的是整型变量a对应的内存单元的地址**。也正是由于指针类型变量存储的是内存单元的地址，指针类型变量的大小与其基类型大小无关，而是和系统地址的表示长度有关。

```go
package main

import "unsafe"

type foo struct {
    id   string
    age  int8
    addr string
}

func main() {
    var p1 *int
    var p2 *bool
    var p3 *byte
    var p4 *[20]int
    var p5 *foo
    var p6 unsafe.Pointer
    println(unsafe.Sizeof(p1)) // 8 
    println(unsafe.Sizeof(p2)) // 8
    println(unsafe.Sizeof(p3)) // 8
    println(unsafe.Sizeof(p4)) // 8
    println(unsafe.Sizeof(p5)) // 8
    println(unsafe.Sizeof(p6)) // 8
}
```

unsafe包的Sizeof函数原型：

```go
func Sizeof(x ArbitraryType) uintptr
```

`uintptr`，是一个Go预定义的标识符，**uintptr是一个整数类型，它的大小足以容纳任何指针的比特模式（bit pattern）**，换句话说就是：**在Go语言中uintptr类型的大小就代表了指针类型的大小**。

> `go doc [package路径]`
>
> ```
> go doc fmt
> go doc fmt.Printf
> go doc builtin.uintptr
> go doc unsafe 
> ```

一旦指针变量得到了正确赋值，也就是指针指向某一个合法类型的变量，我们就可以通过指针读取或修改其指向的内存单元所代表的基类型变量，比如：

```go
var a int = 17
var p *int = &a
println(*p) // 17 
(*p) += 3
println(a)  // 20
```

![](images/image-20250103192048205.png)

通过指针变量读取或修改其指向的内存地址上的变量值，这个操作被称为指针的**解引用（dereference）**，就是在指针类型变量的前面加上一个星号。

要输出指针自身的值，也就是指向的内存单元的地址，可以使用Printf通过%p来实现：

```go
fmt.Printf("%p\n", p) // 0x14000122018
```



指针变量可以变换其指向的内存单元:

```go
var a int = 5
var b int = 6

var p *int = &a  // 指向变量a所在内存单元
println(*p)      // 输出变量a的值
p = &b           // 指向变量b所在内存单元
println(*p)      // 输出变量b的值
```

多个指针变量可以指向同一个变量的内存单元的，这样通过其中一个指针变量对内存单元的修改，是可以通过另外一个指针变量的解引用反映出来的:

```go
var a int = 5
var p1 *int = &a // p1指向变量a所在内存单元
var p2 *int = &a // p2指向变量b所在内存单元
(*p1) += 5       // 通过p1修改变量a的值
println(*p2)     // 10 对变量a的修改可以通过另外一个指针变量p2的解引用反映出来
```

### 45.2 二级指针

```go
func main() {
	var a int = 5
	var p1 *int = &a
	println(*p1)
	var b int = 55
	var p2 *int = &b
	println(&p2)

	var pp **int = &p1
	println(**pp) // 5
	pp = &p2
	println(**pp) // 55
}
```

![](images/image-20241218144756444.png)

对一级指针解引用，得到的其实是指针指向的变量。而对二级指针pp解引用一次，得到将是pp指向的指针变量：

```go
println((*pp) == p1) // true
```

对pp解引用两次，其实就相当于对一级指针解引用一次，得到的是pp指向的指针变量所指向的整型变量：

```go
println((**pp) == (*p1)) // true
println((**pp) == a)     // true
```

一级指针常被用来改变普通变量的值，**二级指针就可以用来改变指针变量的值，也就是指针变量的指向**。



前面提到过，在同一个函数中，改变指针的指向十分容易，只需要给一级指针重新赋值为另外一个变量的地址就可以了。

但是，如果我们需要跨函数改变一个指针变量的指向，我们就不能选择一级指针类型作为形参类型了。因为一级指针只能改变普通变量的值，无法改变指针变量的指向。我们只能选择二级指针类型作为形参类型。

```go
package main

func f(pp **int) {
	var b int = 55
	var p1 *int = &b
	(*pp) = p1
}
func main() {
	var a int = 5
	var p *int = &a
	println(*p) // 5
	f(&p)
	println(*p) // 55
}
```

![](images/image-20250103193513824.png)

### 45.3 Go中的指针用途与使用限制

Go是带有垃圾回收的编程语言，指针在Go中依旧位于C位，它的作用不仅体现在==语法层面==上，更体现在Go==运行时层面==，尤其是==内存管理==与==垃圾回收==这两个地方，这两个运行时机制**只关心指针**。

在语法层面，相对于“指针为王”的C语言来说，Go指针的使用要少不少，这很大程度上是因为Go提供了更灵活和高级的复合类型，比如切片、map等，并将使用指针的复杂性隐藏在运行时的实现层面了。这样，Go程序员自己就不需要在语法层面通过指针来实现这些高级复合类型的功能。

指针无论是在Go中，还是在其他支持指针的编程语言中，存在的意义就是为了是**“==可改变==”**。在Go中，我们使用`*T`类型的变量调用方法、以`*T`类型作为函数或方法的形式参数、返回`*T`类型的返回值等的目的，也都是因为指针可以改变其指向的内存单元的值。

当然，指针的好处，还包括它传递的开销是**常数级**的（在x86-64平台上仅仅是8字节的拷贝），**可控可预测**。无论指针指向的是一个字节大小的变量，还是一个拥有10000个元素的[10000]int型数组，传递指针的开销都是一样的。

不过，虽然Go在语法层面上保留了指针，但Go语言的目标之一是成为一门**安全**的编程语言，因此，它对指针的使用做了一定的限制，包括这两方面：

#### 限制一：限制了显式指针类型转换。

在C语言中，可以实现显式指针类型转换：

```c
#include <stdio.h>
  
int main() {
    int a = 0x12345678;
    int *p = &a;
    char *p1 = (char*)p; // 将一个整型指针显式转换为一个char型指针
    printf("%x\n", *p1); 
}
```

但是在Go中，这样的显式指针转换会得到Go编译器的报错信息：

```go
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var a int = 0x12345678
    var pa *int = &a
    var pb *byte = (*byte)(pa) // 编译器报错：cannot convert pa (variable of type *int) to type *byte
    fmt.Printf("%x\n", *pb)
}
```

如果非要进行这个转换，Go也提供了unsafe的方式：

```go
func main() {                                                                         
    var a int = 0x12345678                                                            
    var pa *int = &a                                                                  
    var pb *byte = (*byte)(unsafe.Pointer(pa)) // ok
    fmt.Printf("%x\n", *pb) // 78                                                          
} 
```



#### 限制二：不支持指针运算。

指针运算是C语言的大杀器。

但指针运算也是安全问题的“滋生地”。为了安全性，Go在语法层面抛弃了指针运算这个特性。

如果非要做指针运算，Go依然提供了unsafe的途径。



## 46 Go语言的GC实现

依赖人去处理复杂的对象内存管理的问题是不科学、不合理的。C 和 C++ 程序员已经被折磨了数十年，我们不应该再重蹈覆辙了，于是，后来的很多编程语言就用上垃圾回收（GC）机制。

### 46.1 GC拯救程序员

垃圾回收（Garbage Collection）也被称为自动内存管理技术，在现代编程语言中使用得相当广泛，常见的 Java、Go、C# 均在语言的 runtime 中集成了相应的实现。

在传统的不带GC的编程语言中，我们需要关注对象的分配位置，要自己去选择对象是分配在堆还是栈上，但在 Go 这门有 GC 的语言中，集成了逃逸分析功能来帮助我们自动判断对象应该在堆上还是栈上，我们可以使用 `go build -gcflags="-m"` 来观察逃逸分析的结果：

```go
package main

func main() {
    var m = make([]int, 10240)
    println(m[0])
}
```

可以看到，较大的对象也会被放在堆上。

```shell
$ go build -gcflags="-m" escape.go 
# command-line-arguments
./escape.go:3:6: can inline main
./escape.go:4:14: make([]int, 10240) escapes to heap
```

若对象被分配在栈上，它的管理成本就比较低，通过挪动栈顶寄存器就可以实现对象的分配和释放。若对象被分配在堆上，就要经历层层的内存申请过程。但这些流程对用户都是透明的，在编写代码时我们并不需要在意它。只有需要优化时，我们才需要研究具体的逃逸分析规则。

逃逸分析与垃圾回收结合在一起，极大地解放了程序员们的心智，我们在编写代码时，似乎再也没必要去担心内存的分配和释放问题了。

**然而，一切抽象皆有成本，这个成本要么花在编译期，要么花在运行期。**

GC 这种方案是选择在运行期来解决问题，不过在极端场景下GC本身引起的问题依然是令人难以忽视的：

![](images/512b9dbc7e904e7d99ccbbcfe46a5547.jpg)

这张图的场景是在内存中缓存了上亿的 kv，这时 GC 使用的 CPU 甚至占到了总 CPU 占用的 90% 以上。简单粗暴地在内存中缓存对象，到头来发现 GC 成为了 CPU 杀手，吃掉了大量的服务器资源，这显然不是我们期望的结果。

想要正确地分析原因，就需要我们对 GC 本身的实现机制有稍微深入一些的理解。

### 46.2 内存管理的三个参与者

- ==mutator==指的是我们的应用，也就是application，我们将堆上的对象看作一个图，跳出应用来看的话，应用的代码就是在不停地修改这张堆对象图里的指向关系。下面的图可以帮我们理解 mutator 对堆上的对象的影响：

![](images/1aa2af74b3784bbc9dfff4d08efe697f.jpg)

- ==allocator==指的是==内存分配器==，应用需要内存的时候都要向allocator申请。allocator要维护好内存分配的数据结构，在多线程场景下工作的内存分配器还需要考虑高并发场景下锁的影响，并针对性地进行设计以降低锁冲突。

- ==collector==是==垃圾回收器==。死掉的堆对象、不用的堆内存都要由 collector 回收，最终归还给操作系统。当 GC 扫描流程开始执行时，collector 需要扫描内存中存活的堆对象，扫描完成后，未被扫描到的对象就是无法访问的堆上垃圾，需要将其占用内存回收掉。

三者的交互过程：

![](images/96723c68fa2a4b3b9f37a87e5835a650.jpg)

应用需要在堆上申请内存时，会由编译器帮程序员自动调用runtime.newobject，这时 allocator 会使用 mmap 这个系统调用从操作系统中申请内存，若 allocator 发现之前申请的内存还有富余，会从本地预先分配的数据结构中划分出一块内存，并把它以指针的形式返回给应用。在内存分配的过程中，allocator 要负责维护内存管理对应的数据结构。

而collector 要扫描的就是 allocator 管理的这些数据结构，应用不再使用的部分便应该被回收，通过 madvise 这个系统调用返还给操作系统。

下面看看这些交互的细节。

### 46.3 分配内存

应用程序使用 mmap 向 OS 申请内存，操作系统提供的接口比较简单，mmap 返回的结果是连续的内存区域。

mutator 申请内存是以应用视角来看问题。比如说，我需要的是某一个 struct和某一个 slice 对应的内存，这与从操作系统中获取内存的接口之间还有一个鸿沟。这就需要由 allocator 进行映射与转换，将以“块”来看待的内存与以“对象”来看待的内存进行映射：

![](images/a2fd04efa8ac43acb4b1eb2c63285708.jpg)

从上面这张图看到，在应用的视角看，我们需要初始化的 a 是一个 1024000 长度的 int 切片；在内存管理的视角来看，我们需要管理的只是 start、offset 对应的一段内存。

在现代 CPU 上，除了内存分配的正确性以外，还要考虑**分配过程的效率问题**，应用执行期间小对象会不断地生成与销毁，如果每一次对象的分配与释放都需要与操作系统交互，那么成本是很高的。这就需要我们在应用层设计好内存分配的多级缓存，尽量减少小对象高频创建与销毁时的锁竞争，这个问题在传统的 C/C++ 语言中已经有了解法，那就是 `tcmalloc`：

![](images/77498d454319407ca377af4c3a0fcd8c.jpg)

tcmalloc 通过维护一套**多级缓存结构**，降低了应用内存分配过程中对全局锁的使用频率，使小对象的内存分配做到了**尽量无锁**。

Go 语言的内存分配器基本是 tcmalloc 的 1:1 搬运……毕竟都是 Google 的项目。

在 Go 语言中，根据对象中是否有指针以及对象的大小，将内存分配过程分为三类：

- `tiny` ：size < 16 bytes && has no pointer(noscan)；
- `small` ：has pointer(scan) || (size >= 16 bytes && size <= 32 KB)；
- `large` ：size > 32 KB。

一个个分析。在内存分配过程中，最复杂的就是 tiny 类型的分配。

我们可以将内存分配的路径与 CPU 的多级缓存作类比，这里 mcache 内部的 tiny 可以类比为 L1 cache，而 alloc 数组中的元素可以类比为 L2 cache，全局的 mheap.mcentral 结构为 L3 cache，mheap.arenas 是 L4，L4 是以页为单位将内存向下派发的，由 pageAlloc 来管理 arena 中的空闲内存。具体你可以看下这张表：

![](images/7ad31d70b606476984ab469060eb3662.jpg)

如果 L4 也没法满足我们的内存分配需求，那我们就需要向操作系统去要内存了。

和 tiny 的四级分配路径相比，small 类型的内存没有本地的 mcache.tiny 缓存，其余的与 tiny 分配路径完全一致：

![image-20250104111500117](images/image-20250104111500117.png)

large 内存分配稍微特殊一些，没有前面这两类这样复杂的缓存流程，而是直接从 mheap.arenas 中要内存，直接走 `pageAlloc` **页分配器**。

页分配器在 Go 语言中迭代了多个版本，从简单的 freelist 结构，到 treap 结构，再到现在最新版本的 radix 结构，它的查找时间复杂度也从 O(N) -> O(log(n)) -> O(1)。

在当前版本中，我们只需要知道常数时间复杂度就可以确定空闲页组成的 radix tree 是否能够满足内存分配需求。若不满足，则要对 arena 继续进行切分，或向操作系统申请更多的 arena。

只看这些分类文字不太好理解，接下来看看 **arenas、page、mspan、alloc** 这些概念是怎么关联在一起组成 Go 的内存分配流程的。

#### 内存分配的数据结构之间的关系

==arenas== 是 Go 向操作系统申请内存时的最小单位，每个 arena 为 64MB 大小，在内存中可以部分连续，但整体是个稀疏结构。

单个 arena 会被切分成以 ==8KB== 为单位的 ==page==，由 page allocator 管理，一个或多个 page 可以组成一个 mspan，每个 mspan 可以按照 sizeclass 再划分成多个 element。同样大小的 mspan 又分为 scan 和 noscan 两种，分别对应内部有指针的 object 和内部没有指针的 object。

之前讲到的四级分配结构如下图：

![](images/acd5348c531a4eb6a3221937005af37c.jpg)

从上图清晰地看到内存分配的多级路径，我们可以再研究一下这里面的 mspan。每一个 mspan 都有一个 allocBits 结构，从 mspan 里分配 element 时，我们只要将 mspan 中对应该 element 位置的 bit 位置一就可以了，其实就是将 mspan 对应 allocBits 中的对应 bit 位置一。每一个 mspan 都会对应一个 allocBits 结构，如下图：

![](images/dc9612e35e114af9a75c52eb8c02bd1c.jpg)

当然，在代码中还有一些位操作优化（如freeIndex、allocCache），之后深入。

### 46.4 垃圾回收

Go 语言使用了**并发标记与清扫算法**作为它的 GC 实现。

标记、清扫算法是一种古老的 GC 算法，是指将内存中正在使用的对象进行标记，之后清扫掉那些未被标记的对象的一种垃圾回收算法。并发标记与清扫重点在**并发**，是指垃圾回收的**标记和清扫过程能够与应用代码并发执行**。但并发标记清扫算法的一大缺陷是无法解决==内存碎片==问题，而 tcmalloc 恰好一定程度上缓解了内存碎片问题，两者配合使用相得益彰。

**但这并不是说 tcmalloc 完全没有内存碎片，不信你可以在代码里搜搜 max waste**。

#### 垃圾分类

进行垃圾回收之前，我们要先对内存垃圾进行分类，主要可以分为==语义垃圾==和==语法垃圾==两类，但并不是所有垃圾都可以被垃圾回收器回收。

**语义垃圾（semantic garbage）**，有些场景也被称为内存泄露，指的是从语法上可达（可以通过局部、全局变量被引用）的对象，但从语义上来讲他们是垃圾，垃圾回收器对此无能为力。

看一个语义垃圾在 Go 语言中的实例：

![](images/939324bb372140f6bc545bcf0e25719b.jpg)

初始化了一个 slice，元素均为指针，每个指针都指向了堆上 10MB 大小的一个对象。

![](images/8ecbf5602460400ea192ca2d5c198df1.jpg)

当这个 slice 缩容时，底层数组的后两个元素已经无法再访问了，但它关联的堆上内存依然是无法释放的。

碰到类似的场景，你可能需要在缩容前，**先将数组元素置为 nil**。

另外一种内存垃圾就是**语法垃圾（syntactic garbage）**，讲的是那些从语法上无法到达的对象，这些才是垃圾收集器主要的收集目标。

一个简单的例子来理解一下语法垃圾：

![图片](images/84a15bf6bc1644b486b60805f25a41a6.jpg)

在 allocOnHeap 返回后，堆上的 a 无法访问，便成为了语法垃圾。

#### GC 流程

Go 的每一轮版本迭代几乎都会对 GC 做优化。经过多次优化后，较新的 GC 流程如下图：

![](images/5e3aedd8e3174aa39dc434a2a0515f65.jpg)

可以看到，在并发标记开始前和并发标记终止时，有两个短暂的 stw，该 stw 可以使用 pprof 的 pauseNs 来观测，也可以直接采集到监控系统中：

![img](images/47bfe57965784cdfb0da16964ae58727.jpg)

监控系统中的 PauseNs 就是每次 stw 的时长。尽管官方声称 Go 的 stw 已经是亚毫秒级了，但我们在高压力的系统中仍然能够看到毫秒级的 stw。

#### 标记流程

Go 语言使用**三色抽象**作为其并发标记的实现。所以这里我们首先要理解三种颜色的抽象：

- 黑表示已经扫描完毕，子节点扫描完毕（gcmarkbits = 1，且在队列外）；
- 灰表示已经扫描完毕，子节点未扫描完毕（gcmarkbits = 1, 在队列内）；
- 白表示未扫描，collector 不知道任何相关信息。

使用三色抽象，主要是为了能让垃圾回收流程与应用流程并发执行，这样将对象扫描过程拆分为多个阶段，不需要一次性完成整个扫描流程。

![](images/03ce3a85020e446b86aaa957b8de0ba6.jpg)

GC 扫描的起点是根对象，忽略掉那些不重要的（finalizer 相关的先省略），常见的根对象可以参见下图：

![](images/35b61ac17230472e8f7325b95bebd9cc.jpg)

所以在 Go 语言中，从根开始扫描的含义是从 .bss 段，.data 段以及 goroutine 的栈开始扫描，最终遍历整个堆上的对象树。

标记过程是一个广度优先的遍历过程。它是扫描节点，将节点的子节点推到任务队列中，然后递归扫描子节点的子节点，直到所有工作队列都被排空为止。

![](images/79dbd155a3b54f39af42a25109cbcdd7.jpg)

标记过程会将白色对象标记，并推进队列中变成灰色对象。scanobject 的具体过程：

![img](images/02f9cf9f4ca2487d8be6bf0faf10cade.jpg)

在标记过程中，gc mark worker 会一边从工作队列（gcw）中弹出对象，一边把它的子对象 push 到工作队列（gcw）中，如果工作队列满了，则要将一部分元素向全局队列转移。

我们知道，堆上对象本质上是图，会存储引用关系互相交叉的时候，在标记过程中也有简单的剪枝逻辑：

![](images/810f1f46c3f04b719b221663fe2d5d8c.jpg)

这里，D 是 A 和 B 的共同子节点，在标记过程中自然会减枝，防止重复标记浪费计算资源：

![图片](images/2eab2ae6f0484073b7d972ca9c7fce18.jpg)

如果多个后台 mark worker 确实产生了并发，标记时使用的是 atomic.Or8，也是并发安全的：

![图片](images/4c2f35d1b1444d6e88b3d18369e4c59e.jpg)

#### 协助标记

当应用分配内存过快时，后台的 mark worker 无法及时完成标记工作，这时应用本身需要进行堆内存分配时，会判断是否需要适当协助 GC 的标记过程，防止应用因为分配过快发生 OOM。

碰到这种情况时，我们会在火焰图中看到对应的协助标记的调用栈：

![](images/d6c6d3860022494392fcb8751be36036.jpg)

不过，协助标记会对应用的响应延迟产生影响，我们可以尝试降低应用的对象分配数量进行优化。Go 内部具体是通过一套记账还账系统来实现协助标记的流程的，这一部分深入可以去看看[这里](https://github.com/golang/go/blob/11b28e7e98bce0d92d8b49c6d222fb66858994ff/src/runtime/mgcmark.go#L407) 。

#### 对象丢失问题

前面我们提到了GC线程/协程与应用线程/协程是并发执行的，在 GC 标记 worker 工作期间，应用还会不断地修改堆上对象的引用关系，这就可能导致对象丢失问题。下面是一个典型的应用与 GC 同时执行时，由于应用对指针的变更导致对象漏标记，从而被 GC 误回收的情况。

![](images/839d36450ae2440687142ff371ba11d5.jpg)

在这张图表现的 GC 标记过程中，应用动态地修改了 A 和 C 的指针，让 A 对象的内部指针指向了 B，C 的内部指针指向了 D。如果标记过程垃圾收集器无法感知到这种变化，最终 B 对象在标记完成后是白色，会被错误地认作内存垃圾被回收。

为了解决漏标，错标的问题，我们先需要定义“**三色不变性**”，如果我们的堆上对象的引用关系不管怎么修改，都能满足三色不变性，那么也不会发生对象丢失问题。三色不变性可以分为强三色不变性和弱三色不变性两种，

首先是强三色不变性（strong tricolor invariant），禁止黑色对象指向白色对象：

![图片](images/cd51fa8f975c4c5bb2bfb7ac5a80aa7f.jpg)

然后是弱三色不变性（weak tricolor invariant），黑色对象可以指向白色对象，但指向的白色对象，必须有能从灰色对象可达的路径：

![图片](images/f2882375ffed431c867863e4807f7bee.jpg)

无论应用在与 GC 并发执行期间如何修改堆上对象的关系，只要修改之后，堆上对象能**满足任意一种不变性**，就不会发生对象的丢失问题。

而实现强/弱三色不变性均需要引入屏障技术。在 Go 语言中，使用写屏障，也就是 write barrier 来解决上述问题。

#### write barrier

这里barrier 的本质是 : snippet of code insert before pointer modify。不过，**在并发编程领域也有 memory barrier，但这个含义与 GC 领域的barrier是完全不同的**，在阅读相关材料时，你一定要注意不要混淆这两个概念。

Go 语言的 GC 只有 write barrier，没有 read barrier。

在应用进入 GC 标记阶段前的 stw 阶段，会将全局变量 runtime.writeBarrier.enabled 修改为 true，这时所有的堆上指针修改操作在修改之前便会额外调用 runtime.gcWriteBarrier：

![图片](images/376bf160b2ab424b9afc97b2f3cc3413.jpg)

在反汇编结果中，我们可以通过行数找到原始的代码位置：

![图片](images/1ffa2401a2804cc9ad3393addac1ad39.jpg)

在GC领域中，常见的 write barrier 有两种：

- Dijistra Insertion Barrier，指针修改时，指向的新对象要标灰：-![](images/a57811bdaba541c9a7d0c19fed26d484.jpg)

- Yuasa Deletion Barrier，指针修改时，修改前指向的对象要标灰：-

​	![](images/70622d4b3e3f432fad9748bc8126a691.jpg)

从理论上来讲，如果 Go 语言的所有对象都在堆上，使用上述两种屏障的任意一种，都不会发生对象丢失的问题。

但我们不要忽略，在 Go 语言中，还有很多对象被分配在栈上。栈上的对象操作极其频繁，给栈上对象增加写屏障成本很高，所以 Go 是不给栈上对象开启屏障的。

只对堆上对象开启写屏障的话，使用上述两种屏障其中的任意一种，都需要在 stw 阶段对栈进行重扫。所以经过多个版本的迭代，现在 Go 的写屏障混合了上述两种屏障，实现是这样的：

![图片](images/efda40c5e1b34809a0a020ae0df40143.jpg)

这和 Go 语言在混合屏障的 proposal 上的实现不太相符，本来 proposal 是这么写的：

![图片](images/1e93ed330d074a22a1eeac1c34627151.jpg)

为什么会有这种差异呢？这主要是因为栈的颜色判断成本是很高的，官方最终还是选择了更为简单的实现，即指针断开的老对象和新对象都标灰的实现。

我们再来详细地看看前面两种屏障的对象丢失问题。

- Dijistra Insertion Barrier 的对象丢失问题：-

![图片](images/b31c628cab2d4236a260a38873e0297b.jpg)

- Yuasa Deletion Barrier 的对象丢失问题：-

![图片](images/a167877dd4ae43fc9f73a3d155e5fe3c.jpg)

早期 Go 只使用了 Dijistra 屏障，但因为会有上述对象丢失问题，需要在第二个 stw 周期进行栈重扫（stack rescan）。当 goroutine 数量较多时，stw 时间会变得很长。

但单独使用任意一种 barrier ，又没法满足 Go 消除栈重扫的要求，所以最新版本中 Go 的混合屏障其实是 Dijistra Insertion Barrier + Yuasa Deletion Barrier。

![图片](images/18de4fea78f7409b8c99ea579896d21e.jpg)

混合 write barrier 会将两个指针推到 p 的 wbBuf 结构去，看看这个过程：

![图片](images/9df9fa4fedb046efb3e3aa5637aab902.jpg)

mutator 和后台的 mark worker 在并发执行时的完整过程了：

![图片](images/fc5260fe14504985818eac2baf15e952.jpg)

#### 回收流程

相比复杂的标记流程，对象的回收和内存释放就简单多了。

进程启动时会有两个特殊 goroutine：

- 一个叫 sweep.g，主要负责清扫死对象，合并相关的空闲页；
- 一个叫 scvg.g，主要负责向操作系统归还内存。

```
(dlv) goroutines
* Goroutine 1 - User: ./int.go:22 main.main (0x10572a6) (thread 5247606)
  Goroutine 2 - User: /usr/local/go/src/runtime/proc.go:367 runtime.gopark (0x102e596) [force gc (idle) 455634h24m29.787802783s]
  Goroutine 3 - User: /usr/local/go/src/runtime/proc.go:367 runtime.gopark (0x102e596) [GC sweep wait]
  Goroutine 4 - User: /usr/local/go/src/runtime/proc.go:367 runtime.gopark (0x102e596) [GC scavenge wait]
```

这里的 GC sweep wait 和 GC scavenge wait， 就是这两个 goroutine。

当 GC 的标记流程结束之后，sweep goroutine 就会被唤醒，进行清扫工作，其实就是循环执行 sweepone -> sweep。针对每个 mspan，sweep.g 的工作是将标记期间生成的 bitmap 替换掉分配时使用的 bitmap：

![图片](images/a8014d7470984fa8b2e7bdd69a66d04f.jpg)

然后根据 mspan 中的槽位情况决定该 mspan 的去向：

- 如果 mspan 中存活对象数 = 0，也就是所有 element 都变成了内存垃圾，那执行 freeSpan -> 归还组成该 mspan 所使用的页，并更新全局的页分配器摘要信息；
- 如果 mspan 中没有空槽，说明所有对象都是存活的，将其放入 fullSwept 队列中；
- 如果 mspan 中有空槽，说明这个 mspan 还可以拿来做内存分配，将其放入 partialSweep 队列中。

之后“清道夫” scvg goroutine 被唤醒，执行线性流程，一路运行到将页内存归还给操作系统，也就是 bgscavenge -> pageAlloc.scavenge -> pageAlloc.scavengeOne -> pageAlloc.scavengeRangeLocked -> sysUnused -> madvise：

![](images/2aee70e015c8412790c9fbb08af6de6c.jpg)



### 46.5 问题分析

从前面的基础知识中，可以总结出 Go 语言垃圾回收的关键点：

- 无分代；
- 与应用执行并发；
- 协助标记流程；
- 并发执行时开启 write barrier。

我们日常编码中就需要考虑这些关键点，进行一些针对性的设计与优化。比如，因为无分代，当我们遇到一些需要在内存中保留几千万 kv map 的场景（比如机器学习的特征系统）时，就需要想办法降低 GC 扫描成本。

又比如，因为有协助标记，当应用的 GC 占用的 CPU 超过 25% 时，会触发大量的协助标记，影响应用的延迟，这时也要对 GC 进行优化。

简单的业务场景，我们使用 sync.Pool 就可以带来较好的优化效果，若碰到一些复杂的业务场景，还要考虑 offheap 之类的欺骗 GC 的方案，比如 [dgraph 的方案](https://dgraph.io/blog/post/manual-memory-management-golang-jemalloc/)。因为我们这讲聚焦于内存分配和 GC 的实现，就不展开介绍这些具体方案了。

另外，这讲中涉及的所有内存管理的名词，你都可以在：[https://memorymanagement.org](https://memorymanagement.org/) 上找到。如果你还对垃圾回收的理论还有什么不解，我推荐你阅读：《[GC Handbook](https://gchandbook.org/)》，它可以解答你所有的疑问。





## Go语言学习资料







## 未来

![](images/image-20240704194647473.png)

字节的开源组织：https://github.com/cloudwego




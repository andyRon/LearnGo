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

### 怎样学才能少走弯路？

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

对于无类型常量参与的表达式求值，Go编译器会**根据上下文中的类型信息**，==把无类型常量自动转换为相应的类型==后，再参与求值计算，这一转型动作是隐式进行的。

但由于转型的对象是一个常量，所以这并不会引发类型安全问题，Go编译器会保证这一转型的安全性。

如果Go编译器在做隐式转型时，发现无法将常量转换为目标类型，Go编译器也会报错：

```go
const m = 1333333333

var k int8 = 1
j := k + m // 编译器报错：constant 1333333333 overflows int8 
```

> 类型常量与常量隐式转型的“珠联璧合”使得在Go这样的具有强类型系统的语言，在处理表达式混合数据类型运算的时候具有比较大的灵活性，代码编写也得到了一定程度的简化。也就是说，我们不需要在求值表达式中做任何显式转型了。
>
> 所以说，**在Go中，使用无类型常量是一种惯用法**。

### 14.4 实现枚举

Go语言其实并没有原生提供枚举类型。

Go语言中，可以使用**const代码块**定义的常量集合，来实现枚举。这是因为枚举类型本质上就是一个由**==有限数量常量==**所构成的集合。

Go将C语言枚举类型的这种基于前一个枚举值加1的特性，分解成了Go中的两个特性：**自动重复上一行，以及引入const块中的行==偏移量指示器==`iota`**，这样它们就可以分别独立使用了。

#### Go的const语法提供了“==隐式重复前一个非空表达式==”的机制

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

如果一个Go源文件中有多个const代码块定义的不同枚举，每个const代码块中的iota也是**独立变化**的:

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

在特定场景下显式字面值（直接赋值）更具优势，比如：

#### 1 需要明确业务含义的离散值

当枚举常量代表具有特定业务含义的离散数值时（例如 HTTP 状态码、错误码等），显式赋值可避免 iota 的自动递增与业务逻辑脱节。

```go
const (
    HTTP_OK         = 200
    HTTP_FORBIDDEN  = 403
    HTTP_NOT_FOUND  = 404
)
```

#### 2 非连续或非整数型枚举值

当枚举值需要非连续整数、字符串或混合类型时，显式赋值是唯一选择：

- 字符串枚举（如状态描述）：

```go
const (
    StatusSuccess = "SUCCESS"
    StatusFailed  = "FAILED"
)
```

iota 仅支持整数生成，无法直接用于字符串。

- 混合类型枚举（如结合整数和字符串）：

```go
const (
    CodeOK      = 0
    MessageOK   = "Operation succeeded"
)
```

#### 3 需要精细化类型控制的场景

当需严格约束常量类型或利用无类型常量的灵活性时，显式赋值更合适：

- 类型别名约束：

```go
type StatusCode int
const (
    Normal   StatusCode = 1
    Error    StatusCode = 2
)
```

显式指定 StatusCode 类型可限制函数参数类型，避免非法值传入。

- 无类型常量需求：

显式字面值允许无类型常量跨类型使用（如 const universal = 10 可用于 int 或 float64 上下文），而 iota 生成的常量默认有底层类型35。

#### 4 避免iota表达式复杂度

当 iota 需要结合复杂表达式（如乘法、位运算）时，显式赋值更易维护：

```go
// 使用 iota 的复杂表达式（易混淆）
const (
    Read   = 1 << iota // 1
    Write              // 2
    Execute            // 4
)

// 显式字面值（更直观）
const (
    Read   = 1
    Write  = 2
    Execute = 4
)
```


权衡：虽然 iota 可简化位掩码生成，但显式赋值在代码审查和后期维护中更清晰。

#### 5 需强制唯一值的稳定性

当枚举值需长期稳定且不允许因代码结构调整（如插入新常量）导致数值变化时：

```go
const (
    LegacyFlag = 100 // 历史遗留值，不可变更
    NewFeature = 200
)
```


风险：使用 iota 时插入新常量会导致后续值偏移，破坏已有逻辑。



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

数组类型变量是一个整体，这就意味着**一个数组变量表示的是==整个数组==**。这点与C语言完全不同，在C语言中，**数组变量可视为指向数组第一个元素的指针**。

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

复合类型的定义一般都是通过**类型字面值**的方式来进行的，作为复合类型之一的结构体类型也不例外:

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

如果一种类型采用零值初始化得到的**零值变量**，是有意义的，而且是直接可用的，我称这种类型为**“==零值可用==”类型**。

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

- **一旦结构体中包含非导出字段，那么这种逐一字段赋值的方式就不再被支持了**，编译器会报错

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

Go结构体类型是既数组类型之后，第二个将它的元素（结构体字段）一个接着一个以“平铺”形式，存放在一个**==连续内存块==**中的。

![](images/image-20240704120037598.png)

在真实情况下，虽然Go编译器没有在结构体变量占用的内存空间中插入额外字段，但结构体字段实际上可能并不是紧密相连的，中间可能存在“缝隙”。这些“缝隙”同样是结构体变量占用的内存空间的一部分，它们是Go编译器插入的“==填充物（Padding）==”。

![](images/image-20240708193452381.png)

填充物是因为需要==内存对齐==，指的就是各种内存对象的内存地址不是随意确定的，必须满足特定要求。

对于各种**基本数据类型**来说，它的变量的内存地址值必须是**其类型本身大小的整数倍**，比如，一个int64类型的变量的内存地址，应该能被int64类型自身的大小，也就是8整除；一个uint16类型的变量的内存地址，应该能被uint16类型自身的大小，也就是2整除。

对于结构体而言，它的变量的内存地址，只要是**它最长字段长度与==系统对齐系数==两者之间较小的那个的整数倍**就可以了。但对于结构体类型来说，还要让它**每个字段的内存地址都严格满足内存对齐要求**。

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

结构体的内存地址为**min(结构体最长字段的长度，系统内存对齐系数)**的整数倍，那么这里结构体T最长字段为i，它的长度为8，而64bit系统上的系统内存对齐系数一般为8，那么整个结构体的对齐系数就是8。

> 为什么上面的示意图还要在结构体的尾部填充了6个字节呢？

结构体T的对齐系数是8，那么就要保证每个结构体T的变量的内存地址，都能被8整除。如果只分配一个T类型变量，不再继续填充，也可能保证其内存地址为8的倍数。但如果考虑分配的是一个元素为T类型的数组，比如下面这行代码，我们虽然可以保证T[0]这个元素地址可以被8整除，但能保证T[1]的地址也可以被8整除吗？

```go
var array [10]T
```

数组是元素连续存储的一种类型，元素T[1]的地址为T[0]地址+T的大小(18)，显然无法被8整除，这将导致T[1]及后续元素的地址都无法对齐，这显然不能满足内存对齐的要求。加6变成24，就能被8整除了。

> 为什么会出现内存对齐的要求呢？

出于对处理器**存取数据效率**的考虑。在早期的一些处理器中，比如Sun公司的Sparc处理器仅支持内存对齐的地址，如果它遇到没有对齐的内存地址，会引发段错误，导致程序崩溃。常见的x86-64架构处理器虽然处理未对齐的内存地址不会出现段错误，但数据的**存取性能**也会受到影响。

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

有些时候，为了保证某个字段的内存地址有更为严格的约束，也会做**==主动填充==**。比如runtime包中的`mstats`结构体定义就采用了主动填充：

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

### 小结

Go不是一门面向对象范式的编程语言，没有C++或Java中的那种class类型。Go中结构体类型与class接近。

Go中的结构体类型提供了一种聚合抽象能力，开发者可以使用它建立对真实世界的事物的抽象。

自定义新类型的方式：

- 类型定义
- 类型别名

结构体类型变量的初始化方式：

- 零值初始化
- 复合字面值初始化
- 使用特定构造函数进行初始化

结构体类型和数组类型一样，都是以平铺形式存放在连续内存块中的类型。但前者有内存对齐的要求。

### 思考题

> Go语言不支持在结构体类型定义中，递归地放入其自身类型字段，但却可以拥有自身类型的指针类型、以自身类型为元素类型的切片类型，以及以自身类型作为value类型的map类型的字段，你能思考一下其中的原因吗？

在 Go 语言中，结构体不能直接包含自身类型的字段，但可以包含指向自身的指针、切片或 map 类型的字段，这背后涉及语言设计中的**内存管理机制和类型系统的限制**。具体原因如下：

#### 一、内存分配确定性要求

Go 编译器需要在编译期间确定所有类型的内存大小，这是实现高效内存分配和类型安全的基础。

1. 直接递归导致无限大小  

   如果结构体直接包含自身类型的字段（例如 type T struct { t T }），编译器会陷入无限递归计算内存大小的困境。假设 T 需要确定大小，就必须先计算其字段 t 的大小，而 t 又是一个 T 类型，导致无限循环。  

2. 指针、切片和 Map 的固定大小  

指针：无论指向的数据类型如何，指针的大小是固定的（32 位系统为 4 字节，64 位系统为 8 字节）。编译器只需为指针分配固定内存，无需关心其指向的具体数据类型大小。  

切片和 Map：切片底层是一个包含指针、长度和容量的结构体（总大小固定），而 Map 底层是哈希表的指针。两者均通过间接引用存储数据，内存大小在编译期即可确定。  

#### 二、引用类型的间接性

指针、切片和 Map 本质上是引用类型，它们存储的是对数据的引用而非数据本身。这种设计避免了结构体直接包含实际数据导致的递归问题：  

1. 指针的间接引用  

   例如 type Node struct { next *Node }，next 字段仅存储一个地址，而非完整的 Node 实例。即使存在递归关系，内存分配仍可完成。  

2. 切片和 Map 的动态特性  

切片底层是数组的视图，实际数据存储在堆中。结构体中的切片字段仅包含指向数组的指针、长度和容量信息（总大小固定）。  

Map 底层是哈希表，结构体中的 Map 字段仅存储哈希表的指针和元数据。  

这些引用类型的字段不会导致结构体自身大小的递归增长。

#### 三、类型系统的设计考量

Go 语言类型系统的设计强调简洁性和安全性，通过以下规则避免复杂性和潜在错误：  

1. 禁止无限递归类型  

直接包含自身类型的字段会导致类型定义无限递归，这违反了类型系统的确定性原则。而指针、切片和 Map 通过间接引用规避了这一问题。  

2. 内存安全与性能优化  

指针、切片和 Map 的内存管理由运行时系统统一处理，避免了用户手动管理递归结构的复杂性。  

编译器可以静态分析引用类型的内存布局，优化数据访问效率（如切片扩容时的内存分配策略）。

#### 四、实际应用场景

允许间接引用自身类型的字段，为常见数据结构提供了实现基础：  

- 链表和树结构  

```go
type TreeNode struct {
   Value    int
   Children []*TreeNode  // 通过切片存储子节点
}
```

- 图结构  

   使用 Map 存储邻接关系：  

 ```go
 type Graph struct {
        Nodes map[string]*Node  // Key为节点标识，Value为节点指针
 }    
 ```

- 复杂嵌套数据  

   通过指针和切片实现递归数据结构（如 JSON 解析时的嵌套对象）。

#### 总结

Go 语言通过限制结构体直接包含自身类型字段，保证了编译期的内存确定性；同时允许指针、切片和 Map 字段的间接引用，既避免了无限递归问题，又为动态数据结构提供了灵活的实现方式。这一设计平衡了类型安全、内存效率与编程灵活性。

## 18 控制结构：if的“快乐路径”原则

Go中程序的分支结构：if和switch-case两种语句形式；

循环结构：只有for。

操作符优先级决定了操作数优先参与哪个操作符的求值运算：

![](images/image-20240709105248548.png)

### if语句的自用变量

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

### if语句的“快乐路径”原则

```go
//伪代码段1：
func doSomething() error {
	if errorCondition1 {
		// some error logic
		... ...
		return err1
	}
	
	// some success logic
	... ...

	if errorCondition2 {
		// some error logic
		... ...
		return err2
	}

	// some success logic
	... ...
	return nil
}

// 伪代码段2：
func doSomething() error {
	if successCondition1 {
		// some success logic
		... ...

		if successCondition2 {
			// some success logic
			... ...

			return nil
		} else {
			// some error logic
			... ...
			return err2
		}
	} else {
		// some error logic
		... ...
		return err1
	}
}
```

伪代码段1有这几个特点：

- 没有使用else分支，失败就立即返回；
- “成功”逻辑始终“居左”并延续到函数结尾，没有被嵌入到if的布尔表达式为true的代码分支中；
- 整个代码段布局扁平，没有深度的缩进；

伪代码段2使用了带有嵌套的二分支结构，特点如下：

- 整个代码段呈现为“锯齿状”，有深度缩进；
- “成功”逻辑被嵌入到if的布尔表达式为true的代码分支中；

Go社区把伪代码段1这种if语句的使用方式称为if语句的**“快乐路径（Happy Path）”原则**，所谓“快乐路径”也就是成功逻辑的代码执行路径，特点：

- 仅使用单分支控制结构；
- 当布尔表达式求值为false时，也就是出现错误时，在单分支中快速返回；
- 正常逻辑在代码布局上始终“靠左”，这样读者可以从上到下一眼看到该函数正常逻辑的全貌；
- 函数执行到最后一行代表一种成功状态。

### 思考题

> 如果一个if语句使用了多分支结构，那么if语句中的几个布尔表达式如何排列能达到最好的效果呢？

在多分支 `if` 语句中，布尔表达式的排列顺序对代码的性能和可读性有显著影响。以下是基于硬件特性、编译器优化和编程实践的综合优化策略：

#### 一、按条件成立概率排序

1. **将最可能成立的条件放在最前面**
   CPU 的分支预测机制会优先预取和执行条件为真的代码路径。若高频条件在前，可减少预测失败导致的流水线清空和性能损失。例如：

   ```
   // 假设条件 A 成立概率为 80%，条件 B 为 15%，其他为 5%
   if (A) { /* 高频路径 */ }
   else if (B) { /* 次高频路径 */ }
   else { /* 低频路径 */ }
   ```

   这种排列能最大化分支预测成功率，减少 CPU 因预测错误导致的性能惩罚（通常 10-20 时钟周期）。

2. **动态概率的应对策略**
   若条件概率无法静态预估，可通过运行时统计（如 Profile-Guided Optimization, PGO）收集数据，指导编译器生成更优代码。

#### 二、按计算成本排序

1. **低计算成本的条件优先判断**
   轻量条件（如布尔变量、简单比较）应放在复杂条件（如函数调用、数学运算）之前，避免执行冗余计算。例如：

   ```
   // 低成本的 `is_valid` 判断优先
   if (!is_valid) { /* 快速失败 */ }
   else if (compute_heavy_check(data)) { /* 复杂计算 */ }
   ```

   这符合“卫语句（Guard Clause）”思想，能提前过滤无效情况。

2. **避免重复计算**
   若多个条件共享部分计算，可提取公共部分到外层。例如：

   ```
   // 提取公共计算
   const int value = compute_value();
   if (value < 10) { ... }
   else if (value >= 10 && value < 20) { ... }
   ```

#### 三、利用逻辑短路特性

1. **逻辑运算符的短路行为**
   在复合条件中，`&&` 和 `||` 会按顺序短路后续判断。将易失败或低成本的条件前置可提前终止计算。例如：

   ```
   // 若 ptr 常为 NULL，则优先判断
   if (ptr != NULL && ptr->is_valid()) { ... }
   ```

   这避免了 `ptr->is_valid()` 在空指针时的崩溃风险。

2. **合并或拆分条件表达式**
   若多个条件有重叠逻辑，可合并简化。例如：

   ```
   // 合并范围判断
   if (x >= 0 && x <= 100) { ... }
   // 优于分开的 `x >= 0` 和 `x <= 100`
   ```

#### 四、考虑可读性与维护性

1. **语义分组优先**
   若性能差异不大，应按逻辑相关性排序。例如用户权限校验：

   ```
   if (is_admin) { ... }
   else if (is_editor) { ... }
   else if (is_viewer) { ... }
   ```

   这种排列更符合业务逻辑的层次。

2. **避免过度优化可读性**
   若条件顺序对性能影响微乎其微（如低频调用代码），优先保证代码清晰性。例如：

   ```
   // 清晰的错误码处理
   if (error == NETWORK_ERROR) { ... }
   else if (error == FILE_NOT_FOUND) { ... }
   ```

------

#### 五、针对特定场景的优化

1. **互斥条件的处理**
   若条件互斥（如枚举值），可用 `switch` 替代 `if-else`，编译器可能生成跳转表优化：

   ```
   switch (mode) {
       case MODE_A: ... break;
       case MODE_B: ... break;
       default: ...;
   }
   ```

2. **无分支编程技巧**
   对性能敏感的场景（如高频循环），可用算术或位运算替代条件判断。例如：

   ```
   // 避免分支：条件成立时加 1，否则加 0
   sum += (condition) * value;
   ```

   现代编译器（如 GCC 12+）会自动优化简单条件为无分支指令（如 `cmov`）。

------

#### 总结建议

1. **性能优先场景**：按 **条件概率 > 计算成本 > 逻辑短路** 排序，结合 PGO 和性能分析工具（如 `perf`）验证优化效果。
2. **可读性优先场景**：按 **逻辑相关性 > 语义分组** 排序，必要时添加注释说明排序理由。
3. **工具辅助**：使用编译器的优化选项（如GCC的 `-O3`、`-fprofile-use`）自动化部分优化。

实际应用中需权衡性能与代码维护成本，高频热点代码可激进优化，低频代码以清晰性为主。

## 19 控制结构：Go的for循环，仅此一种

### 19.1 for语句的经典使用形式

![](images/image-20241225161246544.png)

1️⃣ **循环前置语句**。**仅会被执行一次**

2️⃣ **条件判断表达式**

3️⃣ **循环体**

4️⃣ **循环后置语句**

#### 仅保留循环判断条件表达式

```go
i := 0
for i < 10 {
    println(i)
    i++
}
```

除了循环体之外，我们仅保留循环判断条件表达式。

#### 无限循环

三种等价形式

```go
for { 
   // 循环体代码
}

for true { 
   // 循环体代码
}

for ; ; { 
   // 循环体代码
}
```



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

for range对于string类型来说，**每次循环得到的v值是一个Unicode字符码点**，也就是`rune`类型值，而不是一个字节，返回的第一个值i为**该Unicode字符码点的内存编码（UTF-8）的第一个字节在字符串内存序列中的位置**。

> for与for range，对string类型进行循环操作的语义是不同的。



#### map

在Go语言中，**对map进行循环操作，for range是唯一的方法**。

```go
var m = map[string]int {
	"Rob" : 67,
    "Russ" : 39,
    "John" : 29,
}

for k, v := range m {
    println(k, v)
}
```

for range对于map类型来说，**每次循环，循环变量k和v分别会被赋值为map键值对集合中一个元素的key值和value值**。

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



### 19.5 for语句的常见“坑”与避坑方法 

for range这个语法糖提高了Go的表达能力，同时也可能带入一些问题。

#### 1️⃣循环变量的重用

循环变量重用问题已在 **Go 1.22**修复。

#### 2️⃣参与循环的是range表达式的副本



#### 3️⃣遍历map中元素的随机性



## 20 控制结构：Go中的switch语句有哪些变化？

switch语句是按照case出现的先后顺序对case表达式进行求值的，如果将匹配成功概率高的case表达式排在前面，就会有助于提升switch语句执行效率。这点对于case后面是表达式列表的语句同样有效，可以将匹配概率最高的表达式放在表达式列表的最左侧。

无论default分支出现在什么位置，它都只会在所有case都没有匹配上的情况下才会被执行的。

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






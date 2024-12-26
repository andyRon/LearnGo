Go Web
---

https://learnku.com/docs/build-web-application-with-golang


https://github.com/astaxie/build-web-application-with-golang

## 1 GO环境配置

## 2 GO语言基础 ❤️

`&x`表示x所在地址；`*int`表示对应的指针类型。



二十五个关键字：

```go
break    default      func    interface    select
case     defer        go      map          struct
chan     else         goto    package      switch
const    fallthrough  if      range        type
continue for          import  return       var
```

- var 和 const 变量和常量申明
- package 和 import 
- func 用于定义函数和方法
- return 用于从函数返回
- defer 用于类似析构函数
- go 用于并发
- select 用于选择不同类型的通讯
- interface 用于定义接口
- struct 用于定义抽象数据类型
- break、case、continue、for、fallthrough、else、if、switch、goto、default 流程
- chan 用于 channel 通讯
- type 用于声明自定义类型
- map 用于声明 map 类型数据
- range 用于读取 slice、map、channel 数据



### 2.3 流程和函数



### 2.4 struct



### 2.5 面向对象



### 2.6 interface



#### 嵌入interface



#### 反射



### 2.7 并发

#### goroutine



#### channels

创建：

```go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
```

channel通过操作符`<-`来接收和发送数据:

```go
ch <- v    // 发送v到channel ch.
v := <-ch  // 从ch中接收数据，并赋值给v
```



#### Buffered Channels



#### Range和Close



#### Select



#### 超时



#### runtime goroutine







## 3 web基础

### 3.1 Web工作方式



### 3.2 Go搭建一个Web服务器



### 3.3 Go如何使得Web工作

服务器端的几个概念

Request：用户请求的信息，用来解析用户的请求信息，包括post、get、cookie、url等信息

Response：服务器需要反馈给客户端的信息

Conn：用户的每次请求链接

Handler：处理请求和生成返回信息的处理逻辑



### 3.4 Go的http包详解

Go的http有两个核心功能：Conn、ServeMux



## 4 表单

### 4.1 处理表单的输入



### 4.2 验证表单的输入



### 4.3 预防跨站脚本



### 4.4 防止多次递交表单



### 4.5 处理文件上传







## 5 访问数据库

Go没有内置的驱动支持任何的数据库，但是Go定义了database/sql接口，用户可以基于驱动接口开发相应数据库的驱动。

### 5.1 database/sql接口

#### sql.Register

这个存在于database/sql的函数是用来注册数据库驱动的，当第三方开发者开发数据库驱动时，都会实现init函数，在init里面会调用这个`Register(name string, driver driver.Driver)`完成本驱动的注册。

```go
//https://github.com/mattn/go-sqlite3驱动
func init() {
	sql.Register("sqlite3", &SQLiteDriver{})
}

//https://github.com/mikespook/mymysql驱动
// Driver automatically registered in database/sql
var d = Driver{proto: "tcp", raddr: "127.0.0.1:3306"}
func init() {
	Register("SET NAMES utf8")
	sql.Register("mymysql", &d)
}
```





#### driver.Driver

Driver是一个数据库驱动的接口，他定义了一个`method: Open(name string)`，这个方法返回一个数据库的Conn接口。



#### driver.Conn





#### driver.Stmt



#### driver.Tx



#### driver.Execer



#### driver.Result





#### driver.Rows



#### driver.RowsAffected



#### driver.Value



#### driver.ValueConverter



#### driver.Valuer





### 5.2 使用MySQL数据库

mysql驱动 https://github.com/go-sql-driver/mysql 支持database/sql，全部采用go写



### 5.3 使用SQLite数据库

https://github.com/mattn/go-sqlite3 支持database/sql接口，基于cgo(关于cgo的知识请参看官方文档或者本书后面的章节)写的



### 5.4 使用PostgreSQL数据库

https://github.com/lib/pq 支持database/sql驱动，纯Go写的



### 5.5 使用Beego orm库进行ORM开发



### 5.6 NOSQL数据库操作

#### redis

https://github.com/redis/go-redis

https://github.com/gomodule/redigo



#### mongoDB

MongoDB是一个高性能，开源，无模式的文档型数据库，是一个介于关系数据库和非关系数据库之间的产品，是非关系数据库当中功能最丰富，最像关系数据库的。他支持的数据结构非常松散，采用的是类似json的bjson格式来存储数据，因此可以存储比较复杂的数据类型。Mongo最大的特点是他支持的查询语言非常强大，其语法有点类似于面向对象的查询语言，几乎可以实现类似关系数据库单表查询的绝大部分功能，而且还支持对数据建立索引。

MongoDB和Mysql的操作对比图:

![](images/5.6.mongodb.png)



## 6 session和数据存储



### 6.2 Go如何使用session



### 6.3 session存储





### 6.4 预防session劫持





## 7 文本处理

### 7.1 XML处理



### 7.2 JSON处理



### 7.3 正则处理



### 7.4 模板处理



### 7.5 文件操作



### 7.6 字符串处理



## 8 Web服务

### 8.1 Socket编程



### 8.2 WebSocket



### 8.3 REST



### 8.4 RPC



## 9 安全与加密

### 9.1 预防CSRF攻击



### 9.2 确保输入过滤



### 9.3 避免XSS攻击



### 9.4 避免SQL注入



### 9.5 存储密码



### 9.6 加密和解密数据





## 10 国际化和本地化

### 10.1 设置默认地区



### 10.2 本地化资源



### 10.3 国际化站点



## 11 错误处理，调试和测试



### 11.1 错误处理



### 11.2 使用GDB调试

https://github.com/go-delve/delve



### 11.3 Go怎么写测试用例



## 12 部署与维护

### 12.1 应用日志



### 12.2 网站错误处理



### 12.3 应用部署



### 12.4 备份和恢复



## 13 如何设计一个Web框架



## 14 扩展Web框架

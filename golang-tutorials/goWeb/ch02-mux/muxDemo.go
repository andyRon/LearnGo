package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)                    // 设置响应状态码为 200
	fmt.Fprintf(w, "Hello, %s!!!!", params["name"]) // 发送响应到客户端
}

// 自定义处理器
type HelloWorldHandler struct {
}

func (handler *HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "你好，%s!! 这是自定义处理器。", params["name"])
}

func listPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "文章列表")
}

func createPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "发布文章")
}

func updatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "修改文章")
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "删除文章")
}

func showPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "文章详情")
}

// 日志中间件
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("日志中间件: " + r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

// 简单校验中间体
func checkToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.FormValue("token")
		if token == "9527" {
			log.Printf("Token check success: %s\n", r.RequestURI)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

func main() {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	r.HandleFunc("/hello", sayHelloWorld)
	//r.HandleFunc("/hello/{name}", sayHelloWorld)
	r.HandleFunc("/hello/{name:[a-z]+}", sayHelloWorld).Methods("GET", "POST") // 用正则限制参数的字符
	r.Handle("/hello/zh/{name}", &HelloWorldHandler{}).Methods("GET")

	r.PathPrefix("/x").HandlerFunc(sayHelloWorld) // 路由前缀

	// 限定请求参数
	// curl http://localhost:8080/r/header -H "X-Requested-With: XMLHttpRequest"
	r.HandleFunc("/r/header", func(w http.ResponseWriter, r *http.Request) {
		header := "X-Requested-With"
		fmt.Fprintf(w, "包含指定请求头[%s=%s]", header, r.Header[header])
	}).Headers("X-Requested-With", "XMLHttpRequest")

	// 限定查询字符串。如，查询字符串必须包含 token 且值为 test 才可以匹配到给定路由 /query/string
	// curl http://localhost:8080/query/string\?token\=test
	r.HandleFunc("/query/string", func(w http.ResponseWriter, r *http.Request) {
		query := "token"
		fmt.Fprintf(w, "包含指定查询字符串[%s=%s]", query, r.FormValue(query))
	}).Queries("token", "test")

	// 自定义匹配规则
	// curl http://localhost:8080/custom/matcher -H "Referer:http://andyron.top"
	r.HandleFunc("/custom/matcher", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "请求来自指定域名: %s", r.Referer())
	}).MatcherFunc(func(request *http.Request, match *mux.RouteMatch) bool { // 对请求进行判断、筛选
		return request.Referer() == "http://andyron.top"
	})

	// 路由分组（基于子路由+路径前缀）
	postRouter := r.PathPrefix("/posts").Subrouter()
	//postRouter := r.PathPrefix("/posts").Host("admin.goweb.test").Subrouter()   // 可以结合子域名做进一步划分
	postRouter.Use(checkToken) // 添加一个校验中间体  http://localhost:8080/posts/?token=9527

	// 处理静态资源
	// 解析服务器启动参数 dir 作为静态资源 Web 根目录
	// 默认是当前目录 .
	var dir string
	flag.StringVar(&dir, "dir", ".", "静态资源所在目录，默认为当前目录")
	flag.Parse()
	// 处理形如 http://localhost:8000/static/<filename> 的静态资源路由
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	postRouter.HandleFunc("/", listPosts).Methods("GET").Name("posts.index") // 路由命名
	postRouter.HandleFunc("/create", createPost).Methods("POST")
	postRouter.HandleFunc("/update", updatePost).Methods("PUT")
	postRouter.HandleFunc("/delete", deletePost).Methods("DELETE")
	postRouter.HandleFunc("/show", showPost).Methods("GET")

	// 使用路由名
	indexUrl, _ := r.Get("posts.index").URL()
	log.Println("文章列表链接：", indexUrl)

	log.Fatal(http.ListenAndServe(":8080", r))
}

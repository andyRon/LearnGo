package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//log.Println("这是一条很普通的日志。")
	//v := "很普通的"
	//log.Printf("这是一条%s日志。\n", v)
	//log.Fatalln("这是一条会触发fatal的日志。")
	//log.Panicln("这是一条会触发panic的日志。")

	fmt.Println(log.Ldate)
	fmt.Println(log.Ltime)
	fmt.Println(log.Lmicroseconds)
	fmt.Println(log.Llongfile)
	fmt.Println(log.Lshortfile)
	fmt.Println(log.LUTC)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("这是一条很普通的日志。")

	// 前缀
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
	log.SetPrefix("[pprof]")
	log.Println("这是一条很普通的日志。")

	// 日志输出位置，一般会写在init函数
	//logFile, err := os.OpenFile("./log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	//if err != nil {
	//	log.Println("open log file failed, err:", err)
	//	return
	//}
	//defer logFile.Close()
	//log.SetOutput(logFile)
	//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//log.Println("这是一条很普通的日志。")
	//log.SetPrefix("[andyron]")
	//log.Println("这是一条很普通的日志。")

	// 自定义logger
	logger := log.New(os.Stdout, "[andyron]", log.Ldate|log.Ltime|log.Lshortfile)
	//logger := log.New(logFile, "[andyron]", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("这是一条自定义的logger记录的日志。")
}

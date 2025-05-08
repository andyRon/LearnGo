package main

import (
	"fmt"
	//_ "github.com/go-redis/redis/v7" // “_”为空导入
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Println("hello, go module mode.")
	logrus.Println(uuid.NewString())
}

func init() {
	fmt.Println("123")
}

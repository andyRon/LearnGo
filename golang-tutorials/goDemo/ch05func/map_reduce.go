package main

import (
	"fmt"
	"strconv"
)

// 年龄求和
func ageSum(users []map[string]string) int {
	var sum int
	for _, user := range users {
		num, _ := strconv.Atoi(user["age"])
		sum += num
	}
	return sum
}

// Map 映射转化函数
func mapToString(items []map[string]string, f func(map[string]string) string) []string {
	newSlice := make([]string, len(items))
	for _, item := range items {
		newSlice = append(newSlice, f(item))
	}
	return newSlice
}

// Reduce 求和函数
func fieldSum(items []string, f func(string) int) int {
	var sum int
	for _, item := range items {
		sum += f(item)
	}
	return sum
}

// Filter 函数
func itemsFilter(items []map[string]string, f func(map[string]string) bool) []map[string]string {
	newSlice := make([]map[string]string, len(items))
	for _, item := range items {
		if f(item) {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}

func main() {
	var users = []map[string]string{
		{
			"name": "张三",
			"age":  "17",
		},
		{
			"name": "李四",
			"age":  "19",
		},
		{
			"name": "王五",
			"age":  "38",
		},
		{
			"name": "赵六",
			"age":  "-10",
		},
		{
			"name": "孙七",
			"age":  "60",
		},
		{
			"name": "周八",
			"age":  "10",
		},
	}
	//fmt.Printf("用户年龄累加结果：%d\n", ageSum(users))

	// 应用 Filter 函数对无效用户年龄进行过滤，或者排除指定范围年龄
	validUsers := itemsFilter(users, func(user map[string]string) bool {
		age, ok := user["age"]
		if !ok {
			return false
		}
		intAge, err := strconv.Atoi(age)
		if err != nil {
			return false
		}
		if intAge < 18 || intAge > 35 {
			return false
		}
		return true
	})

	ageSlice := mapToString(validUsers, func(user map[string]string) string {
		return user["age"]
	})
	sum := fieldSum(ageSlice, func(age string) int {
		intAge, _ := strconv.Atoi(age)
		return intAge
	})
	fmt.Printf("用户年龄累加结果：%d\n", sum)
}

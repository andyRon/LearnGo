package main

import (
	"fmt"
	"sort"
	"strconv"
)

func str1() {
	str := "hello, world"
	str1 := str[:5]  // 获取索引5（不含）之前的子串
	str2 := str[7:]  // 获取索引7（含）之后的子串
	str3 := str[0:5] // 获取从索引0（含）到索引5（不含）之间的子串
	fmt.Println("str1:", str1)
	fmt.Println("str2:", str2)
	fmt.Println("str3:", str3)
}

func str2() {
	str := "Hello, 世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，ch 类型为 byte
		fmt.Println(i, ch)
	}
}

func str3() {
	str := "Hello, 世界"
	for i, ch := range str {
		fmt.Println(i, ch) // ch 的类型为 rune
	}

}

func arr() {
	// 通过二维数组生成九九乘法表
	var multi [9][9]string
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			n1 := i + 1
			n2 := j + 1
			if n1 < n2 { // 摒除重复的记录
				continue
			}
			multi[i][j] = fmt.Sprintf("%dx%d=%d", n2, n1, n1*n2)
		}
	}

	// 打印九九乘法表
	for _, v1 := range multi {
		for _, v2 := range v1 {
			fmt.Printf("%-8s", v2) // 位宽为8，左对齐
		}
		fmt.Println()
	}
}

func slice1() {

	// 先定义一个数组
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	// 基于数组创建切片
	q2 := months[3:6]     // 第二季度
	summer := months[5:8] // 夏季

	fmt.Println(q2)
	fmt.Println(summer)

	//all := months[:]
	firsthalf := months[:6]
	//secondhalf := months[6:]
	//
	//mySlice := make([]int, 5, 10)

	q1 := firsthalf[:9]
	fmt.Println(q1)
}

func map1() {
	var testMap map[string]int
	testMap = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	k := "two"
	v, ok := testMap[k]
	if ok {
		fmt.Printf("The element of key %q: %d\n", k, v)
	} else {
		fmt.Println("Not found!")
	}

	for i, value := range testMap {
		fmt.Println(i, value)
	}

	invMap := make(map[int]string, 3)
	for k, v := range testMap {
		invMap[v] = k
	}
	for k, v := range invMap {
		fmt.Println(k, v)
	}

	keys := make([]string, 0)
	for k, _ := range testMap {
		keys = append(keys, k)
	}
	sort.Strings(keys) // 对键进行排序
	fmt.Println("Sorted map by key:")
	for _, k := range keys {
		fmt.Println(k, testMap[k])
	}

	values := make([]int, 0)
	for _, v := range testMap {
		values = append(values, v)
	}
	sort.Ints(values) // 对值进行排序
	fmt.Println("Sorted map by value:")
	for _, v := range values {
		fmt.Println(invMap[v], v)
	}
}

func poiter1() {
	a := 100
	var ptr *int // 声明指针类型
	ptr = &a     // 初始化指针类型值为变量 a
	fmt.Println(ptr)
	fmt.Println(*ptr)
}

func strToNum() {
	v1 := "100"
	v2, _ := strconv.Atoi(v1) // 将字符串转化为整型，v2 = 100

	v3 := 100
	v4 := strconv.Itoa(v3) // 将整型转化为字符串, v4 = "100"

	v5 := "true"
	v6, _ := strconv.ParseBool(v5) // 将字符串转化为布尔型
	v5 = strconv.FormatBool(v6)    // 将布尔值转化为字符串

	v7 := "100"
	v8, _ := strconv.ParseInt(v7, 10, 64) // 将字符串转化为整型，第二个参数表示进制，第三个参数表示最大位数
	v7 = strconv.FormatInt(v8, 10)        // 将整型转化为字符串，第二个参数表示进制

	v9, _ := strconv.ParseUint(v7, 10, 64) // 将字符串转化为无符号整型，参数含义同 ParseInt
	v7 = strconv.FormatUint(v9, 10)        // 将无符号整数型转化为字符串，参数含义同 FormatInt

	v10 := "99.99"
	v11, _ := strconv.ParseFloat(v10, 64) // 将字符串转化为浮点型，第二个参数表示精度
	v10 = strconv.FormatFloat(v11, 'E', -1, 64)

	q := strconv.Quote("Hello, 世界")       // 为字符串加引号
	q = strconv.QuoteToASCII("Hello, 世界") // 将字符串转化为 ASCII 编码

	fmt.Println(v2, v4, q)
}

func copy1() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}

	// 复制 slice1 到 slice 2
	//copy(slice2, slice1) // 只会复制 slice1 的前3个元素到 slice2 中
	// slice2 结果: [1, 2, 3]
	//fmt.Println(slice2)

	// 复制 slice2 到 slice 1
	copy(slice1, slice2) // 只会复制 slice2 的 3 个元素到 slice1 的前 3 个位置
	// slice1 结果：[5, 4, 3, 4, 5]
	fmt.Println(slice1)
}

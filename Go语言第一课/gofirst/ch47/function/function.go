package main

import "fmt"

/*
所有内置函数
*/
/* */

func main() {
	/* 一、基本数据操作类 */
	// 1 append 向切片动态追加元素，支持单元素、多元素或另一个切片的合并
	s := []int{1, 2}
	s = append(s, 3)
	s = append(s, 4, 5)
	s = append(s, []int{6, 7, 8}...) // 合并切片，... 是 Go 语言中的展开操作符，它将切片 []int{6, 7, 8} 展开为单独的元素 6, 7, 8
	fmt.Println(s)                   // [1,2,3,4,5,6,7,8]

	// 2 copy 复制切片元素到目标切片，返回实际复制的元素数量
	src := []int{1, 2, 3}
	dst := make([]int, 2)
	n := copy(dst, src) // dst=[1,2], n=2
	println(n)

	// 3 len 获取字符串、数组、切片、map、通道的长度或元素数量
	s1 := "hello"
	m := map[string]int{"a": 1}
	fmt.Println(len(s1)) // 5
	fmt.Println(len(m))  // 1

	// 4 cap 返回切片或通道的容量（底层数组可容纳的最大元素数）
	ch := make(chan int, 5)
	fmt.Println(cap(ch)) // 5

	/* 二、内存管理类 */
	// 5 new 为值类型（如 int、结构体）分配内存，返回指向零值的指针
	p := new(int)
	fmt.Println(*p) // 0

	// 6 make 初始化引用类型（切片、map、通道），返回可直接使用的实例
	slice := make([]int, 2, 5) // 长度2，容量5的切片
	fmt.Println(slice)

	// 7 clear 清空 map 或切片的所有元素（切片置零值，map 删除所有键）
	map1 := map[string]int{"a": 1, "b": 2}
	fmt.Println(map1)
	clear(map1)
	fmt.Println(map1)

	/* 三、错误处理与并发控制 */
	// 8 panic 触发运行时错误，终止程序流程
	//panic("critical error") // TODO

	// 9 recover 用途：捕获 panic 错误，需与 defer 结合使用
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	// 10 close 关闭通道，阻止继续发送数据（接收操作仍可进行）
	ch1 := make(chan int)
	close(ch1)

	/* 四、类型转换与特殊操作 */
	// 11 delete 从 map 中删除键值对，若键不存在则无操作
	map2 := map[string]int{"a": 1}
	delete(map2, "a") // m=map[]

	// 12 complex、real、imag 用于复数操作，返回实部、虚部
	c := complex(3, 4)   // 3+4i
	fmt.Println(real(c)) // 3

	// 13 print、println 打印输出，print 不换行，println 换行

}

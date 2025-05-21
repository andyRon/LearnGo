package main

import (
	"fmt"
	"sort"
)

// 学生成绩结构体
type StuScore struct {
	name  string
	score int
}

type StuScores []StuScore

func (s StuScores) Len() int {
	return len(s)
}

func (s StuScores) Less(i, j int) bool {
	// 升序
	return s[i].score < s[j].score
	// 降序
	// return s[i].score > s[j].score
}

func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	stus := StuScores{
		{"zs", 90},
		{"ls", 70},
		{"ww", 80},
		{"ar", 95},
	}

	fmt.Println("默认：", stus)
	sort.Sort(stus)
	fmt.Println("是否排序：", sort.IsSorted(stus))
	fmt.Println("已经排过序", stus)

	// 翻转
	sort.Sort(sort.Reverse(stus)) // TODO
	fmt.Println(stus)

	// 查找
	x := 11
	s := []int{3, 6, 8, 11, 45} // 注意已经升序排序
	pos := sort.Search(len(s), func(i int) bool { return s[i] >= x })
	if pos < len(s) && s[pos] == x {
		fmt.Println(x, " 在 s 中的位置为：", pos)
	} else {
		fmt.Println("s 不包含元素 ", x)
	}
	//GuessingGame()

	//
	s1 := []int{5, 2, 6, 3, 1, 4} // 未排序的切片数据
	sort.Ints(s1)
	fmt.Println(s1) // 将会输出[1 2 3 4 5 6]
	sort.Sort(sort.Reverse(sort.IntSlice(s1)))
	fmt.Println(s1)

	// SearchInts使用限制：必须排序
	s2 := []int{5, 2, 6, 3, 1, 4}       // 未排序的切片数据
	fmt.Println(sort.SearchInts(s2, 2)) // 将会输出 0 而不是 1
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

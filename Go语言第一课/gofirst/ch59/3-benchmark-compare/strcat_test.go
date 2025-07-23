package main

import (
	"strings"
	"testing"
)

var sl = []string{
	"Rob Pike ",
	"Robert Griesemer ",
	"Ken Thompson ",
}

func Strcat(sl []string) string {
	return concatStringByJoin(sl)
	//return concatStringByOperator(sl)
}

func concatStringByOperator(sl []string) string {
	var s string
	for _, v := range sl {
		s += v
	}
	return s
}

func concatStringByJoin(sl []string) string {
	return strings.Join(sl, "")
}

func BenchmarkStrcat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Strcat(sl)
	}
}

// 分别 通过Go原生的操作符（"+"）连接的方式和strings.Join函数两种方式实现了字符串的连接，采集一下不同性能基准数据。

/*
go test -run=NONE -bench . strcat_test.go > old.txt
go test -run=NONE -bench . strcat_test.go > new.txt

benchcmp old.txt new.txt
benchmark             old ns/op     new ns/op     delta
BenchmarkStrcat-8     47.0          32.2          -31.55%
*/

/*
// 使用-count对BenchmarkStrcat执行多次
go test -run=NONE -count 5 -bench . strcat_test.go > old2.txt
go test -run=NONE -count 5 -bench . strcat_test.go > new2.txt

benchcmp old2.txt new2.txt
benchmark             old ns/op     new ns/op     delta
BenchmarkStrcat-8     47.7          30.1          -36.92%
BenchmarkStrcat-8     47.2          30.0          -36.46%
BenchmarkStrcat-8     47.2          31.8          -32.68%
BenchmarkStrcat-8     47.0          37.3          -20.66%
BenchmarkStrcat-8     47.0          30.1          -35.90%

// -best命令行选项，benchcmp将分别从old2.txt和new2.txt中挑选性能最好的一条数据，然后进行比较
benchcmp -best old2.txt new2.txt
benchmark             old ns/op     new ns/op     delta
BenchmarkStrcat-8     47.0          30.0          -36.09%
*/

/*
go test -run=NONE -count 5 -bench . strcat_test.go -benchmem > old_with_mem.txt
go test -run=NONE -count 5 -bench . strcat_test.go -benchmem > new_with_mem.txt

benchstat old_with_mem.txt new_with_mem.txt
goos: darwin
goarch: arm64
cpu: Apple M1
         │ old_with_mem.txt │          new_with_mem.txt           │
         │      sec/op      │    sec/op     vs base               │
Strcat-8       47.38n ± ∞ ¹   30.07n ± ∞ ¹  -36.53% (p=0.008 n=5)
¹ need >= 6 samples for confidence interval at level 0.95

         │ old_with_mem.txt │          new_with_mem.txt          │
         │       B/op       │    B/op      vs base               │
Strcat-8        80.00 ± ∞ ¹   48.00 ± ∞ ¹  -40.00% (p=0.008 n=5)
¹ need >= 6 samples for confidence interval at level 0.95

         │ old_with_mem.txt │          new_with_mem.txt          │
         │    allocs/op     │  allocs/op   vs base               │
Strcat-8        2.000 ± ∞ ¹   1.000 ± ∞ ¹  -50.00% (p=0.008 n=5)
¹ need >= 6 samples for confidence interval at level 0.95

*/

package test

import (
	"fmt"
	"strings"
	"testing"
)

// go test -bench=. -benchmem
// -benchmem 选项用于在基准测试时输出内存分配相关的统计信息。具体包含每次操作的平均内存分配量（单位为字节）以及每次操作的平均内存分配次数。

const n = 10000

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < n; j++ {
			s += "a"
		}
		_ = s
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < n; j++ {
			builder.WriteString("a")
		}
		_ = builder.String()
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strs := make([]string, n)
		for j := 0; j < n; j++ {
			strs[j] = "a"
		}
		_ = strings.Join(strs, "")
	}
}

func BenchmarkFmtSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < n; j++ {
			s = fmt.Sprintf("%s%s", s, "a")
		}
		_ = s
	}
}

// BenchmarkBytesAppend 使用 []byte 进行字符串连接的基准测试
func BenchmarkBytesAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var bytes []byte
		for j := 0; j < n; j++ {
			bytes = append(bytes, "a"...)
		}
		_ = string(bytes)
	}
}

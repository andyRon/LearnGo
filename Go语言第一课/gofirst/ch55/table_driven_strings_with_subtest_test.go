package main

import (
	"strings"
	"testing"
)

// 表驱动测试与子测试（subtest）结合

func TestCompare4(t *testing.T) {
	compareTests := []struct {
		name, a, b string
		i          int
	}{
		{`compareTwoEmptyString`, "", "", 0},
		{`compareSecondParamIsEmpty`, "a", "", 1},
		{`compareFirstParamIsEmpty`, "", "a", -1},
	}

	for _, tt := range compareTests {
		t.Run(tt.name, func(t *testing.T) {
			cmp := strings.Compare(tt.a, tt.b)
			if cmp != tt.i {
				t.Errorf(`want %v, but Compare(%q, %q) = %v`, tt.i, tt.a, tt.b, cmp)
			}
		})
	}
}

// 将测试结果的判定逻辑放入一个单独的子测试中，这样可以单独执行表中某项数据的测试。

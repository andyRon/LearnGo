package main

import (
	"strings"
	"testing"
)

// 改编自$GOROOT/src/strings/compare_test.go

func TestCompare(t *testing.T) {
	var a, b string
	var i int

	a, b = "", ""
	i = 0
	cmp := strings.Compare(a, b)
	if cmp != i {
		t.Errorf("want %v, but Compare(%q, %q) = %v", i, a, b, cmp)
	}

	a, b = "a", ""
	i = 1
	cmp = strings.Compare(a, b)
	if cmp != i {
		t.Errorf("want %v, but Compare(%q, %q) = %v", i, a, b, cmp)
	}

	a, b = "", "a"
	i = -1
	cmp = strings.Compare(a, b)
	if cmp != i {
		t.Errorf("want %v, but Compare(%q, %q) = %v", i, a, b, cmp)
	}
}

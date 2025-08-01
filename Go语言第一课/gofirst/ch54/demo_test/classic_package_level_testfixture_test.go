package demo_test

import (
	"fmt"
	"testing"
)

// 包级别测试固件

// Go 1.14版本testing包增加了testing.Cleanup方法，为测试固件的销毁提供了包级原生的支持
// Go 1.4版本引入了TestMain，使得包级别测试固件的创建和销毁终于有了正式的施展舞台。

func setUp(testName string) func() {
	fmt.Printf("\tsetUp fixture for %s\n", testName)
	return func() {
		fmt.Printf("\ttearDown fixture for %s\n", testName)
	}
}

func TestFunc1(t *testing.T) {
	t.Cleanup(setUp(t.Name()))
	fmt.Printf("\tExecute test: %s\n", t.Name())
}

func TestFunc2(t *testing.T) {
	t.Cleanup(setUp(t.Name()))
	fmt.Printf("\tExecute test: %s\n", t.Name())
}

func TestFunc3(t *testing.T) {
	t.Cleanup(setUp(t.Name()))
	fmt.Printf("\tExecute test: %s\n", t.Name())
}

func pkgSetUp(pkgName string) func() {
	fmt.Printf("package SetUp fixture for %s\n", pkgName)
	return func() {
		fmt.Printf("package TearDown fixture for %s\n", pkgName)
	}
}

func TestMain(m *testing.M) {
	defer pkgSetUp("package demo_test")()
	m.Run()
}

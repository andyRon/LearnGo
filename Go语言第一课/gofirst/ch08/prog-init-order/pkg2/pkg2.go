package pkg2

import (
	"fmt"
	_ "github.com/andyron/prog-init-order/pkg3"
)

const (
	c = "c"
)

var (
	_ = constInitCheck()
	v = variableInit("v")
)

func constInitCheck() string {
	if c != "" {
		fmt.Println("pkg2: const c has been initialized")
	}
	return ""
}

func variableInit(name string) string {
	fmt.Printf("pkg2: var %s has been initialized\n", name)
	return name
}

func init() {
	fmt.Println("pkg2: init func invoked")
}

func main() {
	// do nothing
}

package main

import (
	"fmt"
	"github.com/khaitranrh/test-git/module2/pkg"
)

func main() {
	bar := pkg.Bar()
	foo := pkg.Foo()
	fmt.Printf("Hello, %s! I'm %s\n", foo, bar)
}

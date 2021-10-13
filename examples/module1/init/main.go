package main

import (
	"fmt"
	_ "github.com/cncamp/golang/examples/module1/init/a"
	_ "github.com/cncamp/golang/examples/module1/init/b"
)

func init() {
	fmt.Println("main init")
}
func main() {
	fmt.Println("hello world1")
}
func hello() {
	fmt.Println("r2")
}

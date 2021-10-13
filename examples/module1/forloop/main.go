package main

import (
	"fmt"
)

func main() {
	str := "hello go"
	for i, v := range str {
		fmt.Println(i, string(v))
	}
}

func main1() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fullString := "hello world"
	fmt.Println(fullString)
	for i, c := range fullString {
		fmt.Println(i, string(c))
	}

}

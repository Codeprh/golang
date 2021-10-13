package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//defer是finally的关键字，但是遵循栈的规则：先进后出，因此打印是go world hello
	defer fmt.Println("hello")
	defer fmt.Println("world")
	defer fmt.Println("go")

	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		//要使用闭包，匿名函数，函数结束就执行defer
		func(i int) {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("fori")
		}(i)
	}
}

func main1() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	loopFunc()
	time.Sleep(time.Second)
}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		// go func(i int) {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("loopFunc:", i)
		// }(i)
	}
}

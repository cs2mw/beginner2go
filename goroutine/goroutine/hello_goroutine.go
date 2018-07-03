package main

import (
	"fmt"
	"runtime"
)

func sayHello(s string) {
	fmt.Println("sayhello:")
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	go sayHello("world")
	sayHello("Hello")
}

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Outside a goroutine")
	go func() {
		fmt.Println("Inside goroutine")
	}()
	fmt.Println("Outside a goroutine")

	runtime.Gosched() // yields the processor, let's other routines execute before exiting.
}

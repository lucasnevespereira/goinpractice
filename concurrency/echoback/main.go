package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	// calls echo() as a goroutine
	go echo(os.Stdin, os.Stdout)
	// we need to sleep 'cause the main function is also a routine.
	// if we don't main() wil be finished before echo ends.
	time.Sleep(30 * time.Second)
	fmt.Println("Done sleeping")
	os.Exit(1)
}

func echo(in io.Reader, out io.Writer) {
	// copies data to a writer from a reader
	io.Copy(out, in)
}

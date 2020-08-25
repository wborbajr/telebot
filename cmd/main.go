package main

import (
	"runtime"
	"fmt"
)

func main()  {
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("GOOS: %s\n", runtime.GOOS)
	fmt.Printf("GOARCH: %s\n", runtime.GOARCH)
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS)
	fmt.Printf("GOROOT: %d\n", runtime.GOROOT)
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU)
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine)
	
}
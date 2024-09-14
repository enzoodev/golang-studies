package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of CPUs: ", runtime.NumCPU())
}

package main

import "runtime/debug"

func thirdFunction() {
	debug.PrintStack()
}

func secondFunction() {
	thirdFunction()
}

func firstFunction() {
	secondFunction()
}

func main() {
	firstFunction()
}

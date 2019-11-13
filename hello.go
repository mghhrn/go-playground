package main

import "fmt"
import "runtime"

func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf("number of cpu cores = ")
	fmt.Printf("%d\n", runtime.NumCPU())
}


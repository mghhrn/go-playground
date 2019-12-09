package main

import (
	"fmt"
	"runtime"

	"github.com/mghhrn/go-playground/prime-sieve"
	MyStringUtil "github.com/mghhrn/go-stringutil"
)


func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf("number of cpu cores = ")
	fmt.Printf("%d\n", runtime.NumCPU())

	fmt.Println(MyStringUtil.Reverse("!oG ,olleH"))
	prime_sieve.Sieve();

}


package main

import (
	"flag"
	"fmt"
	"main/calc"
	"strings"
	"syscall"
)

func main() {
	flag.Parse()
	expr := strings.Join(flag.Args(), "")

	result, err := calc.Calc(expr);

	if err != nil {
		fmt.Println(nil)
		syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
	}

	fmt.Println(result)
}
package main

import (
	"flag"
	"fmt"
	"main/calc"
	"strings"
)

func main() {
	flag.Parse()
	expr := strings.Join(flag.Args(), "")

	fmt.Println(calc.Calc(expr))
}
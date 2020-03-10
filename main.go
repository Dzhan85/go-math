package main

import (
	"fmt"
	"github.com/Dzhan85/go-math"
)

func main() {
	op := math.Calc("+")
	fmt.Println(op(10,20))
	op = math.Calc("*")
	fmt.Println(op(10,20))
}


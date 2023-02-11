package main

import (
	"fmt"
	"math/rand"

	"github.com/yeesik-kim/golang/packages/calculate"
	"github.com/yeesik-kim/golang/packages/printeo"
)

func main() {
	fmt.Println(rand.Int())

	// 모듈
	pac()
}

func pac() {
	printeo.PrintlnMyView("module test")
	fmt.Println(calculate.Min(5, 2))
}

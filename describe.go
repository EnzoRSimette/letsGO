package main

import (
	"fmt"
)

const version = "v1.0.1"

func describe(name string, age int, height float64, active bool) string {
	var b int
	b = 5
	var n string = name
	a := age
	var (
		h  float64 = height
		ac bool    = active
	)
	fmt.Println(b)
	return fmt.Sprintf("Hi %s, your age is %d, your height is %.2f. Are you active? %t. We are currently in version %s", n, a, h, ac, version)
}

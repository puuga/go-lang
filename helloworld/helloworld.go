package main

import (
	"fmt"
	"time"
	"math"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func main() {
	fmt.Println("Welcome to the playground!")
	fmt.Println("The time is", time.Now())
	fmt.Println("PI", math.Pi)
	fmt.Println("add(2, 3)", add(2, 3))
	a, b := swap("hello", "world")
	fmt.Println("swap('hello', 'world')", a, b)
	fmt.Println(split(17))

	var i int
	fmt.Println(i, c, python, java)
}

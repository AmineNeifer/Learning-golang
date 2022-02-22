package main

import "fmt"

func sayHi() {
	fmt.Println("hi")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func main() {
	var greeterFunction func()
	var mathFunction func(int, int) float64

	greeterFunction = sayHi
	mathFunction = divide

	greeterFunction()
	fmt.Println(mathFunction(2, 5))
}
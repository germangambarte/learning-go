package main

import "fmt"

func main() {
	var x int = 10
	var y float64 = 30.2
	var floatSum = float64(x) + y
	var intSum = x + int(y)
	fmt.Println(floatSum)
	fmt.Println(intSum)
}

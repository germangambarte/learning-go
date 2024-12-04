package main

import "fmt"

func main() {
	var x int = 10
	var b byte = 100
	var intSum = x + int(b)
	var byteSum = byte(x) + b
	fmt.Println(intSum)
	fmt.Println(byteSum)
}

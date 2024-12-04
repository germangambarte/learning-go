package main

import "fmt"

func exerciseOne() {
	fmt.Println("Exercise 1")
	var i int = 20
	var f float64 = float64(i)

	fmt.Printf("i = %d, f = %f\n", i, f)
}

func exerciseTwo() {
	fmt.Println("Exercise 2")

	const value = 20

	var i int = value
	var f float64 = value

	fmt.Printf("i = %d, f = %f\n", i, f)
}

func exerciseThree() {
	fmt.Println("Exercise 3")

	var (
		b      byte
		smallI int32
		bigI   uint64
	)
	b += 1
	smallI += 1
	bigI += 1

	fmt.Printf("b = %d\n", b)
	fmt.Printf("smallI = %d\n", smallI)
	fmt.Printf("bigI = %d\n", bigI)
}

func main() {
	exerciseOne()
	exerciseTwo()
	exerciseThree()
}

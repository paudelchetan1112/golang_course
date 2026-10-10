package main

import (
	"fmt"
	"math"
)

func main() {
	//variable declaration
	var a, b int = 10, 3
	var result int
	result = a + b
	fmt.Println("Addition:", result)

	result = a - b
	fmt.Println("Subtraction:", result)

	result = a * b
	fmt.Println("Multiplication:", result)

	result = a / b
	fmt.Println("Division:", result)
	result = a % b
	fmt.Println("Modulus:", result)

	//increment and decrement
	a++
	fmt.Println("Increment:", a)
	b--
	fmt.Println("Decrement:", b)

	const p float64 = 22.0 / 7

	fmt.Println("Value of pi:", p)

	//overflow and underflow
	//overflow with signed integer
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)

	maxInt++
	fmt.Println(maxInt)
	//overflow with unsigned integer
	var maxUint uint64 = 18446744073709551615 //max value for uint64
	fmt.Println(maxUint)

	maxUint++
	fmt.Println(maxUint)

	var smallfloat float64 = 1.0e-323
	fmt.Println(smallfloat)

	smallfloat = smallfloat / math.MaxFloat64
	fmt.Println(smallfloat)

}

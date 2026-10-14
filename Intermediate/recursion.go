package Intermediate

import "fmt"
func Recursion(){
println("Factorial of 5 is", factorial(5))
println("Factorial of 6 is", factorial(6))
fmt.Println("sum of digit:", sumOfDigit(12))
fmt.Println("sum of digit of 2458", sumOfDigit(2458))

}
func factorial(n int) int{
	if n==0{
		return 1 //Best case : factorial of 0 is 1
	}
	return n*factorial(n-1)  
	//n*(n-1)*(n-2)*factorial(n-3)..................factorial(0)
	

}
func sumOfDigit(n int) int{
	if n==0{
		return n
	}
	return (n%10)+sumOfDigit(n/10)
}
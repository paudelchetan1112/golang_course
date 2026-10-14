package main

import "fmt"

func main(){
	result:=ApplyOperation(10, 0, division)
	fmt.Println(result)

}
func addition(a float64, b float64) float64 {
	return a+b
}
func multiplication(a float64, b float64) float64{
	return a*b
}
func division(a float64, b float64)float64{
	if b!=0{
		return a/b
	}else{
		fmt.Println("Divison by Zero is not possible")
	}
	return 0
}
func subtraction(a float64, b float64)float64{
	return a-b
}
func ApplyOperation(a float64, b float64, operation func(float64, float64)float64)float64{
	return operation(a, b)
}
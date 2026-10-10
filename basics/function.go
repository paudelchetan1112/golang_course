package main

import "fmt"

func main(){
	//func <name> (parameter list) returnType{
	//code block
	//return value
	//}
	// sum:=add(1,2)
	// fmt.Println(sum)


	// greet:=func(){
	// 	fmt.Println("Hello Anomynous function")
	// }
	// greet()
	// operation:=add
	// fmt.Println(operation(50,100))
	result:=applyoperation(5,3, add)
	fmt.Println(result)

	newFunction:=createMultiplier(5)
	fmt.Println(newFunction(3))
	


}
func add(a, b int)int{
		return a+b
	}
	//function that takes a function as a argument
	func applyoperation(x int, y int, operation func(int , int)int)int{
		return operation(x,y)

	}
	//function that return a function
	func createMultiplier(factor int)func(int)int{
		return func(x int)int{
			return x*factor
		}
	}
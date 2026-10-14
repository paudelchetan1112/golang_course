package Intermediate

import "fmt"
func Closure(){
sequence:=adder()
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())
fmt.Println(sequence())

subtracter:=func() func(int)int{
	countdown:=99
	return func(x int)int{
		countdown-=x
		return countdown
	}
}()
fmt.Println(subtracter(1))
fmt.Println(subtracter(1))
fmt.Println(subtracter(1))
fmt.Println(subtracter(1))
fmt.Println(subtracter(1))

}
func adder() func()int{
	i:=0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("Added 1 to i")
		return i
	}
}

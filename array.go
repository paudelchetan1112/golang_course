package main

import "fmt"
func main(){
	//syntax
	//var arrayName [size]datatype

// 	var numbers[5] int
// 	fmt.Println(numbers)
// numbers[len(numbers)-1]=20
// fmt.Println(numbers)
// fruits:=[4]string{"Apple","Banana","Mango","Grapes"}
// fmt.Println("fruits:",fruits)
// originalArray:=[3]int{1,2,3}
// copiedArray:=originalArray
// copiedArray[0]=10
// fmt.Println("originalArray:",originalArray)
// fmt.Println("copiedArray:",copiedArray)
// 	for i:=0;i<len(numbers);i++{
// 		 fmt.Println("Element at index ", i, ":", numbers[i])

// 	}
// 	for _, value:=range numbers{ //underscore is used to store unused values
// 		fmt.Printf("Value:%d\n", value)

// 	}
// 	fmt.Println("length of the numbers array is: ", len(numbers))

// 	//comparing Arrays:
// 	array1:=[3]int{1,2,3}
// 	array2:=[3]int{1,2,3}
// 	if array1==array2{
// 		fmt.Println("Arrays are equal")
// 	}else{
// 		fmt.Println("Arrays are not equal")
// 	}
// 	var matrix [3][3] int =[3][3]int{
// 		{1,2,3},
// 		{4,5,6},
// 		{7,8,9},
// 	}
// 	fmt.Println("Matrix:",matrix)
// 	for i:=0;i<len(matrix);i++{
// 		for j:=0;j<len(matrix[i]);j++{
// 			fmt.Printf("%d ",matrix[i][j])
// 		}
// 		fmt.Println()	
// 	}
originalArray:=[3]int{1,2,3}
var copiedArray *[3]int=&originalArray
copiedArray[0]=10
fmt.Println("originalArray:",originalArray)
fmt.Println("copiedArray:",*copiedArray)
}
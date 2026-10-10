package main

import (
	"fmt"
	"slices"
)

func Slice() {
var numbers []int
numbers = append(numbers, 10)
numbers = append(numbers, 20)
fmt.Println(numbers)
var numbers1=[]int{1,2,3}
numbers2:=[]int{4,5,6}
fmt.Println(numbers1)
fmt.Println(numbers2)
// slice:=make([]int, 5)
a:=[5]int{1,2,3,4,5}
slice1:=a[1:4]
fmt.Println(slice1)
slice1=append(slice1, 6,7)
fmt.Println(slice1)
sliceCopy:=make([]int, len(slice1))
copy(sliceCopy, slice1)
fmt.Println(sliceCopy)

for i, v:=range slice1{
	fmt.Println(i,v)

}
fmt.Println("element at index 3 of slice1", slice1[3])
slice1[3]=100
fmt.Println("element at index 3 of slice1 after modification", slice1[3])
if slices.Equal(slice1, sliceCopy){
	fmt.Println("slices are equal")
} else {
	fmt.Println("slices are not equal")
}
//two dimensional slice
twoD:=make([][]int, 3)
for i:=0;i<3;i++{
	innerLen:=i+1
	twoD[i]=make([]int, innerLen)
	for j:=0;j<innerLen;j++{
		twoD[i][j]=i+j
	}
}
fmt.Println(twoD)					
	for j:=0;j<3;j++{


	}
	fmt.Println("The capacity of slice1 is:", cap(slice1))
	fmt.Println("The len of slice2 is", len(slice1))

}

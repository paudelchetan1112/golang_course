package main

import "fmt"
func main(){
//mon jan 2 15:04:05 MST 2006

layout:="2006-01-02T15:04:05Z07:00"
str:="2024-07-04T14:30:18Z"
t, err:=time.Parse(layout, str)
if err!=nil{
	fmt.Println("Error parsing time:", err)
	return 
}
fmt.Println(t)

}


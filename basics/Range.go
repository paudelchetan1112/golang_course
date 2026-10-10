package main

import (
	"fmt"
	
)

func Range(){
	message:="HelloWorld"
	for i, v:=range message{
		// fmt.Println(i, string(v))
		fmt.Printf("Index:%d, Rune:%c\n", i,v)



	}
}
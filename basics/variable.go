package main

import "fmt"
// middleName:="ABC"
var middleName="Chetan"

func Variables(){
	var age int //unitialized variable
	var name string="Chetan" //initialized variable
	var name1="john"
	count:=10
	lastName:="smith"
	var isActive bool
	var string2="hello"
	string3:="ram"
var num int
var string4 string
fmt.Println(age)

fmt.Println("Empty string", string4)
fmt.Println(num)
fmt.Println(name)
fmt.Println(name1)
fmt.Println(count)
fmt.Println(lastName)
fmt.Println(isActive)
fmt.Println(string2)
fmt.Println(string3)
middleName:="ABC"
fmt.Println(middleName)
//default values
//Numeric Type: 0
//Boolean type: false
//String Type: ""
//pointer, slices, maps, channel, functions, interface, structs: nil


//----- Scope
// fmt.Println(firstName)
}
func printName(){
	firstName:="micheal"
	fmt.Println(firstName)

}
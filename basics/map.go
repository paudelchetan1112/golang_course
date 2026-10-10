package main

import (
	"fmt"
	"maps"
)

func Map(){
	//var m map[keyType]valueType

	//mapVariable:=make(map[keyType]ValueType)
	//using Map literal
		// mapVariable:=map[keyType]ValueType{
		// 	key1:value1,
		// 	key2:value2,
		// }

		myMap:=make(map[string]int)
		fmt.Println(myMap)

		myMap["key1"]=9	
		myMap["Code"]=18
		fmt.Println(myMap["key1"])
delete(myMap, "key1")
		fmt.Println(myMap["key1"])
		// clear(myMap)
		value, ok:=myMap["key1"]
	if !ok{
		fmt.Println("Value is not in this key")
	}else{
		fmt.Println("Value:", value)
	}
myMap2:=map[string]int{
	"a":1, "b":2,
}
fmt.Println(myMap2)
if maps.Equal(myMap, myMap2){
	fmt.Println("Maps are equal")	
}
for k, v:=range myMap2{

	fmt.Println(k, v)
}
var myMap4 map[string]string
if myMap4==nil{
	fmt.Println("Map is initialized to nil value")
}else{
	fmt.Println("Map is not initialized to nil value. ")

}
// myMap4["key"]="value"

// fmt.Println(myMap4)

myMap4=make(map[string]string)
myMap4["key"]="value"
fmt.Println(myMap4)

myMap5:=make(map[string]map[string]string)
myMap5["map1"]=myMap4
fmt.Println(myMap5)



}
package main

import (
	"encoding/json"
	"fmt"
	"os"

	// "io"
	// "net/http"
	// "os"
)

type Person struct{
	FirstName string `json:"first_name"`
	Age int `json:"age"`

}
func main(){
	person:=Person{FirstName: "john", Age: 30}
// 	//marshalling
// 	jsonData, err:=json.Marshal(person)
// 	if err!=nil{
// 		fmt.Println("error marshalling", err)
// 		return
// 	}
// 	fmt.Printf("%s",jsonData)
// 	var data Person
// err=json.Unmarshal(jsonData,&data)
// if err!=nil{
// 	fmt.Println("error marshalling", err)
// 		return
// }
// fmt.Println(data)


err:=json.NewEncoder(os.Stdout).Encode(&person)
if err!=nil{
	fmt.Println(err)

}
// fmt.Println(person)

fmt.Printf("%T", person)


}
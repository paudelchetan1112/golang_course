package Intermediate

import (
	"fmt"
	"unicode/utf8"
)
func main(){
	message:="Hello\nGo!"
	raw_Message:=`Hello\nGo!`
	message1:="Hello, \tGo!"
	fmt.Println(message)
	fmt.Println(raw_Message)
	fmt.Println(message1)
	fmt.Println("Length of message variable is", len(message))

	fmt.Println("length of the raw message is:", len(raw_Message))
	fmt.Println("Length of the message is:", len(message))
	fmt.Println("The first character in message var is", string(message[0]))

	greeting:="Hello"
	name:="Alice"
	fmt.Println(greeting+name)
	// str1:="Apple"
	// str:="apple"

	// str2:="banana"
	for i, char:=range message{
fmt.Printf("Character at index %d is %c\n", i, char)


	}
	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))
	

}
 
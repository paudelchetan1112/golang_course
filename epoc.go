package main

import (
	"fmt"
	"time"
)
func main(){
//00:00:00 UTC on jan 1, 1970
now:=time.Now()
unixTime:=now.Unix()
fmt.Println("current Time :", unixTime)
t:=time.Unix(unixTime, 0)
fmt.Println(t)

}

package main

import (
	"fmt"
	"time"
)
func main(){
	fmt.Println(time.Now())
	//specific time
	specificTime:=time.Date(2024, time.July, 30, 12, 0,0,0,time.UTC) //mon jan 2 2006 15:04:05 MST 2006
	fmt.Println(specificTime)
	parsedDate,err:=time.Parse("2006-01-02","2020-05-01" ) //mon jan 2 2006 15:04:05 MST 2006 "this is reference value"
if err!=nil{
	fmt.Println(err)

}
fmt.Println(parsedDate)
parseTime1, _:=time.Parse("06-01-02", "20-05-01")
fmt.Println(parseTime1)
parseTime2,err:=time.Parse("2006-Jan-02", "2020-Feb-01")
if err!=nil{
	fmt.Println(err)

}
fmt.Println(parseTime2)

parsedTime3,_:=time.Parse("06-1-2 15-04", "20-5-1 18-03")
fmt.Println(parsedTime3)


//formatting time
t:=time.Now()
fmt.Printf("%T", t)
fmt.Println("Formatted Time:", t.Format("06-01-02 04-15-05"))
oneDayLater:=t.Add(time.Hour*24)
fmt.Println(oneDayLater)
fmt.Println(oneDayLater.Weekday())
fmt.Println("RoundTime:", t.Round(time.Hour))
loc,_:=time.LoadLocation("Asia/Kathmandu")
t=time.Date(2024, time.July, 8, 14, 16, 40, 00, time.UTC)
//convert this to the specific time zone
tLocal:=t.In(loc)
//perform rounding
roundedTime:=t.Round(time.Hour)
roundedTimeLocal:=roundedTime.In(loc)

fmt.Println("original Time (UTC):", t)
fmt.Println("orginal time (local):", tLocal)
fmt.Println("Rounded Time (UTC):", roundedTime)
fmt.Println("RoundedTime (local):", roundedTimeLocal) //round , round up the time 

fmt.Println("Truncated Time:",t.Truncate(time.Hour) ) //truncate, round down the time
loc,_=time.LoadLocation("America/New_York")
//convert time to location
tInNy:=time.Now().In(loc)
fmt.Println("new york time:", tInNy)
t1:=time.Date(2024, time.July, 4, 12, 0,0,0,time.UTC)
t2:=time.Date(2024, time.July, 4, 18, 0, 0, 0, time.UTC)
duration:=t2.Sub(t1)
fmt.Println(duration)
fmt.Printf("%T\n", duration)
fmt.Println(t2.After(t1))

}


package main

import (
	"fmt"
	"time"
)

func main(){
	curTime:=time.Now()
	fmt.Println(curTime)

	fmt.Println(curTime.Format("01-02-2006"))
	fmt.Println(curTime.Format("01/02/2006"))
	fmt.Println(curTime.Format("Feb 1,2006,Mon"))

	createdDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate)
	fmt.Println(createdDate.Format("01-02-2006"))
}
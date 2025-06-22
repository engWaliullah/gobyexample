package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2005, time.April, 10, 3, 7, 17, 8, time.UTC)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))

}

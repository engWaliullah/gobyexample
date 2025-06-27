package main

import (
	"fmt"
	"time"
)

func printNumber() {
	for i := 0; i < 5; i++ {

		time.Sleep(250 * time.Millisecond)

		fmt.Println(i)
	}
}

func printLetter() {
	for i := 'a'; i <= 'e'; i++ {

		time.Sleep(400 * time.Millisecond)

		fmt.Printf("%c\n", i)
	}
}

func main() {

	go printNumber()
	go printLetter()

	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Main func finished")

}

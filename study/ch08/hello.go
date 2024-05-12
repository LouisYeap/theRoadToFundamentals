package main

import (
	"fmt"
	"time"
)

var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		//交叉打印，通知
		<-number
		fmt.Println(i, i+1)
		i += 2
		letter <- true
	}
}

func printLetter() {
	i := 0
	str := "abcdefghijklmnopqrstuvwxyzABCDEF"
	for {
		//交叉打印，通知
		<-number
		if i >= len(str) {
			return
		}
		fmt.Println(str[i : i+2])
		i += 2
		number <- true
	}
}
func main() {

	go printNum()

	go printLetter()

	number <- true

	time.Sleep(time.Second * 100)
}

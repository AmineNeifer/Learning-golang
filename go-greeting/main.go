package main

import "fmt"


func greeting(myChannel chan string) {
	myChannel <- "hi"
}
func greeting1(myChannel chan string) {
	myChannel <- "hello"
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)
	go greeting1(myChannel)
	fmt.Println(<-myChannel)
	fmt.Println(<-myChannel)
}
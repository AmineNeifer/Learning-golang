package main

import (
	"fmt"
	"log"
	"liji/calendar/date"
)

func main() {
	date := date.Date{}
	err := date.SetYear(2019)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetDay(20)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date)
}

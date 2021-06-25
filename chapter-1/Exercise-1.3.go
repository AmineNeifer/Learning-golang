package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("with for loop")
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%fs elapsed\n", time.Since(start).Seconds())

	fmt.Println("---")
	fmt.Println("without for loop")
	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("%fs elapsed\n", time.Since(start).Seconds())
}

package main

import (
	"fmt"
	"os"
)

func main() {
	for ind, val := range os.Args[:] {
		fmt.Println(ind, ": "+val)
	}
}

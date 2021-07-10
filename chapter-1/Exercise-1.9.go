package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	var dst = os.Stdout
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")) {
			time.Sleep(1 * time.Second)
			fmt.Println(url + " is not a valid url!")
			fmt.Println("moving on to the next url...")
			time.Sleep(1 * time.Second)
			continue
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Status code: %s\n", resp.Status[:len(resp.Status)-3])
		time.Sleep(1 * time.Second)
		_, err = io.Copy(dst, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

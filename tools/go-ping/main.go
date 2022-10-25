package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		url := os.Args[1]
		start := time.Now()
		_, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		dur := time.Since(start)
		fmt.Printf("Request took %vms\n", dur.Milliseconds())
	}
}

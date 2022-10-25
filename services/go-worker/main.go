package main

import (
	"fmt"
	"time"
)

func main() {
	for true {
		fmt.Println("Working hard...")
		time.Sleep(2 * time.Second)
	}
}

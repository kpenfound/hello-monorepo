package main

import (
	"fmt"
	"time"
)

func main() {
	for true {
		fmt.Println("Serving content...")
		time.Sleep(2 * time.Second)
	}
}

package main

import (
	"fmt"
	"time"

	"github.com/gucio321/go-clear"
)

func main() {
	fmt.Println("Cleaning console in 3 seconds...")
	time.Sleep(3 * time.Second)
	clear.Clear()
}

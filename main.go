package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		defer fmt.Print(os.Args[1], ".\n")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Print("Good morning, ")
	case t.Hour() < 18:
		fmt.Print("Good afternoon, ")
	default:
		fmt.Print("Good evening, ")
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	s, sep := "", ""
	start := time.Now()
	for _, value := range os.Args[1:] {
		s += sep + value
		sep = " "
	}
	fmt.Println(s)
	forRange := time.Since(start).Seconds()

	fmt.Printf("For range: %.10fs\n\n", forRange)

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))

	join := time.Since(start).Seconds()

	fmt.Printf("Join: %.10fs\n\n", join)
}

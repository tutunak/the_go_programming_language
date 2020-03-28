package main

import (
	"time"
	"fmt"
)

func main() {
	var i int64
	i = 0
	for {
		fmt.Println(i)
		i += 1
		time.Sleep(time.Duration(i) * time.Second)
	}
}

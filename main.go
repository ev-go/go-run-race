package main

import (
	"fmt"
	"time"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, val := range values {
		go func() {
			fmt.Println(val)
		}()
	}
	time.Sleep(100 * time.Millisecond)
}

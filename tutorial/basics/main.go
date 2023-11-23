package main

import (
	"fmt"
)

func fibonacci() func() int {
	first, second := 1, 0
	return func() int {
		first, second = second, first+second
		return first
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

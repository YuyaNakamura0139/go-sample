package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// nilインターフェースは具体的な値も型も保持しないので、ランタイムエラーになる
// nakamura@nakamura 13 % go run main.go
// (<nil>, <nil>)
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x10885d9]

// goroutine 1 [running]:
// main.main()
//         /Users/nakamura/go-sample/tutorial/methods_and_interfaces/13/main.go:12 +0x19
// exit status 2

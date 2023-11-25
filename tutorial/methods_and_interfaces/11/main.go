package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	// T型の変数iを宣言
	// Iインターフェースを満たす任意の型の値を代入できる
	var i I

	// &T{"Hello"}はIインターフェースを満たすので代入可能
	i = &T{"Hello"}
	// (&{Hello}, *main.T)
	describe(i)
	i.M()

	// F(math.Pi)もIインターフェースを満たすので代入可能
	i = F(math.Pi)
	// (3.141592653589793, main.F)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

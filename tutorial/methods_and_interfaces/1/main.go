package main

import (
	"fmt"
	"math"
)

// 構造体の定義
type Vertex struct {
	X, Y float64
}

// Goにはクラスが無いが、型にメソッドを定義できる
// funcとAbsメソッド名の間にレシーバ引数をとる
// この場合、Absメソッドはvという名前のVertex型のレシーバを持つ
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

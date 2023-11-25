package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバを使用することで、Vertexの値を変更することができる
// 変数レシーバだとVertexのコピーを作成し、そのコピーのX,Yを変更することになるので、元の値には影響を与えない
// ポインタレシーバを使う方が多いらしい
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}

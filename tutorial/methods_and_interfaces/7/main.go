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

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	// メソッドが変数レシーバの場合、呼び出す時にポインタレシーバも使用できる
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	// pは初期化したVertexへのポインタを指す
	fmt.Println(p)
	// *pでポインタが指す値を取得できる
	fmt.Println(*p)
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

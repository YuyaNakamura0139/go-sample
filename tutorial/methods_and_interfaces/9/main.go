package main

import (
	"fmt"
	"math"
)

// インターフェース型
// メソッドの集まり
type Abser interface {
	Abs() float64
}

func main() {
	// Abserを変数aとして宣言
	// a.<func>でインターフェースで定義したメソッドを使用できる
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	// fのAbsメソッドは変数レシーバなので、aに値を入れても問題ない
	a = f
	// vのAbsメソッドはポインタレシーバなので、aにはvのポインタを入れる必要がある
	a = &v

	// Absメソッドは*Vertexの定義のため、VertexがAbserインターフェースを実装していないとエラーが出る
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

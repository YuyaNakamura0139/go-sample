package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// vはVertex型の値を保持
	v := Vertex{3, 4}
	// Scaleメソッドはポインタレシーバをとるので、自動的に&vとして解釈される
	v.Scale(2)
	ScaleFunc(&v, 10)

	// pはVertex型へのポインタを保持
	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}

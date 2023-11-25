package main

import (
	"fmt"
	"math"
)

// 他のパッケージに定義しているかたに対してレシーバを伴うメソッドを宣言できない
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

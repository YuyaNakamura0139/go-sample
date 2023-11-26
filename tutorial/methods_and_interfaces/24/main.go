package main

import (
	"fmt"
	"image"
)

// package image

// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }

func main() {
	// これは、左上角が(0,0)で右下角が(100,100)の新しいRGBA画像を作成します。
	// したがって、この画像は100x100ピクセルのサイズを持ちます。
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// これは、画像の境界を出力します。
	// 境界は、画像の左上と右下の座標を表す矩形です。
	// この場合、出力はimage.Rect(0, 0, 100, 100)となります。
	fmt.Println(m.Bounds())
	// これは、画像の(0,0)座標にあるピクセルの色をRGBA形式で出力します。
	// RGBAは、赤(Red)、緑(Green)、青(Blue)、アルファ(透明度)を表します。
	// 新しい画像は初期化されているため、すべてのピクセルの色は(0,0,0,0)となります。
	fmt.Println(m.At(0, 0).RGBA())
}

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	// capacityが8の空のリストを作成
	b := make([]byte, 8)
	for {
		// Readメソッドは、データ(r)を与えられたバイトスライス(b)へ入れ、入れたバイトのサイズとエラーを返す
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}

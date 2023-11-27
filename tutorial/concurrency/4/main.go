package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// 明示的にチャネルを閉じないと、main関数で受信を待ちづづけてしまいデッドロック状態になる
	close(c)
}

func main() {
	// サイズ10のバッファチャネルを作成
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// チャネルcから値を受信し続ける
	for i := range c {
		fmt.Println(i)
	}
}

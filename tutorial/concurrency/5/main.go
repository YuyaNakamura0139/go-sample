package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		// ゴルーチン内は順次実行(シーケンシャル)なので、1→2の順で実行される
		for i := 0; i < 10; i++ {
			// 1. cからの受信を待って出力する
			fmt.Println(<-c)
		}
		// 2. quitに0を送信
		quit <- 0
	}()
	fibonacci(c, quit)
}

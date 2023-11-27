package main

import "fmt"

func main() {
	// 引数の二つ目にバッファサイズを入れることで、チャネルをバッファとして使える
	ch := make(chan int, 2)
	// バッファが詰まった時は、チャネルへの送信をブロック
	// バッファが空の時は、チャネルの受信をブロックする
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch <- 3
	fmt.Println(<-ch)
}

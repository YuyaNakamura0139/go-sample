package main

import "fmt"

// ch <- v    // v をチャネル ch へ送信する
// v := <-ch  // ch から受信した変数を v へ割り当てる

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

// ゴルーチン間でデータを安全に交換する
func main() {
	// 1. 整数のスライスとchannelの生成
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	// 2. スライスの前後半をゴルーチンで計算
	// 非同期なので、すぐ次の処理が実行される
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// 3. チャネルcから二つの値を受信する
	// チャネルからの受信はブロッキング操作で、値が送信されるまで実行が停止する
	// よって、sumゴルーチンがそれぞれ計算を完了し、結果をチャネルに送信するまで待つことになる。
	x, y := <-c, <-c

	// 4. 最後に出力
	fmt.Println(x, y, x+y)
}

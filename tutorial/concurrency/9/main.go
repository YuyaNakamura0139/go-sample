package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is safe to concurrently
// sync.Mutex型のmuとmap[string]int型のvを持つ構造体SafeCounter構造体を定義
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
// mu.Lock()とmu.Unlock()で囲まれているため、複数のごルーチンから同時に呼び出されても
// 一度に一つのゴルーチンからしか実行されない
// カウンターの更新が正しく行われ、データの競合状態が防止される
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

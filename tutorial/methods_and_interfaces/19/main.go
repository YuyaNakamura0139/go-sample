// エラーの章
package main

import (
	"fmt"
	"time"
)

// Error用の構造体を定義
type MyError struct {
	When time.Time
	What string
}

// Error型にErrorメソッドを定義したことで組み込みのerrorインターフェースを満たす
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// ポインターがerror型として返される理由は、Go言語のインターフェースの動作によるものです。
// error型は、Error()メソッドを持つ任意の型を表すインターフェースです。
// このため、MyError型のポインターをerror型として返すことができます。
// Go言語では、インターフェース型の変数には、そのインターフェースを実装する具体的な型の値やポインターを代入することができます。
// したがって、MyError型のポインターはerror型の変数に代入することができ、error型として扱うことができます。
// この仕組みにより、run()関数はerror型を返すことができ、エラーが発生した場合にはMyError型のエラー情報を返すこことができます。

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// errはインターフェース型の変数
	// errは
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

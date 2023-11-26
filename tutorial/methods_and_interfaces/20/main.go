package main

import (
	"fmt"
)

// error用の新しい型を定義
type ErrNegativeSqrt float64

// ErrNegativeSqrt型にErrorメソッドを実装し、errorインターフェースを満たすようにする
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// error型を返すので、xが負の場合は、ErrNegativeSqrt型に定義したErrorメソッドが実行される。
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	// 適当にzを決めうちして、ニュートン法でxの平方根の近似値を求める
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	// 1.414213562373095 <nil>
	fmt.Println(Sqrt(2))
	// 0 cannot Sqrt negative number: -2
	fmt.Println(Sqrt(-2))
}

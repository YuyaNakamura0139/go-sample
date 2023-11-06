package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 1; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)

		count := i
		countString := "回目の結果"
		result := z
		fmt.Println(count, countString, result)

		if math.Sqrt(x) == z {
			fmt.Println("clear!")
			break
		}
	}

	return z
}

func main() {
	var num float64 = 2
	fmt.Println(Sqrt(num))
	fmt.Println(math.Sqrt(num))
}

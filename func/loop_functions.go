package main

/**
平方根の計算をしてみよう
z -= (z*z - x) / (2 * z)
上記式で、z^2がどれほどxに近づくかに応じて、zを調整できる
*/

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	i := 1
	for i < 6 {
		z -= (z*z - x) / (2 * z)
		i++
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

package main

import "fmt"

// golangで配列を実装してみる
func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// int型の6個の変数配列
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}

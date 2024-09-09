package main

import "fmt"

// 配列は固定長・スライスは可変長
// 型[]T
func main() {
	// int型の6個の変数配列
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// スライス：primesのindex1~3を取得
	var s []int = primes[1:4]
	fmt.Println(s)
	fmt.Println("=========================")
	sliceReference()
	fmt.Println("=========================")
	sliceLiterals()
	fmt.Println("=========================")
	sliceDefaults()
}

func sliceDefaults() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	// 下限を設定しない
	s = s[:2]
	fmt.Println(s)

	// 上限を設定しない
	s = s[1:]
	fmt.Println(s)
}

// スライスのリテラル(長さのないリテラルのようなもの)
func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		// literalの配列も作成可能
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

func sliceReference() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[1:2]
	b := names[1:3]
	fmt.Println(a, b)

	// スライスの要素を変更
	b[0] = "XXX"
	// 元の配列の対応する要素が同様に変更される
	fmt.Println(a, b)
	fmt.Println(names)
}

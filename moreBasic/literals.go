package main

import "fmt"

/**
 * フィールドの値を列挙することで、新しいstructの初期値の割り当てを示す
 */

type Vertex3 struct {
	X, Y int
}

var (
	// 通常の書き方
	v1 = Vertex3{1, 2}
	// Yは0になる
	v2 = Vertex3{X: 1}
	// X Yともに0になる
	v3 = Vertex3{}
	// 新しく割り当てられたstructのポインタ
	p = &Vertex3{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

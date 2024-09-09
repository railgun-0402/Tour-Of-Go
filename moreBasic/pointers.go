package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(p)  // pがさすメモリアドレス自体
	fmt.Println(*p) // pがさすメモリアドレスの値

	// *pはポインタの参照先の値を取得する
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

package main

import "fmt"

type Vertex2 struct {
	X int
	Y int
}

func main() {
	v := Vertex2{1, 2}
	// ポインタを通してアクセス
	p := &v
	// (*p).Xとも書ける
	p.X = 1e9
	fmt.Println(v)

}

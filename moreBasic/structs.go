package main

import "fmt"

// Vertex fieldの集まり
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// structsのXに値を代入
	v.X = 4
	fmt.Println(Vertex{v.X, 2})
}

package main

import (
	"fmt"
	"time"
)

func formatStringUsingFmt() {
	_ = fmt.Sprintf("%s,%s", "Hello", "world")
}
func formatStringUsingUnion() {
	_ = fmt.Sprintf("Hello" + "," + "world")
}

func main() {
	//fmt.Println("formatStringUsingUnion")
	//then := time.Now()
	//formatStringUsingUnion()
	//fmt.Println(time.Now().Sub(then).Nanoseconds())

	fmt.Println("formatStringUsingFmt")
	then := time.Now()
	formatStringUsingFmt()
	fmt.Println(time.Now().Sub(then).Nanoseconds())
}

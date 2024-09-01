package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on")

	// 実行されているOSの文字列で条件分岐する
	// Goはcaseの条件が一致次第、自動breakする
	switch os := runtime.GOOS; os {
	case "darwin": // MacOS
		fmt.Println("OS X.")
	case "linux": // linux
		fmt.Println("Linux")
	case "windows": // windows
		fmt.Println("Linux")
	default:
		fmt.Println("%s.\n", os)
	}
}

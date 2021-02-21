package main

// パッケージ、ライブラリ
import (
	"fmt"
	"os/user"
	"time"
)

// 最初に呼び出される特殊関数　init
func init() {
	fmt.Println("Init!")
}

func xxx() {
	fmt.Println("うふふ")
}

func main() {
	fmt.Println("Hello world!", time.Now())
	fmt.Println(user.Current())
}

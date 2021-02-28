package main

import (
	"fmt"
)

// varであれば外でも宣言可能
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	// 省略は関数の中のみ
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	// 型確認に
	fmt.Printf("%T\n", xf64)
}

func main() {
	fmt.Println(i, f64, s, t, f)
	foo()
}

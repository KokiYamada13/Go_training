package main

import "fmt"

func main() {
	n := make([]int, 3, 5)
	fmt.Printf("lan=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 0, 0)
	fmt.Printf("lan=%d cap=%d value=%v\n", len(n), cap(n), n)
	n = append(n, 1, 2, 3, 4, 5)
	fmt.Printf("lan=%d cap=%d value=%v\n", len(n), cap(n), n)

	a := make([]int, 3)
	fmt.Printf("lan=%d cap=%d value=%v\n", len(a), cap(a), a)

	// c = make([]int, 0, 5)
	// for i := 0; i < 5; i++ {
	// 	c = append(c, i)
	// 	fmt.Println(c)
	// }
	// fmt.Println(c)
}

package main

import (
	"math/rand"
)

func main() {
	// ランダムで6までの数字を返す
	n := rand.Intn(5) + 1
	println(n)
	if n == 6 {
		println("Lucky!")
	}
	if n == 5 {
		println("Good!")
	}
}

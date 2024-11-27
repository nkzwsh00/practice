package main

func main() {
	const count = 100
	for i := 1; i <= count; i++ {
		if i%2 == 0 {
			println(i, "is even")
		} else {
			println(i, "is odd")
		}
	}
}

package main

import (
	"fmt"
)

func Callatz(i int) int {
	n := 0
	for {
		fmt.Println(n)
		if i == 1 {
			break
		}
		if i%2 != 0 {
			i = (3*i + 1) / 2
		} else {
			i = i / 2
		}
		n++
	}
	return n
}

func main() {
	fmt.Println(Callatz(5))
}

package main

import "fmt"

func main() {
	var n int
	var check bool

	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		check = n%i == 0
		fmt.Println(i, check)
	}
}

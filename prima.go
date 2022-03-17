package main

import "fmt"

func main() {
	var n int
	var prime bool

	fmt.Scanln(&n)

	if n >= 1 {
		for i := 2; i <= n; i++ {
			if (n % i) == 0 {
				prime = true
			}
		}
		if prime {
			fmt.Println("prima")
		} else {
			fmt.Println("bukan prima")
		}
	}
}

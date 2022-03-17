package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)
	i := 0
	for i < n {
		if i%2 == 1 {
			fmt.Println(i)
		}
		i++
		if i == n {
			break
		}
	}
}

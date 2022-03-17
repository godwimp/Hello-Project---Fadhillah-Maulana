package main

import "fmt"

func main() {
	var jum int
	jum = 0
	for i := 1; i <= 1000; i++ {
		jum = i + jum
	}
	fmt.Println(jum)
}

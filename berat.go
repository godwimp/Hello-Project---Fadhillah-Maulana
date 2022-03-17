package main

import "fmt"

func main() {
	var n int
	var berat, rata, tot float64

	fmt.Scanln(&n)
	tot = 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&berat)
		tot = tot + berat
		rata = tot / float64(n)
	}
	fmt.Println(rata, "kg")
}

package main

import "fmt"

func main() {
	var tot, n0, n1, sel, rata float64
	var i int

	fmt.Scan(&n0, &n1)
	tot = 0
	i = 0

	for n0 != n1 {
		sel = n1 - n0
		i++
		tot = tot + sel
		n0 = n1
		fmt.Scan(&n1)
	}
	rata = tot / float64(i)
	fmt.Printf("%.3f", rata)
}

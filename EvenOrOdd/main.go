package main

import "fmt"

func main() {
	max := 10
	var n []int
	for i := 0; i <= max; i++ {
		n = append(n, i)
	}

	for _, check := range n {
		if check%2 == 0 {
			fmt.Printf("%d is Even\n", check)
		} else {
			fmt.Printf("%d is Odd\n", check)
		}
	}
}

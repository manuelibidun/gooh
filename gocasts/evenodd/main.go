package main

import "fmt"

func main() {
	numbers := [11]int{}

	for n := range numbers {
		if n%2 == 0 {
			fmt.Println(n, "is even")
		}
		if n%2 > 0 {
			fmt.Println(n, "is odd")
		}
	}
}

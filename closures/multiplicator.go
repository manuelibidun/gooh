package main

import "fmt"

func main() {
	baseTwo := makeMultiplicator(2)
	baseThree := makeMultiplicator(3)
	fmt.Println(baseTwo(4))
	fmt.Println(baseThree(5))
}

type multClosureType func(int) int

func makeMultiplicator(base int) multClosureType {
	return func(factor int) int {
		return base * factor
	}
}




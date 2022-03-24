package main

import "fmt"

func main() {
	sequence := makeSequence(2)
	
	fmt.Println(sequence())
	fmt.Println(sequence(1))
	
	for i := 1; i <= 5; i++ {
		fmt.Println(sequence(2))
	}
	
	fmt.Println(sequence(10))
}

type intSequenceType func(...int) int

func makeSequence(start int) intSequenceType {
	value := start
	return func(incr ...int) int {
		var incrementBy int = 1
		if len(incr) > 0 {
			incrementBy = incr[0]
		}
		value = value + incrementBy
		return value
	}
}


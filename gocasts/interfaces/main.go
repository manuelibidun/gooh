package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}
type bot interface {
	getGreeting() string
	//getVersion() float64
}

func (eb englishBot) getGreeting() string {
	//fmt.Println("Hello!")
	return "Hello"
}
func (sb spanishBot) getGreeting() string {
	//fmt.Println("Hello!")
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

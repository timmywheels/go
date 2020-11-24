package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func main() {
	english := englishBot{}
	spanish := spanishBot{}

	printGreeting(english)
	printGreeting(spanish)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// imagine a very custom impl
	// for getting the english greeting
	return "Hello, World!"
}

func (spanishBot) getGreeting() string {
	// imagine very custom impl
	// for getting spanish greeting
	return "Hola, Mundo!"
}



//func (eb englishBot) printGreeting() {
//	fmt.Println(eb.getGreeting())
//}
//
//func (sb spanishBot) printGreeting() {
//	fmt.Println(sb.getGreeting())
//}
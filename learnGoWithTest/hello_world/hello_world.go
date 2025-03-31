package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French" 
	 englishHelloPrefix = "Hello, "
	 spanishHelloPrefix = "Hola, "
	 frenchHelloPrefix = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return getPrefix(language) + name
}

func getPrefix(Language string)(prefix string){
	switch Language {
	case spanish: 
	prefix = spanishHelloPrefix
	case french: 
	prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("juthi", spanish))
}
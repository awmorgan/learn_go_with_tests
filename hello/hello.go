package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "
const emoji = "Emoji"
const emojiHelloPrefix = "👋, "

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	switch lang {
	case spanish:
		return spanishHelloPrefix + name
	case french:
		return frenchHelloPrefix + name
	case emoji:
		return emojiHelloPrefix + name
	default:
		return englishHelloPrefix + name
	}
}

func main() {
	fmt.Println(Hello("world", ""))
}

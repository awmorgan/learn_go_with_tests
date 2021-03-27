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
	prefix := greetingPrefix(lang)
	return prefix + name
}

func greetingPrefix(lang string) string {
	prefix := ""
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case emoji:
		prefix = emojiHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", ""))
}

package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Halo, "

func Hello(name string, language string) string {

	if name == "" {
		name = "world"
	}

	return getMultiLangHello(language) + name
}

func getMultiLangHello(language string) (prefix string) {

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}

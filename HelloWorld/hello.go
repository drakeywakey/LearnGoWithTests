package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const german = "German"

const helloPrefix = "Sup, "
const spanishPrefix = "Cenar, "
const frenchPrefix = "Ca va, "
const germanPrefix = "Guten tag, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name + "?"
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchPrefix
	case spanish:
		prefix = spanishPrefix
	case german:
		prefix = germanPrefix
	default:
		prefix = helloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Drake", ""))
}

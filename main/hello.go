package main

import "fmt"

const helloEnglishPrefix = "Hello, "
const helloHindiPrefix = "नमस्ते, "
const helloFrenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return getPrefix(language) + name
}

func getPrefix(language string) string {
	var prefix string

	switch language {
	case "French":
		prefix = helloFrenchPrefix
	case "Hindi":
		prefix = helloHindiPrefix
	default:
		prefix = helloEnglishPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("", ""))
}

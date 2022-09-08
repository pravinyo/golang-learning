package main

import "fmt"

const helloEnglishPrefix = "Hello, "
const helloHindiPrefix = "नमस्ते, "
const helloFrenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	var prefix string
	if name == "" {
		name = "World"
	}

	switch language {
	case "French":
		prefix = helloFrenchPrefix
	case "Hindi":
		prefix = helloHindiPrefix
	default:
		prefix = helloEnglishPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}

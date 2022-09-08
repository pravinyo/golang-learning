package main

import "fmt"

const helloEnglishPrefix = "Hello, "
const helloHindiPrefix = "नमस्ते, "
const helloFrenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	prefix := helloEnglishPrefix
	if name == "" {
		name = "World"
	}
	if language == "Hindi" {
		prefix = helloHindiPrefix
	}

	if language == "French" {
		prefix = helloFrenchPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}

package main

import "fmt"

const helloEnglishPrefix = "Hello, "
const helloHindiPrefix = "नमस्ते, "

func Hello(name string, language string) string {
	prefix := helloEnglishPrefix
	if name == "" {
		name = "World"
	}
	if language == "Hindi" {
		prefix = helloHindiPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}

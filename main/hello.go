package main

const helloEnglishPrefix = "Hello, "
const helloHindiPrefix = "नमस्ते, "
const helloFrenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return getPrefix(language) + name
}

func getPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = helloFrenchPrefix
	case "Hindi":
		prefix = helloHindiPrefix
	default:
		prefix = helloEnglishPrefix
	}
	return
}

//func main() {
//	fmt.Println(Hello("", ""))
//}

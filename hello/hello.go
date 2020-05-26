package hello

import (
	"fmt"
)

const (
	englishHelloPrefix string = "Hello "

	spanishHelloPrefix string = "Hola "

	frenchHelloPrefix string = "Bonjour "
)

func Hello(name string, language string) string {

	if name == "" {
		name = "World!"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(prefix string) string {

	switch prefix {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	//Default Prefix:
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func main() {

	fmt.Println(Hello("Mo", ""))
}

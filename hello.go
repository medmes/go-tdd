package hello

import "fmt"

const (
	englishHelloPrefix string = "Hello "

	spanishHelloPrefix string = "Hola "

	frenchHelloPrefix string = "Bonjour "
)

func Hello(name string, language string) string {
	//Default Prefix:
	prefix := englishHelloPrefix
	if name == "" {
		name = "World!"
	}

	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {

	fmt.Println(Hello("Mo", ""))
}

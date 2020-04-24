package hello

import "fmt"

const (
	englishHelloPrefix string = "Hello "

	spanishHelloPrefix string = "Hola "

	bonjourHelloPrefix string = "Bonjour "
)

func Hello(name string, language string) string {
	if "" == name {
		return englishHelloPrefix + "World!"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name
	}

	if language == "French" {
		return bonjourHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {

	fmt.Println(Hello("Mo", ""))
}

package hello

import "fmt"

const (
	englishHelloPrefix string = "Hello "

	spanishHelloPrefix string = "Hola "
)

func Hello(name string, language string) string {
	if "" == name {
		return englishHelloPrefix + "World!"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {

	fmt.Println(Hello("Mo", ""))
}

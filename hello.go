package hello

import "fmt"

const englishHelloPrefix string = "Hello "

func Hello(name string, language string) string {
	if "" == name {
		return englishHelloPrefix + "World!"
	}

	if language == "Spanish" {
		return "Hola " + name
	}
	return englishHelloPrefix + name
}

func main() {

	fmt.Println(Hello("Mo", ""))
}

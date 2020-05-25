package di

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer /* interface */, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Mo!")

}

func main() {
	fmt.Println("HTTP starting Listening on port: 8080")
	if err := http.ListenAndServe(":8080", http.HandlerFunc(MyGreetHandler)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

package di

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Poem struct {
	content []byte
	// let's consider Poems have to be stored somewhere
	//acmeStorageServices is a service that handle storage operation
	storage acmeStorageServices.PoemNotebook
}

func NewPoem() *Poem {
	return &Poem{
		storage: acmeStorageServices.NewPoemNotebook(),
	}
}

func (p *Poem) Load(title string) {
	p.content = p.storage.Load(title)
}

func (p *Poem) Save(title string) {
	p.storage.Save(title)
}

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

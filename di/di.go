package di

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// PoemStorage Interface
type PoemStorage interface {
	Type() string        // Return a string describing the storage type.
	Load(string) []byte  // Load a poem by name, and return byte table.
	Save(string, []byte) // Save a poem by name.
}

// Poem Struct
type Poem struct {
	content []byte
	// let's consider Poems have to be stored somewhere
	//acmeStorageServices is a service that handle storage operation
	storage PoemStorage
}

func NewPoem(ps PoemStorage) *Poem {
	return &Poem{
		storage: ps,
	}
}

func (p *Poem) Save(title string) {
	p.storage.Save(title, p.content)
}

func (p *Poem) Load(title string) {
	p.content = p.storage.Load(title)
}

func (p *Poem) String() string {
	return string(p.content)
}

// NoteBook Struct
type NoteBook struct {
	poems map[string][]byte
}

func NewNoteBook() *NoteBook {
	return &NoteBook{
		poems: map[string][]byte{},
	}
}

func (n *NoteBook) Save(name string, contents []byte) {
	n.poems[name] = contents
}

func (n *NoteBook) Load(name string) []byte {
	return n.poems[name]
}

func (n *NoteBook) Type() string {
	return "Notebook"
}

// Napkin Struct
type Napkin struct {
	poem []byte
}

func NewNapkin() *Napkin {
	return &Napkin{
		poem: []byte{},
	}
}

func (n *Napkin) Save(name string, contents []byte) {
	n.poem = contents
}

func (n *Napkin) Load(name string) []byte {
	return n.poem
}

func (n *Napkin) Type() string {
	return "Napkin"
}

// Main for Poem example:
func mainPoem() {
	notebook := NewNoteBook()
	napkin := NewNapkin()

	poem := NewPoem(notebook)
	poem.Save("My first poem")

	poem = NewPoem(notebook)
	poem.Load("My first poem")
	fmt.Println(poem)

	poem = NewPoem(napkin)
	poem.Save("My second poem")
	poem = NewPoem(napkin)
	poem.Load("My second poem")
	fmt.Println(poem)
}

// TDD - DI work:
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

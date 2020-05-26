package maps

import "errors"

var NotFoundError error = errors.New("not Found word")

// var m map[string]string // never initialize an empty map, it might generate a nil panic
// Initialize
// var dictionary = map[string]string{}
// OR
// var dictionary = make(map[string]string)
type Dictionary map[string]string

func (d Dictionary) Search(w string /* word */) (string, error) {
	f, ok := d[w]
	if !ok {
		return "", NotFoundError
	}
	return f, nil
}

func (d Dictionary) Add(w /* word */, de /* definition */ string) {
	d[w] = de
}

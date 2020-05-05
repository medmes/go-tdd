package maps

import "errors"

var notFoundError error = errors.New("not Found word")

type Dictionary map[string]string

func (d Dictionary) Search(w string /* word */) (string, error) {
	f, ok := d[w]
	if !ok {
		return "", notFoundError
	}
	return f, nil
}

package main

import "errors"

func main() {

}

type Dict map[string]string

var ErrNotFound = errors.New("could not find the word")

func (d Dict) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dict) Add(word string, def string) {
	d[word] = def
}

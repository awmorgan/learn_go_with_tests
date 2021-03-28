package main

func main() {

}

type Dict map[string]string

type DictErr string

const (
	ErrNotFound   = DictErr("could not find the word")
	ErrWordExists = DictErr("word already exists")
)

func (e DictErr) Error() string {
	return string(e)
}

func (d Dict) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dict) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = def
		return nil
	case nil:
		return ErrWordExists
	default:
		return err
	}
}

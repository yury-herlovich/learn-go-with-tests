package dictionary

import "errors"

var (
	ErrNotFound        = errors.New("could not find the word you were looking for")
	ErrKeyExists       = errors.New("key already exists in dictionary")
	ErrKeyDoesNotExist = errors.New("key does not exist in dictionary")
)

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	word, ok := d[key]

	if !ok {
		return "", ErrNotFound
	}

	return word, nil
}

func (d Dictionary) Add(key, str string) error {
	_, ok := d[key]
	if ok {
		return ErrKeyExists
	}

	d[key] = str

	return nil
}

func (d Dictionary) Update(key, str string) error {
	_, err := d.Search(key)

	if err != nil {
		return ErrKeyDoesNotExist
	}

	d[key] = str
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	if err != nil {
		return ErrKeyDoesNotExist
	}

	delete(d, key)
	return nil
}

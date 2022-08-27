package main

type Dictionary map[string]string

const (
	ErrMissingKey   = DictionaryErr("could not find the word you were looking for")
	ErrExists       = DictionaryErr("key already exists")
	ErrDoesNotExist = DictionaryErr("the map does not contain the supplied key")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrMissingKey
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, ok := d[key]
	if ok {
		return ErrExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	if err == ErrMissingKey {
		return ErrDoesNotExist
	}

	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	if err == ErrMissingKey {
		return ErrMissingKey
	}

	delete(d, key)
	return nil
}

package main

import "errors"

type Dictionary map[string]string

var ErrWordNotFound = errors.New("word not known")

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key string, definition string) {
	d[key] = definition
}

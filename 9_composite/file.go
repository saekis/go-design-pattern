package main

import "github.com/pkg/errors"

type file struct {
	name string
	size int64
}

func NewFile(name string, size int64) Entry {
	return &file{
		name: name,
		size: size,
	}
}

func (f *file) GetName() string {
	return f.name
}

func (f *file) GetSize() int64 {
	return f.size
}

func (f *file) Add(Entry) error {
	return errors.New("This is not directory")
}

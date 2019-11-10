package main

type Entry interface {
	GetName() string
	GetSize() int64
	Add(Entry) error
}

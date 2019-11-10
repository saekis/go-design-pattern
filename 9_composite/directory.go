package main

type directory struct {
	name    string
	size    int64
	entries []Entry
}

func NewDirectory(name string, size int64) Entry {
	return &directory{
		name: name,
		size: size,
	}
}

func (d *directory) GetName() string {
	return d.name
}

func (d *directory) GetSize() int64 {
	return d.size
}

func (d *directory) Add(entry Entry) error {
	d.entries = append(d.entries, entry)
	return nil
}

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()

	for _, v := range walk(getPath()) {
		fmt.Printf("%#v", v)
	}
	return 0
}

func walk(path string) []Entry {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var entries []Entry
	for _, file := range files {
		var entry Entry
		if file.IsDir() {
			entry := NewDirectory(file.Name(), file.Size())
			entries = append(entries, walk(filepath.Join(path, file.Name()))...)
			for _, entryInDir := range entries {
				err := entry.Add(entryInDir)
				if err != nil {
					panic("panic!")
				}
			}
		} else {
			entry = NewFile(file.Name(), file.Size())
		}

		entries = append(entries, entry)
	}
	return entries
}

func getPath() string {
	var path string
	if args := flag.Args(); len(args) > 0 {
		path = args[0]
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
	} else {
		path = "./"
	}
	return path
}

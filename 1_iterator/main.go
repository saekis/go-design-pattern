package main

import (
	"fmt"

	"github.com/saekis/go-design-pattern/1_iterator/bookshelf"
)

func main() {
	bs := bookshelf.NewBookShelf(3)
	bs.AppendBook(bookshelf.NewBook("リーダブルコード"))
	bs.AppendBook(bookshelf.NewBook("達人プログラマ"))
	bs.AppendBook(bookshelf.NewBook("Team Geek"))
	it := bs.Iterator()
	for it.HasNext() {
		book := it.Next()
		fmt.Println(book.GetName())
	}
}

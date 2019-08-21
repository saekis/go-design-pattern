package main

import (
	"fmt"

	"github.com/saekis/go-design-pattern/7_builder/director"

	"github.com/saekis/go-design-pattern/7_builder/builder"
)

func main() {
	// Build text
	b := builder.NewTextBuilder()
	d := director.NewDocumentDirector(b)
	result := d.Build()
	fmt.Println(result)

	// Build html
	b = builder.NewHtmlBuilder()
	d = director.NewDocumentDirector(b)
	result = d.Build()
	fmt.Println(result)
}

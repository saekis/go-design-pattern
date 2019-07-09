package main

import (
	"fmt"

	"github.com/saekis/go-design-pattern/3_template_method/printer"
)

func main() {
	ad := printer.AbstractPrinter{}
	fmt.Println(ad.Display(printer.NewCharDisplay("char_display")))
	fmt.Println(ad.Display(printer.NewStringDisplay("string_display")))
}

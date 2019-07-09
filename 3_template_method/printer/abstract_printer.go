package printer

type AbstractPrinter struct {
}

func (*AbstractPrinter) Display(printer Printer) string {
	result := printer.Open()
	result += printer.Print()
	result += printer.Close()
	return result
}

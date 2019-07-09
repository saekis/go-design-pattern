package printer

type charDisplay struct {
	char string
}

func NewCharDisplay(char string) Printer {
	return &charDisplay{char: char}
}

func (*charDisplay) Open() string {
	return "<<"
}

func (cd *charDisplay) Print() string {
	return cd.char
}

func (*charDisplay) Close() string {
	return ">>"
}

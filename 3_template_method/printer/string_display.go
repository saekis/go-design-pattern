package printer

type stringDisplay struct {
	char string
}

func NewStringDisplay(char string) Printer {
	return &stringDisplay{char: char}
}

func (*stringDisplay) Open() string {
	return "+--+"
}

func (cd *stringDisplay) Print() string {
	return "|" + cd.char + "|"
}

func (*stringDisplay) Close() string {
	return "+--+"
}

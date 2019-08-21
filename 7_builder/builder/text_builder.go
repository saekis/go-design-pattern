package builder

import (
	"bytes"
	"fmt"
)

type textBuilderImple struct {
	buf bytes.Buffer
}

func NewTextBuilder() Builder {
	return &textBuilderImple{buf: bytes.Buffer{}}
}

func (b *textBuilderImple) MakeTitle(title string) {
	b.buf.WriteString(fmt.Sprintf("「%s」\n\n", title))
}

func (b *textBuilderImple) MakeString(str string) {
	b.buf.WriteString(fmt.Sprintf("%s\n", str))
}

func (b *textBuilderImple) MakeItems(items []string) {
	for _, item := range items {
		b.buf.WriteString(fmt.Sprintf("・%s\n", item))
	}
}

func (b *textBuilderImple) Close() {
	// Do nothing
}

func (b *textBuilderImple) GetResult() string {
	return b.buf.String()
}

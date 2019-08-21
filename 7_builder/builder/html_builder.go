package builder

import (
	"bytes"
	"fmt"
)

type htmlBuilderImple struct {
	buf bytes.Buffer
}

func NewHtmlBuilder() Builder {
	return &htmlBuilderImple{buf: bytes.Buffer{}}
}

func (b *htmlBuilderImple) MakeTitle(title string) {
	b.buf.WriteString(fmt.Sprintf("<html><head><title>%s</title></head><body>", title))
	b.buf.WriteString(fmt.Sprintf("<h1>%s</h1>", title))
}

func (b *htmlBuilderImple) MakeString(str string) {
	b.buf.WriteString(fmt.Sprintf("<p>%s</p>", str))
}

func (b *htmlBuilderImple) MakeItems(items []string) {
	b.buf.WriteString("<ul>")
	for _, item := range items {
		b.buf.WriteString(fmt.Sprintf("<li>%s</li>", item))
	}
	b.buf.WriteString("</ul>")
}

func (b *htmlBuilderImple) Close() {
	b.buf.WriteString("</body></html>")
}

func (b *htmlBuilderImple) GetResult() string {
	return b.buf.String()
}

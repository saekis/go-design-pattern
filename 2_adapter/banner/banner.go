package banner

import "fmt"

type Banner struct {
	Str string
}

func (b *Banner) WithParen() string {
	return fmt.Sprintf("(%s)", b.Str)
}

func (b *Banner) WithAster() string {
	return fmt.Sprintf("*%s*", b.Str)
}

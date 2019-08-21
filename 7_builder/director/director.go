package director

import "github.com/saekis/go-design-pattern/7_builder/builder"

type directorImple struct {
	builder builder.Builder
}

type Director interface {
	Build() string
}

func NewDirector(builder builder.Builder) Director {
	return &directorImple{builder: builder}
}

func (director *directorImple) Build() string {
	director.builder.MakeTitle("title")
	director.builder.MakeString("string")
	director.builder.MakeItems([]string{
		"item1",
		"item2",
	})
	director.builder.Close()
	return director.builder.GetResult()
}

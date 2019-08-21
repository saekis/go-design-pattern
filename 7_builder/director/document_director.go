package director

import "github.com/saekis/go-design-pattern/7_builder/builder"

type documentDirector struct {
	builder builder.Builder
}

func NewDocumentDirector(builder builder.Builder) Director {
	return &documentDirector{builder: builder}
}

func (director *documentDirector) Build() string {
	director.builder.MakeTitle("title")
	director.builder.MakeString("string")
	director.builder.MakeItems([]string{
		"item1",
		"item2",
	})
	director.builder.Close()
	return director.builder.GetResult()
}

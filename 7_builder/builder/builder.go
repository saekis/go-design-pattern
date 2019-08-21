package builder

type Builder interface {
	MakeTitle(string)
	MakeString(string)
	MakeItems([]string)
	Close()
	GetResult() string
}

package printer

type Printer interface {
	Open() string
	Print() string
	Close() string
}

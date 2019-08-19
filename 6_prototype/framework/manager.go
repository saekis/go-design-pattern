package framework

type Manager interface {
	Register(string, Product)
	Create(string) Product
}

type managerImple struct {
	showcase map[string]Product
}

func NewManager() Manager {
	return &managerImple{showcase: map[string]Product{}}
}

func (m *managerImple) Register(name string, proto Product) {
	m.showcase[name] = proto
}

func (m *managerImple) Create(protoname string) Product {
	if p, ok := m.showcase[protoname]; ok {
		return p.CreateClone()
	}

	return NewNullProduct()
}

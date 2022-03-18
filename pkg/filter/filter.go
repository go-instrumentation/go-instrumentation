package filter

type Filter interface {
	Allow(targetObject Object) (allow bool)
	GetName() string
}

type Base struct {
	Name string
}

func (b Base) GetName() string {
	return b.Name
}

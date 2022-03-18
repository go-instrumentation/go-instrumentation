package filter

type Filter interface {
	Allow(targetObject Object) (allow bool)
	String() string
}

type Base struct {
	Name string
}

func (b Base) String() string {
	return b.Name
}

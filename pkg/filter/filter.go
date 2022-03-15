package filter

type Filter interface {
	Allow(targetObject Object) (allow bool)
}
